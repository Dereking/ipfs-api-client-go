package IPFSClient

import (
	"fmt"
)

type IPFSClient struct {
	Host string
}

func NewIPFSClient(hostOrIp string, port int) *IPFSClient {
	return &IPFSClient{
		Host: fmt.Sprintf("http://%s:%d", hostOrIp, port),
	}
}
func NewIPFSClientLocal() *IPFSClient {
	return NewIPFSClient("127.0.0.1", 5001)
}
