package gokithelper

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
)

func DecodeResponseToHTTP(_ context.Context, w http.ResponseWriter, response interface{}) error {
	if resp, ok := response.(Response); ok {
		w.Header().Add("Content-Type", "application/json")
		if failedEvent := resp.FailedEvent(); failedEvent != nil {
			return json.NewEncoder(w).Encode(failedEvent)
		}
		return json.NewEncoder(w).Encode(resp.SuccessEvent())
	}
	return errors.New("response should implement gokithelper.Response")
}
