package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/commercionetwork/ibc-go/v3/modules/core/exported"
)

// ClientKeeper expected account IBC client keeper
type ClientKeeper interface {
	GetClientState(ctx sdk.Context, clientID string) (exported.ClientState, bool)
	GetClientConsensusState(ctx sdk.Context, clientID string, height exported.Height) (exported.ConsensusState, bool)
	GetSelfConsensusState(ctx sdk.Context, height exported.Height) (exported.ConsensusState, error)
	ValidateSelfClient(ctx sdk.Context, clientState exported.ClientState) error
	IterateClients(ctx sdk.Context, cb func(string, exported.ClientState) bool)
	ClientStore(ctx sdk.Context, clientID string) sdk.KVStore
}
