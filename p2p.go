package IPFSClient

import "errors"

// - [ ] Experimental RPC commands
//   - [ ] /api/v0/p2p/close
func (c *IPFSClient) P2pClose() (res string, err error) {
	err = errors.New("unimplement api")
	return "", err
}

//  Experimental - [ ] /api/v0/p2p/forward
func (c *IPFSClient) P2pForward() (res string, err error) {
	err = errors.New("unimplement api")
	return "", err
}

// Experimental  - [ ] /api/v0/p2p/listen
func (c *IPFSClient) P2pListen() (res string, err error) {
	err = errors.New("unimplement api")
	return "", err
}

//  Experimental - [ ] /api/v0/p2p/ls
func (c *IPFSClient) P2pLs() (res string, err error) {
	err = errors.New("unimplement api")
	return "", err
}

//  Experimental - [ ] /api/v0/p2p/stream/close
func (c *IPFSClient) P2pStreamClose() (res string, err error) {
	err = errors.New("unimplement api")
	return "", err
}

// Experimental  - [ ] /api/v0/p2p/stream/ls
func (c *IPFSClient) P2pStreamLs() (res string, err error) {
	err = errors.New("unimplement api")
	return "", err
}
