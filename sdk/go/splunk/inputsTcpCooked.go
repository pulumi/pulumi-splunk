// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package splunk

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-splunk/sdk/go/splunk/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// ## # Resource: InputsTcpCooked
//
// Create or update cooked TCP input information and create new containers for managing cooked data.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-splunk/sdk/go/splunk"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := splunk.NewInputsTcpCooked(ctx, "tcpCooked", &splunk.InputsTcpCookedArgs{
//				ConnectionHost: pulumi.String("dns"),
//				Disabled:       pulumi.Bool(false),
//				RestrictToHost: pulumi.String("splunk"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
type InputsTcpCooked struct {
	pulumi.CustomResourceState

	// The app/user context that is the namespace for the resource
	Acl InputsTcpCookedAclOutput `pulumi:"acl"`
	// Valid values: (ip | dns | none)
	// Set the host for the remote server that is sending data.
	// ip sets the host to the IP address of the remote server sending data.
	// dns sets the host to the reverse DNS entry for the IP address of the remote server sending data.
	// none leaves the host as specified in inputs.conf, which is typically the Splunk system hostname.
	// Default value is dns.
	ConnectionHost pulumi.StringOutput `pulumi:"connectionHost"`
	// Indicates if input is disabled.
	Disabled pulumi.BoolOutput `pulumi:"disabled"`
	// Host from which the indexer gets data.
	Host pulumi.StringOutput `pulumi:"host"`
	// The port number of this input.
	Name pulumi.StringOutput `pulumi:"name"`
	// Restrict incoming connections on this port to the host specified here.
	RestrictToHost pulumi.StringOutput `pulumi:"restrictToHost"`
}

// NewInputsTcpCooked registers a new resource with the given unique name, arguments, and options.
func NewInputsTcpCooked(ctx *pulumi.Context,
	name string, args *InputsTcpCookedArgs, opts ...pulumi.ResourceOption) (*InputsTcpCooked, error) {
	if args == nil {
		args = &InputsTcpCookedArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource InputsTcpCooked
	err := ctx.RegisterResource("splunk:index/inputsTcpCooked:InputsTcpCooked", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetInputsTcpCooked gets an existing InputsTcpCooked resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetInputsTcpCooked(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *InputsTcpCookedState, opts ...pulumi.ResourceOption) (*InputsTcpCooked, error) {
	var resource InputsTcpCooked
	err := ctx.ReadResource("splunk:index/inputsTcpCooked:InputsTcpCooked", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering InputsTcpCooked resources.
type inputsTcpCookedState struct {
	// The app/user context that is the namespace for the resource
	Acl *InputsTcpCookedAcl `pulumi:"acl"`
	// Valid values: (ip | dns | none)
	// Set the host for the remote server that is sending data.
	// ip sets the host to the IP address of the remote server sending data.
	// dns sets the host to the reverse DNS entry for the IP address of the remote server sending data.
	// none leaves the host as specified in inputs.conf, which is typically the Splunk system hostname.
	// Default value is dns.
	ConnectionHost *string `pulumi:"connectionHost"`
	// Indicates if input is disabled.
	Disabled *bool `pulumi:"disabled"`
	// Host from which the indexer gets data.
	Host *string `pulumi:"host"`
	// The port number of this input.
	Name *string `pulumi:"name"`
	// Restrict incoming connections on this port to the host specified here.
	RestrictToHost *string `pulumi:"restrictToHost"`
}

type InputsTcpCookedState struct {
	// The app/user context that is the namespace for the resource
	Acl InputsTcpCookedAclPtrInput
	// Valid values: (ip | dns | none)
	// Set the host for the remote server that is sending data.
	// ip sets the host to the IP address of the remote server sending data.
	// dns sets the host to the reverse DNS entry for the IP address of the remote server sending data.
	// none leaves the host as specified in inputs.conf, which is typically the Splunk system hostname.
	// Default value is dns.
	ConnectionHost pulumi.StringPtrInput
	// Indicates if input is disabled.
	Disabled pulumi.BoolPtrInput
	// Host from which the indexer gets data.
	Host pulumi.StringPtrInput
	// The port number of this input.
	Name pulumi.StringPtrInput
	// Restrict incoming connections on this port to the host specified here.
	RestrictToHost pulumi.StringPtrInput
}

func (InputsTcpCookedState) ElementType() reflect.Type {
	return reflect.TypeOf((*inputsTcpCookedState)(nil)).Elem()
}

type inputsTcpCookedArgs struct {
	// The app/user context that is the namespace for the resource
	Acl *InputsTcpCookedAcl `pulumi:"acl"`
	// Valid values: (ip | dns | none)
	// Set the host for the remote server that is sending data.
	// ip sets the host to the IP address of the remote server sending data.
	// dns sets the host to the reverse DNS entry for the IP address of the remote server sending data.
	// none leaves the host as specified in inputs.conf, which is typically the Splunk system hostname.
	// Default value is dns.
	ConnectionHost *string `pulumi:"connectionHost"`
	// Indicates if input is disabled.
	Disabled *bool `pulumi:"disabled"`
	// Host from which the indexer gets data.
	Host *string `pulumi:"host"`
	// The port number of this input.
	Name *string `pulumi:"name"`
	// Restrict incoming connections on this port to the host specified here.
	RestrictToHost *string `pulumi:"restrictToHost"`
}

// The set of arguments for constructing a InputsTcpCooked resource.
type InputsTcpCookedArgs struct {
	// The app/user context that is the namespace for the resource
	Acl InputsTcpCookedAclPtrInput
	// Valid values: (ip | dns | none)
	// Set the host for the remote server that is sending data.
	// ip sets the host to the IP address of the remote server sending data.
	// dns sets the host to the reverse DNS entry for the IP address of the remote server sending data.
	// none leaves the host as specified in inputs.conf, which is typically the Splunk system hostname.
	// Default value is dns.
	ConnectionHost pulumi.StringPtrInput
	// Indicates if input is disabled.
	Disabled pulumi.BoolPtrInput
	// Host from which the indexer gets data.
	Host pulumi.StringPtrInput
	// The port number of this input.
	Name pulumi.StringPtrInput
	// Restrict incoming connections on this port to the host specified here.
	RestrictToHost pulumi.StringPtrInput
}

func (InputsTcpCookedArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*inputsTcpCookedArgs)(nil)).Elem()
}

type InputsTcpCookedInput interface {
	pulumi.Input

	ToInputsTcpCookedOutput() InputsTcpCookedOutput
	ToInputsTcpCookedOutputWithContext(ctx context.Context) InputsTcpCookedOutput
}

func (*InputsTcpCooked) ElementType() reflect.Type {
	return reflect.TypeOf((**InputsTcpCooked)(nil)).Elem()
}

func (i *InputsTcpCooked) ToInputsTcpCookedOutput() InputsTcpCookedOutput {
	return i.ToInputsTcpCookedOutputWithContext(context.Background())
}

func (i *InputsTcpCooked) ToInputsTcpCookedOutputWithContext(ctx context.Context) InputsTcpCookedOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InputsTcpCookedOutput)
}

func (i *InputsTcpCooked) ToOutput(ctx context.Context) pulumix.Output[*InputsTcpCooked] {
	return pulumix.Output[*InputsTcpCooked]{
		OutputState: i.ToInputsTcpCookedOutputWithContext(ctx).OutputState,
	}
}

// InputsTcpCookedArrayInput is an input type that accepts InputsTcpCookedArray and InputsTcpCookedArrayOutput values.
// You can construct a concrete instance of `InputsTcpCookedArrayInput` via:
//
//	InputsTcpCookedArray{ InputsTcpCookedArgs{...} }
type InputsTcpCookedArrayInput interface {
	pulumi.Input

	ToInputsTcpCookedArrayOutput() InputsTcpCookedArrayOutput
	ToInputsTcpCookedArrayOutputWithContext(context.Context) InputsTcpCookedArrayOutput
}

type InputsTcpCookedArray []InputsTcpCookedInput

func (InputsTcpCookedArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*InputsTcpCooked)(nil)).Elem()
}

