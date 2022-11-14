package http

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// writeResponse - вспомогательная функция, которая записывет http статус-код и текстовое сообщение в ответ клиенту.
// Нужна для уменьшения дублирования кода и улучшения читаемости кода вызывающей функции.
func writeResponse(w http.ResponseWriter, status int, message string) {
	w.WriteHeader(status)
	_, _ = w.Write([]byte(message))
	_, _ = w.Write([]byte("\n"))
}

// writeJsonResponse - вспомогательная функция, которая запсывает http статус-код и сообщение в формате json в ответ клиенту.
// Нужна для уменьшения дублирования кода и улучшения читаемости кода вызывающей функции.
func writeJsonResponse(w http.ResponseWriter, status int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		writeResponse(w, http.StatusInternalServerError, fmt.Sprintf("can't marshal data: %s", err))
		return
	}

	w.Header().Set("Content-Type", "application/json")

	writeResponse(w, status, string(response))
}

func convert(js map[string]interface{}) map[string]string {
	result := map[string]string{}
	links, ok := js["data"].(map[string]interface{})
	if !ok {
		return result
	}
	for k := range links {
		if v, ok := links[k].(string); ok {
			fmt.Println(k,v)
			result[k] = v
		}
	}
	return result
}
