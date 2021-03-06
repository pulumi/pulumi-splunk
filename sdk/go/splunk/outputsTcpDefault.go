// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package splunk

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## # Resource: OutputsTcpDefault
//
// Manage to global tcpout properties.
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
// 		_, err := splunk.NewOutputsTcpDefault(ctx, "tcpDefault", &splunk.OutputsTcpDefaultArgs{
// 			DefaultGroup:          pulumi.String("test-indexers"),
// 			Disabled:              pulumi.Bool(false),
// 			DropEventsOnQueueFull: pulumi.Int(60),
// 			IndexAndForward:       pulumi.Bool(true),
// 			MaxQueueSize:          pulumi.String("100KB"),
// 			SendCookedData:        pulumi.Bool(true),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type OutputsTcpDefault struct {
	pulumi.CustomResourceState

	// The app/user context that is the namespace for the resource
	Acl OutputsTcpDefaultAclOutput `pulumi:"acl"`
	// Comma-separated list of one or more target group names, specified later in [tcpout:<target_group>] stanzas of outputs.conf.spec file.
	// The forwarder sends all data to the specified groups. If you do not want to forward data automatically, do not set this attribute. Can be overridden by an inputs.conf _TCP_ROUTING setting, which in turn can be overridden by a props.conf/transforms.conf modifier.
	DefaultGroup pulumi.StringOutput `pulumi:"defaultGroup"`
	// Disables default tcpout settings
	Disabled pulumi.BoolOutput `pulumi:"disabled"`
	// If set to a positive number, wait the specified number of seconds before throwing out all new events until the output queue has space. Defaults to -1 (do not drop events).
	// <br>CAUTION: Do not set this value to a positive integer if you are monitoring files.
	// Setting this to -1 or 0 causes the output queue to block when it gets full, which causes further blocking up the processing chain. If any target group queue is blocked, no more data reaches any other target group.
	// Using auto load-balancing is the best way to minimize this condition, because, in that case, multiple receivers must be down (or jammed up) before queue blocking can occur.
	DropEventsOnQueueFull pulumi.IntOutput `pulumi:"dropEventsOnQueueFull"`
	// How often (in seconds) to send a heartbeat packet to the receiving server.
	// Heartbeats are only sent if sendCookedData=true. Defaults to 30 seconds.
	HeartbeatFrequency pulumi.IntOutput `pulumi:"heartbeatFrequency"`
	// Specifies whether to index all data locally, in addition to forwarding it. Defaults to false.
	// This is known as an "index-and-forward" configuration. This attribute is only available for heavy forwarders. It is available only at the top level [tcpout] stanza in outputs.conf. It cannot be overridden in a target group.
	IndexAndForward pulumi.BoolOutput `pulumi:"indexAndForward"`
	// Specify an integer or integer[KB|MB|GB].
	// <br>Sets the maximum size of the forwarder output queue. It also sets the maximum size of the wait queue to 3x this value, if you have enabled indexer acknowledgment (useACK=true).
	// Although the wait queue and the output queues are both configured by this attribute, they are separate queues. The setting determines the maximum size of the queue in-memory (RAM) buffer.
	// For heavy forwarders sending parsed data, maxQueueSize is the maximum number of events. Since events are typically much shorter than data blocks, the memory consumed by the queue on a parsing forwarder is likely to be much smaller than on a non-parsing forwarder, if you use this version of the setting.
	// If specified as a lone integer (for example, maxQueueSize=100), maxQueueSize indicates the maximum number of queued events (for parsed data) or blocks of data (for unparsed data). A block of data is approximately 64KB. For non-parsing forwarders, such as universal forwarders, that send unparsed data, maxQueueSize is the maximum number of data blocks.
	// If specified as an integer followed by KB, MB, or GB (for example, maxQueueSize=100MB), maxQueueSize indicates the maximum RAM allocated to the queue buffer. Defaults to 500KB (which means a maximum size of 500KB for the output queue and 1500KB for the wait queue, if any).
	MaxQueueSize pulumi.StringOutput `pulumi:"maxQueueSize"`
	// Configuration to be edited. The only valid value is "tcpout".
	Name pulumi.StringOutput `pulumi:"name"`
	// If true, events are cooked (processed by Splunk software). If false, events are raw and untouched prior to sending. Defaults to true.
	// Set to false if you are sending to a third-party system.
	SendCookedData pulumi.BoolOutput `pulumi:"sendCookedData"`
}

