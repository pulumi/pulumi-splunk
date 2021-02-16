// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

/**
 * ## # Resource: splunk.OutputsTcpSyslog
 *
 * Access the configuration of a forwarded server configured to provide data in standard syslog format.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as splunk from "@pulumi/splunk";
 *
 * const tcpSyslog = new splunk.OutputsTcpSyslog("tcp_syslog", {
 *     priority: 5,
 *     server: "new-host-1:1234",
 * });
 * ```
 */
export class OutputsTcpSyslog extends pulumi.CustomResource {
    /**
     * Get an existing OutputsTcpSyslog resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: OutputsTcpSyslogState, opts?: pulumi.CustomResourceOptions): OutputsTcpSyslog {
        return new OutputsTcpSyslog(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'splunk:index/outputsTcpSyslog:OutputsTcpSyslog';

    /**
     * Returns true if the given object is an instance of OutputsTcpSyslog.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is OutputsTcpSyslog {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === OutputsTcpSyslog.__pulumiType;
    }

    /**
     * The app/user context that is the namespace for the resource
     */
    public readonly acl!: pulumi.Output<outputs.OutputsTcpSyslogAcl>;
    /**
     * If true, disables global syslog settings.
     */
    public readonly disabled!: pulumi.Output<boolean>;
    /**
     * Name of the syslog output group. This is name used when creating syslog configuration in outputs.conf.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Sets syslog priority value. The priority value should specified as an integer. See $SPLUNK_HOME/etc/system/README/outputs.conf.spec for details.
     */
    public readonly priority!: pulumi.Output<number>;
    /**
     * host:port of the server where syslog data should be sent
     */
    public readonly server!: pulumi.Output<string>;
    /**
     * Specifies a rule for handling data in addition to that provided by the "syslog" sourcetype. By default, there is no value for syslogSourceType.
     * <br>This string is used as a substring match against the sourcetype key. For example, if the string is set to 'syslog', then all source types containing the string "syslog" receives this special treatment.
     * To match a source type explicitly, use the pattern "sourcetype::sourcetype_name." For example
     * syslogSourcetype = sourcetype::apache_common
     * Data that is "syslog" or matches this setting is assumed to already be in syslog format.
     * Data that does not match the rules has a header, potentially a timestamp, and a hostname added to the front of the event. This is how Splunk software causes arbitrary log data to match syslog expectations.
     */
    public readonly syslogSourcetype!: pulumi.Output<string>;
    /**
     * Format of timestamp to add at start of the events to be forwarded.
     * The format is a strftime-style timestamp formatting string. See $SPLUNK_HOME/etc/system/README/outputs.conf.spec for details.
     */
    public readonly timestampFormat!: pulumi.Output<string>;
    /**
     * Protocol to use to send syslog data. Valid values: (tcp | udp ).
     */
    public readonly type!: pulumi.Output<string>;

