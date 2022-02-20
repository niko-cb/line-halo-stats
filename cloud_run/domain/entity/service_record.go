package entity

type ServiceRecord struct {
	Data struct {
		Core struct {
			Summary struct {
				Kills     int `json:"kills"`
				Deaths    int `json:"deaths"`
				Assists   int `json:"assists"`
				Betrayals int `json:"betrayals"`
				Suicides  int `json:"suicides"`
				Vehicles  struct {
					Destroys int `json:"destroys"`
					Hijacks  int `json:"hijacks"`
				} `json:"vehicles"`
				Medals int `json:"medals"`
			} `json:"summary"`
			Damage struct {
				Taken   int `json:"taken"`
				Dealt   int `json:"dealt"`
				Average int `json:"average"`
			} `json:"damage"`
			Shots struct {
				Fired    int     `json:"fired"`
				Landed   int     `json:"landed"`
				Missed   int     `json:"missed"`
				Accuracy float64 `json:"accuracy"`
			} `json:"shots"`
			Breakdowns struct {
				Kills struct {
					Melee        int `json:"melee"`
					Grenades     int `json:"grenades"`
					Headshots    int `json:"headshots"`
					PowerWeapons int `json:"power_weapons"`
				} `json:"kills"`
				Assists struct {
					Emp      int `json:"emp"`
					Driver   int `json:"driver"`
					Callouts int `json:"callouts"`
				} `json:"assists"`
				Matches struct {
					Wins   int `json:"wins"`
					Losses int `json:"losses"`
					Left   int `json:"left"`
					Draws  int `json:"draws"`
				} `json:"matches"`
				Medals []struct {
					Id        int    `json:"id"`
					Name      string `json:"name"`
					Count     int    `json:"count"`
					ImageUrls struct {
						Small  string `json:"small"`
						Medium string `json:"medium"`
						Large  string `json:"large"`
					} `json:"image_urls"`
				} `json:"medals"`
			} `json:"breakdowns"`
			Kda        float64 `json:"kda"`
			Kdr        float64 `json:"kdr"`
			TotalScore int     `json:"total_score"`
		} `json:"core"`
		MatchesPlayed int `json:"matches_played"`
		TimePlayed    struct {
			Seconds int    `json:"seconds"`
			Human   string `json:"human"`
		} `json:"time_played"`
		WinRate float64 `json:"win_rate"`
	} `json:"data"`
	Additional struct {
		Gamertag string `json:"gamertag"`
		Filter   string `json:"filter"`
	} `json:"additional"`
}
