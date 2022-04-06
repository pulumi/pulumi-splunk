// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

/**
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as splunk from "@pulumi/splunk";
 *
 * const myApp = new splunk.GenericAcl("my_app", {
 *     acl: {
 *         // use app=system, owner=nobody when managing apps, as they have no owner or app context
 *         app: "system",
 *         owner: "nobody",
 *         reads: ["*"],
 *         writes: [
 *             "admin",
 *             "power",
 *         ],
 *     },
 *     // apps are managed via the apps/local/<app> endpoint
 *     path: "apps/local/my_app",
 * });
 * const myDashboard = new splunk.GenericAcl("my_dashboard", {
 *     acl: {
 *         app: "my_app",
 *         owner: "joe_user",
 *         reads: ["team_joe"],
 *         writes: ["team_joe"],
 *     },
 *     path: "data/ui/views/my_dashboard",
 * });
 * ```
 *
 * ## Import
 *
 * Generic ACL resources can be imported by specifying their owner, app, and path with a colon-delimited string as the ID
 *
 * ```sh
 *  $ pulumi import splunk:index/genericAcl:GenericAcl splunk_generic_acl <owner>:<app>:<path>
 * ```
 */
export class GenericAcl extends pulumi.CustomResource {
    /**
     * Get an existing GenericAcl resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: GenericAclState, opts?: pulumi.CustomResourceOptions): GenericAcl {
        return new GenericAcl(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'splunk:index/genericAcl:GenericAcl';

    /**
     * Returns true if the given object is an instance of GenericAcl.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is GenericAcl {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === GenericAcl.__pulumiType;
    }

    /**
     * The ACL to apply to the object, including app/owner to properly identify the object.
     * Though technically optional, it should be explicitly set for this resource to really be valid. Some objects, such as
     * apps, require specific values for app and owner. Consult the REST API documentation regarding which values to use for
     * app and owner for objects that don't fit in the normal namespace.
     */
    public readonly acl!: pulumi.Output<outputs.GenericAclAcl>;
    /**
     * REST API Endpoint path to the object, relative to servicesNS/<owner>/<app>
     */
    public readonly path!: pulumi.Output<string>;

    /**
     * Create a GenericAcl resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: GenericAclArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: GenericAclArgs | GenericAclState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as GenericAclState | undefined;
            resourceInputs["acl"] = state ? state.acl : undefined;
            resourceInputs["path"] = state ? state.path : undefined;
        } else {
            const args = argsOrState as GenericAclArgs | undefined;
            if ((!args || args.path === undefined) && !opts.urn) {
                throw new Error("Missing required property 'path'");
            }
            resourceInputs["acl"] = args ? args.acl : undefined;
            resourceInputs["path"] = args ? args.path : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(GenericAcl.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering GenericAcl resources.
 */
export interface GenericAclState {
    /**
     * The ACL to apply to the object, including app/owner to properly identify the object.
     * Though technically optional, it should be explicitly set for this resource to really be valid. Some objects, such as
     * apps, require specific values for app and owner. Consult the REST API documentation regarding which values to use for
     * app and owner for objects that don't fit in the normal namespace.
     */
    acl?: pulumi.Input<inputs.GenericAclAcl>;
    /**
     * REST API Endpoint path to the object, relative to servicesNS/<owner>/<app>
     */
    path?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a GenericAcl resource.
 */
export interface GenericAclArgs {
    /**
     * The ACL to apply to the object, including app/owner to properly identify the object.
     * Though technically optional, it should be explicitly set for this resource to really be valid. Some objects, such as
     * apps, require specific values for app and owner. Consult the REST API documentation regarding which values to use for
     * app and owner for objects that don't fit in the normal namespace.
     */
    acl?: pulumi.Input<inputs.GenericAclAcl>;
    /**
     * REST API Endpoint path to the object, relative to servicesNS/<owner>/<app>
     */
    path: pulumi.Input<string>;
}
