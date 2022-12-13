package types

import (
	"fmt"

	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

const (
	// DefaultSendEnabled enabled
	DefaultSendEnabled = true
	// DefaultReceiveEnabled enabled
	DefaultReceiveEnabled = true
)

var (
	// KeySendEnabled is store's key for SendEnabled Params
	KeySendEnabled = []byte("SendEnabled")
	// KeyReceiveEnabled is store's key for ReceiveEnabled Params
	KeyReceiveEnabled = []byte("ReceiveEnabled")
	// KeyReceiveEnabled is store's key for ReceiveEnabled Params
	KeyAllowedAddresses = []byte("AllowedAddresses")
)

// ParamKeyTable type declaration for parameters
func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

// NewParams creates a new parameter configuration for the ibc transfer module
func NewParams(enableSend, enableReceive bool, allowedAddresses []string) Params {
	return Params{
		SendEnabled:      enableSend,
		ReceiveEnabled:   enableReceive,
		AllowedAddresses: allowedAddresses,
	}
}

// DefaultParams is the default parameter configuration for the ibc-transfer module
func DefaultParams() Params {
	var DefaultAllowedAddresses []string
	return NewParams(DefaultSendEnabled, DefaultReceiveEnabled, DefaultAllowedAddresses)
}

// Validate all ibc-transfer module parameters
func (p Params) Validate() error {
	if err := validateEnabled(p.SendEnabled); err != nil {
		return err
	}

	return validateEnabled(p.ReceiveEnabled)
}

// ParamSetPairs implements params.ParamSet
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{
		paramtypes.NewParamSetPair(KeySendEnabled, p.SendEnabled, validateEnabled),
		paramtypes.NewParamSetPair(KeyReceiveEnabled, p.ReceiveEnabled, validateEnabled),
		paramtypes.NewParamSetPair(KeyAllowedAddresses, p.AllowedAddresses, validateSlice),
	}
}

func validateEnabled(i interface{}) error {
	_, ok := i.(bool)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	return nil
}

func validateSlice(i interface{}) error {
	_, ok := i.([]string)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	return nil
}
