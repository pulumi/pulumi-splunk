// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package splunk

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## # Resource: InputsTcpRaw
//
// Create or update raw TCP input information for managing raw tcp inputs from forwarders.
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
// 		_, err := splunk.NewInputsTcpRaw(ctx, "tcpRaw", &splunk.InputsTcpRawArgs{
// 			Disabled:   pulumi.Bool(false),
// 			Index:      pulumi.String("main"),
// 			Queue:      pulumi.String("indexQueue"),
// 			Source:     pulumi.String("new"),
// 			Sourcetype: pulumi.String("new"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type InputsTcpRaw struct {
	pulumi.CustomResourceState

	// The app/user context that is the namespace for the resource
	Acl InputsTcpRawAclOutput `pulumi:"acl"`
	// Valid values: (ip | dns | none)
	// Set the host for the remote server that is sending data.
	// ip sets the host to the IP address of the remote server sending data.
	// dns sets the host to the reverse DNS entry for the IP address of the remote server sending data.
	// none leaves the host as specified in inputs.conf, which is typically the Splunk system hostname.
	// Default value is dns.
	ConnectionHost pulumi.StringOutput `pulumi:"connectionHost"`
	// Indicates if input is disabled.
	Disabled pulumi.BoolOutput `pulumi:"disabled"`
	// Host from which the indexer gets data.
	Host pulumi.StringOutput `pulumi:"host"`
	// Index to store generated events. Defaults to default.
	Index pulumi.StringOutput `pulumi:"index"`
	// The input port which receives raw data.
	Name pulumi.StringOutput `pulumi:"name"`
	// Valid values: (parsingQueue | indexQueue)
	// Specifies where the input processor should deposit the events it reads. Defaults to parsingQueue.
	// Set queue to parsingQueue to apply props.conf and other parsing rules to your data. For more information about props.conf and rules for timestamping and linebreaking, refer to props.conf and the online documentation at "Monitor files and directories with inputs.conf"
	// Set queue to indexQueue to send your data directly into the index.
	Queue pulumi.StringOutput `pulumi:"queue"`
	// Specifies in seconds the timeout value for adding a Done-key. Default value is 10 seconds.
	// If a connection over the port specified by name remains idle after receiving data for specified number of seconds, it adds a Done-key. This implies the last event is completely received.
	RawTcpDoneTimeout pulumi.IntOutput `pulumi:"rawTcpDoneTimeout"`
	// Allows for restricting this input to only accept data from the host specified here.
	RestrictToHost pulumi.StringOutput `pulumi:"restrictToHost"`
	// Sets the source key/field for events from this input. Defaults to the input file path.
	// Sets the source key initial value. The key is used during parsing/indexing, in particular to set the source field during indexing. It is also the source field used at search time. As a convenience, the chosen string is prepended with 'source::'.
	Source pulumi.StringOutput `pulumi:"source"`
	// Set the source type for events from this input.
	// "sourcetype=" is automatically prepended to <string>.
	// Defaults to audittrail (if signedaudit=true) or fschange (if signedaudit=false).
	Sourcetype pulumi.StringOutput `pulumi:"sourcetype"`
}

