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

type RconScoreboardResult struct {
	Result RconScoreboard `json:"result"`
	Failed bool
}

type RconScoreboard struct {
	Id           int               `json:"id"`
	PlayerStats  []RconPlayerStats `json:"player_stats"`
	CreationTime MyTime            `json:"creation_time"`
	Start        MyTime            `json:"start"`
	End          MyTime            `json:"end"`
	ServerNumber int               `json:"server_number"`
	MapName      string            `json:"map_name"`
}

type RconPlayerStats struct {
	Id             int            `json:"id,omitempty" db:"id"`
	PlayerId       int            `json:"player_id,omitempty" db:"player_id"`
	Player         string         `json:"player,omitempty" db:"player"`
	SteamInfo      RconSteamInfo  `json:"steaminfo" db:"steaminfo"`
	MapId          int            `json:"map_id,omitempty" db:"map_id"`
	Kills          int            `json:"kills,omitempty" db:"kills"`
	KillStreak     int            `json:"kills_streak,omitempty" db:"kills_streak"`
	Deaths         int            `json:"deaths,omitempty" db:"deaths"`
	DeathStreak    int            `json:"deaths_without_kill_streak,omitempty" db:"deaths_without_kill_streak"`
	TeamKills      int            `json:"teamkills,omitempty" db:"teamkills"`
	TeamKillStreak int            `json:"teamkills_streak,omitempty" db:"teamkills_streak"`
	TKDeaths       int            `json:"deaths_by_tk,omitempty" db:"deaths_by_tk"`
	TKDeathStreak  int            `json:"deaths_by_tk_streak,omitempty" db:"deaths_by_tk_streak"`
	VoteStarted    int            `json:"nb_vote_started,omitempty" db:"nb_vote_started"`
	VotedYes       int            `json:"nb_voted_yes,omitempty" db:"nb_voted_yes"`
	VotedNo        int            `json:"nb_voted_no,omitempty" db:"nb_voted_no"`
	TimeSeconds    int            `json:"time_seconds,omitempty" db:"time_seconds"`
	KPM            float64        `json:"kills_per_minute,omitempty" db:"kills_per_minute"`
	DPM            float64        `json:"deaths_per_minute,omitempty" db:"deaths_per_minute"`
	KDRatio        float64        `json:"kill_death_ratio,omitempty" db:"kill_death_ratio"`
	LongestLife    int            `json:"longest_life_secs,omitempty" db:"longest_life_secs"`
	ShortestLife   int            `json:"shortest_life_secs,omitempty" db:"shortest_life_secs"`
	MostKilled     map[string]int `json:"most_killed,omitempty" db:"most_killed"`
	KilledBy       map[string]int `json:"death_by,omitempty" db:"death_by"`
	Weapons        map[string]int `json:"weapons,omitempty" db:"weapons"`
}

type RconSteamInfo struct {
	Id      int                    `json:"id,omitempty"`
	Created MyTime                 `json:"created,omitempty"`
	Updated MyTime                 `json:"updated,omitempty"`
	Profile steamapi.PlayerSummary `json:"profile"`
	Country string                 `json:"country,omitempty"`
}

type ExternalGame struct {
	Id     int          `json:"id,omitempty"`
	MapId  int          `json:"mapId,omitempty"`
	Name   string       `json:"name,omitempty"`
	URL    string       `json:"URL,omitempty"`
	Date   time.Time    `json:"date"`
	Squads []EventSquad `json:"squads"`
}

type EventSquad struct {
	Name    string        `json:"name"`
	Players []SquadPlayer `json:"players"`
}

type SquadPlayer struct {
	SteamID string `json:"steamId"`
}

type ExternalStats struct {
	SteamID string `json:"steamId" db:"steam_id"`
	Name    string `json:"name"`
	GameStats
}

var Store *session.Store
var DbPool *pgxpool.Pool

type MyTime struct {
	time.Time
}

func (self *MyTime) UnmarshalJSON(b []byte) (err error) {
	s := string(b)

	if s == "null" {
		return
	}

	// Get rid of the quotes "" around the value.
	// A second option would be to include them
	// in the date format string instead, like so below:
	//   time.Parse(`"`+time.RFC3339Nano+`"`, s)
	s = s[1 : len(s)-1]

	t, err := time.Parse(time.RFC3339Nano, s)
	if err != nil {
		t, err = time.Parse("2006-01-02T15:04:05.999999999", s)
	}
	self.Time = t
	return
}
