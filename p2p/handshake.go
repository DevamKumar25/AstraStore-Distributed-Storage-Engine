package p2p

import (
	"errors"
)

// ErrInvalidHandshake is returned if the handskae btw the local
// and remote node could not be established
var ErrInvalidHandshake = errors.New("Invalid handshake")

type Handshaker interface {
	handshake() error
}

type HandshakeFunc func(Peer) error

func NOPHandshakeFunc(Peer) error { return nil }
