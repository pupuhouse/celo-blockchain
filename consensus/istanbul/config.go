// Copyright 2017 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package istanbul

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/p2p/enode"
)

type ProposerPolicy uint64

const (
	RoundRobin ProposerPolicy = iota
	Sticky
	ShuffledRoundRobin
)

type FaultyMode uint64

const (
	// Disabled disables the faulty mode
	Disabled FaultyMode = iota
	// Random will, at the time of any of the following actions, engage in faulty behavior with 50% probability
	Random
	// NotBroadcast doesn't broadcast any messages to other validators
	NotBroadcast
	// SendWrongMsg sends the message with the wrong message code
	SendWrongMsg
	// ModifySig modifies the message signature
	ModifySig
	// AlwaysPropose always proposes a proposal to validators
	AlwaysPropose
	// AlwaysRoundChange always sends round change while receiving messages
	AlwaysRoundChange
	// BadBlock always proposes a block with bad body
	BadBlock
	// SendExtraMessages sends multiple copies of the same message
	SendExtraMessages
	// SendExtraFutureMessages sends multiple copies of a prepare message with a future sequence and round number
	SendExtraFutureMessages
)

func (f FaultyMode) Uint64() uint64 {
	return uint64(f)
}

func (f FaultyMode) String() string {
	switch f {
	case Disabled:
		return "Disabled"
	case Random:
		return "Random"
	case NotBroadcast:
		return "NotBroadcast"
	case SendWrongMsg:
		return "SendWrongMsg"
	case ModifySig:
		return "ModifySig"
	case AlwaysPropose:
		return "AlwaysPropose"
	case AlwaysRoundChange:
		return "AlwaysRoundChange"
	case BadBlock:
		return "BadBlock"
	case SendExtraMessages:
		return "SendExtraMessages"
	case SendExtraFutureMessages:
		return "SendExtraFutureMessages"
	default:
		return "Undefined"
	}
}

type Config struct {
	RequestTimeout              uint64         `toml:",omitempty"` // The timeout for each Istanbul round in milliseconds.
	TimeoutBackoffFactor        uint64         `toml:",omitempty"` // Timeout at subsequent rounds is: RequestTimeout + 2**round * TimeoutBackoffFactor (in milliseconds)
	MinResendRoundChangeTimeout uint64         `toml:",omitempty"` // Minimum interval with which to resend RoundChange messages for same round
	MaxResendRoundChangeTimeout uint64         `toml:",omitempty"` // Maximum interval with which to resend RoundChange messages for same round
	BlockPeriod                 uint64         `toml:",omitempty"` // Default minimum difference between two consecutive block's timestamps in second
	ProposerPolicy              ProposerPolicy `toml:",omitempty"` // The policy for proposer selection
	FaultyMode                  uint64         `toml:",omitempty"` // The faulty node indicates the faulty node's behavior
	Epoch                       uint64         `toml:",omitempty"` // The number of blocks after which to checkpoint and reset the pending votes
	LookbackWindow              uint64         `toml:",omitempty"` // The window of blocks in which a validator is forgived from voting
	ValidatorEnodeDBPath        string         `toml:",omitempty"` // The location for the validator enodes DB
	VersionCertificateDBPath    string         `toml:",omitempty"` // The location for the signed announce version DB
	RoundStateDBPath            string         `toml:",omitempty"` // The location for the round states DB
	Validator                   bool           `toml:",omitempty"` // Specified if this node is configured to validate (specifically if --mine command line is set)

	// Proxy Configs
	Proxy                   bool           `toml:",omitempty"` // Specifies if this node is a proxy
	ProxiedValidatorAddress common.Address `toml:",omitempty"` // The address of the proxied validator

	// Proxied Validator Configs
	Proxied                 bool        `toml:",omitempty"` // Specifies if this node is proxied
	ProxyInternalFacingNode *enode.Node `toml:",omitempty"` // The internal facing node of the proxy that this proxied validator will contect to
	ProxyExternalFacingNode *enode.Node `toml:",omitempty"` // The external facing node of the proxy that the proxied validator will broadcast via the announce message

	// Announce Configs
	AnnounceQueryEnodeGossipPeriod                 uint64 `toml:",omitempty"` // Time duration (in seconds) between gossiped query enode messages
	AnnounceAggressiveQueryEnodeGossipOnEnablement bool   `toml:",omitempty"` // Specifies if this node should aggressively query enodes on announce enablement
	AnnounceAdditionalValidatorsToGossip           int64  `toml:",omitempty"` // Specifies the number of additional non-elected validators to gossip an announce
}

var DefaultConfig = &Config{
	RequestTimeout:                 3000,
	TimeoutBackoffFactor:           1000,
	MinResendRoundChangeTimeout:    15 * 1000,
	MaxResendRoundChangeTimeout:    2 * 60 * 1000,
	BlockPeriod:                    5,
	ProposerPolicy:                 ShuffledRoundRobin,
	FaultyMode:                     Disabled.Uint64(),
	Epoch:                          30000,
	LookbackWindow:                 12,
	ValidatorEnodeDBPath:           "validatorenodes",
	VersionCertificateDBPath:       "versioncertificates",
	RoundStateDBPath:               "roundstates",
	Validator:                      false,
	Proxy:                          false,
	Proxied:                        false,
	AnnounceQueryEnodeGossipPeriod: 300, // 5 minutes
	AnnounceAggressiveQueryEnodeGossipOnEnablement: true,
	AnnounceAdditionalValidatorsToGossip:           10,
}
