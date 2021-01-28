// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package splunk

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
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
// 	"github.com/pulumi/pulumi-splunk/sdk/go/splunk/"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := splunk.NewInputsTcpSsl(ctx, "test", &splunk.InputsTcpSslArgs{
// 			Disabled:          pulumi.Bool(false),
// 			RequireClientCert: pulumi.Bool(true),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
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
	return reflect.TypeOf((*InputsTcpSsl)(nil))
}

func (i *InputsTcpSsl) ToInputsTcpSslOutput() InputsTcpSslOutput {
	return i.ToInputsTcpSslOutputWithContext(context.Background())
}

func (i *InputsTcpSsl) ToInputsTcpSslOutputWithContext(ctx context.Context) InputsTcpSslOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InputsTcpSslOutput)
}

type InputsTcpSslOutput struct {
	*pulumi.OutputState
}

func (InputsTcpSslOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*InputsTcpSsl)(nil))
}

func (o InputsTcpSslOutput) ToInputsTcpSslOutput() InputsTcpSslOutput {
	return o
}

func (o InputsTcpSslOutput) ToInputsTcpSslOutputWithContext(ctx context.Context) InputsTcpSslOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(InputsTcpSslOutput{})
}
