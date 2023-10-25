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

// ## # Resource: GlobalHttpEventCollector
//
// Update Global HTTP Event Collector input configuration.
type GlobalHttpEventCollector struct {
	pulumi.CustomResourceState

	// Number of threads used by HTTP Input server.
	DedicatedIoThreads pulumi.IntOutput `pulumi:"dedicatedIoThreads"`
	// Input disabled indicator.
	Disabled pulumi.BoolOutput `pulumi:"disabled"`
	// Enable SSL protocol for HTTP data input. `true` = SSL enabled, `false` = SSL disabled.
	EnableSsl pulumi.BoolOutput `pulumi:"enableSsl"`
	// Maximum number of simultaneous HTTP connections accepted. Adjusting this value may cause server performance issues and is not generally recommended. Possible values for this setting vary by OS.
	MaxSockets pulumi.IntOutput `pulumi:"maxSockets"`
	// Maximum number of threads that can be used by active HTTP transactions. Adjusting this value may cause server performance issues and is not generally recommended. Possible values for this setting vary by OS.
	MaxThreads pulumi.IntOutput `pulumi:"maxThreads"`
	// HTTP data input IP port.
	Port pulumi.IntOutput `pulumi:"port"`
	// Indicates whether the event collector input writes its configuration to a deployment server repository. When this setting is set to 1 (enabled), the input writes its configuration to the directory specified as repositoryLocation in serverclass.conf.
	// Copy the full contents of the splunkHttpinput app directory to this directory for the configuration to work. When enabled, only the tokens defined in the splunkHttpinput app in this repository are viewable and editable on the API and the Data Inputs page in Splunk Web. When disabled, the input writes its configuration to $SPLUNK_HOME/etc/apps by default. Defaults to 0 (disabled).
	UseDeploymentServer pulumi.IntOutput `pulumi:"useDeploymentServer"`
}

