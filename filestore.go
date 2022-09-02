package IPFSClient

import "errors"

// - [ ] filestore
//   - [ ] /api/v0/filestore/dups
func (c *IPFSClient) FileStoreDups() (res string, err error) {
	err = errors.New("unimplement api")
	return "", err
}

//   - [ ] /api/v0/filestore/ls
func (c *IPFSClient) FileStoreLs() (res string, err error) {
	err = errors.New("unimplement api")
	return "", err
}

//   - [ ] /api/v0/filestore/verify
func (c *IPFSClient) FileStoreVerify() (res string, err error) {
	err = errors.New("unimplement api")
	return "", err
}
