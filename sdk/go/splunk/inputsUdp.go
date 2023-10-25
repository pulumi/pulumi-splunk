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

// ## # Resource: InputsTcpRaw
//
// Create and manage UDP data inputs.
type InputsUdp struct {
	pulumi.CustomResourceState

	// The app/user context that is the namespace for the resource
	Acl InputsUdpAclOutput `pulumi:"acl"`
	// Valid values: (ip | dns | none)
	// Set the host for the remote server that is sending data.
	// ip sets the host to the IP address of the remote server sending data.
	// dns sets the host to the reverse DNS entry for the IP address of the remote server sending data.
	// none leaves the host as specified in inputs.conf, which is typically the Splunk system hostname.
	// Default value is dns.
	ConnectionHost pulumi.StringOutput `pulumi:"connectionHost"`
	// Indicates if input is disabled.
	Disabled pulumi.BoolOutput `pulumi:"disabled"`
	// The value to populate in the host field for incoming events. This is used during parsing/indexing, in particular to set the host field. It is also the host field used at search time.
	Host pulumi.StringOutput `pulumi:"host"`
	// Which index events from this input should be stored in. Defaults to default.
	Index pulumi.StringOutput `pulumi:"index"`
	// The UDP port that this input should listen on.
	Name pulumi.StringOutput `pulumi:"name"`
	// If set to true, prevents Splunk software from prepending a timestamp and hostname to incoming events.
	NoAppendingTimestamp pulumi.BoolOutput `pulumi:"noAppendingTimestamp"`
	// If set to true, Splunk software does not remove the priority field from incoming syslog events.
	NoPriorityStripping pulumi.BoolOutput `pulumi:"noPriorityStripping"`
	// Which queue events from this input should be sent to. Generally this does not need to be changed.
	Queue pulumi.StringOutput `pulumi:"queue"`
	// Restrict incoming connections on this port to the host specified here.
	// If this is not set, the value specified in [udp://<remote server>:<port>] in inputs.conf is used.
	RestrictToHost pulumi.StringOutput `pulumi:"restrictToHost"`
	// The value to populate in the source field for incoming events. The same source should not be used for multiple data inputs.
	Source pulumi.StringOutput `pulumi:"source"`
	// The value to populate in the sourcetype field for incoming events.
	Sourcetype pulumi.StringOutput `pulumi:"sourcetype"`
}

