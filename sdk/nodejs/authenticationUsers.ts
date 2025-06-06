// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * ## # Resource: splunk.AuthenticationUsers
 *
 * Create and update user information or delete the user.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as splunk from "@pulumi/splunk";
 *
 * const user01 = new splunk.AuthenticationUsers("user01", {
 *     name: "user01",
 *     email: "user01@example.com",
 *     password: "password01",
 *     forceChangePass: false,
 *     roles: ["terraform-user01-role"],
 * });
 * ```
 */
export class AuthenticationUsers extends pulumi.CustomResource {
    /**
     * Get an existing AuthenticationUsers resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AuthenticationUsersState, opts?: pulumi.CustomResourceOptions): AuthenticationUsers {
        return new AuthenticationUsers(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'splunk:index/authenticationUsers:AuthenticationUsers';

    /**
     * Returns true if the given object is an instance of AuthenticationUsers.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is AuthenticationUsers {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === AuthenticationUsers.__pulumiType;
    }

    /**
     * User default app. Overrides the default app inherited from the user roles.
     */
    public readonly defaultApp!: pulumi.Output<string>;
    /**
     * User email address.
     */
    public readonly email!: pulumi.Output<string>;
    /**
     * Force user to change password indication
     */
    public readonly forceChangePass!: pulumi.Output<boolean | undefined>;
    /**
     * Unique user login name.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * User login password.
     */
    public readonly password!: pulumi.Output<string | undefined>;
    /**
     * Full user name.
     */
    public readonly realname!: pulumi.Output<string>;
    /**
     * Restart background search job that has not completed when Splunk restarts indication.
     */
    public readonly restartBackgroundJobs!: pulumi.Output<boolean>;
    /**
     * Role to assign to this user. At least one existing role is required.
     */
    public readonly roles!: pulumi.Output<string[]>;
    /**
     * User timezone.
     */
    public readonly tz!: pulumi.Output<string>;

    /**
     * Create a AuthenticationUsers resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: AuthenticationUsersArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AuthenticationUsersArgs | AuthenticationUsersState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as AuthenticationUsersState | undefined;
            resourceInputs["defaultApp"] = state ? state.defaultApp : undefined;
            resourceInputs["email"] = state ? state.email : undefined;
            resourceInputs["forceChangePass"] = state ? state.forceChangePass : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["password"] = state ? state.password : undefined;
            resourceInputs["realname"] = state ? state.realname : undefined;
            resourceInputs["restartBackgroundJobs"] = state ? state.restartBackgroundJobs : undefined;
            resourceInputs["roles"] = state ? state.roles : undefined;
            resourceInputs["tz"] = state ? state.tz : undefined;
        } else {
            const args = argsOrState as AuthenticationUsersArgs | undefined;
            resourceInputs["defaultApp"] = args ? args.defaultApp : undefined;
            resourceInputs["email"] = args ? args.email : undefined;
            resourceInputs["forceChangePass"] = args ? args.forceChangePass : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["password"] = args?.password ? pulumi.secret(args.password) : undefined;
            resourceInputs["realname"] = args ? args.realname : undefined;
            resourceInputs["restartBackgroundJobs"] = args ? args.restartBackgroundJobs : undefined;
            resourceInputs["roles"] = args ? args.roles : undefined;
            resourceInputs["tz"] = args ? args.tz : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["password"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(AuthenticationUsers.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering AuthenticationUsers resources.
 */
export interface AuthenticationUsersState {
    /**
     * User default app. Overrides the default app inherited from the user roles.
     */
    defaultApp?: pulumi.Input<string>;
    /**
     * User email address.
     */
    email?: pulumi.Input<string>;
    /**
     * Force user to change password indication
     */
    forceChangePass?: pulumi.Input<boolean>;
    /**
     * Unique user login name.
     */
    name?: pulumi.Input<string>;
    /**
     * User login password.
     */
    password?: pulumi.Input<string>;
    /**
     * Full user name.
     */
    realname?: pulumi.Input<string>;
    /**
     * Restart background search job that has not completed when Splunk restarts indication.
     */
    restartBackgroundJobs?: pulumi.Input<boolean>;
    /**
     * Role to assign to this user. At least one existing role is required.
     */
    roles?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * User timezone.
     */
    tz?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a AuthenticationUsers resource.
 */
export interface AuthenticationUsersArgs {
    /**
     * User default app. Overrides the default app inherited from the user roles.
     */
    defaultApp?: pulumi.Input<string>;
    /**
     * User email address.
     */
    email?: pulumi.Input<string>;
    /**
     * Force user to change password indication
     */
    forceChangePass?: pulumi.Input<boolean>;
    /**
     * Unique user login name.
     */
    name?: pulumi.Input<string>;
    /**
     * User login password.
     */
    password?: pulumi.Input<string>;
    /**
     * Full user name.
     */
    realname?: pulumi.Input<string>;
    /**
     * Restart background search job that has not completed when Splunk restarts indication.
     */
    restartBackgroundJobs?: pulumi.Input<boolean>;
    /**
     * Role to assign to this user. At least one existing role is required.
     */
    roles?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * User timezone.
     */
    tz?: pulumi.Input<string>;
}
