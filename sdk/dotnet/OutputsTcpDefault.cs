// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Splunk
{
    /// <summary>
    /// ## # Resource: splunk.OutputsTcpDefault
    /// 
    /// Manage to global tcpout properties.
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Splunk = Pulumi.Splunk;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var tcpDefault = new Splunk.OutputsTcpDefault("tcp_default", new()
    ///     {
    ///         Name = "tcpout",
    ///         Disabled = false,
    ///         DefaultGroup = "test-indexers",
    ///         DropEventsOnQueueFull = 60,
    ///         IndexAndForward = true,
    ///         SendCookedData = true,
    ///         MaxQueueSize = "100KB",
    ///     });
    /// 
    /// });
    /// ```
    /// </summary>
    [SplunkResourceType("splunk:index/outputsTcpDefault:OutputsTcpDefault")]
    public partial class OutputsTcpDefault : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The app/user context that is the namespace for the resource
        /// </summary>
        [Output("acl")]
        public Output<Outputs.OutputsTcpDefaultAcl> Acl { get; private set; } = null!;

        /// <summary>
        /// Comma-separated list of one or more target group names, specified later in [tcpout:&lt;target_group&gt;] stanzas of outputs.conf.spec file.
        /// The forwarder sends all data to the specified groups. If you do not want to forward data automatically, do not set this attribute. Can be overridden by an inputs.conf _TCP_ROUTING setting, which in turn can be overridden by a props.conf/transforms.conf modifier.
        /// </summary>
        [Output("defaultGroup")]
        public Output<string> DefaultGroup { get; private set; } = null!;

        /// <summary>
        /// Disables default tcpout settings
        /// </summary>
        [Output("disabled")]
        public Output<bool> Disabled { get; private set; } = null!;

        /// <summary>
        /// If set to a positive number, wait the specified number of seconds before throwing out all new events until the output queue has space. Defaults to -1 (do not drop events).
        /// &lt;br&gt;CAUTION: Do not set this value to a positive integer if you are monitoring files.
        /// Setting this to -1 or 0 causes the output queue to block when it gets full, which causes further blocking up the processing chain. If any target group queue is blocked, no more data reaches any other target group.
        /// Using auto load-balancing is the best way to minimize this condition, because, in that case, multiple receivers must be down (or jammed up) before queue blocking can occur.
        /// </summary>
        [Output("dropEventsOnQueueFull")]
        public Output<int> DropEventsOnQueueFull { get; private set; } = null!;

        /// <summary>
        /// How often (in seconds) to send a heartbeat packet to the receiving server.
        /// Heartbeats are only sent if sendCookedData=true. Defaults to 30 seconds.
        /// </summary>
        [Output("heartbeatFrequency")]
        public Output<int> HeartbeatFrequency { get; private set; } = null!;

        /// <summary>
        /// Specifies whether to index all data locally, in addition to forwarding it. Defaults to false.
        /// This is known as an "index-and-forward" configuration. This attribute is only available for heavy forwarders. It is available only at the top level [tcpout] stanza in outputs.conf. It cannot be overridden in a target group.
        /// </summary>
        [Output("indexAndForward")]
        public Output<bool> IndexAndForward { get; private set; } = null!;

        /// <summary>
        /// Specify an integer or integer[KB|MB|GB].
        /// &lt;br&gt;Sets the maximum size of the forwarder output queue. It also sets the maximum size of the wait queue to 3x this value, if you have enabled indexer acknowledgment (useACK=true).
        /// Although the wait queue and the output queues are both configured by this attribute, they are separate queues. The setting determines the maximum size of the queue in-memory (RAM) buffer.
        /// For heavy forwarders sending parsed data, maxQueueSize is the maximum number of events. Since events are typically much shorter than data blocks, the memory consumed by the queue on a parsing forwarder is likely to be much smaller than on a non-parsing forwarder, if you use this version of the setting.
        /// If specified as a lone integer (for example, maxQueueSize=100), maxQueueSize indicates the maximum number of queued events (for parsed data) or blocks of data (for unparsed data). A block of data is approximately 64KB. For non-parsing forwarders, such as universal forwarders, that send unparsed data, maxQueueSize is the maximum number of data blocks.
        /// If specified as an integer followed by KB, MB, or GB (for example, maxQueueSize=100MB), maxQueueSize indicates the maximum RAM allocated to the queue buffer. Defaults to 500KB (which means a maximum size of 500KB for the output queue and 1500KB for the wait queue, if any).
        /// </summary>
        [Output("maxQueueSize")]
        public Output<string> MaxQueueSize { get; private set; } = null!;

        /// <summary>
        /// Configuration to be edited. The only valid value is "tcpout".
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// If true, events are cooked (processed by Splunk software). If false, events are raw and untouched prior to sending. Defaults to true.
        /// Set to false if you are sending to a third-party system.
        /// </summary>
        [Output("sendCookedData")]
        public Output<bool> SendCookedData { get; private set; } = null!;


        /// <summary>
        /// Create a OutputsTcpDefault resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public OutputsTcpDefault(string name, OutputsTcpDefaultArgs? args = null, CustomResourceOptions? options = null)
            : base("splunk:index/outputsTcpDefault:OutputsTcpDefault", name, args ?? new OutputsTcpDefaultArgs(), MakeResourceOptions(options, ""))
        {
        }

        private OutputsTcpDefault(string name, Input<string> id, OutputsTcpDefaultState? state = null, CustomResourceOptions? options = null)
            : base("splunk:index/outputsTcpDefault:OutputsTcpDefault", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing OutputsTcpDefault resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static OutputsTcpDefault Get(string name, Input<string> id, OutputsTcpDefaultState? state = null, CustomResourceOptions? options = null)
        {
            return new OutputsTcpDefault(name, id, state, options);
        }
    }

    public sealed class OutputsTcpDefaultArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The app/user context that is the namespace for the resource
        /// </summary>
        [Input("acl")]
        public Input<Inputs.OutputsTcpDefaultAclArgs>? Acl { get; set; }

        /// <summary>
        /// Comma-separated list of one or more target group names, specified later in [tcpout:&lt;target_group&gt;] stanzas of outputs.conf.spec file.
        /// The forwarder sends all data to the specified groups. If you do not want to forward data automatically, do not set this attribute. Can be overridden by an inputs.conf _TCP_ROUTING setting, which in turn can be overridden by a props.conf/transforms.conf modifier.
        /// </summary>
        [Input("defaultGroup")]
        public Input<string>? DefaultGroup { get; set; }

        /// <summary>
        /// Disables default tcpout settings
        /// </summary>
        [Input("disabled")]
        public Input<bool>? Disabled { get; set; }

        /// <summary>
        /// If set to a positive number, wait the specified number of seconds before throwing out all new events until the output queue has space. Defaults to -1 (do not drop events).
        /// &lt;br&gt;CAUTION: Do not set this value to a positive integer if you are monitoring files.
        /// Setting this to -1 or 0 causes the output queue to block when it gets full, which causes further blocking up the processing chain. If any target group queue is blocked, no more data reaches any other target group.
        /// Using auto load-balancing is the best way to minimize this condition, because, in that case, multiple receivers must be down (or jammed up) before queue blocking can occur.
        /// </summary>
        [Input("dropEventsOnQueueFull")]
        public Input<int>? DropEventsOnQueueFull { get; set; }

        /// <summary>
        /// How often (in seconds) to send a heartbeat packet to the receiving server.
        /// Heartbeats are only sent if sendCookedData=true. Defaults to 30 seconds.
        /// </summary>
        [Input("heartbeatFrequency")]
        public Input<int>? HeartbeatFrequency { get; set; }

        /// <summary>
        /// Specifies whether to index all data locally, in addition to forwarding it. Defaults to false.
        /// This is known as an "index-and-forward" configuration. This attribute is only available for heavy forwarders. It is available only at the top level [tcpout] stanza in outputs.conf. It cannot be overridden in a target group.
        /// </summary>
        [Input("indexAndForward")]
        public Input<bool>? IndexAndForward { get; set; }

        /// <summary>
        /// Specify an integer or integer[KB|MB|GB].
        /// &lt;br&gt;Sets the maximum size of the forwarder output queue. It also sets the maximum size of the wait queue to 3x this value, if you have enabled indexer acknowledgment (useACK=true).
        /// Although the wait queue and the output queues are both configured by this attribute, they are separate queues. The setting determines the maximum size of the queue in-memory (RAM) buffer.
        /// For heavy forwarders sending parsed data, maxQueueSize is the maximum number of events. Since events are typically much shorter than data blocks, the memory consumed by the queue on a parsing forwarder is likely to be much smaller than on a non-parsing forwarder, if you use this version of the setting.
        /// If specified as a lone integer (for example, maxQueueSize=100), maxQueueSize indicates the maximum number of queued events (for parsed data) or blocks of data (for unparsed data). A block of data is approximately 64KB. For non-parsing forwarders, such as universal forwarders, that send unparsed data, maxQueueSize is the maximum number of data blocks.
        /// If specified as an integer followed by KB, MB, or GB (for example, maxQueueSize=100MB), maxQueueSize indicates the maximum RAM allocated to the queue buffer. Defaults to 500KB (which means a maximum size of 500KB for the output queue and 1500KB for the wait queue, if any).
        /// </summary>
        [Input("maxQueueSize")]
        public Input<string>? MaxQueueSize { get; set; }

        /// <summary>
        /// Configuration to be edited. The only valid value is "tcpout".
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// If true, events are cooked (processed by Splunk software). If false, events are raw and untouched prior to sending. Defaults to true.
        /// Set to false if you are sending to a third-party system.
        /// </summary>
        [Input("sendCookedData")]
        public Input<bool>? SendCookedData { get; set; }

        public OutputsTcpDefaultArgs()
        {
        }
        public static new OutputsTcpDefaultArgs Empty => new OutputsTcpDefaultArgs();
    }

    public sealed class OutputsTcpDefaultState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The app/user context that is the namespace for the resource
        /// </summary>
        [Input("acl")]
        public Input<Inputs.OutputsTcpDefaultAclGetArgs>? Acl { get; set; }

        /// <summary>
        /// Comma-separated list of one or more target group names, specified later in [tcpout:&lt;target_group&gt;] stanzas of outputs.conf.spec file.
        /// The forwarder sends all data to the specified groups. If you do not want to forward data automatically, do not set this attribute. Can be overridden by an inputs.conf _TCP_ROUTING setting, which in turn can be overridden by a props.conf/transforms.conf modifier.
        /// </summary>
        [Input("defaultGroup")]
        public Input<string>? DefaultGroup { get; set; }

        /// <summary>
        /// Disables default tcpout settings
        /// </summary>
        [Input("disabled")]
        public Input<bool>? Disabled { get; set; }

        /// <summary>
        /// If set to a positive number, wait the specified number of seconds before throwing out all new events until the output queue has space. Defaults to -1 (do not drop events).
        /// &lt;br&gt;CAUTION: Do not set this value to a positive integer if you are monitoring files.
        /// Setting this to -1 or 0 causes the output queue to block when it gets full, which causes further blocking up the processing chain. If any target group queue is blocked, no more data reaches any other target group.
        /// Using auto load-balancing is the best way to minimize this condition, because, in that case, multiple receivers must be down (or jammed up) before queue blocking can occur.
        /// </summary>
        [Input("dropEventsOnQueueFull")]
        public Input<int>? DropEventsOnQueueFull { get; set; }

        /// <summary>
        /// How often (in seconds) to send a heartbeat packet to the receiving server.
        /// Heartbeats are only sent if sendCookedData=true. Defaults to 30 seconds.
        /// </summary>
        [Input("heartbeatFrequency")]
        public Input<int>? HeartbeatFrequency { get; set; }

        /// <summary>
        /// Specifies whether to index all data locally, in addition to forwarding it. Defaults to false.
        /// This is known as an "index-and-forward" configuration. This attribute is only available for heavy forwarders. It is available only at the top level [tcpout] stanza in outputs.conf. It cannot be overridden in a target group.
        /// </summary>
        [Input("indexAndForward")]
        public Input<bool>? IndexAndForward { get; set; }

        /// <summary>
        /// Specify an integer or integer[KB|MB|GB].
        /// &lt;br&gt;Sets the maximum size of the forwarder output queue. It also sets the maximum size of the wait queue to 3x this value, if you have enabled indexer acknowledgment (useACK=true).
        /// Although the wait queue and the output queues are both configured by this attribute, they are separate queues. The setting determines the maximum size of the queue in-memory (RAM) buffer.
        /// For heavy forwarders sending parsed data, maxQueueSize is the maximum number of events. Since events are typically much shorter than data blocks, the memory consumed by the queue on a parsing forwarder is likely to be much smaller than on a non-parsing forwarder, if you use this version of the setting.
        /// If specified as a lone integer (for example, maxQueueSize=100), maxQueueSize indicates the maximum number of queued events (for parsed data) or blocks of data (for unparsed data). A block of data is approximately 64KB. For non-parsing forwarders, such as universal forwarders, that send unparsed data, maxQueueSize is the maximum number of data blocks.
        /// If specified as an integer followed by KB, MB, or GB (for example, maxQueueSize=100MB), maxQueueSize indicates the maximum RAM allocated to the queue buffer. Defaults to 500KB (which means a maximum size of 500KB for the output queue and 1500KB for the wait queue, if any).
        /// </summary>
        [Input("maxQueueSize")]
        public Input<string>? MaxQueueSize { get; set; }

        /// <summary>
        /// Configuration to be edited. The only valid value is "tcpout".
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// If true, events are cooked (processed by Splunk software). If false, events are raw and untouched prior to sending. Defaults to true.
        /// Set to false if you are sending to a third-party system.
        /// </summary>
        [Input("sendCookedData")]
        public Input<bool>? SendCookedData { get; set; }

        public OutputsTcpDefaultState()
        {
        }
        public static new OutputsTcpDefaultState Empty => new OutputsTcpDefaultState();
    }
}
