// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package splunk

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
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
// 	"github.com/pulumi/pulumi-splunk/sdk/go/splunk/"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := splunk.NewOutputsTcpSyslog(ctx, "tcpSyslog", &splunk.OutputsTcpSyslogArgs{
// 			Priority: pulumi.Int(5),
// 			Server:   pulumi.String("new-host-1:1234"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
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
	return reflect.TypeOf((*OutputsTcpSyslog)(nil))
}

func (i *OutputsTcpSyslog) ToOutputsTcpSyslogOutput() OutputsTcpSyslogOutput {
	return i.ToOutputsTcpSyslogOutputWithContext(context.Background())
}

func (i *OutputsTcpSyslog) ToOutputsTcpSyslogOutputWithContext(ctx context.Context) OutputsTcpSyslogOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OutputsTcpSyslogOutput)
}

type OutputsTcpSyslogOutput struct {
	*pulumi.OutputState
}

func (OutputsTcpSyslogOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*OutputsTcpSyslog)(nil))
}

func (o OutputsTcpSyslogOutput) ToOutputsTcpSyslogOutput() OutputsTcpSyslogOutput {
	return o
}

func (o OutputsTcpSyslogOutput) ToOutputsTcpSyslogOutputWithContext(ctx context.Context) OutputsTcpSyslogOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(OutputsTcpSyslogOutput{})
}
