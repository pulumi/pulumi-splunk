// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package splunk

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## # Resource: OutputsTcpGroup
//
// Access to the configuration of a group of one or more data forwarding destinations.
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
// 		_, err := splunk.NewOutputsTcpGroup(ctx, "tcpGroup", &splunk.OutputsTcpGroupArgs{
// 			Disabled:              pulumi.Bool(false),
// 			DropEventsOnQueueFull: pulumi.Int(60),
// 			MaxQueueSize:          pulumi.String("100KB"),
// 			SendCookedData:        pulumi.Bool(true),
// 			Servers: pulumi.StringArray{
// 				pulumi.String("1.1.1.1:1234"),
// 				pulumi.String("2.2.2.2:1234"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type OutputsTcpGroup struct {
	pulumi.CustomResourceState

	// The app/user context that is the namespace for the resource
	Acl OutputsTcpGroupAclOutput `pulumi:"acl"`
	// If true, forwarder sends compressed data. If set to true, the receiver port must also have compression turned on.
	Compressed pulumi.BoolOutput `pulumi:"compressed"`
	// If true, disables the group.
	Disabled pulumi.BoolOutput `pulumi:"disabled"`
	// If set to a positive number, wait the specified number of seconds before throwing out all new events until the output queue has space. Defaults to -1 (do not drop events).
	// <br>CAUTION: Do not set this value to a positive integer if you are monitoring files.
	// Setting this to -1 or 0 causes the output queue to block when it gets full, which causes further blocking up the processing chain. If any target group queue is blocked, no more data reaches any other target group.
	// Using auto load-balancing is the best way to minimize this condition, because, in that case, multiple receivers must be down (or jammed up) before queue blocking can occur.
	DropEventsOnQueueFull pulumi.IntOutput `pulumi:"dropEventsOnQueueFull"`
	// How often (in seconds) to send a heartbeat packet to the receiving server.
	// Heartbeats are only sent if sendCookedData=true. Defaults to 30 seconds.
	HeartbeatFrequency pulumi.IntOutput `pulumi:"heartbeatFrequency"`
	// Specify an integer or integer[KB|MB|GB].
	// <br>Sets the maximum size of the forwarder output queue. It also sets the maximum size of the wait queue to 3x this value, if you have enabled indexer acknowledgment (useACK=true).
	// Although the wait queue and the output queues are both configured by this attribute, they are separate queues. The setting determines the maximum size of the queue in-memory (RAM) buffer.
	// For heavy forwarders sending parsed data, maxQueueSize is the maximum number of events. Since events are typically much shorter than data blocks, the memory consumed by the queue on a parsing forwarder is likely to be much smaller than on a non-parsing forwarder, if you use this version of the setting.
	// If specified as a lone integer (for example, maxQueueSize=100), maxQueueSize indicates the maximum number of queued events (for parsed data) or blocks of data (for unparsed data). A block of data is approximately 64KB. For non-parsing forwarders, such as universal forwarders, that send unparsed data, maxQueueSize is the maximum number of data blocks.
	// If specified as an integer followed by KB, MB, or GB (for example, maxQueueSize=100MB), maxQueueSize indicates the maximum RAM allocated to the queue buffer. Defaults to 500KB (which means a maximum size of 500KB for the output queue and 1500KB for the wait queue, if any).
	MaxQueueSize pulumi.StringOutput `pulumi:"maxQueueSize"`
	// Valid values: (tcpout | syslog). Specifies the type of output processor.
	Method pulumi.StringOutput `pulumi:"method"`
	// The name of the group of receivers.
	Name pulumi.StringOutput `pulumi:"name"`
	// If true, events are cooked (processed by Splunk software). If false, events are raw and untouched prior to sending. Defaults to true.
	// Set to false if you are sending to a third-party system.
	SendCookedData pulumi.BoolOutput `pulumi:"sendCookedData"`
	// Comma-separated list of servers to include in the group.
	Servers pulumi.StringArrayOutput `pulumi:"servers"`
	// Token value generated by the indexer after configuration.
	Token pulumi.StringOutput `pulumi:"token"`
}

