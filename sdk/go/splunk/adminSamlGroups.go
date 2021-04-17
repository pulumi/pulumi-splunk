// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package splunk

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## # Resource: AdminSamlGroups
//
// Manage external groups in an IdP response to internal Splunk roles.
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
// 		_, err := splunk.NewAdminSamlGroups(ctx, "saml_group", &splunk.AdminSamlGroupsArgs{
// 			Roles: pulumi.StringArray{
// 				pulumi.String("admin"),
// 				pulumi.String("power"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ## Import
//
// SAML groups can be imported using the id, e.g.
//
// ```sh
//  $ pulumi import splunk:index/adminSamlGroups:AdminSamlGroups saml-group mygroup
// ```
type AdminSamlGroups struct {
	pulumi.CustomResourceState

	// The name of the external group.
	Name pulumi.StringOutput `pulumi:"name"`
	// List of internal roles assigned to the group.
	Roles pulumi.StringArrayOutput `pulumi:"roles"`
}

// NewAdminSamlGroups registers a new resource with the given unique name, arguments, and options.
func NewAdminSamlGroups(ctx *pulumi.Context,
	name string, args *AdminSamlGroupsArgs, opts ...pulumi.ResourceOption) (*AdminSamlGroups, error) {
	if args == nil {
		args = &AdminSamlGroupsArgs{}
	}

	var resource AdminSamlGroups
	err := ctx.RegisterResource("splunk:index/adminSamlGroups:AdminSamlGroups", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAdminSamlGroups gets an existing AdminSamlGroups resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAdminSamlGroups(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AdminSamlGroupsState, opts ...pulumi.ResourceOption) (*AdminSamlGroups, error) {
	var resource AdminSamlGroups
	err := ctx.ReadResource("splunk:index/adminSamlGroups:AdminSamlGroups", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AdminSamlGroups resources.
type adminSamlGroupsState struct {
	// The name of the external group.
	Name *string `pulumi:"name"`
	// List of internal roles assigned to the group.
	Roles []string `pulumi:"roles"`
}

type AdminSamlGroupsState struct {
	// The name of the external group.
	Name pulumi.StringPtrInput
	// List of internal roles assigned to the group.
	Roles pulumi.StringArrayInput
}

func (AdminSamlGroupsState) ElementType() reflect.Type {
	return reflect.TypeOf((*adminSamlGroupsState)(nil)).Elem()
}

type adminSamlGroupsArgs struct {
	// The name of the external group.
	Name *string `pulumi:"name"`
	// List of internal roles assigned to the group.
	Roles []string `pulumi:"roles"`
}

// The set of arguments for constructing a AdminSamlGroups resource.
type AdminSamlGroupsArgs struct {
	// The name of the external group.
	Name pulumi.StringPtrInput
	// List of internal roles assigned to the group.
	Roles pulumi.StringArrayInput
}

func (AdminSamlGroupsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*adminSamlGroupsArgs)(nil)).Elem()
}

type AdminSamlGroupsInput interface {
	pulumi.Input

	ToAdminSamlGroupsOutput() AdminSamlGroupsOutput
	ToAdminSamlGroupsOutputWithContext(ctx context.Context) AdminSamlGroupsOutput
}

func (*AdminSamlGroups) ElementType() reflect.Type {
	return reflect.TypeOf((*AdminSamlGroups)(nil))
}

func (i *AdminSamlGroups) ToAdminSamlGroupsOutput() AdminSamlGroupsOutput {
	return i.ToAdminSamlGroupsOutputWithContext(context.Background())
}

func (i *AdminSamlGroups) ToAdminSamlGroupsOutputWithContext(ctx context.Context) AdminSamlGroupsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AdminSamlGroupsOutput)
}

func (i *AdminSamlGroups) ToAdminSamlGroupsPtrOutput() AdminSamlGroupsPtrOutput {
	return i.ToAdminSamlGroupsPtrOutputWithContext(context.Background())
}

func (i *AdminSamlGroups) ToAdminSamlGroupsPtrOutputWithContext(ctx context.Context) AdminSamlGroupsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AdminSamlGroupsPtrOutput)
}

type AdminSamlGroupsPtrInput interface {
	pulumi.Input

	ToAdminSamlGroupsPtrOutput() AdminSamlGroupsPtrOutput
	ToAdminSamlGroupsPtrOutputWithContext(ctx context.Context) AdminSamlGroupsPtrOutput
}

type adminSamlGroupsPtrType AdminSamlGroupsArgs

func (*adminSamlGroupsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AdminSamlGroups)(nil))
}

func (i *adminSamlGroupsPtrType) ToAdminSamlGroupsPtrOutput() AdminSamlGroupsPtrOutput {
	return i.ToAdminSamlGroupsPtrOutputWithContext(context.Background())
}

func (i *adminSamlGroupsPtrType) ToAdminSamlGroupsPtrOutputWithContext(ctx context.Context) AdminSamlGroupsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AdminSamlGroupsPtrOutput)
}

// AdminSamlGroupsArrayInput is an input type that accepts AdminSamlGroupsArray and AdminSamlGroupsArrayOutput values.
// You can construct a concrete instance of `AdminSamlGroupsArrayInput` via:
//
//          AdminSamlGroupsArray{ AdminSamlGroupsArgs{...} }
type AdminSamlGroupsArrayInput interface {
	pulumi.Input

	ToAdminSamlGroupsArrayOutput() AdminSamlGroupsArrayOutput
	ToAdminSamlGroupsArrayOutputWithContext(context.Context) AdminSamlGroupsArrayOutput
}

