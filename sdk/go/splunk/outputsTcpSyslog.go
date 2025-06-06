// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package splunk

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-splunk/sdk/go/splunk/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## # Resource: OutputsTcpSyslog
//
// Access the configuration of a forwarded server configured to provide data in standard syslog format.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-splunk/sdk/go/splunk"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := splunk.NewOutputsTcpSyslog(ctx, "tcp_syslog", &splunk.OutputsTcpSyslogArgs{
//				Name:     pulumi.String("new-syslog"),
//				Server:   pulumi.String("new-host-1:1234"),
//				Priority: pulumi.Int(5),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
type OutputsTcpSyslog struct {
	pulumi.CustomResourceState

	// The app/user context that is the namespace for the resource
	Acl OutputsTcpSyslogAclOutput `pulumi:"acl"`
	// If true, disables global syslog settings.
	Disabled pulumi.BoolOutput `pulumi:"disabled"`
	// Name of the syslog output group. This is name used when creating syslog configuration in outputs.conf.
	Name pulumi.StringOutput `pulumi:"name"`
	// Sets syslog priority value. The priority value should specified as an integer. See $SPLUNK_HOME/etc/system/README/outputs.conf.spec for details.
	Priority pulumi.IntOutput `pulumi:"priority"`
	// host:port of the server where syslog data should be sent
	Server pulumi.StringOutput `pulumi:"server"`
	// Specifies a rule for handling data in addition to that provided by the "syslog" sourcetype. By default, there is no value for syslogSourceType.
	// <br>This string is used as a substring match against the sourcetype key. For example, if the string is set to 'syslog', then all source types containing the string "syslog" receives this special treatment.
	// To match a source type explicitly, use the pattern "sourcetype::sourcetype_name." For example
	// syslogSourcetype = sourcetype::apache_common
	// Data that is "syslog" or matches this setting is assumed to already be in syslog format.
	// Data that does not match the rules has a header, potentially a timestamp, and a hostname added to the front of the event. This is how Splunk software causes arbitrary log data to match syslog expectations.
	SyslogSourcetype pulumi.StringOutput `pulumi:"syslogSourcetype"`
	// Format of timestamp to add at start of the events to be forwarded.
	// The format is a strftime-style timestamp formatting string. See $SPLUNK_HOME/etc/system/README/outputs.conf.spec for details.
	TimestampFormat pulumi.StringOutput `pulumi:"timestampFormat"`
	// Protocol to use to send syslog data. Valid values: (tcp | udp ).
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewOutputsTcpSyslog registers a new resource with the given unique name, arguments, and options.
func NewOutputsTcpSyslog(ctx *pulumi.Context,
	name string, args *OutputsTcpSyslogArgs, opts ...pulumi.ResourceOption) (*OutputsTcpSyslog, error) {
	if args == nil {
		args = &OutputsTcpSyslogArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource OutputsTcpSyslog
	err := ctx.RegisterResource("splunk:index/outputsTcpSyslog:OutputsTcpSyslog", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetOutputsTcpSyslog gets an existing OutputsTcpSyslog resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetOutputsTcpSyslog(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *OutputsTcpSyslogState, opts ...pulumi.ResourceOption) (*OutputsTcpSyslog, error) {
	var resource OutputsTcpSyslog
	err := ctx.ReadResource("splunk:index/outputsTcpSyslog:OutputsTcpSyslog", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering OutputsTcpSyslog resources.
type outputsTcpSyslogState struct {
	// The app/user context that is the namespace for the resource
	Acl *OutputsTcpSyslogAcl `pulumi:"acl"`
	// If true, disables global syslog settings.
	Disabled *bool `pulumi:"disabled"`
	// Name of the syslog output group. This is name used when creating syslog configuration in outputs.conf.
	Name *string `pulumi:"name"`
	// Sets syslog priority value. The priority value should specified as an integer. See $SPLUNK_HOME/etc/system/README/outputs.conf.spec for details.
	Priority *int `pulumi:"priority"`
	// host:port of the server where syslog data should be sent
	Server *string `pulumi:"server"`
	// Specifies a rule for handling data in addition to that provided by the "syslog" sourcetype. By default, there is no value for syslogSourceType.
	// <br>This string is used as a substring match against the sourcetype key. For example, if the string is set to 'syslog', then all source types containing the string "syslog" receives this special treatment.
	// To match a source type explicitly, use the pattern "sourcetype::sourcetype_name." For example
	// syslogSourcetype = sourcetype::apache_common
	// Data that is "syslog" or matches this setting is assumed to already be in syslog format.
	// Data that does not match the rules has a header, potentially a timestamp, and a hostname added to the front of the event. This is how Splunk software causes arbitrary log data to match syslog expectations.
	SyslogSourcetype *string `pulumi:"syslogSourcetype"`
	// Format of timestamp to add at start of the events to be forwarded.
	// The format is a strftime-style timestamp formatting string. See $SPLUNK_HOME/etc/system/README/outputs.conf.spec for details.
	TimestampFormat *string `pulumi:"timestampFormat"`
	// Protocol to use to send syslog data. Valid values: (tcp | udp ).
	Type *string `pulumi:"type"`
}

type OutputsTcpSyslogState struct {
	// The app/user context that is the namespace for the resource
	Acl OutputsTcpSyslogAclPtrInput
	// If true, disables global syslog settings.
	Disabled pulumi.BoolPtrInput
	// Name of the syslog output group. This is name used when creating syslog configuration in outputs.conf.
	Name pulumi.StringPtrInput
	// Sets syslog priority value. The priority value should specified as an integer. See $SPLUNK_HOME/etc/system/README/outputs.conf.spec for details.
	Priority pulumi.IntPtrInput
	// host:port of the server where syslog data should be sent
	Server pulumi.StringPtrInput
	// Specifies a rule for handling data in addition to that provided by the "syslog" sourcetype. By default, there is no value for syslogSourceType.
	// <br>This string is used as a substring match against the sourcetype key. For example, if the string is set to 'syslog', then all source types containing the string "syslog" receives this special treatment.
	// To match a source type explicitly, use the pattern "sourcetype::sourcetype_name." For example
	// syslogSourcetype = sourcetype::apache_common
	// Data that is "syslog" or matches this setting is assumed to already be in syslog format.
	// Data that does not match the rules has a header, potentially a timestamp, and a hostname added to the front of the event. This is how Splunk software causes arbitrary log data to match syslog expectations.
	SyslogSourcetype pulumi.StringPtrInput
	// Format of timestamp to add at start of the events to be forwarded.
	// The format is a strftime-style timestamp formatting string. See $SPLUNK_HOME/etc/system/README/outputs.conf.spec for details.
	TimestampFormat pulumi.StringPtrInput
	// Protocol to use to send syslog data. Valid values: (tcp | udp ).
	Type pulumi.StringPtrInput
}

func (OutputsTcpSyslogState) ElementType() reflect.Type {
	return reflect.TypeOf((*outputsTcpSyslogState)(nil)).Elem()
}

type outputsTcpSyslogArgs struct {
	// The app/user context that is the namespace for the resource
	Acl *OutputsTcpSyslogAcl `pulumi:"acl"`
	// If true, disables global syslog settings.
	Disabled *bool `pulumi:"disabled"`
	// Name of the syslog output group. This is name used when creating syslog configuration in outputs.conf.
	Name *string `pulumi:"name"`
	// Sets syslog priority value. The priority value should specified as an integer. See $SPLUNK_HOME/etc/system/README/outputs.conf.spec for details.
	Priority *int `pulumi:"priority"`
	// host:port of the server where syslog data should be sent
	Server *string `pulumi:"server"`
	// Specifies a rule for handling data in addition to that provided by the "syslog" sourcetype. By default, there is no value for syslogSourceType.
	// <br>This string is used as a substring match against the sourcetype key. For example, if the string is set to 'syslog', then all source types containing the string "syslog" receives this special treatment.
	// To match a source type explicitly, use the pattern "sourcetype::sourcetype_name." For example
	// syslogSourcetype = sourcetype::apache_common
	// Data that is "syslog" or matches this setting is assumed to already be in syslog format.
	// Data that does not match the rules has a header, potentially a timestamp, and a hostname added to the front of the event. This is how Splunk software causes arbitrary log data to match syslog expectations.
	SyslogSourcetype *string `pulumi:"syslogSourcetype"`
	// Format of timestamp to add at start of the events to be forwarded.
	// The format is a strftime-style timestamp formatting string. See $SPLUNK_HOME/etc/system/README/outputs.conf.spec for details.
	TimestampFormat *string `pulumi:"timestampFormat"`
	// Protocol to use to send syslog data. Valid values: (tcp | udp ).
	Type *string `pulumi:"type"`
}

// The set of arguments for constructing a OutputsTcpSyslog resource.
type OutputsTcpSyslogArgs struct {
	// The app/user context that is the namespace for the resource
	Acl OutputsTcpSyslogAclPtrInput
	// If true, disables global syslog settings.
	Disabled pulumi.BoolPtrInput
	// Name of the syslog output group. This is name used when creating syslog configuration in outputs.conf.
	Name pulumi.StringPtrInput
	// Sets syslog priority value. The priority value should specified as an integer. See $SPLUNK_HOME/etc/system/README/outputs.conf.spec for details.
	Priority pulumi.IntPtrInput
	// host:port of the server where syslog data should be sent
	Server pulumi.StringPtrInput
	// Specifies a rule for handling data in addition to that provided by the "syslog" sourcetype. By default, there is no value for syslogSourceType.
	// <br>This string is used as a substring match against the sourcetype key. For example, if the string is set to 'syslog', then all source types containing the string "syslog" receives this special treatment.
	// To match a source type explicitly, use the pattern "sourcetype::sourcetype_name." For example
	// syslogSourcetype = sourcetype::apache_common
	// Data that is "syslog" or matches this setting is assumed to already be in syslog format.
	// Data that does not match the rules has a header, potentially a timestamp, and a hostname added to the front of the event. This is how Splunk software causes arbitrary log data to match syslog expectations.
	SyslogSourcetype pulumi.StringPtrInput
	// Format of timestamp to add at start of the events to be forwarded.
	// The format is a strftime-style timestamp formatting string. See $SPLUNK_HOME/etc/system/README/outputs.conf.spec for details.
	TimestampFormat pulumi.StringPtrInput
	// Protocol to use to send syslog data. Valid values: (tcp | udp ).
	Type pulumi.StringPtrInput
}

func (OutputsTcpSyslogArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*outputsTcpSyslogArgs)(nil)).Elem()
}

type OutputsTcpSyslogInput interface {
	pulumi.Input

	ToOutputsTcpSyslogOutput() OutputsTcpSyslogOutput
	ToOutputsTcpSyslogOutputWithContext(ctx context.Context) OutputsTcpSyslogOutput
}

func (*OutputsTcpSyslog) ElementType() reflect.Type {
	return reflect.TypeOf((**OutputsTcpSyslog)(nil)).Elem()
}

func (i *OutputsTcpSyslog) ToOutputsTcpSyslogOutput() OutputsTcpSyslogOutput {
	return i.ToOutputsTcpSyslogOutputWithContext(context.Background())
}

func (i *OutputsTcpSyslog) ToOutputsTcpSyslogOutputWithContext(ctx context.Context) OutputsTcpSyslogOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OutputsTcpSyslogOutput)
}

// OutputsTcpSyslogArrayInput is an input type that accepts OutputsTcpSyslogArray and OutputsTcpSyslogArrayOutput values.
// You can construct a concrete instance of `OutputsTcpSyslogArrayInput` via:
//
//	OutputsTcpSyslogArray{ OutputsTcpSyslogArgs{...} }
type OutputsTcpSyslogArrayInput interface {
	pulumi.Input

	ToOutputsTcpSyslogArrayOutput() OutputsTcpSyslogArrayOutput
	ToOutputsTcpSyslogArrayOutputWithContext(context.Context) OutputsTcpSyslogArrayOutput
}

type OutputsTcpSyslogArray []OutputsTcpSyslogInput

func (OutputsTcpSyslogArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*OutputsTcpSyslog)(nil)).Elem()
}

