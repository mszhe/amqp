package amqp

import (
	"fmt"
	"testing"
	"time"
)

func TestGetPool(t *testing.T) {
	initCap := 5
	maxCap := 10
	url := "amqp://root:123456@192.168.1.254:5672/"

	pool, err := GetPool(url, initCap, maxCap)
	if err != nil {
		panic(err)
	}

	c, err := pool.Get()
	if err != nil {
		panic(err)
	}

	conn, err := DialWithConn(c, url)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	fmt.Println(conn)
	time.Sleep(time.Second*999999)
}
