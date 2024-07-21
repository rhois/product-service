package presenter

type ChatRequest struct {
	Query string `json:"query"`
}

type ChatResponse struct {
	Response string `json:"response"`
}