// NewOutputsTcpDefault registers a new resource with the given unique name, arguments, and options.
func NewOutputsTcpDefault(ctx *pulumi.Context,
	name string, args *OutputsTcpDefaultArgs, opts ...pulumi.ResourceOption) (*OutputsTcpDefault, error) {
	if args == nil {
		args = &OutputsTcpDefaultArgs{}
	}

	var resource OutputsTcpDefault
	err := ctx.RegisterResource("splunk:index/outputsTcpDefault:OutputsTcpDefault", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetOutputsTcpDefault gets an existing OutputsTcpDefault resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetOutputsTcpDefault(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *OutputsTcpDefaultState, opts ...pulumi.ResourceOption) (*OutputsTcpDefault, error) {
	var resource OutputsTcpDefault
	err := ctx.ReadResource("splunk:index/outputsTcpDefault:OutputsTcpDefault", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering OutputsTcpDefault resources.
type outputsTcpDefaultState struct {
	// The app/user context that is the namespace for the resource
	Acl *OutputsTcpDefaultAcl `pulumi:"acl"`
	// Comma-separated list of one or more target group names, specified later in [tcpout:<target_group>] stanzas of outputs.conf.spec file.
	// The forwarder sends all data to the specified groups. If you do not want to forward data automatically, do not set this attribute. Can be overridden by an inputs.conf _TCP_ROUTING setting, which in turn can be overridden by a props.conf/transforms.conf modifier.
	DefaultGroup *string `pulumi:"defaultGroup"`
	// Disables default tcpout settings
	Disabled *bool `pulumi:"disabled"`
	// If set to a positive number, wait the specified number of seconds before throwing out all new events until the output queue has space. Defaults to -1 (do not drop events).
	// <br>CAUTION: Do not set this value to a positive integer if you are monitoring files.
	// Setting this to -1 or 0 causes the output queue to block when it gets full, which causes further blocking up the processing chain. If any target group queue is blocked, no more data reaches any other target group.
	// Using auto load-balancing is the best way to minimize this condition, because, in that case, multiple receivers must be down (or jammed up) before queue blocking can occur.
	DropEventsOnQueueFull *int `pulumi:"dropEventsOnQueueFull"`
	// How often (in seconds) to send a heartbeat packet to the receiving server.
	// Heartbeats are only sent if sendCookedData=true. Defaults to 30 seconds.
	HeartbeatFrequency *int `pulumi:"heartbeatFrequency"`
	// Specifies whether to index all data locally, in addition to forwarding it. Defaults to false.
	// This is known as an "index-and-forward" configuration. This attribute is only available for heavy forwarders. It is available only at the top level [tcpout] stanza in outputs.conf. It cannot be overridden in a target group.
	IndexAndForward *bool `pulumi:"indexAndForward"`
	// Specify an integer or integer[KB|MB|GB].
	// <br>Sets the maximum size of the forwarder output queue. It also sets the maximum size of the wait queue to 3x this value, if you have enabled indexer acknowledgment (useACK=true).
	// Although the wait queue and the output queues are both configured by this attribute, they are separate queues. The setting determines the maximum size of the queue in-memory (RAM) buffer.
	// For heavy forwarders sending parsed data, maxQueueSize is the maximum number of events. Since events are typically much shorter than data blocks, the memory consumed by the queue on a parsing forwarder is likely to be much smaller than on a non-parsing forwarder, if you use this version of the setting.
	// If specified as a lone integer (for example, maxQueueSize=100), maxQueueSize indicates the maximum number of queued events (for parsed data) or blocks of data (for unparsed data). A block of data is approximately 64KB. For non-parsing forwarders, such as universal forwarders, that send unparsed data, maxQueueSize is the maximum number of data blocks.
	// If specified as an integer followed by KB, MB, or GB (for example, maxQueueSize=100MB), maxQueueSize indicates the maximum RAM allocated to the queue buffer. Defaults to 500KB (which means a maximum size of 500KB for the output queue and 1500KB for the wait queue, if any).
	MaxQueueSize *string `pulumi:"maxQueueSize"`
	// Configuration to be edited. The only valid value is "tcpout".
	Name *string `pulumi:"name"`
	// If true, events are cooked (processed by Splunk software). If false, events are raw and untouched prior to sending. Defaults to true.
	// Set to false if you are sending to a third-party system.
	SendCookedData *bool `pulumi:"sendCookedData"`
}

type OutputsTcpDefaultState struct {
	// The app/user context that is the namespace for the resource
	Acl OutputsTcpDefaultAclPtrInput
	// Comma-separated list of one or more target group names, specified later in [tcpout:<target_group>] stanzas of outputs.conf.spec file.
	// The forwarder sends all data to the specified groups. If you do not want to forward data automatically, do not set this attribute. Can be overridden by an inputs.conf _TCP_ROUTING setting, which in turn can be overridden by a props.conf/transforms.conf modifier.
	DefaultGroup pulumi.StringPtrInput
	// Disables default tcpout settings
	Disabled pulumi.BoolPtrInput
	// If set to a positive number, wait the specified number of seconds before throwing out all new events until the output queue has space. Defaults to -1 (do not drop events).
	// <br>CAUTION: Do not set this value to a positive integer if you are monitoring files.
	// Setting this to -1 or 0 causes the output queue to block when it gets full, which causes further blocking up the processing chain. If any target group queue is blocked, no more data reaches any other target group.
	// Using auto load-balancing is the best way to minimize this condition, because, in that case, multiple receivers must be down (or jammed up) before queue blocking can occur.
	DropEventsOnQueueFull pulumi.IntPtrInput
	// How often (in seconds) to send a heartbeat packet to the receiving server.
	// Heartbeats are only sent if sendCookedData=true. Defaults to 30 seconds.
	HeartbeatFrequency pulumi.IntPtrInput
	// Specifies whether to index all data locally, in addition to forwarding it. Defaults to false.
	// This is known as an "index-and-forward" configuration. This attribute is only available for heavy forwarders. It is available only at the top level [tcpout] stanza in outputs.conf. It cannot be overridden in a target group.
	IndexAndForward pulumi.BoolPtrInput
	// Specify an integer or integer[KB|MB|GB].
	// <br>Sets the maximum size of the forwarder output queue. It also sets the maximum size of the wait queue to 3x this value, if you have enabled indexer acknowledgment (useACK=true).
	// Although the wait queue and the output queues are both configured by this attribute, they are separate queues. The setting determines the maximum size of the queue in-memory (RAM) buffer.
	// For heavy forwarders sending parsed data, maxQueueSize is the maximum number of events. Since events are typically much shorter than data blocks, the memory consumed by the queue on a parsing forwarder is likely to be much smaller than on a non-parsing forwarder, if you use this version of the setting.
	// If specified as a lone integer (for example, maxQueueSize=100), maxQueueSize indicates the maximum number of queued events (for parsed data) or blocks of data (for unparsed data). A block of data is approximately 64KB. For non-parsing forwarders, such as universal forwarders, that send unparsed data, maxQueueSize is the maximum number of data blocks.
	// If specified as an integer followed by KB, MB, or GB (for example, maxQueueSize=100MB), maxQueueSize indicates the maximum RAM allocated to the queue buffer. Defaults to 500KB (which means a maximum size of 500KB for the output queue and 1500KB for the wait queue, if any).
	MaxQueueSize pulumi.StringPtrInput
	// Configuration to be edited. The only valid value is "tcpout".
	Name pulumi.StringPtrInput
	// If true, events are cooked (processed by Splunk software). If false, events are raw and untouched prior to sending. Defaults to true.
	// Set to false if you are sending to a third-party system.
	SendCookedData pulumi.BoolPtrInput
}

func (OutputsTcpDefaultState) ElementType() reflect.Type {
	return reflect.TypeOf((*outputsTcpDefaultState)(nil)).Elem()
}

type outputsTcpDefaultArgs struct {
	// The app/user context that is the namespace for the resource
	Acl *OutputsTcpDefaultAcl `pulumi:"acl"`
	// Comma-separated list of one or more target group names, specified later in [tcpout:<target_group>] stanzas of outputs.conf.spec file.
	// The forwarder sends all data to the specified groups. If you do not want to forward data automatically, do not set this attribute. Can be overridden by an inputs.conf _TCP_ROUTING setting, which in turn can be overridden by a props.conf/transforms.conf modifier.
	DefaultGroup *string `pulumi:"defaultGroup"`
	// Disables default tcpout settings
	Disabled *bool `pulumi:"disabled"`
	// If set to a positive number, wait the specified number of seconds before throwing out all new events until the output queue has space. Defaults to -1 (do not drop events).
	// <br>CAUTION: Do not set this value to a positive integer if you are monitoring files.
	// Setting this to -1 or 0 causes the output queue to block when it gets full, which causes further blocking up the processing chain. If any target group queue is blocked, no more data reaches any other target group.
	// Using auto load-balancing is the best way to minimize this condition, because, in that case, multiple receivers must be down (or jammed up) before queue blocking can occur.
	DropEventsOnQueueFull *int `pulumi:"dropEventsOnQueueFull"`
	// How often (in seconds) to send a heartbeat packet to the receiving server.
	// Heartbeats are only sent if sendCookedData=true. Defaults to 30 seconds.
	HeartbeatFrequency *int `pulumi:"heartbeatFrequency"`
	// Specifies whether to index all data locally, in addition to forwarding it. Defaults to false.
	// This is known as an "index-and-forward" configuration. This attribute is only available for heavy forwarders. It is available only at the top level [tcpout] stanza in outputs.conf. It cannot be overridden in a target group.
	IndexAndForward *bool `pulumi:"indexAndForward"`
	// Specify an integer or integer[KB|MB|GB].
	// <br>Sets the maximum size of the forwarder output queue. It also sets the maximum size of the wait queue to 3x this value, if you have enabled indexer acknowledgment (useACK=true).
	// Although the wait queue and the output queues are both configured by this attribute, they are separate queues. The setting determines the maximum size of the queue in-memory (RAM) buffer.
	// For heavy forwarders sending parsed data, maxQueueSize is the maximum number of events. Since events are typically much shorter than data blocks, the memory consumed by the queue on a parsing forwarder is likely to be much smaller than on a non-parsing forwarder, if you use this version of the setting.
	// If specified as a lone integer (for example, maxQueueSize=100), maxQueueSize indicates the maximum number of queued events (for parsed data) or blocks of data (for unparsed data). A block of data is approximately 64KB. For non-parsing forwarders, such as universal forwarders, that send unparsed data, maxQueueSize is the maximum number of data blocks.
	// If specified as an integer followed by KB, MB, or GB (for example, maxQueueSize=100MB), maxQueueSize indicates the maximum RAM allocated to the queue buffer. Defaults to 500KB (which means a maximum size of 500KB for the output queue and 1500KB for the wait queue, if any).
	MaxQueueSize *string `pulumi:"maxQueueSize"`
	// Configuration to be edited. The only valid value is "tcpout".
	Name *string `pulumi:"name"`
	// If true, events are cooked (processed by Splunk software). If false, events are raw and untouched prior to sending. Defaults to true.
	// Set to false if you are sending to a third-party system.
	SendCookedData *bool `pulumi:"sendCookedData"`
}

// The set of arguments for constructing a OutputsTcpDefault resource.
type OutputsTcpDefaultArgs struct {
	// The app/user context that is the namespace for the resource
	Acl OutputsTcpDefaultAclPtrInput
	// Comma-separated list of one or more target group names, specified later in [tcpout:<target_group>] stanzas of outputs.conf.spec file.
	// The forwarder sends all data to the specified groups. If you do not want to forward data automatically, do not set this attribute. Can be overridden by an inputs.conf _TCP_ROUTING setting, which in turn can be overridden by a props.conf/transforms.conf modifier.
	DefaultGroup pulumi.StringPtrInput
	// Disables default tcpout settings
	Disabled pulumi.BoolPtrInput
	// If set to a positive number, wait the specified number of seconds before throwing out all new events until the output queue has space. Defaults to -1 (do not drop events).
	// <br>CAUTION: Do not set this value to a positive integer if you are monitoring files.
	// Setting this to -1 or 0 causes the output queue to block when it gets full, which causes further blocking up the processing chain. If any target group queue is blocked, no more data reaches any other target group.
	// Using auto load-balancing is the best way to minimize this condition, because, in that case, multiple receivers must be down (or jammed up) before queue blocking can occur.
	DropEventsOnQueueFull pulumi.IntPtrInput
	// How often (in seconds) to send a heartbeat packet to the receiving server.
	// Heartbeats are only sent if sendCookedData=true. Defaults to 30 seconds.
	HeartbeatFrequency pulumi.IntPtrInput
	// Specifies whether to index all data locally, in addition to forwarding it. Defaults to false.
	// This is known as an "index-and-forward" configuration. This attribute is only available for heavy forwarders. It is available only at the top level [tcpout] stanza in outputs.conf. It cannot be overridden in a target group.
	IndexAndForward pulumi.BoolPtrInput
	// Specify an integer or integer[KB|MB|GB].
	// <br>Sets the maximum size of the forwarder output queue. It also sets the maximum size of the wait queue to 3x this value, if you have enabled indexer acknowledgment (useACK=true).
	// Although the wait queue and the output queues are both configured by this attribute, they are separate queues. The setting determines the maximum size of the queue in-memory (RAM) buffer.
	// For heavy forwarders sending parsed data, maxQueueSize is the maximum number of events. Since events are typically much shorter than data blocks, the memory consumed by the queue on a parsing forwarder is likely to be much smaller than on a non-parsing forwarder, if you use this version of the setting.
	// If specified as a lone integer (for example, maxQueueSize=100), maxQueueSize indicates the maximum number of queued events (for parsed data) or blocks of data (for unparsed data). A block of data is approximately 64KB. For non-parsing forwarders, such as universal forwarders, that send unparsed data, maxQueueSize is the maximum number of data blocks.
	// If specified as an integer followed by KB, MB, or GB (for example, maxQueueSize=100MB), maxQueueSize indicates the maximum RAM allocated to the queue buffer. Defaults to 500KB (which means a maximum size of 500KB for the output queue and 1500KB for the wait queue, if any).
	MaxQueueSize pulumi.StringPtrInput
	// Configuration to be edited. The only valid value is "tcpout".
	Name pulumi.StringPtrInput
	// If true, events are cooked (processed by Splunk software). If false, events are raw and untouched prior to sending. Defaults to true.
	// Set to false if you are sending to a third-party system.
	SendCookedData pulumi.BoolPtrInput
}

func (OutputsTcpDefaultArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*outputsTcpDefaultArgs)(nil)).Elem()
}

type OutputsTcpDefaultInput interface {
	pulumi.Input

	ToOutputsTcpDefaultOutput() OutputsTcpDefaultOutput
	ToOutputsTcpDefaultOutputWithContext(ctx context.Context) OutputsTcpDefaultOutput
}

func (*OutputsTcpDefault) ElementType() reflect.Type {
	return reflect.TypeOf((*OutputsTcpDefault)(nil))
}

func (i *OutputsTcpDefault) ToOutputsTcpDefaultOutput() OutputsTcpDefaultOutput {
	return i.ToOutputsTcpDefaultOutputWithContext(context.Background())
}

func (i *OutputsTcpDefault) ToOutputsTcpDefaultOutputWithContext(ctx context.Context) OutputsTcpDefaultOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OutputsTcpDefaultOutput)
}

