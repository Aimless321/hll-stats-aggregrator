package handlers

import (
	"backend/models"
	"context"
	"github.com/Philipp15b/go-steamapi"
	"github.com/gofiber/fiber/v2"
	"github.com/gookit/config/v2"
	"github.com/jackc/pgx/v4"
	"github.com/randallmlough/pgxscan"
	"log"
	"strconv"
)

func GetData(ctx *fiber.Ctx) error {
	ids := make([]uint64, 1)
	steamId, err := strconv.ParseUint(ctx.Params("steamid"), 10, 64)
	if err != nil {
		log.Println("Invalid steamid", ctx.Params("steamid"))
		ctx.SendStatus(400)
	}

	ids = append(ids, steamId)
	playerSummaries, err := steamapi.GetPlayerSummaries(ids, config.String("steamApiKey"))
	if err != nil {
		log.Panic("Cannot get steam info")
	}

	rows, _ := models.DbPool.Query(context.Background(),
		"select mh.creation_time as date,"+
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
			"from player_stats "+
			"         join steam_id_64 s on player_stats.playersteamid_id = s.id "+
			"         join map_history mh on player_stats.map_id = mh.id "+
			"         join (select map_id, count(player_stats.playersteamid_id) as player_count "+
			"               from player_stats "+
			"               group by player_stats.map_id) as player_count on player_count.map_id = player_stats.map_id "+
			"where s.steam_id_64 = $1"+
			"  and kills > 0 "+
			"  and player_count > 80 "+
			"order by mh.creation_time desc "+
			"limit 30", ctx.Params("steamid"))

	var stats []models.GameStats
	if err := pgxscan.NewScanner(rows).Scan(&stats); err != nil {
		switch err {
		case pgx.ErrNoRows:
			return ctx.SendStatus(404)
		default:
			log.Panic(err)
		}
	}

	rows, _ = models.DbPool.Query(context.Background(),
		"select avg(stats.kills)             as kills, "+
			"       avg(stats.deaths)            as deaths, "+
			"       avg(stats.kills_per_minute)  as kills_per_minute, "+
			"       avg(stats.deaths_per_minute) as deaths_per_minute, "+
			"       avg(stats.kill_death_ratio)  as kill_death_ratio "+
			"from (select kills, "+
			"             deaths, "+
			"             kills_per_minute, "+
			"             deaths_per_minute, "+
			"             kill_death_ratio "+
			"      from player_stats "+
			"               join steam_id_64 s on player_stats.playersteamid_id = s.id "+
			"               join map_history mh on player_stats.map_id = mh.id "+
			"               join (select map_id, count(player_stats.playersteamid_id) as player_count "+
			"                     from player_stats "+
			"                     group by player_stats.map_id) as player_count on player_count.map_id = player_stats.map_id "+
			"      where s.steam_id_64 = $1 "+
			"        and kills > 0 "+
			"        and player_count > 80"+
			"      order by mh.creation_time desc "+
			"      limit 30) stats;", ctx.Params("steamid"))
	var avgStats models.AvgStats
	if err := pgxscan.NewScanner(rows).Scan(&avgStats); err != nil {
		log.Panic(err)
	}

	return ctx.Status(200).JSON(models.PlayerStats{
		SteamInfo: playerSummaries[0],
		Avg:       avgStats,
		LastGames: stats,
	})
}
