package etcd

import (
	"github.com/alexeyco/leader"
	client "github.com/coreos/etcd/clientv3"
)

// NewLeader returns etcd leader instance.
func NewLeader(c *client.Client, opts ...leader.LeaderOption) leader.Leader {
	var options leader.LeaderOptions
	for _, o := range opts {
		o(&options)
	}

	return &etcdLeader{
		client: c,
		opts:   options,
		path:   "/srv/leader",
	}
}