// NewOutputsTcpGroup registers a new resource with the given unique name, arguments, and options.
func NewOutputsTcpGroup(ctx *pulumi.Context,
	name string, args *OutputsTcpGroupArgs, opts ...pulumi.ResourceOption) (*OutputsTcpGroup, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Servers == nil {
		return nil, errors.New("invalid value for required argument 'Servers'")
	}
	var resource OutputsTcpGroup
	err := ctx.RegisterResource("splunk:index/outputsTcpGroup:OutputsTcpGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetOutputsTcpGroup gets an existing OutputsTcpGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetOutputsTcpGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *OutputsTcpGroupState, opts ...pulumi.ResourceOption) (*OutputsTcpGroup, error) {
	var resource OutputsTcpGroup
	err := ctx.ReadResource("splunk:index/outputsTcpGroup:OutputsTcpGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering OutputsTcpGroup resources.
type outputsTcpGroupState struct {
	// The app/user context that is the namespace for the resource
	Acl *OutputsTcpGroupAcl `pulumi:"acl"`
	// If true, forwarder sends compressed data. If set to true, the receiver port must also have compression turned on.
	Compressed *bool `pulumi:"compressed"`
	// If true, disables the group.
	Disabled *bool `pulumi:"disabled"`
	// If set to a positive number, wait the specified number of seconds before throwing out all new events until the output queue has space. Defaults to -1 (do not drop events).
	// <br>CAUTION: Do not set this value to a positive integer if you are monitoring files.
	// Setting this to -1 or 0 causes the output queue to block when it gets full, which causes further blocking up the processing chain. If any target group queue is blocked, no more data reaches any other target group.
	// Using auto load-balancing is the best way to minimize this condition, because, in that case, multiple receivers must be down (or jammed up) before queue blocking can occur.
	DropEventsOnQueueFull *int `pulumi:"dropEventsOnQueueFull"`
	// How often (in seconds) to send a heartbeat packet to the receiving server.
	// Heartbeats are only sent if sendCookedData=true. Defaults to 30 seconds.
	HeartbeatFrequency *int `pulumi:"heartbeatFrequency"`
	// Specify an integer or integer[KB|MB|GB].
	// <br>Sets the maximum size of the forwarder output queue. It also sets the maximum size of the wait queue to 3x this value, if you have enabled indexer acknowledgment (useACK=true).
	// Although the wait queue and the output queues are both configured by this attribute, they are separate queues. The setting determines the maximum size of the queue in-memory (RAM) buffer.
	// For heavy forwarders sending parsed data, maxQueueSize is the maximum number of events. Since events are typically much shorter than data blocks, the memory consumed by the queue on a parsing forwarder is likely to be much smaller than on a non-parsing forwarder, if you use this version of the setting.
	// If specified as a lone integer (for example, maxQueueSize=100), maxQueueSize indicates the maximum number of queued events (for parsed data) or blocks of data (for unparsed data). A block of data is approximately 64KB. For non-parsing forwarders, such as universal forwarders, that send unparsed data, maxQueueSize is the maximum number of data blocks.
	// If specified as an integer followed by KB, MB, or GB (for example, maxQueueSize=100MB), maxQueueSize indicates the maximum RAM allocated to the queue buffer. Defaults to 500KB (which means a maximum size of 500KB for the output queue and 1500KB for the wait queue, if any).
	MaxQueueSize *string `pulumi:"maxQueueSize"`
	// Valid values: (tcpout | syslog). Specifies the type of output processor.
	Method *string `pulumi:"method"`
	// The name of the group of receivers.
	Name *string `pulumi:"name"`
	// If true, events are cooked (processed by Splunk software). If false, events are raw and untouched prior to sending. Defaults to true.
	// Set to false if you are sending to a third-party system.
	SendCookedData *bool `pulumi:"sendCookedData"`
	// Comma-separated list of servers to include in the group.
	Servers []string `pulumi:"servers"`
	// Token value generated by the indexer after configuration.
	Token *string `pulumi:"token"`
}

type OutputsTcpGroupState struct {
	// The app/user context that is the namespace for the resource
	Acl OutputsTcpGroupAclPtrInput
	// If true, forwarder sends compressed data. If set to true, the receiver port must also have compression turned on.
	Compressed pulumi.BoolPtrInput
	// If true, disables the group.
	Disabled pulumi.BoolPtrInput
	// If set to a positive number, wait the specified number of seconds before throwing out all new events until the output queue has space. Defaults to -1 (do not drop events).
	// <br>CAUTION: Do not set this value to a positive integer if you are monitoring files.
	// Setting this to -1 or 0 causes the output queue to block when it gets full, which causes further blocking up the processing chain. If any target group queue is blocked, no more data reaches any other target group.
	// Using auto load-balancing is the best way to minimize this condition, because, in that case, multiple receivers must be down (or jammed up) before queue blocking can occur.
	DropEventsOnQueueFull pulumi.IntPtrInput
	// How often (in seconds) to send a heartbeat packet to the receiving server.
	// Heartbeats are only sent if sendCookedData=true. Defaults to 30 seconds.
	HeartbeatFrequency pulumi.IntPtrInput
	// Specify an integer or integer[KB|MB|GB].
	// <br>Sets the maximum size of the forwarder output queue. It also sets the maximum size of the wait queue to 3x this value, if you have enabled indexer acknowledgment (useACK=true).
	// Although the wait queue and the output queues are both configured by this attribute, they are separate queues. The setting determines the maximum size of the queue in-memory (RAM) buffer.
	// For heavy forwarders sending parsed data, maxQueueSize is the maximum number of events. Since events are typically much shorter than data blocks, the memory consumed by the queue on a parsing forwarder is likely to be much smaller than on a non-parsing forwarder, if you use this version of the setting.
	// If specified as a lone integer (for example, maxQueueSize=100), maxQueueSize indicates the maximum number of queued events (for parsed data) or blocks of data (for unparsed data). A block of data is approximately 64KB. For non-parsing forwarders, such as universal forwarders, that send unparsed data, maxQueueSize is the maximum number of data blocks.
	// If specified as an integer followed by KB, MB, or GB (for example, maxQueueSize=100MB), maxQueueSize indicates the maximum RAM allocated to the queue buffer. Defaults to 500KB (which means a maximum size of 500KB for the output queue and 1500KB for the wait queue, if any).
	MaxQueueSize pulumi.StringPtrInput
	// Valid values: (tcpout | syslog). Specifies the type of output processor.
	Method pulumi.StringPtrInput
	// The name of the group of receivers.
	Name pulumi.StringPtrInput
	// If true, events are cooked (processed by Splunk software). If false, events are raw and untouched prior to sending. Defaults to true.
	// Set to false if you are sending to a third-party system.
	SendCookedData pulumi.BoolPtrInput
	// Comma-separated list of servers to include in the group.
	Servers pulumi.StringArrayInput
	// Token value generated by the indexer after configuration.
	Token pulumi.StringPtrInput
}

func (OutputsTcpGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*outputsTcpGroupState)(nil)).Elem()
}

type outputsTcpGroupArgs struct {
	// The app/user context that is the namespace for the resource
	Acl *OutputsTcpGroupAcl `pulumi:"acl"`
	// If true, forwarder sends compressed data. If set to true, the receiver port must also have compression turned on.
	Compressed *bool `pulumi:"compressed"`
	// If true, disables the group.
	Disabled *bool `pulumi:"disabled"`
	// If set to a positive number, wait the specified number of seconds before throwing out all new events until the output queue has space. Defaults to -1 (do not drop events).
	// <br>CAUTION: Do not set this value to a positive integer if you are monitoring files.
	// Setting this to -1 or 0 causes the output queue to block when it gets full, which causes further blocking up the processing chain. If any target group queue is blocked, no more data reaches any other target group.
	// Using auto load-balancing is the best way to minimize this condition, because, in that case, multiple receivers must be down (or jammed up) before queue blocking can occur.
	DropEventsOnQueueFull *int `pulumi:"dropEventsOnQueueFull"`
	// How often (in seconds) to send a heartbeat packet to the receiving server.
	// Heartbeats are only sent if sendCookedData=true. Defaults to 30 seconds.
	HeartbeatFrequency *int `pulumi:"heartbeatFrequency"`
	// Specify an integer or integer[KB|MB|GB].
	// <br>Sets the maximum size of the forwarder output queue. It also sets the maximum size of the wait queue to 3x this value, if you have enabled indexer acknowledgment (useACK=true).
	// Although the wait queue and the output queues are both configured by this attribute, they are separate queues. The setting determines the maximum size of the queue in-memory (RAM) buffer.
	// For heavy forwarders sending parsed data, maxQueueSize is the maximum number of events. Since events are typically much shorter than data blocks, the memory consumed by the queue on a parsing forwarder is likely to be much smaller than on a non-parsing forwarder, if you use this version of the setting.
	// If specified as a lone integer (for example, maxQueueSize=100), maxQueueSize indicates the maximum number of queued events (for parsed data) or blocks of data (for unparsed data). A block of data is approximately 64KB. For non-parsing forwarders, such as universal forwarders, that send unparsed data, maxQueueSize is the maximum number of data blocks.
	// If specified as an integer followed by KB, MB, or GB (for example, maxQueueSize=100MB), maxQueueSize indicates the maximum RAM allocated to the queue buffer. Defaults to 500KB (which means a maximum size of 500KB for the output queue and 1500KB for the wait queue, if any).
	MaxQueueSize *string `pulumi:"maxQueueSize"`
	// Valid values: (tcpout | syslog). Specifies the type of output processor.
	Method *string `pulumi:"method"`
	// The name of the group of receivers.
	Name *string `pulumi:"name"`
	// If true, events are cooked (processed by Splunk software). If false, events are raw and untouched prior to sending. Defaults to true.
	// Set to false if you are sending to a third-party system.
	SendCookedData *bool `pulumi:"sendCookedData"`
	// Comma-separated list of servers to include in the group.
	Servers []string `pulumi:"servers"`
	// Token value generated by the indexer after configuration.
	Token *string `pulumi:"token"`
}

// The set of arguments for constructing a OutputsTcpGroup resource.
type OutputsTcpGroupArgs struct {
	// The app/user context that is the namespace for the resource
	Acl OutputsTcpGroupAclPtrInput
	// If true, forwarder sends compressed data. If set to true, the receiver port must also have compression turned on.
	Compressed pulumi.BoolPtrInput
	// If true, disables the group.
	Disabled pulumi.BoolPtrInput
	// If set to a positive number, wait the specified number of seconds before throwing out all new events until the output queue has space. Defaults to -1 (do not drop events).
	// <br>CAUTION: Do not set this value to a positive integer if you are monitoring files.
	// Setting this to -1 or 0 causes the output queue to block when it gets full, which causes further blocking up the processing chain. If any target group queue is blocked, no more data reaches any other target group.
	// Using auto load-balancing is the best way to minimize this condition, because, in that case, multiple receivers must be down (or jammed up) before queue blocking can occur.
	DropEventsOnQueueFull pulumi.IntPtrInput
	// How often (in seconds) to send a heartbeat packet to the receiving server.
	// Heartbeats are only sent if sendCookedData=true. Defaults to 30 seconds.
	HeartbeatFrequency pulumi.IntPtrInput
	// Specify an integer or integer[KB|MB|GB].
	// <br>Sets the maximum size of the forwarder output queue. It also sets the maximum size of the wait queue to 3x this value, if you have enabled indexer acknowledgment (useACK=true).
	// Although the wait queue and the output queues are both configured by this attribute, they are separate queues. The setting determines the maximum size of the queue in-memory (RAM) buffer.
	// For heavy forwarders sending parsed data, maxQueueSize is the maximum number of events. Since events are typically much shorter than data blocks, the memory consumed by the queue on a parsing forwarder is likely to be much smaller than on a non-parsing forwarder, if you use this version of the setting.
	// If specified as a lone integer (for example, maxQueueSize=100), maxQueueSize indicates the maximum number of queued events (for parsed data) or blocks of data (for unparsed data). A block of data is approximately 64KB. For non-parsing forwarders, such as universal forwarders, that send unparsed data, maxQueueSize is the maximum number of data blocks.
	// If specified as an integer followed by KB, MB, or GB (for example, maxQueueSize=100MB), maxQueueSize indicates the maximum RAM allocated to the queue buffer. Defaults to 500KB (which means a maximum size of 500KB for the output queue and 1500KB for the wait queue, if any).
	MaxQueueSize pulumi.StringPtrInput
	// Valid values: (tcpout | syslog). Specifies the type of output processor.
	Method pulumi.StringPtrInput
	// The name of the group of receivers.
	Name pulumi.StringPtrInput
	// If true, events are cooked (processed by Splunk software). If false, events are raw and untouched prior to sending. Defaults to true.
	// Set to false if you are sending to a third-party system.
	SendCookedData pulumi.BoolPtrInput
	// Comma-separated list of servers to include in the group.
	Servers pulumi.StringArrayInput
	// Token value generated by the indexer after configuration.
	Token pulumi.StringPtrInput
}

func (OutputsTcpGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*outputsTcpGroupArgs)(nil)).Elem()
}

