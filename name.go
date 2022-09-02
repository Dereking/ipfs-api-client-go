package IPFSClient

import "errors"

// - [ ] name
//   - [ ] /api/v0/name/publish
func (c *IPFSClient) NamePublish() (res string, err error) {
	err = errors.New("unimplement api")
	return "", err
}

// / - [ ] /api/v0/name/resolve
func (c *IPFSClient) NameResolve() (res string, err error) {
	err = errors.New("unimplement api")
	return "", err
}

// - [ ] Experimental RPC commands
//   - [ ] /api/v0/name/pubsub/cancel
func (c *IPFSClient) NamePubsubCancel() (res string, err error) {
	err = errors.New("unimplement api")
	return "", err
}

// Experimental RPC commands  - [ ] /api/v0/name/pubsub/state
func (c *IPFSClient) NamePubsubState() (res string, err error) {
	err = errors.New("unimplement api")
	return "", err
}

//  Experimental RPC commands - [ ] /api/v0/name/pubsub/subs
func (c *IPFSClient) NamePubsubSubs() (res string, err error) {
	err = errors.New("unimplement api")
	return "", err
}