func (i *OutputsTcpDefault) ToOutputsTcpDefaultPtrOutput() OutputsTcpDefaultPtrOutput {
	return i.ToOutputsTcpDefaultPtrOutputWithContext(context.Background())
}

func (i *OutputsTcpDefault) ToOutputsTcpDefaultPtrOutputWithContext(ctx context.Context) OutputsTcpDefaultPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OutputsTcpDefaultPtrOutput)
}

type OutputsTcpDefaultPtrInput interface {
	pulumi.Input

	ToOutputsTcpDefaultPtrOutput() OutputsTcpDefaultPtrOutput
	ToOutputsTcpDefaultPtrOutputWithContext(ctx context.Context) OutputsTcpDefaultPtrOutput
}

type outputsTcpDefaultPtrType OutputsTcpDefaultArgs

func (*outputsTcpDefaultPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**OutputsTcpDefault)(nil))
}

func (i *outputsTcpDefaultPtrType) ToOutputsTcpDefaultPtrOutput() OutputsTcpDefaultPtrOutput {
	return i.ToOutputsTcpDefaultPtrOutputWithContext(context.Background())
}

func (i *outputsTcpDefaultPtrType) ToOutputsTcpDefaultPtrOutputWithContext(ctx context.Context) OutputsTcpDefaultPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OutputsTcpDefaultPtrOutput)
}

