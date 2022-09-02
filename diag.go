package IPFSClient

import "errors"

// - [ ] diag
//   - [ ] /api/v0/diag/cmds
func (c *IPFSClient) DiagCmds() (res string, err error) {
	err = errors.New("unimplement api")
	return "", err
}

//   - [ ] /api/v0/diag/cmds/clear
func (c *IPFSClient) DiagClear() (res string, err error) {
	err = errors.New("unimplement api")
	return "", err
}

//   - [ ] /api/v0/diag/cmds/set-time
func (c *IPFSClient) DiagSetTime() (res string, err error) {
	err = errors.New("unimplement api")
	return "", err
}

//   - [ ] /api/v0/diag/profile
func (c *IPFSClient) DiagProfile() (res string, err error) {
	err = errors.New("unimplement api")
	return "", err
}

//   - [ ] /api/v0/diag/sys
func (c *IPFSClient) DiagSys() (res string, err error) {
	err = errors.New("unimplement api")
	return "", err
}
