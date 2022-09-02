package IPFSClient

import "errors"

// - [ ] swarm
//   - [ ] /api/v0/swarm/addrs
func (c *IPFSClient) SwarmAddrs() (res string, err error) {
	err = errors.New("unimplement api")
	return "", err
}

//   - [ ] /api/v0/swarm/addrs/listen
func (c *IPFSClient) SwarmAddrsListen() (res string, err error) {
	err = errors.New("unimplement api")
	return "", err
}

//   - [ ] /api/v0/swarm/addrs/local
func (c *IPFSClient) SwarmAddrsLocal() (res string, err error) {
	err = errors.New("unimplement api")
	return "", err
}

//   - [ ] /api/v0/swarm/connect
func (c *IPFSClient) SwarmConnect() (res string, err error) {
	err = errors.New("unimplement api")
	return "", err
}

//   - [ ] /api/v0/swarm/disconnect
func (c *IPFSClient) SwarmDisconect() (res string, err error) {
	err = errors.New("unimplement api")
	return "", err
}

//   - [ ] /api/v0/swarm/filters
func (c *IPFSClient) SwarmFilters() (res string, err error) {
	err = errors.New("unimplement api")
	return "", err
}

//   - [ ] /api/v0/swarm/filters/add
func (c *IPFSClient) SwarmFiltersAdd() (res string, err error) {
	err = errors.New("unimplement api")
	return "", err
}

//   - [ ] /api/v0/swarm/filters/rm
func (c *IPFSClient) SwarmFiltersRm() (res string, err error) {
	err = errors.New("unimplement api")
	return "", err
}

//   - [ ] /api/v0/swarm/peering/add
func (c *IPFSClient) SwarmPeeringAdd() (res string, err error) {
	err = errors.New("unimplement api")
	return "", err
}

//   - [ ] /api/v0/swarm/peering/ls
func (c *IPFSClient) SwarmPeeringLs() (res string, err error) {
	err = errors.New("unimplement api")
	return "", err
}

//   - [ ] /api/v0/swarm/peering/rm
func (c *IPFSClient) SwarmPeeringRm() (res string, err error) {
	err = errors.New("unimplement api")
	return "", err
}

//   - [ ] /api/v0/swarm/peers
func (c *IPFSClient) SwarmPeers() (res string, err error) {
	err = errors.New("unimplement api")
	return "", err
}

// - [ ] Experimental RPC commands
//  Experimental - [ ] /api/v0/swarm/limit
func (c *IPFSClient) SwarmLimit() (res string, err error) {
	err = errors.New("unimplement api")
	return "", err
}

//  Experimental - [ ] /api/v0/swarm/stats
func (c *IPFSClient) SwarmStats() (res string, err error) {
	err = errors.New("unimplement api")
	return "", err
}