// NewInputsUdp registers a new resource with the given unique name, arguments, and options.
func NewInputsUdp(ctx *pulumi.Context,
	name string, args *InputsUdpArgs, opts ...pulumi.ResourceOption) (*InputsUdp, error) {
	if args == nil {
		args = &InputsUdpArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource InputsUdp
	err := ctx.RegisterResource("splunk:index/inputsUdp:InputsUdp", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetInputsUdp gets an existing InputsUdp resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetInputsUdp(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *InputsUdpState, opts ...pulumi.ResourceOption) (*InputsUdp, error) {
	var resource InputsUdp
	err := ctx.ReadResource("splunk:index/inputsUdp:InputsUdp", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering InputsUdp resources.
type inputsUdpState struct {
	// The app/user context that is the namespace for the resource
	Acl *InputsUdpAcl `pulumi:"acl"`
	// Valid values: (ip | dns | none)
	// Set the host for the remote server that is sending data.
	// ip sets the host to the IP address of the remote server sending data.
	// dns sets the host to the reverse DNS entry for the IP address of the remote server sending data.
	// none leaves the host as specified in inputs.conf, which is typically the Splunk system hostname.
	// Default value is dns.
	ConnectionHost *string `pulumi:"connectionHost"`
	// Indicates if input is disabled.
	Disabled *bool `pulumi:"disabled"`
	// The value to populate in the host field for incoming events. This is used during parsing/indexing, in particular to set the host field. It is also the host field used at search time.
	Host *string `pulumi:"host"`
	// Which index events from this input should be stored in. Defaults to default.
	Index *string `pulumi:"index"`
	// The UDP port that this input should listen on.
	Name *string `pulumi:"name"`
	// If set to true, prevents Splunk software from prepending a timestamp and hostname to incoming events.
	NoAppendingTimestamp *bool `pulumi:"noAppendingTimestamp"`
	// If set to true, Splunk software does not remove the priority field from incoming syslog events.
	NoPriorityStripping *bool `pulumi:"noPriorityStripping"`
	// Which queue events from this input should be sent to. Generally this does not need to be changed.
	Queue *string `pulumi:"queue"`
	// Restrict incoming connections on this port to the host specified here.
	// If this is not set, the value specified in [udp://<remote server>:<port>] in inputs.conf is used.
	RestrictToHost *string `pulumi:"restrictToHost"`
	// The value to populate in the source field for incoming events. The same source should not be used for multiple data inputs.
	Source *string `pulumi:"source"`
	// The value to populate in the sourcetype field for incoming events.
	Sourcetype *string `pulumi:"sourcetype"`
}

type InputsUdpState struct {
	// The app/user context that is the namespace for the resource
	Acl InputsUdpAclPtrInput
	// Valid values: (ip | dns | none)
	// Set the host for the remote server that is sending data.
	// ip sets the host to the IP address of the remote server sending data.
	// dns sets the host to the reverse DNS entry for the IP address of the remote server sending data.
	// none leaves the host as specified in inputs.conf, which is typically the Splunk system hostname.
	// Default value is dns.
	ConnectionHost pulumi.StringPtrInput
	// Indicates if input is disabled.
	Disabled pulumi.BoolPtrInput
	// The value to populate in the host field for incoming events. This is used during parsing/indexing, in particular to set the host field. It is also the host field used at search time.
	Host pulumi.StringPtrInput
	// Which index events from this input should be stored in. Defaults to default.
	Index pulumi.StringPtrInput
	// The UDP port that this input should listen on.
	Name pulumi.StringPtrInput
	// If set to true, prevents Splunk software from prepending a timestamp and hostname to incoming events.
	NoAppendingTimestamp pulumi.BoolPtrInput
	// If set to true, Splunk software does not remove the priority field from incoming syslog events.
	NoPriorityStripping pulumi.BoolPtrInput
	// Which queue events from this input should be sent to. Generally this does not need to be changed.
	Queue pulumi.StringPtrInput
	// Restrict incoming connections on this port to the host specified here.
	// If this is not set, the value specified in [udp://<remote server>:<port>] in inputs.conf is used.
	RestrictToHost pulumi.StringPtrInput
	// The value to populate in the source field for incoming events. The same source should not be used for multiple data inputs.
	Source pulumi.StringPtrInput
	// The value to populate in the sourcetype field for incoming events.
	Sourcetype pulumi.StringPtrInput
}

func (InputsUdpState) ElementType() reflect.Type {
	return reflect.TypeOf((*inputsUdpState)(nil)).Elem()
}

type inputsUdpArgs struct {
	// The app/user context that is the namespace for the resource
	Acl *InputsUdpAcl `pulumi:"acl"`
	// Valid values: (ip | dns | none)
	// Set the host for the remote server that is sending data.
	// ip sets the host to the IP address of the remote server sending data.
	// dns sets the host to the reverse DNS entry for the IP address of the remote server sending data.
	// none leaves the host as specified in inputs.conf, which is typically the Splunk system hostname.
	// Default value is dns.
	ConnectionHost *string `pulumi:"connectionHost"`
	// Indicates if input is disabled.
	Disabled *bool `pulumi:"disabled"`
	// The value to populate in the host field for incoming events. This is used during parsing/indexing, in particular to set the host field. It is also the host field used at search time.
	Host *string `pulumi:"host"`
	// Which index events from this input should be stored in. Defaults to default.
	Index *string `pulumi:"index"`
	// The UDP port that this input should listen on.
	Name *string `pulumi:"name"`
	// If set to true, prevents Splunk software from prepending a timestamp and hostname to incoming events.
	NoAppendingTimestamp *bool `pulumi:"noAppendingTimestamp"`
	// If set to true, Splunk software does not remove the priority field from incoming syslog events.
	NoPriorityStripping *bool `pulumi:"noPriorityStripping"`
	// Which queue events from this input should be sent to. Generally this does not need to be changed.
	Queue *string `pulumi:"queue"`
	// Restrict incoming connections on this port to the host specified here.
	// If this is not set, the value specified in [udp://<remote server>:<port>] in inputs.conf is used.
	RestrictToHost *string `pulumi:"restrictToHost"`
	// The value to populate in the source field for incoming events. The same source should not be used for multiple data inputs.
	Source *string `pulumi:"source"`
	// The value to populate in the sourcetype field for incoming events.
	Sourcetype *string `pulumi:"sourcetype"`
}

// The set of arguments for constructing a InputsUdp resource.
type InputsUdpArgs struct {
	// The app/user context that is the namespace for the resource
	Acl InputsUdpAclPtrInput
	// Valid values: (ip | dns | none)
	// Set the host for the remote server that is sending data.
	// ip sets the host to the IP address of the remote server sending data.
	// dns sets the host to the reverse DNS entry for the IP address of the remote server sending data.
	// none leaves the host as specified in inputs.conf, which is typically the Splunk system hostname.
	// Default value is dns.
	ConnectionHost pulumi.StringPtrInput
	// Indicates if input is disabled.
	Disabled pulumi.BoolPtrInput
	// The value to populate in the host field for incoming events. This is used during parsing/indexing, in particular to set the host field. It is also the host field used at search time.
	Host pulumi.StringPtrInput
	// Which index events from this input should be stored in. Defaults to default.
	Index pulumi.StringPtrInput
	// The UDP port that this input should listen on.
	Name pulumi.StringPtrInput
	// If set to true, prevents Splunk software from prepending a timestamp and hostname to incoming events.
	NoAppendingTimestamp pulumi.BoolPtrInput
	// If set to true, Splunk software does not remove the priority field from incoming syslog events.
	NoPriorityStripping pulumi.BoolPtrInput
	// Which queue events from this input should be sent to. Generally this does not need to be changed.
	Queue pulumi.StringPtrInput
	// Restrict incoming connections on this port to the host specified here.
	// If this is not set, the value specified in [udp://<remote server>:<port>] in inputs.conf is used.
	RestrictToHost pulumi.StringPtrInput
	// The value to populate in the source field for incoming events. The same source should not be used for multiple data inputs.
	Source pulumi.StringPtrInput
	// The value to populate in the sourcetype field for incoming events.
	Sourcetype pulumi.StringPtrInput
}

func (InputsUdpArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*inputsUdpArgs)(nil)).Elem()
}

type InputsUdpInput interface {
	pulumi.Input

	ToInputsUdpOutput() InputsUdpOutput
	ToInputsUdpOutputWithContext(ctx context.Context) InputsUdpOutput
}

func (*InputsUdp) ElementType() reflect.Type {
	return reflect.TypeOf((**InputsUdp)(nil)).Elem()
}

func (i *InputsUdp) ToInputsUdpOutput() InputsUdpOutput {
	return i.ToInputsUdpOutputWithContext(context.Background())
}

func (i *InputsUdp) ToInputsUdpOutputWithContext(ctx context.Context) InputsUdpOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InputsUdpOutput)
}

func (i *InputsUdp) ToOutput(ctx context.Context) pulumix.Output[*InputsUdp] {
	return pulumix.Output[*InputsUdp]{
		OutputState: i.ToInputsUdpOutputWithContext(ctx).OutputState,
	}
}

// InputsUdpArrayInput is an input type that accepts InputsUdpArray and InputsUdpArrayOutput values.
// You can construct a concrete instance of `InputsUdpArrayInput` via:
//
//	InputsUdpArray{ InputsUdpArgs{...} }
type InputsUdpArrayInput interface {
	pulumi.Input

	ToInputsUdpArrayOutput() InputsUdpArrayOutput
	ToInputsUdpArrayOutputWithContext(context.Context) InputsUdpArrayOutput
}

type InputsUdpArray []InputsUdpInput

func (InputsUdpArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*InputsUdp)(nil)).Elem()
}

