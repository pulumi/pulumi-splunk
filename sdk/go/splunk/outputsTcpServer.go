// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package splunk

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## # Resource: OutputsTcpServer
//
// Access data forwarding configurations.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-splunk/sdk/go/splunk"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := splunk.NewOutputsTcpServer(ctx, "tcpServer", &splunk.OutputsTcpServerArgs{
// 			SslAltNameToCheck: pulumi.String("old-host"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type OutputsTcpServer struct {
	pulumi.CustomResourceState

	// The app/user context that is the namespace for the resource
	Acl OutputsTcpServerAclOutput `pulumi:"acl"`
	// If true, disables the group.
	Disabled pulumi.BoolOutput `pulumi:"disabled"`
	// Valid values: (clone | balance | autobalance)
	// The data distribution method used when two or more servers exist in the same forwarder group.
	Method pulumi.StringOutput `pulumi:"method"`
	// <host>:<port> of the Splunk receiver. <host> can be either an ip address or server name. <port> is the that port that the Splunk receiver is listening on.
	Name pulumi.StringOutput `pulumi:"name"`
	// The alternate name to match in the remote server's SSL certificate.
	SslAltNameToCheck pulumi.StringOutput `pulumi:"sslAltNameToCheck"`
	// Path to the client certificate. If specified, connection uses SSL.
	SslCertPath pulumi.StringOutput `pulumi:"sslCertPath"`
	// SSL Cipher in the form ALL:!aNULL:!eNULL:!LOW:!EXP:RC4+RSA:+HIGH:+MEDIUM
	SslCipher pulumi.StringOutput `pulumi:"sslCipher"`
	// Check the common name of the server's certificate against this name.
	// If there is no match, assume that Splunk Enterprise is not authenticated against this server. You must specify this setting if sslVerifyServerCert is true.
	SslCommonNameToCheck pulumi.StringOutput `pulumi:"sslCommonNameToCheck"`
	// The password associated with the CAcert.
	// The default Splunk Enterprise CAcert uses the password "password."
	SslPassword pulumi.StringOutput `pulumi:"sslPassword"`
	// The path to the root certificate authority file.
	SslRootCaPath pulumi.StringOutput `pulumi:"sslRootCaPath"`
	// If true, make sure that the server you are connecting to is a valid one (authenticated). Both the common name and the alternate name of the server are then checked for a match.
	SslVerifyServerCert pulumi.BoolOutput `pulumi:"sslVerifyServerCert"`
}

