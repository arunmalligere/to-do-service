package web

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"net/http"
)



func Respond(ctx context.Context, w http.ResponseWriter, data interface{}, statusCode int ) error {
 
	v, ok := ctx.Value(KeyValues).(*Values)

	if !ok {
		return errors.New("web value missing from context")
	}
	v.StatusCode = statusCode

	if statusCode == http.StatusNoContent {
		w.WriteHeader(statusCode)
		return nil
	}

	jsonData, err := json.Marshal(data)

	if err != nil {
		log.Println(err)
		return err
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	if _, err := w.Write(jsonData); err != nil {
		return err
	}

	return nil
}