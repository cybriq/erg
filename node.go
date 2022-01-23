package erg

// This asserts the interface is satisfied if it is changed or an implementation
// has had its signature changed from the interface, this flags a compiler error
var _ Node = (*node)(nil)

// Node is the API for an erg Node
type Node interface {
	Add(n Node)
	Request(n Node)
	Respond(query [32]byte) []byte
}

// node is the initial implementation for which Node interface will be
// implemented
type node struct{}

func NewNode() *node {
	return &node{}
}

// Add appends a new peer to the peer list. There is no remove, but least
// recently seen can be pruned to bound the size. It is a single table key value
// store, so it is neither large nor slow to access.
func (n *node) Add(no Node) {
	// TODO implement me
	panic("implement me")
}

// Request sends a hash of the last known state of a peer's recent event log, to
// which the peer will report no change or send a new log that has a different
// hash
func (n *node) Request(no Node) {
	// TODO implement me
	panic("implement me")
}

// Respond returns an update if its log has changed from the given hash
func (n *node) Respond(query [32]byte) []byte {
	// TODO implement me
	panic("implement me")
}