// NewGlobalHttpEventCollector registers a new resource with the given unique name, arguments, and options.
func NewGlobalHttpEventCollector(ctx *pulumi.Context,
	name string, args *GlobalHttpEventCollectorArgs, opts ...pulumi.ResourceOption) (*GlobalHttpEventCollector, error) {
	if args == nil {
		args = &GlobalHttpEventCollectorArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource GlobalHttpEventCollector
	err := ctx.RegisterResource("splunk:index/globalHttpEventCollector:GlobalHttpEventCollector", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetGlobalHttpEventCollector gets an existing GlobalHttpEventCollector resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetGlobalHttpEventCollector(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GlobalHttpEventCollectorState, opts ...pulumi.ResourceOption) (*GlobalHttpEventCollector, error) {
	var resource GlobalHttpEventCollector
	err := ctx.ReadResource("splunk:index/globalHttpEventCollector:GlobalHttpEventCollector", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering GlobalHttpEventCollector resources.
type globalHttpEventCollectorState struct {
	// Number of threads used by HTTP Input server.
	DedicatedIoThreads *int `pulumi:"dedicatedIoThreads"`
	// Input disabled indicator.
	Disabled *bool `pulumi:"disabled"`
	// Enable SSL protocol for HTTP data input. `true` = SSL enabled, `false` = SSL disabled.
	EnableSsl *bool `pulumi:"enableSsl"`
	// Maximum number of simultaneous HTTP connections accepted. Adjusting this value may cause server performance issues and is not generally recommended. Possible values for this setting vary by OS.
	MaxSockets *int `pulumi:"maxSockets"`
	// Maximum number of threads that can be used by active HTTP transactions. Adjusting this value may cause server performance issues and is not generally recommended. Possible values for this setting vary by OS.
	MaxThreads *int `pulumi:"maxThreads"`
	// HTTP data input IP port.
	Port *int `pulumi:"port"`
	// Indicates whether the event collector input writes its configuration to a deployment server repository. When this setting is set to 1 (enabled), the input writes its configuration to the directory specified as repositoryLocation in serverclass.conf.
	// Copy the full contents of the splunkHttpinput app directory to this directory for the configuration to work. When enabled, only the tokens defined in the splunkHttpinput app in this repository are viewable and editable on the API and the Data Inputs page in Splunk Web. When disabled, the input writes its configuration to $SPLUNK_HOME/etc/apps by default. Defaults to 0 (disabled).
	UseDeploymentServer *int `pulumi:"useDeploymentServer"`
}

type GlobalHttpEventCollectorState struct {
	// Number of threads used by HTTP Input server.
	DedicatedIoThreads pulumi.IntPtrInput
	// Input disabled indicator.
	Disabled pulumi.BoolPtrInput
	// Enable SSL protocol for HTTP data input. `true` = SSL enabled, `false` = SSL disabled.
	EnableSsl pulumi.BoolPtrInput
	// Maximum number of simultaneous HTTP connections accepted. Adjusting this value may cause server performance issues and is not generally recommended. Possible values for this setting vary by OS.
	MaxSockets pulumi.IntPtrInput
	// Maximum number of threads that can be used by active HTTP transactions. Adjusting this value may cause server performance issues and is not generally recommended. Possible values for this setting vary by OS.
	MaxThreads pulumi.IntPtrInput
	// HTTP data input IP port.
	Port pulumi.IntPtrInput
	// Indicates whether the event collector input writes its configuration to a deployment server repository. When this setting is set to 1 (enabled), the input writes its configuration to the directory specified as repositoryLocation in serverclass.conf.
	// Copy the full contents of the splunkHttpinput app directory to this directory for the configuration to work. When enabled, only the tokens defined in the splunkHttpinput app in this repository are viewable and editable on the API and the Data Inputs page in Splunk Web. When disabled, the input writes its configuration to $SPLUNK_HOME/etc/apps by default. Defaults to 0 (disabled).
	UseDeploymentServer pulumi.IntPtrInput
}

func (GlobalHttpEventCollectorState) ElementType() reflect.Type {
	return reflect.TypeOf((*globalHttpEventCollectorState)(nil)).Elem()
}

type globalHttpEventCollectorArgs struct {
	// Number of threads used by HTTP Input server.
	DedicatedIoThreads *int `pulumi:"dedicatedIoThreads"`
	// Input disabled indicator.
	Disabled *bool `pulumi:"disabled"`
	// Enable SSL protocol for HTTP data input. `true` = SSL enabled, `false` = SSL disabled.
	EnableSsl *bool `pulumi:"enableSsl"`
	// Maximum number of simultaneous HTTP connections accepted. Adjusting this value may cause server performance issues and is not generally recommended. Possible values for this setting vary by OS.
	MaxSockets *int `pulumi:"maxSockets"`
	// Maximum number of threads that can be used by active HTTP transactions. Adjusting this value may cause server performance issues and is not generally recommended. Possible values for this setting vary by OS.
	MaxThreads *int `pulumi:"maxThreads"`
	// HTTP data input IP port.
	Port *int `pulumi:"port"`
	// Indicates whether the event collector input writes its configuration to a deployment server repository. When this setting is set to 1 (enabled), the input writes its configuration to the directory specified as repositoryLocation in serverclass.conf.
	// Copy the full contents of the splunkHttpinput app directory to this directory for the configuration to work. When enabled, only the tokens defined in the splunkHttpinput app in this repository are viewable and editable on the API and the Data Inputs page in Splunk Web. When disabled, the input writes its configuration to $SPLUNK_HOME/etc/apps by default. Defaults to 0 (disabled).
	UseDeploymentServer *int `pulumi:"useDeploymentServer"`
}

// The set of arguments for constructing a GlobalHttpEventCollector resource.
type GlobalHttpEventCollectorArgs struct {
	// Number of threads used by HTTP Input server.
	DedicatedIoThreads pulumi.IntPtrInput
	// Input disabled indicator.
	Disabled pulumi.BoolPtrInput
	// Enable SSL protocol for HTTP data input. `true` = SSL enabled, `false` = SSL disabled.
	EnableSsl pulumi.BoolPtrInput
	// Maximum number of simultaneous HTTP connections accepted. Adjusting this value may cause server performance issues and is not generally recommended. Possible values for this setting vary by OS.
	MaxSockets pulumi.IntPtrInput
	// Maximum number of threads that can be used by active HTTP transactions. Adjusting this value may cause server performance issues and is not generally recommended. Possible values for this setting vary by OS.
	MaxThreads pulumi.IntPtrInput
	// HTTP data input IP port.
	Port pulumi.IntPtrInput
	// Indicates whether the event collector input writes its configuration to a deployment server repository. When this setting is set to 1 (enabled), the input writes its configuration to the directory specified as repositoryLocation in serverclass.conf.
	// Copy the full contents of the splunkHttpinput app directory to this directory for the configuration to work. When enabled, only the tokens defined in the splunkHttpinput app in this repository are viewable and editable on the API and the Data Inputs page in Splunk Web. When disabled, the input writes its configuration to $SPLUNK_HOME/etc/apps by default. Defaults to 0 (disabled).
	UseDeploymentServer pulumi.IntPtrInput
}

func (GlobalHttpEventCollectorArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*globalHttpEventCollectorArgs)(nil)).Elem()
}

type GlobalHttpEventCollectorInput interface {
	pulumi.Input

	ToGlobalHttpEventCollectorOutput() GlobalHttpEventCollectorOutput
	ToGlobalHttpEventCollectorOutputWithContext(ctx context.Context) GlobalHttpEventCollectorOutput
}

func (*GlobalHttpEventCollector) ElementType() reflect.Type {
	return reflect.TypeOf((**GlobalHttpEventCollector)(nil)).Elem()
}

func (i *GlobalHttpEventCollector) ToGlobalHttpEventCollectorOutput() GlobalHttpEventCollectorOutput {
	return i.ToGlobalHttpEventCollectorOutputWithContext(context.Background())
}

func (i *GlobalHttpEventCollector) ToGlobalHttpEventCollectorOutputWithContext(ctx context.Context) GlobalHttpEventCollectorOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GlobalHttpEventCollectorOutput)
}

func (i *GlobalHttpEventCollector) ToOutput(ctx context.Context) pulumix.Output[*GlobalHttpEventCollector] {
	return pulumix.Output[*GlobalHttpEventCollector]{
		OutputState: i.ToGlobalHttpEventCollectorOutputWithContext(ctx).OutputState,
	}
}

// GlobalHttpEventCollectorArrayInput is an input type that accepts GlobalHttpEventCollectorArray and GlobalHttpEventCollectorArrayOutput values.
// You can construct a concrete instance of `GlobalHttpEventCollectorArrayInput` via:
//
//	GlobalHttpEventCollectorArray{ GlobalHttpEventCollectorArgs{...} }
type GlobalHttpEventCollectorArrayInput interface {
	pulumi.Input

	ToGlobalHttpEventCollectorArrayOutput() GlobalHttpEventCollectorArrayOutput
	ToGlobalHttpEventCollectorArrayOutputWithContext(context.Context) GlobalHttpEventCollectorArrayOutput
}

type GlobalHttpEventCollectorArray []GlobalHttpEventCollectorInput

func (GlobalHttpEventCollectorArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*GlobalHttpEventCollector)(nil)).Elem()
}

