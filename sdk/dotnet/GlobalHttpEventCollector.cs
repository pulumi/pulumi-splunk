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
    /// ## # Resource: splunk.GlobalHttpEventCollector
    /// 
    /// Update Global HTTP Event Collector input configuration.
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
    ///     var http = new Splunk.GlobalHttpEventCollector("http", new()
    ///     {
    ///         Disabled = false,
    ///         EnableSsl = true,
    ///         Port = 8088,
    ///     });
    /// 
    /// });
    /// ```
    /// </summary>
    [SplunkResourceType("splunk:index/globalHttpEventCollector:GlobalHttpEventCollector")]
    public partial class GlobalHttpEventCollector : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Number of threads used by HTTP Input server.
        /// </summary>
        [Output("dedicatedIoThreads")]
        public Output<int> DedicatedIoThreads { get; private set; } = null!;

        /// <summary>
        /// Input disabled indicator.
        /// </summary>
        [Output("disabled")]
        public Output<bool> Disabled { get; private set; } = null!;

        /// <summary>
        /// Enable SSL protocol for HTTP data input. `true` = SSL enabled, `false` = SSL disabled.
        /// </summary>
        [Output("enableSsl")]
        public Output<bool> EnableSsl { get; private set; } = null!;

        /// <summary>
        /// Maximum number of simultaneous HTTP connections accepted. Adjusting this value may cause server performance issues and is not generally recommended. Possible values for this setting vary by OS.
        /// </summary>
        [Output("maxSockets")]
        public Output<int> MaxSockets { get; private set; } = null!;

        /// <summary>
        /// Maximum number of threads that can be used by active HTTP transactions. Adjusting this value may cause server performance issues and is not generally recommended. Possible values for this setting vary by OS.
        /// </summary>
        [Output("maxThreads")]
        public Output<int> MaxThreads { get; private set; } = null!;

        /// <summary>
        /// HTTP data input IP port.
        /// </summary>
        [Output("port")]
        public Output<int> Port { get; private set; } = null!;

        /// <summary>
        /// Indicates whether the event collector input writes its configuration to a deployment server repository. When this setting is set to 1 (enabled), the input writes its configuration to the directory specified as repositoryLocation in serverclass.conf.
        /// Copy the full contents of the splunk_httpinput app directory to this directory for the configuration to work. When enabled, only the tokens defined in the splunk_httpinput app in this repository are viewable and editable on the API and the Data Inputs page in Splunk Web. When disabled, the input writes its configuration to $SPLUNK_HOME/etc/apps by default. Defaults to 0 (disabled).
        /// </summary>
        [Output("useDeploymentServer")]
        public Output<int> UseDeploymentServer { get; private set; } = null!;


        /// <summary>
        /// Create a GlobalHttpEventCollector resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public GlobalHttpEventCollector(string name, GlobalHttpEventCollectorArgs? args = null, CustomResourceOptions? options = null)
            : base("splunk:index/globalHttpEventCollector:GlobalHttpEventCollector", name, args ?? new GlobalHttpEventCollectorArgs(), MakeResourceOptions(options, ""))
        {
        }

        private GlobalHttpEventCollector(string name, Input<string> id, GlobalHttpEventCollectorState? state = null, CustomResourceOptions? options = null)
            : base("splunk:index/globalHttpEventCollector:GlobalHttpEventCollector", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing GlobalHttpEventCollector resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static GlobalHttpEventCollector Get(string name, Input<string> id, GlobalHttpEventCollectorState? state = null, CustomResourceOptions? options = null)
        {
            return new GlobalHttpEventCollector(name, id, state, options);
        }
    }

    public sealed class GlobalHttpEventCollectorArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Number of threads used by HTTP Input server.
        /// </summary>
        [Input("dedicatedIoThreads")]
        public Input<int>? DedicatedIoThreads { get; set; }

        /// <summary>
        /// Input disabled indicator.
        /// </summary>
        [Input("disabled")]
        public Input<bool>? Disabled { get; set; }

        /// <summary>
        /// Enable SSL protocol for HTTP data input. `true` = SSL enabled, `false` = SSL disabled.
        /// </summary>
        [Input("enableSsl")]
        public Input<bool>? EnableSsl { get; set; }

        /// <summary>
        /// Maximum number of simultaneous HTTP connections accepted. Adjusting this value may cause server performance issues and is not generally recommended. Possible values for this setting vary by OS.
        /// </summary>
        [Input("maxSockets")]
        public Input<int>? MaxSockets { get; set; }

        /// <summary>
        /// Maximum number of threads that can be used by active HTTP transactions. Adjusting this value may cause server performance issues and is not generally recommended. Possible values for this setting vary by OS.
        /// </summary>
        [Input("maxThreads")]
        public Input<int>? MaxThreads { get; set; }

        /// <summary>
        /// HTTP data input IP port.
        /// </summary>
        [Input("port")]
        public Input<int>? Port { get; set; }

        /// <summary>
        /// Indicates whether the event collector input writes its configuration to a deployment server repository. When this setting is set to 1 (enabled), the input writes its configuration to the directory specified as repositoryLocation in serverclass.conf.
        /// Copy the full contents of the splunk_httpinput app directory to this directory for the configuration to work. When enabled, only the tokens defined in the splunk_httpinput app in this repository are viewable and editable on the API and the Data Inputs page in Splunk Web. When disabled, the input writes its configuration to $SPLUNK_HOME/etc/apps by default. Defaults to 0 (disabled).
        /// </summary>
        [Input("useDeploymentServer")]
        public Input<int>? UseDeploymentServer { get; set; }

        public GlobalHttpEventCollectorArgs()
        {
        }
        public static new GlobalHttpEventCollectorArgs Empty => new GlobalHttpEventCollectorArgs();
    }

    public sealed class GlobalHttpEventCollectorState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Number of threads used by HTTP Input server.
        /// </summary>
        [Input("dedicatedIoThreads")]
        public Input<int>? DedicatedIoThreads { get; set; }

        /// <summary>
        /// Input disabled indicator.
        /// </summary>
        [Input("disabled")]
        public Input<bool>? Disabled { get; set; }

        /// <summary>
        /// Enable SSL protocol for HTTP data input. `true` = SSL enabled, `false` = SSL disabled.
        /// </summary>
        [Input("enableSsl")]
        public Input<bool>? EnableSsl { get; set; }

        /// <summary>
        /// Maximum number of simultaneous HTTP connections accepted. Adjusting this value may cause server performance issues and is not generally recommended. Possible values for this setting vary by OS.
        /// </summary>
        [Input("maxSockets")]
        public Input<int>? MaxSockets { get; set; }

        /// <summary>
        /// Maximum number of threads that can be used by active HTTP transactions. Adjusting this value may cause server performance issues and is not generally recommended. Possible values for this setting vary by OS.
        /// </summary>
        [Input("maxThreads")]
        public Input<int>? MaxThreads { get; set; }

        /// <summary>
        /// HTTP data input IP port.
        /// </summary>
        [Input("port")]
        public Input<int>? Port { get; set; }

        /// <summary>
        /// Indicates whether the event collector input writes its configuration to a deployment server repository. When this setting is set to 1 (enabled), the input writes its configuration to the directory specified as repositoryLocation in serverclass.conf.
        /// Copy the full contents of the splunk_httpinput app directory to this directory for the configuration to work. When enabled, only the tokens defined in the splunk_httpinput app in this repository are viewable and editable on the API and the Data Inputs page in Splunk Web. When disabled, the input writes its configuration to $SPLUNK_HOME/etc/apps by default. Defaults to 0 (disabled).
        /// </summary>
        [Input("useDeploymentServer")]
        public Input<int>? UseDeploymentServer { get; set; }

        public GlobalHttpEventCollectorState()
        {
        }
        public static new GlobalHttpEventCollectorState Empty => new GlobalHttpEventCollectorState();
    }
}
