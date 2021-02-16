// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

/**
 * ## # Resource: splunk.InputsTcpSplunkTcpToken
 *
 * Manage receiver access using tokens.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as splunk from "@pulumi/splunk";
 *
 * const tcpSplunkTcpToken = new splunk.InputsTcpSplunkTcpToken("tcp_splunk_tcp_token", {
 *     token: "D66C45B3-7C28-48A1-A13A-027914146501",
 * });
 * ```
 */
export class InputsTcpSplunkTcpToken extends pulumi.CustomResource {
    /**
     * Get an existing InputsTcpSplunkTcpToken resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: InputsTcpSplunkTcpTokenState, opts?: pulumi.CustomResourceOptions): InputsTcpSplunkTcpToken {
        return new InputsTcpSplunkTcpToken(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'splunk:index/inputsTcpSplunkTcpToken:InputsTcpSplunkTcpToken';

    /**
     * Returns true if the given object is an instance of InputsTcpSplunkTcpToken.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is InputsTcpSplunkTcpToken {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === InputsTcpSplunkTcpToken.__pulumiType;
    }

    /**
     * The app/user context that is the namespace for the resource
     */
    public readonly acl!: pulumi.Output<outputs.InputsTcpSplunkTcpTokenAcl>;
    /**
     * Required. Name for the token to create.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Optional. Token value to use. If unspecified, a token is generated automatically.
     */
    public readonly token!: pulumi.Output<string>;

    /**
     * Create a InputsTcpSplunkTcpToken resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: InputsTcpSplunkTcpTokenArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: InputsTcpSplunkTcpTokenArgs | InputsTcpSplunkTcpTokenState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as InputsTcpSplunkTcpTokenState | undefined;
            inputs["acl"] = state ? state.acl : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["token"] = state ? state.token : undefined;
        } else {
            const args = argsOrState as InputsTcpSplunkTcpTokenArgs | undefined;
            inputs["acl"] = args ? args.acl : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["token"] = args ? args.token : undefined;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(InputsTcpSplunkTcpToken.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering InputsTcpSplunkTcpToken resources.
 */
export interface InputsTcpSplunkTcpTokenState {
    /**
     * The app/user context that is the namespace for the resource
     */
    readonly acl?: pulumi.Input<inputs.InputsTcpSplunkTcpTokenAcl>;
    /**
     * Required. Name for the token to create.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Optional. Token value to use. If unspecified, a token is generated automatically.
     */
    readonly token?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a InputsTcpSplunkTcpToken resource.
 */
export interface InputsTcpSplunkTcpTokenArgs {
    /**
     * The app/user context that is the namespace for the resource
     */
    readonly acl?: pulumi.Input<inputs.InputsTcpSplunkTcpTokenAcl>;
    /**
     * Required. Name for the token to create.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Optional. Token value to use. If unspecified, a token is generated automatically.
     */
    readonly token?: pulumi.Input<string>;
}
