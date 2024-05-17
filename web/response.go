package web

type ResponseStatusOk struct {
	Status bool `json:"status"`
	Data   any  `json:"data"`
}
