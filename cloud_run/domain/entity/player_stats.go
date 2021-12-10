package entity

import "fmt"

type PlayerStatsParam struct {
	KD            float64
	KDA           float64
	WinRate       float64
	TimePlayed    string
	ShotAccuracy  float64
	CrossplayRank int
	SoloDuoRank   int
}

type PlayerStats struct {
	PlayerStatsParam
	Gamertag string
}

func NewPlayerStats(param PlayerStatsParam, gt string) *PlayerStats {
	return &PlayerStats{
		PlayerStatsParam: param,
		Gamertag:         gt,
	}
}

func (s *PlayerStats) String() string {
	return fmt.Sprintf("Stats for %+v"+
		"\n\n"+
		"Crossplay Rank: %+v"+
		"\n"+
		"Solo/Duo Rank: %+v"+
		"\n"+
		"KD: %+v"+
		"\n"+
		"KDA: %+v"+
		"\n"+
		"Accuracy: %+v"+
		"\n"+
		"Win Rate: %+v"+
		"\n"+
		"Time Played: %+v",
		s.Gamertag,
		s.CrossplayRank,
		s.SoloDuoRank,
		s.KD,
		s.KDA,
		s.ShotAccuracy,
		s.WinRate,
		s.TimePlayed)
}