// OutputsTcpDefaultArrayInput is an input type that accepts OutputsTcpDefaultArray and OutputsTcpDefaultArrayOutput values.
// You can construct a concrete instance of `OutputsTcpDefaultArrayInput` via:
//
//          OutputsTcpDefaultArray{ OutputsTcpDefaultArgs{...} }
type OutputsTcpDefaultArrayInput interface {
	pulumi.Input

	ToOutputsTcpDefaultArrayOutput() OutputsTcpDefaultArrayOutput
	ToOutputsTcpDefaultArrayOutputWithContext(context.Context) OutputsTcpDefaultArrayOutput
}

type OutputsTcpDefaultArray []OutputsTcpDefaultInput

func (OutputsTcpDefaultArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*OutputsTcpDefault)(nil))
}

func (i OutputsTcpDefaultArray) ToOutputsTcpDefaultArrayOutput() OutputsTcpDefaultArrayOutput {
	return i.ToOutputsTcpDefaultArrayOutputWithContext(context.Background())
}

func (i OutputsTcpDefaultArray) ToOutputsTcpDefaultArrayOutputWithContext(ctx context.Context) OutputsTcpDefaultArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OutputsTcpDefaultArrayOutput)
}

// OutputsTcpDefaultMapInput is an input type that accepts OutputsTcpDefaultMap and OutputsTcpDefaultMapOutput values.
// You can construct a concrete instance of `OutputsTcpDefaultMapInput` via:
//
//          OutputsTcpDefaultMap{ "key": OutputsTcpDefaultArgs{...} }
type OutputsTcpDefaultMapInput interface {
	pulumi.Input

	ToOutputsTcpDefaultMapOutput() OutputsTcpDefaultMapOutput
	ToOutputsTcpDefaultMapOutputWithContext(context.Context) OutputsTcpDefaultMapOutput
}