// NewInputsTcpRaw registers a new resource with the given unique name, arguments, and options.
func NewInputsTcpRaw(ctx *pulumi.Context,
	name string, args *InputsTcpRawArgs, opts ...pulumi.ResourceOption) (*InputsTcpRaw, error) {
	if args == nil {
		args = &InputsTcpRawArgs{}
	}

	var resource InputsTcpRaw
	err := ctx.RegisterResource("splunk:index/inputsTcpRaw:InputsTcpRaw", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetInputsTcpRaw gets an existing InputsTcpRaw resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetInputsTcpRaw(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *InputsTcpRawState, opts ...pulumi.ResourceOption) (*InputsTcpRaw, error) {
	var resource InputsTcpRaw
	err := ctx.ReadResource("splunk:index/inputsTcpRaw:InputsTcpRaw", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering InputsTcpRaw resources.
type inputsTcpRawState struct {
	// The app/user context that is the namespace for the resource
	Acl *InputsTcpRawAcl `pulumi:"acl"`
	// Valid values: (ip | dns | none)
	// Set the host for the remote server that is sending data.
	// ip sets the host to the IP address of the remote server sending data.
	// dns sets the host to the reverse DNS entry for the IP address of the remote server sending data.
	// none leaves the host as specified in inputs.conf, which is typically the Splunk system hostname.
	// Default value is dns.
	ConnectionHost *string `pulumi:"connectionHost"`
	// Indicates if input is disabled.
	Disabled *bool `pulumi:"disabled"`
	// Host from which the indexer gets data.
	Host *string `pulumi:"host"`
	// Index to store generated events. Defaults to default.
	Index *string `pulumi:"index"`
	// The input port which receives raw data.
	Name *string `pulumi:"name"`
	// Valid values: (parsingQueue | indexQueue)
	// Specifies where the input processor should deposit the events it reads. Defaults to parsingQueue.
	// Set queue to parsingQueue to apply props.conf and other parsing rules to your data. For more information about props.conf and rules for timestamping and linebreaking, refer to props.conf and the online documentation at "Monitor files and directories with inputs.conf"
	// Set queue to indexQueue to send your data directly into the index.
	Queue *string `pulumi:"queue"`
	// Specifies in seconds the timeout value for adding a Done-key. Default value is 10 seconds.
	// If a connection over the port specified by name remains idle after receiving data for specified number of seconds, it adds a Done-key. This implies the last event is completely received.
	RawTcpDoneTimeout *int `pulumi:"rawTcpDoneTimeout"`
	// Allows for restricting this input to only accept data from the host specified here.
	RestrictToHost *string `pulumi:"restrictToHost"`
	// Sets the source key/field for events from this input. Defaults to the input file path.
	// Sets the source key initial value. The key is used during parsing/indexing, in particular to set the source field during indexing. It is also the source field used at search time. As a convenience, the chosen string is prepended with 'source::'.
	Source *string `pulumi:"source"`
	// Set the source type for events from this input.
	// "sourcetype=" is automatically prepended to <string>.
	// Defaults to audittrail (if signedaudit=true) or fschange (if signedaudit=false).
	Sourcetype *string `pulumi:"sourcetype"`
}

type InputsTcpRawState struct {
	// The app/user context that is the namespace for the resource
	Acl InputsTcpRawAclPtrInput
	// Valid values: (ip | dns | none)
	// Set the host for the remote server that is sending data.
	// ip sets the host to the IP address of the remote server sending data.
	// dns sets the host to the reverse DNS entry for the IP address of the remote server sending data.
	// none leaves the host as specified in inputs.conf, which is typically the Splunk system hostname.
	// Default value is dns.
	ConnectionHost pulumi.StringPtrInput
	// Indicates if input is disabled.
	Disabled pulumi.BoolPtrInput
	// Host from which the indexer gets data.
	Host pulumi.StringPtrInput
	// Index to store generated events. Defaults to default.
	Index pulumi.StringPtrInput
	// The input port which receives raw data.
	Name pulumi.StringPtrInput
	// Valid values: (parsingQueue | indexQueue)
	// Specifies where the input processor should deposit the events it reads. Defaults to parsingQueue.
	// Set queue to parsingQueue to apply props.conf and other parsing rules to your data. For more information about props.conf and rules for timestamping and linebreaking, refer to props.conf and the online documentation at "Monitor files and directories with inputs.conf"
	// Set queue to indexQueue to send your data directly into the index.
	Queue pulumi.StringPtrInput
	// Specifies in seconds the timeout value for adding a Done-key. Default value is 10 seconds.
	// If a connection over the port specified by name remains idle after receiving data for specified number of seconds, it adds a Done-key. This implies the last event is completely received.
	RawTcpDoneTimeout pulumi.IntPtrInput
	// Allows for restricting this input to only accept data from the host specified here.
	RestrictToHost pulumi.StringPtrInput
	// Sets the source key/field for events from this input. Defaults to the input file path.
	// Sets the source key initial value. The key is used during parsing/indexing, in particular to set the source field during indexing. It is also the source field used at search time. As a convenience, the chosen string is prepended with 'source::'.
	Source pulumi.StringPtrInput
	// Set the source type for events from this input.
	// "sourcetype=" is automatically prepended to <string>.
	// Defaults to audittrail (if signedaudit=true) or fschange (if signedaudit=false).
	Sourcetype pulumi.StringPtrInput
}

func (InputsTcpRawState) ElementType() reflect.Type {
	return reflect.TypeOf((*inputsTcpRawState)(nil)).Elem()
}

type inputsTcpRawArgs struct {
	// The app/user context that is the namespace for the resource
	Acl *InputsTcpRawAcl `pulumi:"acl"`
	// Valid values: (ip | dns | none)
	// Set the host for the remote server that is sending data.
	// ip sets the host to the IP address of the remote server sending data.
	// dns sets the host to the reverse DNS entry for the IP address of the remote server sending data.
	// none leaves the host as specified in inputs.conf, which is typically the Splunk system hostname.
	// Default value is dns.
	ConnectionHost *string `pulumi:"connectionHost"`
	// Indicates if input is disabled.
	Disabled *bool `pulumi:"disabled"`
	// Host from which the indexer gets data.
	Host *string `pulumi:"host"`
	// Index to store generated events. Defaults to default.
	Index *string `pulumi:"index"`
	// The input port which receives raw data.
	Name *string `pulumi:"name"`
	// Valid values: (parsingQueue | indexQueue)
	// Specifies where the input processor should deposit the events it reads. Defaults to parsingQueue.
	// Set queue to parsingQueue to apply props.conf and other parsing rules to your data. For more information about props.conf and rules for timestamping and linebreaking, refer to props.conf and the online documentation at "Monitor files and directories with inputs.conf"
	// Set queue to indexQueue to send your data directly into the index.
	Queue *string `pulumi:"queue"`
	// Specifies in seconds the timeout value for adding a Done-key. Default value is 10 seconds.
	// If a connection over the port specified by name remains idle after receiving data for specified number of seconds, it adds a Done-key. This implies the last event is completely received.
	RawTcpDoneTimeout *int `pulumi:"rawTcpDoneTimeout"`
	// Allows for restricting this input to only accept data from the host specified here.
	RestrictToHost *string `pulumi:"restrictToHost"`
	// Sets the source key/field for events from this input. Defaults to the input file path.
	// Sets the source key initial value. The key is used during parsing/indexing, in particular to set the source field during indexing. It is also the source field used at search time. As a convenience, the chosen string is prepended with 'source::'.
	Source *string `pulumi:"source"`
	// Set the source type for events from this input.
	// "sourcetype=" is automatically prepended to <string>.
	// Defaults to audittrail (if signedaudit=true) or fschange (if signedaudit=false).
	Sourcetype *string `pulumi:"sourcetype"`
}

// The set of arguments for constructing a InputsTcpRaw resource.
type InputsTcpRawArgs struct {
	// The app/user context that is the namespace for the resource
	Acl InputsTcpRawAclPtrInput
	// Valid values: (ip | dns | none)
	// Set the host for the remote server that is sending data.
	// ip sets the host to the IP address of the remote server sending data.
	// dns sets the host to the reverse DNS entry for the IP address of the remote server sending data.
	// none leaves the host as specified in inputs.conf, which is typically the Splunk system hostname.
	// Default value is dns.
	ConnectionHost pulumi.StringPtrInput
	// Indicates if input is disabled.
	Disabled pulumi.BoolPtrInput
	// Host from which the indexer gets data.
	Host pulumi.StringPtrInput
	// Index to store generated events. Defaults to default.
	Index pulumi.StringPtrInput
	// The input port which receives raw data.
	Name pulumi.StringPtrInput
	// Valid values: (parsingQueue | indexQueue)
	// Specifies where the input processor should deposit the events it reads. Defaults to parsingQueue.
	// Set queue to parsingQueue to apply props.conf and other parsing rules to your data. For more information about props.conf and rules for timestamping and linebreaking, refer to props.conf and the online documentation at "Monitor files and directories with inputs.conf"
	// Set queue to indexQueue to send your data directly into the index.
	Queue pulumi.StringPtrInput
	// Specifies in seconds the timeout value for adding a Done-key. Default value is 10 seconds.
	// If a connection over the port specified by name remains idle after receiving data for specified number of seconds, it adds a Done-key. This implies the last event is completely received.
	RawTcpDoneTimeout pulumi.IntPtrInput
	// Allows for restricting this input to only accept data from the host specified here.
	RestrictToHost pulumi.StringPtrInput
	// Sets the source key/field for events from this input. Defaults to the input file path.
	// Sets the source key initial value. The key is used during parsing/indexing, in particular to set the source field during indexing. It is also the source field used at search time. As a convenience, the chosen string is prepended with 'source::'.
	Source pulumi.StringPtrInput
	// Set the source type for events from this input.
	// "sourcetype=" is automatically prepended to <string>.
	// Defaults to audittrail (if signedaudit=true) or fschange (if signedaudit=false).
	Sourcetype pulumi.StringPtrInput
}

func (InputsTcpRawArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*inputsTcpRawArgs)(nil)).Elem()
}

type InputsTcpRawInput interface {
	pulumi.Input

	ToInputsTcpRawOutput() InputsTcpRawOutput
	ToInputsTcpRawOutputWithContext(ctx context.Context) InputsTcpRawOutput
}

func (*InputsTcpRaw) ElementType() reflect.Type {
	return reflect.TypeOf((**InputsTcpRaw)(nil)).Elem()
}

func (i *InputsTcpRaw) ToInputsTcpRawOutput() InputsTcpRawOutput {
	return i.ToInputsTcpRawOutputWithContext(context.Background())
}

func (i *InputsTcpRaw) ToInputsTcpRawOutputWithContext(ctx context.Context) InputsTcpRawOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InputsTcpRawOutput)
}

