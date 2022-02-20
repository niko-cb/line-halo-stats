package entity

type PlayerCSRS struct {
	Data       []Data     `json:"data"`
	Additional Additional `json:"additional"`
}

type Data struct {
	Queue    string   `json:"queue"`
	Input    string   `json:"input"`
	Response Response `json:"response"`
}

type Additional struct {
	Gamertag string `json:"gamertag"`
	Season   int    `json:"season"`
}

type Response struct {
	Current Current `json:"current"`
	Season  Season  `json:"season"`
	AllTime AllTime `json:"all_time"`
}

type Current struct {
	Value                       int    `json:"value"`
	MeasurementMatchesRemaining int    `json:"measurement_matches_remaining"`
	Tier                        string `json:"tier"`
	TierStart                   int    `json:"tier_start"`
	SubTier                     int    `json:"sub_tier"`
	NextTier                    string `json:"next_tier"`
	NextTierStart               int    `json:"next_tier_start"`
	NextSubTier                 int    `json:"next_sub_tier"`
	InitialMeasurementMatches   int    `json:"initial_measurement_matches"`
	TierImageUrl                string `json:"tier_image_url"`
}

type Season struct {
	Value                       int    `json:"value"`
	MeasurementMatchesRemaining int    `json:"measurement_matches_remaining"`
	Tier                        string `json:"tier"`
	TierStart                   int    `json:"tier_start"`
	SubTier                     int    `json:"sub_tier"`
	NextTier                    string `json:"next_tier"`
	NextTierStart               int    `json:"next_tier_start"`
	NextSubTier                 int    `json:"next_sub_tier"`
	InitialMeasurementMatches   int    `json:"initial_measurement_matches"`
	TierImageUrl                string `json:"tier_image_url"`
}

type AllTime struct {
	Value                       int    `json:"value"`
	MeasurementMatchesRemaining int    `json:"measurement_matches_remaining"`
	Tier                        string `json:"tier"`
	TierStart                   int    `json:"tier_start"`
	SubTier                     int    `json:"sub_tier"`
	NextTier                    string `json:"next_tier"`
	NextTierStart               int    `json:"next_tier_start"`
	NextSubTier                 int    `json:"next_sub_tier"`
	InitialMeasurementMatches   int    `json:"initial_measurement_matches"`
	TierImageUrl                string `json:"tier_image_url"`
}