func (i InputsUdpArray) ToInputsUdpArrayOutput() InputsUdpArrayOutput {
	return i.ToInputsUdpArrayOutputWithContext(context.Background())
}

func (i InputsUdpArray) ToInputsUdpArrayOutputWithContext(ctx context.Context) InputsUdpArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InputsUdpArrayOutput)
}

func (i InputsUdpArray) ToOutput(ctx context.Context) pulumix.Output[[]*InputsUdp] {
	return pulumix.Output[[]*InputsUdp]{
		OutputState: i.ToInputsUdpArrayOutputWithContext(ctx).OutputState,
	}
}

// InputsUdpMapInput is an input type that accepts InputsUdpMap and InputsUdpMapOutput values.
// You can construct a concrete instance of `InputsUdpMapInput` via:
//
//	InputsUdpMap{ "key": InputsUdpArgs{...} }
type InputsUdpMapInput interface {
	pulumi.Input

	ToInputsUdpMapOutput() InputsUdpMapOutput
	ToInputsUdpMapOutputWithContext(context.Context) InputsUdpMapOutput
}

type InputsUdpMap map[string]InputsUdpInput

func (InputsUdpMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*InputsUdp)(nil)).Elem()
}

func (i InputsUdpMap) ToInputsUdpMapOutput() InputsUdpMapOutput {
	return i.ToInputsUdpMapOutputWithContext(context.Background())
}

