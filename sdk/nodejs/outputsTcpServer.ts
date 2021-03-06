// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

/**
 * ## # Resource: splunk.OutputsTcpServer
 *
 * Access data forwarding configurations.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as splunk from "@pulumi/splunk";
 *
 * const tcpServer = new splunk.OutputsTcpServer("tcp_server", {
 *     sslAltNameToCheck: "old-host",
 * });
 * ```
 */
export class OutputsTcpServer extends pulumi.CustomResource {
    /**
     * Get an existing OutputsTcpServer resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: OutputsTcpServerState, opts?: pulumi.CustomResourceOptions): OutputsTcpServer {
        return new OutputsTcpServer(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'splunk:index/outputsTcpServer:OutputsTcpServer';

    /**
     * Returns true if the given object is an instance of OutputsTcpServer.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is OutputsTcpServer {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === OutputsTcpServer.__pulumiType;
    }

    /**
     * The app/user context that is the namespace for the resource
     */
    public readonly acl!: pulumi.Output<outputs.OutputsTcpServerAcl>;
    /**
     * If true, disables the group.
     */
    public readonly disabled!: pulumi.Output<boolean>;
    /**
     * Valid values: (clone | balance | autobalance)
     * The data distribution method used when two or more servers exist in the same forwarder group.
     */
    public readonly method!: pulumi.Output<string>;
    /**
     * <host>:<port> of the Splunk receiver. <host> can be either an ip address or server name. <port> is the that port that the Splunk receiver is listening on.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The alternate name to match in the remote server's SSL certificate.
     */
    public readonly sslAltNameToCheck!: pulumi.Output<string>;
    /**
     * Path to the client certificate. If specified, connection uses SSL.
     */
    public readonly sslCertPath!: pulumi.Output<string>;
    /**
     * SSL Cipher in the form ALL:!aNULL:!eNULL:!LOW:!EXP:RC4+RSA:+HIGH:+MEDIUM
     */
    public readonly sslCipher!: pulumi.Output<string>;
    /**
     * Check the common name of the server's certificate against this name.
     * If there is no match, assume that Splunk Enterprise is not authenticated against this server. You must specify this setting if sslVerifyServerCert is true.
     */
    public readonly sslCommonNameToCheck!: pulumi.Output<string>;
    /**
     * The password associated with the CAcert.
     * The default Splunk Enterprise CAcert uses the password "password."
     */
    public readonly sslPassword!: pulumi.Output<string>;
    /**
     * The path to the root certificate authority file.
     */
    public readonly sslRootCaPath!: pulumi.Output<string>;
    /**
     * If true, make sure that the server you are connecting to is a valid one (authenticated). Both the common name and the alternate name of the server are then checked for a match.
     */
    public readonly sslVerifyServerCert!: pulumi.Output<boolean>;

