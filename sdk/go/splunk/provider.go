// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package splunk

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// The provider type for the splunk package. By default, resources use package-wide configuration
// settings, however an explicit `Provider` instance may be created and passed during resource
// construction to achieve fine-grained programmatic control over provider settings. See the
// [documentation](https://www.pulumi.com/docs/reference/programming-model/#providers) for more information.
type Provider struct {
	pulumi.ProviderResourceState
}

// NewProvider registers a new resource with the given unique name, arguments, and options.
func NewProvider(ctx *pulumi.Context,
	name string, args *ProviderArgs, opts ...pulumi.ResourceOption) (*Provider, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Url == nil {
		return nil, errors.New("invalid value for required argument 'Url'")
	}
	var resource Provider
	err := ctx.RegisterResource("pulumi:providers:splunk", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

type providerArgs struct {
	// Authentication tokens, also known as JSON Web Tokens (JWT), are a method for authenticating Splunk platform users into
	// the Splunk platform
	AuthToken *string `pulumi:"authToken"`
	// insecure skip verification flag
	InsecureSkipVerify *bool `pulumi:"insecureSkipVerify"`
	// Splunk instance password
	Password *string `pulumi:"password"`
	// Timeout when making calls to Splunk server. Defaults to 60 seconds
	Timeout *int `pulumi:"timeout"`
	// Splunk instance URL
	Url string `pulumi:"url"`
	// Splunk instance admin username
	Username *string `pulumi:"username"`
}

// The set of arguments for constructing a Provider resource.
type ProviderArgs struct {
	// Authentication tokens, also known as JSON Web Tokens (JWT), are a method for authenticating Splunk platform users into
	// the Splunk platform
	AuthToken pulumi.StringPtrInput
	// insecure skip verification flag
	InsecureSkipVerify pulumi.BoolPtrInput
	// Splunk instance password
	Password pulumi.StringPtrInput
	// Timeout when making calls to Splunk server. Defaults to 60 seconds
	Timeout pulumi.IntPtrInput
	// Splunk instance URL
	Url pulumi.StringInput
	// Splunk instance admin username
	Username pulumi.StringPtrInput
}

func (ProviderArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*providerArgs)(nil)).Elem()
}

type ProviderInput interface {
	pulumi.Input

	ToProviderOutput() ProviderOutput
	ToProviderOutputWithContext(ctx context.Context) ProviderOutput
}

func (Provider) ElementType() reflect.Type {
	return reflect.TypeOf((*Provider)(nil)).Elem()
}

func (i Provider) ToProviderOutput() ProviderOutput {
	return i.ToProviderOutputWithContext(context.Background())
}

func (i Provider) ToProviderOutputWithContext(ctx context.Context) ProviderOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProviderOutput)
}

type ProviderOutput struct {
	*pulumi.OutputState
}

func (ProviderOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ProviderOutput)(nil)).Elem()
}

func (o ProviderOutput) ToProviderOutput() ProviderOutput {
	return o
}

func (o ProviderOutput) ToProviderOutputWithContext(ctx context.Context) ProviderOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(ProviderOutput{})
}
