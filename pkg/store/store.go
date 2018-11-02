/*
 * Copyright 2018 COSMOPlat-DEV
 *
 * This program is free software; you can redistribute and/or modify it
 * under the terms of the standard MIT license. See LICENSE for more details
 */

package store

import (
	"github.com/cosmoplat-dev/anchor/pkg/utils"
	"net"
)

// Store is the store interface for anchor
type Store interface {
	Lock() error
	Unlock() error
	Close() error
	Reserve(id string, ip net.IP, podName string, podNamespace string, controller string) (bool, error)
	Release(id string) error

	RetrieveGateway(subnet *net.IPNet) net.IP // return nil if error
	RetrieveAllocated(namespace string, subnet *net.IPNet) (*utils.RangeSet, error)
	RetrieveUsed(namespace string, subnet *net.IPNet) (*utils.RangeSet, error)
}
