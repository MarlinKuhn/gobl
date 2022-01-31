package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"gitlab.com/flimzy/testy"
)

func Test_serve_build(t *testing.T) {
	tests := []struct {
		name string
		req  *http.Request
		err  string
	}{
		{
			name: "wrong content type",
			req: func() *http.Request {
				req, _ := http.NewRequest(http.MethodPost, "/build", nil)
				req.Header.Set("Content-Type", "text/plain")
				return req
			}(),
			err: "code=415, message=Unsupported Media Type",
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			e := echo.New()
			rec := httptest.NewRecorder()
			c := e.NewContext(tt.req, rec)

			err := serve().build()(c)
			if tt.err == "" {
				assert.Nil(t, err)
			} else {
				assert.EqualError(t, err, tt.err)
			}
			if err != nil {
				return
			}
			if d := testy.DiffHTTPResponse(testy.Snapshot(t), rec.Result()); d != nil {
				t.Error(d)
			}
		})
	}
}
