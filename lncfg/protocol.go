//go:build !integration

package lncfg

import (
	"github.com/lightningnetwork/lnd/feature"
	"github.com/lightningnetwork/lnd/lnwire"
)

// ProtocolOptions is a struct that we use to be able to test backwards
// compatibility of protocol additions, while defaulting to the latest within
// lnd, or to enable experimental protocol changes.
//
//nolint:lll
type ProtocolOptions struct {
	// LegacyProtocol is a sub-config that houses all the legacy protocol
	// options.  These are mostly used for integration tests as most modern
	// nodes should always run with them on by default.
	LegacyProtocol `group:"legacy" namespace:"legacy"`

	// WumboChans should be set if we want to enable support for wumbo
	// (channels larger than 0.16 BTC) channels, which is the opposite of
	// mini.
	WumboChans bool `long:"wumbo-channels" description:"if set, then lnd will create and accept requests for channels larger chan 0.16 BTC"`

	// TaprootChans should be set if we want to enable support for the
	// experimental simple taproot chans commitment type.
	TaprootChans bool `long:"simple-taproot-chans" description:"if set, then lnd will create and accept requests for channels using the simple taproot commitment type"`

	// NoAnchors should be set if we don't want to support opening or accepting
	// channels having the anchor commitment type.
	NoAnchors bool `long:"no-anchors" description:"disable support for anchor commitments"`

	// NoScriptEnforcedLease should be set if we don't want to support
	// opening or accepting channels having the script enforced commitment
	// type for leased channel.
	NoScriptEnforcedLease bool `long:"no-script-enforced-lease" description:"disable support for script enforced lease commitments"`

	// OptionScidAlias should be set if we want to signal the
	// option-scid-alias feature bit. This allows scid aliases and the
	// option-scid-alias channel-type.
	OptionScidAlias bool `long:"option-scid-alias" description:"enable support for option_scid_alias channels"`

	// OptionZeroConf should be set if we want to signal the zero-conf
	// feature bit.
	OptionZeroConf bool `long:"zero-conf" description:"enable support for zero-conf channels, must have option-scid-alias set also"`

	// NoOptionAnySegwit should be set to true if we don't want to use any
	// Taproot (and beyond) addresses for co-op closing.
	NoOptionAnySegwit bool `long:"no-any-segwit" description:"disallow using any segiwt witness version as a co-op close address"`

	// CustomMessage allows the custom message APIs to handle messages with
	// the provided protocol numbers, which fall outside the custom message
	// number range.
	CustomMessage []uint16 `long:"custom-message" description:"allows the custom message apis to send and report messages with the protocol number provided that fall outside of the custom message number range."`

	// CustomInit specifies feature bits to advertise in the node's init
	// message.
	CustomInit []uint16 `long:"custom-init" description:"custom feature bits to advertise in the node's init message"`

	// CustomNodeAnn specifies custom feature bits to advertise in the
	// node's announcement message.
	CustomNodeAnn []uint16 `long:"custom-nodeann" description:"custom feature bits to advertise in the node's announcement message"`

	// CustomInvoice specifies custom feature bits to advertise in the
	// node's invoices.
	CustomInvoice []uint16 `long:"custom-invoice" description:"custom feature bits to advertise in the node's invoices"`
}

// Wumbo returns true if lnd should permit the creation and acceptance of wumbo
// channels.
func (l *ProtocolOptions) Wumbo() bool {
	return l.WumboChans
}

// NoAnchorCommitments returns true if we have disabled support for the anchor
// commitment type.
func (l *ProtocolOptions) NoAnchorCommitments() bool {
	return l.NoAnchors
}

// NoScriptEnforcementLease returns true if we have disabled support for the
// script enforcement commitment type for leased channels.
func (l *ProtocolOptions) NoScriptEnforcementLease() bool {
	return l.NoScriptEnforcedLease
}

// ScidAlias returns true if we have enabled the option-scid-alias feature bit.
func (l *ProtocolOptions) ScidAlias() bool {
	return l.OptionScidAlias
}

// ZeroConf returns true if we have enabled the zero-conf feature bit.
func (l *ProtocolOptions) ZeroConf() bool {
	return l.OptionZeroConf
}

// NoAnySegwit returns true if we don't signal that we understand other newer
// segwit witness versions for co-op close addresses.
func (l *ProtocolOptions) NoAnySegwit() bool {
	return l.NoOptionAnySegwit
}

// CustomMessageOverrides returns the set of protocol messages that we override
// to allow custom handling.
func (p ProtocolOptions) CustomMessageOverrides() []uint16 {
	return p.CustomMessage
}

// CustomFeatures returns a custom set of feature bits to advertise.
//
//nolint:lll
func (p ProtocolOptions) CustomFeatures() map[feature.Set][]lnwire.FeatureBit {
	customFeatures := make(map[feature.Set][]lnwire.FeatureBit)

	setFeatures := func(set feature.Set, bits []uint16) {
		for _, customFeature := range bits {
			customFeatures[set] = append(
				customFeatures[set],
				lnwire.FeatureBit(customFeature),
			)
		}
	}

	setFeatures(feature.SetInit, p.CustomInit)
	setFeatures(feature.SetNodeAnn, p.CustomNodeAnn)
	setFeatures(feature.SetInvoice, p.CustomInvoice)

	return customFeatures
}
