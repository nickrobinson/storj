// Copyright (C) 2019 Storj Labs, Inc.
// See LICENSE for copying information.

package kademlia

import (
	"storj.io/storj/pkg/pb"
)

func (rt *RoutingTable) addToReplacementCache(kadBucketID bucketID, node *pb.Node) {
	rt.rcMutex.Lock()
	defer rt.rcMutex.Unlock()
	nodes := rt.replacementCache[kadBucketID]
	nodes = append(nodes, node)

	if len(nodes) > rt.rcBucketSize {
		copy(nodes, nodes[1:])
		nodes = nodes[:len(nodes)-1]
	}
	rt.replacementCache[kadBucketID] = nodes
}

func (rt *RoutingTable) removeFromReplacementCache(kadBucketID bucketID, node *pb.Node) {
	rt.rcMutex.Lock()
	defer rt.rcMutex.Unlock()
	nodes := rt.replacementCache[kadBucketID]
	for i, n := range nodes {
		if n.Id == node.Id && n.Address.GetAddress() == node.Address.GetAddress() {
			nodes = append(nodes[:i], nodes[i+1:]...)
			break
		}
	}
	rt.replacementCache[kadBucketID] = nodes
}
