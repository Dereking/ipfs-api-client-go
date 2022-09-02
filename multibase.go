package IPFSClient

import "errors"

// - [ ] multibase
//   - [ ] /api/v0/multibase/decode
func (c *IPFSClient) MultibaseDecode() (res string, err error) {
	err = errors.New("unimplement api")
	return "", err
}

//   - [ ] /api/v0/multibase/encode
func (c *IPFSClient) MultibaseEncode() (res string, err error) {
	err = errors.New("unimplement api")
	return "", err
}

//   - [ ] /api/v0/multibase/list
func (c *IPFSClient) MultibaseList() (res string, err error) {
	err = errors.New("unimplement api")
	return "", err
}

//   - [ ] /api/v0/multibase/transcode
func (c *IPFSClient) MultibaseTranscode() (res string, err error) {
	err = errors.New("unimplement api")
	return "", err
}
