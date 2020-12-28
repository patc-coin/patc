// Copyright 2019 The go-pc Authors
// This file is part of the go-pc library.
//
// The go-pc library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-pc library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-pc library. If not, see <http://www.gnu.org/licenses/>.

package les

import (
	"pc.dp.tc/p2p/dnsdisc"
	"pc.dp.tc/p2p/enode"
	"pc.dp.tc/rlp"
)

// lesEntry is the "les" ENR entry. This is set for LES servers only.
type lesEntry struct {
	// Ignore additional fields (for forward compatibility).
	Rest []rlp.RawValue `rlp:"tail"`
}

// ENRKey implements enr.Entry.
func (e lesEntry) ENRKey() string {
	return "les"
}

// setupDiscovery creates the node discovery source for the eth protocol.
func (eth *Lightpc) setupDiscovery() (enode.Iterator, error) {
	if len(eth.config.EthDiscoveryURLs) == 0 {
		return nil, nil
	}
	client := dnsdisc.NewClient(dnsdisc.Config{})
	return client.NewIterator(eth.config.EthDiscoveryURLs...)
}
