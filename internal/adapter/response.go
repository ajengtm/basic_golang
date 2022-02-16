package adapter

import (
	"encoding/json"
	"net/http"

	zaplogger "basic_golang/internal/adapter/zap"

	"go.uber.org/zap"
)

type empty struct{}

var EmptyResponse empty

func Response(w http.ResponseWriter, httpStatusHeader int, data interface{}, errors interface{}, meta interface{}, code int) {
	logger := zaplogger.GetLogger()
	apiResponse := struct {
		Data   interface{} `json:"data"`
		Errors interface{} `json:"errors"`
		Meta   interface{} `json:"meta"`
	}{
		data,
		errors,
		meta,
	}
	if code == http.StatusOK {
		logger.Info("API Response", zap.Any("API Response", apiResponse))
	} else {
		logger.Error("API Response", zap.Any("API Response", apiResponse))
	}

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpStatusHeader)

	_ = json.NewEncoder(w).Encode(apiResponse)
}