// InputsTcpRawArrayInput is an input type that accepts InputsTcpRawArray and InputsTcpRawArrayOutput values.
// You can construct a concrete instance of `InputsTcpRawArrayInput` via:
//
//          InputsTcpRawArray{ InputsTcpRawArgs{...} }
type InputsTcpRawArrayInput interface {
	pulumi.Input

	ToInputsTcpRawArrayOutput() InputsTcpRawArrayOutput
	ToInputsTcpRawArrayOutputWithContext(context.Context) InputsTcpRawArrayOutput
}

type InputsTcpRawArray []InputsTcpRawInput

func (InputsTcpRawArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*InputsTcpRaw)(nil)).Elem()
}

func (i InputsTcpRawArray) ToInputsTcpRawArrayOutput() InputsTcpRawArrayOutput {
	return i.ToInputsTcpRawArrayOutputWithContext(context.Background())
}

func (i InputsTcpRawArray) ToInputsTcpRawArrayOutputWithContext(ctx context.Context) InputsTcpRawArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InputsTcpRawArrayOutput)
}

// InputsTcpRawMapInput is an input type that accepts InputsTcpRawMap and InputsTcpRawMapOutput values.
// You can construct a concrete instance of `InputsTcpRawMapInput` via:
//
//          InputsTcpRawMap{ "key": InputsTcpRawArgs{...} }
type InputsTcpRawMapInput interface {
	pulumi.Input

	ToInputsTcpRawMapOutput() InputsTcpRawMapOutput
	ToInputsTcpRawMapOutputWithContext(context.Context) InputsTcpRawMapOutput
}

