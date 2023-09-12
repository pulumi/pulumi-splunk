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

// ## # Resource: InputsTcpSsl
//
// Access or update the SSL configuration for the host.
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
//			_, err := splunk.NewInputsTcpSsl(ctx, "test", &splunk.InputsTcpSslArgs{
//				Disabled:          pulumi.Bool(false),
//				RequireClientCert: pulumi.Bool(true),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
type InputsTcpSsl struct {
	pulumi.CustomResourceState

	// Indicates if input is disabled.
	Disabled pulumi.BoolOutput `pulumi:"disabled"`
	// Server certificate password, if any.
	Password pulumi.StringOutput `pulumi:"password"`
	// Determines whether a client must authenticate.
	RequireClientCert pulumi.BoolOutput `pulumi:"requireClientCert"`
	// Certificate authority list (root file)
	RootCa pulumi.StringOutput `pulumi:"rootCa"`
	// Full path to the server certificate.
	ServerCert pulumi.StringOutput `pulumi:"serverCert"`
}

// NewInputsTcpSsl registers a new resource with the given unique name, arguments, and options.
func NewInputsTcpSsl(ctx *pulumi.Context,
	name string, args *InputsTcpSslArgs, opts ...pulumi.ResourceOption) (*InputsTcpSsl, error) {
	if args == nil {
		args = &InputsTcpSslArgs{}
	}

	if args.Password != nil {
		args.Password = pulumi.ToSecret(args.Password).(pulumi.StringPtrInput)
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"password",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource InputsTcpSsl
	err := ctx.RegisterResource("splunk:index/inputsTcpSsl:InputsTcpSsl", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetInputsTcpSsl gets an existing InputsTcpSsl resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetInputsTcpSsl(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *InputsTcpSslState, opts ...pulumi.ResourceOption) (*InputsTcpSsl, error) {
	var resource InputsTcpSsl
	err := ctx.ReadResource("splunk:index/inputsTcpSsl:InputsTcpSsl", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering InputsTcpSsl resources.
type inputsTcpSslState struct {
	// Indicates if input is disabled.
	Disabled *bool `pulumi:"disabled"`
	// Server certificate password, if any.
	Password *string `pulumi:"password"`
	// Determines whether a client must authenticate.
	RequireClientCert *bool `pulumi:"requireClientCert"`
	// Certificate authority list (root file)
	RootCa *string `pulumi:"rootCa"`
	// Full path to the server certificate.
	ServerCert *string `pulumi:"serverCert"`
}

type InputsTcpSslState struct {
	// Indicates if input is disabled.
	Disabled pulumi.BoolPtrInput
	// Server certificate password, if any.
	Password pulumi.StringPtrInput
	// Determines whether a client must authenticate.
	RequireClientCert pulumi.BoolPtrInput
	// Certificate authority list (root file)
	RootCa pulumi.StringPtrInput
	// Full path to the server certificate.
	ServerCert pulumi.StringPtrInput
}

func (InputsTcpSslState) ElementType() reflect.Type {
	return reflect.TypeOf((*inputsTcpSslState)(nil)).Elem()
}

type inputsTcpSslArgs struct {
	// Indicates if input is disabled.
	Disabled *bool `pulumi:"disabled"`
	// Server certificate password, if any.
	Password *string `pulumi:"password"`
	// Determines whether a client must authenticate.
	RequireClientCert *bool `pulumi:"requireClientCert"`
	// Certificate authority list (root file)
	RootCa *string `pulumi:"rootCa"`
	// Full path to the server certificate.
	ServerCert *string `pulumi:"serverCert"`
}

// The set of arguments for constructing a InputsTcpSsl resource.
type InputsTcpSslArgs struct {
	// Indicates if input is disabled.
	Disabled pulumi.BoolPtrInput
	// Server certificate password, if any.
	Password pulumi.StringPtrInput
	// Determines whether a client must authenticate.
	RequireClientCert pulumi.BoolPtrInput
	// Certificate authority list (root file)
	RootCa pulumi.StringPtrInput
	// Full path to the server certificate.
	ServerCert pulumi.StringPtrInput
}

func (InputsTcpSslArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*inputsTcpSslArgs)(nil)).Elem()
}

type InputsTcpSslInput interface {
	pulumi.Input

	ToInputsTcpSslOutput() InputsTcpSslOutput
	ToInputsTcpSslOutputWithContext(ctx context.Context) InputsTcpSslOutput
}

func (*InputsTcpSsl) ElementType() reflect.Type {
	return reflect.TypeOf((**InputsTcpSsl)(nil)).Elem()
}

func (i *InputsTcpSsl) ToInputsTcpSslOutput() InputsTcpSslOutput {
	return i.ToInputsTcpSslOutputWithContext(context.Background())
}

func (i *InputsTcpSsl) ToInputsTcpSslOutputWithContext(ctx context.Context) InputsTcpSslOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InputsTcpSslOutput)
}

func (i *InputsTcpSsl) ToOutput(ctx context.Context) pulumix.Output[*InputsTcpSsl] {
	return pulumix.Output[*InputsTcpSsl]{
		OutputState: i.ToInputsTcpSslOutputWithContext(ctx).OutputState,
	}
}

// InputsTcpSslArrayInput is an input type that accepts InputsTcpSslArray and InputsTcpSslArrayOutput values.
// You can construct a concrete instance of `InputsTcpSslArrayInput` via:
//
//	InputsTcpSslArray{ InputsTcpSslArgs{...} }
type InputsTcpSslArrayInput interface {
	pulumi.Input

	ToInputsTcpSslArrayOutput() InputsTcpSslArrayOutput
	ToInputsTcpSslArrayOutputWithContext(context.Context) InputsTcpSslArrayOutput
}

type InputsTcpSslArray []InputsTcpSslInput

func (InputsTcpSslArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*InputsTcpSsl)(nil)).Elem()
}

