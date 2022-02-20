package dialogflow_request

import (
	"context"
	"errors"
	"fmt"
	"line-halo-stats/cloud_run/lib/log"
	"strings"
)

type DialogflowRequest struct {
	ResponseId  string `json:"responseId"`
	QueryResult struct {
		QueryText  string `json:"queryText"`
		Action     string `json:"action"`
		Parameters struct {
		} `json:"parameters"`
		AllRequiredParamsPresent bool   `json:"allRequiredParamsPresent"`
		FulfillmentText          string `json:"fulfillmentText"`
		FulfillmentMessages      []struct {
			Text struct {
				Text []string `json:"text"`
			} `json:"text"`
			Platform string `json:"platform,omitempty"`
		} `json:"fulfillmentMessages"`
		OutputContexts []struct {
			Name          string `json:"name"`
			LifespanCount int    `json:"lifespanCount"`
			Parameters    struct {
				NoInput float64 `json:"no-input"`
				NoMatch float64 `json:"no-match"`
			} `json:"parameters"`
		} `json:"outputContexts"`
		Intent struct {
			Name        string `json:"name"`
			DisplayName string `json:"displayName"`
			IsFallback  bool   `json:"isFallback"`
		} `json:"intent"`
		IntentDetectionConfidence float64 `json:"intentDetectionConfidence"`
		LanguageCode              string  `json:"languageCode"`
	} `json:"queryResult"`
	OriginalDetectIntentRequest struct {
		Source  string `json:"source"`
		Payload struct {
			Data struct {
				Timestamp  string `json:"timestamp"`
				ReplyToken string `json:"replyToken"`
				Source     struct {
					UserId string `json:"userId"`
					Type   string `json:"type"`
				} `json:"source"`
				Message struct {
					Id   string `json:"id"`
					Type string `json:"type"`
					Text string `json:"text"`
				} `json:"message"`
				Type string `json:"type"`
			} `json:"data"`
		} `json:"payload"`
	} `json:"originalDetectIntentRequest"`
	Session string `json:"session"`
}

func (r *DialogflowRequest) ParseGamertag(ctx context.Context) (string, error) {
	q := strings.Split(r.QueryResult.QueryText, " ")
	if len(q) <= 1 {
		log.Errorf(ctx, "invalid query: %+v", q)
		return "", errors.New(fmt.Sprintf("invalid query: %+v", q))
	}

	var newQ []string
	for i := 0; i < len(q)-1; i++ {
		newQ = append(newQ, q[i])
	}

	return strings.Join(newQ, " "), nil
}
