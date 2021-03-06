// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
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
 *     email: "user01@example.com",
 *     forceChangePass: false,
 *     password: "password01",
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
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as AuthenticationUsersState | undefined;
            inputs["defaultApp"] = state ? state.defaultApp : undefined;
            inputs["email"] = state ? state.email : undefined;
            inputs["forceChangePass"] = state ? state.forceChangePass : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["password"] = state ? state.password : undefined;
            inputs["realname"] = state ? state.realname : undefined;
            inputs["restartBackgroundJobs"] = state ? state.restartBackgroundJobs : undefined;
            inputs["roles"] = state ? state.roles : undefined;
            inputs["tz"] = state ? state.tz : undefined;
        } else {
            const args = argsOrState as AuthenticationUsersArgs | undefined;
            inputs["defaultApp"] = args ? args.defaultApp : undefined;
            inputs["email"] = args ? args.email : undefined;
            inputs["forceChangePass"] = args ? args.forceChangePass : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["password"] = args ? args.password : undefined;
            inputs["realname"] = args ? args.realname : undefined;
            inputs["restartBackgroundJobs"] = args ? args.restartBackgroundJobs : undefined;
            inputs["roles"] = args ? args.roles : undefined;
            inputs["tz"] = args ? args.tz : undefined;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(AuthenticationUsers.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering AuthenticationUsers resources.
 */
export interface AuthenticationUsersState {
    /**
     * User default app. Overrides the default app inherited from the user roles.
     */
    readonly defaultApp?: pulumi.Input<string>;
    /**
     * User email address.
     */
    readonly email?: pulumi.Input<string>;
    /**
     * Force user to change password indication
     */
    readonly forceChangePass?: pulumi.Input<boolean>;
    /**
     * Unique user login name.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * User login password.
     */
    readonly password?: pulumi.Input<string>;
    /**
     * Full user name.
     */
    readonly realname?: pulumi.Input<string>;
    /**
     * Restart background search job that has not completed when Splunk restarts indication.
     */
    readonly restartBackgroundJobs?: pulumi.Input<boolean>;
    /**
     * Role to assign to this user. At least one existing role is required.
     */
    readonly roles?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * User timezone.
     */
    readonly tz?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a AuthenticationUsers resource.
 */
export interface AuthenticationUsersArgs {
    /**
     * User default app. Overrides the default app inherited from the user roles.
     */
    readonly defaultApp?: pulumi.Input<string>;
    /**
     * User email address.
     */
    readonly email?: pulumi.Input<string>;
    /**
     * Force user to change password indication
     */
    readonly forceChangePass?: pulumi.Input<boolean>;
    /**
     * Unique user login name.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * User login password.
     */
    readonly password?: pulumi.Input<string>;
    /**
     * Full user name.
     */
    readonly realname?: pulumi.Input<string>;
    /**
     * Restart background search job that has not completed when Splunk restarts indication.
     */
    readonly restartBackgroundJobs?: pulumi.Input<boolean>;
    /**
     * Role to assign to this user. At least one existing role is required.
     */
    readonly roles?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * User timezone.
     */
    readonly tz?: pulumi.Input<string>;
}
