package erg

type node struct{}

type Node interface {
	Add(n Node)
	Request(n Node)
	Respond(query [32]byte) []byte
}

// Add appends a new peer to the peer list. There is no remove, but least
// recently seen can be pruned to bound the size. It is a single table key value
// store, so it is neither large nor slow to access.
func (n2 *node) Add(n Node) {
	// TODO implement me
	panic("implement me")
}

// Request sends a hash of the last known state of a peer's recent event log, to
// which the peer will report no change or send a new log that has a different
// hash
func (n2 *node) Request(n Node) {
	// TODO implement me
	panic("implement me")
}

// Respond returns an update if its log has changed from the given hash
func (n2 *node) Respond(query [32]byte) []byte {
	// TODO implement me
	panic("implement me")
}