func (i InputsTcpSslArray) ToInputsTcpSslArrayOutput() InputsTcpSslArrayOutput {
	return i.ToInputsTcpSslArrayOutputWithContext(context.Background())
}

func (i InputsTcpSslArray) ToInputsTcpSslArrayOutputWithContext(ctx context.Context) InputsTcpSslArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InputsTcpSslArrayOutput)
}

func (i InputsTcpSslArray) ToOutput(ctx context.Context) pulumix.Output[[]*InputsTcpSsl] {
	return pulumix.Output[[]*InputsTcpSsl]{
		OutputState: i.ToInputsTcpSslArrayOutputWithContext(ctx).OutputState,
	}
}

// InputsTcpSslMapInput is an input type that accepts InputsTcpSslMap and InputsTcpSslMapOutput values.
// You can construct a concrete instance of `InputsTcpSslMapInput` via:
//
//	InputsTcpSslMap{ "key": InputsTcpSslArgs{...} }
type InputsTcpSslMapInput interface {
	pulumi.Input

	ToInputsTcpSslMapOutput() InputsTcpSslMapOutput
	ToInputsTcpSslMapOutputWithContext(context.Context) InputsTcpSslMapOutput
}

type InputsTcpSslMap map[string]InputsTcpSslInput

func (InputsTcpSslMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*InputsTcpSsl)(nil)).Elem()
}

func (i InputsTcpSslMap) ToInputsTcpSslMapOutput() InputsTcpSslMapOutput {
	return i.ToInputsTcpSslMapOutputWithContext(context.Background())
}

func (i InputsTcpSslMap) ToInputsTcpSslMapOutputWithContext(ctx context.Context) InputsTcpSslMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InputsTcpSslMapOutput)
}

func (i InputsTcpSslMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*InputsTcpSsl] {
	return pulumix.Output[map[string]*InputsTcpSsl]{
		OutputState: i.ToInputsTcpSslMapOutputWithContext(ctx).OutputState,
	}
}

type InputsTcpSslOutput struct{ *pulumi.OutputState }

func (InputsTcpSslOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**InputsTcpSsl)(nil)).Elem()
}

func (o InputsTcpSslOutput) ToInputsTcpSslOutput() InputsTcpSslOutput {
	return o
}

