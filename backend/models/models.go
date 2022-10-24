package models

import (
	"github.com/Philipp15b/go-steamapi"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/jackc/pgx/v4/pgxpool"
	"time"
)

type DiscordGuildMember struct {
	User                       DiscordUser `json:"user"`
	Nick                       string      `json:"nick,omitempty"`
	Avatar                     string      `json:"avatar" json:"avatar,omitempty"`
	Roles                      []string    `json:"roles,omitempty"`
	JoinedAt                   time.Time   `json:"joined_at,omitempty"`
	PremiumSince               time.Time   `json:"premium_since,omitempty"`
	Deaf                       bool        `json:"deaf,omitempty"`
	Mute                       bool        `json:"mute,omitempty"`
	Pending                    bool        `json:"is_pending,omitempty"`
	Permissions                string      `json:"permissions,omitempty"`
	CommunicationDisabledUntil time.Time   `json:"communication_disabled_until,omitempty"`
	Flags                      int         `json:"flags,omitempty"`
}

type DiscordUser struct {
	Id               string `json:"id,omitempty"`
	Username         string `json:"username,omitempty"`
	Avatar           string `json:"avatar,omitempty"`
	Discriminator    string `json:"discriminator,omitempty"`
	AvatarDecoration string `json:"avatar_decoration,omitempty"`
	PublicFlags      int    `json:"public_flags,omitempty"`
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

type KPMStats struct {
	SteamID         string  `db:"steam_id_64" json:"steamId"`
	Name            string  `db:"used_name" json:"name"`
	Kills           float64 `db:"kills" json:"kills"`
	Deaths          float64 `db:"deaths" json:"deaths"`
	KillsPerMinute  float64 `db:"kills_per_minute" json:"killsPerMinute"`
	DeathsPerMinute float64 `db:"deaths_per_minute" json:"deathsPerMinute"`
	KDRatio         float64 `db:"kill_death_ratio" json:"KDRatio"`
}

type HighKillData struct {
	SteamID string `db:"steam_id_64" json:"steamId"`
	Name    string `db:"used_name" json:"name"`
	Kills   int    `db:"max_kill_game" json:"kills"`
}

var Store *session.Store
var DbPool *pgxpool.Pool
