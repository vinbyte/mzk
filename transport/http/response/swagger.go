package response

type ExampleSuccessResponse struct {
	Code    int         `json:"code" example:"200"`
	Msg     string      `json:"msg" example:"success"`
	Records interface{} `json:"records"`
}

type ExampleBadRequestResponse struct {
	Code    int         `json:"code" example:"400"`
	Msg     string      `json:"msg" example:"startDate must less or equal than endDate"`
	Records interface{} `json:"records"`
}

type ExampleInternalErrorResponse struct {
	Code    int         `json:"code" example:"500"`
	Msg     string      `json:"msg" example:"internal server error"`
	Records interface{} `json:"records"`
}
