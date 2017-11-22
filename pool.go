package amqp

import (
	"errors"
	"gopkg.in/fatih/pool.v2"
	"net"
	"strconv"
)

func GetPool(url string, initialCap, maxCap int) (pool.Pool, error) {
	uri, err := ParseURI(url)
	if err != nil {
		return nil, err
	}
	addr := net.JoinHostPort(uri.Host, strconv.FormatInt(int64(uri.Port), 10))
	dialer := defaultDial
	if 0 != initialCap && 0 != maxCap {
		if initialCap < 0 || maxCap <= 0 || initialCap > maxCap {
			return nil, errors.New("invalid capacity settings")
		}
		factory := func() (net.Conn, error) {
			return dialer("tcp", addr)
		}
		return pool.NewChannelPool(initialCap, maxCap, factory)
	} else {
		return nil, errors.New("invalid capacity settings")
	}
}
