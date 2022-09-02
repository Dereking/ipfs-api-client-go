package IPFSClient

import "errors"

// - [ ] repo
//   - [ ] /api/v0/repo/gc
func (c *IPFSClient) RepoGc() (res string, err error) {
	err = errors.New("unimplement api")
	return "", err
}

//   - [ ] /api/v0/repo/migrate
func (c *IPFSClient) RepoMigrate() (res string, err error) {
	err = errors.New("unimplement api")
	return "", err
}

//   - [ ] /api/v0/repo/stat
func (c *IPFSClient) RepoStat() (res string, err error) {
	err = errors.New("unimplement api")
	return "", err
}

//   - [ ] /api/v0/repo/verify
func (c *IPFSClient) RepoVerify() (res string, err error) {
	err = errors.New("unimplement api")
	return "", err
}

//   - [ ] /api/v0/repo/version
func (c *IPFSClient) RepoVersion() (res string, err error) {
	err = errors.New("unimplement api")
	return "", err
}
