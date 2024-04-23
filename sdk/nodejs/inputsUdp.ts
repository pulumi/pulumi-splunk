// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * ## # Resource: splunk.InputsTcpRaw
 *
 * Create and manage UDP data inputs.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as splunk from "@pulumi/splunk";
 *
 * const udp = new splunk.InputsUdp("udp", {
 *     name: "41000",
 *     index: "main",
 *     source: "new",
 *     sourcetype: "new",
 *     disabled: false,
 * });
 * ```
 */
export class InputsUdp extends pulumi.CustomResource {
    /**
     * Get an existing InputsUdp resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: InputsUdpState, opts?: pulumi.CustomResourceOptions): InputsUdp {
        return new InputsUdp(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'splunk:index/inputsUdp:InputsUdp';

    /**
     * Returns true if the given object is an instance of InputsUdp.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is InputsUdp {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === InputsUdp.__pulumiType;
    }

    /**
     * The app/user context that is the namespace for the resource
     */
    public readonly acl!: pulumi.Output<outputs.InputsUdpAcl>;
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
     * The value to populate in the host field for incoming events. This is used during parsing/indexing, in particular to set the host field. It is also the host field used at search time.
     */
    public readonly host!: pulumi.Output<string>;
    /**
     * Which index events from this input should be stored in. Defaults to default.
     */
    public readonly index!: pulumi.Output<string>;
    /**
     * The UDP port that this input should listen on.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * If set to true, prevents Splunk software from prepending a timestamp and hostname to incoming events.
     */
    public readonly noAppendingTimestamp!: pulumi.Output<boolean>;
    /**
     * If set to true, Splunk software does not remove the priority field from incoming syslog events.
     */
    public readonly noPriorityStripping!: pulumi.Output<boolean>;
    /**
     * Which queue events from this input should be sent to. Generally this does not need to be changed.
     */
    public readonly queue!: pulumi.Output<string>;
    /**
     * Restrict incoming connections on this port to the host specified here.
     * If this is not set, the value specified in [udp://<remote server>:<port>] in inputs.conf is used.
     */
    public readonly restrictToHost!: pulumi.Output<string>;
    /**
     * The value to populate in the source field for incoming events. The same source should not be used for multiple data inputs.
     */
    public readonly source!: pulumi.Output<string>;
    /**
     * The value to populate in the sourcetype field for incoming events.
     */
    public readonly sourcetype!: pulumi.Output<string>;

    /**
     * Create a InputsUdp resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: InputsUdpArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: InputsUdpArgs | InputsUdpState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as InputsUdpState | undefined;
            resourceInputs["acl"] = state ? state.acl : undefined;
            resourceInputs["connectionHost"] = state ? state.connectionHost : undefined;
            resourceInputs["disabled"] = state ? state.disabled : undefined;
            resourceInputs["host"] = state ? state.host : undefined;
            resourceInputs["index"] = state ? state.index : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["noAppendingTimestamp"] = state ? state.noAppendingTimestamp : undefined;
            resourceInputs["noPriorityStripping"] = state ? state.noPriorityStripping : undefined;
            resourceInputs["queue"] = state ? state.queue : undefined;
            resourceInputs["restrictToHost"] = state ? state.restrictToHost : undefined;
            resourceInputs["source"] = state ? state.source : undefined;
            resourceInputs["sourcetype"] = state ? state.sourcetype : undefined;
        } else {
            const args = argsOrState as InputsUdpArgs | undefined;
            resourceInputs["acl"] = args ? args.acl : undefined;
            resourceInputs["connectionHost"] = args ? args.connectionHost : undefined;
            resourceInputs["disabled"] = args ? args.disabled : undefined;
            resourceInputs["host"] = args ? args.host : undefined;
            resourceInputs["index"] = args ? args.index : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["noAppendingTimestamp"] = args ? args.noAppendingTimestamp : undefined;
            resourceInputs["noPriorityStripping"] = args ? args.noPriorityStripping : undefined;
            resourceInputs["queue"] = args ? args.queue : undefined;
            resourceInputs["restrictToHost"] = args ? args.restrictToHost : undefined;
            resourceInputs["source"] = args ? args.source : undefined;
            resourceInputs["sourcetype"] = args ? args.sourcetype : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(InputsUdp.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering InputsUdp resources.
 */
export interface InputsUdpState {
    /**
     * The app/user context that is the namespace for the resource
     */
    acl?: pulumi.Input<inputs.InputsUdpAcl>;
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
     * The value to populate in the host field for incoming events. This is used during parsing/indexing, in particular to set the host field. It is also the host field used at search time.
     */
    host?: pulumi.Input<string>;
    /**
     * Which index events from this input should be stored in. Defaults to default.
     */
    index?: pulumi.Input<string>;
    /**
     * The UDP port that this input should listen on.
     */
    name?: pulumi.Input<string>;
    /**
     * If set to true, prevents Splunk software from prepending a timestamp and hostname to incoming events.
     */
    noAppendingTimestamp?: pulumi.Input<boolean>;
    /**
     * If set to true, Splunk software does not remove the priority field from incoming syslog events.
     */
    noPriorityStripping?: pulumi.Input<boolean>;
    /**
     * Which queue events from this input should be sent to. Generally this does not need to be changed.
     */
    queue?: pulumi.Input<string>;
    /**
     * Restrict incoming connections on this port to the host specified here.
     * If this is not set, the value specified in [udp://<remote server>:<port>] in inputs.conf is used.
     */
    restrictToHost?: pulumi.Input<string>;
    /**
     * The value to populate in the source field for incoming events. The same source should not be used for multiple data inputs.
     */
    source?: pulumi.Input<string>;
    /**
     * The value to populate in the sourcetype field for incoming events.
     */
    sourcetype?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a InputsUdp resource.
 */
export interface InputsUdpArgs {
    /**
     * The app/user context that is the namespace for the resource
     */
    acl?: pulumi.Input<inputs.InputsUdpAcl>;
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
     * The value to populate in the host field for incoming events. This is used during parsing/indexing, in particular to set the host field. It is also the host field used at search time.
     */
    host?: pulumi.Input<string>;
    /**
     * Which index events from this input should be stored in. Defaults to default.
     */
    index?: pulumi.Input<string>;
    /**
     * The UDP port that this input should listen on.
     */
    name?: pulumi.Input<string>;
    /**
     * If set to true, prevents Splunk software from prepending a timestamp and hostname to incoming events.
     */
    noAppendingTimestamp?: pulumi.Input<boolean>;
    /**
     * If set to true, Splunk software does not remove the priority field from incoming syslog events.
     */
    noPriorityStripping?: pulumi.Input<boolean>;
    /**
     * Which queue events from this input should be sent to. Generally this does not need to be changed.
     */
    queue?: pulumi.Input<string>;
    /**
     * Restrict incoming connections on this port to the host specified here.
     * If this is not set, the value specified in [udp://<remote server>:<port>] in inputs.conf is used.
     */
    restrictToHost?: pulumi.Input<string>;
    /**
     * The value to populate in the source field for incoming events. The same source should not be used for multiple data inputs.
     */
    source?: pulumi.Input<string>;
    /**
     * The value to populate in the sourcetype field for incoming events.
     */
    sourcetype?: pulumi.Input<string>;
}
