// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package splunk

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## # Resource: InputsMonitor
//
// Create or update a new file or directory monitor input.
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
// 		_, err := splunk.NewInputsMonitor(ctx, "monitor", &splunk.InputsMonitorArgs{
// 			Recursive:  pulumi.Bool(true),
// 			Sourcetype: pulumi.String("text"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type InputsMonitor struct {
	pulumi.CustomResourceState

	// The app/user context that is the namespace for the resource
	Acl InputsMonitorAclOutput `pulumi:"acl"`
	// Specify a regular expression for a file path. The file path that matches this regular expression is not indexed.
	Blacklist pulumi.StringOutput `pulumi:"blacklist"`
	// A string that modifies the file tracking identity for files in this input. The magic value <SOURCE> invokes special behavior.
	CrcSalt pulumi.StringOutput `pulumi:"crcSalt"`
	// Indicates if input monitoring is disabled.
	Disabled pulumi.BoolOutput `pulumi:"disabled"`
	// If set to true, files that are seen for the first time is read from the end.
	FollowTail pulumi.BoolOutput `pulumi:"followTail"`
	// The value to populate in the host field for events from this data input.
	Host pulumi.StringOutput `pulumi:"host"`
	// Specify a regular expression for a file path. If the path for a file matches this regular expression, the captured value is used to populate the host field for events from this data input. The regular expression must have one capture group.
	HostRegex pulumi.StringOutput `pulumi:"hostRegex"`
	// Use the specified slash-separate segment of the filepath as the host field value.
	HostSegment pulumi.IntOutput `pulumi:"hostSegment"`
	// Specify a time value. If the modification time of a file being monitored falls outside of this rolling time window, the file is no longer being monitored.
	IgnoreOlderThan pulumi.StringOutput `pulumi:"ignoreOlderThan"`
	// Which index events from this input should be stored in. Defaults to default.
	Index pulumi.StringOutput `pulumi:"index"`
	// The file or directory path to monitor on the system.
	Name pulumi.StringOutput `pulumi:"name"`
	// Setting this to false prevents monitoring of any subdirectories encountered within this data input.
	Recursive pulumi.BoolOutput `pulumi:"recursive"`
	// The value to populate in the source field for events from this data input. The same source should not be used for multiple data inputs.
	RenameSource pulumi.StringOutput `pulumi:"renameSource"`
	// The value to populate in the sourcetype field for incoming events.
	Sourcetype pulumi.StringOutput `pulumi:"sourcetype"`
	// When Splunk software reaches the end of a file that is being read, the file is kept open for a minimum of the number of seconds specified in this value. After this period has elapsed, the file is checked again for more data.
	TimeBeforeClose pulumi.IntOutput `pulumi:"timeBeforeClose"`
	// Specify a regular expression for a file path. Only file paths that match this regular expression are indexed.
	Whitelist pulumi.StringOutput `pulumi:"whitelist"`
}

