package IPFSClient

import "errors"

// - [ ] cid
//   - [ ] /api/v0/cid/base32
func (c *IPFSClient) CidBase32() (res string, err error) {
	err = errors.New("unimplement api")
	return "", err
}

//   - [ ] /api/v0/cid/bases
func (c *IPFSClient) CidBases() (res string, err error) {
	err = errors.New("unimplement api")
	return "", err
}

//   - [ ] /api/v0/cid/codecs
func (c *IPFSClient) CidCodecs() (res string, err error) {
	err = errors.New("unimplement api")
	return "", err
}

//   - [ ] /api/v0/cid/format
func (c *IPFSClient) CidFormat() (res string, err error) {
	err = errors.New("unimplement api")
	return "", err
}

//   - [ ] /api/v0/cid/hashes
func (c *IPFSClient) CidHashes() (res string, err error) {
	err = errors.New("unimplement api")
	return "", err
}
