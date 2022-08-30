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
	Kills           float64 `db:"kills" json:"kills,omitempty"`
	Deaths          float64 `db:"deaths" json:"deaths,omitempty"`
	KillsPerMinute  float64 `db:"kills_per_minute" json:"killsPerMinute,omitempty"`
	DeathsPerMinute float64 `db:"deaths_per_minute" json:"deathsPerMinute,omitempty"`
	KDRatio         float64 `db:"kill_death_ratio" json:"KDRatio,omitempty"`
}

type GameStats struct {
	Date                time.Time      `db:"date" json:"date,omitempty"`
	Kills               int            `db:"kills" json:"kills,omitempty"`
	KillStreak          int            `db:"kills_streak" json:"killStreak,omitempty"`
	Deaths              int            `db:"deaths" json:"deaths,omitempty"`
	DeathStreak         int            `db:"deaths_without_kill_streak" json:"deathStreak,omitempty"`
	GameLength          int            `db:"time_seconds" json:"gameLength,omitempty"`
	KillsPerMinute      float64        `db:"kills_per_minute" json:"killsPerMinute,omitempty"`
	DeathsPerMinute     float64        `db:"deaths_per_minute" json:"deathsPerMinute,omitempty"`
	KDRatio             float64        `db:"kill_death_ratio" json:"KDRatio,omitempty"`
	LongestLifeSeconds  int            `db:"longest_life_secs" json:"longestLifeSeconds,omitempty"`
	ShortestLifeSeconds int            `db:"shortest_life_secs" json:"shortestLifeSeconds,omitempty"`
	Weapons             map[string]int `db:"weapons" json:"weapons,omitempty"`
	MapName             string         `db:"map_name" json:"mapName,omitempty"`
	ServerNumber        int            `db:"server_number" json:"serverNumber,omitempty"`
}

type PlayerStats struct {
	SteamInfo steamapi.PlayerSummary `json:"steamInfo"`
	Avg       AvgStats               `json:"avg" json:"avg"`
	LastGames []GameStats            `json:"lastGames,omitempty" json:"lastGames,omitempty"`
}

var Store *session.Store
var DbPool *pgxpool.Pool
