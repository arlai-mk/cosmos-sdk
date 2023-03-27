package v1

import (
	"sigs.k8s.io/yaml"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	// DefaultStartingRepresentativeID is 1
	DefaultStartingRepresentativeID uint64 = 1

	// TODO: Why can't we just have one string description which can be JSON by convention
	MaxMonikerLength  = 70
	MaxIdentityLength = 3000
	MaxWebsiteLength  = 140
	MaxDetailsLength  = 280
)

// ValidatorI expected validator functions
type RepresentativeI interface {
	GetMoniker() string            // moniker of the representative
	GetRepAddress() sdk.AccAddress // account address of the representative
}

var _ RepresentativeI = Representative{}

// NewRepresentative constructs a new Representative
//
//nolint:interfacer
func NewRepresentative(repAddr sdk.AccAddress, description RepDescription) (Representative, error) {
	return Representative{
		RepresentativeAddress: repAddr.String(),
		Description:           description,
	}, nil
}

// Representatives is a collection of Representative objects
type Representatives []*Representative

// String implements the Stringer interface for a Representative object.
func (r Representative) String() string {
	bz, err := codec.ProtoMarshalJSON(&r, nil)
	if err != nil {
		panic(err)
	}

	out, err := yaml.JSONToYAML(bz)
	if err != nil {
		panic(err)
	}

	return string(out)
}

func NewRepDescription(moniker, identity, website, details string) RepDescription {
	return RepDescription{
		Moniker:  moniker,
		Identity: identity,
		Website:  website,
		Details:  details,
	}
}

// String implements the Stringer interface for a RepDescription object.
func (d RepDescription) String() string {
	out, _ := yaml.Marshal(d)
	return string(out)
}

// Equal checks if the receiver equals the parameter
func (r *Representative) Equal(r2 *Representative) bool {
	return r.RepresentativeAddress == r2.RepresentativeAddress &&
		r.Description.Equal(r2.Description)
}

func (r Representative) GetMoniker() string { return r.Description.Moniker }
func (r Representative) GetRepAddress() sdk.AccAddress {
	if r.RepresentativeAddress == "" {
		return nil
	}
	return sdk.MustAccAddressFromBech32(r.RepresentativeAddress)
}

// return the representative as store value
func MustMarshalRepresentative(cdc codec.BinaryCodec, representative *Representative) []byte {
	return cdc.MustMarshal(representative)
}

// unmarshal a representative from a store value
func MustUnmarshalRepresentative(cdc codec.BinaryCodec, value []byte) Representative {
	representative, err := UnmarshalRepresentative(cdc, value)
	if err != nil {
		panic(err)
	}

	return representative
}

// unmarshal a representative from a store value
func UnmarshalRepresentative(cdc codec.BinaryCodec, value []byte) (r Representative, err error) {
	err = cdc.Unmarshal(value, &r)
	return r, err
}
