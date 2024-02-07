package main

import "github.com/gocql/gocql"

var Session *gocql.Session

func init() {
    // Connect to ScyllaDB
    cluster := gocql.NewCluster("127.0.0.1")
    cluster.Keyspace = "todos_keyspace"
    var err error
    Session, err = cluster.CreateSession()
    if err != nil {
        panic(err)
    }
}
