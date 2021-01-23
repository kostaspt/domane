package http

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/tidwall/gjson"
)

func TestHandler_DomainsExtensions(t *testing.T) {
	initTestCase := func(body io.Reader) (rec *httptest.ResponseRecorder, err error) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodPost, "/", body)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec = httptest.NewRecorder()

		err = New().DomainsExtensions(e.NewContext(req, rec))
		return
	}

	t.Run("no data", func(t *testing.T) {
		rec, err := initTestCase(nil)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusBadRequest, rec.Code)
	})

	t.Run("no query", func(t *testing.T) {
		rec, err := initTestCase(strings.NewReader(`{"query":""}`))

		assert.NoError(t, err)
		assert.Equal(t, http.StatusBadRequest, rec.Code)
	})

	t.Run("valid", func(t *testing.T) {
		rec, err := initTestCase(strings.NewReader(`{"query":"hello world"}`))

		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, "helloworld.com", gjson.Get(rec.Body.String(), "0.domain").String())
		assert.Equal(t, "1", gjson.Get(rec.Body.String(), "0.position").String())
		assert.Equal(t, "2", gjson.Get(rec.Body.String(), "1.position").String())
	})
}

func TestHandler_DomainsSimilar(t *testing.T) {
	initTestCase := func(body io.Reader) (rec *httptest.ResponseRecorder, err error) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodPost, "/", body)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec = httptest.NewRecorder()

		err = New().DomainsSimilars(e.NewContext(req, rec))
		return
	}

	t.Run("no data", func(t *testing.T) {
		rec, err := initTestCase(nil)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusBadRequest, rec.Code)
	})

	t.Run("no query", func(t *testing.T) {
		rec, err := initTestCase(strings.NewReader(`{"query":""}`))

		assert.NoError(t, err)
		assert.Equal(t, http.StatusBadRequest, rec.Code)
	})

	t.Run("valid", func(t *testing.T) {
		rec, err := initTestCase(strings.NewReader(`{"query":"hello world"}`))

		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.NotEqual(t, "helloworld.com", gjson.Get(rec.Body.String(), "*.domain").String())
		assert.Equal(t, "1", gjson.Get(rec.Body.String(), "0.position").String())
		assert.Equal(t, "2", gjson.Get(rec.Body.String(), "1.position").String())
	})
}