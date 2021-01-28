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
    /// ## # Resource: splunk.InputsTcpSplunkTcpToken
    /// 
    /// Manage receiver access using tokens.
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Splunk = Pulumi.Splunk;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var tcpSplunkTcpToken = new Splunk.InputsTcpSplunkTcpToken("tcpSplunkTcpToken", new Splunk.InputsTcpSplunkTcpTokenArgs
    ///         {
    ///             Token = "D66C45B3-7C28-48A1-A13A-027914146501",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// </summary>
    [SplunkResourceType("splunk:index/inputsTcpSplunkTcpToken:InputsTcpSplunkTcpToken")]
    public partial class InputsTcpSplunkTcpToken : Pulumi.CustomResource
    {
        /// <summary>
        /// The app/user context that is the namespace for the resource
        /// </summary>
        [Output("acl")]
        public Output<Outputs.InputsTcpSplunkTcpTokenAcl> Acl { get; private set; } = null!;

        /// <summary>
        /// Required. Name for the token to create.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Optional. Token value to use. If unspecified, a token is generated automatically.
        /// </summary>
        [Output("token")]
        public Output<string> Token { get; private set; } = null!;


        /// <summary>
        /// Create a InputsTcpSplunkTcpToken resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public InputsTcpSplunkTcpToken(string name, InputsTcpSplunkTcpTokenArgs? args = null, CustomResourceOptions? options = null)
            : base("splunk:index/inputsTcpSplunkTcpToken:InputsTcpSplunkTcpToken", name, args ?? new InputsTcpSplunkTcpTokenArgs(), MakeResourceOptions(options, ""))
        {
        }

        private InputsTcpSplunkTcpToken(string name, Input<string> id, InputsTcpSplunkTcpTokenState? state = null, CustomResourceOptions? options = null)
            : base("splunk:index/inputsTcpSplunkTcpToken:InputsTcpSplunkTcpToken", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing InputsTcpSplunkTcpToken resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static InputsTcpSplunkTcpToken Get(string name, Input<string> id, InputsTcpSplunkTcpTokenState? state = null, CustomResourceOptions? options = null)
        {
            return new InputsTcpSplunkTcpToken(name, id, state, options);
        }
    }

    public sealed class InputsTcpSplunkTcpTokenArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The app/user context that is the namespace for the resource
        /// </summary>
        [Input("acl")]
        public Input<Inputs.InputsTcpSplunkTcpTokenAclArgs>? Acl { get; set; }

        /// <summary>
        /// Required. Name for the token to create.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Optional. Token value to use. If unspecified, a token is generated automatically.
        /// </summary>
        [Input("token")]
        public Input<string>? Token { get; set; }

        public InputsTcpSplunkTcpTokenArgs()
        {
        }
    }

    public sealed class InputsTcpSplunkTcpTokenState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The app/user context that is the namespace for the resource
        /// </summary>
        [Input("acl")]
        public Input<Inputs.InputsTcpSplunkTcpTokenAclGetArgs>? Acl { get; set; }

        /// <summary>
        /// Required. Name for the token to create.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Optional. Token value to use. If unspecified, a token is generated automatically.
        /// </summary>
        [Input("token")]
        public Input<string>? Token { get; set; }

        public InputsTcpSplunkTcpTokenState()
        {
        }
    }
}
