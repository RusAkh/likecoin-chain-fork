// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: likechain/likenft/v1/genesis.proto

package types

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
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

// GenesisState defines the likenft module's genesis state.
type GenesisState struct {
	Params                   Params                    `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
	ClassesByIscnList        []ClassesByISCN           `protobuf:"bytes,2,rep,name=classes_by_iscn_list,json=classesByIscnList,proto3" json:"classes_by_iscn_list"`
	ClassesByAccountList     []ClassesByAccount        `protobuf:"bytes,3,rep,name=classes_by_account_list,json=classesByAccountList,proto3" json:"classes_by_account_list"`
	BlindBoxContentList      []BlindBoxContent         `protobuf:"bytes,4,rep,name=blind_box_content_list,json=blindBoxContentList,proto3" json:"blind_box_content_list"`
	ClassRevealQueue         []ClassRevealQueueEntry   `protobuf:"bytes,5,rep,name=class_reveal_queue,json=classRevealQueue,proto3" json:"class_reveal_queue"`
	OfferList                []Offer                   `protobuf:"bytes,6,rep,name=offer_list,json=offerList,proto3" json:"offer_list"`
	ListingList              []Listing                 `protobuf:"bytes,7,rep,name=listing_list,json=listingList,proto3" json:"listing_list"`
	OfferExpireQueue         []OfferExpireQueueEntry   `protobuf:"bytes,8,rep,name=offer_expire_queue,json=offerExpireQueue,proto3" json:"offer_expire_queue"`
	ListingExpireQueue       []ListingExpireQueueEntry `protobuf:"bytes,9,rep,name=listing_expire_queue,json=listingExpireQueue,proto3" json:"listing_expire_queue"`
	RoyaltyConfigByClassList []RoyaltyConfigByClass    `protobuf:"bytes,10,rep,name=royalty_config_by_class_list,json=royaltyConfigByClassList,proto3" json:"royalty_config_by_class_list"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_e01c79fdac411e6f, []int{0}
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

func (m *GenesisState) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

func (m *GenesisState) GetClassesByIscnList() []ClassesByISCN {
	if m != nil {
		return m.ClassesByIscnList
	}
	return nil
}

func (m *GenesisState) GetClassesByAccountList() []ClassesByAccount {
	if m != nil {
		return m.ClassesByAccountList
	}
	return nil
}

func (m *GenesisState) GetBlindBoxContentList() []BlindBoxContent {
	if m != nil {
		return m.BlindBoxContentList
	}
	return nil
}

func (m *GenesisState) GetClassRevealQueue() []ClassRevealQueueEntry {
	if m != nil {
		return m.ClassRevealQueue
	}
	return nil
}

func (m *GenesisState) GetOfferList() []Offer {
	if m != nil {
		return m.OfferList
	}
	return nil
}

func (m *GenesisState) GetListingList() []Listing {
	if m != nil {
		return m.ListingList
	}
	return nil
}

func (m *GenesisState) GetOfferExpireQueue() []OfferExpireQueueEntry {
	if m != nil {
		return m.OfferExpireQueue
	}
	return nil
}

func (m *GenesisState) GetListingExpireQueue() []ListingExpireQueueEntry {
	if m != nil {
		return m.ListingExpireQueue
	}
	return nil
}

func (m *GenesisState) GetRoyaltyConfigByClassList() []RoyaltyConfigByClass {
	if m != nil {
		return m.RoyaltyConfigByClassList
	}
	return nil
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "likechain.likenft.v1.GenesisState")
}

func init() {
	proto.RegisterFile("likechain/likenft/v1/genesis.proto", fileDescriptor_e01c79fdac411e6f)
}

