package p2p

// Peer is an interface that represents the remote node.
type Peer interface {
}

// transport is anything that can handles the communications
// between the nodes in the network. This can be of the form( TCP,UDP, websockets....)

type Transport interface {
	ListenAndAccept() error
}
