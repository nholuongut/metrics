/*
Copyright Nho Luong DevOps.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: k8s.io/metrics/pkg/apis/external_metrics/v1beta1/generated.proto

package v1beta1

import (
	fmt "fmt"

	io "io"

	proto "github.com/gogo/protobuf/proto"
	github_com_gogo_protobuf_sortkeys "github.com/gogo/protobuf/sortkeys"

	math "math"
	math_bits "math/bits"
	reflect "reflect"
	strings "strings"
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

func (m *ExternalMetricValue) Reset()      { *m = ExternalMetricValue{} }
func (*ExternalMetricValue) ProtoMessage() {}
func (*ExternalMetricValue) Descriptor() ([]byte, []int) {
	return fileDescriptor_56302cb587e818f0, []int{0}
}
func (m *ExternalMetricValue) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ExternalMetricValue) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *ExternalMetricValue) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExternalMetricValue.Merge(m, src)
}
func (m *ExternalMetricValue) XXX_Size() int {
	return m.Size()
}
func (m *ExternalMetricValue) XXX_DiscardUnknown() {
	xxx_messageInfo_ExternalMetricValue.DiscardUnknown(m)
}

var xxx_messageInfo_ExternalMetricValue proto.InternalMessageInfo

func (m *ExternalMetricValueList) Reset()      { *m = ExternalMetricValueList{} }
func (*ExternalMetricValueList) ProtoMessage() {}
func (*ExternalMetricValueList) Descriptor() ([]byte, []int) {
	return fileDescriptor_56302cb587e818f0, []int{1}
}
func (m *ExternalMetricValueList) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ExternalMetricValueList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *ExternalMetricValueList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExternalMetricValueList.Merge(m, src)
}
func (m *ExternalMetricValueList) XXX_Size() int {
	return m.Size()
}
func (m *ExternalMetricValueList) XXX_DiscardUnknown() {
	xxx_messageInfo_ExternalMetricValueList.DiscardUnknown(m)
}

var xxx_messageInfo_ExternalMetricValueList proto.InternalMessageInfo

func init() {
	proto.RegisterType((*ExternalMetricValue)(nil), "k8s.io.metrics.pkg.apis.external_metrics.v1beta1.ExternalMetricValue")
	proto.RegisterMapType((map[string]string)(nil), "k8s.io.metrics.pkg.apis.external_metrics.v1beta1.ExternalMetricValue.MetricLabelsEntry")
	proto.RegisterType((*ExternalMetricValueList)(nil), "k8s.io.metrics.pkg.apis.external_metrics.v1beta1.ExternalMetricValueList")
}

func init() {
	proto.RegisterFile("k8s.io/metrics/pkg/apis/external_metrics/v1beta1/generated.proto", fileDescriptor_56302cb587e818f0)
}

var fileDescriptor_56302cb587e818f0 = []byte{
	// 539 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x94, 0xcf, 0x6b, 0xdb, 0x30,
	0x14, 0xc7, 0xad, 0x78, 0x19, 0x8d, 0xda, 0x42, 0xa3, 0x15, 0x66, 0x72, 0x70, 0x42, 0x4f, 0xd9,
	0x60, 0x52, 0x13, 0xc6, 0x28, 0xbb, 0x6c, 0x18, 0x72, 0x18, 0x34, 0x83, 0xb9, 0xa5, 0x85, 0x6d,
	0x30, 0x14, 0x47, 0x73, 0xb4, 0xc4, 0x3f, 0xb0, 0xe4, 0x74, 0xb9, 0xed, 0x4f, 0xd8, 0xfe, 0xab,
	0x1c, 0x7b, 0xec, 0x29, 0x2c, 0xee, 0x9f, 0xb1, 0xcb, 0xb0, 0x6c, 0x37, 0x59, 0xd2, 0xb5, 0x04,
	0x7a, 0xf3, 0x93, 0xf4, 0xbe, 0xef, 0xf3, 0xbe, 0xef, 0x61, 0xf8, 0x76, 0x78, 0x24, 0x30, 0x0f,
	0x88, 0xc7, 0x64, 0xc4, 0x1d, 0x41, 0xc2, 0xa1, 0x4b, 0x68, 0xc8, 0x05, 0x61, 0xdf, 0x25, 0x8b,
	0x7c, 0x3a, 0xfa, 0x52, 0xdc, 0x8c, 0x5b, 0x3d, 0x26, 0x69, 0x8b, 0xb8, 0xcc, 0x67, 0x11, 0x95,
	0xac, 0x8f, 0xc3, 0x28, 0x90, 0x01, 0x3a, 0xcc, 0x14, 0x70, 0xfe, 0x0e, 0x87, 0x43, 0x17, 0xa7,
	0x0a, 0x78, 0x55, 0x01, 0xe7, 0x0a, 0xb5, 0x17, 0x2e, 0x97, 0x83, 0xb8, 0x87, 0x9d, 0xc0, 0x23,
	0x6e, 0xe0, 0x06, 0x44, 0x09, 0xf5, 0xe2, 0xaf, 0x2a, 0x52, 0x81, 0xfa, 0xca, 0x0a, 0xd4, 0x5e,
	0xe6, 0x88, 0x34, 0xe4, 0x1e, 0x75, 0x06, 0xdc, 0x67, 0xd1, 0xa4, 0xe0, 0x24, 0x11, 0x13, 0x41,
	0x1c, 0x39, 0x6c, 0x15, 0xeb, 0xce, 0x2c, 0x91, 0xb6, 0x4b, 0xc9, 0x78, 0xad, 0x99, 0x1a, 0xf9,
	0x5f, 0x56, 0x14, 0xfb, 0x92, 0x7b, 0xeb, 0x65, 0x5e, 0xdd, 0x97, 0x20, 0x9c, 0x01, 0xf3, 0xe8,
	0x6a, 0xde, 0xc1, 0x1f, 0x1d, 0x3e, 0xe9, 0xe4, 0x06, 0x75, 0x95, 0x3f, 0x67, 0x74, 0x14, 0x33,
	0xd4, 0x86, 0x30, 0xb3, 0xeb, 0x3d, 0xf5, 0x98, 0x01, 0x1a, 0xa0, 0x59, 0xb1, 0xd0, 0x74, 0x56,
	0xd7, 0x92, 0x59, 0x1d, 0x76, 0x6f, 0x6e, 0xec, 0xa5, 0x57, 0xe8, 0x17, 0x80, 0x3b, 0x59, 0x78,
	0x4c, 0x7b, 0x6c, 0x24, 0x8c, 0x52, 0x43, 0x6f, 0x6e, 0xb7, 0xcf, 0xf1, 0xa6, 0x93, 0xc1, 0xb7,
	0x10, 0xe1, 0xee, 0x92, 0x72, 0xc7, 0x97, 0xd1, 0xc4, 0xda, 0xcf, 0x79, 0x76, 0x96, 0xaf, 0xec,
	0x7f, 0x10, 0xd0, 0x27, 0x58, 0x49, 0xdb, 0x17, 0x92, 0x7a, 0xa1, 0xa1, 0x37, 0x40, 0x73, 0xbb,
	0xfd, 0xbc, 0xe0, 0x59, 0xf6, 0x6a, 0x01, 0x95, 0x8e, 0x04, 0x8f, 0x5b, 0xf8, 0x94, 0x7b, 0xcc,
	0xaa, 0xe6, 0x25, 0x2a, 0xa7, 0x85, 0x88, 0xbd, 0xd0, 0x43, 0xcf, 0xe0, 0xe3, 0x0b, 0xee, 0xf7,
	0x83, 0x0b, 0xe3, 0x51, 0x03, 0x34, 0x75, 0xab, 0x9a, 0xcc, 0xea, 0xbb, 0xe7, 0xea, 0xe4, 0x84,
	0x39, 0x81, 0xdf, 0x17, 0x76, 0xfe, 0x00, 0x9d, 0xc0, 0xf2, 0x38, 0x6d, 0xc3, 0x28, 0x2b, 0x06,
	0x7c, 0x17, 0x03, 0x2e, 0x96, 0x09, 0x7f, 0x88, 0xa9, 0x2f, 0xb9, 0x9c, 0x58, 0xbb, 0x39, 0x47,
	0x59, 0x79, 0x61, 0x67, 0x5a, 0xb5, 0x37, 0xb0, 0xba, 0xe6, 0x0a, 0xda, 0x83, 0xfa, 0x90, 0x4d,
	0xb2, 0x91, 0xd9, 0xe9, 0x27, 0xda, 0x2f, 0x6a, 0x97, 0xd4, 0x59, 0x16, 0xbc, 0x2e, 0x1d, 0x81,
	0x83, 0x6b, 0x00, 0x9f, 0xde, 0xe2, 0xf5, 0x31, 0x17, 0x12, 0x7d, 0x86, 0x5b, 0xa9, 0x15, 0x7d,
	0x2a, 0xa9, 0x12, 0xbb, 0x07, 0x7a, 0x61, 0x5c, 0x9a, 0xdd, 0x65, 0x92, 0x5a, 0x7b, 0x39, 0xf4,
	0x56, 0x71, 0x62, 0xdf, 0x28, 0xa2, 0x6f, 0xb0, 0xcc, 0x25, 0xf3, 0x8a, 0x1d, 0xe9, 0x3c, 0xc8,
	0x8e, 0x2c, 0x6c, 0x7a, 0x97, 0x6a, 0xdb, 0x59, 0x09, 0xeb, 0x6c, 0x3a, 0x37, 0xb5, 0xcb, 0xb9,
	0xa9, 0x5d, 0xcd, 0x4d, 0xed, 0x47, 0x62, 0x82, 0x69, 0x62, 0x82, 0xcb, 0xc4, 0x04, 0x57, 0x89,
	0x09, 0x7e, 0x27, 0x26, 0xf8, 0x79, 0x6d, 0x6a, 0x1f, 0x0f, 0x37, 0xfd, 0x03, 0xfd, 0x0d, 0x00,
	0x00, 0xff, 0xff, 0x68, 0x2d, 0x7f, 0xed, 0xb4, 0x04, 0x00, 0x00,
}

func (m *ExternalMetricValue) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ExternalMetricValue) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ExternalMetricValue) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Value.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenerated(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x2a
	if m.WindowSeconds != nil {
		i = encodeVarintGenerated(dAtA, i, uint64(*m.WindowSeconds))
		i--
		dAtA[i] = 0x20
	}
	{
		size, err := m.Timestamp.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenerated(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	if len(m.MetricLabels) > 0 {
		keysForMetricLabels := make([]string, 0, len(m.MetricLabels))
		for k := range m.MetricLabels {
			keysForMetricLabels = append(keysForMetricLabels, string(k))
		}
		github_com_gogo_protobuf_sortkeys.Strings(keysForMetricLabels)
		for iNdEx := len(keysForMetricLabels) - 1; iNdEx >= 0; iNdEx-- {
			v := m.MetricLabels[string(keysForMetricLabels[iNdEx])]
			baseI := i
			i -= len(v)
			copy(dAtA[i:], v)
			i = encodeVarintGenerated(dAtA, i, uint64(len(v)))
			i--
			dAtA[i] = 0x12
			i -= len(keysForMetricLabels[iNdEx])
			copy(dAtA[i:], keysForMetricLabels[iNdEx])
			i = encodeVarintGenerated(dAtA, i, uint64(len(keysForMetricLabels[iNdEx])))
			i--
			dAtA[i] = 0xa
			i = encodeVarintGenerated(dAtA, i, uint64(baseI-i))
			i--
			dAtA[i] = 0x12
		}
	}
	i -= len(m.MetricName)
	copy(dAtA[i:], m.MetricName)
	i = encodeVarintGenerated(dAtA, i, uint64(len(m.MetricName)))
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *ExternalMetricValueList) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ExternalMetricValueList) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ExternalMetricValueList) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Items) > 0 {
		for iNdEx := len(m.Items) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Items[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenerated(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	{
		size, err := m.ListMeta.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenerated(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintGenerated(dAtA []byte, offset int, v uint64) int {
	offset -= sovGenerated(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ExternalMetricValue) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.MetricName)
	n += 1 + l + sovGenerated(uint64(l))
	if len(m.MetricLabels) > 0 {
		for k, v := range m.MetricLabels {
			_ = k
			_ = v
			mapEntrySize := 1 + len(k) + sovGenerated(uint64(len(k))) + 1 + len(v) + sovGenerated(uint64(len(v)))
			n += mapEntrySize + 1 + sovGenerated(uint64(mapEntrySize))
		}
	}
	l = m.Timestamp.Size()
	n += 1 + l + sovGenerated(uint64(l))
	if m.WindowSeconds != nil {
		n += 1 + sovGenerated(uint64(*m.WindowSeconds))
	}
	l = m.Value.Size()
	n += 1 + l + sovGenerated(uint64(l))
	return n
}

func (m *ExternalMetricValueList) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.ListMeta.Size()
	n += 1 + l + sovGenerated(uint64(l))
	if len(m.Items) > 0 {
		for _, e := range m.Items {
			l = e.Size()
			n += 1 + l + sovGenerated(uint64(l))
		}
	}
	return n
}

func sovGenerated(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGenerated(x uint64) (n int) {
	return sovGenerated(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *ExternalMetricValue) String() string {
	if this == nil {
		return "nil"
	}
	keysForMetricLabels := make([]string, 0, len(this.MetricLabels))
	for k := range this.MetricLabels {
		keysForMetricLabels = append(keysForMetricLabels, k)
	}
	github_com_gogo_protobuf_sortkeys.Strings(keysForMetricLabels)
	mapStringForMetricLabels := "map[string]string{"
	for _, k := range keysForMetricLabels {
		mapStringForMetricLabels += fmt.Sprintf("%v: %v,", k, this.MetricLabels[k])
	}
	mapStringForMetricLabels += "}"
	s := strings.Join([]string{`&ExternalMetricValue{`,
		`MetricName:` + fmt.Sprintf("%v", this.MetricName) + `,`,
		`MetricLabels:` + mapStringForMetricLabels + `,`,
		`Timestamp:` + strings.Replace(strings.Replace(fmt.Sprintf("%v", this.Timestamp), "Time", "v1.Time", 1), `&`, ``, 1) + `,`,
		`WindowSeconds:` + valueToStringGenerated(this.WindowSeconds) + `,`,
		`Value:` + strings.Replace(strings.Replace(fmt.Sprintf("%v", this.Value), "Quantity", "resource.Quantity", 1), `&`, ``, 1) + `,`,
		`}`,
	}, "")
	return s
}
func (this *ExternalMetricValueList) String() string {
	if this == nil {
		return "nil"
	}
	repeatedStringForItems := "[]ExternalMetricValue{"
	for _, f := range this.Items {
		repeatedStringForItems += strings.Replace(strings.Replace(f.String(), "ExternalMetricValue", "ExternalMetricValue", 1), `&`, ``, 1) + ","
	}
	repeatedStringForItems += "}"
	s := strings.Join([]string{`&ExternalMetricValueList{`,
		`ListMeta:` + strings.Replace(strings.Replace(fmt.Sprintf("%v", this.ListMeta), "ListMeta", "v1.ListMeta", 1), `&`, ``, 1) + `,`,
		`Items:` + repeatedStringForItems + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringGenerated(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *ExternalMetricValue) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenerated
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
			return fmt.Errorf("proto: ExternalMetricValue: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ExternalMetricValue: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MetricName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
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
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenerated
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MetricName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MetricLabels", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
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
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenerated
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.MetricLabels == nil {
				m.MetricLabels = make(map[string]string)
			}
			var mapkey string
			var mapvalue string
			for iNdEx < postIndex {
				entryPreIndex := iNdEx
				var wire uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowGenerated
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
				if fieldNum == 1 {
					var stringLenmapkey uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowGenerated
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapkey |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapkey := int(stringLenmapkey)
					if intStringLenmapkey < 0 {
						return ErrInvalidLengthGenerated
					}
					postStringIndexmapkey := iNdEx + intStringLenmapkey
					if postStringIndexmapkey < 0 {
						return ErrInvalidLengthGenerated
					}
					if postStringIndexmapkey > l {
						return io.ErrUnexpectedEOF
					}
					mapkey = string(dAtA[iNdEx:postStringIndexmapkey])
					iNdEx = postStringIndexmapkey
				} else if fieldNum == 2 {
					var stringLenmapvalue uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowGenerated
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapvalue |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapvalue := int(stringLenmapvalue)
					if intStringLenmapvalue < 0 {
						return ErrInvalidLengthGenerated
					}
					postStringIndexmapvalue := iNdEx + intStringLenmapvalue
					if postStringIndexmapvalue < 0 {
						return ErrInvalidLengthGenerated
					}
					if postStringIndexmapvalue > l {
						return io.ErrUnexpectedEOF
					}
					mapvalue = string(dAtA[iNdEx:postStringIndexmapvalue])
					iNdEx = postStringIndexmapvalue
				} else {
					iNdEx = entryPreIndex
					skippy, err := skipGenerated(dAtA[iNdEx:])
					if err != nil {
						return err
					}
					if (skippy < 0) || (iNdEx+skippy) < 0 {
						return ErrInvalidLengthGenerated
					}
					if (iNdEx + skippy) > postIndex {
						return io.ErrUnexpectedEOF
					}
					iNdEx += skippy
				}
			}
			m.MetricLabels[mapkey] = mapvalue
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Timestamp", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
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
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenerated
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Timestamp.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field WindowSeconds", wireType)
			}
			var v int64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.WindowSeconds = &v
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
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
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenerated
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Value.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenerated
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
func (m *ExternalMetricValueList) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenerated
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
			return fmt.Errorf("proto: ExternalMetricValueList: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ExternalMetricValueList: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ListMeta", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
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
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenerated
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ListMeta.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Items", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
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
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenerated
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Items = append(m.Items, ExternalMetricValue{})
			if err := m.Items[len(m.Items)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenerated
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
func skipGenerated(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenerated
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
					return 0, ErrIntOverflowGenerated
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
					return 0, ErrIntOverflowGenerated
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
				return 0, ErrInvalidLengthGenerated
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGenerated
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGenerated
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGenerated        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenerated          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGenerated = fmt.Errorf("proto: unexpected end of group")
)
