// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
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
 * const new_conf_stanza = new splunk.ConfigsConf("new-conf-stanza", {variables: {
 *     disabled: "false",
 *     custom_key: "value",
 * }});
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

    /**
     * The app/user context that is the namespace for the resource
     */
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
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ConfigsConfState | undefined;
            inputs["acl"] = state ? state.acl : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["variables"] = state ? state.variables : undefined;
        } else {
            const args = argsOrState as ConfigsConfArgs | undefined;
            inputs["acl"] = args ? args.acl : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["variables"] = args ? args.variables : undefined;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(ConfigsConf.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ConfigsConf resources.
 */
export interface ConfigsConfState {
    /**
     * The app/user context that is the namespace for the resource
     */
    readonly acl?: pulumi.Input<inputs.ConfigsConfAcl>;
    /**
     * A '/' separated string consisting of {conf_file_name}/{stanza_name} ex. props/custom_stanza
     */
    readonly name?: pulumi.Input<string>;
    /**
     * A map of key value pairs for a stanza.
     */
    readonly variables?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}

/**
 * The set of arguments for constructing a ConfigsConf resource.
 */
export interface ConfigsConfArgs {
    /**
     * The app/user context that is the namespace for the resource
     */
    readonly acl?: pulumi.Input<inputs.ConfigsConfAcl>;
    /**
     * A '/' separated string consisting of {conf_file_name}/{stanza_name} ex. props/custom_stanza
     */
    readonly name?: pulumi.Input<string>;
    /**
     * A map of key value pairs for a stanza.
     */
    readonly variables?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