type AdminSamlGroupsArray []AdminSamlGroupsInput

func (AdminSamlGroupsArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*AdminSamlGroups)(nil))
}

func (i AdminSamlGroupsArray) ToAdminSamlGroupsArrayOutput() AdminSamlGroupsArrayOutput {
	return i.ToAdminSamlGroupsArrayOutputWithContext(context.Background())
}

func (i AdminSamlGroupsArray) ToAdminSamlGroupsArrayOutputWithContext(ctx context.Context) AdminSamlGroupsArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AdminSamlGroupsArrayOutput)
}

// AdminSamlGroupsMapInput is an input type that accepts AdminSamlGroupsMap and AdminSamlGroupsMapOutput values.
// You can construct a concrete instance of `AdminSamlGroupsMapInput` via:
//
//          AdminSamlGroupsMap{ "key": AdminSamlGroupsArgs{...} }
type AdminSamlGroupsMapInput interface {
	pulumi.Input

	ToAdminSamlGroupsMapOutput() AdminSamlGroupsMapOutput
	ToAdminSamlGroupsMapOutputWithContext(context.Context) AdminSamlGroupsMapOutput
}

type AdminSamlGroupsMap map[string]AdminSamlGroupsInput

func (AdminSamlGroupsMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*AdminSamlGroups)(nil))
}

func (i AdminSamlGroupsMap) ToAdminSamlGroupsMapOutput() AdminSamlGroupsMapOutput {
	return i.ToAdminSamlGroupsMapOutputWithContext(context.Background())
}

func (i AdminSamlGroupsMap) ToAdminSamlGroupsMapOutputWithContext(ctx context.Context) AdminSamlGroupsMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AdminSamlGroupsMapOutput)
}

type AdminSamlGroupsOutput struct {
	*pulumi.OutputState
}

func (AdminSamlGroupsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AdminSamlGroups)(nil))
}

func (o AdminSamlGroupsOutput) ToAdminSamlGroupsOutput() AdminSamlGroupsOutput {
	return o
}

func (o AdminSamlGroupsOutput) ToAdminSamlGroupsOutputWithContext(ctx context.Context) AdminSamlGroupsOutput {
	return o
}

func (o AdminSamlGroupsOutput) ToAdminSamlGroupsPtrOutput() AdminSamlGroupsPtrOutput {
	return o.ToAdminSamlGroupsPtrOutputWithContext(context.Background())
}

func (o AdminSamlGroupsOutput) ToAdminSamlGroupsPtrOutputWithContext(ctx context.Context) AdminSamlGroupsPtrOutput {
	return o.ApplyT(func(v AdminSamlGroups) *AdminSamlGroups {
		return &v
	}).(AdminSamlGroupsPtrOutput)
}

type AdminSamlGroupsPtrOutput struct {
	*pulumi.OutputState
}

func (AdminSamlGroupsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AdminSamlGroups)(nil))
}

func (o AdminSamlGroupsPtrOutput) ToAdminSamlGroupsPtrOutput() AdminSamlGroupsPtrOutput {
	return o
}

func (o AdminSamlGroupsPtrOutput) ToAdminSamlGroupsPtrOutputWithContext(ctx context.Context) AdminSamlGroupsPtrOutput {
	return o
}

type AdminSamlGroupsArrayOutput struct{ *pulumi.OutputState }

func (AdminSamlGroupsArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AdminSamlGroups)(nil))
}

func (o AdminSamlGroupsArrayOutput) ToAdminSamlGroupsArrayOutput() AdminSamlGroupsArrayOutput {
	return o
}

func (o AdminSamlGroupsArrayOutput) ToAdminSamlGroupsArrayOutputWithContext(ctx context.Context) AdminSamlGroupsArrayOutput {
	return o
}

func (o AdminSamlGroupsArrayOutput) Index(i pulumi.IntInput) AdminSamlGroupsOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) AdminSamlGroups {
		return vs[0].([]AdminSamlGroups)[vs[1].(int)]
	}).(AdminSamlGroupsOutput)
}

type AdminSamlGroupsMapOutput struct{ *pulumi.OutputState }

func (AdminSamlGroupsMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]AdminSamlGroups)(nil))
}

func (o AdminSamlGroupsMapOutput) ToAdminSamlGroupsMapOutput() AdminSamlGroupsMapOutput {
	return o
}

func (o AdminSamlGroupsMapOutput) ToAdminSamlGroupsMapOutputWithContext(ctx context.Context) AdminSamlGroupsMapOutput {
	return o
}

func (o AdminSamlGroupsMapOutput) MapIndex(k pulumi.StringInput) AdminSamlGroupsOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) AdminSamlGroups {
		return vs[0].(map[string]AdminSamlGroups)[vs[1].(string)]
	}).(AdminSamlGroupsOutput)
}

func init() {
	pulumi.RegisterOutputType(AdminSamlGroupsOutput{})
	pulumi.RegisterOutputType(AdminSamlGroupsPtrOutput{})
	pulumi.RegisterOutputType(AdminSamlGroupsArrayOutput{})
	pulumi.RegisterOutputType(AdminSamlGroupsMapOutput{})
}
