// Package leader provides leader election
package leader

// Leader provides leadership election
type Leader interface {
	// elect leader
	Elect(id string, opts ...ElectOption) (Elected, error)
	// follow the leader
	Follow() chan string
}

type Elected interface {
	// id of leader
	Id() string
	// seek re-election
	Reelect() error
	// resign leadership
	Resign() error
	// observe leadership revocation
	Revoked() chan bool
}

type LeaderOptions struct {
	Nodes []string
	Group string
}

type LeaderOption func(o *LeaderOptions)

type ElectOptions struct{}

type ElectOption func(o *ElectOptions)

// Group sets the group name for coordinating leadership
func Group(g string) LeaderOption {
	return func(o *LeaderOptions) {
		o.Group = g
	}
}
