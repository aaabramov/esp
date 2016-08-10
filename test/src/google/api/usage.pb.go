// Code generated by protoc-gen-go.
// source: google/api/usage.proto
// DO NOT EDIT!

package api

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Service access types.
//
// Access to restricted API features is always controlled by
// [visibility][google.api.Visibility], independent of the ServiceAccess type.
//
type Usage_ServiceAccess int32

const (
	// The service can only be seen/used by users identified in the service's
	// access control policy.
	//
	// If the service has not been whitelisted by your domain administrator
	// for out-of-org publishing, then this mode will be treated like
	// ORG_RESTRICTED.
	Usage_RESTRICTED Usage_ServiceAccess = 0
	// The service can be seen/used by anyone.
	//
	// If the service has not been whitelisted by your domain administrator
	// for out-of-org publishing, then this mode will be treated like
	// ORG_PUBLIC.
	//
	// The discovery document for the service will also be public and allow
	// unregistered access.
	Usage_PUBLIC Usage_ServiceAccess = 1
	// The service can be seen/used by users identified in the service's
	// access control policy and they are within the organization that owns the
	// service.
	//
	// Access is further constrained to the group
	// controlled by the administrator of the project/org that owns the
	// service.
	Usage_ORG_RESTRICTED Usage_ServiceAccess = 2
	// The service can be seen/used by the group of users controlled by the
	// administrator of the project/org that owns the service.
	Usage_ORG_PUBLIC Usage_ServiceAccess = 3
)

var Usage_ServiceAccess_name = map[int32]string{
	0: "RESTRICTED",
	1: "PUBLIC",
	2: "ORG_RESTRICTED",
	3: "ORG_PUBLIC",
}
var Usage_ServiceAccess_value = map[string]int32{
	"RESTRICTED":     0,
	"PUBLIC":         1,
	"ORG_RESTRICTED": 2,
	"ORG_PUBLIC":     3,
}

func (x Usage_ServiceAccess) String() string {
	return proto.EnumName(Usage_ServiceAccess_name, int32(x))
}
func (Usage_ServiceAccess) EnumDescriptor() ([]byte, []int) { return fileDescriptor18, []int{0, 0} }

// Configuration controlling usage of a service.
type Usage struct {
	// Controls which users can see or activate the service.
	ServiceAccess Usage_ServiceAccess `protobuf:"varint,4,opt,name=service_access,json=serviceAccess,enum=google.api.Usage_ServiceAccess" json:"service_access,omitempty"`
	// Requirements that must be satisfied before a consumer project can use the
	// service. Each requirement is of the form <service.name>/<requirement-id>;
	// for example 'serviceusage.googleapis.com/billing-enabled'.
	Requirements []string `protobuf:"bytes,1,rep,name=requirements" json:"requirements,omitempty"`
	// Services that must be activated in order for this service to be used.
	// The set of services activated as a result of these relations are all
	// activated in parallel with no guaranteed order of activation.
	// Each string is a service name, e.g. `calendar.googleapis.com`.
	DependsOnServices []string `protobuf:"bytes,2,rep,name=depends_on_services,json=dependsOnServices" json:"depends_on_services,omitempty"`
	// Services that must be contacted before a consumer can begin using the
	// service. Each service will be contacted in sequence, and, if any activation
	// call fails, the entire activation will fail. Each hook is of the form
	// <service.name>/<hook-id>, where <hook-id> is optional; for example:
	// 'robotservice.googleapis.com/default'.
	ActivationHooks []string `protobuf:"bytes,3,rep,name=activation_hooks,json=activationHooks" json:"activation_hooks,omitempty"`
	// Services that must be contacted before a consumer can deactivate a
	// service. Each service will be contacted in sequence, and, if any
	// deactivation call fails, the entire deactivation will fail. Each hook is
	// of the form <service.name>/<hook-id>, where <hook-id> is optional; for
	// example:
	// 'compute.googleapis.com/'.
	DeactivationHooks []string `protobuf:"bytes,5,rep,name=deactivation_hooks,json=deactivationHooks" json:"deactivation_hooks,omitempty"`
	// Individual rules for configuring usage on selected methods.
	Rules []*UsageRule `protobuf:"bytes,6,rep,name=rules" json:"rules,omitempty"`
}

func (m *Usage) Reset()                    { *m = Usage{} }
func (m *Usage) String() string            { return proto.CompactTextString(m) }
func (*Usage) ProtoMessage()               {}
func (*Usage) Descriptor() ([]byte, []int) { return fileDescriptor18, []int{0} }

func (m *Usage) GetRules() []*UsageRule {
	if m != nil {
		return m.Rules
	}
	return nil
}

// Usage configuration rules for the service.
//
// NOTE: Under development.
//
//
// Use this rule to configure unregistered calls for the service. Unregistered
// calls are calls that do not contain consumer project identity.
// (Example: calls that do not contain an API key).
// By default, API methods do not allow unregistered calls, and each method call
// must be identified by a consumer project identity. Use this rule to
// allow/disallow unregistered calls.
//
// Example of an API that wants to allow unregistered calls for entire service.
//
//     usage:
//       rules:
//       - selector: "*"
//         allow_unregistered_calls: true
//
// Example of a method that wants to allow unregistered calls.
//
//     usage:
//       rules:
//       - selector: "google.example.library.v1.LibraryService.CreateBook"
//         allow_unregistered_calls: true
type UsageRule struct {
	// Selects the methods to which this rule applies. Use '*' to indicate all
	// methods in all APIs.
	//
	// Refer to [selector][google.api.DocumentationRule.selector] for syntax details.
	Selector string `protobuf:"bytes,1,opt,name=selector" json:"selector,omitempty"`
	// True, if the method allows unregistered calls; false otherwise.
	AllowUnregisteredCalls bool `protobuf:"varint,2,opt,name=allow_unregistered_calls,json=allowUnregisteredCalls" json:"allow_unregistered_calls,omitempty"`
}