// NewInputsMonitor registers a new resource with the given unique name, arguments, and options.
func NewInputsMonitor(ctx *pulumi.Context,
	name string, args *InputsMonitorArgs, opts ...pulumi.ResourceOption) (*InputsMonitor, error) {
	if args == nil {
		args = &InputsMonitorArgs{}
	}

	var resource InputsMonitor
	err := ctx.RegisterResource("splunk:index/inputsMonitor:InputsMonitor", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetInputsMonitor gets an existing InputsMonitor resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetInputsMonitor(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *InputsMonitorState, opts ...pulumi.ResourceOption) (*InputsMonitor, error) {
	var resource InputsMonitor
	err := ctx.ReadResource("splunk:index/inputsMonitor:InputsMonitor", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering InputsMonitor resources.
type inputsMonitorState struct {
	// The app/user context that is the namespace for the resource
	Acl *InputsMonitorAcl `pulumi:"acl"`
	// Specify a regular expression for a file path. The file path that matches this regular expression is not indexed.
	Blacklist *string `pulumi:"blacklist"`
	// A string that modifies the file tracking identity for files in this input. The magic value <SOURCE> invokes special behavior.
	CrcSalt *string `pulumi:"crcSalt"`
	// Indicates if input monitoring is disabled.
	Disabled *bool `pulumi:"disabled"`
	// If set to true, files that are seen for the first time is read from the end.
	FollowTail *bool `pulumi:"followTail"`
	// The value to populate in the host field for events from this data input.
	Host *string `pulumi:"host"`
	// Specify a regular expression for a file path. If the path for a file matches this regular expression, the captured value is used to populate the host field for events from this data input. The regular expression must have one capture group.
	HostRegex *string `pulumi:"hostRegex"`
	// Use the specified slash-separate segment of the filepath as the host field value.
	HostSegment *int `pulumi:"hostSegment"`
	// Specify a time value. If the modification time of a file being monitored falls outside of this rolling time window, the file is no longer being monitored.
	IgnoreOlderThan *string `pulumi:"ignoreOlderThan"`
	// Which index events from this input should be stored in. Defaults to default.
	Index *string `pulumi:"index"`
	// The file or directory path to monitor on the system.
	Name *string `pulumi:"name"`
	// Setting this to false prevents monitoring of any subdirectories encountered within this data input.
	Recursive *bool `pulumi:"recursive"`
	// The value to populate in the source field for events from this data input. The same source should not be used for multiple data inputs.
	RenameSource *string `pulumi:"renameSource"`
	// The value to populate in the sourcetype field for incoming events.
	Sourcetype *string `pulumi:"sourcetype"`
	// When Splunk software reaches the end of a file that is being read, the file is kept open for a minimum of the number of seconds specified in this value. After this period has elapsed, the file is checked again for more data.
	TimeBeforeClose *int `pulumi:"timeBeforeClose"`
	// Specify a regular expression for a file path. Only file paths that match this regular expression are indexed.
	Whitelist *string `pulumi:"whitelist"`
}

type InputsMonitorState struct {
	// The app/user context that is the namespace for the resource
	Acl InputsMonitorAclPtrInput
	// Specify a regular expression for a file path. The file path that matches this regular expression is not indexed.
	Blacklist pulumi.StringPtrInput
	// A string that modifies the file tracking identity for files in this input. The magic value <SOURCE> invokes special behavior.
	CrcSalt pulumi.StringPtrInput
	// Indicates if input monitoring is disabled.
	Disabled pulumi.BoolPtrInput
	// If set to true, files that are seen for the first time is read from the end.
	FollowTail pulumi.BoolPtrInput
	// The value to populate in the host field for events from this data input.
	Host pulumi.StringPtrInput
	// Specify a regular expression for a file path. If the path for a file matches this regular expression, the captured value is used to populate the host field for events from this data input. The regular expression must have one capture group.
	HostRegex pulumi.StringPtrInput
	// Use the specified slash-separate segment of the filepath as the host field value.
	HostSegment pulumi.IntPtrInput
	// Specify a time value. If the modification time of a file being monitored falls outside of this rolling time window, the file is no longer being monitored.
	IgnoreOlderThan pulumi.StringPtrInput
	// Which index events from this input should be stored in. Defaults to default.
	Index pulumi.StringPtrInput
	// The file or directory path to monitor on the system.
	Name pulumi.StringPtrInput
	// Setting this to false prevents monitoring of any subdirectories encountered within this data input.
	Recursive pulumi.BoolPtrInput
	// The value to populate in the source field for events from this data input. The same source should not be used for multiple data inputs.
	RenameSource pulumi.StringPtrInput
	// The value to populate in the sourcetype field for incoming events.
	Sourcetype pulumi.StringPtrInput
	// When Splunk software reaches the end of a file that is being read, the file is kept open for a minimum of the number of seconds specified in this value. After this period has elapsed, the file is checked again for more data.
	TimeBeforeClose pulumi.IntPtrInput
	// Specify a regular expression for a file path. Only file paths that match this regular expression are indexed.
	Whitelist pulumi.StringPtrInput
}

func (InputsMonitorState) ElementType() reflect.Type {
	return reflect.TypeOf((*inputsMonitorState)(nil)).Elem()
}

type inputsMonitorArgs struct {
	// The app/user context that is the namespace for the resource
	Acl *InputsMonitorAcl `pulumi:"acl"`
	// Specify a regular expression for a file path. The file path that matches this regular expression is not indexed.
	Blacklist *string `pulumi:"blacklist"`
	// A string that modifies the file tracking identity for files in this input. The magic value <SOURCE> invokes special behavior.
	CrcSalt *string `pulumi:"crcSalt"`
	// Indicates if input monitoring is disabled.
	Disabled *bool `pulumi:"disabled"`
	// If set to true, files that are seen for the first time is read from the end.
	FollowTail *bool `pulumi:"followTail"`
	// The value to populate in the host field for events from this data input.
	Host *string `pulumi:"host"`
	// Specify a regular expression for a file path. If the path for a file matches this regular expression, the captured value is used to populate the host field for events from this data input. The regular expression must have one capture group.
	HostRegex *string `pulumi:"hostRegex"`
	// Use the specified slash-separate segment of the filepath as the host field value.
	HostSegment *int `pulumi:"hostSegment"`
	// Specify a time value. If the modification time of a file being monitored falls outside of this rolling time window, the file is no longer being monitored.
	IgnoreOlderThan *string `pulumi:"ignoreOlderThan"`
	// Which index events from this input should be stored in. Defaults to default.
	Index *string `pulumi:"index"`
	// The file or directory path to monitor on the system.
	Name *string `pulumi:"name"`
	// Setting this to false prevents monitoring of any subdirectories encountered within this data input.
	Recursive *bool `pulumi:"recursive"`
	// The value to populate in the source field for events from this data input. The same source should not be used for multiple data inputs.
	RenameSource *string `pulumi:"renameSource"`
	// The value to populate in the sourcetype field for incoming events.
	Sourcetype *string `pulumi:"sourcetype"`
	// When Splunk software reaches the end of a file that is being read, the file is kept open for a minimum of the number of seconds specified in this value. After this period has elapsed, the file is checked again for more data.
	TimeBeforeClose *int `pulumi:"timeBeforeClose"`
	// Specify a regular expression for a file path. Only file paths that match this regular expression are indexed.
	Whitelist *string `pulumi:"whitelist"`
}

// The set of arguments for constructing a InputsMonitor resource.
type InputsMonitorArgs struct {
	// The app/user context that is the namespace for the resource
	Acl InputsMonitorAclPtrInput
	// Specify a regular expression for a file path. The file path that matches this regular expression is not indexed.
	Blacklist pulumi.StringPtrInput
	// A string that modifies the file tracking identity for files in this input. The magic value <SOURCE> invokes special behavior.
	CrcSalt pulumi.StringPtrInput
	// Indicates if input monitoring is disabled.
	Disabled pulumi.BoolPtrInput
	// If set to true, files that are seen for the first time is read from the end.
	FollowTail pulumi.BoolPtrInput
	// The value to populate in the host field for events from this data input.
	Host pulumi.StringPtrInput
	// Specify a regular expression for a file path. If the path for a file matches this regular expression, the captured value is used to populate the host field for events from this data input. The regular expression must have one capture group.
	HostRegex pulumi.StringPtrInput
	// Use the specified slash-separate segment of the filepath as the host field value.
	HostSegment pulumi.IntPtrInput
	// Specify a time value. If the modification time of a file being monitored falls outside of this rolling time window, the file is no longer being monitored.
	IgnoreOlderThan pulumi.StringPtrInput
	// Which index events from this input should be stored in. Defaults to default.
	Index pulumi.StringPtrInput
	// The file or directory path to monitor on the system.
	Name pulumi.StringPtrInput
	// Setting this to false prevents monitoring of any subdirectories encountered within this data input.
	Recursive pulumi.BoolPtrInput
	// The value to populate in the source field for events from this data input. The same source should not be used for multiple data inputs.
	RenameSource pulumi.StringPtrInput
	// The value to populate in the sourcetype field for incoming events.
	Sourcetype pulumi.StringPtrInput
	// When Splunk software reaches the end of a file that is being read, the file is kept open for a minimum of the number of seconds specified in this value. After this period has elapsed, the file is checked again for more data.
	TimeBeforeClose pulumi.IntPtrInput
	// Specify a regular expression for a file path. Only file paths that match this regular expression are indexed.
	Whitelist pulumi.StringPtrInput
}

func (InputsMonitorArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*inputsMonitorArgs)(nil)).Elem()
}

type InputsMonitorInput interface {
	pulumi.Input

	ToInputsMonitorOutput() InputsMonitorOutput
	ToInputsMonitorOutputWithContext(ctx context.Context) InputsMonitorOutput
}

func (*InputsMonitor) ElementType() reflect.Type {
	return reflect.TypeOf((*InputsMonitor)(nil))
}

func (i *InputsMonitor) ToInputsMonitorOutput() InputsMonitorOutput {
	return i.ToInputsMonitorOutputWithContext(context.Background())
}

func (i *InputsMonitor) ToInputsMonitorOutputWithContext(ctx context.Context) InputsMonitorOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InputsMonitorOutput)
}

