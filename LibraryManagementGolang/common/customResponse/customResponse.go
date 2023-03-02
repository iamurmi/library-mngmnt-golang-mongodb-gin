package customresponse

type Response struct {
	Sucess bool        `json:"sucess"`
	Data   interface{} `json:"data"`
	Error  string      `json:"error,omitempty"`
}

func SuccessResponse(data interface{}, err error) (resp Response) {
	resp = Response{
		Sucess: true,
		Data:   data,
		Error:  "",
	}
	return resp
}
func FailedResponse(data interface{}, err error) (resp Response) {
	errStr := err.Error()
	resp = Response{
		Sucess: false,
		Data:   nil,
		Error:  errStr,
	}
	return resp
}
