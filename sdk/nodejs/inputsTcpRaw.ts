// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

/**
 * ## # Resource: splunk.InputsTcpRaw
 *
 * Create or update raw TCP input information for managing raw tcp inputs from forwarders.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as splunk from "@pulumi/splunk";
 *
 * const tcpRaw = new splunk.InputsTcpRaw("tcp_raw", {
 *     disabled: false,
 *     index: "main",
 *     queue: "indexQueue",
 *     source: "new",
 *     sourcetype: "new",
 * });
 * ```
 */
export class InputsTcpRaw extends pulumi.CustomResource {
    /**
     * Get an existing InputsTcpRaw resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: InputsTcpRawState, opts?: pulumi.CustomResourceOptions): InputsTcpRaw {
        return new InputsTcpRaw(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'splunk:index/inputsTcpRaw:InputsTcpRaw';

    /**
     * Returns true if the given object is an instance of InputsTcpRaw.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is InputsTcpRaw {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === InputsTcpRaw.__pulumiType;
    }

    /**
     * The app/user context that is the namespace for the resource
     */
    public readonly acl!: pulumi.Output<outputs.InputsTcpRawAcl>;
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
     * Index to store generated events. Defaults to default.
     */
    public readonly index!: pulumi.Output<string>;
    /**
     * The input port which receives raw data.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Valid values: (parsingQueue | indexQueue)
     * Specifies where the input processor should deposit the events it reads. Defaults to parsingQueue.
     * Set queue to parsingQueue to apply props.conf and other parsing rules to your data. For more information about props.conf and rules for timestamping and linebreaking, refer to props.conf and the online documentation at "Monitor files and directories with inputs.conf"
     * Set queue to indexQueue to send your data directly into the index.
     */
    public readonly queue!: pulumi.Output<string>;
    /**
     * Specifies in seconds the timeout value for adding a Done-key. Default value is 10 seconds.
     * If a connection over the port specified by name remains idle after receiving data for specified number of seconds, it adds a Done-key. This implies the last event is completely received.
     */
    public readonly rawTcpDoneTimeout!: pulumi.Output<number>;
    /**
     * Allows for restricting this input to only accept data from the host specified here.
     */
    public readonly restrictToHost!: pulumi.Output<string>;
    /**
     * Sets the source key/field for events from this input. Defaults to the input file path.
     * Sets the source key initial value. The key is used during parsing/indexing, in particular to set the source field during indexing. It is also the source field used at search time. As a convenience, the chosen string is prepended with 'source::'.
     */
    public readonly source!: pulumi.Output<string>;
    /**
     * Set the source type for events from this input.
     * "sourcetype=" is automatically prepended to <string>.
     * Defaults to audittrail (if signedaudit=true) or fschange (if signedaudit=false).
     */
    public readonly sourcetype!: pulumi.Output<string>;