func (i GlobalHttpEventCollectorArray) ToGlobalHttpEventCollectorArrayOutput() GlobalHttpEventCollectorArrayOutput {
	return i.ToGlobalHttpEventCollectorArrayOutputWithContext(context.Background())
}

func (i GlobalHttpEventCollectorArray) ToGlobalHttpEventCollectorArrayOutputWithContext(ctx context.Context) GlobalHttpEventCollectorArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GlobalHttpEventCollectorArrayOutput)
}

func (i GlobalHttpEventCollectorArray) ToOutput(ctx context.Context) pulumix.Output[[]*GlobalHttpEventCollector] {
	return pulumix.Output[[]*GlobalHttpEventCollector]{
		OutputState: i.ToGlobalHttpEventCollectorArrayOutputWithContext(ctx).OutputState,
	}
}

// GlobalHttpEventCollectorMapInput is an input type that accepts GlobalHttpEventCollectorMap and GlobalHttpEventCollectorMapOutput values.
// You can construct a concrete instance of `GlobalHttpEventCollectorMapInput` via:
//
//	GlobalHttpEventCollectorMap{ "key": GlobalHttpEventCollectorArgs{...} }
type GlobalHttpEventCollectorMapInput interface {
	pulumi.Input

	ToGlobalHttpEventCollectorMapOutput() GlobalHttpEventCollectorMapOutput
	ToGlobalHttpEventCollectorMapOutputWithContext(context.Context) GlobalHttpEventCollectorMapOutput
}

