package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/cosmos/ibc-go/v3/modules/apps/transfer/types"
)

// GetSendEnabled retrieves the send enabled boolean from the paramstore
func (k Keeper) GetSendEnabled(ctx sdk.Context) bool {
	var res bool
	k.paramSpace.Get(ctx, types.KeySendEnabled, &res)
	return res
}

// GetReceiveEnabled retrieves the receive enabled boolean from the paramstore
func (k Keeper) GetReceiveEnabled(ctx sdk.Context) bool {
	var res bool
	k.paramSpace.Get(ctx, types.KeyReceiveEnabled, &res)
	return res
}

// GetReceiveEnabled retrieves the receive enabled boolean from the paramstore
func (k Keeper) GetAllowedAddresses(ctx sdk.Context) []string {
	var res []string
	k.paramSpace.Get(ctx, types.KeyAllowedAddresses, &res)
	return res
}

func (k Keeper) IsAddressEnabled(ctx sdk.Context, sender sdk.Address) bool {
	str := sender.String()
	for _, v := range k.GetAllowedAddresses(ctx) {
		if v == str {
			return true
		}
	}
	return false
}

// GetParams returns the total set of ibc-transfer parameters.
func (k Keeper) GetParams(ctx sdk.Context) types.Params {
	return types.NewParams(k.GetSendEnabled(ctx), k.GetReceiveEnabled(ctx), k.GetAllowedAddresses(ctx))
}

// SetParams sets the total set of ibc-transfer parameters.
func (k Keeper) SetParams(ctx sdk.Context, params types.Params) {
	k.paramSpace.SetParamSet(ctx, &params)
}
