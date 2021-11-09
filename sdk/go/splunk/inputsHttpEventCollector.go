// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package splunk

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## # Resource: InputsHttpEventCollector
//
// Create or update HTTP Event Collector input configuration tokens.
type InputsHttpEventCollector struct {
	pulumi.CustomResourceState

	// The app/user context that is the namespace for the resource
	Acl InputsHttpEventCollectorAclOutput `pulumi:"acl"`
	// Input disabled indicator
	Disabled pulumi.BoolOutput `pulumi:"disabled"`
	// Default host value for events with this token
	Host pulumi.StringOutput `pulumi:"host"`
	// Index to store generated events
	Index pulumi.StringOutput `pulumi:"index"`
	// Set of indexes allowed for events with this token
	Indexes pulumi.StringArrayOutput `pulumi:"indexes"`
	// Token name (inputs.conf key)
	Name pulumi.StringOutput `pulumi:"name"`
	// Default source for events with this token
	Source pulumi.StringOutput `pulumi:"source"`
	// Default source type for events with this token
	Sourcetype pulumi.StringOutput `pulumi:"sourcetype"`
	// Token value for sending data to collector/event endpoint
	Token pulumi.StringOutput `pulumi:"token"`
	// Indexer acknowledgement for this token
	UseAck pulumi.IntOutput `pulumi:"useAck"`
}