func (i OutputsTcpSyslogArray) ToOutputsTcpSyslogArrayOutput() OutputsTcpSyslogArrayOutput {
	return i.ToOutputsTcpSyslogArrayOutputWithContext(context.Background())
}

func (i OutputsTcpSyslogArray) ToOutputsTcpSyslogArrayOutputWithContext(ctx context.Context) OutputsTcpSyslogArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OutputsTcpSyslogArrayOutput)
}

// OutputsTcpSyslogMapInput is an input type that accepts OutputsTcpSyslogMap and OutputsTcpSyslogMapOutput values.
// You can construct a concrete instance of `OutputsTcpSyslogMapInput` via:
//
//	OutputsTcpSyslogMap{ "key": OutputsTcpSyslogArgs{...} }
type OutputsTcpSyslogMapInput interface {
	pulumi.Input

	ToOutputsTcpSyslogMapOutput() OutputsTcpSyslogMapOutput
	ToOutputsTcpSyslogMapOutputWithContext(context.Context) OutputsTcpSyslogMapOutput
}

type OutputsTcpSyslogMap map[string]OutputsTcpSyslogInput

func (OutputsTcpSyslogMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*OutputsTcpSyslog)(nil)).Elem()
}

func (i OutputsTcpSyslogMap) ToOutputsTcpSyslogMapOutput() OutputsTcpSyslogMapOutput {
	return i.ToOutputsTcpSyslogMapOutputWithContext(context.Background())
}