type OutputsTcpGroupInput interface {
	pulumi.Input

	ToOutputsTcpGroupOutput() OutputsTcpGroupOutput
	ToOutputsTcpGroupOutputWithContext(ctx context.Context) OutputsTcpGroupOutput
}

func (*OutputsTcpGroup) ElementType() reflect.Type {
	return reflect.TypeOf((**OutputsTcpGroup)(nil)).Elem()
}

func (i *OutputsTcpGroup) ToOutputsTcpGroupOutput() OutputsTcpGroupOutput {
	return i.ToOutputsTcpGroupOutputWithContext(context.Background())
}

func (i *OutputsTcpGroup) ToOutputsTcpGroupOutputWithContext(ctx context.Context) OutputsTcpGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OutputsTcpGroupOutput)
}

// OutputsTcpGroupArrayInput is an input type that accepts OutputsTcpGroupArray and OutputsTcpGroupArrayOutput values.
// You can construct a concrete instance of `OutputsTcpGroupArrayInput` via:
//
//          OutputsTcpGroupArray{ OutputsTcpGroupArgs{...} }
type OutputsTcpGroupArrayInput interface {
	pulumi.Input

	ToOutputsTcpGroupArrayOutput() OutputsTcpGroupArrayOutput
	ToOutputsTcpGroupArrayOutputWithContext(context.Context) OutputsTcpGroupArrayOutput
}