type GlobalHttpEventCollectorMap map[string]GlobalHttpEventCollectorInput

func (GlobalHttpEventCollectorMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*GlobalHttpEventCollector)(nil)).Elem()
}

func (i GlobalHttpEventCollectorMap) ToGlobalHttpEventCollectorMapOutput() GlobalHttpEventCollectorMapOutput {
	return i.ToGlobalHttpEventCollectorMapOutputWithContext(context.Background())
}

func (i GlobalHttpEventCollectorMap) ToGlobalHttpEventCollectorMapOutputWithContext(ctx context.Context) GlobalHttpEventCollectorMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GlobalHttpEventCollectorMapOutput)
}

func (i GlobalHttpEventCollectorMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*GlobalHttpEventCollector] {
	return pulumix.Output[map[string]*GlobalHttpEventCollector]{
		OutputState: i.ToGlobalHttpEventCollectorMapOutputWithContext(ctx).OutputState,
	}
}

type GlobalHttpEventCollectorOutput struct{ *pulumi.OutputState }

func (GlobalHttpEventCollectorOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GlobalHttpEventCollector)(nil)).Elem()
}

func (o GlobalHttpEventCollectorOutput) ToGlobalHttpEventCollectorOutput() GlobalHttpEventCollectorOutput {
	return o
}

func (o GlobalHttpEventCollectorOutput) ToGlobalHttpEventCollectorOutputWithContext(ctx context.Context) GlobalHttpEventCollectorOutput {
	return o
}

func (o GlobalHttpEventCollectorOutput) ToOutput(ctx context.Context) pulumix.Output[*GlobalHttpEventCollector] {
	return pulumix.Output[*GlobalHttpEventCollector]{
		OutputState: o.OutputState,
	}
}

// Number of threads used by HTTP Input server.
func (o GlobalHttpEventCollectorOutput) DedicatedIoThreads() pulumi.IntOutput {
	return o.ApplyT(func(v *GlobalHttpEventCollector) pulumi.IntOutput { return v.DedicatedIoThreads }).(pulumi.IntOutput)
}

// Input disabled indicator.
func (o GlobalHttpEventCollectorOutput) Disabled() pulumi.BoolOutput {
	return o.ApplyT(func(v *GlobalHttpEventCollector) pulumi.BoolOutput { return v.Disabled }).(pulumi.BoolOutput)
}

// Enable SSL protocol for HTTP data input. `true` = SSL enabled, `false` = SSL disabled.
func (o GlobalHttpEventCollectorOutput) EnableSsl() pulumi.BoolOutput {
	return o.ApplyT(func(v *GlobalHttpEventCollector) pulumi.BoolOutput { return v.EnableSsl }).(pulumi.BoolOutput)
}

// Maximum number of simultaneous HTTP connections accepted. Adjusting this value may cause server performance issues and is not generally recommended. Possible values for this setting vary by OS.
func (o GlobalHttpEventCollectorOutput) MaxSockets() pulumi.IntOutput {
	return o.ApplyT(func(v *GlobalHttpEventCollector) pulumi.IntOutput { return v.MaxSockets }).(pulumi.IntOutput)
}

// Maximum number of threads that can be used by active HTTP transactions. Adjusting this value may cause server performance issues and is not generally recommended. Possible values for this setting vary by OS.
func (o GlobalHttpEventCollectorOutput) MaxThreads() pulumi.IntOutput {
	return o.ApplyT(func(v *GlobalHttpEventCollector) pulumi.IntOutput { return v.MaxThreads }).(pulumi.IntOutput)
}

