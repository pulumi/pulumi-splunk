// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package splunk

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## # Resource: ShIndexesManager
//
// Create indexes on Splunk Cloud instances. [BETA]
//
// ## Authorization and authentication
//
// As of now there is no support to create indexes in user-specified workspaces on Splunk Cloud.
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
// 		_, err := splunk.NewShIndexesManager(ctx, "tf_index", &splunk.ShIndexesManagerArgs{
// 			Datatype:               pulumi.String("event"),
// 			FrozenTimePeriodInSecs: pulumi.String("94608000"),
// 			MaxGlobalRawDataSizeMb: pulumi.String("100"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type ShIndexesManager struct {
	pulumi.CustomResourceState

	Acl ShIndexesManagerAclOutput `pulumi:"acl"`
	// Valid values: (event | metric). Specifies the type of index.
	Datatype pulumi.StringOutput `pulumi:"datatype"`
	// Number of seconds after which indexed data rolls to frozen.
	// Defaults to 94608000 (3 years).Freezing data means it is removed from the index. If you need to archive your data, refer to coldToFrozenDir and coldToFrozenScript parameter documentation.
	FrozenTimePeriodInSecs pulumi.StringPtrOutput `pulumi:"frozenTimePeriodInSecs"`
	// The maximum size of an index (in MB). If an index grows larger than the maximum size, the oldest data is frozen.
	// Defaults to 100 MB.
	MaxGlobalRawDataSizeMb pulumi.StringPtrOutput `pulumi:"maxGlobalRawDataSizeMb"`
	// The name of the index to create.
	Name pulumi.StringOutput `pulumi:"name"`
}

// NewShIndexesManager registers a new resource with the given unique name, arguments, and options.
func NewShIndexesManager(ctx *pulumi.Context,
	name string, args *ShIndexesManagerArgs, opts ...pulumi.ResourceOption) (*ShIndexesManager, error) {
	if args == nil {
		args = &ShIndexesManagerArgs{}
	}

	var resource ShIndexesManager
	err := ctx.RegisterResource("splunk:index/shIndexesManager:ShIndexesManager", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetShIndexesManager gets an existing ShIndexesManager resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetShIndexesManager(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ShIndexesManagerState, opts ...pulumi.ResourceOption) (*ShIndexesManager, error) {
	var resource ShIndexesManager
	err := ctx.ReadResource("splunk:index/shIndexesManager:ShIndexesManager", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ShIndexesManager resources.
type shIndexesManagerState struct {
	Acl *ShIndexesManagerAcl `pulumi:"acl"`
	// Valid values: (event | metric). Specifies the type of index.
	Datatype *string `pulumi:"datatype"`
	// Number of seconds after which indexed data rolls to frozen.
	// Defaults to 94608000 (3 years).Freezing data means it is removed from the index. If you need to archive your data, refer to coldToFrozenDir and coldToFrozenScript parameter documentation.
	FrozenTimePeriodInSecs *string `pulumi:"frozenTimePeriodInSecs"`
	// The maximum size of an index (in MB). If an index grows larger than the maximum size, the oldest data is frozen.
	// Defaults to 100 MB.
	MaxGlobalRawDataSizeMb *string `pulumi:"maxGlobalRawDataSizeMb"`
	// The name of the index to create.
	Name *string `pulumi:"name"`
}

type ShIndexesManagerState struct {
	Acl ShIndexesManagerAclPtrInput
	// Valid values: (event | metric). Specifies the type of index.
	Datatype pulumi.StringPtrInput
	// Number of seconds after which indexed data rolls to frozen.
	// Defaults to 94608000 (3 years).Freezing data means it is removed from the index. If you need to archive your data, refer to coldToFrozenDir and coldToFrozenScript parameter documentation.
	FrozenTimePeriodInSecs pulumi.StringPtrInput
	// The maximum size of an index (in MB). If an index grows larger than the maximum size, the oldest data is frozen.
	// Defaults to 100 MB.
	MaxGlobalRawDataSizeMb pulumi.StringPtrInput
	// The name of the index to create.
	Name pulumi.StringPtrInput
}

func (ShIndexesManagerState) ElementType() reflect.Type {
	return reflect.TypeOf((*shIndexesManagerState)(nil)).Elem()
}

type shIndexesManagerArgs struct {
	Acl *ShIndexesManagerAcl `pulumi:"acl"`
	// Valid values: (event | metric). Specifies the type of index.
	Datatype *string `pulumi:"datatype"`
	// Number of seconds after which indexed data rolls to frozen.
	// Defaults to 94608000 (3 years).Freezing data means it is removed from the index. If you need to archive your data, refer to coldToFrozenDir and coldToFrozenScript parameter documentation.
	FrozenTimePeriodInSecs *string `pulumi:"frozenTimePeriodInSecs"`
	// The maximum size of an index (in MB). If an index grows larger than the maximum size, the oldest data is frozen.
	// Defaults to 100 MB.
	MaxGlobalRawDataSizeMb *string `pulumi:"maxGlobalRawDataSizeMb"`
	// The name of the index to create.
	Name *string `pulumi:"name"`
}

// The set of arguments for constructing a ShIndexesManager resource.
type ShIndexesManagerArgs struct {
	Acl ShIndexesManagerAclPtrInput
	// Valid values: (event | metric). Specifies the type of index.
	Datatype pulumi.StringPtrInput
	// Number of seconds after which indexed data rolls to frozen.
	// Defaults to 94608000 (3 years).Freezing data means it is removed from the index. If you need to archive your data, refer to coldToFrozenDir and coldToFrozenScript parameter documentation.
	FrozenTimePeriodInSecs pulumi.StringPtrInput
	// The maximum size of an index (in MB). If an index grows larger than the maximum size, the oldest data is frozen.
	// Defaults to 100 MB.
	MaxGlobalRawDataSizeMb pulumi.StringPtrInput
	// The name of the index to create.
	Name pulumi.StringPtrInput
}

func (ShIndexesManagerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*shIndexesManagerArgs)(nil)).Elem()
}

type ShIndexesManagerInput interface {
	pulumi.Input

	ToShIndexesManagerOutput() ShIndexesManagerOutput
	ToShIndexesManagerOutputWithContext(ctx context.Context) ShIndexesManagerOutput
}

func (*ShIndexesManager) ElementType() reflect.Type {
	return reflect.TypeOf((*ShIndexesManager)(nil))
}

func (i *ShIndexesManager) ToShIndexesManagerOutput() ShIndexesManagerOutput {
	return i.ToShIndexesManagerOutputWithContext(context.Background())
}

func (i *ShIndexesManager) ToShIndexesManagerOutputWithContext(ctx context.Context) ShIndexesManagerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ShIndexesManagerOutput)
}

