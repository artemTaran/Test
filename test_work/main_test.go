package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_request(t *testing.T) {

	handler := http.HandlerFunc(_request)
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/request", nil)
	handler.ServeHTTP(rec, req)
}

func Test_admin_requests(t *testing.T) {
	handler := http.HandlerFunc(_adminRequests)
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/admin/requests", nil)
	handler.ServeHTTP(rec, req)
}
