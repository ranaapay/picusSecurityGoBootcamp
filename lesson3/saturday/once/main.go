package main

import "sync"

//Singleton

/*
type DbConnection struct {}

var(
	conn *DbConnection
)

func GetConnection() *DbConnection {
	if conn == nil {
		conn = &DbConnection{}
	}
	return conn
}
*/

type DbConnection struct{}

var (
	dbConnectionOnce sync.Once
	conn             *DbConnection
)

func GetConnection() *DbConnection {
	dbConnectionOnce.Do(func() {
		conn = &DbConnection{}
	})
	return conn
}
