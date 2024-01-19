package infrastructure

import (
	"time"

	"github.com/lodashventure/nlp/model"
)

var statusText = map[bool]string{
	true:  "Success",
	false: "Failed",
}

func Response(status bool, errorResponses []model.ErrorResponse) *model.Response {
	return &model.Response{
		Status:    statusText[status],
		Timestamp: time.Now().Format(time.RFC3339),
		Error:     errorResponses,
	}
}
