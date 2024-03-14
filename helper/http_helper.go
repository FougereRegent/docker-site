package helper

import (
	"errors"
	"net"
	"net/http"
)

var __client *http.Client

func buildFakeDial(unixPath string) func(proto, addr string) (conn net.Conn, err error) {
	result := func(proto, addr string) (conn net.Conn, err error) {
		return net.Dial("unix", unixPath)
	}
	return result
}

func InitClient(unixSocket string) *http.Client {
	transport := &http.Transport{
		Dial: buildFakeDial(unixSocket),
	}
	__client = &http.Client{
		Transport: transport,
	}

	return __client
}

func GetClient() (*http.Client, error) {
	if __client == nil {
		return nil, errors.New("Http client is null")
	}
	return __client, nil
}
