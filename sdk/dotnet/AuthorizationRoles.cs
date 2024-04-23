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
    /// ## # Resource: splunk.AuthorizationRoles
    /// 
    /// Create and update role information.
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
    ///     var role01 = new Splunk.AuthorizationRoles("role01", new()
    ///     {
    ///         Name = "terraform-user01-role",
    ///         DefaultApp = "search",
    ///         ImportedRoles = new[]
    ///         {
    ///             "power",
    ///             "user",
    ///         },
    ///         Capabilities = new[]
    ///         {
    ///             "accelerate_datamodel",
    ///             "change_authentication",
    ///             "restart_splunkd",
    ///         },
    ///         SearchIndexesAlloweds = new[]
    ///         {
    ///             "_audit",
    ///             "_internal",
    ///             "main",
    ///         },
    ///         SearchIndexesDefaults = new[]
    ///         {
    ///             "_audit",
    ///             "_internal",
    ///             "main",
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// </summary>
    [SplunkResourceType("splunk:index/authorizationRoles:AuthorizationRoles")]
    public partial class AuthorizationRoles : global::Pulumi.CustomResource
    {
        /// <summary>
        /// List of capabilities assigned to role.
        /// </summary>
        [Output("capabilities")]
        public Output<ImmutableArray<string>> Capabilities { get; private set; } = null!;

        /// <summary>
        /// Maximum number of concurrently running real-time searches that all members of this role can have.
        /// </summary>
        [Output("cumulativeRealtimeSearchJobsQuota")]
        public Output<int?> CumulativeRealtimeSearchJobsQuota { get; private set; } = null!;

        /// <summary>
        /// Maximum number of concurrently running searches for all role members. Warning message logged when limit is reached.
        /// </summary>
        [Output("cumulativeSearchJobsQuota")]
        public Output<int?> CumulativeSearchJobsQuota { get; private set; } = null!;

        /// <summary>
        /// Specify the folder name of the default app to use for this role. A user-specific default app overrides this.
        /// </summary>
        [Output("defaultApp")]
        public Output<string> DefaultApp { get; private set; } = null!;

        /// <summary>
        /// List of imported roles for this role. &lt;br&gt;Importing other roles imports all aspects of that role, such as capabilities and allowed indexes to search. In combining multiple roles, the effective value for each attribute is value with the broadest permissions.
        /// </summary>
        [Output("importedRoles")]
        public Output<ImmutableArray<string>> ImportedRoles { get; private set; } = null!;

        /// <summary>
        /// The name of the user role to create.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Specify the maximum number of concurrent real-time search jobs for this role. This count is independent from the normal search jobs limit.
        /// </summary>
        [Output("realtimeSearchJobsQuota")]
        public Output<int?> RealtimeSearchJobsQuota { get; private set; } = null!;

        /// <summary>
        /// Specifies the maximum disk space in MB that can be used by a user's search jobs. For example, a value of 100 limits this role to 100 MB total.
        /// </summary>
        [Output("searchDiskQuota")]
        public Output<int?> SearchDiskQuota { get; private set; } = null!;

        /// <summary>
        /// Specify a search string that restricts the scope of searches run by this role. Search results for this role only show events that also match the search string you specify. In the case that a user has multiple roles with different search filters, they are combined with an OR.
        /// </summary>
        [Output("searchFilter")]
        public Output<string> SearchFilter { get; private set; } = null!;

        /// <summary>
        /// List of indexes that this role has permissions to search. These may be wildcarded, but the index name must begin with an underscore to match internal indexes.
        /// </summary>
        [Output("searchIndexesAlloweds")]
        public Output<ImmutableArray<string>> SearchIndexesAlloweds { get; private set; } = null!;

        /// <summary>
        /// List of indexes to search when no index is specified. These indexes can be wildcarded, with the exception that '*' does not match internal indexes. To match internal indexes, start with '_'. All internal indexes are represented by '_*'. A user with this role can search other indexes using "index= "
        /// </summary>
        [Output("searchIndexesDefaults")]
        public Output<ImmutableArray<string>> SearchIndexesDefaults { get; private set; } = null!;

        /// <summary>
        /// The maximum number of concurrent searches a user with this role is allowed to run. For users with multiple roles, the maximum quota value among all of the roles applies.
        /// </summary>
        [Output("searchJobsQuota")]
        public Output<int?> SearchJobsQuota { get; private set; } = null!;

        /// <summary>
        /// Maximum time span of a search, in seconds. By default, searches are not limited to any specific time window. To override any search time windows from imported roles, set srchTimeWin to '0', as the 'admin' role does.
        /// </summary>
        [Output("searchTimeWin")]
        public Output<int?> SearchTimeWin { get; private set; } = null!;


        /// <summary>
        /// Create a AuthorizationRoles resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public AuthorizationRoles(string name, AuthorizationRolesArgs? args = null, CustomResourceOptions? options = null)
            : base("splunk:index/authorizationRoles:AuthorizationRoles", name, args ?? new AuthorizationRolesArgs(), MakeResourceOptions(options, ""))
        {
        }

        private AuthorizationRoles(string name, Input<string> id, AuthorizationRolesState? state = null, CustomResourceOptions? options = null)
            : base("splunk:index/authorizationRoles:AuthorizationRoles", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing AuthorizationRoles resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static AuthorizationRoles Get(string name, Input<string> id, AuthorizationRolesState? state = null, CustomResourceOptions? options = null)
        {
            return new AuthorizationRoles(name, id, state, options);
        }
    }

    public sealed class AuthorizationRolesArgs : global::Pulumi.ResourceArgs
    {
        [Input("capabilities")]
        private InputList<string>? _capabilities;

        /// <summary>
        /// List of capabilities assigned to role.
        /// </summary>
        public InputList<string> Capabilities
        {
            get => _capabilities ?? (_capabilities = new InputList<string>());
            set => _capabilities = value;
        }

        /// <summary>
        /// Maximum number of concurrently running real-time searches that all members of this role can have.
        /// </summary>
        [Input("cumulativeRealtimeSearchJobsQuota")]
        public Input<int>? CumulativeRealtimeSearchJobsQuota { get; set; }

        /// <summary>
        /// Maximum number of concurrently running searches for all role members. Warning message logged when limit is reached.
        /// </summary>
        [Input("cumulativeSearchJobsQuota")]
        public Input<int>? CumulativeSearchJobsQuota { get; set; }

        /// <summary>
        /// Specify the folder name of the default app to use for this role. A user-specific default app overrides this.
        /// </summary>
        [Input("defaultApp")]
        public Input<string>? DefaultApp { get; set; }

        [Input("importedRoles")]
        private InputList<string>? _importedRoles;

        /// <summary>
        /// List of imported roles for this role. &lt;br&gt;Importing other roles imports all aspects of that role, such as capabilities and allowed indexes to search. In combining multiple roles, the effective value for each attribute is value with the broadest permissions.
        /// </summary>
        public InputList<string> ImportedRoles
        {
            get => _importedRoles ?? (_importedRoles = new InputList<string>());
            set => _importedRoles = value;
        }

        /// <summary>
        /// The name of the user role to create.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Specify the maximum number of concurrent real-time search jobs for this role. This count is independent from the normal search jobs limit.
        /// </summary>
        [Input("realtimeSearchJobsQuota")]
        public Input<int>? RealtimeSearchJobsQuota { get; set; }

        /// <summary>
        /// Specifies the maximum disk space in MB that can be used by a user's search jobs. For example, a value of 100 limits this role to 100 MB total.
        /// </summary>
        [Input("searchDiskQuota")]
        public Input<int>? SearchDiskQuota { get; set; }

        /// <summary>
        /// Specify a search string that restricts the scope of searches run by this role. Search results for this role only show events that also match the search string you specify. In the case that a user has multiple roles with different search filters, they are combined with an OR.
        /// </summary>
        [Input("searchFilter")]
        public Input<string>? SearchFilter { get; set; }

        [Input("searchIndexesAlloweds")]
        private InputList<string>? _searchIndexesAlloweds;

        /// <summary>
        /// List of indexes that this role has permissions to search. These may be wildcarded, but the index name must begin with an underscore to match internal indexes.
        /// </summary>
        public InputList<string> SearchIndexesAlloweds
        {
            get => _searchIndexesAlloweds ?? (_searchIndexesAlloweds = new InputList<string>());
            set => _searchIndexesAlloweds = value;
        }

        [Input("searchIndexesDefaults")]
        private InputList<string>? _searchIndexesDefaults;

        /// <summary>
        /// List of indexes to search when no index is specified. These indexes can be wildcarded, with the exception that '*' does not match internal indexes. To match internal indexes, start with '_'. All internal indexes are represented by '_*'. A user with this role can search other indexes using "index= "
        /// </summary>
        public InputList<string> SearchIndexesDefaults
        {
            get => _searchIndexesDefaults ?? (_searchIndexesDefaults = new InputList<string>());
            set => _searchIndexesDefaults = value;
        }

        /// <summary>
        /// The maximum number of concurrent searches a user with this role is allowed to run. For users with multiple roles, the maximum quota value among all of the roles applies.
        /// </summary>
        [Input("searchJobsQuota")]
        public Input<int>? SearchJobsQuota { get; set; }

        /// <summary>
        /// Maximum time span of a search, in seconds. By default, searches are not limited to any specific time window. To override any search time windows from imported roles, set srchTimeWin to '0', as the 'admin' role does.
        /// </summary>
        [Input("searchTimeWin")]
        public Input<int>? SearchTimeWin { get; set; }

        public AuthorizationRolesArgs()
        {
        }
        public static new AuthorizationRolesArgs Empty => new AuthorizationRolesArgs();
    }

    public sealed class AuthorizationRolesState : global::Pulumi.ResourceArgs
    {
        [Input("capabilities")]
        private InputList<string>? _capabilities;

        /// <summary>
        /// List of capabilities assigned to role.
        /// </summary>
        public InputList<string> Capabilities
        {
            get => _capabilities ?? (_capabilities = new InputList<string>());
            set => _capabilities = value;
        }

        /// <summary>
        /// Maximum number of concurrently running real-time searches that all members of this role can have.
        /// </summary>
        [Input("cumulativeRealtimeSearchJobsQuota")]
        public Input<int>? CumulativeRealtimeSearchJobsQuota { get; set; }

        /// <summary>
        /// Maximum number of concurrently running searches for all role members. Warning message logged when limit is reached.
        /// </summary>
        [Input("cumulativeSearchJobsQuota")]
        public Input<int>? CumulativeSearchJobsQuota { get; set; }

        /// <summary>
        /// Specify the folder name of the default app to use for this role. A user-specific default app overrides this.
        /// </summary>
        [Input("defaultApp")]
        public Input<string>? DefaultApp { get; set; }

        [Input("importedRoles")]
        private InputList<string>? _importedRoles;

        /// <summary>
        /// List of imported roles for this role. &lt;br&gt;Importing other roles imports all aspects of that role, such as capabilities and allowed indexes to search. In combining multiple roles, the effective value for each attribute is value with the broadest permissions.
        /// </summary>
        public InputList<string> ImportedRoles
        {
            get => _importedRoles ?? (_importedRoles = new InputList<string>());
            set => _importedRoles = value;
        }

        /// <summary>
        /// The name of the user role to create.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Specify the maximum number of concurrent real-time search jobs for this role. This count is independent from the normal search jobs limit.
        /// </summary>
        [Input("realtimeSearchJobsQuota")]
        public Input<int>? RealtimeSearchJobsQuota { get; set; }

        /// <summary>
        /// Specifies the maximum disk space in MB that can be used by a user's search jobs. For example, a value of 100 limits this role to 100 MB total.
        /// </summary>
        [Input("searchDiskQuota")]
        public Input<int>? SearchDiskQuota { get; set; }

        /// <summary>
        /// Specify a search string that restricts the scope of searches run by this role. Search results for this role only show events that also match the search string you specify. In the case that a user has multiple roles with different search filters, they are combined with an OR.
        /// </summary>
        [Input("searchFilter")]
        public Input<string>? SearchFilter { get; set; }

        [Input("searchIndexesAlloweds")]
        private InputList<string>? _searchIndexesAlloweds;

        /// <summary>
        /// List of indexes that this role has permissions to search. These may be wildcarded, but the index name must begin with an underscore to match internal indexes.
        /// </summary>
        public InputList<string> SearchIndexesAlloweds
        {
            get => _searchIndexesAlloweds ?? (_searchIndexesAlloweds = new InputList<string>());
            set => _searchIndexesAlloweds = value;
        }

        [Input("searchIndexesDefaults")]
        private InputList<string>? _searchIndexesDefaults;

        /// <summary>
        /// List of indexes to search when no index is specified. These indexes can be wildcarded, with the exception that '*' does not match internal indexes. To match internal indexes, start with '_'. All internal indexes are represented by '_*'. A user with this role can search other indexes using "index= "
        /// </summary>
        public InputList<string> SearchIndexesDefaults
        {
            get => _searchIndexesDefaults ?? (_searchIndexesDefaults = new InputList<string>());
            set => _searchIndexesDefaults = value;
        }

        /// <summary>
        /// The maximum number of concurrent searches a user with this role is allowed to run. For users with multiple roles, the maximum quota value among all of the roles applies.
        /// </summary>
        [Input("searchJobsQuota")]
        public Input<int>? SearchJobsQuota { get; set; }

        /// <summary>
        /// Maximum time span of a search, in seconds. By default, searches are not limited to any specific time window. To override any search time windows from imported roles, set srchTimeWin to '0', as the 'admin' role does.
        /// </summary>
        [Input("searchTimeWin")]
        public Input<int>? SearchTimeWin { get; set; }

        public AuthorizationRolesState()
        {
        }
        public static new AuthorizationRolesState Empty => new AuthorizationRolesState();
    }
}
