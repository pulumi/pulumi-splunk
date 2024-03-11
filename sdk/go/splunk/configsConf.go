// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package splunk

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-splunk/sdk/go/splunk/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## # Resource: ConfigsConf
//
// Create and manage configuration file stanzas.
//
// ## Example Usage
//
// <!--Start PulumiCodeChooser -->
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
//			_, err := splunk.NewConfigsConf(ctx, "new-conf-stanza", &splunk.ConfigsConfArgs{
//				Variables: pulumi.StringMap{
//					"disabled":   pulumi.String("false"),
//					"custom_key": pulumi.String("value"),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// <!--End PulumiCodeChooser -->
type ConfigsConf struct {
	pulumi.CustomResourceState

	Acl ConfigsConfAclOutput `pulumi:"acl"`
	// A '/' separated string consisting of {conf_file_name}/{stanza_name} ex. props/custom_stanza
	Name pulumi.StringOutput `pulumi:"name"`
	// A map of key value pairs for a stanza.
	Variables pulumi.StringMapOutput `pulumi:"variables"`
}

// NewConfigsConf registers a new resource with the given unique name, arguments, and options.
func NewConfigsConf(ctx *pulumi.Context,
	name string, args *ConfigsConfArgs, opts ...pulumi.ResourceOption) (*ConfigsConf, error) {
	if args == nil {
		args = &ConfigsConfArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ConfigsConf
	err := ctx.RegisterResource("splunk:index/configsConf:ConfigsConf", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetConfigsConf gets an existing ConfigsConf resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetConfigsConf(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ConfigsConfState, opts ...pulumi.ResourceOption) (*ConfigsConf, error) {
	var resource ConfigsConf
	err := ctx.ReadResource("splunk:index/configsConf:ConfigsConf", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ConfigsConf resources.
type configsConfState struct {
	Acl *ConfigsConfAcl `pulumi:"acl"`
	// A '/' separated string consisting of {conf_file_name}/{stanza_name} ex. props/custom_stanza
	Name *string `pulumi:"name"`
	// A map of key value pairs for a stanza.
	Variables map[string]string `pulumi:"variables"`
}

type ConfigsConfState struct {
	Acl ConfigsConfAclPtrInput
	// A '/' separated string consisting of {conf_file_name}/{stanza_name} ex. props/custom_stanza
	Name pulumi.StringPtrInput
	// A map of key value pairs for a stanza.
	Variables pulumi.StringMapInput
}

func (ConfigsConfState) ElementType() reflect.Type {
	return reflect.TypeOf((*configsConfState)(nil)).Elem()
}

type configsConfArgs struct {
	Acl *ConfigsConfAcl `pulumi:"acl"`
	// A '/' separated string consisting of {conf_file_name}/{stanza_name} ex. props/custom_stanza
	Name *string `pulumi:"name"`
	// A map of key value pairs for a stanza.
	Variables map[string]string `pulumi:"variables"`
}

// The set of arguments for constructing a ConfigsConf resource.
type ConfigsConfArgs struct {
	Acl ConfigsConfAclPtrInput
	// A '/' separated string consisting of {conf_file_name}/{stanza_name} ex. props/custom_stanza
	Name pulumi.StringPtrInput
	// A map of key value pairs for a stanza.
	Variables pulumi.StringMapInput
}

func (ConfigsConfArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*configsConfArgs)(nil)).Elem()
}

type ConfigsConfInput interface {
	pulumi.Input

	ToConfigsConfOutput() ConfigsConfOutput
	ToConfigsConfOutputWithContext(ctx context.Context) ConfigsConfOutput
}

func (*ConfigsConf) ElementType() reflect.Type {
	return reflect.TypeOf((**ConfigsConf)(nil)).Elem()
}

func (i *ConfigsConf) ToConfigsConfOutput() ConfigsConfOutput {
	return i.ToConfigsConfOutputWithContext(context.Background())
}

func (i *ConfigsConf) ToConfigsConfOutputWithContext(ctx context.Context) ConfigsConfOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConfigsConfOutput)
}

// ConfigsConfArrayInput is an input type that accepts ConfigsConfArray and ConfigsConfArrayOutput values.
// You can construct a concrete instance of `ConfigsConfArrayInput` via:
//
//	ConfigsConfArray{ ConfigsConfArgs{...} }
type ConfigsConfArrayInput interface {
	pulumi.Input

	ToConfigsConfArrayOutput() ConfigsConfArrayOutput
	ToConfigsConfArrayOutputWithContext(context.Context) ConfigsConfArrayOutput
}

type ConfigsConfArray []ConfigsConfInput

func (ConfigsConfArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ConfigsConf)(nil)).Elem()
}