func (i *InputsMonitor) ToInputsMonitorPtrOutput() InputsMonitorPtrOutput {
	return i.ToInputsMonitorPtrOutputWithContext(context.Background())
}

func (i *InputsMonitor) ToInputsMonitorPtrOutputWithContext(ctx context.Context) InputsMonitorPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InputsMonitorPtrOutput)
}

type InputsMonitorPtrInput interface {
	pulumi.Input

	ToInputsMonitorPtrOutput() InputsMonitorPtrOutput
	ToInputsMonitorPtrOutputWithContext(ctx context.Context) InputsMonitorPtrOutput
}

type inputsMonitorPtrType InputsMonitorArgs

func (*inputsMonitorPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**InputsMonitor)(nil))
}

func (i *inputsMonitorPtrType) ToInputsMonitorPtrOutput() InputsMonitorPtrOutput {
	return i.ToInputsMonitorPtrOutputWithContext(context.Background())
}

func (i *inputsMonitorPtrType) ToInputsMonitorPtrOutputWithContext(ctx context.Context) InputsMonitorPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InputsMonitorPtrOutput)
}

// InputsMonitorArrayInput is an input type that accepts InputsMonitorArray and InputsMonitorArrayOutput values.
// You can construct a concrete instance of `InputsMonitorArrayInput` via:
//
//          InputsMonitorArray{ InputsMonitorArgs{...} }
type InputsMonitorArrayInput interface {
	pulumi.Input

	ToInputsMonitorArrayOutput() InputsMonitorArrayOutput
	ToInputsMonitorArrayOutputWithContext(context.Context) InputsMonitorArrayOutput
}