type OutputsTcpGroupArray []OutputsTcpGroupInput

func (OutputsTcpGroupArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*OutputsTcpGroup)(nil)).Elem()
}

func (i OutputsTcpGroupArray) ToOutputsTcpGroupArrayOutput() OutputsTcpGroupArrayOutput {
	return i.ToOutputsTcpGroupArrayOutputWithContext(context.Background())
}

func (i OutputsTcpGroupArray) ToOutputsTcpGroupArrayOutputWithContext(ctx context.Context) OutputsTcpGroupArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OutputsTcpGroupArrayOutput)
}

// OutputsTcpGroupMapInput is an input type that accepts OutputsTcpGroupMap and OutputsTcpGroupMapOutput values.
// You can construct a concrete instance of `OutputsTcpGroupMapInput` via:
//
//          OutputsTcpGroupMap{ "key": OutputsTcpGroupArgs{...} }
type OutputsTcpGroupMapInput interface {
	pulumi.Input

	ToOutputsTcpGroupMapOutput() OutputsTcpGroupMapOutput
	ToOutputsTcpGroupMapOutputWithContext(context.Context) OutputsTcpGroupMapOutput
}

type OutputsTcpGroupMap map[string]OutputsTcpGroupInput

func (OutputsTcpGroupMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*OutputsTcpGroup)(nil)).Elem()
}

