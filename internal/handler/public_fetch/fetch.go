package public_fetch

import (
	"net/http"

	"basic_golang/internal/adapter"

	"github.com/opentracing/opentracing-go"

	zaplogger "basic_golang/internal/adapter/zap"
)

func (s *Server) GetResources(w http.ResponseWriter, r *http.Request) {
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

	res, err := s.FetchDomain.GetResources(ctx, token)
	if err != nil {
		adapter.Response(w, http.StatusNotAcceptable, adapter.EmptyResponse, err.Error(), adapter.EmptyResponse, http.StatusNotAcceptable)
		return
	}

	adapter.Response(w, http.StatusOK, res, adapter.EmptyResponse, adapter.EmptyResponse, http.StatusOK)
}

func (s *Server) GetResourcesAdmin(w http.ResponseWriter, r *http.Request) {
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

	res, err := s.FetchDomain.GetResourcesAdmin(ctx, token)
	if err != nil {
		adapter.Response(w, http.StatusNotAcceptable, adapter.EmptyResponse, err.Error(), adapter.EmptyResponse, http.StatusNotAcceptable)
		return
	}

	adapter.Response(w, http.StatusOK, res, adapter.EmptyResponse, adapter.EmptyResponse, http.StatusOK)
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
