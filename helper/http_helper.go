package helper

import (
	"fmt"
	"io"
	"net"
	"net/http"
)

const (
	POST Method = 0
	GET  Method = 1
)

type Method int

type MethodAction interface {
	Send(path string) (result []byte, err error)
}

type HttpClient struct {
	Client  *http.Client
	UrlBase string
}

type GetMethod struct {
}

type PostMethod struct {
}

var __client *HttpClient

func buildFakeDial(unixPath string) func(proto, addr string) (conn net.Conn, err error) {
	result := func(proto, addr string) (conn net.Conn, err error) {
		return net.Dial("unix", unixPath)
	}
	return result
}

func InitClient(unixSocket string, base string) *HttpClient {
	transport := &http.Transport{
		Dial: buildFakeDial(unixSocket),
	}
	__client = &HttpClient{
		Client: &http.Client{
			Transport: transport,
		},
		UrlBase: base,
	}

	return __client
}

func (meth *GetMethod) Send(path string) (result []byte, err error) {
	url := fmt.Sprintf("%s/%s", __client.UrlBase, path)
	resp, err := __client.Client.Get(url)

	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}
	return body, nil
}

func (meth *PostMethod) Send(path string) (result []byte, err error) {
	url := fmt.Sprintf("%s/%s", __client.UrlBase, path)
	resp, err := __client.Client.Post(url, "application/json", nil)

	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}
	return body, nil
}
