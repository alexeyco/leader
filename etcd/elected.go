package etcd

import (
	"context"

	cc "github.com/coreos/etcd/clientv3/concurrency"
)

type etcdElected struct {
	id string
	s  *cc.Session
	e  *cc.Election
}

func (e *etcdElected) Id() string {
	return e.id
}

func (e *etcdElected) Reelect() error {
	return e.e.Campaign(context.TODO(), e.id)
}

func (e *etcdElected) Revoked() chan bool {
	ch := make(chan bool, 1)
	ech := e.e.Observe(context.Background())

	go func() {
		for r := range ech {
			if string(r.Kvs[0].Value) != e.id {
				ch <- true
				close(ch)
				return
			}
		}
	}()

	return ch
}

func (e *etcdElected) Resign() error {
	return e.e.Resign(context.Background())
}
