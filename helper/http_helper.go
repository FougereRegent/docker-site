package helper

import (
	"errors"
	"fmt"
	"io"
	"net"
	"net/http"
	"strings"
)

type Method int

const (
	POST Method = 0
	GET  Method = 1
)

type MethodAction interface {
	Send(path string, expectStatusCode int) (result []byte, err error)
	ReceiveStream(path string, expectStatusCode int, dataQueue chan string, quit chan int) error
}

type HttpClient struct {
	Client  *http.Client
	UrlBase string
}

type BaseMethod struct {
	header map[string]interface{}
}

type GetMethod struct {
	*BaseMethod
}

type PostMethod struct {
	*BaseMethod
	data string
}

type DeleteMethod struct {
	*BaseMethod
}

type PutMethod struct {
	*BaseMethod
	data string
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

func MakeRequest(method Method) MethodAction {
	var result MethodAction
	switch method {
	case POST:
		result = &PostMethod{}
		break
	case GET:
		result = &GetMethod{}
		break
	}
	return result
}

func (meth *GetMethod) Send(path string, expectStatusCode int) (result []byte, err error) {
	url := fmt.Sprintf("%s/%s", __client.UrlBase, path)
	resp, err := __client.Client.Get(url)

	if err != nil {
		return nil, err
	}

	if resp.StatusCode != expectStatusCode {
		return nil, errors.New(fmt.Sprintf("%s : %d", "Status non attendu", resp.StatusCode))
	}

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}
	return body, nil
}

func (meth *PostMethod) Send(path string, expectStatusCode int) (result []byte, err error) {
	url := fmt.Sprintf("%s/%s", __client.UrlBase, path)
	resp, err := __client.Client.Post(url, "application/json", nil)

	if err != nil {
		return nil, err
	}

	if resp.StatusCode != expectStatusCode {
		return nil, errors.New(fmt.Sprintf("%s : %d", "Status non attendu", resp.StatusCode))
	}

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}
	return body, nil
}

func (meth *GetMethod) ReceiveStream(path string, expectStatusCode int, dataQueue chan string, quit chan int) error {
	url := fmt.Sprintf("%s%s", __client.UrlBase, path)

	if resp, err := __client.Client.Get(url); err != nil {
		return err
	} else {
		for {
			buf := make([]byte, 1024)
			strBuilder := strings.Builder{}
			nbBytes, _ := resp.Body.Read(buf)
			strBuilder.Write(buf)

			for buf[nbBytes-1] != '\n' {
				nbBytes, _ = resp.Body.Read(buf)
				strBuilder.Write(buf[:nbBytes])
			}

			select {
			case <-quit:
				return nil
			default:
				dataQueue <- strBuilder.String()
			}
		}
	}
}

func (meth *PostMethod) ReceiveStream(path string, expectStatusCode int, dataQueue chan string, quit chan int) error {
	return errors.ErrUnsupported
}

func (meth *DeleteMethod) Send(path string, expectStatusCode int) (result []byte, err error) {
	return nil, nil
}

func (meth *PutMethod) Send(path string, expectStatusCode int) (result []byte, err error) {
	return nil, nil
}