// NewOutputsTcpServer registers a new resource with the given unique name, arguments, and options.
func NewOutputsTcpServer(ctx *pulumi.Context,
	name string, args *OutputsTcpServerArgs, opts ...pulumi.ResourceOption) (*OutputsTcpServer, error) {
	if args == nil {
		args = &OutputsTcpServerArgs{}
	}

	var resource OutputsTcpServer
	err := ctx.RegisterResource("splunk:index/outputsTcpServer:OutputsTcpServer", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetOutputsTcpServer gets an existing OutputsTcpServer resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetOutputsTcpServer(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *OutputsTcpServerState, opts ...pulumi.ResourceOption) (*OutputsTcpServer, error) {
	var resource OutputsTcpServer
	err := ctx.ReadResource("splunk:index/outputsTcpServer:OutputsTcpServer", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering OutputsTcpServer resources.
type outputsTcpServerState struct {
	// The app/user context that is the namespace for the resource
	Acl *OutputsTcpServerAcl `pulumi:"acl"`
	// If true, disables the group.
	Disabled *bool `pulumi:"disabled"`
	// Valid values: (clone | balance | autobalance)
	// The data distribution method used when two or more servers exist in the same forwarder group.
	Method *string `pulumi:"method"`
	// <host>:<port> of the Splunk receiver. <host> can be either an ip address or server name. <port> is the that port that the Splunk receiver is listening on.
	Name *string `pulumi:"name"`
	// The alternate name to match in the remote server's SSL certificate.
	SslAltNameToCheck *string `pulumi:"sslAltNameToCheck"`
	// Path to the client certificate. If specified, connection uses SSL.
	SslCertPath *string `pulumi:"sslCertPath"`
	// SSL Cipher in the form ALL:!aNULL:!eNULL:!LOW:!EXP:RC4+RSA:+HIGH:+MEDIUM
	SslCipher *string `pulumi:"sslCipher"`
	// Check the common name of the server's certificate against this name.
	// If there is no match, assume that Splunk Enterprise is not authenticated against this server. You must specify this setting if sslVerifyServerCert is true.
	SslCommonNameToCheck *string `pulumi:"sslCommonNameToCheck"`
	// The password associated with the CAcert.
	// The default Splunk Enterprise CAcert uses the password "password."
	SslPassword *string `pulumi:"sslPassword"`
	// The path to the root certificate authority file.
	SslRootCaPath *string `pulumi:"sslRootCaPath"`
	// If true, make sure that the server you are connecting to is a valid one (authenticated). Both the common name and the alternate name of the server are then checked for a match.
	SslVerifyServerCert *bool `pulumi:"sslVerifyServerCert"`
}

type OutputsTcpServerState struct {
	// The app/user context that is the namespace for the resource
	Acl OutputsTcpServerAclPtrInput
	// If true, disables the group.
	Disabled pulumi.BoolPtrInput
	// Valid values: (clone | balance | autobalance)
	// The data distribution method used when two or more servers exist in the same forwarder group.
	Method pulumi.StringPtrInput
	// <host>:<port> of the Splunk receiver. <host> can be either an ip address or server name. <port> is the that port that the Splunk receiver is listening on.
	Name pulumi.StringPtrInput
	// The alternate name to match in the remote server's SSL certificate.
	SslAltNameToCheck pulumi.StringPtrInput
	// Path to the client certificate. If specified, connection uses SSL.
	SslCertPath pulumi.StringPtrInput
	// SSL Cipher in the form ALL:!aNULL:!eNULL:!LOW:!EXP:RC4+RSA:+HIGH:+MEDIUM
	SslCipher pulumi.StringPtrInput
	// Check the common name of the server's certificate against this name.
	// If there is no match, assume that Splunk Enterprise is not authenticated against this server. You must specify this setting if sslVerifyServerCert is true.
	SslCommonNameToCheck pulumi.StringPtrInput
	// The password associated with the CAcert.
	// The default Splunk Enterprise CAcert uses the password "password."
	SslPassword pulumi.StringPtrInput
	// The path to the root certificate authority file.
	SslRootCaPath pulumi.StringPtrInput
	// If true, make sure that the server you are connecting to is a valid one (authenticated). Both the common name and the alternate name of the server are then checked for a match.
	SslVerifyServerCert pulumi.BoolPtrInput
}

func (OutputsTcpServerState) ElementType() reflect.Type {
	return reflect.TypeOf((*outputsTcpServerState)(nil)).Elem()
}

type outputsTcpServerArgs struct {
	// The app/user context that is the namespace for the resource
	Acl *OutputsTcpServerAcl `pulumi:"acl"`
	// If true, disables the group.
	Disabled *bool `pulumi:"disabled"`
	// Valid values: (clone | balance | autobalance)
	// The data distribution method used when two or more servers exist in the same forwarder group.
	Method *string `pulumi:"method"`
	// <host>:<port> of the Splunk receiver. <host> can be either an ip address or server name. <port> is the that port that the Splunk receiver is listening on.
	Name *string `pulumi:"name"`
	// The alternate name to match in the remote server's SSL certificate.
	SslAltNameToCheck *string `pulumi:"sslAltNameToCheck"`
	// Path to the client certificate. If specified, connection uses SSL.
	SslCertPath *string `pulumi:"sslCertPath"`
	// SSL Cipher in the form ALL:!aNULL:!eNULL:!LOW:!EXP:RC4+RSA:+HIGH:+MEDIUM
	SslCipher *string `pulumi:"sslCipher"`
	// Check the common name of the server's certificate against this name.
	// If there is no match, assume that Splunk Enterprise is not authenticated against this server. You must specify this setting if sslVerifyServerCert is true.
	SslCommonNameToCheck *string `pulumi:"sslCommonNameToCheck"`
	// The password associated with the CAcert.
	// The default Splunk Enterprise CAcert uses the password "password."
	SslPassword *string `pulumi:"sslPassword"`
	// The path to the root certificate authority file.
	SslRootCaPath *string `pulumi:"sslRootCaPath"`
	// If true, make sure that the server you are connecting to is a valid one (authenticated). Both the common name and the alternate name of the server are then checked for a match.
	SslVerifyServerCert *bool `pulumi:"sslVerifyServerCert"`
}

// The set of arguments for constructing a OutputsTcpServer resource.
type OutputsTcpServerArgs struct {
	// The app/user context that is the namespace for the resource
	Acl OutputsTcpServerAclPtrInput
	// If true, disables the group.
	Disabled pulumi.BoolPtrInput
	// Valid values: (clone | balance | autobalance)
	// The data distribution method used when two or more servers exist in the same forwarder group.
	Method pulumi.StringPtrInput
	// <host>:<port> of the Splunk receiver. <host> can be either an ip address or server name. <port> is the that port that the Splunk receiver is listening on.
	Name pulumi.StringPtrInput
	// The alternate name to match in the remote server's SSL certificate.
	SslAltNameToCheck pulumi.StringPtrInput
	// Path to the client certificate. If specified, connection uses SSL.
	SslCertPath pulumi.StringPtrInput
	// SSL Cipher in the form ALL:!aNULL:!eNULL:!LOW:!EXP:RC4+RSA:+HIGH:+MEDIUM
	SslCipher pulumi.StringPtrInput
	// Check the common name of the server's certificate against this name.
	// If there is no match, assume that Splunk Enterprise is not authenticated against this server. You must specify this setting if sslVerifyServerCert is true.
	SslCommonNameToCheck pulumi.StringPtrInput
	// The password associated with the CAcert.
	// The default Splunk Enterprise CAcert uses the password "password."
	SslPassword pulumi.StringPtrInput
	// The path to the root certificate authority file.
	SslRootCaPath pulumi.StringPtrInput
	// If true, make sure that the server you are connecting to is a valid one (authenticated). Both the common name and the alternate name of the server are then checked for a match.
	SslVerifyServerCert pulumi.BoolPtrInput
}

func (OutputsTcpServerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*outputsTcpServerArgs)(nil)).Elem()
}

type OutputsTcpServerInput interface {
	pulumi.Input

	ToOutputsTcpServerOutput() OutputsTcpServerOutput
	ToOutputsTcpServerOutputWithContext(ctx context.Context) OutputsTcpServerOutput
}

func (*OutputsTcpServer) ElementType() reflect.Type {
	return reflect.TypeOf((*OutputsTcpServer)(nil))
}

func (i *OutputsTcpServer) ToOutputsTcpServerOutput() OutputsTcpServerOutput {
	return i.ToOutputsTcpServerOutputWithContext(context.Background())
}

func (i *OutputsTcpServer) ToOutputsTcpServerOutputWithContext(ctx context.Context) OutputsTcpServerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OutputsTcpServerOutput)
}

