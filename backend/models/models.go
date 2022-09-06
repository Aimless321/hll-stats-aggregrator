package models

import (
	"github.com/Philipp15b/go-steamapi"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/jackc/pgx/v4/pgxpool"
	"time"
)

type DiscordUserResponse struct {
	Id            string `json:"id"`
	Username      string `json:"username"`
	Discriminator string `json:"discriminator"`
	Avatar        string `json:"avatar"`
	Verified      bool   `json:"verified"`
	Email         string `json:"email"`
	Flags         int    `json:"flags"`
	Banner        string `json:"banner"`
	AccentColor   int    `json:"accent_color"`
	PremiumType   int    `json:"premium_type"`
	PublicFlags   int    `json:"public_flags"`
}

type AvgStats struct {
	Kills           float64 `db:"kills" json:"kills"`
	Deaths          float64 `db:"deaths" json:"deaths"`
	KillsPerMinute  float64 `db:"kills_per_minute" json:"killsPerMinute"`
	DeathsPerMinute float64 `db:"deaths_per_minute" json:"deathsPerMinute"`
	KDRatio         float64 `db:"kill_death_ratio" json:"KDRatio"`
}

type GameStats struct {
	Date                time.Time      `db:"date" json:"date"`
	GameId              int            `db:"map_id" json:"gameId"`
	Kills               int            `db:"kills" json:"kills"`
	KillStreak          int            `db:"kills_streak" json:"killStreak"`
	Deaths              int            `db:"deaths" json:"deaths"`
	DeathStreak         int            `db:"deaths_without_kill_streak" json:"deathStreak"`
	GameLength          int            `db:"time_seconds" json:"gameLength"`
	KillsPerMinute      float64        `db:"kills_per_minute" json:"killsPerMinute"`
	DeathsPerMinute     float64        `db:"deaths_per_minute" json:"deathsPerMinute"`
	KDRatio             float64        `db:"kill_death_ratio" json:"KDRatio"`
	LongestLifeSeconds  int            `db:"longest_life_secs" json:"longestLifeSeconds"`
	ShortestLifeSeconds int            `db:"shortest_life_secs" json:"shortestLifeSeconds"`
	Weapons             map[string]int `db:"weapons" json:"weapons,omitempty"`
	MapName             string         `db:"map_name" json:"mapName"`
	ServerNumber        int            `db:"server_number" json:"serverNumber"`
}

type PlayerStats struct {
	SteamInfo       steamapi.PlayerSummary `json:"steamInfo"`
	PublicAvg       AvgStats               `json:"publicAvg"`
	CompAvg         AvgStats               `json:"compAvg"`
	LastPublicGames []GameStats            `json:"publicGames,omitempty"`
	LastCompGames   []GameStats            `json:"compGames,omitempty"`
}

var Store *session.Store
var DbPool *pgxpool.Pool