// NewInputsHttpEventCollector registers a new resource with the given unique name, arguments, and options.
func NewInputsHttpEventCollector(ctx *pulumi.Context,
	name string, args *InputsHttpEventCollectorArgs, opts ...pulumi.ResourceOption) (*InputsHttpEventCollector, error) {
	if args == nil {
		args = &InputsHttpEventCollectorArgs{}
	}

	var resource InputsHttpEventCollector
	err := ctx.RegisterResource("splunk:index/inputsHttpEventCollector:InputsHttpEventCollector", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetInputsHttpEventCollector gets an existing InputsHttpEventCollector resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetInputsHttpEventCollector(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *InputsHttpEventCollectorState, opts ...pulumi.ResourceOption) (*InputsHttpEventCollector, error) {
	var resource InputsHttpEventCollector
	err := ctx.ReadResource("splunk:index/inputsHttpEventCollector:InputsHttpEventCollector", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering InputsHttpEventCollector resources.
type inputsHttpEventCollectorState struct {
	// The app/user context that is the namespace for the resource
	Acl *InputsHttpEventCollectorAcl `pulumi:"acl"`
	// Input disabled indicator
	Disabled *bool `pulumi:"disabled"`
	// Default host value for events with this token
	Host *string `pulumi:"host"`
	// Index to store generated events
	Index *string `pulumi:"index"`
	// Set of indexes allowed for events with this token
	Indexes []string `pulumi:"indexes"`
	// Token name (inputs.conf key)
	Name *string `pulumi:"name"`
	// Default source for events with this token
	Source *string `pulumi:"source"`
	// Default source type for events with this token
	Sourcetype *string `pulumi:"sourcetype"`
	// Token value for sending data to collector/event endpoint
	Token *string `pulumi:"token"`
	// Indexer acknowledgement for this token
	UseAck *int `pulumi:"useAck"`
}

type InputsHttpEventCollectorState struct {
	// The app/user context that is the namespace for the resource
	Acl InputsHttpEventCollectorAclPtrInput
	// Input disabled indicator
	Disabled pulumi.BoolPtrInput
	// Default host value for events with this token
	Host pulumi.StringPtrInput
	// Index to store generated events
	Index pulumi.StringPtrInput
	// Set of indexes allowed for events with this token
	Indexes pulumi.StringArrayInput
	// Token name (inputs.conf key)
	Name pulumi.StringPtrInput
	// Default source for events with this token
	Source pulumi.StringPtrInput
	// Default source type for events with this token
	Sourcetype pulumi.StringPtrInput
	// Token value for sending data to collector/event endpoint
	Token pulumi.StringPtrInput
	// Indexer acknowledgement for this token
	UseAck pulumi.IntPtrInput
}

func (InputsHttpEventCollectorState) ElementType() reflect.Type {
	return reflect.TypeOf((*inputsHttpEventCollectorState)(nil)).Elem()
}

type inputsHttpEventCollectorArgs struct {
	// The app/user context that is the namespace for the resource
	Acl *InputsHttpEventCollectorAcl `pulumi:"acl"`
	// Input disabled indicator
	Disabled *bool `pulumi:"disabled"`
	// Default host value for events with this token
	Host *string `pulumi:"host"`
	// Index to store generated events
	Index *string `pulumi:"index"`
	// Set of indexes allowed for events with this token
	Indexes []string `pulumi:"indexes"`
	// Token name (inputs.conf key)
	Name *string `pulumi:"name"`
	// Default source for events with this token
	Source *string `pulumi:"source"`
	// Default source type for events with this token
	Sourcetype *string `pulumi:"sourcetype"`
	// Token value for sending data to collector/event endpoint
	Token *string `pulumi:"token"`
	// Indexer acknowledgement for this token
	UseAck *int `pulumi:"useAck"`
}

// The set of arguments for constructing a InputsHttpEventCollector resource.
type InputsHttpEventCollectorArgs struct {
	// The app/user context that is the namespace for the resource
	Acl InputsHttpEventCollectorAclPtrInput
	// Input disabled indicator
	Disabled pulumi.BoolPtrInput
	// Default host value for events with this token
	Host pulumi.StringPtrInput
	// Index to store generated events
	Index pulumi.StringPtrInput
	// Set of indexes allowed for events with this token
	Indexes pulumi.StringArrayInput
	// Token name (inputs.conf key)
	Name pulumi.StringPtrInput
	// Default source for events with this token
	Source pulumi.StringPtrInput
	// Default source type for events with this token
	Sourcetype pulumi.StringPtrInput
	// Token value for sending data to collector/event endpoint
	Token pulumi.StringPtrInput
	// Indexer acknowledgement for this token
	UseAck pulumi.IntPtrInput
}

func (InputsHttpEventCollectorArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*inputsHttpEventCollectorArgs)(nil)).Elem()
}

type InputsHttpEventCollectorInput interface {
	pulumi.Input

	ToInputsHttpEventCollectorOutput() InputsHttpEventCollectorOutput
	ToInputsHttpEventCollectorOutputWithContext(ctx context.Context) InputsHttpEventCollectorOutput
}

func (*InputsHttpEventCollector) ElementType() reflect.Type {
	return reflect.TypeOf((*InputsHttpEventCollector)(nil))
}

func (i *InputsHttpEventCollector) ToInputsHttpEventCollectorOutput() InputsHttpEventCollectorOutput {
	return i.ToInputsHttpEventCollectorOutputWithContext(context.Background())
}

func (i *InputsHttpEventCollector) ToInputsHttpEventCollectorOutputWithContext(ctx context.Context) InputsHttpEventCollectorOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InputsHttpEventCollectorOutput)
}

func (i *InputsHttpEventCollector) ToInputsHttpEventCollectorPtrOutput() InputsHttpEventCollectorPtrOutput {
	return i.ToInputsHttpEventCollectorPtrOutputWithContext(context.Background())
}

func (i *InputsHttpEventCollector) ToInputsHttpEventCollectorPtrOutputWithContext(ctx context.Context) InputsHttpEventCollectorPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InputsHttpEventCollectorPtrOutput)
}

type InputsHttpEventCollectorPtrInput interface {
	pulumi.Input

	ToInputsHttpEventCollectorPtrOutput() InputsHttpEventCollectorPtrOutput
	ToInputsHttpEventCollectorPtrOutputWithContext(ctx context.Context) InputsHttpEventCollectorPtrOutput
}

