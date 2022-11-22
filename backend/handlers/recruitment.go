package handlers

import (
	"backend/models"
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/randallmlough/pgxscan"
	"log"
)

func GetRecruitmentData(ctx *fiber.Ctx) error {
	highKpm, err := getHighKPMPlayers()
	if err != nil {
		log.Panic(err)
	}

	highKillGames, err := getHighKillGames()
	if err != nil {
		log.Panic(err)
	}

	return ctx.Status(200).JSON(map[string]interface{}{"highKillGames": highKillGames, "highKPMGamers": highKpm})
}

func getHighKPMPlayers() ([]models.KPMStats, error) {
	rows, _ := models.DbPool.Query(context.Background(),
		"select steam_id_64,"+
			"       MODE() WITHIN GROUP (ORDER BY name) as used_name,"+
			"       avg(kills)                          as kills,"+
			"       avg(deaths)                         as deaths,"+
			"       avg(kills_per_minute)               as kills_per_minute,"+
			"       avg(deaths_per_minute)              as deaths_per_minute,"+
			"       avg(kill_death_ratio)               as kill_death_ratio "+
			"from player_stats"+
			"         join steam_id_64 s on player_stats.playersteamid_id = s.id"+
			"         join map_history mh on player_stats.map_id = mh.id"+
			"         join (select map_id, count(player_stats.playersteamid_id) as player_count"+
			"               from player_stats"+
			"               group by player_stats.map_id) as player_count on player_count.map_id = player_stats.map_id "+
			"where kills > 0"+
			"  and player_count > 80"+
			"  and mh.creation_time > current_date - interval '7 days' "+
			"  and mh.server_number IN (1,2) "+
			"group by steam_id_64 "+
			"having count(steam_id_64) > 5 "+
			"order by kills_per_minute desc "+
			"limit 50;")

	var kpmStats []models.KPMStats
	err := pgxscan.NewScanner(rows).Scan(&kpmStats)

	return kpmStats, err
}

func getHighKillGames() ([]models.HighKillData, error) {
	rows, _ := models.DbPool.Query(context.Background(),
		"select steam_id_64,"+
			"       MODE() WITHIN GROUP (ORDER BY name) as used_name,"+
			"       max(kills)                          as max_kill_game "+
			"from player_stats "+
			"         join steam_id_64 s on player_stats.playersteamid_id = s.id"+
			"         join map_history mh on player_stats.map_id = mh.id"+
			"         join (select map_id, count(player_stats.playersteamid_id) as player_count"+
			"               from player_stats"+
			"               group by player_stats.map_id) as player_count on player_count.map_id = player_stats.map_id "+
			"where kills > 0"+
			"  and player_count > 80"+
			"  and mh.creation_time > current_date - interval '7 days'"+
			"  and mh.server_number IN (1,2) "+
			"group by steam_id_64 "+
			"order by max_kill_game desc "+
			"limit 100;")

	var gameData []models.HighKillData
	err := pgxscan.NewScanner(rows).Scan(&gameData)

	return gameData, err
}