var fileDescriptor_e01c79fdac411e6f = []byte{
	// 555 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x94, 0x41, 0x6f, 0xd3, 0x30,
	0x14, 0xc7, 0x5b, 0x36, 0x3a, 0xe6, 0xee, 0x00, 0xa1, 0x82, 0xa8, 0x8c, 0x50, 0x86, 0x40, 0x63,
	0xd0, 0x44, 0xdb, 0xc4, 0x85, 0x13, 0xa4, 0x1a, 0x08, 0x69, 0x62, 0xd0, 0xdd, 0x76, 0x09, 0x89,
	0x71, 0x3b, 0x8b, 0xcc, 0x0e, 0xb1, 0x5b, 0x35, 0xdf, 0x02, 0xf1, 0xa9, 0x76, 0xdc, 0x91, 0x13,
	0x42, 0xed, 0x17, 0x41, 0x7e, 0xf6, 0x2a, 0xb5, 0xf1, 0xd2, 0x5b, 0x64, 0xff, 0xde, 0xff, 0xff,
	0x9e, 0xdf, 0x7b, 0x41, 0x3b, 0x29, 0xfd, 0x41, 0xf0, 0x79, 0x4c, 0x59, 0xa0, 0xbe, 0xd8, 0x40,
	0x06, 0xe3, 0xfd, 0x60, 0x48, 0x18, 0x11, 0x54, 0xf8, 0x59, 0xce, 0x25, 0x77, 0x5a, 0x73, 0xc6,
	0x37, 0x8c, 0x3f, 0xde, 0x6f, 0xb7, 0x86, 0x7c, 0xc8, 0x01, 0x08, 0xd4, 0x97, 0x66, 0xdb, 0xaf,
	0xad, 0x7a, 0x49, 0x4a, 0xd9, 0xf7, 0x28, 0xe1, 0x93, 0x08, 0x73, 0x26, 0x09, 0x93, 0x86, 0xee,
	0x5a, 0x69, 0x9c, 0xc6, 0x42, 0x44, 0x39, 0x19, 0x93, 0x38, 0x8d, 0x7e, 0x8e, 0xc8, 0x88, 0xac,
	0xc6, 0x89, 0x88, 0x92, 0x22, 0x8a, 0x31, 0xe6, 0xa3, 0xb9, 0xfa, 0xde, 0x2a, 0x9c, 0x0a, 0xcc,
	0x0c, 0x6b, 0x7f, 0x87, 0x94, 0x0a, 0x49, 0xd9, 0xd0, 0x30, 0x41, 0x15, 0x13, 0x91, 0x49, 0x46,
	0x73, 0xb2, 0x90, 0x6f, 0xc7, 0x1a, 0xc0, 0x07, 0x03, 0x92, 0x57, 0x56, 0x04, 0x84, 0x4d, 0xf0,
	0xa9, 0x15, 0xcf, 0xe2, 0x3c, 0xbe, 0x30, 0xcd, 0x6a, 0xbf, 0xb4, 0x22, 0x39, 0x2f, 0xe2, 0x54,
	0x16, 0xea, 0xf9, 0x07, 0xd4, 0xd4, 0xb3, 0xf3, 0x7b, 0x03, 0x6d, 0x7d, 0xd4, 0x9d, 0x3e, 0x95,
	0xb1, 0x24, 0xce, 0x5b, 0xd4, 0xd0, 0x5a, 0x6e, 0xbd, 0x53, 0xdf, 0x6d, 0x1e, 0x6c, 0xfb, 0xb6,
	0xce, 0xfb, 0x5f, 0x80, 0x09, 0xd7, 0x2f, 0xff, 0x3e, 0xa9, 0xf5, 0x4d, 0x84, 0x73, 0x86, 0x5a,
	0x4b, 0x2f, 0x1b, 0xa9, 0x97, 0x71, 0x6f, 0x75, 0xd6, 0x76, 0x9b, 0x07, 0xcf, 0xec, 0x4a, 0x3d,
	0x1d, 0x11, 0x16, 0x9f, 0x4e, 0x7b, 0x9f, 0x8d, 0xe0, 0x3d, 0x3c, 0x3f, 0x14, 0x98, 0x1d, 0x53,
	0x21, 0x1d, 0x8c, 0x1e, 0x96, 0x9b, 0xac, 0xe5, 0xd7, 0x40, 0xfe, 0xc5, 0x0a, 0xf9, 0xf7, 0x3a,
	0xc4, 0x38, 0xb4, 0xf0, 0xd2, 0x39, 0x98, 0x7c, 0x43, 0x0f, 0x4a, 0x63, 0xaa, 0x3d, 0xd6, 0xc1,
	0xe3, 0xb9, 0xdd, 0x23, 0x54, 0x31, 0x21, 0x9f, 0xf4, 0x74, 0x84, 0xb1, 0xb8, 0x9f, 0x2c, 0x1e,
	0x83, 0x43, 0x84, 0x9c, 0xf2, 0x68, 0xbb, 0xb7, 0x41, 0xfd, 0x55, 0x45, 0x05, 0x7d, 0xc0, 0xbf,
	0x2a, 0xfa, 0x88, 0xc9, 0xbc, 0x30, 0x1e, 0x77, 0xf1, 0xd2, 0xa5, 0xf3, 0x0e, 0x21, 0x3d, 0x3a,
	0x90, 0x76, 0x03, 0x84, 0x1f, 0xd9, 0x85, 0x4f, 0x14, 0x67, 0x84, 0x36, 0x21, 0x08, 0x52, 0xfc,
	0x80, 0xb6, 0xae, 0xe7, 0x19, 0x34, 0x36, 0x40, 0xe3, 0xb1, 0x5d, 0xe3, 0x58, 0x93, 0x46, 0xa5,
	0x69, 0x02, 0xaf, 0x4b, 0x2d, 0x0f, 0xb1, 0x7b, 0xa7, 0xaa, 0x54, 0xc8, 0xe8, 0x08, 0xf0, 0x72,
	0xa9, 0x7c, 0xe9, 0xd2, 0x21, 0xa8, 0x65, 0x5b, 0x3c, 0x77, 0x13, 0x2c, 0xba, 0x95, 0x09, 0xdf,
	0x60, 0xe2, 0xa4, 0xa5, 0x6b, 0x27, 0x43, 0xdb, 0x8b, 0xab, 0xa3, 0x06, 0x50, 0x37, 0x11, 0xde,
	0x07, 0x81, 0xdd, 0x9e, 0xdd, 0xae, 0xaf, 0x23, 0x7b, 0x10, 0x18, 0x16, 0xd0, 0x4b, 0xe3, 0xe5,
	0xe6, 0x96, 0x3b, 0x95, 0x5e, 0x78, 0x72, 0x39, 0xf5, 0xea, 0x57, 0x53, 0xaf, 0xfe, 0x6f, 0xea,
	0xd5, 0x7f, 0xcd, 0xbc, 0xda, 0xd5, 0xcc, 0xab, 0xfd, 0x99, 0x79, 0xb5, 0xb3, 0x37, 0x43, 0x2a,
	0xcf, 0x47, 0x89, 0x8f, 0xf9, 0x85, 0xfe, 0x13, 0x71, 0xb3, 0xe3, 0xea, 0xa3, 0xab, 0x57, 0x7e,
	0x7c, 0x18, 0x4c, 0xe6, 0x7b, 0x2f, 0x8b, 0x8c, 0x88, 0xa4, 0x01, 0xcb, 0x7e, 0xf8, 0x3f, 0x00,
	0x00, 0xff, 0xff, 0x4d, 0xf9, 0x69, 0x02, 0xea, 0x05, 0x00, 0x00,
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
	if len(m.RoyaltyConfigByClassList) > 0 {
		for iNdEx := len(m.RoyaltyConfigByClassList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.RoyaltyConfigByClassList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x52
		}
	}
	if len(m.ListingExpireQueue) > 0 {
		for iNdEx := len(m.ListingExpireQueue) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ListingExpireQueue[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x4a
		}
	}
	if len(m.OfferExpireQueue) > 0 {
		for iNdEx := len(m.OfferExpireQueue) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.OfferExpireQueue[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x42
		}
	}
	if len(m.ListingList) > 0 {
		for iNdEx := len(m.ListingList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ListingList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x3a
		}
	}
	if len(m.OfferList) > 0 {
		for iNdEx := len(m.OfferList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.OfferList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x32
		}
	}
	if len(m.ClassRevealQueue) > 0 {
		for iNdEx := len(m.ClassRevealQueue) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ClassRevealQueue[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x2a
		}
	}
	if len(m.BlindBoxContentList) > 0 {
		for iNdEx := len(m.BlindBoxContentList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.BlindBoxContentList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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
	if len(m.ClassesByAccountList) > 0 {
		for iNdEx := len(m.ClassesByAccountList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ClassesByAccountList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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
	if len(m.ClassesByIscnList) > 0 {
		for iNdEx := len(m.ClassesByIscnList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ClassesByIscnList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
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
	l = m.Params.Size()
	n += 1 + l + sovGenesis(uint64(l))
	if len(m.ClassesByIscnList) > 0 {
		for _, e := range m.ClassesByIscnList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.ClassesByAccountList) > 0 {
		for _, e := range m.ClassesByAccountList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.BlindBoxContentList) > 0 {
		for _, e := range m.BlindBoxContentList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.ClassRevealQueue) > 0 {
		for _, e := range m.ClassRevealQueue {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.OfferList) > 0 {
		for _, e := range m.OfferList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.ListingList) > 0 {
		for _, e := range m.ListingList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.OfferExpireQueue) > 0 {
		for _, e := range m.OfferExpireQueue {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.ListingExpireQueue) > 0 {
		for _, e := range m.ListingExpireQueue {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.RoyaltyConfigByClassList) > 0 {
		for _, e := range m.RoyaltyConfigByClassList {
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
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClassesByIscnList", wireType)
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
			m.ClassesByIscnList = append(m.ClassesByIscnList, ClassesByISCN{})
			if err := m.ClassesByIscnList[len(m.ClassesByIscnList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClassesByAccountList", wireType)
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
			m.ClassesByAccountList = append(m.ClassesByAccountList, ClassesByAccount{})
			if err := m.ClassesByAccountList[len(m.ClassesByAccountList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BlindBoxContentList", wireType)
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
			m.BlindBoxContentList = append(m.BlindBoxContentList, BlindBoxContent{})
			if err := m.BlindBoxContentList[len(m.BlindBoxContentList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClassRevealQueue", wireType)
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
			m.ClassRevealQueue = append(m.ClassRevealQueue, ClassRevealQueueEntry{})
			if err := m.ClassRevealQueue[len(m.ClassRevealQueue)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OfferList", wireType)
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
			m.OfferList = append(m.OfferList, Offer{})
			if err := m.OfferList[len(m.OfferList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ListingList", wireType)
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
			m.ListingList = append(m.ListingList, Listing{})
			if err := m.ListingList[len(m.ListingList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OfferExpireQueue", wireType)
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
			m.OfferExpireQueue = append(m.OfferExpireQueue, OfferExpireQueueEntry{})
			if err := m.OfferExpireQueue[len(m.OfferExpireQueue)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ListingExpireQueue", wireType)
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
			m.ListingExpireQueue = append(m.ListingExpireQueue, ListingExpireQueueEntry{})
			if err := m.ListingExpireQueue[len(m.ListingExpireQueue)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RoyaltyConfigByClassList", wireType)
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
			m.RoyaltyConfigByClassList = append(m.RoyaltyConfigByClassList, RoyaltyConfigByClass{})
			if err := m.RoyaltyConfigByClassList[len(m.RoyaltyConfigByClassList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