    /**
     * Create a OutputsTcpSyslog resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: OutputsTcpSyslogArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: OutputsTcpSyslogArgs | OutputsTcpSyslogState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as OutputsTcpSyslogState | undefined;
            inputs["acl"] = state ? state.acl : undefined;
            inputs["disabled"] = state ? state.disabled : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["priority"] = state ? state.priority : undefined;
            inputs["server"] = state ? state.server : undefined;
            inputs["syslogSourcetype"] = state ? state.syslogSourcetype : undefined;
            inputs["timestampFormat"] = state ? state.timestampFormat : undefined;
            inputs["type"] = state ? state.type : undefined;
        } else {
            const args = argsOrState as OutputsTcpSyslogArgs | undefined;
            inputs["acl"] = args ? args.acl : undefined;
            inputs["disabled"] = args ? args.disabled : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["priority"] = args ? args.priority : undefined;
            inputs["server"] = args ? args.server : undefined;
            inputs["syslogSourcetype"] = args ? args.syslogSourcetype : undefined;
            inputs["timestampFormat"] = args ? args.timestampFormat : undefined;
            inputs["type"] = args ? args.type : undefined;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(OutputsTcpSyslog.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering OutputsTcpSyslog resources.
 */
export interface OutputsTcpSyslogState {
    /**
     * The app/user context that is the namespace for the resource
     */
    readonly acl?: pulumi.Input<inputs.OutputsTcpSyslogAcl>;
    /**
     * If true, disables global syslog settings.
     */
    readonly disabled?: pulumi.Input<boolean>;
    /**
     * Name of the syslog output group. This is name used when creating syslog configuration in outputs.conf.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Sets syslog priority value. The priority value should specified as an integer. See $SPLUNK_HOME/etc/system/README/outputs.conf.spec for details.
     */
    readonly priority?: pulumi.Input<number>;
    /**
     * host:port of the server where syslog data should be sent
     */
    readonly server?: pulumi.Input<string>;
    /**
     * Specifies a rule for handling data in addition to that provided by the "syslog" sourcetype. By default, there is no value for syslogSourceType.
     * <br>This string is used as a substring match against the sourcetype key. For example, if the string is set to 'syslog', then all source types containing the string "syslog" receives this special treatment.
     * To match a source type explicitly, use the pattern "sourcetype::sourcetype_name." For example
     * syslogSourcetype = sourcetype::apache_common
     * Data that is "syslog" or matches this setting is assumed to already be in syslog format.
     * Data that does not match the rules has a header, potentially a timestamp, and a hostname added to the front of the event. This is how Splunk software causes arbitrary log data to match syslog expectations.
     */
    readonly syslogSourcetype?: pulumi.Input<string>;
    /**
     * Format of timestamp to add at start of the events to be forwarded.
     * The format is a strftime-style timestamp formatting string. See $SPLUNK_HOME/etc/system/README/outputs.conf.spec for details.
     */
    readonly timestampFormat?: pulumi.Input<string>;
    /**
     * Protocol to use to send syslog data. Valid values: (tcp | udp ).
     */
    readonly type?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a OutputsTcpSyslog resource.
 */
export interface OutputsTcpSyslogArgs {
    /**
     * The app/user context that is the namespace for the resource
     */
    readonly acl?: pulumi.Input<inputs.OutputsTcpSyslogAcl>;
    /**
     * If true, disables global syslog settings.
     */
    readonly disabled?: pulumi.Input<boolean>;
    /**
     * Name of the syslog output group. This is name used when creating syslog configuration in outputs.conf.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Sets syslog priority value. The priority value should specified as an integer. See $SPLUNK_HOME/etc/system/README/outputs.conf.spec for details.
     */
    readonly priority?: pulumi.Input<number>;
    /**
     * host:port of the server where syslog data should be sent
     */
    readonly server?: pulumi.Input<string>;
    /**
     * Specifies a rule for handling data in addition to that provided by the "syslog" sourcetype. By default, there is no value for syslogSourceType.
     * <br>This string is used as a substring match against the sourcetype key. For example, if the string is set to 'syslog', then all source types containing the string "syslog" receives this special treatment.
     * To match a source type explicitly, use the pattern "sourcetype::sourcetype_name." For example
     * syslogSourcetype = sourcetype::apache_common
     * Data that is "syslog" or matches this setting is assumed to already be in syslog format.
     * Data that does not match the rules has a header, potentially a timestamp, and a hostname added to the front of the event. This is how Splunk software causes arbitrary log data to match syslog expectations.
     */
    readonly syslogSourcetype?: pulumi.Input<string>;
    /**
     * Format of timestamp to add at start of the events to be forwarded.
     * The format is a strftime-style timestamp formatting string. See $SPLUNK_HOME/etc/system/README/outputs.conf.spec for details.
     */
    readonly timestampFormat?: pulumi.Input<string>;
    /**
     * Protocol to use to send syslog data. Valid values: (tcp | udp ).
     */
    readonly type?: pulumi.Input<string>;
}
