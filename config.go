package IPFSClient

import "errors"

// - [ ] config
//   - [ ] /api/v0/config
func (c *IPFSClient) Config() (res string, err error) {
	err = errors.New("unimplement api")
	return "", err
}

//   - [ ] /api/v0/config/edit
func (c *IPFSClient) ConfigEdit() (res string, err error) {
	err = errors.New("unimplement api")
	return "", err
}

//   - [ ] /api/v0/config/profile/apply
func (c *IPFSClient) ConfigProfileApply() (res string, err error) {
	err = errors.New("unimplement api")
	return "", err
}

//   - [ ] /api/v0/config/replace
func (c *IPFSClient) ConfigReplace() (res string, err error) {
	err = errors.New("unimplement api")
	return "", err
}

//   - [ ] /api/v0/config/show
func (c *IPFSClient) ConfigShow() (res string, err error) {
	err = errors.New("unimplement api")
	return "", err
}
