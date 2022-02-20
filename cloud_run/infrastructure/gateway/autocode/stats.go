package autocode

import (
	"bytes"
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

func NewAutocodeRepository(cfg *config.ServerConfig) repository.AutocodeRepository {
	return &AutocodeRepository{
		autocodeURL:  entity.AutocodeURL(cfg.AutocodeURL),
		autocodeAuth: entity.AutocodeAuth(cfg.AutocodeAuth),
	}
}

type AutocodeRepository struct {
	autocodeURL  entity.AutocodeURL
	autocodeAuth entity.AutocodeAuth
}

func (a AutocodeRepository) GetPlayerCSRS(ctx context.Context, gamertag string) (*entity.PlayerCSRS, error) {
	c := http.Client{}
	autocodeURLSuffix := "stats/csrs/"

	reqURL := fmt.Sprintf("%s/%s", a.autocodeURL.String(), autocodeURLSuffix)

	data := &entity.Payload{
		Gamertag: gamertag,
	}

	payload, _ := json.Marshal(data)

	req, err := http.NewRequest(http.MethodPost, reqURL, bytes.NewBuffer(payload))
	if err != nil {
		log.Errorf(ctx, "failed to create request to autocode: %+v", err)
		return nil, err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", a.autocodeAuth))
	resp, err := c.Do(req)

	if err != nil {
		return nil, fmt.Errorf("autocode request failed: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %v", err)
	}

	var playerCsrs *entity.PlayerCSRS
	err = json.Unmarshal(body, &playerCsrs)

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("request to autocode failed: StatusCode = %d, ResponseBody = %+v", resp.StatusCode, string(body))
	}

	return playerCsrs, nil
}

func (a AutocodeRepository) GetPlayerServiceRecord(ctx context.Context, gamertag string) (*entity.ServiceRecord, error) {

	c := http.Client{}
	autocodeURLSuffix := "stats/service-record/multiplayer"

	reqURL := fmt.Sprintf("%s/%s", a.autocodeURL, autocodeURLSuffix)

	data := &entity.Payload{
		Gamertag: gamertag,
		Filter:   "matchmade:pvp",
	}

	payload, _ := json.Marshal(data)

	req, err := http.NewRequest("POST", reqURL, bytes.NewBuffer(payload))
	if err != nil {
		log.Errorf(ctx, "failed to create request to autocode: %+v", err)
		return nil, err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", a.autocodeAuth))
	resp, err := c.Do(req)

	if err != nil {
		return nil, fmt.Errorf("autocode request failed: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %v", err)
	}

	log.Infof(ctx, "body ====== %+v", string(body))

	var sr *entity.ServiceRecord
	err = json.Unmarshal(body, &sr)

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("request to autocode failed: StatusCode = %d, ResponseBody = %+v", resp.StatusCode, string(body))
	}

	return sr, nil
}

func (a AutocodeRepository) GetPlayerStats(ctx context.Context, csrs *entity.PlayerCSRS, sr *entity.ServiceRecord) (*entity.PlayerStats, error) {
	return &entity.PlayerStats{
		PlayerStatsParam: entity.PlayerStatsParam{
			KD:            sr.Data.Core.Kdr,
			KDA:           sr.Data.Core.Kda,
			WinRate:       sr.Data.WinRate,
			TimePlayed:    sr.Data.TimePlayed.Human,
			ShotAccuracy:  sr.Data.Core.Shots.Accuracy,
			CrossplayRank: csrs.Data[0].Response.Current.Value,
			SoloDuoRank:   csrs.Data[1].Response.Current.Value,
		},
		Gamertag: csrs.Additional.Gamertag,
	}, nil
}
