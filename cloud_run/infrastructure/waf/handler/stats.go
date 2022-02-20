package handler

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi"
	"io"
	"line-halo-stats/cloud_run/domain/entity/dialogflow_request"
	"line-halo-stats/cloud_run/domain/entity/dialogflow_response"
	"line-halo-stats/cloud_run/interface/controller"
	"line-halo-stats/cloud_run/lib/context"
	"line-halo-stats/cloud_run/lib/log"
	"net/http"
)

const (
	PlayerStatsBasePath   = "/stats"
	playerStatsAllAPIPath = "/all"
)

type IPlayerStatsHandler interface {
	PlayerStats(r chi.Router)
}

var _ IPlayerStatsHandler = &PlayerStatsHandler{}

type PlayerStatsHandler struct {
	con *controller.PlayerStatsController
}

func NewPlayerStatsHandler(con *controller.PlayerStatsController) IPlayerStatsHandler {
	return &PlayerStatsHandler{con: con}
}

func (h *PlayerStatsHandler) PlayerStats(r chi.Router) {
	r.Post(playerStatsAllAPIPath, h.getAllPlayerStats)
}

func (h *PlayerStatsHandler) getAllPlayerStats(w http.ResponseWriter, r *http.Request) {
	ctx := context.New(r)
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Errorf(ctx, "error reading request body: %+v", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("couldn't read data"))
		return
	}

	dfRequest := &dialogflow_request.DialogflowRequest{}
	err = json.Unmarshal(body, dfRequest)
	if err != nil {
		log.Errorf(ctx, "couldn't unmarshal the dialogflow request: %+v", err)
		log.Infof(ctx, "dialogflow request body: %+v", string(body))
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("couldn't unmarshal data"))
		return
	}

	gt, err := dfRequest.ParseGamertag(ctx)
	if err != nil {
		log.Errorf(ctx, "couldn't parse gamertag: %+v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	stats, err := h.con.GetPlayerStats(ctx, gt)
	if err != nil {
		log.Errorf(ctx, "couldn't get player stats for %+v", gt)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// TODO fix this
	var res string
	if stats.SoloDuoRank == -1 && stats.CrossplayRank == -1 {
		res = fmt.Sprintf("%s's stats are either hidden or unavailable", stats.Gamertag)
	} else {
		res = stats.String()
	}

	playerStatsResponse := dialogflow_response.NewResponse(res)
	response, err := json.Marshal(playerStatsResponse)
	if err != nil {
		log.Errorf(ctx, "couldn't marshal response: %+v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(response)
	w.WriteHeader(http.StatusOK)
}
