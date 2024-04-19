// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * ## # Resource: splunk.ShIndexesManager
 *
 * Create indexes on Splunk Cloud instances. [BETA]
 *
 * ## Authorization and authentication
 *
 * As of now there is no support to create indexes in user-specified workspaces on Splunk Cloud.
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as splunk from "@pulumi/splunk";
 *
 * const tf_index = new splunk.ShIndexesManager("tf-index", {
 *     name: "tf-test-index-0",
 *     datatype: "event",
 *     frozenTimePeriodInSecs: "94608000",
 *     maxGlobalRawDataSizeMb: "100",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 */
export class ShIndexesManager extends pulumi.CustomResource {
    /**
     * Get an existing ShIndexesManager resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ShIndexesManagerState, opts?: pulumi.CustomResourceOptions): ShIndexesManager {
        return new ShIndexesManager(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'splunk:index/shIndexesManager:ShIndexesManager';

    /**
     * Returns true if the given object is an instance of ShIndexesManager.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ShIndexesManager {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ShIndexesManager.__pulumiType;
    }

    public readonly acl!: pulumi.Output<outputs.ShIndexesManagerAcl>;
    /**
     * Valid values: (event | metric). Specifies the type of index.
     */
    public readonly datatype!: pulumi.Output<string>;
    /**
     * Number of seconds after which indexed data rolls to frozen.
     * Defaults to 94608000 (3 years).Freezing data means it is removed from the index. If you need to archive your data, refer to coldToFrozenDir and coldToFrozenScript parameter documentation.
     */
    public readonly frozenTimePeriodInSecs!: pulumi.Output<string | undefined>;
    /**
     * The maximum size of an index (in MB). If an index grows larger than the maximum size, the oldest data is frozen.
     * Defaults to 100 MB.
     */
    public readonly maxGlobalRawDataSizeMb!: pulumi.Output<string | undefined>;
    /**
     * The name of the index to create.
     */
    public readonly name!: pulumi.Output<string>;

    /**
     * Create a ShIndexesManager resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: ShIndexesManagerArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ShIndexesManagerArgs | ShIndexesManagerState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ShIndexesManagerState | undefined;
            resourceInputs["acl"] = state ? state.acl : undefined;
            resourceInputs["datatype"] = state ? state.datatype : undefined;
            resourceInputs["frozenTimePeriodInSecs"] = state ? state.frozenTimePeriodInSecs : undefined;
            resourceInputs["maxGlobalRawDataSizeMb"] = state ? state.maxGlobalRawDataSizeMb : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
        } else {
            const args = argsOrState as ShIndexesManagerArgs | undefined;
            resourceInputs["acl"] = args ? args.acl : undefined;
            resourceInputs["datatype"] = args ? args.datatype : undefined;
            resourceInputs["frozenTimePeriodInSecs"] = args ? args.frozenTimePeriodInSecs : undefined;
            resourceInputs["maxGlobalRawDataSizeMb"] = args ? args.maxGlobalRawDataSizeMb : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(ShIndexesManager.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ShIndexesManager resources.
 */
export interface ShIndexesManagerState {
    acl?: pulumi.Input<inputs.ShIndexesManagerAcl>;
    /**
     * Valid values: (event | metric). Specifies the type of index.
     */
    datatype?: pulumi.Input<string>;
    /**
     * Number of seconds after which indexed data rolls to frozen.
     * Defaults to 94608000 (3 years).Freezing data means it is removed from the index. If you need to archive your data, refer to coldToFrozenDir and coldToFrozenScript parameter documentation.
     */
    frozenTimePeriodInSecs?: pulumi.Input<string>;
    /**
     * The maximum size of an index (in MB). If an index grows larger than the maximum size, the oldest data is frozen.
     * Defaults to 100 MB.
     */
    maxGlobalRawDataSizeMb?: pulumi.Input<string>;
    /**
     * The name of the index to create.
     */
    name?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ShIndexesManager resource.
 */
export interface ShIndexesManagerArgs {
    acl?: pulumi.Input<inputs.ShIndexesManagerAcl>;
    /**
     * Valid values: (event | metric). Specifies the type of index.
     */
    datatype?: pulumi.Input<string>;
    /**
     * Number of seconds after which indexed data rolls to frozen.
     * Defaults to 94608000 (3 years).Freezing data means it is removed from the index. If you need to archive your data, refer to coldToFrozenDir and coldToFrozenScript parameter documentation.
     */
    frozenTimePeriodInSecs?: pulumi.Input<string>;
    /**
     * The maximum size of an index (in MB). If an index grows larger than the maximum size, the oldest data is frozen.
     * Defaults to 100 MB.
     */
    maxGlobalRawDataSizeMb?: pulumi.Input<string>;
    /**
     * The name of the index to create.
     */
    name?: pulumi.Input<string>;
}
