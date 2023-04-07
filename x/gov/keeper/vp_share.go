package keeper

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	v1 "github.com/cosmos/cosmos-sdk/x/gov/types/v1"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
)

// GetVotingPowerDelShares retrieves the Voting Power shares by the delegator
func (k Keeper) GetVotingPowerDelShares(ctx sdk.Context, delegator sdk.AccAddress) (vpShares v1.VotingPowerDelShares, found bool) {
	store := ctx.KVStore(k.storeKey)

	value := store.Get(govtypes.VotingPowerSharesKey(delegator))
	if value == nil {
		return vpShares, false
	}

	vpShares = v1.MustUnmarshalVotingPowerDelShares(k.cdc, value)
	return vpShares, true
}

func (k Keeper) SetVotingPowerDelShares(ctx sdk.Context, vpShares v1.VotingPowerDelShares) error {
	store := ctx.KVStore(k.storeKey)

	bz := v1.MustMarshalVotingPowerDelShares(k.cdc, &vpShares)
	store.Set(govtypes.VotingPowerSharesKey(sdk.AccAddress(vpShares.DelegatorAddress)), bz)

	return nil
}

func (keeper Keeper) convertVPSharesBasicToFinal(ctx sdk.Context, vpSharesBasic []*v1.VotingPowerDelShareBasic) ([]v1.VotingPowerDelShare, error) {

	vpShares := make([]v1.VotingPowerDelShare, len(vpSharesBasic))

	for idx, vpShareBasic := range vpSharesBasic {
		repAddr, err := sdk.AccAddressFromBech32(vpShareBasic.RepresentativeAddress)
		if err != nil {
			return []v1.VotingPowerDelShare{}, err
		}
		representativeID, found := keeper.GetRepresentativeIDByAddr(ctx, repAddr)
		if !found {
			return []v1.VotingPowerDelShare{}, fmt.Errorf("Cannot find the representative with address %s", vpShareBasic.RepresentativeAddress)
		}

		_, found = keeper.GetRepresentative(ctx, representativeID)
		if !found {
			return []v1.VotingPowerDelShare{}, fmt.Errorf("Cannot find the representative with ID %d", representativeID)
		}

		vpShares[idx] = v1.VotingPowerDelShare{
			RepresentativeId: representativeID,
			Percentage:       vpShareBasic.Percentage,
		}
	}

	return vpShares, nil
}

type VPRepChangeByAddr map[string]v1.VotingPowerRepChange

// calculateVPRepChanges calculates the voting power change for representatives if a delegator changed their representative shares
//
// - first use calculateVPRepChangeUnstakeAll
//
// - then call calculateVPRepChangeStakeAll
//
// - then call combineVPRepChangesByAddr to combine the VPRepChangeByAddr
func (keeper Keeper) calculateVPRepChanges(ctx sdk.Context, delegator sdk.AccAddress, newVpShares v1.VotingPowerDelShares) ([]v1.VotingPowerRepChange, error) {
	vpRepChanges1, _ := keeper.calculateVPRepChangeUnstakeAll(ctx, delegator)

	vpRepChanges2, _ := keeper.calculateVPRepChangeStakeAll(ctx, delegator, newVpShares)

	vpRepChange := combineVPRepChangesByAddr(vpRepChanges1, vpRepChanges2)

	return vpRepChange, nil
}

// calculateVPRepChangeStake calculates the voting power change for representatives if a delegator stakes with a validator
func (keeper Keeper) calculateVPRepChangeStake(ctx sdk.Context, delegator sdk.AccAddress, validator sdk.ValAddress, amount sdk.Int) (VPRepChangeByAddr, error) {
	vpRepChange := make(VPRepChangeByAddr)
	// TODO: code the function

	return vpRepChange, nil
}

// calculateVPRepChangeUnstake calculates the voting power change for representatives if a delegator unstakes from a validator
func (keeper Keeper) calculateVPRepChangeUnstake(ctx sdk.Context, delegator sdk.AccAddress, validator sdk.ValAddress, amount sdk.Int) (VPRepChangeByAddr, error) {
	vpRepChange := make(VPRepChangeByAddr)
	// TODO: code the function

	return vpRepChange, nil
}

func combineVPRepChangesByAddr(vpRepChanges1 VPRepChangeByAddr, vpRepChanges2 VPRepChangeByAddr) []v1.VotingPowerRepChange {
	vpRepChanges := make([]v1.VotingPowerRepChange, 0)

	// index is the concatenation of representative address and validator address
	for index, vpRepChange1 := range vpRepChanges1 {
		if vpRepChange2, found := vpRepChanges2[index]; found {
			vpRepChanges = append(vpRepChanges, v1.VotingPowerRepChange{
				RepresentativeAddress: vpRepChange1.RepresentativeAddress,
				ValidatorAddress:      vpRepChange2.ValidatorAddress,
				VotingPowerChange:     vpRepChange1.VotingPowerChange.Add(vpRepChange2.VotingPowerChange),
			})
		} else {
			vpRepChanges = append(vpRepChanges, vpRepChange1)
		}
	}

	for index, vpRepChange2 := range vpRepChanges2 {
		if _, found := vpRepChanges1[index]; !found {
			vpRepChanges = append(vpRepChanges, vpRepChange2)
		}
	}

	return vpRepChanges
}