func (i OutputsTcpSyslogMap) ToOutputsTcpSyslogMapOutputWithContext(ctx context.Context) OutputsTcpSyslogMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OutputsTcpSyslogMapOutput)
}

type OutputsTcpSyslogOutput struct{ *pulumi.OutputState }

func (OutputsTcpSyslogOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**OutputsTcpSyslog)(nil)).Elem()
}

func (o OutputsTcpSyslogOutput) ToOutputsTcpSyslogOutput() OutputsTcpSyslogOutput {
	return o
}

func (o OutputsTcpSyslogOutput) ToOutputsTcpSyslogOutputWithContext(ctx context.Context) OutputsTcpSyslogOutput {
	return o
}

// The app/user context that is the namespace for the resource
func (o OutputsTcpSyslogOutput) Acl() OutputsTcpSyslogAclOutput {
	return o.ApplyT(func(v *OutputsTcpSyslog) OutputsTcpSyslogAclOutput { return v.Acl }).(OutputsTcpSyslogAclOutput)
}

// If true, disables global syslog settings.
func (o OutputsTcpSyslogOutput) Disabled() pulumi.BoolOutput {
	return o.ApplyT(func(v *OutputsTcpSyslog) pulumi.BoolOutput { return v.Disabled }).(pulumi.BoolOutput)
}

