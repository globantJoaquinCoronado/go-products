package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func RespondWithError(w http.ResponseWriter, code int, msg string) {
	fmt.Println("Error", msg)
	params := map[string]string{
		"message": msg,
		"code":    strconv.Itoa(code),
		"error":   http.StatusText(code),
	}
	RespondWithJson(w, code, params)
}

func RespondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)
	_, _ = w.Write(response)
}
