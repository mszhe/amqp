package amqp

import (
	"errors"
	"gopkg.in/fatih/pool.v2"
	"net"
	"strconv"
)

func GetPool(url string, initCap, maxCap int) (pool.Pool, error) {
	uri, err := ParseURI(url)
	if err != nil {
		return nil, err
	}
	if 0 != initCap && 0 != maxCap {
		return pool.NewChannelPool(initCap, maxCap, func() (net.Conn, error) {
			return defaultDial("tcp", net.JoinHostPort(uri.Host, strconv.FormatInt(int64(uri.Port), 10)))
		})
	} else {
		return nil, errors.New("invalid capacity settings")
	}
}