func (keeper Keeper) GetVotingPowerByValidator(ctx sdk.Context, delegator sdk.AccAddress) (map[string]sdk.Dec, map[string]sdk.ValAddress) {
	vpByValidator := make(map[string]sdk.Dec)
	valByAddrStr := make(map[string]sdk.ValAddress)

	// iterate over all delegations from delegator
	keeper.sk.IterateDelegations(ctx, delegator, func(index int64, delegation stakingtypes.DelegationI) (stop bool) {
		validator := keeper.sk.Validator(ctx, delegation.GetValidatorAddr())
		valOper := validator.GetOperator()

		votingPower := delegation.GetShares().MulInt(validator.GetTokens()).Quo(validator.GetDelegatorShares())

		vpByValidator[valOper.String()] = votingPower
		valByAddrStr[valOper.String()] = valOper

		return false
	})

	return vpByValidator, valByAddrStr
}

// calculateVPRepChangeUnstakeAll calculates the voting power change for representatives if a delegator was unstaking all their tokens
//
// - If there is no VotingPowerDelShares then unstaking means removing from validators
//
// - If there is VotingPowerDelShares then unstaking means removing from representatives based on the shares
func (keeper Keeper) calculateVPRepChangeUnstakeAll(ctx sdk.Context, delegator sdk.AccAddress) (VPRepChangeByAddr, error) {

	vpRepChange := make(VPRepChangeByAddr)

	vpShares, hasVPShares := keeper.GetVotingPowerDelShares(ctx, delegator)

	vpByValidator, valByAddrStr := keeper.GetVotingPowerByValidator(ctx, delegator)

	if !hasVPShares {
		for valAddrStr, votingPower := range vpByValidator {
			valOper := valByAddrStr[valAddrStr]
			repAddr := sdk.AccAddress(valOper)

			vpRepChange[repAddr.String()+valOper.String()] = v1.VotingPowerRepChange{
				RepresentativeAddress: repAddr.String(),
				ValidatorAddress:      valOper.String(),
				VotingPowerChange:     votingPower.Mul(sdk.NewDec(-1)).TruncateInt(),
			}
		}
	} else {
		for _, vpShare := range vpShares.VpShare {
			representative, found := keeper.GetRepresentative(ctx, vpShare.RepresentativeId)
			// This should never happen, but still don't want to panic here
			if !found {
				continue
			}

			repAddr := representative.GetRepAddress()
			for valAddrStr, votingPower := range vpByValidator {
				vpRepChange[repAddr.String()+valAddrStr] = v1.VotingPowerRepChange{
					RepresentativeAddress: repAddr.String(),
					ValidatorAddress:      valAddrStr,
					VotingPowerChange:     votingPower.Mul(vpShare.Percentage).Mul(sdk.NewDec(-1)).TruncateInt(),
				}
			}
		}
	}

	return vpRepChange, nil
}

func (keeper Keeper) calculateVPRepChangeStakeAll(
	ctx sdk.Context,
	delegator sdk.AccAddress,
	newVpShares v1.VotingPowerDelShares,
) (VPRepChangeByAddr, error) {

	vpRepChange := make(VPRepChangeByAddr)

	vpByValidator, _ := keeper.GetVotingPowerByValidator(ctx, delegator)

	for _, vpShare := range newVpShares.VpShare {
		representative, found := keeper.GetRepresentative(ctx, vpShare.RepresentativeId)
		// This should never happen, but still don't want to panic here
		if !found {
			continue
		}

		repAddr := representative.GetRepAddress()
		for valAddrStr, votingPower := range vpByValidator {
			vpRepChange[repAddr.String()+valAddrStr] = v1.VotingPowerRepChange{
				RepresentativeAddress: repAddr.String(),
				ValidatorAddress:      valAddrStr,
				VotingPowerChange:     votingPower.Mul(vpShare.Percentage).TruncateInt(),
			}
		}
	}

	return vpRepChange, nil
}

// The function applyVPRepChanges will update the VotingPowerRep in the store
// based on the list of VotingPowerRepChange
func (keeper Keeper) applyVPRepChanges(ctx sdk.Context, vpRepChanges []v1.VotingPowerRepChange) {

	store := ctx.KVStore(keeper.storeKey)

	for _, vpRepChange := range vpRepChanges {
		// get the representative ID
		representativeID, found := keeper.GetRepresentativeIDByAddr(ctx, sdk.AccAddress(vpRepChange.RepresentativeAddress))

		// This should never happen, but still don't want to panic here
		if !found {
			continue
		}

		// check if there's a VotingPowerRep for that representative / validator
		value := store.Get(govtypes.VotingPowerKey(representativeID, sdk.ValAddress(vpRepChange.ValidatorAddress)))

		var vpPowerRep v1.VotingPowerRep
		// if nothing, then we just save the VotingPowerRep
		if value == nil {
			vpPowerRep, _ = v1.NewVotingPowerRep(
				representativeID,
				sdk.ValAddress(vpRepChange.ValidatorAddress),
				vpRepChange.VotingPowerChange,
			)
			// store.Set
		} else { // if already something, then we adjust and save
			vpPowerRep = v1.MustUnmarshalVotingPowerRep(keeper.cdc, value)

			// Add the vpRepChange to the current VotingPowerRep
			vpPowerRep.VotingPower = vpPowerRep.VotingPower.Add(vpRepChange.VotingPowerChange)
		}

		// we need to make sure it cannot go negative
		if vpPowerRep.VotingPower.LT(sdk.ZeroInt()) {
			panic("the voting power can never be negative")
		}

		bz := v1.MustMarshalVotingPowerRep(keeper.cdc, &vpPowerRep)

		store.Set(govtypes.VotingPowerKey(representativeID, sdk.ValAddress(vpRepChange.ValidatorAddress)), bz)
	}
}