type InputsMonitorArray []InputsMonitorInput

func (InputsMonitorArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*InputsMonitor)(nil)).Elem()
}

func (i InputsMonitorArray) ToInputsMonitorArrayOutput() InputsMonitorArrayOutput {
	return i.ToInputsMonitorArrayOutputWithContext(context.Background())
}

func (i InputsMonitorArray) ToInputsMonitorArrayOutputWithContext(ctx context.Context) InputsMonitorArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InputsMonitorArrayOutput)
}

// InputsMonitorMapInput is an input type that accepts InputsMonitorMap and InputsMonitorMapOutput values.
// You can construct a concrete instance of `InputsMonitorMapInput` via:
//
//          InputsMonitorMap{ "key": InputsMonitorArgs{...} }
type InputsMonitorMapInput interface {
	pulumi.Input

	ToInputsMonitorMapOutput() InputsMonitorMapOutput
	ToInputsMonitorMapOutputWithContext(context.Context) InputsMonitorMapOutput
}

type InputsMonitorMap map[string]InputsMonitorInput

func (InputsMonitorMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*InputsMonitor)(nil)).Elem()
}

func (i InputsMonitorMap) ToInputsMonitorMapOutput() InputsMonitorMapOutput {
	return i.ToInputsMonitorMapOutputWithContext(context.Background())
}

