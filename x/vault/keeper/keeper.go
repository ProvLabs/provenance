package keeper

import (
	"fmt"

	"cosmossdk.io/log"
	storetypes "cosmossdk.io/store/types"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/gogoproto/proto"

	"github.com/provenance-io/provenance/x/vault/types"
)

// Keeper provides the vault module's core state handling functionality.
type Keeper struct {
	cdc      codec.BinaryCodec
	storeKey storetypes.StoreKey
}

// NewKeeper creates a new vault Keeper.
func NewKeeper(cdc codec.BinaryCodec, storeKey storetypes.StoreKey) Keeper {
	return Keeper{
		cdc:      cdc,
		storeKey: storeKey,
	}
}

// getStore returns the module's KVStore.
func (k Keeper) getStore(ctx sdk.Context) storetypes.KVStore {
	return ctx.KVStore(k.storeKey)
}

// getLogger returns a logger with vault module context.
func (k Keeper) getLogger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", "x/"+types.ModuleName)
}

// emitEvent emits the provided event and writes any error to the error log.
// If you have multiple events to emit, consider using emitEvents.
func (k Keeper) emitEvent(ctx sdk.Context, event proto.Message) {
	err := ctx.EventManager().EmitTypedEvent(event)
	if err != nil {
		k.logErrorf(ctx, "error emitting event %#v: %v", event, err)
	}
}

// logErrorf uses fmt.Sprintf to combine the msg and args, and logs the result as an error from this module.
// Note that this is different from the logging .Error(msg string, keyvals ...interface{}) syntax.
func (k Keeper) logErrorf(ctx sdk.Context, msg string, args ...interface{}) {
	k.getLogger(ctx).Error(fmt.Sprintf(msg, args...))
}
