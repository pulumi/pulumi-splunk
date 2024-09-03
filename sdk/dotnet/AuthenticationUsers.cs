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
    /// # Resource: splunk.AuthenticationUsers
    /// Create and update user information or delete the user.
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
    ///     var user01 = new Splunk.AuthenticationUsers("user01", new()
    ///     {
    ///         Name = "user01",
    ///         Email = "user01@example.com",
    ///         Password = "password01",
    ///         ForceChangePass = false,
    ///         Roles = new[]
    ///         {
    ///             "terraform-user01-role",
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// </summary>
    [SplunkResourceType("splunk:index/authenticationUsers:AuthenticationUsers")]
    public partial class AuthenticationUsers : global::Pulumi.CustomResource
    {
        /// <summary>
        /// User default app. Overrides the default app inherited from the user roles.
        /// </summary>
        [Output("defaultApp")]
        public Output<string> DefaultApp { get; private set; } = null!;

        /// <summary>
        /// User email address.
        /// </summary>
        [Output("email")]
        public Output<string> Email { get; private set; } = null!;

        /// <summary>
        /// Force user to change password indication
        /// </summary>
        [Output("forceChangePass")]
        public Output<bool?> ForceChangePass { get; private set; } = null!;

        /// <summary>
        /// Unique user login name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// User login password.
        /// </summary>
        [Output("password")]
        public Output<string?> Password { get; private set; } = null!;

        /// <summary>
        /// Full user name.
        /// </summary>
        [Output("realname")]
        public Output<string> Realname { get; private set; } = null!;

        /// <summary>
        /// Restart background search job that has not completed when Splunk restarts indication.
        /// </summary>
        [Output("restartBackgroundJobs")]
        public Output<bool> RestartBackgroundJobs { get; private set; } = null!;

        /// <summary>
        /// Role to assign to this user. At least one existing role is required.
        /// </summary>
        [Output("roles")]
        public Output<ImmutableArray<string>> Roles { get; private set; } = null!;

        /// <summary>
        /// User timezone.
        /// </summary>
        [Output("tz")]
        public Output<string> Tz { get; private set; } = null!;


        /// <summary>
        /// Create a AuthenticationUsers resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public AuthenticationUsers(string name, AuthenticationUsersArgs? args = null, CustomResourceOptions? options = null)
            : base("splunk:index/authenticationUsers:AuthenticationUsers", name, args ?? new AuthenticationUsersArgs(), MakeResourceOptions(options, ""))
        {
        }

        private AuthenticationUsers(string name, Input<string> id, AuthenticationUsersState? state = null, CustomResourceOptions? options = null)
            : base("splunk:index/authenticationUsers:AuthenticationUsers", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing AuthenticationUsers resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static AuthenticationUsers Get(string name, Input<string> id, AuthenticationUsersState? state = null, CustomResourceOptions? options = null)
        {
            return new AuthenticationUsers(name, id, state, options);
        }
    }

    public sealed class AuthenticationUsersArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// User default app. Overrides the default app inherited from the user roles.
        /// </summary>
        [Input("defaultApp")]
        public Input<string>? DefaultApp { get; set; }

        /// <summary>
        /// User email address.
        /// </summary>
        [Input("email")]
        public Input<string>? Email { get; set; }

        /// <summary>
        /// Force user to change password indication
        /// </summary>
        [Input("forceChangePass")]
        public Input<bool>? ForceChangePass { get; set; }

        /// <summary>
        /// Unique user login name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("password")]
        private Input<string>? _password;

        /// <summary>
        /// User login password.
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
        /// Full user name.
        /// </summary>
        [Input("realname")]
        public Input<string>? Realname { get; set; }

        /// <summary>
        /// Restart background search job that has not completed when Splunk restarts indication.
        /// </summary>
        [Input("restartBackgroundJobs")]
        public Input<bool>? RestartBackgroundJobs { get; set; }

        [Input("roles")]
        private InputList<string>? _roles;

        /// <summary>
        /// Role to assign to this user. At least one existing role is required.
        /// </summary>
        public InputList<string> Roles
        {
            get => _roles ?? (_roles = new InputList<string>());
            set => _roles = value;
        }

        /// <summary>
        /// User timezone.
        /// </summary>
        [Input("tz")]
        public Input<string>? Tz { get; set; }

        public AuthenticationUsersArgs()
        {
        }
        public static new AuthenticationUsersArgs Empty => new AuthenticationUsersArgs();
    }

    public sealed class AuthenticationUsersState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// User default app. Overrides the default app inherited from the user roles.
        /// </summary>
        [Input("defaultApp")]
        public Input<string>? DefaultApp { get; set; }

        /// <summary>
        /// User email address.
        /// </summary>
        [Input("email")]
        public Input<string>? Email { get; set; }

        /// <summary>
        /// Force user to change password indication
        /// </summary>
        [Input("forceChangePass")]
        public Input<bool>? ForceChangePass { get; set; }

        /// <summary>
        /// Unique user login name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("password")]
        private Input<string>? _password;

        /// <summary>
        /// User login password.
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
        /// Full user name.
        /// </summary>
        [Input("realname")]
        public Input<string>? Realname { get; set; }

        /// <summary>
        /// Restart background search job that has not completed when Splunk restarts indication.
        /// </summary>
        [Input("restartBackgroundJobs")]
        public Input<bool>? RestartBackgroundJobs { get; set; }

        [Input("roles")]
        private InputList<string>? _roles;

        /// <summary>
        /// Role to assign to this user. At least one existing role is required.
        /// </summary>
        public InputList<string> Roles
        {
            get => _roles ?? (_roles = new InputList<string>());
            set => _roles = value;
        }

        /// <summary>
        /// User timezone.
        /// </summary>
        [Input("tz")]
        public Input<string>? Tz { get; set; }

        public AuthenticationUsersState()
        {
        }
        public static new AuthenticationUsersState Empty => new AuthenticationUsersState();
    }
}