func (i OutputsTcpGroupMap) ToOutputsTcpGroupMapOutput() OutputsTcpGroupMapOutput {
	return i.ToOutputsTcpGroupMapOutputWithContext(context.Background())
}

func (i OutputsTcpGroupMap) ToOutputsTcpGroupMapOutputWithContext(ctx context.Context) OutputsTcpGroupMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OutputsTcpGroupMapOutput)
}

type OutputsTcpGroupOutput struct{ *pulumi.OutputState }

func (OutputsTcpGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**OutputsTcpGroup)(nil)).Elem()
}

func (o OutputsTcpGroupOutput) ToOutputsTcpGroupOutput() OutputsTcpGroupOutput {
	return o
}

func (o OutputsTcpGroupOutput) ToOutputsTcpGroupOutputWithContext(ctx context.Context) OutputsTcpGroupOutput {
	return o
}

type OutputsTcpGroupArrayOutput struct{ *pulumi.OutputState }

func (OutputsTcpGroupArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*OutputsTcpGroup)(nil)).Elem()
}

func (o OutputsTcpGroupArrayOutput) ToOutputsTcpGroupArrayOutput() OutputsTcpGroupArrayOutput {
	return o
}

func (o OutputsTcpGroupArrayOutput) ToOutputsTcpGroupArrayOutputWithContext(ctx context.Context) OutputsTcpGroupArrayOutput {
	return o
}

