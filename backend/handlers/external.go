package handlers

import (
	"backend/models"
	"context"
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/randallmlough/pgxscan"
	"log"
	"regexp"
	"strconv"
)

func CreateTables() {
	_, err := models.DbPool.Exec(context.Background(), "CREATE TABLE IF NOT EXISTS circle_external_events ("+
		"id SERIAL PRIMARY KEY, "+
		"map_id integer REFERENCES map_history(id), "+
		"name TEXT, "+
		"date TIMESTAMP, "+
		"squads JSONB, "+
		"url TEXT)")
	if err != nil {
		log.Fatal(err)
	}
}

func ImportExternal(ctx *fiber.Ctx) error {
	params := new(struct {
		StatsURL string `json:"statsURL"`
		Name     string `json:"name"`
	})

	if err := ctx.BodyParser(params); err != nil {
		return err
	}

	re := regexp.MustCompile(`(https?://.*)/#/gamescoreboard/(\d+)`)
	matches := re.FindStringSubmatch(params.StatsURL)
	if len(matches) < 3 {
		return ctx.SendStatus(400)
	}

	// Call API for stats
	a := fiber.AcquireAgent()
	a.UserAgent("YourMomIsAwesome")
	req := a.Request()
	req.SetRequestURI(fmt.Sprintf("%v/api/get_map_scoreboard?map_id=%v", matches[1], matches[2]))
	if err := a.Parse(); err != nil {
		log.Panic(err)
	}
	stats := models.RconScoreboardResult{}
	var body []byte
	var errs []error
	if _, body, errs = a.Bytes(); len(errs) > 0 {
		log.Panic(errs)
	}
	err := json.Unmarshal(body, &stats)
	if err != nil {
		return err
	}

	var mapId int
	err = models.DbPool.QueryRow(context.Background(),
		"INSERT INTO map_history (creation_time, start, \"end\", server_number, map_name) "+
			"VALUES (now(), $1, $2, 100, $3) RETURNING id",
		stats.Result.Start.Time, stats.Result.End.Time, stats.Result.MapName).Scan(&mapId)
	if err != nil {
		log.Panic(err)
	}

	tx, err := models.DbPool.Begin(context.Background())
	for _, player := range stats.Result.PlayerStats {
		var playerId int
		err = models.DbPool.QueryRow(context.Background(), "INSERT INTO steam_id_64 (steam_id_64, created) "+
			"VALUES ($1, now()) "+
			"ON CONFLICT(steam_id_64) DO UPDATE SET steam_id_64 = excluded.steam_id_64 RETURNING id", strconv.Itoa(int(player.SteamInfo.Profile.SteamID))).Scan(&playerId)
		if err != nil {
			log.Panic(err)
		}

		_, err = tx.Exec(
			context.Background(),
			"INSERT INTO player_names (playersteamid_id, name, created, last_seen) VALUES ($1, $2, now(), now()) ON CONFLICT DO NOTHING",
			playerId,
			player.Player,
		)
		if err != nil {
			log.Panic(err)
		}

		_, err = tx.Exec(context.Background(), "INSERT INTO player_stats ("+
			"playersteamid_id, "+
			"map_id, "+
			"kills, "+
			"kills_streak, "+
			"deaths, "+
			"deaths_without_kill_streak, "+
			"teamkills, "+
			"teamkills_streak, "+
			"deaths_by_tk, "+
			"deaths_by_tk_streak, "+
			"nb_vote_started, "+
			"nb_voted_yes, "+
			"nb_voted_no, "+
			"time_seconds, "+
			"kills_per_minute, "+
			"deaths_per_minute, "+
			"kill_death_ratio, "+
			"longest_life_secs, "+
			"shortest_life_secs, "+
			"death_by, "+
			"most_killed, "+
			"name, "+
			"weapons) "+
			"VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23)",
			playerId, mapId, player.Kills, player.KillStreak, player.Deaths, player.DeathStreak, player.TeamKills,
			player.TeamKillStreak, player.TKDeaths, player.TKDeathStreak, player.VoteStarted, player.VotedNo, player.VotedYes,
			player.TimeSeconds, player.KPM, player.DPM, player.KDRatio, player.LongestLife, player.ShortestLife, player.KilledBy,
			player.MostKilled, player.Player, player.Weapons)
		if err != nil {
			log.Panic(err)
		}
	}
	err = tx.Commit(context.Background())
	if err != nil {
		log.Panic(err)
	}

	var eventId int
	err = models.DbPool.QueryRow(
		context.Background(),
		"INSERT INTO circle_external_events (map_id, name, url, date, squads) "+
			"VALUES ($1, $2, $3, now(), '[]') RETURNING id",
		mapId,
		params.Name,
		params.StatsURL,
	).Scan(&eventId)
	if err != nil {
		log.Panic(err)
	}

	return ctx.Status(200).JSON(struct {
		EventId int `json:"eventId"`
	}{eventId})
}

func GetAllExternals(ctx *fiber.Ctx) error {
	rows, _ := models.DbPool.Query(context.Background(),
		"SELECT * FROM circle_external_events LIMIT 50")

	var games []models.ExternalGame
	err := pgxscan.NewScanner(rows).Scan(&games)
	if err != nil {
		log.Panic(err)
	}

	return ctx.Status(200).JSON(games)
}

func GetExternal(ctx *fiber.Ctx) error {
	rows, _ := models.DbPool.Query(context.Background(),
		"SELECT * FROM circle_external_events WHERE id = $1", ctx.Params("gameid"))

	var game models.ExternalGame
	err := pgxscan.NewScanner(rows).Scan(&game)
	if err != nil {
		log.Panic(err)
	}

	rows, _ = models.DbPool.Query(
		context.Background(),
		"select s.steam_id_64 as steam_id,"+
			"       pn.name,"+
			"       kills,"+
			"       kills_streak,"+
			"       deaths,"+
			"       deaths_without_kill_streak,"+
			"       time_seconds,"+
			"       kills_per_minute,"+
			"       deaths_per_minute,"+
			"       kill_death_ratio,"+
			"       longest_life_secs,"+
			"       shortest_life_secs,"+
			"       weapons,"+
			"       map_name,"+
			"       server_number "+
			"from player_stats"+
			"         join steam_id_64 s on player_stats.playersteamid_id = s.id"+
			"         join map_history mh on player_stats.map_id = mh.id"+
			"         join player_names pn on s.id = pn.playersteamid_id"+
			"         LEFT JOIN player_names pn2 ON (s.id = pn2.playersteamid_id AND (pn.last_seen < pn2.last_seen or"+
			"                                                                        (pn.last_seen = pn2.last_seen and pn.id < pn2.id))) "+
			"where player_stats.map_id = $1 "+
			"and pn2.id is null",
		game.MapId,
	)

	var stats []models.ExternalStats
	err = pgxscan.NewScanner(rows).Scan(&stats)
	if err != nil {
		log.Panic(err)
	}

	return ctx.Status(200).JSON(struct {
		Game    models.ExternalGame    `json:"game"`
		Players []models.ExternalStats `json:"players"`
	}{
		Game:    game,
		Players: stats,
	})
}

func UpdateExternal(ctx *fiber.Ctx) error {
	model := new(models.ExternalGame)

	if err := ctx.BodyParser(model); err != nil {
		return err
	}

	_, err := models.DbPool.Exec(
		context.Background(),
		"UPDATE circle_external_events SET squads = $1 WHERE id = $2",
		model.Squads,
		model.Id,
	)
	if err != nil {
		log.Panic(err)
	}

	return ctx.SendStatus(201)
}
