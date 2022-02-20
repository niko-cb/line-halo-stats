package dialogflow_response

type DialogflowResponseCarousel struct {
	Type     string `json:"type"`
	AltText  string `json:"altText"`
	Template struct {
		Columns []struct {
			ImageBackgroundColor string `json:"imageBackgroundColor"`
			DefaultAction        struct {
				Uri   string `json:"uri"`
				Type  string `json:"type"`
				Label string `json:"label"`
			} `json:"defaultAction"`
			Title   string `json:"title"`
			Actions []struct {
				Type  string `json:"type"`
				Data  string `json:"data,omitempty"`
				Label string `json:"label"`
				Uri   string `json:"uri,omitempty"`
			} `json:"actions"`
			Text              string `json:"text"`
			ThumbnailImageUrl string `json:"thumbnailImageUrl"`
		} `json:"columns"`
		Type             string `json:"type"`
		ImageAspectRatio string `json:"imageAspectRatio"`
		ImageSize        string `json:"imageSize"`
	} `json:"template"`
}
