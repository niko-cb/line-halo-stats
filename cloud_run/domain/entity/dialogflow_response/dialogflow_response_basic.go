package dialogflow_response

type DialogflowResponseBasic struct {
	FulfillmentMessages []FulfillmentMessages `json:"fulfillmentMessages"`
}

type FulfillmentMessages struct {
	Text Text `json:"text"`
}

type Text struct {
	Text []string `json:"text"`
}

func NewResponse(response string) *DialogflowResponseBasic {
	fm := FulfillmentMessages{Text: Text{Text: []string{response}}}

	return &DialogflowResponseBasic{FulfillmentMessages: []FulfillmentMessages{fm}}
}