type inputsHttpEventCollectorPtrType InputsHttpEventCollectorArgs

func (*inputsHttpEventCollectorPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**InputsHttpEventCollector)(nil))
}

func (i *inputsHttpEventCollectorPtrType) ToInputsHttpEventCollectorPtrOutput() InputsHttpEventCollectorPtrOutput {
	return i.ToInputsHttpEventCollectorPtrOutputWithContext(context.Background())
}

func (i *inputsHttpEventCollectorPtrType) ToInputsHttpEventCollectorPtrOutputWithContext(ctx context.Context) InputsHttpEventCollectorPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InputsHttpEventCollectorPtrOutput)
}

// InputsHttpEventCollectorArrayInput is an input type that accepts InputsHttpEventCollectorArray and InputsHttpEventCollectorArrayOutput values.
// You can construct a concrete instance of `InputsHttpEventCollectorArrayInput` via:
//
//          InputsHttpEventCollectorArray{ InputsHttpEventCollectorArgs{...} }
type InputsHttpEventCollectorArrayInput interface {
	pulumi.Input

	ToInputsHttpEventCollectorArrayOutput() InputsHttpEventCollectorArrayOutput
	ToInputsHttpEventCollectorArrayOutputWithContext(context.Context) InputsHttpEventCollectorArrayOutput
}

type InputsHttpEventCollectorArray []InputsHttpEventCollectorInput

func (InputsHttpEventCollectorArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*InputsHttpEventCollector)(nil)).Elem()
}

func (i InputsHttpEventCollectorArray) ToInputsHttpEventCollectorArrayOutput() InputsHttpEventCollectorArrayOutput {
	return i.ToInputsHttpEventCollectorArrayOutputWithContext(context.Background())
}

func (i InputsHttpEventCollectorArray) ToInputsHttpEventCollectorArrayOutputWithContext(ctx context.Context) InputsHttpEventCollectorArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InputsHttpEventCollectorArrayOutput)
}

// InputsHttpEventCollectorMapInput is an input type that accepts InputsHttpEventCollectorMap and InputsHttpEventCollectorMapOutput values.
// You can construct a concrete instance of `InputsHttpEventCollectorMapInput` via:
//
//          InputsHttpEventCollectorMap{ "key": InputsHttpEventCollectorArgs{...} }
type InputsHttpEventCollectorMapInput interface {
	pulumi.Input

	ToInputsHttpEventCollectorMapOutput() InputsHttpEventCollectorMapOutput
	ToInputsHttpEventCollectorMapOutputWithContext(context.Context) InputsHttpEventCollectorMapOutput
}

type InputsHttpEventCollectorMap map[string]InputsHttpEventCollectorInput

func (InputsHttpEventCollectorMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*InputsHttpEventCollector)(nil)).Elem()
}

func (i InputsHttpEventCollectorMap) ToInputsHttpEventCollectorMapOutput() InputsHttpEventCollectorMapOutput {
	return i.ToInputsHttpEventCollectorMapOutputWithContext(context.Background())
}

func (i InputsHttpEventCollectorMap) ToInputsHttpEventCollectorMapOutputWithContext(ctx context.Context) InputsHttpEventCollectorMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InputsHttpEventCollectorMapOutput)
}

type InputsHttpEventCollectorOutput struct{ *pulumi.OutputState }

func (InputsHttpEventCollectorOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*InputsHttpEventCollector)(nil))
}

func (o InputsHttpEventCollectorOutput) ToInputsHttpEventCollectorOutput() InputsHttpEventCollectorOutput {
	return o
}

func (o InputsHttpEventCollectorOutput) ToInputsHttpEventCollectorOutputWithContext(ctx context.Context) InputsHttpEventCollectorOutput {
	return o
}

func (o InputsHttpEventCollectorOutput) ToInputsHttpEventCollectorPtrOutput() InputsHttpEventCollectorPtrOutput {
	return o.ToInputsHttpEventCollectorPtrOutputWithContext(context.Background())
}

