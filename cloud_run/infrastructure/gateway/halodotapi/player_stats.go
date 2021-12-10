package halodotapi

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"line-halo-stats/cloud_run/domain/entity"
	"line-halo-stats/cloud_run/domain/repository"
	"line-halo-stats/cloud_run/infrastructure/waf/config"
	"line-halo-stats/cloud_run/lib/log"
	"net/http"
)

func NewHalodotapiRepository(cfg *config.ServerConfig) repository.HalodotapiRepository {
	return &HalodotapiRepository{
		halodotapiURL: entity.HalodotapiURL(cfg.HalodotapiURL),
		CryptumAuth:   entity.CryptumAuth(cfg.CryptumAuth),
	}
}

type HalodotapiRepository struct {
	halodotapiURL entity.HalodotapiURL
	CryptumAuth   entity.CryptumAuth
}

func (h HalodotapiRepository) GetPlayerCSRS(ctx context.Context, gamertag string) (*entity.PlayerCSRS, error) {
	c := http.Client{}
	csrsSuffix := "csrs?season=1"

	reqURL := fmt.Sprintf("%+v%s/%s", h.halodotapiURL, gamertag, csrsSuffix)
	req, err := http.NewRequest("GET", reqURL, nil)
	if err != nil {
		log.Errorf(ctx, "failed to create request to cryptum: %+v", err)
		return nil, err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", string(h.CryptumAuth))
	req.Header.Add("Cryptum-API-Version", "2.3-alpha")
	resp, err := c.Do(req)

	if err != nil {
		return nil, fmt.Errorf("halodotapi request failed: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %v", err)
	}

	var playerCsrs *entity.PlayerCSRS
	err = json.Unmarshal(body, &playerCsrs)

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("request to halodotapi failed: StatusCode = %d, ResponseBody = %+v", resp.StatusCode, string(body))
	}

	return playerCsrs, nil
}

func (h HalodotapiRepository) GetPlayerServiceRecord(ctx context.Context, gamertag string) (*entity.PlayerServiceRecord, error) {

	c := http.Client{}
	srSuffix := "service-record/global"

	reqURL := fmt.Sprintf("%+v%s/%s", h.halodotapiURL, gamertag, srSuffix)
	req, err := http.NewRequest("GET", reqURL, nil)
	if err != nil {
		log.Errorf(ctx, "failed to create request to cryptum: %+v", err)
		return nil, err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", string(h.CryptumAuth))
	req.Header.Add("Cryptum-API-Version", "2.3-alpha")
	resp, err := c.Do(req)

	if err != nil {
		return nil, fmt.Errorf("halodotapi request failed: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %v", err)
	}

	var sr *entity.PlayerServiceRecord
	err = json.Unmarshal(body, &sr)

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("request to halodotapi failed: StatusCode = %d, ResponseBody = %+v", resp.StatusCode, string(body))
	}

	return sr, nil
}

func (h HalodotapiRepository) GetPlayerStats(ctx context.Context, csrs *entity.PlayerCSRS, sr *entity.PlayerServiceRecord) (*entity.PlayerStats, error) {
	return &entity.PlayerStats{
		PlayerStatsParam: entity.PlayerStatsParam{
			KD:            sr.Data.Kdr,
			KDA:           sr.Data.Kda,
			WinRate:       sr.Data.WinRate,
			TimePlayed:    sr.Data.TimePlayed.Human,
			ShotAccuracy:  sr.Data.Shots.Accuracy,
			CrossplayRank: csrs.Data[0].Response.Current.Value,
			SoloDuoRank:   csrs.Data[1].Response.Current.Value,
		},
		Gamertag: csrs.Additional.Gamertag,
	}, nil
}