// HTTP data input IP port.
func (o GlobalHttpEventCollectorOutput) Port() pulumi.IntOutput {
	return o.ApplyT(func(v *GlobalHttpEventCollector) pulumi.IntOutput { return v.Port }).(pulumi.IntOutput)
}

// Indicates whether the event collector input writes its configuration to a deployment server repository. When this setting is set to 1 (enabled), the input writes its configuration to the directory specified as repositoryLocation in serverclass.conf.
// Copy the full contents of the splunkHttpinput app directory to this directory for the configuration to work. When enabled, only the tokens defined in the splunkHttpinput app in this repository are viewable and editable on the API and the Data Inputs page in Splunk Web. When disabled, the input writes its configuration to $SPLUNK_HOME/etc/apps by default. Defaults to 0 (disabled).
func (o GlobalHttpEventCollectorOutput) UseDeploymentServer() pulumi.IntOutput {
	return o.ApplyT(func(v *GlobalHttpEventCollector) pulumi.IntOutput { return v.UseDeploymentServer }).(pulumi.IntOutput)
}

type GlobalHttpEventCollectorArrayOutput struct{ *pulumi.OutputState }

func (GlobalHttpEventCollectorArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*GlobalHttpEventCollector)(nil)).Elem()
}

func (o GlobalHttpEventCollectorArrayOutput) ToGlobalHttpEventCollectorArrayOutput() GlobalHttpEventCollectorArrayOutput {
	return o
}

func (o GlobalHttpEventCollectorArrayOutput) ToGlobalHttpEventCollectorArrayOutputWithContext(ctx context.Context) GlobalHttpEventCollectorArrayOutput {
	return o
}

func (o GlobalHttpEventCollectorArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*GlobalHttpEventCollector] {
	return pulumix.Output[[]*GlobalHttpEventCollector]{
		OutputState: o.OutputState,
	}
}

func (o GlobalHttpEventCollectorArrayOutput) Index(i pulumi.IntInput) GlobalHttpEventCollectorOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *GlobalHttpEventCollector {
		return vs[0].([]*GlobalHttpEventCollector)[vs[1].(int)]
	}).(GlobalHttpEventCollectorOutput)
}

type GlobalHttpEventCollectorMapOutput struct{ *pulumi.OutputState }

func (GlobalHttpEventCollectorMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*GlobalHttpEventCollector)(nil)).Elem()
}

func (o GlobalHttpEventCollectorMapOutput) ToGlobalHttpEventCollectorMapOutput() GlobalHttpEventCollectorMapOutput {
	return o
}

func (o GlobalHttpEventCollectorMapOutput) ToGlobalHttpEventCollectorMapOutputWithContext(ctx context.Context) GlobalHttpEventCollectorMapOutput {
	return o
}

func (o GlobalHttpEventCollectorMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*GlobalHttpEventCollector] {
	return pulumix.Output[map[string]*GlobalHttpEventCollector]{
		OutputState: o.OutputState,
	}
}

func (o GlobalHttpEventCollectorMapOutput) MapIndex(k pulumi.StringInput) GlobalHttpEventCollectorOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *GlobalHttpEventCollector {
		return vs[0].(map[string]*GlobalHttpEventCollector)[vs[1].(string)]
	}).(GlobalHttpEventCollectorOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*GlobalHttpEventCollectorInput)(nil)).Elem(), &GlobalHttpEventCollector{})
	pulumi.RegisterInputType(reflect.TypeOf((*GlobalHttpEventCollectorArrayInput)(nil)).Elem(), GlobalHttpEventCollectorArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*GlobalHttpEventCollectorMapInput)(nil)).Elem(), GlobalHttpEventCollectorMap{})
	pulumi.RegisterOutputType(GlobalHttpEventCollectorOutput{})
	pulumi.RegisterOutputType(GlobalHttpEventCollectorArrayOutput{})
	pulumi.RegisterOutputType(GlobalHttpEventCollectorMapOutput{})
}
