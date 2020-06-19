package vless

import (
	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	protocol "v2ray.com/core/common/protocol"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type Account struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID of the account, in the form of a UUID, e.g., "66ad4540-b58c-4ad2-9926-ea63445a9b57".
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Mess settings. Only applies to client side for now.
	Mess string `protobuf:"bytes,2,opt,name=mess,proto3" json:"mess,omitempty"`
	// Security settings. Only applies to client side.
	SecuritySettings *protocol.SecurityConfig `protobuf:"bytes,3,opt,name=security_settings,json=securitySettings,proto3" json:"security_settings,omitempty"`
}

func (x *Account) Reset() {
	*x = Account{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v2ray_com_core_proxy_vless_account_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Account) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Account) ProtoMessage() {}

func (x *Account) ProtoReflect() protoreflect.Message {
	mi := &file_v2ray_com_core_proxy_vless_account_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Account.ProtoReflect.Descriptor instead.
func (*Account) Descriptor() ([]byte, []int) {
	return file_v2ray_com_core_proxy_vless_account_proto_rawDescGZIP(), []int{0}
}

func (x *Account) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Account) GetMess() string {
	if x != nil {
		return x.Mess
	}
	return ""
}

func (x *Account) GetSecuritySettings() *protocol.SecurityConfig {
	if x != nil {
		return x.SecuritySettings
	}
	return nil
}

var File_v2ray_com_core_proxy_vless_account_proto protoreflect.FileDescriptor

var file_v2ray_com_core_proxy_vless_account_proto_rawDesc = []byte{
	0x0a, 0x28, 0x76, 0x32, 0x72, 0x61, 0x79, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2f, 0x76, 0x6c, 0x65, 0x73, 0x73, 0x2f, 0x61, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x16, 0x76, 0x32, 0x72, 0x61,
	0x79, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x76, 0x6c, 0x65,
	0x73, 0x73, 0x1a, 0x2c, 0x76, 0x32, 0x72, 0x61, 0x79, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f,
	0x72, 0x65, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63,
	0x6f, 0x6c, 0x2f, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x86, 0x01, 0x0a, 0x07, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04,
	0x6d, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6d, 0x65, 0x73, 0x73,
	0x12, 0x57, 0x0a, 0x11, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x5f, 0x73, 0x65, 0x74,
	0x74, 0x69, 0x6e, 0x67, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x76, 0x32,
	0x72, 0x61, 0x79, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x53, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74,
	0x79, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x10, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74,
	0x79, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x42, 0x3e, 0x0a, 0x1a, 0x63, 0x6f, 0x6d,
	0x2e, 0x76, 0x32, 0x72, 0x61, 0x79, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x78,
	0x79, 0x2e, 0x76, 0x6c, 0x65, 0x73, 0x73, 0x50, 0x01, 0x5a, 0x05, 0x76, 0x6c, 0x65, 0x73, 0x73,
	0xaa, 0x02, 0x16, 0x56, 0x32, 0x52, 0x61, 0x79, 0x2e, 0x43, 0x6f, 0x72, 0x65, 0x2e, 0x50, 0x72,
	0x6f, 0x78, 0x79, 0x2e, 0x56, 0x6c, 0x65, 0x73, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_v2ray_com_core_proxy_vless_account_proto_rawDescOnce sync.Once
	file_v2ray_com_core_proxy_vless_account_proto_rawDescData = file_v2ray_com_core_proxy_vless_account_proto_rawDesc
)

func file_v2ray_com_core_proxy_vless_account_proto_rawDescGZIP() []byte {
	file_v2ray_com_core_proxy_vless_account_proto_rawDescOnce.Do(func() {
		file_v2ray_com_core_proxy_vless_account_proto_rawDescData = protoimpl.X.CompressGZIP(file_v2ray_com_core_proxy_vless_account_proto_rawDescData)
	})
	return file_v2ray_com_core_proxy_vless_account_proto_rawDescData
}

var file_v2ray_com_core_proxy_vless_account_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_v2ray_com_core_proxy_vless_account_proto_goTypes = []interface{}{
	(*Account)(nil),                 // 0: v2ray.core.proxy.vless.Account
	(*protocol.SecurityConfig)(nil), // 1: v2ray.core.common.protocol.SecurityConfig
}
var file_v2ray_com_core_proxy_vless_account_proto_depIdxs = []int32{
	1, // 0: v2ray.core.proxy.vless.Account.security_settings:type_name -> v2ray.core.common.protocol.SecurityConfig
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_v2ray_com_core_proxy_vless_account_proto_init() }
func file_v2ray_com_core_proxy_vless_account_proto_init() {
	if File_v2ray_com_core_proxy_vless_account_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_v2ray_com_core_proxy_vless_account_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Account); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_v2ray_com_core_proxy_vless_account_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_v2ray_com_core_proxy_vless_account_proto_goTypes,
		DependencyIndexes: file_v2ray_com_core_proxy_vless_account_proto_depIdxs,
		MessageInfos:      file_v2ray_com_core_proxy_vless_account_proto_msgTypes,
	}.Build()
	File_v2ray_com_core_proxy_vless_account_proto = out.File
	file_v2ray_com_core_proxy_vless_account_proto_rawDesc = nil
	file_v2ray_com_core_proxy_vless_account_proto_goTypes = nil
	file_v2ray_com_core_proxy_vless_account_proto_depIdxs = nil
}
