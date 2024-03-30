package main

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	writer  http.ResponseWriter
	request *http.Request
	body    JSONApiResponse
	headers map[string]string
}

func newResponse(writer http.ResponseWriter, request *http.Request) Response {
	return Response{
		writer:  writer,
		request: request,
		body:    (JSONApiResponse{}).success(nil),
		headers: (map[string]string{"Content-Type": "application/json"}),
	}
}

func (r *Response) SetHeaders(h map[string]string) Response {
	for key, value := range h {
		r.writer.Header().Add(key, value)
	}
	return *r
}

func (r *Response) SetBody(body JSONApiResponse) *Response {
	r.body = body

	return r
}

func (r *Response) Send() {
	json.NewEncoder(r.writer).Encode(r.body)
}

type JSONApiResponse struct {
	Message string `json:"message"`
	Status  uint8  `json:"status"`
	Data    any    `json:"Data"`
}

func (r JSONApiResponse) success(data any) JSONApiResponse {
	return CreateJSONResponse(JSONApiResponse{
		Data: data,
	})
}

func CreateJSONResponse(b JSONApiResponse) JSONApiResponse {
	if b.Message == "" {
		b.Message = "ok"
	}

	if b.Status == 0 {
		b.Status = http.StatusOK
	}

	if b.Data == nil {
		b.Data = struct{}{}
	}

	return JSONApiResponse{
		Message: b.Message,
		Status:  b.Status,
		Data:    b.Data,
	}
}
