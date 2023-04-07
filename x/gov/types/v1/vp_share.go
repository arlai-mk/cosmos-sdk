package v1

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// VotingPowerRepChange is used to track changes in voting power for a representative
type VotingPowerRepChange struct {
	RepresentativeAddress string
	ValidatorAddress      string
	VotingPowerChange     sdk.Int
}

// NewVotingPowerDelShares constructs a new VotingPowerDelShares
//
//nolint:interfacer
func NewVotingPowerDelShares(delegator sdk.AccAddress, vpShares []VotingPowerDelShare) (VotingPowerDelShares, error) {
	return VotingPowerDelShares{
		DelegatorAddress: delegator.String(),
		VpShare:          vpShares,
	}, nil
}

// NewVotingPowerDelShare constructs a new VotingPowerDelShares
//
//nolint:interfacer
func NewVotingPowerDelShare(representativeID uint64, percentage sdk.Dec) (VotingPowerDelShare, error) {
	return VotingPowerDelShare{
		RepresentativeId: representativeID,
		Percentage:       percentage,
	}, nil
}

// return the representative as store value
func MustMarshalVotingPowerDelShares(cdc codec.BinaryCodec, vpDelShares *VotingPowerDelShares) []byte {
	return cdc.MustMarshal(vpDelShares)
}

// unmarshal a representative from a store value
func MustUnmarshalVotingPowerDelShares(cdc codec.BinaryCodec, value []byte) VotingPowerDelShares {
	vpDelShares, err := UnmarshalVotingPowerDelShares(cdc, value)
	if err != nil {
		panic(err)
	}

	return vpDelShares
}

// unmarshal a representative from a store value
func UnmarshalVotingPowerDelShares(cdc codec.BinaryCodec, value []byte) (v VotingPowerDelShares, err error) {
	err = cdc.Unmarshal(value, &v)
	return v, err
}

// NewVotingPowerRep constructs a new VotingPowerRep
//
//nolint:interfacer
func NewVotingPowerRep(
	representativeID uint64,
	validator sdk.ValAddress,
	votingPower sdk.Int,
) (VotingPowerRep, error) {
	return VotingPowerRep{
		RepresentativeId: representativeID,
		ValidatorAddress: validator.String(),
		VotingPower:      votingPower,
	}, nil
}

// return the representative as store value
func MustMarshalVotingPowerRep(cdc codec.BinaryCodec, vpRep *VotingPowerRep) []byte {
	return cdc.MustMarshal(vpRep)
}

// unmarshal a representative from a store value
func MustUnmarshalVotingPowerRep(cdc codec.BinaryCodec, value []byte) VotingPowerRep {
	vpRep, err := UnmarshalVotingPowerRep(cdc, value)
	if err != nil {
		panic(err)
	}

	return vpRep
}

// unmarshal a representative from a store value
func UnmarshalVotingPowerRep(cdc codec.BinaryCodec, value []byte) (v VotingPowerRep, err error) {
	err = cdc.Unmarshal(value, &v)
	return v, err
}
