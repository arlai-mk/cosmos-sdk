package keeper

import (
	errorsmod "cosmossdk.io/errors"
	storetypes "cosmossdk.io/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/gov/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	v1 "github.com/cosmos/cosmos-sdk/x/gov/types/v1"
)

// get a single representative
func (k Keeper) GetRepresentative(ctx sdk.Context, representativeID uint64) (representative v1.Representative, found bool) {
	store := ctx.KVStore(k.storeKey)

	value := store.Get(govtypes.RepresentativeKey(representativeID))
	if value == nil {
		return representative, false
	}

	representative = v1.MustUnmarshalRepresentative(k.cdc, value)
	return representative, true
}

func (k Keeper) GetRepresentativeIDByAddr(ctx sdk.Context, repAddr sdk.AccAddress) (representativeID uint64, found bool) {
	store := ctx.KVStore(k.storeKey)

	representativeIDBytes := store.Get(govtypes.RepresentativeIDByAddrKey(repAddr))
	if representativeIDBytes == nil {
		return representativeID, false
	}

	return govtypes.GetRepresentativeIDFromBytes(representativeIDBytes), true
}

func (k Keeper) GetRepresentativeByAddr(ctx sdk.Context, repAddr sdk.AccAddress) (representative v1.Representative, found bool) {
	representativeID, found := k.GetRepresentativeIDByAddr(ctx, repAddr)
	if found == false {
		return representative, false
	}

	return k.GetRepresentative(ctx, representativeID)
}

// set the main record holding representative details
func (k Keeper) SetRepresentative(ctx sdk.Context, representative v1.Representative) (uint64, error) {
	store := ctx.KVStore(k.storeKey)

	// Get next ID
	nextRepresentativeID, err := k.GetRepresentativeID(ctx)
	if err != nil {
		return 0, err
	}

	bz := v1.MustMarshalRepresentative(k.cdc, &representative)

	store.Set(govtypes.RepresentativeKey(nextRepresentativeID), bz)

	// Increase next ID
	k.SetRepresentativeID(ctx, nextRepresentativeID+1)

	return nextRepresentativeID, nil
}

// validator index
func (k Keeper) SetRepresentativeIDByAddr(ctx sdk.Context, representativeID uint64, representative v1.Representative) error {
	store := ctx.KVStore(k.storeKey)
	repAddr := representative.GetRepAddress()
	store.Set(govtypes.RepresentativeIDByAddrKey(repAddr), govtypes.GetRepresentativeIDBytes(representativeID))
	return nil
}

// GetRepresentativeID gets the highest representative ID
func (keeper Keeper) GetRepresentativeID(ctx sdk.Context) (representativeID uint64, err error) {
	store := ctx.KVStore(keeper.storeKey)
	bz := store.Get(govtypes.RepresentativeIDKey)
	if bz == nil {
		return 0, errorsmod.Wrap(types.ErrInvalidGenesis, "initial representative ID hasn't been set")
	}

	representativeID = govtypes.GetRepresentativeIDFromBytes(bz)
	return representativeID, nil
}

// SetRepresentativeID sets the new representative ID to the store
func (keeper Keeper) SetRepresentativeID(ctx sdk.Context, representativeID uint64) {
	store := ctx.KVStore(keeper.storeKey)
	store.Set(govtypes.RepresentativeIDKey, govtypes.GetRepresentativeIDBytes(representativeID))
}

// get the set of all representatives with no limits, used during genesis dump
func (k Keeper) GetAllRepresentatives(ctx sdk.Context) (representatives v1.Representatives) {
	store := ctx.KVStore(k.storeKey)

	iterator := storetypes.KVStorePrefixIterator(store, types.RepresentativesKeyPrefix)
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		representative := v1.MustUnmarshalRepresentative(k.cdc, iterator.Value())
		representatives = append(representatives, &representative)
	}

	return representatives
}
