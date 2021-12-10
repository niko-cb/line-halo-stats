package dialogflow_response

type DialogflowResponse struct {
	FulfillmentMessages []FulfillmentMessages `json:"fulfillmentMessages"`
}

type FulfillmentMessages struct {
	Text Text `json:"text"`
}

type Text struct {
	Text []string `json:"text"`
}

func NewResponse(response string) *DialogflowResponse {
	fm := FulfillmentMessages{Text: Text{Text: []string{response}}}

	return &DialogflowResponse{FulfillmentMessages: []FulfillmentMessages{fm}}
}