type OutputsTcpDefaultMap map[string]OutputsTcpDefaultInput

func (OutputsTcpDefaultMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*OutputsTcpDefault)(nil))
}

func (i OutputsTcpDefaultMap) ToOutputsTcpDefaultMapOutput() OutputsTcpDefaultMapOutput {
	return i.ToOutputsTcpDefaultMapOutputWithContext(context.Background())
}

func (i OutputsTcpDefaultMap) ToOutputsTcpDefaultMapOutputWithContext(ctx context.Context) OutputsTcpDefaultMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OutputsTcpDefaultMapOutput)
}

type OutputsTcpDefaultOutput struct {
	*pulumi.OutputState
}

func (OutputsTcpDefaultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*OutputsTcpDefault)(nil))
}

func (o OutputsTcpDefaultOutput) ToOutputsTcpDefaultOutput() OutputsTcpDefaultOutput {
	return o
}

func (o OutputsTcpDefaultOutput) ToOutputsTcpDefaultOutputWithContext(ctx context.Context) OutputsTcpDefaultOutput {
	return o
}

func (o OutputsTcpDefaultOutput) ToOutputsTcpDefaultPtrOutput() OutputsTcpDefaultPtrOutput {
	return o.ToOutputsTcpDefaultPtrOutputWithContext(context.Background())
}

