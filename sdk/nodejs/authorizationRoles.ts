// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * ## # Resource: splunk.AuthorizationRoles
 *
 * Create and update role information.
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as splunk from "@pulumi/splunk";
 *
 * const role01 = new splunk.AuthorizationRoles("role01", {
 *     capabilities: [
 *         "accelerate_datamodel",
 *         "change_authentication",
 *         "restart_splunkd",
 *     ],
 *     defaultApp: "search",
 *     importedRoles: [
 *         "power",
 *         "user",
 *     ],
 *     searchIndexesAlloweds: [
 *         "_audit",
 *         "_internal",
 *         "main",
 *     ],
 *     searchIndexesDefaults: [
 *         "_audit",
 *         "_internal",
 *         "main",
 *     ],
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 */
export class AuthorizationRoles extends pulumi.CustomResource {
    /**
     * Get an existing AuthorizationRoles resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AuthorizationRolesState, opts?: pulumi.CustomResourceOptions): AuthorizationRoles {
        return new AuthorizationRoles(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'splunk:index/authorizationRoles:AuthorizationRoles';

    /**
     * Returns true if the given object is an instance of AuthorizationRoles.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is AuthorizationRoles {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === AuthorizationRoles.__pulumiType;
    }

    /**
     * List of capabilities assigned to role.
     */
    public readonly capabilities!: pulumi.Output<string[]>;
    /**
     * Maximum number of concurrently running real-time searches that all members of this role can have.
     */
    public readonly cumulativeRealtimeSearchJobsQuota!: pulumi.Output<number | undefined>;
    /**
     * Maximum number of concurrently running searches for all role members. Warning message logged when limit is reached.
     */
    public readonly cumulativeSearchJobsQuota!: pulumi.Output<number | undefined>;
    /**
     * Specify the folder name of the default app to use for this role. A user-specific default app overrides this.
     */
    public readonly defaultApp!: pulumi.Output<string>;
    /**
     * List of imported roles for this role. <br>Importing other roles imports all aspects of that role, such as capabilities and allowed indexes to search. In combining multiple roles, the effective value for each attribute is value with the broadest permissions.
     */
    public readonly importedRoles!: pulumi.Output<string[]>;
    /**
     * The name of the user role to create.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Specify the maximum number of concurrent real-time search jobs for this role. This count is independent from the normal search jobs limit.
     */
    public readonly realtimeSearchJobsQuota!: pulumi.Output<number | undefined>;
    /**
     * Specifies the maximum disk space in MB that can be used by a user's search jobs. For example, a value of 100 limits this role to 100 MB total.
     */
    public readonly searchDiskQuota!: pulumi.Output<number | undefined>;
    /**
     * Specify a search string that restricts the scope of searches run by this role. Search results for this role only show events that also match the search string you specify. In the case that a user has multiple roles with different search filters, they are combined with an OR.
     */
    public readonly searchFilter!: pulumi.Output<string>;
    /**
     * List of indexes that this role has permissions to search. These may be wildcarded, but the index name must begin with an underscore to match internal indexes.
     */
    public readonly searchIndexesAlloweds!: pulumi.Output<string[]>;
    /**
     * List of indexes to search when no index is specified. These indexes can be wildcarded, with the exception that '*' does not match internal indexes. To match internal indexes, start with '_'. All internal indexes are represented by '_*'. A user with this role can search other indexes using "index= "
     */
    public readonly searchIndexesDefaults!: pulumi.Output<string[]>;
    /**
     * The maximum number of concurrent searches a user with this role is allowed to run. For users with multiple roles, the maximum quota value among all of the roles applies.
     */
    public readonly searchJobsQuota!: pulumi.Output<number | undefined>;
    /**
     * Maximum time span of a search, in seconds. By default, searches are not limited to any specific time window. To override any search time windows from imported roles, set srchTimeWin to '0', as the 'admin' role does.
     */
    public readonly searchTimeWin!: pulumi.Output<number | undefined>;

    /**
     * Create a AuthorizationRoles resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: AuthorizationRolesArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AuthorizationRolesArgs | AuthorizationRolesState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as AuthorizationRolesState | undefined;
            resourceInputs["capabilities"] = state ? state.capabilities : undefined;
            resourceInputs["cumulativeRealtimeSearchJobsQuota"] = state ? state.cumulativeRealtimeSearchJobsQuota : undefined;
            resourceInputs["cumulativeSearchJobsQuota"] = state ? state.cumulativeSearchJobsQuota : undefined;
            resourceInputs["defaultApp"] = state ? state.defaultApp : undefined;
            resourceInputs["importedRoles"] = state ? state.importedRoles : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["realtimeSearchJobsQuota"] = state ? state.realtimeSearchJobsQuota : undefined;
            resourceInputs["searchDiskQuota"] = state ? state.searchDiskQuota : undefined;
            resourceInputs["searchFilter"] = state ? state.searchFilter : undefined;
            resourceInputs["searchIndexesAlloweds"] = state ? state.searchIndexesAlloweds : undefined;
            resourceInputs["searchIndexesDefaults"] = state ? state.searchIndexesDefaults : undefined;
            resourceInputs["searchJobsQuota"] = state ? state.searchJobsQuota : undefined;
            resourceInputs["searchTimeWin"] = state ? state.searchTimeWin : undefined;
        } else {
            const args = argsOrState as AuthorizationRolesArgs | undefined;
            resourceInputs["capabilities"] = args ? args.capabilities : undefined;
            resourceInputs["cumulativeRealtimeSearchJobsQuota"] = args ? args.cumulativeRealtimeSearchJobsQuota : undefined;
            resourceInputs["cumulativeSearchJobsQuota"] = args ? args.cumulativeSearchJobsQuota : undefined;
            resourceInputs["defaultApp"] = args ? args.defaultApp : undefined;
            resourceInputs["importedRoles"] = args ? args.importedRoles : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["realtimeSearchJobsQuota"] = args ? args.realtimeSearchJobsQuota : undefined;
            resourceInputs["searchDiskQuota"] = args ? args.searchDiskQuota : undefined;
            resourceInputs["searchFilter"] = args ? args.searchFilter : undefined;
            resourceInputs["searchIndexesAlloweds"] = args ? args.searchIndexesAlloweds : undefined;
            resourceInputs["searchIndexesDefaults"] = args ? args.searchIndexesDefaults : undefined;
            resourceInputs["searchJobsQuota"] = args ? args.searchJobsQuota : undefined;
            resourceInputs["searchTimeWin"] = args ? args.searchTimeWin : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(AuthorizationRoles.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering AuthorizationRoles resources.
 */
export interface AuthorizationRolesState {
    /**
     * List of capabilities assigned to role.
     */
    capabilities?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Maximum number of concurrently running real-time searches that all members of this role can have.
     */
    cumulativeRealtimeSearchJobsQuota?: pulumi.Input<number>;
    /**
     * Maximum number of concurrently running searches for all role members. Warning message logged when limit is reached.
     */
    cumulativeSearchJobsQuota?: pulumi.Input<number>;
    /**
     * Specify the folder name of the default app to use for this role. A user-specific default app overrides this.
     */
    defaultApp?: pulumi.Input<string>;
    /**
     * List of imported roles for this role. <br>Importing other roles imports all aspects of that role, such as capabilities and allowed indexes to search. In combining multiple roles, the effective value for each attribute is value with the broadest permissions.
     */
    importedRoles?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The name of the user role to create.
     */
    name?: pulumi.Input<string>;
    /**
     * Specify the maximum number of concurrent real-time search jobs for this role. This count is independent from the normal search jobs limit.
     */
    realtimeSearchJobsQuota?: pulumi.Input<number>;
    /**
     * Specifies the maximum disk space in MB that can be used by a user's search jobs. For example, a value of 100 limits this role to 100 MB total.
     */
    searchDiskQuota?: pulumi.Input<number>;
    /**
     * Specify a search string that restricts the scope of searches run by this role. Search results for this role only show events that also match the search string you specify. In the case that a user has multiple roles with different search filters, they are combined with an OR.
     */
    searchFilter?: pulumi.Input<string>;
    /**
     * List of indexes that this role has permissions to search. These may be wildcarded, but the index name must begin with an underscore to match internal indexes.
     */
    searchIndexesAlloweds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * List of indexes to search when no index is specified. These indexes can be wildcarded, with the exception that '*' does not match internal indexes. To match internal indexes, start with '_'. All internal indexes are represented by '_*'. A user with this role can search other indexes using "index= "
     */
    searchIndexesDefaults?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The maximum number of concurrent searches a user with this role is allowed to run. For users with multiple roles, the maximum quota value among all of the roles applies.
     */
    searchJobsQuota?: pulumi.Input<number>;
    /**
     * Maximum time span of a search, in seconds. By default, searches are not limited to any specific time window. To override any search time windows from imported roles, set srchTimeWin to '0', as the 'admin' role does.
     */
    searchTimeWin?: pulumi.Input<number>;
}

/**
 * The set of arguments for constructing a AuthorizationRoles resource.
 */
export interface AuthorizationRolesArgs {
    /**
     * List of capabilities assigned to role.
     */
    capabilities?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Maximum number of concurrently running real-time searches that all members of this role can have.
     */
    cumulativeRealtimeSearchJobsQuota?: pulumi.Input<number>;
    /**
     * Maximum number of concurrently running searches for all role members. Warning message logged when limit is reached.
     */
    cumulativeSearchJobsQuota?: pulumi.Input<number>;
    /**
     * Specify the folder name of the default app to use for this role. A user-specific default app overrides this.
     */
    defaultApp?: pulumi.Input<string>;
    /**
     * List of imported roles for this role. <br>Importing other roles imports all aspects of that role, such as capabilities and allowed indexes to search. In combining multiple roles, the effective value for each attribute is value with the broadest permissions.
     */
    importedRoles?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The name of the user role to create.
     */
    name?: pulumi.Input<string>;
    /**
     * Specify the maximum number of concurrent real-time search jobs for this role. This count is independent from the normal search jobs limit.
     */
    realtimeSearchJobsQuota?: pulumi.Input<number>;
    /**
     * Specifies the maximum disk space in MB that can be used by a user's search jobs. For example, a value of 100 limits this role to 100 MB total.
     */
    searchDiskQuota?: pulumi.Input<number>;
    /**
     * Specify a search string that restricts the scope of searches run by this role. Search results for this role only show events that also match the search string you specify. In the case that a user has multiple roles with different search filters, they are combined with an OR.
     */
    searchFilter?: pulumi.Input<string>;
    /**
     * List of indexes that this role has permissions to search. These may be wildcarded, but the index name must begin with an underscore to match internal indexes.
     */
    searchIndexesAlloweds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * List of indexes to search when no index is specified. These indexes can be wildcarded, with the exception that '*' does not match internal indexes. To match internal indexes, start with '_'. All internal indexes are represented by '_*'. A user with this role can search other indexes using "index= "
     */
    searchIndexesDefaults?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The maximum number of concurrent searches a user with this role is allowed to run. For users with multiple roles, the maximum quota value among all of the roles applies.
     */
    searchJobsQuota?: pulumi.Input<number>;
    /**
     * Maximum time span of a search, in seconds. By default, searches are not limited to any specific time window. To override any search time windows from imported roles, set srchTimeWin to '0', as the 'admin' role does.
     */
    searchTimeWin?: pulumi.Input<number>;
}