func (m *UsageRule) Reset()                    { *m = UsageRule{} }
func (m *UsageRule) String() string            { return proto.CompactTextString(m) }
func (*UsageRule) ProtoMessage()               {}
func (*UsageRule) Descriptor() ([]byte, []int) { return fileDescriptor18, []int{1} }

func init() {
	proto.RegisterType((*Usage)(nil), "google.api.Usage")
	proto.RegisterType((*UsageRule)(nil), "google.api.UsageRule")
	proto.RegisterEnum("google.api.Usage_ServiceAccess", Usage_ServiceAccess_name, Usage_ServiceAccess_value)
}

func init() { proto.RegisterFile("google/api/usage.proto", fileDescriptor18) }

var fileDescriptor18 = []byte{
	// 348 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x64, 0x92, 0xc1, 0x4f, 0xe2, 0x40,
	0x14, 0xc6, 0xb7, 0x74, 0x21, 0xf0, 0x76, 0xe9, 0xb2, 0x63, 0x24, 0x8d, 0x07, 0x6d, 0x7a, 0xaa,
	0x31, 0x96, 0x04, 0x2f, 0x5e, 0x05, 0x51, 0x49, 0x4c, 0x20, 0x03, 0x9c, 0x9b, 0xb1, 0x7d, 0xa9,
	0x8d, 0x43, 0xa7, 0xce, 0xb4, 0xf8, 0xc7, 0x7b, 0x31, 0x1d, 0x2a, 0x14, 0x39, 0xbe, 0xef, 0xf7,
	0x6b, 0x5f, 0xbe, 0xc9, 0x83, 0x7e, 0x2c, 0x44, 0xcc, 0x71, 0xc0, 0xb2, 0x64, 0x50, 0x28, 0x16,
	0xa3, 0x9f, 0x49, 0x91, 0x0b, 0x02, 0xdb, 0xdc, 0x67, 0x59, 0xe2, 0x7e, 0x36, 0xa0, 0xb9, 0x2a,
	0x19, 0x79, 0x00, 0x4b, 0xa1, 0xdc, 0x24, 0x21, 0x06, 0x2c, 0x0c, 0x51, 0x29, 0xfb, 0xb7, 0x63,
	0x78, 0xd6, 0xf0, 0xc2, 0xdf, 0xeb, 0xbe, 0x56, 0xfd, 0xc5, 0xd6, 0xbb, 0xd3, 0x1a, 0xed, 0xaa,
	0xfa, 0x48, 0x5c, 0xf8, 0x2b, 0xf1, 0xbd, 0x48, 0x24, 0xae, 0x31, 0xcd, 0x95, 0x6d, 0x38, 0xa6,
	0xd7, 0xa1, 0x07, 0x19, 0xf1, 0xe1, 0x24, 0xc2, 0x0c, 0xd3, 0x48, 0x05, 0x22, 0x0d, 0xaa, 0xef,
	0x95, 0xdd, 0xd0, 0xea, 0xff, 0x0a, 0xcd, 0xd2, 0x6a, 0x8f, 0x22, 0x97, 0xd0, 0x63, 0x61, 0x9e,
	0x6c, 0x58, 0x9e, 0x88, 0x34, 0x78, 0x15, 0xe2, 0x4d, 0xd9, 0xa6, 0x96, 0xff, 0xed, 0xf3, 0xa7,
	0x32, 0x26, 0xd7, 0x40, 0x22, 0x3c, 0x92, 0x9b, 0xdf, 0x7f, 0xfe, 0xa9, 0x5f, 0x41, 0x53, 0x16,
	0x1c, 0x95, 0xdd, 0x72, 0x4c, 0xef, 0xcf, 0xf0, 0xf4, 0xa8, 0x2c, 0x2d, 0x38, 0xd2, 0xad, 0xe3,
	0xce, 0xa0, 0x7b, 0x50, 0x9d, 0x58, 0x00, 0x74, 0xb2, 0x58, 0xd2, 0xe9, 0x78, 0x39, 0xb9, 0xef,
	0xfd, 0x22, 0x00, 0xad, 0xf9, 0x6a, 0xf4, 0x3c, 0x1d, 0xf7, 0x0c, 0x42, 0xc0, 0x9a, 0xd1, 0xc7,
	0xa0, 0xc6, 0x1b, 0xa5, 0x5f, 0x66, 0x95, 0x63, 0xba, 0x0c, 0x3a, 0xbb, 0x25, 0xe4, 0x0c, 0xda,
	0x0a, 0x39, 0x86, 0xb9, 0x90, 0xb6, 0xe1, 0x18, 0x5e, 0x87, 0xee, 0x66, 0x72, 0x0b, 0x36, 0xe3,
	0x5c, 0x7c, 0x04, 0x45, 0x2a, 0x31, 0x4e, 0x54, 0x8e, 0x12, 0xa3, 0x20, 0x64, 0x9c, 0x97, 0xaf,
	0x66, 0x78, 0x6d, 0xda, 0xd7, 0x7c, 0x55, 0xc3, 0xe3, 0x92, 0x8e, 0xce, 0xc1, 0x0a, 0xc5, 0xba,
	0x56, 0x6b, 0x04, 0x7a, 0xe5, 0xbc, 0x3c, 0x85, 0xb9, 0xf1, 0xd2, 0xd2, 0x37, 0x71, 0xf3, 0x15,
	0x00, 0x00, 0xff, 0xff, 0xf6, 0xb2, 0x86, 0xb5, 0x2d, 0x02, 0x00, 0x00,
}