func (o OutputsTcpDefaultOutput) ToOutputsTcpDefaultPtrOutputWithContext(ctx context.Context) OutputsTcpDefaultPtrOutput {
	return o.ApplyT(func(v OutputsTcpDefault) *OutputsTcpDefault {
		return &v
	}).(OutputsTcpDefaultPtrOutput)
}

type OutputsTcpDefaultPtrOutput struct {
	*pulumi.OutputState
}

func (OutputsTcpDefaultPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**OutputsTcpDefault)(nil))
}

func (o OutputsTcpDefaultPtrOutput) ToOutputsTcpDefaultPtrOutput() OutputsTcpDefaultPtrOutput {
	return o
}

func (o OutputsTcpDefaultPtrOutput) ToOutputsTcpDefaultPtrOutputWithContext(ctx context.Context) OutputsTcpDefaultPtrOutput {
	return o
}

type OutputsTcpDefaultArrayOutput struct{ *pulumi.OutputState }

func (OutputsTcpDefaultArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]OutputsTcpDefault)(nil))
}

func (o OutputsTcpDefaultArrayOutput) ToOutputsTcpDefaultArrayOutput() OutputsTcpDefaultArrayOutput {
	return o
}

func (o OutputsTcpDefaultArrayOutput) ToOutputsTcpDefaultArrayOutputWithContext(ctx context.Context) OutputsTcpDefaultArrayOutput {
	return o
}

