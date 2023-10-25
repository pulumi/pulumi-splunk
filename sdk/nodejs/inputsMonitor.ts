// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * ## # Resource: splunk.InputsMonitor
 *
 * Create or update a new file or directory monitor input.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as splunk from "@pulumi/splunk";
 *
 * const monitor = new splunk.InputsMonitor("monitor", {
 *     recursive: true,
 *     sourcetype: "text",
 * });
 * ```
 */
export class InputsMonitor extends pulumi.CustomResource {
    /**
     * Get an existing InputsMonitor resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: InputsMonitorState, opts?: pulumi.CustomResourceOptions): InputsMonitor {
        return new InputsMonitor(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'splunk:index/inputsMonitor:InputsMonitor';

    /**
     * Returns true if the given object is an instance of InputsMonitor.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is InputsMonitor {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === InputsMonitor.__pulumiType;
    }

    /**
     * The app/user context that is the namespace for the resource
     */
    public readonly acl!: pulumi.Output<outputs.InputsMonitorAcl>;
    /**
     * Specify a regular expression for a file path. The file path that matches this regular expression is not indexed.
     */
    public readonly blacklist!: pulumi.Output<string>;
    /**
     * A string that modifies the file tracking identity for files in this input. The magic value <SOURCE> invokes special behavior.
     */
    public readonly crcSalt!: pulumi.Output<string>;
    /**
     * Indicates if input monitoring is disabled.
     */
    public readonly disabled!: pulumi.Output<boolean>;
    /**
     * If set to true, files that are seen for the first time is read from the end.
     */
    public readonly followTail!: pulumi.Output<boolean>;
    /**
     * The value to populate in the host field for events from this data input.
     */
    public readonly host!: pulumi.Output<string>;
    /**
     * Specify a regular expression for a file path. If the path for a file matches this regular expression, the captured value is used to populate the host field for events from this data input. The regular expression must have one capture group.
     */
    public readonly hostRegex!: pulumi.Output<string>;
    /**
     * Use the specified slash-separate segment of the filepath as the host field value.
     */
    public readonly hostSegment!: pulumi.Output<number>;
    /**
     * Specify a time value. If the modification time of a file being monitored falls outside of this rolling time window, the file is no longer being monitored.
     */
    public readonly ignoreOlderThan!: pulumi.Output<string>;
    /**
     * Which index events from this input should be stored in. Defaults to default.
     */
    public readonly index!: pulumi.Output<string>;
    /**
     * The file or directory path to monitor on the system.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Setting this to false prevents monitoring of any subdirectories encountered within this data input.
     */
    public readonly recursive!: pulumi.Output<boolean>;
    /**
     * The value to populate in the source field for events from this data input. The same source should not be used for multiple data inputs.
     */
    public readonly renameSource!: pulumi.Output<string>;
    /**
     * The value to populate in the sourcetype field for incoming events.
     */
    public readonly sourcetype!: pulumi.Output<string>;
    /**
     * When Splunk software reaches the end of a file that is being read, the file is kept open for a minimum of the number of seconds specified in this value. After this period has elapsed, the file is checked again for more data.
     */
    public readonly timeBeforeClose!: pulumi.Output<number>;
    /**
     * Specify a regular expression for a file path. Only file paths that match this regular expression are indexed.
     */
    public readonly whitelist!: pulumi.Output<string>;

    /**
     * Create a InputsMonitor resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: InputsMonitorArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: InputsMonitorArgs | InputsMonitorState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as InputsMonitorState | undefined;
            resourceInputs["acl"] = state ? state.acl : undefined;
            resourceInputs["blacklist"] = state ? state.blacklist : undefined;
            resourceInputs["crcSalt"] = state ? state.crcSalt : undefined;
            resourceInputs["disabled"] = state ? state.disabled : undefined;
            resourceInputs["followTail"] = state ? state.followTail : undefined;
            resourceInputs["host"] = state ? state.host : undefined;
            resourceInputs["hostRegex"] = state ? state.hostRegex : undefined;
            resourceInputs["hostSegment"] = state ? state.hostSegment : undefined;
            resourceInputs["ignoreOlderThan"] = state ? state.ignoreOlderThan : undefined;
            resourceInputs["index"] = state ? state.index : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["recursive"] = state ? state.recursive : undefined;
            resourceInputs["renameSource"] = state ? state.renameSource : undefined;
            resourceInputs["sourcetype"] = state ? state.sourcetype : undefined;
            resourceInputs["timeBeforeClose"] = state ? state.timeBeforeClose : undefined;
            resourceInputs["whitelist"] = state ? state.whitelist : undefined;
        } else {
            const args = argsOrState as InputsMonitorArgs | undefined;
            resourceInputs["acl"] = args ? args.acl : undefined;
            resourceInputs["blacklist"] = args ? args.blacklist : undefined;
            resourceInputs["crcSalt"] = args ? args.crcSalt : undefined;
            resourceInputs["disabled"] = args ? args.disabled : undefined;
            resourceInputs["followTail"] = args ? args.followTail : undefined;
            resourceInputs["host"] = args ? args.host : undefined;
            resourceInputs["hostRegex"] = args ? args.hostRegex : undefined;
            resourceInputs["hostSegment"] = args ? args.hostSegment : undefined;
            resourceInputs["ignoreOlderThan"] = args ? args.ignoreOlderThan : undefined;
            resourceInputs["index"] = args ? args.index : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["recursive"] = args ? args.recursive : undefined;
            resourceInputs["renameSource"] = args ? args.renameSource : undefined;
            resourceInputs["sourcetype"] = args ? args.sourcetype : undefined;
            resourceInputs["timeBeforeClose"] = args ? args.timeBeforeClose : undefined;
            resourceInputs["whitelist"] = args ? args.whitelist : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(InputsMonitor.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering InputsMonitor resources.
 */
export interface InputsMonitorState {
    /**
     * The app/user context that is the namespace for the resource
     */
    acl?: pulumi.Input<inputs.InputsMonitorAcl>;
    /**
     * Specify a regular expression for a file path. The file path that matches this regular expression is not indexed.
     */
    blacklist?: pulumi.Input<string>;
    /**
     * A string that modifies the file tracking identity for files in this input. The magic value <SOURCE> invokes special behavior.
     */
    crcSalt?: pulumi.Input<string>;
    /**
     * Indicates if input monitoring is disabled.
     */
    disabled?: pulumi.Input<boolean>;
    /**
     * If set to true, files that are seen for the first time is read from the end.
     */
    followTail?: pulumi.Input<boolean>;
    /**
     * The value to populate in the host field for events from this data input.
     */
    host?: pulumi.Input<string>;
    /**
     * Specify a regular expression for a file path. If the path for a file matches this regular expression, the captured value is used to populate the host field for events from this data input. The regular expression must have one capture group.
     */
    hostRegex?: pulumi.Input<string>;
    /**
     * Use the specified slash-separate segment of the filepath as the host field value.
     */
    hostSegment?: pulumi.Input<number>;
    /**
     * Specify a time value. If the modification time of a file being monitored falls outside of this rolling time window, the file is no longer being monitored.
     */
    ignoreOlderThan?: pulumi.Input<string>;
    /**
     * Which index events from this input should be stored in. Defaults to default.
     */
    index?: pulumi.Input<string>;
    /**
     * The file or directory path to monitor on the system.
     */
    name?: pulumi.Input<string>;
    /**
     * Setting this to false prevents monitoring of any subdirectories encountered within this data input.
     */
    recursive?: pulumi.Input<boolean>;
    /**
     * The value to populate in the source field for events from this data input. The same source should not be used for multiple data inputs.
     */
    renameSource?: pulumi.Input<string>;
    /**
     * The value to populate in the sourcetype field for incoming events.
     */
    sourcetype?: pulumi.Input<string>;
    /**
     * When Splunk software reaches the end of a file that is being read, the file is kept open for a minimum of the number of seconds specified in this value. After this period has elapsed, the file is checked again for more data.
     */
    timeBeforeClose?: pulumi.Input<number>;
    /**
     * Specify a regular expression for a file path. Only file paths that match this regular expression are indexed.
     */
    whitelist?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a InputsMonitor resource.
 */
export interface InputsMonitorArgs {
    /**
     * The app/user context that is the namespace for the resource
     */
    acl?: pulumi.Input<inputs.InputsMonitorAcl>;
    /**
     * Specify a regular expression for a file path. The file path that matches this regular expression is not indexed.
     */
    blacklist?: pulumi.Input<string>;
    /**
     * A string that modifies the file tracking identity for files in this input. The magic value <SOURCE> invokes special behavior.
     */
    crcSalt?: pulumi.Input<string>;
    /**
     * Indicates if input monitoring is disabled.
     */
    disabled?: pulumi.Input<boolean>;
    /**
     * If set to true, files that are seen for the first time is read from the end.
     */
    followTail?: pulumi.Input<boolean>;
    /**
     * The value to populate in the host field for events from this data input.
     */
    host?: pulumi.Input<string>;
    /**
     * Specify a regular expression for a file path. If the path for a file matches this regular expression, the captured value is used to populate the host field for events from this data input. The regular expression must have one capture group.
     */
    hostRegex?: pulumi.Input<string>;
    /**
     * Use the specified slash-separate segment of the filepath as the host field value.
     */
    hostSegment?: pulumi.Input<number>;
    /**
     * Specify a time value. If the modification time of a file being monitored falls outside of this rolling time window, the file is no longer being monitored.
     */
    ignoreOlderThan?: pulumi.Input<string>;
    /**
     * Which index events from this input should be stored in. Defaults to default.
     */
    index?: pulumi.Input<string>;
    /**
     * The file or directory path to monitor on the system.
     */
    name?: pulumi.Input<string>;
    /**
     * Setting this to false prevents monitoring of any subdirectories encountered within this data input.
     */
    recursive?: pulumi.Input<boolean>;
    /**
     * The value to populate in the source field for events from this data input. The same source should not be used for multiple data inputs.
     */
    renameSource?: pulumi.Input<string>;
    /**
     * The value to populate in the sourcetype field for incoming events.
     */
    sourcetype?: pulumi.Input<string>;
    /**
     * When Splunk software reaches the end of a file that is being read, the file is kept open for a minimum of the number of seconds specified in this value. After this period has elapsed, the file is checked again for more data.
     */
    timeBeforeClose?: pulumi.Input<number>;
    /**
     * Specify a regular expression for a file path. Only file paths that match this regular expression are indexed.
     */
    whitelist?: pulumi.Input<string>;
}
