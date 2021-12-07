// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

/**
 * ## # Resource: splunk.InputsHttpEventCollector
 *
 * Create or update HTTP Event Collector input configuration tokens.
 */
export class InputsHttpEventCollector extends pulumi.CustomResource {
    /**
     * Get an existing InputsHttpEventCollector resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: InputsHttpEventCollectorState, opts?: pulumi.CustomResourceOptions): InputsHttpEventCollector {
        return new InputsHttpEventCollector(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'splunk:index/inputsHttpEventCollector:InputsHttpEventCollector';

    /**
     * Returns true if the given object is an instance of InputsHttpEventCollector.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is InputsHttpEventCollector {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === InputsHttpEventCollector.__pulumiType;
    }

    /**
     * The app/user context that is the namespace for the resource
     */
    public readonly acl!: pulumi.Output<outputs.InputsHttpEventCollectorAcl>;
    /**
     * Input disabled indicator
     */
    public readonly disabled!: pulumi.Output<boolean>;
    /**
     * Default host value for events with this token
     */
    public readonly host!: pulumi.Output<string>;
    /**
     * Index to store generated events
     */
    public readonly index!: pulumi.Output<string>;
    /**
     * Set of indexes allowed for events with this token
     */
    public readonly indexes!: pulumi.Output<string[]>;
    /**
     * Token name (inputs.conf key)
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Default source for events with this token
     */
    public readonly source!: pulumi.Output<string>;
    /**
     * Default source type for events with this token
     */
    public readonly sourcetype!: pulumi.Output<string>;
    /**
     * Token value for sending data to collector/event endpoint
     */
    public readonly token!: pulumi.Output<string>;
    /**
     * Indexer acknowledgement for this token
     */
    public readonly useAck!: pulumi.Output<number>;

    /**
     * Create a InputsHttpEventCollector resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: InputsHttpEventCollectorArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: InputsHttpEventCollectorArgs | InputsHttpEventCollectorState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as InputsHttpEventCollectorState | undefined;
            resourceInputs["acl"] = state ? state.acl : undefined;
            resourceInputs["disabled"] = state ? state.disabled : undefined;
            resourceInputs["host"] = state ? state.host : undefined;
            resourceInputs["index"] = state ? state.index : undefined;
            resourceInputs["indexes"] = state ? state.indexes : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["source"] = state ? state.source : undefined;
            resourceInputs["sourcetype"] = state ? state.sourcetype : undefined;
            resourceInputs["token"] = state ? state.token : undefined;
            resourceInputs["useAck"] = state ? state.useAck : undefined;
        } else {
            const args = argsOrState as InputsHttpEventCollectorArgs | undefined;
            resourceInputs["acl"] = args ? args.acl : undefined;
            resourceInputs["disabled"] = args ? args.disabled : undefined;
            resourceInputs["host"] = args ? args.host : undefined;
            resourceInputs["index"] = args ? args.index : undefined;
            resourceInputs["indexes"] = args ? args.indexes : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["source"] = args ? args.source : undefined;
            resourceInputs["sourcetype"] = args ? args.sourcetype : undefined;
            resourceInputs["token"] = args ? args.token : undefined;
            resourceInputs["useAck"] = args ? args.useAck : undefined;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(InputsHttpEventCollector.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering InputsHttpEventCollector resources.
 */
export interface InputsHttpEventCollectorState {
    /**
     * The app/user context that is the namespace for the resource
     */
    acl?: pulumi.Input<inputs.InputsHttpEventCollectorAcl>;
    /**
     * Input disabled indicator
     */
    disabled?: pulumi.Input<boolean>;
    /**
     * Default host value for events with this token
     */
    host?: pulumi.Input<string>;
    /**
     * Index to store generated events
     */
    index?: pulumi.Input<string>;
    /**
     * Set of indexes allowed for events with this token
     */
    indexes?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Token name (inputs.conf key)
     */
    name?: pulumi.Input<string>;
    /**
     * Default source for events with this token
     */
    source?: pulumi.Input<string>;
    /**
     * Default source type for events with this token
     */
    sourcetype?: pulumi.Input<string>;
    /**
     * Token value for sending data to collector/event endpoint
     */
    token?: pulumi.Input<string>;
    /**
     * Indexer acknowledgement for this token
     */
    useAck?: pulumi.Input<number>;
}

/**
 * The set of arguments for constructing a InputsHttpEventCollector resource.
 */
export interface InputsHttpEventCollectorArgs {
    /**
     * The app/user context that is the namespace for the resource
     */
    acl?: pulumi.Input<inputs.InputsHttpEventCollectorAcl>;
    /**
     * Input disabled indicator
     */
    disabled?: pulumi.Input<boolean>;
    /**
     * Default host value for events with this token
     */
    host?: pulumi.Input<string>;
    /**
     * Index to store generated events
     */
    index?: pulumi.Input<string>;
    /**
     * Set of indexes allowed for events with this token
     */
    indexes?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Token name (inputs.conf key)
     */
    name?: pulumi.Input<string>;
    /**
     * Default source for events with this token
     */
    source?: pulumi.Input<string>;
    /**
     * Default source type for events with this token
     */
    sourcetype?: pulumi.Input<string>;
    /**
     * Token value for sending data to collector/event endpoint
     */
    token?: pulumi.Input<string>;
    /**
     * Indexer acknowledgement for this token
     */
    useAck?: pulumi.Input<number>;
}