func (o InputsTcpSslOutput) ToInputsTcpSslOutputWithContext(ctx context.Context) InputsTcpSslOutput {
	return o
}

func (o InputsTcpSslOutput) ToOutput(ctx context.Context) pulumix.Output[*InputsTcpSsl] {
	return pulumix.Output[*InputsTcpSsl]{
		OutputState: o.OutputState,
	}
}

// Indicates if input is disabled.
func (o InputsTcpSslOutput) Disabled() pulumi.BoolOutput {
	return o.ApplyT(func(v *InputsTcpSsl) pulumi.BoolOutput { return v.Disabled }).(pulumi.BoolOutput)
}

// Server certificate password, if any.
func (o InputsTcpSslOutput) Password() pulumi.StringOutput {
	return o.ApplyT(func(v *InputsTcpSsl) pulumi.StringOutput { return v.Password }).(pulumi.StringOutput)
}

// Determines whether a client must authenticate.
func (o InputsTcpSslOutput) RequireClientCert() pulumi.BoolOutput {
	return o.ApplyT(func(v *InputsTcpSsl) pulumi.BoolOutput { return v.RequireClientCert }).(pulumi.BoolOutput)
}

// Certificate authority list (root file)
func (o InputsTcpSslOutput) RootCa() pulumi.StringOutput {
	return o.ApplyT(func(v *InputsTcpSsl) pulumi.StringOutput { return v.RootCa }).(pulumi.StringOutput)
}

// Full path to the server certificate.
func (o InputsTcpSslOutput) ServerCert() pulumi.StringOutput {
	return o.ApplyT(func(v *InputsTcpSsl) pulumi.StringOutput { return v.ServerCert }).(pulumi.StringOutput)
}

type InputsTcpSslArrayOutput struct{ *pulumi.OutputState }

func (InputsTcpSslArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*InputsTcpSsl)(nil)).Elem()
}

func (o InputsTcpSslArrayOutput) ToInputsTcpSslArrayOutput() InputsTcpSslArrayOutput {
	return o
}

func (o InputsTcpSslArrayOutput) ToInputsTcpSslArrayOutputWithContext(ctx context.Context) InputsTcpSslArrayOutput {
	return o
}

func (o InputsTcpSslArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*InputsTcpSsl] {
	return pulumix.Output[[]*InputsTcpSsl]{
		OutputState: o.OutputState,
	}
}

func (o InputsTcpSslArrayOutput) Index(i pulumi.IntInput) InputsTcpSslOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *InputsTcpSsl {
		return vs[0].([]*InputsTcpSsl)[vs[1].(int)]
	}).(InputsTcpSslOutput)
}

type InputsTcpSslMapOutput struct{ *pulumi.OutputState }

func (InputsTcpSslMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*InputsTcpSsl)(nil)).Elem()
}

func (o InputsTcpSslMapOutput) ToInputsTcpSslMapOutput() InputsTcpSslMapOutput {
	return o
}

func (o InputsTcpSslMapOutput) ToInputsTcpSslMapOutputWithContext(ctx context.Context) InputsTcpSslMapOutput {
	return o
}

func (o InputsTcpSslMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*InputsTcpSsl] {
	return pulumix.Output[map[string]*InputsTcpSsl]{
		OutputState: o.OutputState,
	}
}

func (o InputsTcpSslMapOutput) MapIndex(k pulumi.StringInput) InputsTcpSslOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *InputsTcpSsl {
		return vs[0].(map[string]*InputsTcpSsl)[vs[1].(string)]
	}).(InputsTcpSslOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*InputsTcpSslInput)(nil)).Elem(), &InputsTcpSsl{})
	pulumi.RegisterInputType(reflect.TypeOf((*InputsTcpSslArrayInput)(nil)).Elem(), InputsTcpSslArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*InputsTcpSslMapInput)(nil)).Elem(), InputsTcpSslMap{})
	pulumi.RegisterOutputType(InputsTcpSslOutput{})
	pulumi.RegisterOutputType(InputsTcpSslArrayOutput{})
	pulumi.RegisterOutputType(InputsTcpSslMapOutput{})
}
