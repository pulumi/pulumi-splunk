// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * ## # Resource: splunk.InputsTcpCooked
 *
 * Create or update cooked TCP input information and create new containers for managing cooked data.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as splunk from "@pulumi/splunk";
 *
 * const tcpCooked = new splunk.InputsTcpCooked("tcp_cooked", {
 *     name: "50000",
 *     disabled: false,
 *     connectionHost: "dns",
 *     restrictToHost: "splunk",
 * });
 * ```
 */
export class InputsTcpCooked extends pulumi.CustomResource {
    /**
     * Get an existing InputsTcpCooked resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: InputsTcpCookedState, opts?: pulumi.CustomResourceOptions): InputsTcpCooked {
        return new InputsTcpCooked(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'splunk:index/inputsTcpCooked:InputsTcpCooked';

    /**
     * Returns true if the given object is an instance of InputsTcpCooked.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is InputsTcpCooked {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === InputsTcpCooked.__pulumiType;
    }

    /**
     * The app/user context that is the namespace for the resource
     */
    public readonly acl!: pulumi.Output<outputs.InputsTcpCookedAcl>;
    /**
     * Valid values: (ip | dns | none)
     * Set the host for the remote server that is sending data.
     * ip sets the host to the IP address of the remote server sending data.
     * dns sets the host to the reverse DNS entry for the IP address of the remote server sending data.
     * none leaves the host as specified in inputs.conf, which is typically the Splunk system hostname.
     * Default value is dns.
     */
    public readonly connectionHost!: pulumi.Output<string>;
    /**
     * Indicates if input is disabled.
     */
    public readonly disabled!: pulumi.Output<boolean>;
    /**
     * Host from which the indexer gets data.
     */
    public readonly host!: pulumi.Output<string>;
    /**
     * The port number of this input.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Restrict incoming connections on this port to the host specified here.
     */
    public readonly restrictToHost!: pulumi.Output<string>;

    /**
     * Create a InputsTcpCooked resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: InputsTcpCookedArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: InputsTcpCookedArgs | InputsTcpCookedState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as InputsTcpCookedState | undefined;
            resourceInputs["acl"] = state ? state.acl : undefined;
            resourceInputs["connectionHost"] = state ? state.connectionHost : undefined;
            resourceInputs["disabled"] = state ? state.disabled : undefined;
            resourceInputs["host"] = state ? state.host : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["restrictToHost"] = state ? state.restrictToHost : undefined;
        } else {
            const args = argsOrState as InputsTcpCookedArgs | undefined;
            resourceInputs["acl"] = args ? args.acl : undefined;
            resourceInputs["connectionHost"] = args ? args.connectionHost : undefined;
            resourceInputs["disabled"] = args ? args.disabled : undefined;
            resourceInputs["host"] = args ? args.host : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["restrictToHost"] = args ? args.restrictToHost : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(InputsTcpCooked.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering InputsTcpCooked resources.
 */
export interface InputsTcpCookedState {
    /**
     * The app/user context that is the namespace for the resource
     */
    acl?: pulumi.Input<inputs.InputsTcpCookedAcl>;
    /**
     * Valid values: (ip | dns | none)
     * Set the host for the remote server that is sending data.
     * ip sets the host to the IP address of the remote server sending data.
     * dns sets the host to the reverse DNS entry for the IP address of the remote server sending data.
     * none leaves the host as specified in inputs.conf, which is typically the Splunk system hostname.
     * Default value is dns.
     */
    connectionHost?: pulumi.Input<string>;
    /**
     * Indicates if input is disabled.
     */
    disabled?: pulumi.Input<boolean>;
    /**
     * Host from which the indexer gets data.
     */
    host?: pulumi.Input<string>;
    /**
     * The port number of this input.
     */
    name?: pulumi.Input<string>;
    /**
     * Restrict incoming connections on this port to the host specified here.
     */
    restrictToHost?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a InputsTcpCooked resource.
 */
export interface InputsTcpCookedArgs {
    /**
     * The app/user context that is the namespace for the resource
     */
    acl?: pulumi.Input<inputs.InputsTcpCookedAcl>;
    /**
     * Valid values: (ip | dns | none)
     * Set the host for the remote server that is sending data.
     * ip sets the host to the IP address of the remote server sending data.
     * dns sets the host to the reverse DNS entry for the IP address of the remote server sending data.
     * none leaves the host as specified in inputs.conf, which is typically the Splunk system hostname.
     * Default value is dns.
     */
    connectionHost?: pulumi.Input<string>;
    /**
     * Indicates if input is disabled.
     */
    disabled?: pulumi.Input<boolean>;
    /**
     * Host from which the indexer gets data.
     */
    host?: pulumi.Input<string>;
    /**
     * The port number of this input.
     */
    name?: pulumi.Input<string>;
    /**
     * Restrict incoming connections on this port to the host specified here.
     */
    restrictToHost?: pulumi.Input<string>;
}