func (i ConfigsConfArray) ToConfigsConfArrayOutput() ConfigsConfArrayOutput {
	return i.ToConfigsConfArrayOutputWithContext(context.Background())
}

func (i ConfigsConfArray) ToConfigsConfArrayOutputWithContext(ctx context.Context) ConfigsConfArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConfigsConfArrayOutput)
}

// ConfigsConfMapInput is an input type that accepts ConfigsConfMap and ConfigsConfMapOutput values.
// You can construct a concrete instance of `ConfigsConfMapInput` via:
//
//	ConfigsConfMap{ "key": ConfigsConfArgs{...} }
type ConfigsConfMapInput interface {
	pulumi.Input

	ToConfigsConfMapOutput() ConfigsConfMapOutput
	ToConfigsConfMapOutputWithContext(context.Context) ConfigsConfMapOutput
}

type ConfigsConfMap map[string]ConfigsConfInput

func (ConfigsConfMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ConfigsConf)(nil)).Elem()
}

func (i ConfigsConfMap) ToConfigsConfMapOutput() ConfigsConfMapOutput {
	return i.ToConfigsConfMapOutputWithContext(context.Background())
}

func (i ConfigsConfMap) ToConfigsConfMapOutputWithContext(ctx context.Context) ConfigsConfMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConfigsConfMapOutput)
}

type ConfigsConfOutput struct{ *pulumi.OutputState }

func (ConfigsConfOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ConfigsConf)(nil)).Elem()
}

func (o ConfigsConfOutput) ToConfigsConfOutput() ConfigsConfOutput {
	return o
}

func (o ConfigsConfOutput) ToConfigsConfOutputWithContext(ctx context.Context) ConfigsConfOutput {
	return o
}

func (o ConfigsConfOutput) Acl() ConfigsConfAclOutput {
	return o.ApplyT(func(v *ConfigsConf) ConfigsConfAclOutput { return v.Acl }).(ConfigsConfAclOutput)
}

// A '/' separated string consisting of {conf_file_name}/{stanza_name} ex. props/custom_stanza
func (o ConfigsConfOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ConfigsConf) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// A map of key value pairs for a stanza.
func (o ConfigsConfOutput) Variables() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ConfigsConf) pulumi.StringMapOutput { return v.Variables }).(pulumi.StringMapOutput)
}

type ConfigsConfArrayOutput struct{ *pulumi.OutputState }

func (ConfigsConfArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ConfigsConf)(nil)).Elem()
}

func (o ConfigsConfArrayOutput) ToConfigsConfArrayOutput() ConfigsConfArrayOutput {
	return o
}

func (o ConfigsConfArrayOutput) ToConfigsConfArrayOutputWithContext(ctx context.Context) ConfigsConfArrayOutput {
	return o
}

func (o ConfigsConfArrayOutput) Index(i pulumi.IntInput) ConfigsConfOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ConfigsConf {
		return vs[0].([]*ConfigsConf)[vs[1].(int)]
	}).(ConfigsConfOutput)
}

type ConfigsConfMapOutput struct{ *pulumi.OutputState }

func (ConfigsConfMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ConfigsConf)(nil)).Elem()
}

func (o ConfigsConfMapOutput) ToConfigsConfMapOutput() ConfigsConfMapOutput {
	return o
}

func (o ConfigsConfMapOutput) ToConfigsConfMapOutputWithContext(ctx context.Context) ConfigsConfMapOutput {
	return o
}

func (o ConfigsConfMapOutput) MapIndex(k pulumi.StringInput) ConfigsConfOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ConfigsConf {
		return vs[0].(map[string]*ConfigsConf)[vs[1].(string)]
	}).(ConfigsConfOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ConfigsConfInput)(nil)).Elem(), &ConfigsConf{})
	pulumi.RegisterInputType(reflect.TypeOf((*ConfigsConfArrayInput)(nil)).Elem(), ConfigsConfArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ConfigsConfMapInput)(nil)).Elem(), ConfigsConfMap{})
	pulumi.RegisterOutputType(ConfigsConfOutput{})
	pulumi.RegisterOutputType(ConfigsConfArrayOutput{})
	pulumi.RegisterOutputType(ConfigsConfMapOutput{})
}