    /**
     * Create a InputsTcpRaw resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: InputsTcpRawArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: InputsTcpRawArgs | InputsTcpRawState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as InputsTcpRawState | undefined;
            inputs["acl"] = state ? state.acl : undefined;
            inputs["connectionHost"] = state ? state.connectionHost : undefined;
            inputs["disabled"] = state ? state.disabled : undefined;
            inputs["host"] = state ? state.host : undefined;
            inputs["index"] = state ? state.index : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["queue"] = state ? state.queue : undefined;
            inputs["rawTcpDoneTimeout"] = state ? state.rawTcpDoneTimeout : undefined;
            inputs["restrictToHost"] = state ? state.restrictToHost : undefined;
            inputs["source"] = state ? state.source : undefined;
            inputs["sourcetype"] = state ? state.sourcetype : undefined;
        } else {
            const args = argsOrState as InputsTcpRawArgs | undefined;
            inputs["acl"] = args ? args.acl : undefined;
            inputs["connectionHost"] = args ? args.connectionHost : undefined;
            inputs["disabled"] = args ? args.disabled : undefined;
            inputs["host"] = args ? args.host : undefined;
            inputs["index"] = args ? args.index : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["queue"] = args ? args.queue : undefined;
            inputs["rawTcpDoneTimeout"] = args ? args.rawTcpDoneTimeout : undefined;
            inputs["restrictToHost"] = args ? args.restrictToHost : undefined;
            inputs["source"] = args ? args.source : undefined;
            inputs["sourcetype"] = args ? args.sourcetype : undefined;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(InputsTcpRaw.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering InputsTcpRaw resources.
 */
export interface InputsTcpRawState {
    /**
     * The app/user context that is the namespace for the resource
     */
    readonly acl?: pulumi.Input<inputs.InputsTcpRawAcl>;
    /**
     * Valid values: (ip | dns | none)
     * Set the host for the remote server that is sending data.
     * ip sets the host to the IP address of the remote server sending data.
     * dns sets the host to the reverse DNS entry for the IP address of the remote server sending data.
     * none leaves the host as specified in inputs.conf, which is typically the Splunk system hostname.
     * Default value is dns.
     */
    readonly connectionHost?: pulumi.Input<string>;
    /**
     * Indicates if input is disabled.
     */
    readonly disabled?: pulumi.Input<boolean>;
    /**
     * Host from which the indexer gets data.
     */
    readonly host?: pulumi.Input<string>;
    /**
     * Index to store generated events. Defaults to default.
     */
    readonly index?: pulumi.Input<string>;
    /**
     * The input port which receives raw data.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Valid values: (parsingQueue | indexQueue)
     * Specifies where the input processor should deposit the events it reads. Defaults to parsingQueue.
     * Set queue to parsingQueue to apply props.conf and other parsing rules to your data. For more information about props.conf and rules for timestamping and linebreaking, refer to props.conf and the online documentation at "Monitor files and directories with inputs.conf"
     * Set queue to indexQueue to send your data directly into the index.
     */
    readonly queue?: pulumi.Input<string>;
    /**
     * Specifies in seconds the timeout value for adding a Done-key. Default value is 10 seconds.
     * If a connection over the port specified by name remains idle after receiving data for specified number of seconds, it adds a Done-key. This implies the last event is completely received.
     */
    readonly rawTcpDoneTimeout?: pulumi.Input<number>;
    /**
     * Allows for restricting this input to only accept data from the host specified here.
     */
    readonly restrictToHost?: pulumi.Input<string>;
    /**
     * Sets the source key/field for events from this input. Defaults to the input file path.
     * Sets the source key initial value. The key is used during parsing/indexing, in particular to set the source field during indexing. It is also the source field used at search time. As a convenience, the chosen string is prepended with 'source::'.
     */
    readonly source?: pulumi.Input<string>;
    /**
     * Set the source type for events from this input.
     * "sourcetype=" is automatically prepended to <string>.
     * Defaults to audittrail (if signedaudit=true) or fschange (if signedaudit=false).
     */
    readonly sourcetype?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a InputsTcpRaw resource.
 */
export interface InputsTcpRawArgs {
    /**
     * The app/user context that is the namespace for the resource
     */
    readonly acl?: pulumi.Input<inputs.InputsTcpRawAcl>;
    /**
     * Valid values: (ip | dns | none)
     * Set the host for the remote server that is sending data.
     * ip sets the host to the IP address of the remote server sending data.
     * dns sets the host to the reverse DNS entry for the IP address of the remote server sending data.
     * none leaves the host as specified in inputs.conf, which is typically the Splunk system hostname.
     * Default value is dns.
     */
    readonly connectionHost?: pulumi.Input<string>;
    /**
     * Indicates if input is disabled.
     */
    readonly disabled?: pulumi.Input<boolean>;
    /**
     * Host from which the indexer gets data.
     */
    readonly host?: pulumi.Input<string>;
    /**
     * Index to store generated events. Defaults to default.
     */
    readonly index?: pulumi.Input<string>;
    /**
     * The input port which receives raw data.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Valid values: (parsingQueue | indexQueue)
     * Specifies where the input processor should deposit the events it reads. Defaults to parsingQueue.
     * Set queue to parsingQueue to apply props.conf and other parsing rules to your data. For more information about props.conf and rules for timestamping and linebreaking, refer to props.conf and the online documentation at "Monitor files and directories with inputs.conf"
     * Set queue to indexQueue to send your data directly into the index.
     */
    readonly queue?: pulumi.Input<string>;
    /**
     * Specifies in seconds the timeout value for adding a Done-key. Default value is 10 seconds.
     * If a connection over the port specified by name remains idle after receiving data for specified number of seconds, it adds a Done-key. This implies the last event is completely received.
     */
    readonly rawTcpDoneTimeout?: pulumi.Input<number>;
    /**
     * Allows for restricting this input to only accept data from the host specified here.
     */
    readonly restrictToHost?: pulumi.Input<string>;
    /**
     * Sets the source key/field for events from this input. Defaults to the input file path.
     * Sets the source key initial value. The key is used during parsing/indexing, in particular to set the source field during indexing. It is also the source field used at search time. As a convenience, the chosen string is prepended with 'source::'.
     */
    readonly source?: pulumi.Input<string>;
    /**
     * Set the source type for events from this input.
     * "sourcetype=" is automatically prepended to <string>.
     * Defaults to audittrail (if signedaudit=true) or fschange (if signedaudit=false).
     */
    readonly sourcetype?: pulumi.Input<string>;
}