func (o InputsHttpEventCollectorOutput) ToInputsHttpEventCollectorPtrOutputWithContext(ctx context.Context) InputsHttpEventCollectorPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v InputsHttpEventCollector) *InputsHttpEventCollector {
		return &v
	}).(InputsHttpEventCollectorPtrOutput)
}

type InputsHttpEventCollectorPtrOutput struct{ *pulumi.OutputState }

func (InputsHttpEventCollectorPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**InputsHttpEventCollector)(nil))
}

func (o InputsHttpEventCollectorPtrOutput) ToInputsHttpEventCollectorPtrOutput() InputsHttpEventCollectorPtrOutput {
	return o
}

func (o InputsHttpEventCollectorPtrOutput) ToInputsHttpEventCollectorPtrOutputWithContext(ctx context.Context) InputsHttpEventCollectorPtrOutput {
	return o
}

func (o InputsHttpEventCollectorPtrOutput) Elem() InputsHttpEventCollectorOutput {
	return o.ApplyT(func(v *InputsHttpEventCollector) InputsHttpEventCollector {
		if v != nil {
			return *v
		}
		var ret InputsHttpEventCollector
		return ret
	}).(InputsHttpEventCollectorOutput)
}

type InputsHttpEventCollectorArrayOutput struct{ *pulumi.OutputState }

func (InputsHttpEventCollectorArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]InputsHttpEventCollector)(nil))
}

func (o InputsHttpEventCollectorArrayOutput) ToInputsHttpEventCollectorArrayOutput() InputsHttpEventCollectorArrayOutput {
	return o
}

func (o InputsHttpEventCollectorArrayOutput) ToInputsHttpEventCollectorArrayOutputWithContext(ctx context.Context) InputsHttpEventCollectorArrayOutput {
	return o
}

func (o InputsHttpEventCollectorArrayOutput) Index(i pulumi.IntInput) InputsHttpEventCollectorOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) InputsHttpEventCollector {
		return vs[0].([]InputsHttpEventCollector)[vs[1].(int)]
	}).(InputsHttpEventCollectorOutput)
}

type InputsHttpEventCollectorMapOutput struct{ *pulumi.OutputState }

func (InputsHttpEventCollectorMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]InputsHttpEventCollector)(nil))
}

func (o InputsHttpEventCollectorMapOutput) ToInputsHttpEventCollectorMapOutput() InputsHttpEventCollectorMapOutput {
	return o
}

func (o InputsHttpEventCollectorMapOutput) ToInputsHttpEventCollectorMapOutputWithContext(ctx context.Context) InputsHttpEventCollectorMapOutput {
	return o
}

func (o InputsHttpEventCollectorMapOutput) MapIndex(k pulumi.StringInput) InputsHttpEventCollectorOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) InputsHttpEventCollector {
		return vs[0].(map[string]InputsHttpEventCollector)[vs[1].(string)]
	}).(InputsHttpEventCollectorOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*InputsHttpEventCollectorInput)(nil)).Elem(), &InputsHttpEventCollector{})
	pulumi.RegisterInputType(reflect.TypeOf((*InputsHttpEventCollectorPtrInput)(nil)).Elem(), &InputsHttpEventCollector{})
	pulumi.RegisterInputType(reflect.TypeOf((*InputsHttpEventCollectorArrayInput)(nil)).Elem(), InputsHttpEventCollectorArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*InputsHttpEventCollectorMapInput)(nil)).Elem(), InputsHttpEventCollectorMap{})
	pulumi.RegisterOutputType(InputsHttpEventCollectorOutput{})
	pulumi.RegisterOutputType(InputsHttpEventCollectorPtrOutput{})
	pulumi.RegisterOutputType(InputsHttpEventCollectorArrayOutput{})
	pulumi.RegisterOutputType(InputsHttpEventCollectorMapOutput{})
}
