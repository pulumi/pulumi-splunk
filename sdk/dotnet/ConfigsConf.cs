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
    /// # Resource: splunk.ConfigsConf
    /// Create and manage configuration file stanzas.
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
    ///     var new_conf_stanza = new Splunk.ConfigsConf("new-conf-stanza", new()
    ///     {
    ///         Name = "custom-conf/custom",
    ///         Variables = 
    ///         {
    ///             { "disabled", "false" },
    ///             { "custom_key", "value" },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// </summary>
    [SplunkResourceType("splunk:index/configsConf:ConfigsConf")]
    public partial class ConfigsConf : global::Pulumi.CustomResource
    {
        [Output("acl")]
        public Output<Outputs.ConfigsConfAcl> Acl { get; private set; } = null!;

        /// <summary>
        /// A '/' separated string consisting of {conf_file_name}/{stanza_name} ex. props/custom_stanza
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// A map of key value pairs for a stanza.
        /// </summary>
        [Output("variables")]
        public Output<ImmutableDictionary<string, string>> Variables { get; private set; } = null!;


        /// <summary>
        /// Create a ConfigsConf resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ConfigsConf(string name, ConfigsConfArgs? args = null, CustomResourceOptions? options = null)
            : base("splunk:index/configsConf:ConfigsConf", name, args ?? new ConfigsConfArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ConfigsConf(string name, Input<string> id, ConfigsConfState? state = null, CustomResourceOptions? options = null)
            : base("splunk:index/configsConf:ConfigsConf", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ConfigsConf resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ConfigsConf Get(string name, Input<string> id, ConfigsConfState? state = null, CustomResourceOptions? options = null)
        {
            return new ConfigsConf(name, id, state, options);
        }
    }

    public sealed class ConfigsConfArgs : global::Pulumi.ResourceArgs
    {
        [Input("acl")]
        public Input<Inputs.ConfigsConfAclArgs>? Acl { get; set; }

        /// <summary>
        /// A '/' separated string consisting of {conf_file_name}/{stanza_name} ex. props/custom_stanza
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("variables")]
        private InputMap<string>? _variables;

        /// <summary>
        /// A map of key value pairs for a stanza.
        /// </summary>
        public InputMap<string> Variables
        {
            get => _variables ?? (_variables = new InputMap<string>());
            set => _variables = value;
        }

        public ConfigsConfArgs()
        {
        }
        public static new ConfigsConfArgs Empty => new ConfigsConfArgs();
    }

    public sealed class ConfigsConfState : global::Pulumi.ResourceArgs
    {
        [Input("acl")]
        public Input<Inputs.ConfigsConfAclGetArgs>? Acl { get; set; }

        /// <summary>
        /// A '/' separated string consisting of {conf_file_name}/{stanza_name} ex. props/custom_stanza
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("variables")]
        private InputMap<string>? _variables;

        /// <summary>
        /// A map of key value pairs for a stanza.
        /// </summary>
        public InputMap<string> Variables
        {
            get => _variables ?? (_variables = new InputMap<string>());
            set => _variables = value;
        }

        public ConfigsConfState()
        {
        }
        public static new ConfigsConfState Empty => new ConfigsConfState();
    }
}