func (i InputsUdpMap) ToInputsUdpMapOutputWithContext(ctx context.Context) InputsUdpMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InputsUdpMapOutput)
}

func (i InputsUdpMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*InputsUdp] {
	return pulumix.Output[map[string]*InputsUdp]{
		OutputState: i.ToInputsUdpMapOutputWithContext(ctx).OutputState,
	}
}

type InputsUdpOutput struct{ *pulumi.OutputState }

func (InputsUdpOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**InputsUdp)(nil)).Elem()
}

func (o InputsUdpOutput) ToInputsUdpOutput() InputsUdpOutput {
	return o
}

func (o InputsUdpOutput) ToInputsUdpOutputWithContext(ctx context.Context) InputsUdpOutput {
	return o
}

func (o InputsUdpOutput) ToOutput(ctx context.Context) pulumix.Output[*InputsUdp] {
	return pulumix.Output[*InputsUdp]{
		OutputState: o.OutputState,
	}
}

// The app/user context that is the namespace for the resource
func (o InputsUdpOutput) Acl() InputsUdpAclOutput {
	return o.ApplyT(func(v *InputsUdp) InputsUdpAclOutput { return v.Acl }).(InputsUdpAclOutput)
}

// Valid values: (ip | dns | none)
// Set the host for the remote server that is sending data.
// ip sets the host to the IP address of the remote server sending data.
// dns sets the host to the reverse DNS entry for the IP address of the remote server sending data.
// none leaves the host as specified in inputs.conf, which is typically the Splunk system hostname.
// Default value is dns.
func (o InputsUdpOutput) ConnectionHost() pulumi.StringOutput {
	return o.ApplyT(func(v *InputsUdp) pulumi.StringOutput { return v.ConnectionHost }).(pulumi.StringOutput)
}

// Indicates if input is disabled.
func (o InputsUdpOutput) Disabled() pulumi.BoolOutput {
	return o.ApplyT(func(v *InputsUdp) pulumi.BoolOutput { return v.Disabled }).(pulumi.BoolOutput)
}

// The value to populate in the host field for incoming events. This is used during parsing/indexing, in particular to set the host field. It is also the host field used at search time.
func (o InputsUdpOutput) Host() pulumi.StringOutput {
	return o.ApplyT(func(v *InputsUdp) pulumi.StringOutput { return v.Host }).(pulumi.StringOutput)
}

// Which index events from this input should be stored in. Defaults to default.
func (o InputsUdpOutput) Index() pulumi.StringOutput {
	return o.ApplyT(func(v *InputsUdp) pulumi.StringOutput { return v.Index }).(pulumi.StringOutput)
}

