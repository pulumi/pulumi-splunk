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
    /// ## # Resource: splunk.InputsTcpSsl
    /// 
    /// Access or update the SSL configuration for the host.
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
    ///     var test = new Splunk.InputsTcpSsl("test", new()
    ///     {
    ///         Disabled = false,
    ///         RequireClientCert = true,
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// </summary>
    [SplunkResourceType("splunk:index/inputsTcpSsl:InputsTcpSsl")]
    public partial class InputsTcpSsl : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Indicates if input is disabled.
        /// </summary>
        [Output("disabled")]
        public Output<bool> Disabled { get; private set; } = null!;

        /// <summary>
        /// Server certificate password, if any.
        /// </summary>
        [Output("password")]
        public Output<string> Password { get; private set; } = null!;

        /// <summary>
        /// Determines whether a client must authenticate.
        /// </summary>
        [Output("requireClientCert")]
        public Output<bool> RequireClientCert { get; private set; } = null!;

        /// <summary>
        /// Certificate authority list (root file)
        /// </summary>
        [Output("rootCa")]
        public Output<string> RootCa { get; private set; } = null!;

        /// <summary>
        /// Full path to the server certificate.
        /// </summary>
        [Output("serverCert")]
        public Output<string> ServerCert { get; private set; } = null!;


        /// <summary>
        /// Create a InputsTcpSsl resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public InputsTcpSsl(string name, InputsTcpSslArgs? args = null, CustomResourceOptions? options = null)
            : base("splunk:index/inputsTcpSsl:InputsTcpSsl", name, args ?? new InputsTcpSslArgs(), MakeResourceOptions(options, ""))
        {
        }

        private InputsTcpSsl(string name, Input<string> id, InputsTcpSslState? state = null, CustomResourceOptions? options = null)
            : base("splunk:index/inputsTcpSsl:InputsTcpSsl", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                AdditionalSecretOutputs =
                {
                    "password",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing InputsTcpSsl resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static InputsTcpSsl Get(string name, Input<string> id, InputsTcpSslState? state = null, CustomResourceOptions? options = null)
        {
            return new InputsTcpSsl(name, id, state, options);
        }
    }

    public sealed class InputsTcpSslArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Indicates if input is disabled.
        /// </summary>
        [Input("disabled")]
        public Input<bool>? Disabled { get; set; }

        [Input("password")]
        private Input<string>? _password;

        /// <summary>
        /// Server certificate password, if any.
        /// </summary>
        public Input<string>? Password
        {
            get => _password;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _password = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// Determines whether a client must authenticate.
        /// </summary>
        [Input("requireClientCert")]
        public Input<bool>? RequireClientCert { get; set; }

        /// <summary>
        /// Certificate authority list (root file)
        /// </summary>
        [Input("rootCa")]
        public Input<string>? RootCa { get; set; }

        /// <summary>
        /// Full path to the server certificate.
        /// </summary>
        [Input("serverCert")]
        public Input<string>? ServerCert { get; set; }

        public InputsTcpSslArgs()
        {
        }
        public static new InputsTcpSslArgs Empty => new InputsTcpSslArgs();
    }

    public sealed class InputsTcpSslState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Indicates if input is disabled.
        /// </summary>
        [Input("disabled")]
        public Input<bool>? Disabled { get; set; }

        [Input("password")]
        private Input<string>? _password;

        /// <summary>
        /// Server certificate password, if any.
        /// </summary>
        public Input<string>? Password
        {
            get => _password;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _password = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// Determines whether a client must authenticate.
        /// </summary>
        [Input("requireClientCert")]
        public Input<bool>? RequireClientCert { get; set; }

        /// <summary>
        /// Certificate authority list (root file)
        /// </summary>
        [Input("rootCa")]
        public Input<string>? RootCa { get; set; }

        /// <summary>
        /// Full path to the server certificate.
        /// </summary>
        [Input("serverCert")]
        public Input<string>? ServerCert { get; set; }

        public InputsTcpSslState()
        {
        }
        public static new InputsTcpSslState Empty => new InputsTcpSslState();
    }
}
