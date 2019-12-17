package views

// Response -> API Return Default
type Response struct {
	Code    int         `json:"code"`
	Status  bool        `json:"status"`
	Message interface{} `json:"message"`
}

// ResponseWithData -> API Return With Data
type ResponseWithData struct {
	Code    int         `json:"code"`
	Status  bool        `json:"status"`
	Message interface{} `json:"message"`
	Data    interface{} `json:"data"`
}