    /**
     * Create a OutputsTcpServer resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: OutputsTcpServerArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: OutputsTcpServerArgs | OutputsTcpServerState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as OutputsTcpServerState | undefined;
            inputs["acl"] = state ? state.acl : undefined;
            inputs["disabled"] = state ? state.disabled : undefined;
            inputs["method"] = state ? state.method : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["sslAltNameToCheck"] = state ? state.sslAltNameToCheck : undefined;
            inputs["sslCertPath"] = state ? state.sslCertPath : undefined;
            inputs["sslCipher"] = state ? state.sslCipher : undefined;
            inputs["sslCommonNameToCheck"] = state ? state.sslCommonNameToCheck : undefined;
            inputs["sslPassword"] = state ? state.sslPassword : undefined;
            inputs["sslRootCaPath"] = state ? state.sslRootCaPath : undefined;
            inputs["sslVerifyServerCert"] = state ? state.sslVerifyServerCert : undefined;
        } else {
            const args = argsOrState as OutputsTcpServerArgs | undefined;
            inputs["acl"] = args ? args.acl : undefined;
            inputs["disabled"] = args ? args.disabled : undefined;
            inputs["method"] = args ? args.method : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["sslAltNameToCheck"] = args ? args.sslAltNameToCheck : undefined;
            inputs["sslCertPath"] = args ? args.sslCertPath : undefined;
            inputs["sslCipher"] = args ? args.sslCipher : undefined;
            inputs["sslCommonNameToCheck"] = args ? args.sslCommonNameToCheck : undefined;
            inputs["sslPassword"] = args ? args.sslPassword : undefined;
            inputs["sslRootCaPath"] = args ? args.sslRootCaPath : undefined;
            inputs["sslVerifyServerCert"] = args ? args.sslVerifyServerCert : undefined;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(OutputsTcpServer.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering OutputsTcpServer resources.
 */
export interface OutputsTcpServerState {
    /**
     * The app/user context that is the namespace for the resource
     */
    readonly acl?: pulumi.Input<inputs.OutputsTcpServerAcl>;
    /**
     * If true, disables the group.
     */
    readonly disabled?: pulumi.Input<boolean>;
    /**
     * Valid values: (clone | balance | autobalance)
     * The data distribution method used when two or more servers exist in the same forwarder group.
     */
    readonly method?: pulumi.Input<string>;
    /**
     * <host>:<port> of the Splunk receiver. <host> can be either an ip address or server name. <port> is the that port that the Splunk receiver is listening on.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The alternate name to match in the remote server's SSL certificate.
     */
    readonly sslAltNameToCheck?: pulumi.Input<string>;
    /**
     * Path to the client certificate. If specified, connection uses SSL.
     */
    readonly sslCertPath?: pulumi.Input<string>;
    /**
     * SSL Cipher in the form ALL:!aNULL:!eNULL:!LOW:!EXP:RC4+RSA:+HIGH:+MEDIUM
     */
    readonly sslCipher?: pulumi.Input<string>;
    /**
     * Check the common name of the server's certificate against this name.
     * If there is no match, assume that Splunk Enterprise is not authenticated against this server. You must specify this setting if sslVerifyServerCert is true.
     */
    readonly sslCommonNameToCheck?: pulumi.Input<string>;
    /**
     * The password associated with the CAcert.
     * The default Splunk Enterprise CAcert uses the password "password."
     */
    readonly sslPassword?: pulumi.Input<string>;
    /**
     * The path to the root certificate authority file.
     */
    readonly sslRootCaPath?: pulumi.Input<string>;
    /**
     * If true, make sure that the server you are connecting to is a valid one (authenticated). Both the common name and the alternate name of the server are then checked for a match.
     */
    readonly sslVerifyServerCert?: pulumi.Input<boolean>;
}

/**
 * The set of arguments for constructing a OutputsTcpServer resource.
 */
export interface OutputsTcpServerArgs {
    /**
     * The app/user context that is the namespace for the resource
     */
    readonly acl?: pulumi.Input<inputs.OutputsTcpServerAcl>;
    /**
     * If true, disables the group.
     */
    readonly disabled?: pulumi.Input<boolean>;
    /**
     * Valid values: (clone | balance | autobalance)
     * The data distribution method used when two or more servers exist in the same forwarder group.
     */
    readonly method?: pulumi.Input<string>;
    /**
     * <host>:<port> of the Splunk receiver. <host> can be either an ip address or server name. <port> is the that port that the Splunk receiver is listening on.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The alternate name to match in the remote server's SSL certificate.
     */
    readonly sslAltNameToCheck?: pulumi.Input<string>;
    /**
     * Path to the client certificate. If specified, connection uses SSL.
     */
    readonly sslCertPath?: pulumi.Input<string>;
    /**
     * SSL Cipher in the form ALL:!aNULL:!eNULL:!LOW:!EXP:RC4+RSA:+HIGH:+MEDIUM
     */
    readonly sslCipher?: pulumi.Input<string>;
    /**
     * Check the common name of the server's certificate against this name.
     * If there is no match, assume that Splunk Enterprise is not authenticated against this server. You must specify this setting if sslVerifyServerCert is true.
     */
    readonly sslCommonNameToCheck?: pulumi.Input<string>;
    /**
     * The password associated with the CAcert.
     * The default Splunk Enterprise CAcert uses the password "password."
     */
    readonly sslPassword?: pulumi.Input<string>;
    /**
     * The path to the root certificate authority file.
     */
    readonly sslRootCaPath?: pulumi.Input<string>;
    /**
     * If true, make sure that the server you are connecting to is a valid one (authenticated). Both the common name and the alternate name of the server are then checked for a match.
     */
    readonly sslVerifyServerCert?: pulumi.Input<boolean>;
}