func (o OutputsTcpDefaultArrayOutput) Index(i pulumi.IntInput) OutputsTcpDefaultOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) OutputsTcpDefault {
		return vs[0].([]OutputsTcpDefault)[vs[1].(int)]
	}).(OutputsTcpDefaultOutput)
}

type OutputsTcpDefaultMapOutput struct{ *pulumi.OutputState }

func (OutputsTcpDefaultMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]OutputsTcpDefault)(nil))
}

func (o OutputsTcpDefaultMapOutput) ToOutputsTcpDefaultMapOutput() OutputsTcpDefaultMapOutput {
	return o
}

func (o OutputsTcpDefaultMapOutput) ToOutputsTcpDefaultMapOutputWithContext(ctx context.Context) OutputsTcpDefaultMapOutput {
	return o
}

func (o OutputsTcpDefaultMapOutput) MapIndex(k pulumi.StringInput) OutputsTcpDefaultOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) OutputsTcpDefault {
		return vs[0].(map[string]OutputsTcpDefault)[vs[1].(string)]
	}).(OutputsTcpDefaultOutput)
}

func init() {
	pulumi.RegisterOutputType(OutputsTcpDefaultOutput{})
	pulumi.RegisterOutputType(OutputsTcpDefaultPtrOutput{})
	pulumi.RegisterOutputType(OutputsTcpDefaultArrayOutput{})
	pulumi.RegisterOutputType(OutputsTcpDefaultMapOutput{})
}
