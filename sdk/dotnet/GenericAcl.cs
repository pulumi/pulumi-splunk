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
    ///     var myApp = new Splunk.GenericAcl("my_app", new()
    ///     {
    ///         Path = "apps/local/my_app",
    ///         Acl = new Splunk.Inputs.GenericAclAclArgs
    ///         {
    ///             App = "system",
    ///             Owner = "nobody",
    ///             Reads = new[]
    ///             {
    ///                 "*",
    ///             },
    ///             Writes = new[]
    ///             {
    ///                 "admin",
    ///                 "power",
    ///             },
    ///         },
    ///     });
    /// 
    ///     var myDashboard = new Splunk.GenericAcl("my_dashboard", new()
    ///     {
    ///         Path = "data/ui/views/my_dashboard",
    ///         Acl = new Splunk.Inputs.GenericAclAclArgs
    ///         {
    ///             App = "my_app",
    ///             Owner = "joe_user",
    ///             Reads = new[]
    ///             {
    ///                 "team_joe",
    ///             },
    ///             Writes = new[]
    ///             {
    ///                 "team_joe",
    ///             },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// 
    /// ## Import
    /// 
    /// Generic ACL resources can be imported by specifying their owner, app, and path with a colon-delimited string as the ID:
    /// 
    /// ```sh
    /// $ pulumi import splunk:index/genericAcl:GenericAcl splunk_generic_acl &lt;owner&gt;:&lt;app&gt;:&lt;path&gt;
    /// ```
    /// </summary>
    [SplunkResourceType("splunk:index/genericAcl:GenericAcl")]
    public partial class GenericAcl : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The ACL to apply to the object, including app/owner to properly identify the object.
        /// Though technically optional, it should be explicitly set for this resource to really be valid. Some objects, such as
        /// apps, require specific values for app and owner. Consult the REST API documentation regarding which values to use for
        /// app and owner for objects that don't fit in the normal namespace.
        /// </summary>
        [Output("acl")]
        public Output<Outputs.GenericAclAcl> Acl { get; private set; } = null!;

        /// <summary>
        /// REST API Endpoint path to the object, relative to servicesNS/&lt;owner&gt;/&lt;app&gt;
        /// </summary>
        [Output("path")]
        public Output<string> Path { get; private set; } = null!;


        /// <summary>
        /// Create a GenericAcl resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public GenericAcl(string name, GenericAclArgs args, CustomResourceOptions? options = null)
            : base("splunk:index/genericAcl:GenericAcl", name, args ?? new GenericAclArgs(), MakeResourceOptions(options, ""))
        {
        }

        private GenericAcl(string name, Input<string> id, GenericAclState? state = null, CustomResourceOptions? options = null)
            : base("splunk:index/genericAcl:GenericAcl", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing GenericAcl resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static GenericAcl Get(string name, Input<string> id, GenericAclState? state = null, CustomResourceOptions? options = null)
        {
            return new GenericAcl(name, id, state, options);
        }
    }

    public sealed class GenericAclArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ACL to apply to the object, including app/owner to properly identify the object.
        /// Though technically optional, it should be explicitly set for this resource to really be valid. Some objects, such as
        /// apps, require specific values for app and owner. Consult the REST API documentation regarding which values to use for
        /// app and owner for objects that don't fit in the normal namespace.
        /// </summary>
        [Input("acl")]
        public Input<Inputs.GenericAclAclArgs>? Acl { get; set; }

        /// <summary>
        /// REST API Endpoint path to the object, relative to servicesNS/&lt;owner&gt;/&lt;app&gt;
        /// </summary>
        [Input("path", required: true)]
        public Input<string> Path { get; set; } = null!;

        public GenericAclArgs()
        {
        }
        public static new GenericAclArgs Empty => new GenericAclArgs();
    }

    public sealed class GenericAclState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ACL to apply to the object, including app/owner to properly identify the object.
        /// Though technically optional, it should be explicitly set for this resource to really be valid. Some objects, such as
        /// apps, require specific values for app and owner. Consult the REST API documentation regarding which values to use for
        /// app and owner for objects that don't fit in the normal namespace.
        /// </summary>
        [Input("acl")]
        public Input<Inputs.GenericAclAclGetArgs>? Acl { get; set; }

        /// <summary>
        /// REST API Endpoint path to the object, relative to servicesNS/&lt;owner&gt;/&lt;app&gt;
        /// </summary>
        [Input("path")]
        public Input<string>? Path { get; set; }

        public GenericAclState()
        {
        }
        public static new GenericAclState Empty => new GenericAclState();
    }
}
