package p2p

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTCPTransport(t *testing.T) {
	listenAddr := ":4000"
	tr := NewTCPTransport(TCPTransportOpts{ListenAddr: listenAddr})

	assert.Equal(t, tr.ListenAddr, listenAddr)

	// server
	// tr.Accept()

	assert.Nil(t, tr.ListenAndAccept())

}
