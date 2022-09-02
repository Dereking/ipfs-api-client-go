package IPFSClient

import "errors"

// - [ ] Experimental RPC commands
//  Experimental - [ ] /api/v0/pubsub/ls
func (c *IPFSClient) PubsubLs() (res string, err error) {
	err = errors.New("unimplement api")
	return "", err
}

// Experimental  - [ ] /api/v0/pubsub/peers
func (c *IPFSClient) PubsubPeers() (res string, err error) {
	err = errors.New("unimplement api")
	return "", err
}

//  Experimental - [ ] /api/v0/pubsub/pub
func (c *IPFSClient) PubsubPub() (res string, err error) {
	err = errors.New("unimplement api")
	return "", err
}

//  Experimental - [ ] /api/v0/pubsub/sub
func (c *IPFSClient) PubsubSub() (res string, err error) {
	err = errors.New("unimplement api")
	return "", err
}