func (i *OutputsTcpServer) ToOutputsTcpServerPtrOutput() OutputsTcpServerPtrOutput {
	return i.ToOutputsTcpServerPtrOutputWithContext(context.Background())
}

func (i *OutputsTcpServer) ToOutputsTcpServerPtrOutputWithContext(ctx context.Context) OutputsTcpServerPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OutputsTcpServerPtrOutput)
}

type OutputsTcpServerPtrInput interface {
	pulumi.Input

	ToOutputsTcpServerPtrOutput() OutputsTcpServerPtrOutput
	ToOutputsTcpServerPtrOutputWithContext(ctx context.Context) OutputsTcpServerPtrOutput
}

type outputsTcpServerPtrType OutputsTcpServerArgs

func (*outputsTcpServerPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**OutputsTcpServer)(nil))
}

func (i *outputsTcpServerPtrType) ToOutputsTcpServerPtrOutput() OutputsTcpServerPtrOutput {
	return i.ToOutputsTcpServerPtrOutputWithContext(context.Background())
}

func (i *outputsTcpServerPtrType) ToOutputsTcpServerPtrOutputWithContext(ctx context.Context) OutputsTcpServerPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OutputsTcpServerPtrOutput)
}

// OutputsTcpServerArrayInput is an input type that accepts OutputsTcpServerArray and OutputsTcpServerArrayOutput values.
// You can construct a concrete instance of `OutputsTcpServerArrayInput` via:
//
//          OutputsTcpServerArray{ OutputsTcpServerArgs{...} }
type OutputsTcpServerArrayInput interface {
	pulumi.Input

	ToOutputsTcpServerArrayOutput() OutputsTcpServerArrayOutput
	ToOutputsTcpServerArrayOutputWithContext(context.Context) OutputsTcpServerArrayOutput
}

type OutputsTcpServerArray []OutputsTcpServerInput

func (OutputsTcpServerArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*OutputsTcpServer)(nil)).Elem()
}

func (i OutputsTcpServerArray) ToOutputsTcpServerArrayOutput() OutputsTcpServerArrayOutput {
	return i.ToOutputsTcpServerArrayOutputWithContext(context.Background())
}

func (i OutputsTcpServerArray) ToOutputsTcpServerArrayOutputWithContext(ctx context.Context) OutputsTcpServerArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OutputsTcpServerArrayOutput)
}

// OutputsTcpServerMapInput is an input type that accepts OutputsTcpServerMap and OutputsTcpServerMapOutput values.
// You can construct a concrete instance of `OutputsTcpServerMapInput` via:
//
//          OutputsTcpServerMap{ "key": OutputsTcpServerArgs{...} }
type OutputsTcpServerMapInput interface {
	pulumi.Input

	ToOutputsTcpServerMapOutput() OutputsTcpServerMapOutput
	ToOutputsTcpServerMapOutputWithContext(context.Context) OutputsTcpServerMapOutput
}

