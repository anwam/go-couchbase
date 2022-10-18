package main

import (
	"log"

	"github.com/couchbase/gocb/v2"
)

type Store struct {
	cb *gocb.Cluster
}

func NewStore(username, password string) *Store {
	cluster, err := gocb.Connect("couchbase://couchbasedb", gocb.ClusterOptions{
		Authenticator: gocb.PasswordAuthenticator{
			Username: username,
			Password: password,
		},
	})
	if err != nil {
		log.Fatal(err)
	}
	return &Store{
		cb: cluster,
	}
}
