package IPFSClient

import "errors"

// - [ ] stats
//   - [ ] /api/v0/stats/bitswap
func (c *IPFSClient) StatsBitswap() (res string, err error) {
	err = errors.New("unimplement api")
	return "", err
}

//   - [ ] /api/v0/stats/bw
func (c *IPFSClient) StatsBw() (res string, err error) {
	err = errors.New("unimplement api")
	return "", err
}

//   - [ ] /api/v0/stats/dht
func (c *IPFSClient) StatsDht() (res string, err error) {
	err = errors.New("unimplement api")
	return "", err
}

//   - [ ] /api/v0/stats/provide
func (c *IPFSClient) StatsProvide() (res string, err error) {
	err = errors.New("unimplement api")
	return "", err
}

//   - [ ] /api/v0/stats/repo
func (c *IPFSClient) StatsRepo() (res string, err error) {
	err = errors.New("unimplement api")
	return "", err
}
