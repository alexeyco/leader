package etcd

import (
	"context"
	"path"
	"strings"

	"github.com/alexeyco/leader"
	client "github.com/coreos/etcd/clientv3"
	cc "github.com/coreos/etcd/clientv3/concurrency"
)

type etcdLeader struct {
	client *client.Client
	opts   leader.LeaderOptions
	path   string
}

func (e *etcdLeader) Elect(id string, opts ...leader.ElectOption) (leader.Elected, error) {
	var options leader.ElectOptions
	for _, o := range opts {
		o(&options)
	}

	// make path
	p := path.Join(e.path, strings.Replace(id, "/", "-", -1))

	s, err := cc.NewSession(e.client)
	if err != nil {
		return nil, err
	}

	l := cc.NewElection(s, p)

	if err := l.Campaign(context.TODO(), id); err != nil {
		return nil, err
	}

	return &etcdElected{
		e:  l,
		id: id,
	}, nil
}

func (e *etcdLeader) Follow() chan string {
	ch := make(chan string)

	s, err := cc.NewSession(e.client)
	if err != nil {
		return ch
	}

	l := cc.NewElection(s, e.path)
	ech := l.Observe(context.Background())

	go func() {
		for r := range ech {
			ch <- string(r.Kvs[0].Value)
		}
	}()

	return ch
}
