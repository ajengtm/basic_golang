package public

import (
	"encoding/json"
	"fmt"
	"net/http"

	"basic_golang/internal/adapter"
	"basic_golang/internal/domain/auth/services"

	"github.com/opentracing/opentracing-go"
	"go.uber.org/zap"

	zaplogger "basic_golang/internal/adapter/zap"
)

func (s *Server) Login(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	span, ctx := opentracing.StartSpanFromContext(ctx, "httphandler_Login")
	defer span.Finish()
	logger := zaplogger.GetLogger()

	decoder := json.NewDecoder(r.Body)
	var request services.LoginRequest
	if err := decoder.Decode(&request); err != nil {
		err = fmt.Errorf("Failed to parse JSON Format")
		logger.Error(err.Error(), zap.Error(err))
		adapter.Response(w, http.StatusNotAcceptable, adapter.EmptyResponse, err.Error(), adapter.EmptyResponse, http.StatusNotAcceptable)
		return
	}

	jwtToken, err := s.AuthDomain.Login(ctx, &request)
	if err != nil {
		adapter.Response(w, http.StatusNotAcceptable, adapter.EmptyResponse, err.Error(), adapter.EmptyResponse, http.StatusNotAcceptable)
		return
	}

	adapter.Response(w, http.StatusCreated, jwtToken, adapter.EmptyResponse, adapter.EmptyResponse, http.StatusOK)
}

func (s *Server) CheckToken(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	span, ctx := opentracing.StartSpanFromContext(ctx, "httphandler_CheckToken")
	defer span.Finish()

	logger := zaplogger.GetLogger()

	q := r.URL.Query()
	token := q.Get("token")
	if token == "" {
		logger.Error("Invalid Token")
		adapter.Response(w, http.StatusNotAcceptable, adapter.EmptyResponse, "Error Parse JSON", adapter.EmptyResponse, http.StatusNotAcceptable)
		return
	}

	res, err := s.AuthDomain.CheckToken(ctx, token)
	if err != nil {
		adapter.Response(w, http.StatusNotAcceptable, adapter.EmptyResponse, err.Error(), adapter.EmptyResponse, http.StatusNotAcceptable)
		return
	}

	adapter.Response(w, http.StatusOK, res, adapter.EmptyResponse, adapter.EmptyResponse, http.StatusOK)
}

func (s *Server) UpsertUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	span, ctx := opentracing.StartSpanFromContext(ctx, "httphandler_UpsertUser")
	defer span.Finish()
	logger := zaplogger.GetLogger()

	decoder := json.NewDecoder(r.Body)
	var request services.UserRequest
	if err := decoder.Decode(&request); err != nil {
		err = fmt.Errorf("Failed to parse JSON Format")
		logger.Error(err.Error(), zap.Error(err))
		adapter.Response(w, http.StatusNotAcceptable, adapter.EmptyResponse, err.Error(), adapter.EmptyResponse, http.StatusNotAcceptable)
		return
	}

	user, err := s.AuthDomain.UpsertUser(ctx, &request)
	if err != nil {
		adapter.Response(w, http.StatusNotAcceptable, adapter.EmptyResponse, err.Error(), adapter.EmptyResponse, http.StatusNotAcceptable)
		return
	}

	adapter.Response(w, http.StatusCreated, user, adapter.EmptyResponse, adapter.EmptyResponse, http.StatusOK)
}
