// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * ## # Resource: splunk.ConfigsConf
 *
 * Create and manage configuration file stanzas.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as splunk from "@pulumi/splunk";
 *
 * const new_conf_stanza = new splunk.ConfigsConf("new-conf-stanza", {
 *     name: "custom-conf/custom",
 *     variables: {
 *         disabled: "false",
 *         custom_key: "value",
 *     },
 * });
 * ```
 */
export class ConfigsConf extends pulumi.CustomResource {
    /**
     * Get an existing ConfigsConf resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ConfigsConfState, opts?: pulumi.CustomResourceOptions): ConfigsConf {
        return new ConfigsConf(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'splunk:index/configsConf:ConfigsConf';

    /**
     * Returns true if the given object is an instance of ConfigsConf.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ConfigsConf {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ConfigsConf.__pulumiType;
    }

    public readonly acl!: pulumi.Output<outputs.ConfigsConfAcl>;
    /**
     * A '/' separated string consisting of {conf_file_name}/{stanza_name} ex. props/custom_stanza
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * A map of key value pairs for a stanza.
     */
    public readonly variables!: pulumi.Output<{[key: string]: string}>;

    /**
     * Create a ConfigsConf resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: ConfigsConfArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ConfigsConfArgs | ConfigsConfState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ConfigsConfState | undefined;
            resourceInputs["acl"] = state ? state.acl : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["variables"] = state ? state.variables : undefined;
        } else {
            const args = argsOrState as ConfigsConfArgs | undefined;
            resourceInputs["acl"] = args ? args.acl : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["variables"] = args ? args.variables : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(ConfigsConf.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ConfigsConf resources.
 */
export interface ConfigsConfState {
    acl?: pulumi.Input<inputs.ConfigsConfAcl>;
    /**
     * A '/' separated string consisting of {conf_file_name}/{stanza_name} ex. props/custom_stanza
     */
    name?: pulumi.Input<string>;
    /**
     * A map of key value pairs for a stanza.
     */
    variables?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}

/**
 * The set of arguments for constructing a ConfigsConf resource.
 */
export interface ConfigsConfArgs {
    acl?: pulumi.Input<inputs.ConfigsConfAcl>;
    /**
     * A '/' separated string consisting of {conf_file_name}/{stanza_name} ex. props/custom_stanza
     */
    name?: pulumi.Input<string>;
    /**
     * A map of key value pairs for a stanza.
     */
    variables?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