func (i InputsMonitorMap) ToInputsMonitorMapOutputWithContext(ctx context.Context) InputsMonitorMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InputsMonitorMapOutput)
}

type InputsMonitorOutput struct{ *pulumi.OutputState }

func (InputsMonitorOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*InputsMonitor)(nil))
}

func (o InputsMonitorOutput) ToInputsMonitorOutput() InputsMonitorOutput {
	return o
}

func (o InputsMonitorOutput) ToInputsMonitorOutputWithContext(ctx context.Context) InputsMonitorOutput {
	return o
}

func (o InputsMonitorOutput) ToInputsMonitorPtrOutput() InputsMonitorPtrOutput {
	return o.ToInputsMonitorPtrOutputWithContext(context.Background())
}

func (o InputsMonitorOutput) ToInputsMonitorPtrOutputWithContext(ctx context.Context) InputsMonitorPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v InputsMonitor) *InputsMonitor {
		return &v
	}).(InputsMonitorPtrOutput)
}

type InputsMonitorPtrOutput struct{ *pulumi.OutputState }

func (InputsMonitorPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**InputsMonitor)(nil))
}

func (o InputsMonitorPtrOutput) ToInputsMonitorPtrOutput() InputsMonitorPtrOutput {
	return o
}

func (o InputsMonitorPtrOutput) ToInputsMonitorPtrOutputWithContext(ctx context.Context) InputsMonitorPtrOutput {
	return o
}

func (o InputsMonitorPtrOutput) Elem() InputsMonitorOutput {
	return o.ApplyT(func(v *InputsMonitor) InputsMonitor {
		if v != nil {
			return *v
		}
		var ret InputsMonitor
		return ret
	}).(InputsMonitorOutput)
}

type InputsMonitorArrayOutput struct{ *pulumi.OutputState }

func (InputsMonitorArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]InputsMonitor)(nil))
}

func (o InputsMonitorArrayOutput) ToInputsMonitorArrayOutput() InputsMonitorArrayOutput {
	return o
}

func (o InputsMonitorArrayOutput) ToInputsMonitorArrayOutputWithContext(ctx context.Context) InputsMonitorArrayOutput {
	return o
}

func (o InputsMonitorArrayOutput) Index(i pulumi.IntInput) InputsMonitorOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) InputsMonitor {
		return vs[0].([]InputsMonitor)[vs[1].(int)]
	}).(InputsMonitorOutput)
}

type InputsMonitorMapOutput struct{ *pulumi.OutputState }

func (InputsMonitorMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]InputsMonitor)(nil))
}

func (o InputsMonitorMapOutput) ToInputsMonitorMapOutput() InputsMonitorMapOutput {
	return o
}

func (o InputsMonitorMapOutput) ToInputsMonitorMapOutputWithContext(ctx context.Context) InputsMonitorMapOutput {
	return o
}

func (o InputsMonitorMapOutput) MapIndex(k pulumi.StringInput) InputsMonitorOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) InputsMonitor {
		return vs[0].(map[string]InputsMonitor)[vs[1].(string)]
	}).(InputsMonitorOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*InputsMonitorInput)(nil)).Elem(), &InputsMonitor{})
	pulumi.RegisterInputType(reflect.TypeOf((*InputsMonitorPtrInput)(nil)).Elem(), &InputsMonitor{})
	pulumi.RegisterInputType(reflect.TypeOf((*InputsMonitorArrayInput)(nil)).Elem(), InputsMonitorArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*InputsMonitorMapInput)(nil)).Elem(), InputsMonitorMap{})
	pulumi.RegisterOutputType(InputsMonitorOutput{})
	pulumi.RegisterOutputType(InputsMonitorPtrOutput{})
	pulumi.RegisterOutputType(InputsMonitorArrayOutput{})
	pulumi.RegisterOutputType(InputsMonitorMapOutput{})
}
