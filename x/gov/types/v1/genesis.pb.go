// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: cosmos/gov/v1/genesis.proto

package v1

import (
	fmt "fmt"
	proto "github.com/cosmos/gogoproto/proto"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// GenesisState defines the gov module's genesis state.
type GenesisState struct {
	// starting_proposal_id is the ID of the starting proposal.
	StartingProposalId uint64 `protobuf:"varint,1,opt,name=starting_proposal_id,json=startingProposalId,proto3" json:"starting_proposal_id,omitempty"`
	// deposits defines all the deposits present at genesis.
	Deposits []*Deposit `protobuf:"bytes,2,rep,name=deposits,proto3" json:"deposits,omitempty"`
	// votes defines all the votes present at genesis.
	Votes []*Vote `protobuf:"bytes,3,rep,name=votes,proto3" json:"votes,omitempty"`
	// proposals defines all the proposals present at genesis.
	Proposals []*Proposal `protobuf:"bytes,4,rep,name=proposals,proto3" json:"proposals,omitempty"`
	// Deprecated: Prefer to use `params` instead.
	// deposit_params defines all the paramaters of related to deposit.
	DepositParams *DepositParams `protobuf:"bytes,5,opt,name=deposit_params,json=depositParams,proto3" json:"deposit_params,omitempty"` // Deprecated: Do not use.
	// Deprecated: Prefer to use `params` instead.
	// voting_params defines all the paramaters of related to voting.
	VotingParams *VotingParams `protobuf:"bytes,6,opt,name=voting_params,json=votingParams,proto3" json:"voting_params,omitempty"` // Deprecated: Do not use.
	// Deprecated: Prefer to use `params` instead.
	// tally_params defines all the paramaters of related to tally.
	TallyParams *TallyParams `protobuf:"bytes,7,opt,name=tally_params,json=tallyParams,proto3" json:"tally_params,omitempty"` // Deprecated: Do not use.
	// params defines all the paramaters of x/gov module.
	//
	// Since: cosmos-sdk 0.47
	Params *Params `protobuf:"bytes,8,opt,name=params,proto3" json:"params,omitempty"`
	// The constitution allows builders to lay a foundation and define purpose.
	// This is an immutable string set in genesis.
	// There are no amendments, to go outside of scope, just fork.
	// constitution is an immutable string in genesis for a chain builder to lay out their vision, ideas and ideals.
	//
	// Since: cosmos-sdk 0.48
	Constitution string `protobuf:"bytes,9,opt,name=constitution,proto3" json:"constitution,omitempty"`
	// starting_representative_id is the ID of the starting representative.
	StartingRepresentativeId uint64 `protobuf:"varint,10,opt,name=starting_representative_id,json=startingRepresentativeId,proto3" json:"starting_representative_id,omitempty"`
	// representatives defines all the representatives present at genesis.
	Representatives []*Representative `protobuf:"bytes,11,rep,name=representatives,proto3" json:"representatives,omitempty"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_ef7cfd15e3ded621, []int{0}
}
func (m *GenesisState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenesisState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenesisState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisState.Merge(m, src)
}
func (m *GenesisState) XXX_Size() int {
	return m.Size()
}
func (m *GenesisState) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisState.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisState proto.InternalMessageInfo

func (m *GenesisState) GetStartingProposalId() uint64 {
	if m != nil {
		return m.StartingProposalId
	}
	return 0
}

func (m *GenesisState) GetDeposits() []*Deposit {
	if m != nil {
		return m.Deposits
	}
	return nil
}

func (m *GenesisState) GetVotes() []*Vote {
	if m != nil {
		return m.Votes
	}
	return nil
}

func (m *GenesisState) GetProposals() []*Proposal {
	if m != nil {
		return m.Proposals
	}
	return nil
}

// Deprecated: Do not use.
func (m *GenesisState) GetDepositParams() *DepositParams {
	if m != nil {
		return m.DepositParams
	}
	return nil
}

// Deprecated: Do not use.
func (m *GenesisState) GetVotingParams() *VotingParams {
	if m != nil {
		return m.VotingParams
	}
	return nil
}

// Deprecated: Do not use.
func (m *GenesisState) GetTallyParams() *TallyParams {
	if m != nil {
		return m.TallyParams
	}
	return nil
}

func (m *GenesisState) GetParams() *Params {
	if m != nil {
		return m.Params
	}
	return nil
}

func (m *GenesisState) GetConstitution() string {
	if m != nil {
		return m.Constitution
	}
	return ""
}

func (m *GenesisState) GetStartingRepresentativeId() uint64 {
	if m != nil {
		return m.StartingRepresentativeId
	}
	return 0
}

func (m *GenesisState) GetRepresentatives() []*Representative {
	if m != nil {
		return m.Representatives
	}
	return nil
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "cosmos.gov.v1.GenesisState")
}

func init() { proto.RegisterFile("cosmos/gov/v1/genesis.proto", fileDescriptor_ef7cfd15e3ded621) }

var fileDescriptor_ef7cfd15e3ded621 = []byte{
	// 427 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x92, 0xcf, 0x6e, 0xd3, 0x40,
	0x10, 0xc6, 0xe3, 0xb6, 0x09, 0xcd, 0x26, 0x01, 0x69, 0xf9, 0xd3, 0x55, 0x0a, 0x56, 0xd4, 0x93,
	0x11, 0xaa, 0x4d, 0x82, 0xb8, 0x71, 0xa1, 0x2a, 0x8a, 0x7a, 0xab, 0x16, 0xc4, 0x81, 0x4b, 0xe4,
	0xc6, 0x2b, 0xb3, 0x22, 0xf1, 0x58, 0x9e, 0xe9, 0x8a, 0x5e, 0x79, 0x02, 0x1e, 0x8b, 0x63, 0x8f,
	0x1c, 0x51, 0xf2, 0x22, 0x28, 0xbb, 0x76, 0x12, 0x1b, 0x4e, 0x2b, 0xcd, 0xfc, 0xbe, 0x6f, 0x3f,
	0xcd, 0x0c, 0x3b, 0x9d, 0x03, 0x2e, 0x01, 0xa3, 0x14, 0x4c, 0x64, 0xc6, 0x51, 0xaa, 0x32, 0x85,
	0x1a, 0xc3, 0xbc, 0x00, 0x02, 0x3e, 0x70, 0xcd, 0x30, 0x05, 0x13, 0x9a, 0xf1, 0xf0, 0xa4, 0xc1,
	0x82, 0x71, 0xdc, 0xd9, 0x8f, 0x36, 0xeb, 0x4f, 0x9d, 0xf2, 0x23, 0xc5, 0xa4, 0xf8, 0x6b, 0xf6,
	0x04, 0x29, 0x2e, 0x48, 0x67, 0xe9, 0x2c, 0x2f, 0x20, 0x07, 0x8c, 0x17, 0x33, 0x9d, 0x08, 0x6f,
	0xe4, 0x05, 0x47, 0x92, 0x57, 0xbd, 0xeb, 0xb2, 0x75, 0x95, 0xf0, 0x09, 0x3b, 0x4e, 0x54, 0x0e,
	0xa8, 0x09, 0xc5, 0xc1, 0xe8, 0x30, 0xe8, 0x4d, 0x9e, 0x85, 0xb5, 0xdf, 0xc3, 0x4b, 0xd7, 0x96,
	0x5b, 0x8e, 0xbf, 0x64, 0x6d, 0x03, 0xa4, 0x50, 0x1c, 0x5a, 0xc1, 0xe3, 0x86, 0xe0, 0x33, 0x90,
	0x92, 0x8e, 0xe0, 0x6f, 0x59, 0xb7, 0xca, 0x81, 0xe2, 0xc8, 0xe2, 0x27, 0x0d, 0xbc, 0x0a, 0x23,
	0x77, 0x24, 0x9f, 0xb2, 0x87, 0xe5, 0x6f, 0xb3, 0x3c, 0x2e, 0xe2, 0x25, 0x8a, 0xf6, 0xc8, 0x0b,
	0x7a, 0x93, 0xe7, 0xff, 0xcf, 0x76, 0x6d, 0x99, 0x8b, 0x03, 0xe1, 0xc9, 0x41, 0xb2, 0x5f, 0xe2,
	0x97, 0x6c, 0x60, 0xc0, 0x8d, 0xc3, 0xf9, 0x74, 0xac, 0xcf, 0xe9, 0xbf, 0x91, 0x37, 0x63, 0xd9,
	0xd9, 0xf4, 0xcd, 0x5e, 0x85, 0xbf, 0x67, 0x7d, 0x8a, 0x17, 0x8b, 0xbb, 0xca, 0xe4, 0x81, 0x35,
	0x19, 0x36, 0x4c, 0x3e, 0x6d, 0x90, 0x3d, 0x8f, 0x1e, 0xed, 0x0a, 0xfc, 0x9c, 0x75, 0x4a, 0xf1,
	0xb1, 0x15, 0x3f, 0x6d, 0x4e, 0xc1, 0x36, 0x65, 0x09, 0xf1, 0x33, 0xd6, 0x9f, 0x43, 0x86, 0xa4,
	0xe9, 0x96, 0x34, 0x64, 0xa2, 0x3b, 0xf2, 0x82, 0xae, 0xac, 0xd5, 0xf8, 0x3b, 0x36, 0xdc, 0x2e,
	0xbb, 0x50, 0x79, 0xa1, 0x50, 0x65, 0x14, 0x93, 0x36, 0x6a, 0xb3, 0x72, 0x66, 0x57, 0x2e, 0x2a,
	0x42, 0xd6, 0x80, 0xab, 0x84, 0x4f, 0xd9, 0xa3, 0xba, 0x08, 0x45, 0xcf, 0xee, 0xe7, 0x45, 0x23,
	0x59, 0x5d, 0x29, 0x9b, 0xaa, 0x8b, 0x0f, 0xbf, 0x56, 0xbe, 0x77, 0xbf, 0xf2, 0xbd, 0x3f, 0x2b,
	0xdf, 0xfb, 0xb9, 0xf6, 0x5b, 0xf7, 0x6b, 0xbf, 0xf5, 0x7b, 0xed, 0xb7, 0xbe, 0xbc, 0x4a, 0x35,
	0x7d, 0xbd, 0xbd, 0x09, 0xe7, 0xb0, 0x8c, 0xca, 0x13, 0x76, 0xcf, 0x39, 0x26, 0xdf, 0xa2, 0xef,
	0xf6, 0x9e, 0xe9, 0x2e, 0x57, 0x18, 0x99, 0xf1, 0x4d, 0xc7, 0x9e, 0xf4, 0x9b, 0xbf, 0x01, 0x00,
	0x00, 0xff, 0xff, 0x6e, 0xf2, 0x6d, 0xec, 0x19, 0x03, 0x00, 0x00,
}

func (m *GenesisState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Representatives) > 0 {
		for iNdEx := len(m.Representatives) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Representatives[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x5a
		}
	}
	if m.StartingRepresentativeId != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.StartingRepresentativeId))
		i--
		dAtA[i] = 0x50
	}
	if len(m.Constitution) > 0 {
		i -= len(m.Constitution)
		copy(dAtA[i:], m.Constitution)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.Constitution)))
		i--
		dAtA[i] = 0x4a
	}
	if m.Params != nil {
		{
			size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintGenesis(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x42
	}
	if m.TallyParams != nil {
		{
			size, err := m.TallyParams.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintGenesis(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x3a
	}
	if m.VotingParams != nil {
		{
			size, err := m.VotingParams.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintGenesis(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x32
	}
	if m.DepositParams != nil {
		{
			size, err := m.DepositParams.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintGenesis(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x2a
	}
	if len(m.Proposals) > 0 {
		for iNdEx := len(m.Proposals) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Proposals[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.Votes) > 0 {
		for iNdEx := len(m.Votes) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Votes[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.Deposits) > 0 {
		for iNdEx := len(m.Deposits) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Deposits[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if m.StartingProposalId != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.StartingProposalId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintGenesis(dAtA []byte, offset int, v uint64) int {
	offset -= sovGenesis(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GenesisState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.StartingProposalId != 0 {
		n += 1 + sovGenesis(uint64(m.StartingProposalId))
	}
	if len(m.Deposits) > 0 {
		for _, e := range m.Deposits {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.Votes) > 0 {
		for _, e := range m.Votes {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.Proposals) > 0 {
		for _, e := range m.Proposals {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if m.DepositParams != nil {
		l = m.DepositParams.Size()
		n += 1 + l + sovGenesis(uint64(l))
	}
	if m.VotingParams != nil {
		l = m.VotingParams.Size()
		n += 1 + l + sovGenesis(uint64(l))
	}
	if m.TallyParams != nil {
		l = m.TallyParams.Size()
		n += 1 + l + sovGenesis(uint64(l))
	}
	if m.Params != nil {
		l = m.Params.Size()
		n += 1 + l + sovGenesis(uint64(l))
	}
	l = len(m.Constitution)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	if m.StartingRepresentativeId != 0 {
		n += 1 + sovGenesis(uint64(m.StartingRepresentativeId))
	}
	if len(m.Representatives) > 0 {
		for _, e := range m.Representatives {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	return n
}

func sovGenesis(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGenesis(x uint64) (n int) {
	return sovGenesis(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GenesisState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: GenesisState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenesisState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field StartingProposalId", wireType)
			}
			m.StartingProposalId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.StartingProposalId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Deposits", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Deposits = append(m.Deposits, &Deposit{})
			if err := m.Deposits[len(m.Deposits)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Votes", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Votes = append(m.Votes, &Vote{})
			if err := m.Votes[len(m.Votes)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Proposals", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Proposals = append(m.Proposals, &Proposal{})
			if err := m.Proposals[len(m.Proposals)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DepositParams", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.DepositParams == nil {
				m.DepositParams = &DepositParams{}
			}
			if err := m.DepositParams.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field VotingParams", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.VotingParams == nil {
				m.VotingParams = &VotingParams{}
			}
			if err := m.VotingParams.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TallyParams", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.TallyParams == nil {
				m.TallyParams = &TallyParams{}
			}
			if err := m.TallyParams.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Params == nil {
				m.Params = &Params{}
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Constitution", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Constitution = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 10:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field StartingRepresentativeId", wireType)
			}
			m.StartingRepresentativeId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.StartingRepresentativeId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 11:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Representatives", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Representatives = append(m.Representatives, &Representative{})
			if err := m.Representatives[len(m.Representatives)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipGenesis(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenesis
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthGenesis
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGenesis
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGenesis
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGenesis        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenesis          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGenesis = fmt.Errorf("proto: unexpected end of group")
)
