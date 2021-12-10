package entity

type PlayerServiceRecord struct {
	Data       DataSR       `json:"data"`
	Additional AdditionalSR `json:"additional"`
}

type DataSR struct {
	Summary       Summary    `json:"summary"`
	Damage        Damage     `json:"damage"`
	Shots         Shots      `json:"shots"`
	Breakdowns    Breakdowns `json:"breakdowns"`
	Kda           float64    `json:"kda"`
	Kdr           float64    `json:"kdr"`
	TotalScore    int        `json:"total_score"`
	MatchesPlayed int        `json:"matches_played"`
	TimePlayed    TimePlayed `json:"time_played"`
	WinRate       float64    `json:"win_rate"`
}

type Summary struct {
	Kills     int      `json:"kills"`
	Deaths    int      `json:"deaths"`
	Assists   int      `json:"assists"`
	Betrayals int      `json:"betrayals"`
	Suicides  int      `json:"suicides"`
	Vehicles  Vehicles `json:"vehicles"`
	Medals    int      `json:"medals"`
}

type Vehicles struct {
	Destroys int `json:"destroys"`
	Hijacks  int `json:"hijacks"`
}

type Damage struct {
	Taken   int `json:"taken"`
	Dealt   int `json:"dealt"`
	Average int `json:"average"`
}

type Shots struct {
	Fired    int     `json:"fired"`
	Landed   int     `json:"landed"`
	Missed   int     `json:"missed"`
	Accuracy float64 `json:"accuracy"`
}

type Breakdowns struct {
	Kills   Kills   `json:"kills"`
	Assists Assists `json:"assists"`
	Matches Matches `json:"matches"`
}

type Kills struct {
	Melee        int `json:"melee"`
	Grenades     int `json:"grenades"`
	Headshots    int `json:"headshots"`
	PowerWeapons int `json:"power_weapons"`
}

type Assists struct {
	Emp      int `json:"emp"`
	Driver   int `json:"driver"`
	Callouts int `json:"callouts"`
}

type Matches struct {
	Wins   int `json:"wins"`
	Losses int `json:"losses"`
	Left   int `json:"left"`
	Draws  int `json:"draws"`
}

type TimePlayed struct {
	Seconds int    `json:"seconds"`
	Human   string `json:"human"`
}

type AdditionalSR struct {
	Gamertag string `json:"gamertag"`
}