// Name of the syslog output group. This is name used when creating syslog configuration in outputs.conf.
func (o OutputsTcpSyslogOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *OutputsTcpSyslog) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Sets syslog priority value. The priority value should specified as an integer. See $SPLUNK_HOME/etc/system/README/outputs.conf.spec for details.
func (o OutputsTcpSyslogOutput) Priority() pulumi.IntOutput {
	return o.ApplyT(func(v *OutputsTcpSyslog) pulumi.IntOutput { return v.Priority }).(pulumi.IntOutput)
}

// host:port of the server where syslog data should be sent
func (o OutputsTcpSyslogOutput) Server() pulumi.StringOutput {
	return o.ApplyT(func(v *OutputsTcpSyslog) pulumi.StringOutput { return v.Server }).(pulumi.StringOutput)
}

// Specifies a rule for handling data in addition to that provided by the "syslog" sourcetype. By default, there is no value for syslogSourceType.
// <br>This string is used as a substring match against the sourcetype key. For example, if the string is set to 'syslog', then all source types containing the string "syslog" receives this special treatment.
// To match a source type explicitly, use the pattern "sourcetype::sourcetype_name." For example
// syslogSourcetype = sourcetype::apache_common
// Data that is "syslog" or matches this setting is assumed to already be in syslog format.
// Data that does not match the rules has a header, potentially a timestamp, and a hostname added to the front of the event. This is how Splunk software causes arbitrary log data to match syslog expectations.
func (o OutputsTcpSyslogOutput) SyslogSourcetype() pulumi.StringOutput {
	return o.ApplyT(func(v *OutputsTcpSyslog) pulumi.StringOutput { return v.SyslogSourcetype }).(pulumi.StringOutput)
}

