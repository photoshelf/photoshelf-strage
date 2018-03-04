package server

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestNewServer(t *testing.T) {
	s := NewServer()
	assert.Implements(t, (*http.Handler)(nil), s)
}

func TestNewGateway(t *testing.T) {
	gw, err := NewGateway(":1234")
	if assert.NoError(t, err) {
		assert.Implements(t, (*http.Handler)(nil), gw)
	}
}
