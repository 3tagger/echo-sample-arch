package dto

type HttpResponse struct {
	Message string      `json:"message,omitempty" example:"ok"`
	Result  interface{} `json:"result,omitempty"`
}

func GenericBadResponseResponse() *HttpResponse {
	return SimpleMessageResponse("please check your parameters")
}

func SimpleMessageResponse(arg ...string) *HttpResponse {
	msg := "ok"
	if len(arg) >= 1 {
		msg = arg[0]
	}

	return &HttpResponse{
		Message: msg,
	}
}

func SimpleResponse(result interface{}, arg ...string) *HttpResponse {
	resp := SimpleMessageResponse(arg...)

	resp.Result = result

	return resp
}
