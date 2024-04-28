package tests

import (
	"io"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHello(t *testing.T) {
	resp, err := http.Get("http://localhost:8080/hello")
	assert.Nil(t, err, "Request failed")

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	assert.Nil(t, err, "Failed to read the body")

	assert.Equal(t, 200, resp.StatusCode, "Wrong status code")
	assert.Equal(t, "world", string(body), "Wrong body result")
}