type InputsTcpRawMap map[string]InputsTcpRawInput

func (InputsTcpRawMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*InputsTcpRaw)(nil)).Elem()
}

func (i InputsTcpRawMap) ToInputsTcpRawMapOutput() InputsTcpRawMapOutput {
	return i.ToInputsTcpRawMapOutputWithContext(context.Background())
}

func (i InputsTcpRawMap) ToInputsTcpRawMapOutputWithContext(ctx context.Context) InputsTcpRawMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InputsTcpRawMapOutput)
}

type InputsTcpRawOutput struct{ *pulumi.OutputState }

func (InputsTcpRawOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**InputsTcpRaw)(nil)).Elem()
}

func (o InputsTcpRawOutput) ToInputsTcpRawOutput() InputsTcpRawOutput {
	return o
}

func (o InputsTcpRawOutput) ToInputsTcpRawOutputWithContext(ctx context.Context) InputsTcpRawOutput {
	return o
}

type InputsTcpRawArrayOutput struct{ *pulumi.OutputState }

func (InputsTcpRawArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*InputsTcpRaw)(nil)).Elem()
}

func (o InputsTcpRawArrayOutput) ToInputsTcpRawArrayOutput() InputsTcpRawArrayOutput {
	return o
}

func (o InputsTcpRawArrayOutput) ToInputsTcpRawArrayOutputWithContext(ctx context.Context) InputsTcpRawArrayOutput {
	return o
}

func (o InputsTcpRawArrayOutput) Index(i pulumi.IntInput) InputsTcpRawOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *InputsTcpRaw {
		return vs[0].([]*InputsTcpRaw)[vs[1].(int)]
	}).(InputsTcpRawOutput)
}

type InputsTcpRawMapOutput struct{ *pulumi.OutputState }

func (InputsTcpRawMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*InputsTcpRaw)(nil)).Elem()
}

func (o InputsTcpRawMapOutput) ToInputsTcpRawMapOutput() InputsTcpRawMapOutput {
	return o
}

func (o InputsTcpRawMapOutput) ToInputsTcpRawMapOutputWithContext(ctx context.Context) InputsTcpRawMapOutput {
	return o
}

func (o InputsTcpRawMapOutput) MapIndex(k pulumi.StringInput) InputsTcpRawOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *InputsTcpRaw {
		return vs[0].(map[string]*InputsTcpRaw)[vs[1].(string)]
	}).(InputsTcpRawOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*InputsTcpRawInput)(nil)).Elem(), &InputsTcpRaw{})
	pulumi.RegisterInputType(reflect.TypeOf((*InputsTcpRawArrayInput)(nil)).Elem(), InputsTcpRawArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*InputsTcpRawMapInput)(nil)).Elem(), InputsTcpRawMap{})
	pulumi.RegisterOutputType(InputsTcpRawOutput{})
	pulumi.RegisterOutputType(InputsTcpRawArrayOutput{})
	pulumi.RegisterOutputType(InputsTcpRawMapOutput{})
}