func (i *ShIndexesManager) ToShIndexesManagerPtrOutput() ShIndexesManagerPtrOutput {
	return i.ToShIndexesManagerPtrOutputWithContext(context.Background())
}

func (i *ShIndexesManager) ToShIndexesManagerPtrOutputWithContext(ctx context.Context) ShIndexesManagerPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ShIndexesManagerPtrOutput)
}

type ShIndexesManagerPtrInput interface {
	pulumi.Input

	ToShIndexesManagerPtrOutput() ShIndexesManagerPtrOutput
	ToShIndexesManagerPtrOutputWithContext(ctx context.Context) ShIndexesManagerPtrOutput
}

type shIndexesManagerPtrType ShIndexesManagerArgs

func (*shIndexesManagerPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ShIndexesManager)(nil))
}

func (i *shIndexesManagerPtrType) ToShIndexesManagerPtrOutput() ShIndexesManagerPtrOutput {
	return i.ToShIndexesManagerPtrOutputWithContext(context.Background())
}

func (i *shIndexesManagerPtrType) ToShIndexesManagerPtrOutputWithContext(ctx context.Context) ShIndexesManagerPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ShIndexesManagerPtrOutput)
}

// ShIndexesManagerArrayInput is an input type that accepts ShIndexesManagerArray and ShIndexesManagerArrayOutput values.
// You can construct a concrete instance of `ShIndexesManagerArrayInput` via:
//
//          ShIndexesManagerArray{ ShIndexesManagerArgs{...} }
type ShIndexesManagerArrayInput interface {
	pulumi.Input

	ToShIndexesManagerArrayOutput() ShIndexesManagerArrayOutput
	ToShIndexesManagerArrayOutputWithContext(context.Context) ShIndexesManagerArrayOutput
}

type ShIndexesManagerArray []ShIndexesManagerInput

func (ShIndexesManagerArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*ShIndexesManager)(nil))
}

func (i ShIndexesManagerArray) ToShIndexesManagerArrayOutput() ShIndexesManagerArrayOutput {
	return i.ToShIndexesManagerArrayOutputWithContext(context.Background())
}

