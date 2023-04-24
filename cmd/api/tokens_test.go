package main

import (
	"encoding/json"
	"net/http"
	"testing"

	"greenlight.bcc/internal/assert"
)

func TestCreateAuthenticationToken(t *testing.T) {
	app := newTestApplication(t)
	ts := newTestServer(t, app.routesTest())
	defer ts.Close()

	const (
		validEmail    = "alice@example.com"
		validPassword = "pa55word"
	)

	// initData := struct {
	// 	Email    string `json:"email"`
	// 	Password string `json:"password"`
	// }{
	// 	Email:    validEmail,
	// 	Password: validPassword,
	// }

	// initBytes, err := json.Marshal(&initData)
	// if err != nil {
	// 	t.Fatalf("Failed to init data model: %v", err)
	// }

	// ts.postForm(t, "/v1/users", initBytes)

	tests := []struct {
		name     string
		Email    string
		Password string
		wantCode int
	}{
		{
			name:     "Wrong input",
			Email:    validEmail,
			Password: validPassword,
			wantCode: http.StatusBadRequest,
		},
		{
			name:     "No password",
			Email:    validEmail,
			Password: "",
			wantCode: http.StatusUnprocessableEntity,
		},
		// {
		// 	name:     "Non-existent email",
		// 	Email:    "non-existent@exaxample.com",
		// 	Password: validPassword,
		// 	wantCode: http.StatusUnauthorized,
		// },
		// {
		// 	name:     "Wrong password",
		// 	Email:    validEmail,
		// 	Password: "wrong password",
		// 	wantCode: http.StatusUnauthorized,
		// },
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			inputData := struct {
				Email    string `json:"email"`
				Password string `json:"password"`
			}{
				Email:    tt.Email,
				Password: tt.Password,
			}

			b, err := json.Marshal(&inputData)
			if err != nil {
				t.Fatal("wrong input data")
			}
			if tt.name == "Wrong input" {
				b = append(b, 'a')
			}

			code, header, body := ts.postForm(t, "/v1/tokens/authentication", b)
			t.Log(code, header, body)

			// var reader io.Reader
			// reader = strings.NewReader(`{"email": "alice@example.com","password":"pa55word"}`)
			// req, _ := http.NewRequest(http.MethodPost, "/v1/users", reader)
			// rr := httptest.NewRecorder()
			// handler := http.HandlerFunc(app.createAuthenticationTokenHandler)

			// handler.ServeHTTP(rr, req)

			assert.Equal(t, code, tt.wantCode)
		})
	}
}
