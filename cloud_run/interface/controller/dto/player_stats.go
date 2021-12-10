package dto

import (
	"line-halo-stats/cloud_run/domain/entity"
)

type PlayerStats struct {
	entity.PlayerStatsParam
	Gamertag string `json:"gamertag"` // OrderIDは、OpportunityのSalesforceのIDと同じ
}

func (ps *PlayerStats) ToEntity() *entity.PlayerStats {
	return entity.NewPlayerStats(
		ps.PlayerStatsParam,
		ps.Gamertag,
	)
}
