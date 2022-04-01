package test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Brianllp/go_practice/router"
	"github.com/stretchr/testify/assert"
)

func TestHello(t *testing.T) {
	e := router.NewRouter()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()

	e.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, "Hello, World!!!", rec.Body.String())
}
