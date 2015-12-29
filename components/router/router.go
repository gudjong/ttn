// Copyright © 2015 The Things Network
// Use of this source code is governed by the MIT license that can be found in the LICENSE file.

package router

import (
	"github.com/thethingsnetwork/core"
	"github.com/thethingsnetwork/core/lorawan/semtech"
	"time"
)

type Router struct{}

func (r *Router) HandleUplink(packet semtech.Packet, connId core.ConnectionId) {
	/* PULL_DATA
	 *
	 * Send PULL_ACK
	 * Store the gateway in known gateway
	 */

	/* PUSH_DATA
	 *
	 * Send PUSH_ACK
	 * Stores the gateway connection id for later response
	 * Lookup for an existing broker associated to the device address
	 * Forward data to that broker
	 */

	/* Else
	 *
	 * Ignore / Raise an error
	 */

}

func (r *Router) HandleDownlink(packet semtech.Packet) {

}

func (r *Router) RegisterDevice(devAddr core.DeviceAddress, broAddrs ...core.BrokerAddress) {

}

func (r *Router) HandleError(err error) {

}

// --------------- Address keeper
type addressKeeper interface {
	lookup(devAddr core.DeviceAddress) (core.BrokerAddress, error)
	invalidate(broAddrs ...core.BrokerAddress) error
}

type reddisAddressKeeper struct{} // In a second time

type localAddressKeeper struct {
	addresses map[core.DeviceAddress]time.Time
}

func (a *localAddressKeeper) lookup(devAddr core.DeviceAddress) (*core.BrokerAddress, error) {
	return nil, nil
}

func (a *localAddressKeeper) invalidate(broAddrs ...core.BrokerAddress) error {
	return nil
}

// --------------- Routers Adapters

type UpAdapter struct {
}

func (u *UpAdapter) Ack(packet semtech.Packet, gid core.GatewayId) {
}

type DownAdapter struct {
}

func (d *DownAdapter) Broadcast(packet semtech.Packet) {

}

func (d *DownAdapter) Forward(packet semtech.Packet, broAddrs ...core.BrokerAddress) {

}
