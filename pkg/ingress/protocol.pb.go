// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.3
// source: pkg/ingress/protocol.proto

package ingress

import (
	empty "github.com/golang/protobuf/ptypes/empty"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_pkg_ingress_protocol_proto protoreflect.FileDescriptor

var file_pkg_ingress_protocol_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x70, 0x6b, 0x67, 0x2f, 0x69, 0x6e, 0x67, 0x72, 0x65, 0x73, 0x73, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x69, 0x6e,
	0x67, 0x72, 0x65, 0x73, 0x73, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1f, 0x70, 0x6b, 0x67, 0x2f, 0x69, 0x6e, 0x67, 0x72, 0x65, 0x73, 0x73, 0x2f,
	0x61, 0x64, 0x64, 0x2d, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x22, 0x70, 0x6b, 0x67, 0x2f, 0x69, 0x6e, 0x67, 0x72, 0x65, 0x73, 0x73,
	0x2f, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x2d, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x70, 0x6b, 0x67, 0x2f, 0x69, 0x6e, 0x67,
	0x72, 0x65, 0x73, 0x73, 0x2f, 0x67, 0x65, 0x74, 0x2d, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61,
	0x63, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x70, 0x6b, 0x67, 0x2f, 0x69,
	0x6e, 0x67, 0x72, 0x65, 0x73, 0x73, 0x2f, 0x61, 0x64, 0x64, 0x2d, 0x77, 0x6f, 0x72, 0x6b, 0x66,
	0x6c, 0x6f, 0x77, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x21, 0x70, 0x6b, 0x67, 0x2f, 0x69,
	0x6e, 0x67, 0x72, 0x65, 0x73, 0x73, 0x2f, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x2d, 0x77, 0x6f,
	0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x21, 0x70, 0x6b,
	0x67, 0x2f, 0x69, 0x6e, 0x67, 0x72, 0x65, 0x73, 0x73, 0x2f, 0x63, 0x61, 0x6e, 0x63, 0x65, 0x6c,
	0x2d, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1e, 0x70, 0x6b, 0x67, 0x2f, 0x69, 0x6e, 0x67, 0x72, 0x65, 0x73, 0x73, 0x2f, 0x67, 0x65, 0x74,
	0x2d, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1f, 0x70, 0x6b, 0x67, 0x2f, 0x69, 0x6e, 0x67, 0x72, 0x65, 0x73, 0x73, 0x2f, 0x67, 0x65, 0x74,
	0x2d, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x2b, 0x70, 0x6b, 0x67, 0x2f, 0x69, 0x6e, 0x67, 0x72, 0x65, 0x73, 0x73, 0x2f, 0x67, 0x65,
	0x74, 0x2d, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x2d, 0x62, 0x79, 0x2d, 0x77,
	0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x23, 0x70,
	0x6b, 0x67, 0x2f, 0x69, 0x6e, 0x67, 0x72, 0x65, 0x73, 0x73, 0x2f, 0x67, 0x65, 0x74, 0x2d, 0x69,
	0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x2d, 0x6c, 0x6f, 0x67, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x25, 0x70, 0x6b, 0x67, 0x2f, 0x69, 0x6e, 0x67, 0x72, 0x65, 0x73, 0x73, 0x2f,
	0x77, 0x61, 0x74, 0x63, 0x68, 0x2d, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x2d, 0x6c,
	0x6f, 0x67, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x23, 0x70, 0x6b, 0x67, 0x2f, 0x69,
	0x6e, 0x67, 0x72, 0x65, 0x73, 0x73, 0x2f, 0x67, 0x65, 0x74, 0x2d, 0x77, 0x6f, 0x72, 0x6b, 0x66,
	0x6c, 0x6f, 0x77, 0x2d, 0x6e, 0x61, 0x6d, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x22,
	0x70, 0x6b, 0x67, 0x2f, 0x69, 0x6e, 0x67, 0x72, 0x65, 0x73, 0x73, 0x2f, 0x67, 0x65, 0x74, 0x2d,
	0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x2d, 0x75, 0x69, 0x64, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1f, 0x70, 0x6b, 0x67, 0x2f, 0x69, 0x6e, 0x67, 0x72, 0x65, 0x73, 0x73, 0x2f,
	0x67, 0x65, 0x74, 0x2d, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x18, 0x70, 0x6b, 0x67, 0x2f, 0x69, 0x6e, 0x67, 0x72, 0x65, 0x73, 0x73,
	0x2f, 0x69, 0x6e, 0x76, 0x6f, 0x6b, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x21, 0x70,
	0x6b, 0x67, 0x2f, 0x69, 0x6e, 0x67, 0x72, 0x65, 0x73, 0x73, 0x2f, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x2d, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x21, 0x70, 0x6b, 0x67, 0x2f, 0x69, 0x6e, 0x67, 0x72, 0x65, 0x73, 0x73, 0x2f, 0x62, 0x72,
	0x6f, 0x61, 0x64, 0x63, 0x61, 0x73, 0x74, 0x2d, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x70, 0x6b, 0x67, 0x2f, 0x69, 0x6e, 0x67, 0x72, 0x65, 0x73, 0x73,
	0x2f, 0x67, 0x65, 0x74, 0x2d, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1f, 0x70, 0x6b, 0x67, 0x2f, 0x69, 0x6e, 0x67, 0x72, 0x65, 0x73, 0x73, 0x2f,
	0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x2d, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x70, 0x6b, 0x67, 0x2f, 0x69, 0x6e, 0x67, 0x72, 0x65, 0x73, 0x73,
	0x2f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2d, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x24, 0x70, 0x6b, 0x67, 0x2f, 0x69, 0x6e, 0x67, 0x72, 0x65, 0x73, 0x73,
	0x2f, 0x67, 0x65, 0x74, 0x2d, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x2d, 0x6c,
	0x6f, 0x67, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x22, 0x70, 0x6b, 0x67, 0x2f, 0x69,
	0x6e, 0x67, 0x72, 0x65, 0x73, 0x73, 0x2f, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x2d,
	0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2a, 0x70,
	0x6b, 0x67, 0x2f, 0x69, 0x6e, 0x67, 0x72, 0x65, 0x73, 0x73, 0x2f, 0x6c, 0x69, 0x73, 0x74, 0x2d,
	0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x2d, 0x76, 0x61, 0x72, 0x69, 0x61, 0x62,
	0x6c, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x29, 0x70, 0x6b, 0x67, 0x2f, 0x69,
	0x6e, 0x67, 0x72, 0x65, 0x73, 0x73, 0x2f, 0x6c, 0x69, 0x73, 0x74, 0x2d, 0x77, 0x6f, 0x72, 0x6b,
	0x66, 0x6c, 0x6f, 0x77, 0x2d, 0x76, 0x61, 0x72, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x28, 0x70, 0x6b, 0x67, 0x2f, 0x69, 0x6e, 0x67, 0x72, 0x65, 0x73,
	0x73, 0x2f, 0x67, 0x65, 0x74, 0x2d, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x2d,
	0x76, 0x61, 0x72, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x27,
	0x70, 0x6b, 0x67, 0x2f, 0x69, 0x6e, 0x67, 0x72, 0x65, 0x73, 0x73, 0x2f, 0x67, 0x65, 0x74, 0x2d,
	0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x2d, 0x76, 0x61, 0x72, 0x69, 0x61, 0x62, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x28, 0x70, 0x6b, 0x67, 0x2f, 0x69, 0x6e, 0x67,
	0x72, 0x65, 0x73, 0x73, 0x2f, 0x73, 0x65, 0x74, 0x2d, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61,
	0x63, 0x65, 0x2d, 0x76, 0x61, 0x72, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x27, 0x70, 0x6b, 0x67, 0x2f, 0x69, 0x6e, 0x67, 0x72, 0x65, 0x73, 0x73, 0x2f, 0x73,
	0x65, 0x74, 0x2d, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x2d, 0x76, 0x61, 0x72, 0x69,
	0x61, 0x62, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0x87, 0x14, 0x0a, 0x0f, 0x44,
	0x69, 0x72, 0x65, 0x6b, 0x74, 0x69, 0x76, 0x49, 0x6e, 0x67, 0x72, 0x65, 0x73, 0x73, 0x12, 0x4d,
	0x0a, 0x0c, 0x41, 0x64, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x1c,
	0x2e, 0x69, 0x6e, 0x67, 0x72, 0x65, 0x73, 0x73, 0x2e, 0x41, 0x64, 0x64, 0x4e, 0x61, 0x6d, 0x65,
	0x73, 0x70, 0x61, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x69,
	0x6e, 0x67, 0x72, 0x65, 0x73, 0x73, 0x2e, 0x41, 0x64, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70,
	0x61, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x56, 0x0a,
	0x0f, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65,
	0x12, 0x1f, 0x2e, 0x69, 0x6e, 0x67, 0x72, 0x65, 0x73, 0x73, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x20, 0x2e, 0x69, 0x6e, 0x67, 0x72, 0x65, 0x73, 0x73, 0x2e, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x50, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x4e, 0x61, 0x6d, 0x65,
	0x73, 0x70, 0x61, 0x63, 0x65, 0x73, 0x12, 0x1d, 0x2e, 0x69, 0x6e, 0x67, 0x72, 0x65, 0x73, 0x73,
	0x2e, 0x47, 0x65, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x69, 0x6e, 0x67, 0x72, 0x65, 0x73, 0x73, 0x2e,
	0x47, 0x65, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4a, 0x0a, 0x0b, 0x41, 0x64, 0x64, 0x57, 0x6f,
	0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x12, 0x1b, 0x2e, 0x69, 0x6e, 0x67, 0x72, 0x65, 0x73, 0x73,
	0x2e, 0x41, 0x64, 0x64, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x69, 0x6e, 0x67, 0x72, 0x65, 0x73, 0x73, 0x2e, 0x41, 0x64,
	0x64, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x53, 0x0a, 0x0e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x57, 0x6f, 0x72,
	0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x12, 0x1e, 0x2e, 0x69, 0x6e, 0x67, 0x72, 0x65, 0x73, 0x73, 0x2e,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x69, 0x6e, 0x67, 0x72, 0x65, 0x73, 0x73, 0x2e,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x5c, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x57,
	0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x42, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x21, 0x2e,
	0x69, 0x6e, 0x67, 0x72, 0x65, 0x73, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x57, 0x6f, 0x72, 0x6b, 0x66,
	0x6c, 0x6f, 0x77, 0x42, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x22, 0x2e, 0x69, 0x6e, 0x67, 0x72, 0x65, 0x73, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x57, 0x6f,
	0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x42, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x59, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x57, 0x6f, 0x72,
	0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x42, 0x79, 0x55, 0x69, 0x64, 0x12, 0x20, 0x2e, 0x69, 0x6e, 0x67,
	0x72, 0x65, 0x73, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77,
	0x42, 0x79, 0x55, 0x69, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x69,
	0x6e, 0x67, 0x72, 0x65, 0x73, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c,
	0x6f, 0x77, 0x42, 0x79, 0x55, 0x69, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x62, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77,
	0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x23, 0x2e, 0x69, 0x6e, 0x67, 0x72, 0x65,
	0x73, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x49, 0x6e,
	0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e,
	0x69, 0x6e, 0x67, 0x72, 0x65, 0x73, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x57, 0x6f, 0x72, 0x6b, 0x66,
	0x6c, 0x6f, 0x77, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x65, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x57, 0x6f, 0x72, 0x6b,
	0x66, 0x6c, 0x6f, 0x77, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x12, 0x24, 0x2e,
	0x69, 0x6e, 0x67, 0x72, 0x65, 0x73, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x57, 0x6f, 0x72, 0x6b, 0x66,
	0x6c, 0x6f, 0x77, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x69, 0x6e, 0x67, 0x72, 0x65, 0x73, 0x73, 0x2e, 0x47, 0x65,
	0x74, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63,
	0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x59, 0x0a, 0x10,
	0x47, 0x65, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x4c, 0x6f, 0x67, 0x73,
	0x12, 0x20, 0x2e, 0x69, 0x6e, 0x67, 0x72, 0x65, 0x73, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x4e, 0x61,
	0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x4c, 0x6f, 0x67, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x21, 0x2e, 0x69, 0x6e, 0x67, 0x72, 0x65, 0x73, 0x73, 0x2e, 0x47, 0x65, 0x74,
	0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x4c, 0x6f, 0x67, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x6b, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x49, 0x6e,
	0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x42, 0x79, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f,
	0x77, 0x12, 0x26, 0x2e, 0x69, 0x6e, 0x67, 0x72, 0x65, 0x73, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x49,
	0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x42, 0x79, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c,
	0x6f, 0x77, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e, 0x69, 0x6e, 0x67, 0x72,
	0x65, 0x73, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x73,
	0x42, 0x79, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x6e, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x57, 0x6f, 0x72, 0x6b, 0x66,
	0x6c, 0x6f, 0x77, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x4c, 0x6f, 0x67, 0x73, 0x12,
	0x27, 0x2e, 0x69, 0x6e, 0x67, 0x72, 0x65, 0x73, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x57, 0x6f, 0x72,
	0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x4c, 0x6f, 0x67,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x28, 0x2e, 0x69, 0x6e, 0x67, 0x72, 0x65,
	0x73, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x49, 0x6e,
	0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x4c, 0x6f, 0x67, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x76, 0x0a, 0x19, 0x57, 0x61, 0x74, 0x63, 0x68, 0x57, 0x6f, 0x72,
	0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x4c, 0x6f, 0x67,
	0x73, 0x12, 0x29, 0x2e, 0x69, 0x6e, 0x67, 0x72, 0x65, 0x73, 0x73, 0x2e, 0x57, 0x61, 0x74, 0x63,
	0x68, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63,
	0x65, 0x4c, 0x6f, 0x67, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2a, 0x2e, 0x69,
	0x6e, 0x67, 0x72, 0x65, 0x73, 0x73, 0x2e, 0x57, 0x61, 0x74, 0x63, 0x68, 0x57, 0x6f, 0x72, 0x6b,
	0x66, 0x6c, 0x6f, 0x77, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x4c, 0x6f, 0x67, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x30, 0x01, 0x12, 0x5a, 0x0a, 0x16,
	0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x49, 0x6e,
	0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x26, 0x2e, 0x69, 0x6e, 0x67, 0x72, 0x65, 0x73, 0x73,
	0x2e, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x49,
	0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x4d, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x57,
	0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x73, 0x12, 0x1c, 0x2e, 0x69, 0x6e, 0x67, 0x72, 0x65,
	0x73, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x69, 0x6e, 0x67, 0x72, 0x65, 0x73, 0x73,
	0x2e, 0x47, 0x65, 0x74, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x53, 0x0a, 0x0e, 0x49, 0x6e, 0x76, 0x6f, 0x6b,
	0x65, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x12, 0x1e, 0x2e, 0x69, 0x6e, 0x67, 0x72,
	0x65, 0x73, 0x73, 0x2e, 0x49, 0x6e, 0x76, 0x6f, 0x6b, 0x65, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c,
	0x6f, 0x77, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x69, 0x6e, 0x67, 0x72,
	0x65, 0x73, 0x73, 0x2e, 0x49, 0x6e, 0x76, 0x6f, 0x6b, 0x65, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c,
	0x6f, 0x77, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x53, 0x0a, 0x0e,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x12, 0x1e,
	0x2e, 0x69, 0x6e, 0x67, 0x72, 0x65, 0x73, 0x73, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x57,
	0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f,
	0x2e, 0x69, 0x6e, 0x67, 0x72, 0x65, 0x73, 0x73, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x57,
	0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x4a, 0x0a, 0x0e, 0x42, 0x72, 0x6f, 0x61, 0x64, 0x63, 0x61, 0x73, 0x74, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x12, 0x1e, 0x2e, 0x69, 0x6e, 0x67, 0x72, 0x65, 0x73, 0x73, 0x2e, 0x42, 0x72,
	0x6f, 0x61, 0x64, 0x63, 0x61, 0x73, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x47, 0x0a,
	0x0a, 0x47, 0x65, 0x74, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x73, 0x12, 0x1a, 0x2e, 0x69, 0x6e,
	0x67, 0x72, 0x65, 0x73, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x69, 0x6e, 0x67, 0x72, 0x65, 0x73,
	0x73, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x46, 0x0a, 0x0c, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x12, 0x1c, 0x2e, 0x69, 0x6e, 0x67, 0x72, 0x65, 0x73, 0x73,
	0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x44,
	0x0a, 0x0b, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x12, 0x1b, 0x2e,
	0x69, 0x6e, 0x67, 0x72, 0x65, 0x73, 0x73, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x53, 0x65, 0x63,
	0x72, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x22, 0x00, 0x12, 0x56, 0x0a, 0x0f, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77,
	0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x12, 0x1f, 0x2e, 0x69, 0x6e, 0x67, 0x72, 0x65, 0x73,
	0x73, 0x2e, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x69, 0x6e, 0x67, 0x72, 0x65,
	0x73, 0x73, 0x2e, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x4d, 0x65, 0x74, 0x72, 0x69,
	0x63, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x6b, 0x0a, 0x16,
	0x4c, 0x69, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x56, 0x61, 0x72,
	0x69, 0x61, 0x62, 0x6c, 0x65, 0x73, 0x12, 0x26, 0x2e, 0x69, 0x6e, 0x67, 0x72, 0x65, 0x73, 0x73,
	0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x56, 0x61,
	0x72, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x27,
	0x2e, 0x69, 0x6e, 0x67, 0x72, 0x65, 0x73, 0x73, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4e, 0x61, 0x6d,
	0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x56, 0x61, 0x72, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x68, 0x0a, 0x15, 0x4c, 0x69, 0x73,
	0x74, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x56, 0x61, 0x72, 0x69, 0x61, 0x62, 0x6c,
	0x65, 0x73, 0x12, 0x25, 0x2e, 0x69, 0x6e, 0x67, 0x72, 0x65, 0x73, 0x73, 0x2e, 0x4c, 0x69, 0x73,
	0x74, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x56, 0x61, 0x72, 0x69, 0x61, 0x62, 0x6c,
	0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x69, 0x6e, 0x67, 0x72,
	0x65, 0x73, 0x73, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77,
	0x56, 0x61, 0x72, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x67, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70,
	0x61, 0x63, 0x65, 0x56, 0x61, 0x72, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x24, 0x2e, 0x69, 0x6e,
	0x67, 0x72, 0x65, 0x73, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61,
	0x63, 0x65, 0x56, 0x61, 0x72, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x25, 0x2e, 0x69, 0x6e, 0x67, 0x72, 0x65, 0x73, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x4e,
	0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x56, 0x61, 0x72, 0x69, 0x61, 0x62, 0x6c, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x30, 0x01, 0x12, 0x64, 0x0a, 0x13,
	0x47, 0x65, 0x74, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x56, 0x61, 0x72, 0x69, 0x61,
	0x62, 0x6c, 0x65, 0x12, 0x23, 0x2e, 0x69, 0x6e, 0x67, 0x72, 0x65, 0x73, 0x73, 0x2e, 0x47, 0x65,
	0x74, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x56, 0x61, 0x72, 0x69, 0x61, 0x62, 0x6c,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x69, 0x6e, 0x67, 0x72, 0x65,
	0x73, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x56, 0x61,
	0x72, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x30, 0x01, 0x12, 0x58, 0x0a, 0x14, 0x53, 0x65, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61,
	0x63, 0x65, 0x56, 0x61, 0x72, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x24, 0x2e, 0x69, 0x6e, 0x67,
	0x72, 0x65, 0x73, 0x73, 0x2e, 0x53, 0x65, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63,
	0x65, 0x56, 0x61, 0x72, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x28, 0x01, 0x12, 0x56, 0x0a, 0x13,
	0x53, 0x65, 0x74, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x56, 0x61, 0x72, 0x69, 0x61,
	0x62, 0x6c, 0x65, 0x12, 0x23, 0x2e, 0x69, 0x6e, 0x67, 0x72, 0x65, 0x73, 0x73, 0x2e, 0x53, 0x65,
	0x74, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x56, 0x61, 0x72, 0x69, 0x61, 0x62, 0x6c,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x22, 0x00, 0x28, 0x01, 0x42, 0x29, 0x5a, 0x27, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x76, 0x6f, 0x72, 0x74, 0x65, 0x69, 0x6c, 0x2f, 0x64, 0x69, 0x72, 0x65, 0x6b,
	0x74, 0x69, 0x76, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x69, 0x6e, 0x67, 0x72, 0x65, 0x73, 0x73, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_pkg_ingress_protocol_proto_goTypes = []interface{}{
	(*AddNamespaceRequest)(nil),               // 0: ingress.AddNamespaceRequest
	(*DeleteNamespaceRequest)(nil),            // 1: ingress.DeleteNamespaceRequest
	(*GetNamespacesRequest)(nil),              // 2: ingress.GetNamespacesRequest
	(*AddWorkflowRequest)(nil),                // 3: ingress.AddWorkflowRequest
	(*DeleteWorkflowRequest)(nil),             // 4: ingress.DeleteWorkflowRequest
	(*GetWorkflowByNameRequest)(nil),          // 5: ingress.GetWorkflowByNameRequest
	(*GetWorkflowByUidRequest)(nil),           // 6: ingress.GetWorkflowByUidRequest
	(*GetWorkflowInstanceRequest)(nil),        // 7: ingress.GetWorkflowInstanceRequest
	(*GetWorkflowInstancesRequest)(nil),       // 8: ingress.GetWorkflowInstancesRequest
	(*GetNamespaceLogsRequest)(nil),           // 9: ingress.GetNamespaceLogsRequest
	(*GetInstancesByWorkflowRequest)(nil),     // 10: ingress.GetInstancesByWorkflowRequest
	(*GetWorkflowInstanceLogsRequest)(nil),    // 11: ingress.GetWorkflowInstanceLogsRequest
	(*WatchWorkflowInstanceLogsRequest)(nil),  // 12: ingress.WatchWorkflowInstanceLogsRequest
	(*CancelWorkflowInstanceRequest)(nil),     // 13: ingress.CancelWorkflowInstanceRequest
	(*GetWorkflowsRequest)(nil),               // 14: ingress.GetWorkflowsRequest
	(*InvokeWorkflowRequest)(nil),             // 15: ingress.InvokeWorkflowRequest
	(*UpdateWorkflowRequest)(nil),             // 16: ingress.UpdateWorkflowRequest
	(*BroadcastEventRequest)(nil),             // 17: ingress.BroadcastEventRequest
	(*GetSecretsRequest)(nil),                 // 18: ingress.GetSecretsRequest
	(*DeleteSecretRequest)(nil),               // 19: ingress.DeleteSecretRequest
	(*StoreSecretRequest)(nil),                // 20: ingress.StoreSecretRequest
	(*WorkflowMetricsRequest)(nil),            // 21: ingress.WorkflowMetricsRequest
	(*ListNamespaceVariablesRequest)(nil),     // 22: ingress.ListNamespaceVariablesRequest
	(*ListWorkflowVariablesRequest)(nil),      // 23: ingress.ListWorkflowVariablesRequest
	(*GetNamespaceVariableRequest)(nil),       // 24: ingress.GetNamespaceVariableRequest
	(*GetWorkflowVariableRequest)(nil),        // 25: ingress.GetWorkflowVariableRequest
	(*SetNamespaceVariableRequest)(nil),       // 26: ingress.SetNamespaceVariableRequest
	(*SetWorkflowVariableRequest)(nil),        // 27: ingress.SetWorkflowVariableRequest
	(*AddNamespaceResponse)(nil),              // 28: ingress.AddNamespaceResponse
	(*DeleteNamespaceResponse)(nil),           // 29: ingress.DeleteNamespaceResponse
	(*GetNamespacesResponse)(nil),             // 30: ingress.GetNamespacesResponse
	(*AddWorkflowResponse)(nil),               // 31: ingress.AddWorkflowResponse
	(*DeleteWorkflowResponse)(nil),            // 32: ingress.DeleteWorkflowResponse
	(*GetWorkflowByNameResponse)(nil),         // 33: ingress.GetWorkflowByNameResponse
	(*GetWorkflowByUidResponse)(nil),          // 34: ingress.GetWorkflowByUidResponse
	(*GetWorkflowInstanceResponse)(nil),       // 35: ingress.GetWorkflowInstanceResponse
	(*GetWorkflowInstancesResponse)(nil),      // 36: ingress.GetWorkflowInstancesResponse
	(*GetNamespaceLogsResponse)(nil),          // 37: ingress.GetNamespaceLogsResponse
	(*GetInstancesByWorkflowResponse)(nil),    // 38: ingress.GetInstancesByWorkflowResponse
	(*GetWorkflowInstanceLogsResponse)(nil),   // 39: ingress.GetWorkflowInstanceLogsResponse
	(*WatchWorkflowInstanceLogsResponse)(nil), // 40: ingress.WatchWorkflowInstanceLogsResponse
	(*empty.Empty)(nil),                       // 41: google.protobuf.Empty
	(*GetWorkflowsResponse)(nil),              // 42: ingress.GetWorkflowsResponse
	(*InvokeWorkflowResponse)(nil),            // 43: ingress.InvokeWorkflowResponse
	(*UpdateWorkflowResponse)(nil),            // 44: ingress.UpdateWorkflowResponse
	(*GetSecretsResponse)(nil),                // 45: ingress.GetSecretsResponse
	(*WorkflowMetricsResponse)(nil),           // 46: ingress.WorkflowMetricsResponse
	(*ListNamespaceVariablesResponse)(nil),    // 47: ingress.ListNamespaceVariablesResponse
	(*ListWorkflowVariablesResponse)(nil),     // 48: ingress.ListWorkflowVariablesResponse
	(*GetNamespaceVariableResponse)(nil),      // 49: ingress.GetNamespaceVariableResponse
	(*GetWorkflowVariableResponse)(nil),       // 50: ingress.GetWorkflowVariableResponse
}
var file_pkg_ingress_protocol_proto_depIdxs = []int32{
	0,  // 0: ingress.DirektivIngress.AddNamespace:input_type -> ingress.AddNamespaceRequest
	1,  // 1: ingress.DirektivIngress.DeleteNamespace:input_type -> ingress.DeleteNamespaceRequest
	2,  // 2: ingress.DirektivIngress.GetNamespaces:input_type -> ingress.GetNamespacesRequest
	3,  // 3: ingress.DirektivIngress.AddWorkflow:input_type -> ingress.AddWorkflowRequest
	4,  // 4: ingress.DirektivIngress.DeleteWorkflow:input_type -> ingress.DeleteWorkflowRequest
	5,  // 5: ingress.DirektivIngress.GetWorkflowByName:input_type -> ingress.GetWorkflowByNameRequest
	6,  // 6: ingress.DirektivIngress.GetWorkflowByUid:input_type -> ingress.GetWorkflowByUidRequest
	7,  // 7: ingress.DirektivIngress.GetWorkflowInstance:input_type -> ingress.GetWorkflowInstanceRequest
	8,  // 8: ingress.DirektivIngress.GetWorkflowInstances:input_type -> ingress.GetWorkflowInstancesRequest
	9,  // 9: ingress.DirektivIngress.GetNamespaceLogs:input_type -> ingress.GetNamespaceLogsRequest
	10, // 10: ingress.DirektivIngress.GetInstancesByWorkflow:input_type -> ingress.GetInstancesByWorkflowRequest
	11, // 11: ingress.DirektivIngress.GetWorkflowInstanceLogs:input_type -> ingress.GetWorkflowInstanceLogsRequest
	12, // 12: ingress.DirektivIngress.WatchWorkflowInstanceLogs:input_type -> ingress.WatchWorkflowInstanceLogsRequest
	13, // 13: ingress.DirektivIngress.CancelWorkflowInstance:input_type -> ingress.CancelWorkflowInstanceRequest
	14, // 14: ingress.DirektivIngress.GetWorkflows:input_type -> ingress.GetWorkflowsRequest
	15, // 15: ingress.DirektivIngress.InvokeWorkflow:input_type -> ingress.InvokeWorkflowRequest
	16, // 16: ingress.DirektivIngress.UpdateWorkflow:input_type -> ingress.UpdateWorkflowRequest
	17, // 17: ingress.DirektivIngress.BroadcastEvent:input_type -> ingress.BroadcastEventRequest
	18, // 18: ingress.DirektivIngress.GetSecrets:input_type -> ingress.GetSecretsRequest
	19, // 19: ingress.DirektivIngress.DeleteSecret:input_type -> ingress.DeleteSecretRequest
	20, // 20: ingress.DirektivIngress.StoreSecret:input_type -> ingress.StoreSecretRequest
	21, // 21: ingress.DirektivIngress.WorkflowMetrics:input_type -> ingress.WorkflowMetricsRequest
	22, // 22: ingress.DirektivIngress.ListNamespaceVariables:input_type -> ingress.ListNamespaceVariablesRequest
	23, // 23: ingress.DirektivIngress.ListWorkflowVariables:input_type -> ingress.ListWorkflowVariablesRequest
	24, // 24: ingress.DirektivIngress.GetNamespaceVariable:input_type -> ingress.GetNamespaceVariableRequest
	25, // 25: ingress.DirektivIngress.GetWorkflowVariable:input_type -> ingress.GetWorkflowVariableRequest
	26, // 26: ingress.DirektivIngress.SetNamespaceVariable:input_type -> ingress.SetNamespaceVariableRequest
	27, // 27: ingress.DirektivIngress.SetWorkflowVariable:input_type -> ingress.SetWorkflowVariableRequest
	28, // 28: ingress.DirektivIngress.AddNamespace:output_type -> ingress.AddNamespaceResponse
	29, // 29: ingress.DirektivIngress.DeleteNamespace:output_type -> ingress.DeleteNamespaceResponse
	30, // 30: ingress.DirektivIngress.GetNamespaces:output_type -> ingress.GetNamespacesResponse
	31, // 31: ingress.DirektivIngress.AddWorkflow:output_type -> ingress.AddWorkflowResponse
	32, // 32: ingress.DirektivIngress.DeleteWorkflow:output_type -> ingress.DeleteWorkflowResponse
	33, // 33: ingress.DirektivIngress.GetWorkflowByName:output_type -> ingress.GetWorkflowByNameResponse
	34, // 34: ingress.DirektivIngress.GetWorkflowByUid:output_type -> ingress.GetWorkflowByUidResponse
	35, // 35: ingress.DirektivIngress.GetWorkflowInstance:output_type -> ingress.GetWorkflowInstanceResponse
	36, // 36: ingress.DirektivIngress.GetWorkflowInstances:output_type -> ingress.GetWorkflowInstancesResponse
	37, // 37: ingress.DirektivIngress.GetNamespaceLogs:output_type -> ingress.GetNamespaceLogsResponse
	38, // 38: ingress.DirektivIngress.GetInstancesByWorkflow:output_type -> ingress.GetInstancesByWorkflowResponse
	39, // 39: ingress.DirektivIngress.GetWorkflowInstanceLogs:output_type -> ingress.GetWorkflowInstanceLogsResponse
	40, // 40: ingress.DirektivIngress.WatchWorkflowInstanceLogs:output_type -> ingress.WatchWorkflowInstanceLogsResponse
	41, // 41: ingress.DirektivIngress.CancelWorkflowInstance:output_type -> google.protobuf.Empty
	42, // 42: ingress.DirektivIngress.GetWorkflows:output_type -> ingress.GetWorkflowsResponse
	43, // 43: ingress.DirektivIngress.InvokeWorkflow:output_type -> ingress.InvokeWorkflowResponse
	44, // 44: ingress.DirektivIngress.UpdateWorkflow:output_type -> ingress.UpdateWorkflowResponse
	41, // 45: ingress.DirektivIngress.BroadcastEvent:output_type -> google.protobuf.Empty
	45, // 46: ingress.DirektivIngress.GetSecrets:output_type -> ingress.GetSecretsResponse
	41, // 47: ingress.DirektivIngress.DeleteSecret:output_type -> google.protobuf.Empty
	41, // 48: ingress.DirektivIngress.StoreSecret:output_type -> google.protobuf.Empty
	46, // 49: ingress.DirektivIngress.WorkflowMetrics:output_type -> ingress.WorkflowMetricsResponse
	47, // 50: ingress.DirektivIngress.ListNamespaceVariables:output_type -> ingress.ListNamespaceVariablesResponse
	48, // 51: ingress.DirektivIngress.ListWorkflowVariables:output_type -> ingress.ListWorkflowVariablesResponse
	49, // 52: ingress.DirektivIngress.GetNamespaceVariable:output_type -> ingress.GetNamespaceVariableResponse
	50, // 53: ingress.DirektivIngress.GetWorkflowVariable:output_type -> ingress.GetWorkflowVariableResponse
	41, // 54: ingress.DirektivIngress.SetNamespaceVariable:output_type -> google.protobuf.Empty
	41, // 55: ingress.DirektivIngress.SetWorkflowVariable:output_type -> google.protobuf.Empty
	28, // [28:56] is the sub-list for method output_type
	0,  // [0:28] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_pkg_ingress_protocol_proto_init() }
func file_pkg_ingress_protocol_proto_init() {
	if File_pkg_ingress_protocol_proto != nil {
		return
	}
	file_pkg_ingress_add_namespace_proto_init()
	file_pkg_ingress_delete_namespace_proto_init()
	file_pkg_ingress_get_namespaces_proto_init()
	file_pkg_ingress_add_workflow_proto_init()
	file_pkg_ingress_delete_workflow_proto_init()
	file_pkg_ingress_cancel_instance_proto_init()
	file_pkg_ingress_get_instance_proto_init()
	file_pkg_ingress_get_instances_proto_init()
	file_pkg_ingress_get_instances_by_workflow_proto_init()
	file_pkg_ingress_get_instance_logs_proto_init()
	file_pkg_ingress_watch_instance_logs_proto_init()
	file_pkg_ingress_get_workflow_name_proto_init()
	file_pkg_ingress_get_workflow_uid_proto_init()
	file_pkg_ingress_get_workflows_proto_init()
	file_pkg_ingress_invoke_proto_init()
	file_pkg_ingress_update_workflow_proto_init()
	file_pkg_ingress_broadcast_event_proto_init()
	file_pkg_ingress_get_secrets_proto_init()
	file_pkg_ingress_delete_secret_proto_init()
	file_pkg_ingress_store_secret_proto_init()
	file_pkg_ingress_get_namespace_logs_proto_init()
	file_pkg_ingress_workflow_metrics_proto_init()
	file_pkg_ingress_list_namespace_variables_proto_init()
	file_pkg_ingress_list_workflow_variables_proto_init()
	file_pkg_ingress_get_namespace_variable_proto_init()
	file_pkg_ingress_get_workflow_variable_proto_init()
	file_pkg_ingress_set_namespace_variable_proto_init()
	file_pkg_ingress_set_workflow_variable_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_pkg_ingress_protocol_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pkg_ingress_protocol_proto_goTypes,
		DependencyIndexes: file_pkg_ingress_protocol_proto_depIdxs,
	}.Build()
	File_pkg_ingress_protocol_proto = out.File
	file_pkg_ingress_protocol_proto_rawDesc = nil
	file_pkg_ingress_protocol_proto_goTypes = nil
	file_pkg_ingress_protocol_proto_depIdxs = nil
}