func (o OutputsTcpGroupArrayOutput) Index(i pulumi.IntInput) OutputsTcpGroupOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *OutputsTcpGroup {
		return vs[0].([]*OutputsTcpGroup)[vs[1].(int)]
	}).(OutputsTcpGroupOutput)
}

type OutputsTcpGroupMapOutput struct{ *pulumi.OutputState }

func (OutputsTcpGroupMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*OutputsTcpGroup)(nil)).Elem()
}

func (o OutputsTcpGroupMapOutput) ToOutputsTcpGroupMapOutput() OutputsTcpGroupMapOutput {
	return o
}

func (o OutputsTcpGroupMapOutput) ToOutputsTcpGroupMapOutputWithContext(ctx context.Context) OutputsTcpGroupMapOutput {
	return o
}

func (o OutputsTcpGroupMapOutput) MapIndex(k pulumi.StringInput) OutputsTcpGroupOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *OutputsTcpGroup {
		return vs[0].(map[string]*OutputsTcpGroup)[vs[1].(string)]
	}).(OutputsTcpGroupOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*OutputsTcpGroupInput)(nil)).Elem(), &OutputsTcpGroup{})
	pulumi.RegisterInputType(reflect.TypeOf((*OutputsTcpGroupArrayInput)(nil)).Elem(), OutputsTcpGroupArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*OutputsTcpGroupMapInput)(nil)).Elem(), OutputsTcpGroupMap{})
	pulumi.RegisterOutputType(OutputsTcpGroupOutput{})
	pulumi.RegisterOutputType(OutputsTcpGroupArrayOutput{})
	pulumi.RegisterOutputType(OutputsTcpGroupMapOutput{})
}
