package tests

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRepeatQuery(t *testing.T) {
	tests := []struct {
		query          string
		expectedStatus int
		expectedBody   string
	}{
		{"?message=SoftwarePoolPoC2023", 200, "SoftwarePoolPoC2023"},
		{"?random=SoftwarePoolPoC2023", 400, ""},
		{"", 400, ""},
	}

	for _, test := range tests {
		resp, err := http.Get("http://localhost:8080/repeat-my-query" + test.query)
		assert.Nil(t, err, "Request failed")

		defer resp.Body.Close()
		body, err := io.ReadAll(resp.Body)
		assert.Nil(t, err, "Failed to read the body")

		assert.Equal(t, test.expectedStatus, resp.StatusCode, "Wrong status code")
		assert.Equal(t, test.expectedBody, string(body), "Wrong body result")
	}
}

func TestRepeatParam(t *testing.T) {
	tests := []struct {
		query          string
		expectedStatus int
		expectedBody   string
	}{
		{"SoftwarePoolPoC2023", 200, "SoftwarePoolPoC2023"},
		{"", 404, "404 page not found"},
	}

	for _, test := range tests {
		resp, err := http.Get("http://localhost:8080/repeat-my-param/" + test.query)
		assert.Nil(t, err, "Request failed")

		defer resp.Body.Close()
		body, err := io.ReadAll(resp.Body)
		assert.Nil(t, err, "Failed to read the body")

		assert.Equal(t, test.expectedStatus, resp.StatusCode, "Wrong status code")
		assert.Equal(t, test.expectedBody, string(body), "Wrong body result")
	}
}

func TestRepeatBody(t *testing.T) {
	tests := []struct {
		body           string
		expectedStatus int
		expectedBody   string
	}{
		{"SoftwarePoolPoC2023", 200, "SoftwarePoolPoC2023"},
		{"", 400, ""},
	}

	for _, test := range tests {
		reqBody, _ := json.Marshal(test.body)

		resp, err := http.Post("http://localhost:8080/repeat-my-body", "application/json", bytes.NewBuffer(reqBody))
		assert.Nil(t, err, "Request failed")

		defer resp.Body.Close()
		body, err := io.ReadAll(resp.Body)
		assert.Nil(t, err, "Failed to read the body")

		assert.Equal(t, test.expectedStatus, resp.StatusCode, "Wrong status code")
		assert.Equal(t, test.expectedBody, string(body), "Wrong body result")
	}
}

func TestRepeatHeader(t *testing.T) {
	tests := []struct {
		key            string
		value          string
		expectedStatus int
		expectedBody   string
	}{
		{"X-Message", "SoftwarePoolPoC2023", 200, "SoftwarePoolPoC2023"},
		{"message", "SoftwarePoolPoC2023", 400, ""},
	}

	for _, test := range tests {
		req, err := http.NewRequest("GET", "http://localhost:8080/repeat-my-header", nil)
		assert.Nil(t, err, "Request failed")

		req.Header.Set(test.key, test.value)
		client := &http.Client{}
		resp, err := client.Do(req)
		assert.Nil(t, err, "Request failed")

		defer resp.Body.Close()
		body, err := io.ReadAll(resp.Body)
		assert.Nil(t, err, "Failed to read the body")

		assert.Equal(t, test.expectedStatus, resp.StatusCode, "Wrong status code")
		assert.Equal(t, test.expectedBody, string(body), "Wrong body result")
	}

}

func TestRepeatCookie(t *testing.T) {
	tests := []struct {
		cookie         http.Cookie
		expectedStatus int
		expectedBody   string
	}{
		{http.Cookie{Name: "message", Value: "SoftwarePoolPoC2023"}, 200, "SoftwarePoolPoC2023"},
		{http.Cookie{Name: "random", Value: "SoftwarePoolPoC2023"}, 400, ""},
		{http.Cookie{}, 400, ""},
	}

	for _, test := range tests {
		req, err := http.NewRequest("GET", "http://localhost:8080/repeat-my-cookie", nil)
		assert.Nil(t, err, "Request failed")

		req.AddCookie(&test.cookie)
		client := &http.Client{}
		resp, err := client.Do(req)
		assert.Nil(t, err, "Request failed")

		defer resp.Body.Close()
		body, err := io.ReadAll(resp.Body)
		assert.Nil(t, err, "Failed to read the body")

		assert.Equal(t, test.expectedStatus, resp.StatusCode, "Wrong status code")
		assert.Equal(t, test.expectedBody, string(body), "Wrong body result")
	}

}