type OutputsTcpServerMap map[string]OutputsTcpServerInput

func (OutputsTcpServerMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*OutputsTcpServer)(nil)).Elem()
}

func (i OutputsTcpServerMap) ToOutputsTcpServerMapOutput() OutputsTcpServerMapOutput {
	return i.ToOutputsTcpServerMapOutputWithContext(context.Background())
}

func (i OutputsTcpServerMap) ToOutputsTcpServerMapOutputWithContext(ctx context.Context) OutputsTcpServerMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OutputsTcpServerMapOutput)
}

type OutputsTcpServerOutput struct{ *pulumi.OutputState }

func (OutputsTcpServerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*OutputsTcpServer)(nil))
}

func (o OutputsTcpServerOutput) ToOutputsTcpServerOutput() OutputsTcpServerOutput {
	return o
}

func (o OutputsTcpServerOutput) ToOutputsTcpServerOutputWithContext(ctx context.Context) OutputsTcpServerOutput {
	return o
}

func (o OutputsTcpServerOutput) ToOutputsTcpServerPtrOutput() OutputsTcpServerPtrOutput {
	return o.ToOutputsTcpServerPtrOutputWithContext(context.Background())
}

func (o OutputsTcpServerOutput) ToOutputsTcpServerPtrOutputWithContext(ctx context.Context) OutputsTcpServerPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v OutputsTcpServer) *OutputsTcpServer {
		return &v
	}).(OutputsTcpServerPtrOutput)
}

type OutputsTcpServerPtrOutput struct{ *pulumi.OutputState }

func (OutputsTcpServerPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**OutputsTcpServer)(nil))
}

func (o OutputsTcpServerPtrOutput) ToOutputsTcpServerPtrOutput() OutputsTcpServerPtrOutput {
	return o
}

func (o OutputsTcpServerPtrOutput) ToOutputsTcpServerPtrOutputWithContext(ctx context.Context) OutputsTcpServerPtrOutput {
	return o
}

func (o OutputsTcpServerPtrOutput) Elem() OutputsTcpServerOutput {
	return o.ApplyT(func(v *OutputsTcpServer) OutputsTcpServer {
		if v != nil {
			return *v
		}
		var ret OutputsTcpServer
		return ret
	}).(OutputsTcpServerOutput)
}

type OutputsTcpServerArrayOutput struct{ *pulumi.OutputState }

func (OutputsTcpServerArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]OutputsTcpServer)(nil))
}

func (o OutputsTcpServerArrayOutput) ToOutputsTcpServerArrayOutput() OutputsTcpServerArrayOutput {
	return o
}

func (o OutputsTcpServerArrayOutput) ToOutputsTcpServerArrayOutputWithContext(ctx context.Context) OutputsTcpServerArrayOutput {
	return o
}

func (o OutputsTcpServerArrayOutput) Index(i pulumi.IntInput) OutputsTcpServerOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) OutputsTcpServer {
		return vs[0].([]OutputsTcpServer)[vs[1].(int)]
	}).(OutputsTcpServerOutput)
}

type OutputsTcpServerMapOutput struct{ *pulumi.OutputState }

func (OutputsTcpServerMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]OutputsTcpServer)(nil))
}

func (o OutputsTcpServerMapOutput) ToOutputsTcpServerMapOutput() OutputsTcpServerMapOutput {
	return o
}

func (o OutputsTcpServerMapOutput) ToOutputsTcpServerMapOutputWithContext(ctx context.Context) OutputsTcpServerMapOutput {
	return o
}

func (o OutputsTcpServerMapOutput) MapIndex(k pulumi.StringInput) OutputsTcpServerOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) OutputsTcpServer {
		return vs[0].(map[string]OutputsTcpServer)[vs[1].(string)]
	}).(OutputsTcpServerOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*OutputsTcpServerInput)(nil)).Elem(), &OutputsTcpServer{})
	pulumi.RegisterInputType(reflect.TypeOf((*OutputsTcpServerPtrInput)(nil)).Elem(), &OutputsTcpServer{})
	pulumi.RegisterInputType(reflect.TypeOf((*OutputsTcpServerArrayInput)(nil)).Elem(), OutputsTcpServerArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*OutputsTcpServerMapInput)(nil)).Elem(), OutputsTcpServerMap{})
	pulumi.RegisterOutputType(OutputsTcpServerOutput{})
	pulumi.RegisterOutputType(OutputsTcpServerPtrOutput{})
	pulumi.RegisterOutputType(OutputsTcpServerArrayOutput{})
	pulumi.RegisterOutputType(OutputsTcpServerMapOutput{})
}
