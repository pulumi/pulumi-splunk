// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package splunk

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
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
// 	"github.com/pulumi/pulumi-splunk/sdk/go/splunk/"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
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

func init() {
	pulumi.RegisterOutputType(AdminSamlGroupsOutput{})
}