// The UDP port that this input should listen on.
func (o InputsUdpOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *InputsUdp) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// If set to true, prevents Splunk software from prepending a timestamp and hostname to incoming events.
func (o InputsUdpOutput) NoAppendingTimestamp() pulumi.BoolOutput {
	return o.ApplyT(func(v *InputsUdp) pulumi.BoolOutput { return v.NoAppendingTimestamp }).(pulumi.BoolOutput)
}

// If set to true, Splunk software does not remove the priority field from incoming syslog events.
func (o InputsUdpOutput) NoPriorityStripping() pulumi.BoolOutput {
	return o.ApplyT(func(v *InputsUdp) pulumi.BoolOutput { return v.NoPriorityStripping }).(pulumi.BoolOutput)
}

// Which queue events from this input should be sent to. Generally this does not need to be changed.
func (o InputsUdpOutput) Queue() pulumi.StringOutput {
	return o.ApplyT(func(v *InputsUdp) pulumi.StringOutput { return v.Queue }).(pulumi.StringOutput)
}

// Restrict incoming connections on this port to the host specified here.
// If this is not set, the value specified in [udp://<remote server>:<port>] in inputs.conf is used.
func (o InputsUdpOutput) RestrictToHost() pulumi.StringOutput {
	return o.ApplyT(func(v *InputsUdp) pulumi.StringOutput { return v.RestrictToHost }).(pulumi.StringOutput)
}

// The value to populate in the source field for incoming events. The same source should not be used for multiple data inputs.
func (o InputsUdpOutput) Source() pulumi.StringOutput {
	return o.ApplyT(func(v *InputsUdp) pulumi.StringOutput { return v.Source }).(pulumi.StringOutput)
}

// The value to populate in the sourcetype field for incoming events.
func (o InputsUdpOutput) Sourcetype() pulumi.StringOutput {
	return o.ApplyT(func(v *InputsUdp) pulumi.StringOutput { return v.Sourcetype }).(pulumi.StringOutput)
}

type InputsUdpArrayOutput struct{ *pulumi.OutputState }

func (InputsUdpArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*InputsUdp)(nil)).Elem()
}

func (o InputsUdpArrayOutput) ToInputsUdpArrayOutput() InputsUdpArrayOutput {
	return o
}

func (o InputsUdpArrayOutput) ToInputsUdpArrayOutputWithContext(ctx context.Context) InputsUdpArrayOutput {
	return o
}

func (o InputsUdpArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*InputsUdp] {
	return pulumix.Output[[]*InputsUdp]{
		OutputState: o.OutputState,
	}
}

func (o InputsUdpArrayOutput) Index(i pulumi.IntInput) InputsUdpOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *InputsUdp {
		return vs[0].([]*InputsUdp)[vs[1].(int)]
	}).(InputsUdpOutput)
}

type InputsUdpMapOutput struct{ *pulumi.OutputState }

func (InputsUdpMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*InputsUdp)(nil)).Elem()
}

func (o InputsUdpMapOutput) ToInputsUdpMapOutput() InputsUdpMapOutput {
	return o
}

func (o InputsUdpMapOutput) ToInputsUdpMapOutputWithContext(ctx context.Context) InputsUdpMapOutput {
	return o
}

func (o InputsUdpMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*InputsUdp] {
	return pulumix.Output[map[string]*InputsUdp]{
		OutputState: o.OutputState,
	}
}

func (o InputsUdpMapOutput) MapIndex(k pulumi.StringInput) InputsUdpOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *InputsUdp {
		return vs[0].(map[string]*InputsUdp)[vs[1].(string)]
	}).(InputsUdpOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*InputsUdpInput)(nil)).Elem(), &InputsUdp{})
	pulumi.RegisterInputType(reflect.TypeOf((*InputsUdpArrayInput)(nil)).Elem(), InputsUdpArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*InputsUdpMapInput)(nil)).Elem(), InputsUdpMap{})
	pulumi.RegisterOutputType(InputsUdpOutput{})
	pulumi.RegisterOutputType(InputsUdpArrayOutput{})
	pulumi.RegisterOutputType(InputsUdpMapOutput{})
}
