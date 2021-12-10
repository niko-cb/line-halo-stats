package dialogflow_request

import (
	"context"
	"errors"
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

//type DialogflowRequest struct {
//	Session    string     `json:"session"`
//	QueryInput QueryInput `json:"query_input"`
//	Timezone   string     `json:"timezone"`
//}
//
//type QueryInput struct {
//	Text Text `json:"text"`
//}
//
//type Text struct {
//	TextInputs []TextInputs `json:"textInputs"`
//}
//
//type TextInputs struct {
//	Text string `json:"text"`
//}

//type DialogflowRequest struct {
//	QueryResult                 QueryResult                 `json:"queryResult"`
//	OriginalDetectIntentRequest OriginalDetectIntentRequest `json:"originalDetectIntentRequest"`
//}
//
//type QueryResult struct {
//	QueryText                 string               `json:"queryText"`
//	Parameters                Parameters           `json:"parameters"`
//	AllRequiredParamsPresent  bool                 `json:"allRequiredParamsPresent"`
//	FulfillmentText           string               `json:"fulfillmentText"`
//	FulfillmentMessages       FullfillmentMessages `json:"fulfillmentMessages"`
//	OutputContexts            OutputContexts       `json:"outputContexts"`
//	Intent                    Intent               `json:"intent"`
//	IntentDetectionConfidence int                  `json:"intentDetectionConfidence"`
//	DiagnosticInfo            DiagnosticInfo       `json:"diagnosticInfo"`
//	LanguageCode              string               `json:"languageCode"`
//}
//
//type FullfillmentMessages struct {
//	Text Text `json:"text"`
//}
//
//type Text struct {
//	Text []string `json:"text"`
//}
//
//type Parameters struct {
//	ParamName string `json:"param-name"`
//}
//
//type OutputContexts struct {
//	Name          string     `json:"name"`
//	LifespanCount int        `json:"lifespanCount"`
//	Parameters    Parameters `json:"parameters"`
//}
//
//type Intent struct {
//	Name        string `json:"name"`
//	DisplayName string `json:"displayName"`
//}
//
//type DiagnosticInfo struct{}
//
//type OriginalDetectIntentRequest struct{}
//
func (r *DialogflowRequest) ParseGamertag() (string, error) {
	ctx := context.Background()

	queryText := r.QueryResult.QueryText

	if len(queryText) == 0 {
		log.Errorf(ctx, "no query to parse: queryText = %+v", queryText)
		return "", errors.New("no query to parse")
	}

	if queryText == "" {
		return "", errors.New("no query to parse")
	}

	query := strings.Split(queryText, " ")
	if len(query) < 1 {
		return "", errors.New("invalid query")
	}

	gamertag := query[0]

	return gamertag, nil

}
