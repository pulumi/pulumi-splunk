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
    /// ## # Resource: splunk.AdminSamlGroups
    /// 
    /// Manage external groups in an IdP response to internal Splunk roles.
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
    ///     var saml_group = new Splunk.AdminSamlGroups("saml-group", new()
    ///     {
    ///         Name = "mygroup",
    ///         Roles = new[]
    ///         {
    ///             "admin",
    ///             "power",
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// SAML groups can be imported using the id, e.g.
    /// 
    /// ```sh
    /// $ pulumi import splunk:index/adminSamlGroups:AdminSamlGroups saml-group mygroup
    /// ```
    /// </summary>
    [SplunkResourceType("splunk:index/adminSamlGroups:AdminSamlGroups")]
    public partial class AdminSamlGroups : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The name of the external group.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// List of internal roles assigned to the group.
        /// </summary>
        [Output("roles")]
        public Output<ImmutableArray<string>> Roles { get; private set; } = null!;


        /// <summary>
        /// Create a AdminSamlGroups resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public AdminSamlGroups(string name, AdminSamlGroupsArgs? args = null, CustomResourceOptions? options = null)
            : base("splunk:index/adminSamlGroups:AdminSamlGroups", name, args ?? new AdminSamlGroupsArgs(), MakeResourceOptions(options, ""))
        {
        }

        private AdminSamlGroups(string name, Input<string> id, AdminSamlGroupsState? state = null, CustomResourceOptions? options = null)
            : base("splunk:index/adminSamlGroups:AdminSamlGroups", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing AdminSamlGroups resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static AdminSamlGroups Get(string name, Input<string> id, AdminSamlGroupsState? state = null, CustomResourceOptions? options = null)
        {
            return new AdminSamlGroups(name, id, state, options);
        }
    }

    public sealed class AdminSamlGroupsArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the external group.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("roles")]
        private InputList<string>? _roles;

        /// <summary>
        /// List of internal roles assigned to the group.
        /// </summary>
        public InputList<string> Roles
        {
            get => _roles ?? (_roles = new InputList<string>());
            set => _roles = value;
        }

        public AdminSamlGroupsArgs()
        {
        }
        public static new AdminSamlGroupsArgs Empty => new AdminSamlGroupsArgs();
    }

    public sealed class AdminSamlGroupsState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the external group.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("roles")]
        private InputList<string>? _roles;

        /// <summary>
        /// List of internal roles assigned to the group.
        /// </summary>
        public InputList<string> Roles
        {
            get => _roles ?? (_roles = new InputList<string>());
            set => _roles = value;
        }

        public AdminSamlGroupsState()
        {
        }
        public static new AdminSamlGroupsState Empty => new AdminSamlGroupsState();
    }
}
