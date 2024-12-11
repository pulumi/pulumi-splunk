// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * ## 
 *
 * # Resource: splunk.LookupDefinition
 *
 * Manage lookup definitions in Splunk. For more information on lookup definitions, refer to the official Splunk documentation: https://docs.splunk.com/Documentation/Splunk/latest/Knowledge/Aboutlookupsandfieldactions
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as splunk from "@pulumi/splunk";
 *
 * const example = new splunk.LookupDefinition("example", {
 *     name: "example_lookup_definition",
 *     filename: "example_lookup_file.csv",
 *     acl: {
 *         owner: "admin",
 *         app: "search",
 *         sharing: "app",
 *         reads: ["*"],
 *         writes: ["admin"],
 *     },
 * });
 * ```
 *
 * ## Validation Rules
 *
 * When `acl.sharing` is set to `user`, the `acl.read` and `acl.write` fields must not be explicitly set. Setting them will trigger a validation error.
 */
export class LookupDefinition extends pulumi.CustomResource {
    /**
     * Get an existing LookupDefinition resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: LookupDefinitionState, opts?: pulumi.CustomResourceOptions): LookupDefinition {
        return new LookupDefinition(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'splunk:index/lookupDefinition:LookupDefinition';

    /**
     * Returns true if the given object is an instance of LookupDefinition.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is LookupDefinition {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === LookupDefinition.__pulumiType;
    }

    /**
     * Defines the access control list (ACL) for the lookup definition. See acl.md for more details.
     */
    public readonly acl!: pulumi.Output<outputs.LookupDefinitionAcl>;
    /**
     * The filename for the lookup table, usually ending in `.csv`.
     */
    public readonly filename!: pulumi.Output<string>;
    /**
     * A unique name for the lookup definition within the app context.
     */
    public readonly name!: pulumi.Output<string>;

    /**
     * Create a LookupDefinition resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: LookupDefinitionArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: LookupDefinitionArgs | LookupDefinitionState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as LookupDefinitionState | undefined;
            resourceInputs["acl"] = state ? state.acl : undefined;
            resourceInputs["filename"] = state ? state.filename : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
        } else {
            const args = argsOrState as LookupDefinitionArgs | undefined;
            if ((!args || args.filename === undefined) && !opts.urn) {
                throw new Error("Missing required property 'filename'");
            }
            if ((!args || args.name === undefined) && !opts.urn) {
                throw new Error("Missing required property 'name'");
            }
            resourceInputs["acl"] = args ? args.acl : undefined;
            resourceInputs["filename"] = args ? args.filename : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(LookupDefinition.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering LookupDefinition resources.
 */
export interface LookupDefinitionState {
    /**
     * Defines the access control list (ACL) for the lookup definition. See acl.md for more details.
     */
    acl?: pulumi.Input<inputs.LookupDefinitionAcl>;
    /**
     * The filename for the lookup table, usually ending in `.csv`.
     */
    filename?: pulumi.Input<string>;
    /**
     * A unique name for the lookup definition within the app context.
     */
    name?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a LookupDefinition resource.
 */
export interface LookupDefinitionArgs {
    /**
     * Defines the access control list (ACL) for the lookup definition. See acl.md for more details.
     */
    acl?: pulumi.Input<inputs.LookupDefinitionAcl>;
    /**
     * The filename for the lookup table, usually ending in `.csv`.
     */
    filename: pulumi.Input<string>;
    /**
     * A unique name for the lookup definition within the app context.
     */
    name: pulumi.Input<string>;
}