// Format of timestamp to add at start of the events to be forwarded.
// The format is a strftime-style timestamp formatting string. See $SPLUNK_HOME/etc/system/README/outputs.conf.spec for details.
func (o OutputsTcpSyslogOutput) TimestampFormat() pulumi.StringOutput {
	return o.ApplyT(func(v *OutputsTcpSyslog) pulumi.StringOutput { return v.TimestampFormat }).(pulumi.StringOutput)
}

// Protocol to use to send syslog data. Valid values: (tcp | udp ).
func (o OutputsTcpSyslogOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *OutputsTcpSyslog) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

type OutputsTcpSyslogArrayOutput struct{ *pulumi.OutputState }

func (OutputsTcpSyslogArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*OutputsTcpSyslog)(nil)).Elem()
}

func (o OutputsTcpSyslogArrayOutput) ToOutputsTcpSyslogArrayOutput() OutputsTcpSyslogArrayOutput {
	return o
}

func (o OutputsTcpSyslogArrayOutput) ToOutputsTcpSyslogArrayOutputWithContext(ctx context.Context) OutputsTcpSyslogArrayOutput {
	return o
}

func (o OutputsTcpSyslogArrayOutput) Index(i pulumi.IntInput) OutputsTcpSyslogOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *OutputsTcpSyslog {
		return vs[0].([]*OutputsTcpSyslog)[vs[1].(int)]
	}).(OutputsTcpSyslogOutput)
}

type OutputsTcpSyslogMapOutput struct{ *pulumi.OutputState }

func (OutputsTcpSyslogMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*OutputsTcpSyslog)(nil)).Elem()
}

func (o OutputsTcpSyslogMapOutput) ToOutputsTcpSyslogMapOutput() OutputsTcpSyslogMapOutput {
	return o
}

func (o OutputsTcpSyslogMapOutput) ToOutputsTcpSyslogMapOutputWithContext(ctx context.Context) OutputsTcpSyslogMapOutput {
	return o
}

func (o OutputsTcpSyslogMapOutput) MapIndex(k pulumi.StringInput) OutputsTcpSyslogOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *OutputsTcpSyslog {
		return vs[0].(map[string]*OutputsTcpSyslog)[vs[1].(string)]
	}).(OutputsTcpSyslogOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*OutputsTcpSyslogInput)(nil)).Elem(), &OutputsTcpSyslog{})
	pulumi.RegisterInputType(reflect.TypeOf((*OutputsTcpSyslogArrayInput)(nil)).Elem(), OutputsTcpSyslogArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*OutputsTcpSyslogMapInput)(nil)).Elem(), OutputsTcpSyslogMap{})
	pulumi.RegisterOutputType(OutputsTcpSyslogOutput{})
	pulumi.RegisterOutputType(OutputsTcpSyslogArrayOutput{})
	pulumi.RegisterOutputType(OutputsTcpSyslogMapOutput{})
}