func (i InputsTcpCookedArray) ToInputsTcpCookedArrayOutput() InputsTcpCookedArrayOutput {
	return i.ToInputsTcpCookedArrayOutputWithContext(context.Background())
}

func (i InputsTcpCookedArray) ToInputsTcpCookedArrayOutputWithContext(ctx context.Context) InputsTcpCookedArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InputsTcpCookedArrayOutput)
}

func (i InputsTcpCookedArray) ToOutput(ctx context.Context) pulumix.Output[[]*InputsTcpCooked] {
	return pulumix.Output[[]*InputsTcpCooked]{
		OutputState: i.ToInputsTcpCookedArrayOutputWithContext(ctx).OutputState,
	}
}

// InputsTcpCookedMapInput is an input type that accepts InputsTcpCookedMap and InputsTcpCookedMapOutput values.
// You can construct a concrete instance of `InputsTcpCookedMapInput` via:
//
//	InputsTcpCookedMap{ "key": InputsTcpCookedArgs{...} }
type InputsTcpCookedMapInput interface {
	pulumi.Input

	ToInputsTcpCookedMapOutput() InputsTcpCookedMapOutput
	ToInputsTcpCookedMapOutputWithContext(context.Context) InputsTcpCookedMapOutput
}

type InputsTcpCookedMap map[string]InputsTcpCookedInput

func (InputsTcpCookedMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*InputsTcpCooked)(nil)).Elem()
}

func (i InputsTcpCookedMap) ToInputsTcpCookedMapOutput() InputsTcpCookedMapOutput {
	return i.ToInputsTcpCookedMapOutputWithContext(context.Background())
}

func (i InputsTcpCookedMap) ToInputsTcpCookedMapOutputWithContext(ctx context.Context) InputsTcpCookedMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InputsTcpCookedMapOutput)
}

func (i InputsTcpCookedMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*InputsTcpCooked] {
	return pulumix.Output[map[string]*InputsTcpCooked]{
		OutputState: i.ToInputsTcpCookedMapOutputWithContext(ctx).OutputState,
	}
}

