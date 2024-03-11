// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Splunk
{
    /// <summary>
    /// ## # Resource: splunk.OutputsTcpSyslog
    /// 
    /// Access the configuration of a forwarded server configured to provide data in standard syslog format.
    /// 
    /// ## Example Usage
    /// 
    /// &lt;!--Start PulumiCodeChooser --&gt;
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Splunk = Pulumi.Splunk;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var tcpSyslog = new Splunk.OutputsTcpSyslog("tcpSyslog", new()
    ///     {
    ///         Priority = 5,
    ///         Server = "new-host-1:1234",
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// </summary>
    [SplunkResourceType("splunk:index/outputsTcpSyslog:OutputsTcpSyslog")]
    public partial class OutputsTcpSyslog : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The app/user context that is the namespace for the resource
        /// </summary>
        [Output("acl")]
        public Output<Outputs.OutputsTcpSyslogAcl> Acl { get; private set; } = null!;

        /// <summary>
        /// If true, disables global syslog settings.
        /// </summary>
        [Output("disabled")]
        public Output<bool> Disabled { get; private set; } = null!;

        /// <summary>
        /// Name of the syslog output group. This is name used when creating syslog configuration in outputs.conf.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Sets syslog priority value. The priority value should specified as an integer. See $SPLUNK_HOME/etc/system/README/outputs.conf.spec for details.
        /// </summary>
        [Output("priority")]
        public Output<int> Priority { get; private set; } = null!;

        /// <summary>
        /// host:port of the server where syslog data should be sent
        /// </summary>
        [Output("server")]
        public Output<string> Server { get; private set; } = null!;

        /// <summary>
        /// Specifies a rule for handling data in addition to that provided by the "syslog" sourcetype. By default, there is no value for syslogSourceType.
        /// &lt;br&gt;This string is used as a substring match against the sourcetype key. For example, if the string is set to 'syslog', then all source types containing the string "syslog" receives this special treatment.
        /// To match a source type explicitly, use the pattern "sourcetype::sourcetype_name." For example
        /// syslogSourcetype = sourcetype::apache_common
        /// Data that is "syslog" or matches this setting is assumed to already be in syslog format.
        /// Data that does not match the rules has a header, potentially a timestamp, and a hostname added to the front of the event. This is how Splunk software causes arbitrary log data to match syslog expectations.
        /// </summary>
        [Output("syslogSourcetype")]
        public Output<string> SyslogSourcetype { get; private set; } = null!;

        /// <summary>
        /// Format of timestamp to add at start of the events to be forwarded.
        /// The format is a strftime-style timestamp formatting string. See $SPLUNK_HOME/etc/system/README/outputs.conf.spec for details.
        /// </summary>
        [Output("timestampFormat")]
        public Output<string> TimestampFormat { get; private set; } = null!;

        /// <summary>
        /// Protocol to use to send syslog data. Valid values: (tcp | udp ).
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a OutputsTcpSyslog resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public OutputsTcpSyslog(string name, OutputsTcpSyslogArgs? args = null, CustomResourceOptions? options = null)
            : base("splunk:index/outputsTcpSyslog:OutputsTcpSyslog", name, args ?? new OutputsTcpSyslogArgs(), MakeResourceOptions(options, ""))
        {
        }

        private OutputsTcpSyslog(string name, Input<string> id, OutputsTcpSyslogState? state = null, CustomResourceOptions? options = null)
            : base("splunk:index/outputsTcpSyslog:OutputsTcpSyslog", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing OutputsTcpSyslog resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static OutputsTcpSyslog Get(string name, Input<string> id, OutputsTcpSyslogState? state = null, CustomResourceOptions? options = null)
        {
            return new OutputsTcpSyslog(name, id, state, options);
        }
    }

    public sealed class OutputsTcpSyslogArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The app/user context that is the namespace for the resource
        /// </summary>
        [Input("acl")]
        public Input<Inputs.OutputsTcpSyslogAclArgs>? Acl { get; set; }

        /// <summary>
        /// If true, disables global syslog settings.
        /// </summary>
        [Input("disabled")]
        public Input<bool>? Disabled { get; set; }

        /// <summary>
        /// Name of the syslog output group. This is name used when creating syslog configuration in outputs.conf.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Sets syslog priority value. The priority value should specified as an integer. See $SPLUNK_HOME/etc/system/README/outputs.conf.spec for details.
        /// </summary>
        [Input("priority")]
        public Input<int>? Priority { get; set; }

        /// <summary>
        /// host:port of the server where syslog data should be sent
        /// </summary>
        [Input("server")]
        public Input<string>? Server { get; set; }

        /// <summary>
        /// Specifies a rule for handling data in addition to that provided by the "syslog" sourcetype. By default, there is no value for syslogSourceType.
        /// &lt;br&gt;This string is used as a substring match against the sourcetype key. For example, if the string is set to 'syslog', then all source types containing the string "syslog" receives this special treatment.
        /// To match a source type explicitly, use the pattern "sourcetype::sourcetype_name." For example
        /// syslogSourcetype = sourcetype::apache_common
        /// Data that is "syslog" or matches this setting is assumed to already be in syslog format.
        /// Data that does not match the rules has a header, potentially a timestamp, and a hostname added to the front of the event. This is how Splunk software causes arbitrary log data to match syslog expectations.
        /// </summary>
        [Input("syslogSourcetype")]
        public Input<string>? SyslogSourcetype { get; set; }

        /// <summary>
        /// Format of timestamp to add at start of the events to be forwarded.
        /// The format is a strftime-style timestamp formatting string. See $SPLUNK_HOME/etc/system/README/outputs.conf.spec for details.
        /// </summary>
        [Input("timestampFormat")]
        public Input<string>? TimestampFormat { get; set; }

        /// <summary>
        /// Protocol to use to send syslog data. Valid values: (tcp | udp ).
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        public OutputsTcpSyslogArgs()
        {
        }
        public static new OutputsTcpSyslogArgs Empty => new OutputsTcpSyslogArgs();
    }

    public sealed class OutputsTcpSyslogState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The app/user context that is the namespace for the resource
        /// </summary>
        [Input("acl")]
        public Input<Inputs.OutputsTcpSyslogAclGetArgs>? Acl { get; set; }

        /// <summary>
        /// If true, disables global syslog settings.
        /// </summary>
        [Input("disabled")]
        public Input<bool>? Disabled { get; set; }

        /// <summary>
        /// Name of the syslog output group. This is name used when creating syslog configuration in outputs.conf.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Sets syslog priority value. The priority value should specified as an integer. See $SPLUNK_HOME/etc/system/README/outputs.conf.spec for details.
        /// </summary>
        [Input("priority")]
        public Input<int>? Priority { get; set; }

        /// <summary>
        /// host:port of the server where syslog data should be sent
        /// </summary>
        [Input("server")]
        public Input<string>? Server { get; set; }

        /// <summary>
        /// Specifies a rule for handling data in addition to that provided by the "syslog" sourcetype. By default, there is no value for syslogSourceType.
        /// &lt;br&gt;This string is used as a substring match against the sourcetype key. For example, if the string is set to 'syslog', then all source types containing the string "syslog" receives this special treatment.
        /// To match a source type explicitly, use the pattern "sourcetype::sourcetype_name." For example
        /// syslogSourcetype = sourcetype::apache_common
        /// Data that is "syslog" or matches this setting is assumed to already be in syslog format.
        /// Data that does not match the rules has a header, potentially a timestamp, and a hostname added to the front of the event. This is how Splunk software causes arbitrary log data to match syslog expectations.
        /// </summary>
        [Input("syslogSourcetype")]
        public Input<string>? SyslogSourcetype { get; set; }

        /// <summary>
        /// Format of timestamp to add at start of the events to be forwarded.
        /// The format is a strftime-style timestamp formatting string. See $SPLUNK_HOME/etc/system/README/outputs.conf.spec for details.
        /// </summary>
        [Input("timestampFormat")]
        public Input<string>? TimestampFormat { get; set; }

        /// <summary>
        /// Protocol to use to send syslog data. Valid values: (tcp | udp ).
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        public OutputsTcpSyslogState()
        {
        }
        public static new OutputsTcpSyslogState Empty => new OutputsTcpSyslogState();
    }
}