func (i ShIndexesManagerArray) ToShIndexesManagerArrayOutputWithContext(ctx context.Context) ShIndexesManagerArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ShIndexesManagerArrayOutput)
}

// ShIndexesManagerMapInput is an input type that accepts ShIndexesManagerMap and ShIndexesManagerMapOutput values.
// You can construct a concrete instance of `ShIndexesManagerMapInput` via:
//
//          ShIndexesManagerMap{ "key": ShIndexesManagerArgs{...} }
type ShIndexesManagerMapInput interface {
	pulumi.Input

	ToShIndexesManagerMapOutput() ShIndexesManagerMapOutput
	ToShIndexesManagerMapOutputWithContext(context.Context) ShIndexesManagerMapOutput
}

type ShIndexesManagerMap map[string]ShIndexesManagerInput

func (ShIndexesManagerMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*ShIndexesManager)(nil))
}

func (i ShIndexesManagerMap) ToShIndexesManagerMapOutput() ShIndexesManagerMapOutput {
	return i.ToShIndexesManagerMapOutputWithContext(context.Background())
}

func (i ShIndexesManagerMap) ToShIndexesManagerMapOutputWithContext(ctx context.Context) ShIndexesManagerMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ShIndexesManagerMapOutput)
}

type ShIndexesManagerOutput struct {
	*pulumi.OutputState
}

func (ShIndexesManagerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ShIndexesManager)(nil))
}

func (o ShIndexesManagerOutput) ToShIndexesManagerOutput() ShIndexesManagerOutput {
	return o
}

func (o ShIndexesManagerOutput) ToShIndexesManagerOutputWithContext(ctx context.Context) ShIndexesManagerOutput {
	return o
}

func (o ShIndexesManagerOutput) ToShIndexesManagerPtrOutput() ShIndexesManagerPtrOutput {
	return o.ToShIndexesManagerPtrOutputWithContext(context.Background())
}

func (o ShIndexesManagerOutput) ToShIndexesManagerPtrOutputWithContext(ctx context.Context) ShIndexesManagerPtrOutput {
	return o.ApplyT(func(v ShIndexesManager) *ShIndexesManager {
		return &v
	}).(ShIndexesManagerPtrOutput)
}

type ShIndexesManagerPtrOutput struct {
	*pulumi.OutputState
}

func (ShIndexesManagerPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ShIndexesManager)(nil))
}

func (o ShIndexesManagerPtrOutput) ToShIndexesManagerPtrOutput() ShIndexesManagerPtrOutput {
	return o
}

func (o ShIndexesManagerPtrOutput) ToShIndexesManagerPtrOutputWithContext(ctx context.Context) ShIndexesManagerPtrOutput {
	return o
}

type ShIndexesManagerArrayOutput struct{ *pulumi.OutputState }

func (ShIndexesManagerArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ShIndexesManager)(nil))
}

func (o ShIndexesManagerArrayOutput) ToShIndexesManagerArrayOutput() ShIndexesManagerArrayOutput {
	return o
}

func (o ShIndexesManagerArrayOutput) ToShIndexesManagerArrayOutputWithContext(ctx context.Context) ShIndexesManagerArrayOutput {
	return o
}

func (o ShIndexesManagerArrayOutput) Index(i pulumi.IntInput) ShIndexesManagerOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ShIndexesManager {
		return vs[0].([]ShIndexesManager)[vs[1].(int)]
	}).(ShIndexesManagerOutput)
}

type ShIndexesManagerMapOutput struct{ *pulumi.OutputState }

func (ShIndexesManagerMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ShIndexesManager)(nil))
}

func (o ShIndexesManagerMapOutput) ToShIndexesManagerMapOutput() ShIndexesManagerMapOutput {
	return o
}

func (o ShIndexesManagerMapOutput) ToShIndexesManagerMapOutputWithContext(ctx context.Context) ShIndexesManagerMapOutput {
	return o
}

func (o ShIndexesManagerMapOutput) MapIndex(k pulumi.StringInput) ShIndexesManagerOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) ShIndexesManager {
		return vs[0].(map[string]ShIndexesManager)[vs[1].(string)]
	}).(ShIndexesManagerOutput)
}

func init() {
	pulumi.RegisterOutputType(ShIndexesManagerOutput{})
	pulumi.RegisterOutputType(ShIndexesManagerPtrOutput{})
	pulumi.RegisterOutputType(ShIndexesManagerArrayOutput{})
	pulumi.RegisterOutputType(ShIndexesManagerMapOutput{})
}