type InputsTcpCookedOutput struct{ *pulumi.OutputState }

func (InputsTcpCookedOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**InputsTcpCooked)(nil)).Elem()
}

func (o InputsTcpCookedOutput) ToInputsTcpCookedOutput() InputsTcpCookedOutput {
	return o
}

func (o InputsTcpCookedOutput) ToInputsTcpCookedOutputWithContext(ctx context.Context) InputsTcpCookedOutput {
	return o
}

func (o InputsTcpCookedOutput) ToOutput(ctx context.Context) pulumix.Output[*InputsTcpCooked] {
	return pulumix.Output[*InputsTcpCooked]{
		OutputState: o.OutputState,
	}
}

// The app/user context that is the namespace for the resource
func (o InputsTcpCookedOutput) Acl() InputsTcpCookedAclOutput {
	return o.ApplyT(func(v *InputsTcpCooked) InputsTcpCookedAclOutput { return v.Acl }).(InputsTcpCookedAclOutput)
}

// Valid values: (ip | dns | none)
// Set the host for the remote server that is sending data.
// ip sets the host to the IP address of the remote server sending data.
// dns sets the host to the reverse DNS entry for the IP address of the remote server sending data.
// none leaves the host as specified in inputs.conf, which is typically the Splunk system hostname.
// Default value is dns.
func (o InputsTcpCookedOutput) ConnectionHost() pulumi.StringOutput {
	return o.ApplyT(func(v *InputsTcpCooked) pulumi.StringOutput { return v.ConnectionHost }).(pulumi.StringOutput)
}

// Indicates if input is disabled.
func (o InputsTcpCookedOutput) Disabled() pulumi.BoolOutput {
	return o.ApplyT(func(v *InputsTcpCooked) pulumi.BoolOutput { return v.Disabled }).(pulumi.BoolOutput)
}

// Host from which the indexer gets data.
func (o InputsTcpCookedOutput) Host() pulumi.StringOutput {
	return o.ApplyT(func(v *InputsTcpCooked) pulumi.StringOutput { return v.Host }).(pulumi.StringOutput)
}

// The port number of this input.
func (o InputsTcpCookedOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *InputsTcpCooked) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Restrict incoming connections on this port to the host specified here.
func (o InputsTcpCookedOutput) RestrictToHost() pulumi.StringOutput {
	return o.ApplyT(func(v *InputsTcpCooked) pulumi.StringOutput { return v.RestrictToHost }).(pulumi.StringOutput)
}

type InputsTcpCookedArrayOutput struct{ *pulumi.OutputState }

func (InputsTcpCookedArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*InputsTcpCooked)(nil)).Elem()
}

func (o InputsTcpCookedArrayOutput) ToInputsTcpCookedArrayOutput() InputsTcpCookedArrayOutput {
	return o
}

func (o InputsTcpCookedArrayOutput) ToInputsTcpCookedArrayOutputWithContext(ctx context.Context) InputsTcpCookedArrayOutput {
	return o
}

func (o InputsTcpCookedArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*InputsTcpCooked] {
	return pulumix.Output[[]*InputsTcpCooked]{
		OutputState: o.OutputState,
	}
}

func (o InputsTcpCookedArrayOutput) Index(i pulumi.IntInput) InputsTcpCookedOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *InputsTcpCooked {
		return vs[0].([]*InputsTcpCooked)[vs[1].(int)]
	}).(InputsTcpCookedOutput)
}

type InputsTcpCookedMapOutput struct{ *pulumi.OutputState }

func (InputsTcpCookedMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*InputsTcpCooked)(nil)).Elem()
}

func (o InputsTcpCookedMapOutput) ToInputsTcpCookedMapOutput() InputsTcpCookedMapOutput {
	return o
}

func (o InputsTcpCookedMapOutput) ToInputsTcpCookedMapOutputWithContext(ctx context.Context) InputsTcpCookedMapOutput {
	return o
}

func (o InputsTcpCookedMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*InputsTcpCooked] {
	return pulumix.Output[map[string]*InputsTcpCooked]{
		OutputState: o.OutputState,
	}
}

func (o InputsTcpCookedMapOutput) MapIndex(k pulumi.StringInput) InputsTcpCookedOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *InputsTcpCooked {
		return vs[0].(map[string]*InputsTcpCooked)[vs[1].(string)]
	}).(InputsTcpCookedOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*InputsTcpCookedInput)(nil)).Elem(), &InputsTcpCooked{})
	pulumi.RegisterInputType(reflect.TypeOf((*InputsTcpCookedArrayInput)(nil)).Elem(), InputsTcpCookedArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*InputsTcpCookedMapInput)(nil)).Elem(), InputsTcpCookedMap{})
	pulumi.RegisterOutputType(InputsTcpCookedOutput{})
	pulumi.RegisterOutputType(InputsTcpCookedArrayOutput{})
	pulumi.RegisterOutputType(InputsTcpCookedMapOutput{})
}
