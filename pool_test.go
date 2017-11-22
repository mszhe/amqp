package amqp

import (
	"fmt"
	"testing"
)

func TestGetPool(t *testing.T) {
	url := "amqp://root:123456@localhost:5672/"

	pool, err := GetPool(url, 5, 10)
	if err != nil {
		panic(err)
	}

	conn, err := pool.Get()
	if err != nil {
		panic(err)
	}

	connection, err := DialWithConn(conn, url)
	if err != nil {
		panic(err)
	}
	defer connection.Close()

	fmt.Println(connection)
}
