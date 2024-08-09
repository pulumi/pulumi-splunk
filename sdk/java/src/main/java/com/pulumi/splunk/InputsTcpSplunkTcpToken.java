// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.splunk;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.splunk.InputsTcpSplunkTcpTokenArgs;
import com.pulumi.splunk.Utilities;
import com.pulumi.splunk.inputs.InputsTcpSplunkTcpTokenState;
import com.pulumi.splunk.outputs.InputsTcpSplunkTcpTokenAcl;
import java.lang.String;
import javax.annotation.Nullable;

/**
 * ## # Resource: splunk.InputsTcpSplunkTcpToken
 * 
 * Manage receiver access using tokens.
 * 
 * ## Example Usage
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.splunk.InputsTcpSplunkTcpToken;
 * import com.pulumi.splunk.InputsTcpSplunkTcpTokenArgs;
 * import java.util.List;
 * import java.util.ArrayList;
 * import java.util.Map;
 * import java.io.File;
 * import java.nio.file.Files;
 * import java.nio.file.Paths;
 * 
 * public class App {
 *     public static void main(String[] args) {
 *         Pulumi.run(App::stack);
 *     }
 * 
 *     public static void stack(Context ctx) {
 *         var tcpSplunkTcpToken = new InputsTcpSplunkTcpToken("tcpSplunkTcpToken", InputsTcpSplunkTcpTokenArgs.builder()
 *             .name("new-splunk-tcp-token")
 *             .token("D66C45B3-7C28-48A1-A13A-027914146501")
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 */
@ResourceType(type="splunk:index/inputsTcpSplunkTcpToken:InputsTcpSplunkTcpToken")
public class InputsTcpSplunkTcpToken extends com.pulumi.resources.CustomResource {
    /**
     * The app/user context that is the namespace for the resource
     * 
     */
    @Export(name="acl", refs={InputsTcpSplunkTcpTokenAcl.class}, tree="[0]")
    private Output<InputsTcpSplunkTcpTokenAcl> acl;

    /**
     * @return The app/user context that is the namespace for the resource
     * 
     */
    public Output<InputsTcpSplunkTcpTokenAcl> acl() {
        return this.acl;
    }
    /**
     * Required. Name for the token to create.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return Required. Name for the token to create.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * Optional. Token value to use. If unspecified, a token is generated automatically.
     * 
     */
    @Export(name="token", refs={String.class}, tree="[0]")
    private Output<String> token;

    /**
     * @return Optional. Token value to use. If unspecified, a token is generated automatically.
     * 
     */
    public Output<String> token() {
        return this.token;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public InputsTcpSplunkTcpToken(java.lang.String name) {
        this(name, InputsTcpSplunkTcpTokenArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public InputsTcpSplunkTcpToken(java.lang.String name, @Nullable InputsTcpSplunkTcpTokenArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public InputsTcpSplunkTcpToken(java.lang.String name, @Nullable InputsTcpSplunkTcpTokenArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("splunk:index/inputsTcpSplunkTcpToken:InputsTcpSplunkTcpToken", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private InputsTcpSplunkTcpToken(java.lang.String name, Output<java.lang.String> id, @Nullable InputsTcpSplunkTcpTokenState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("splunk:index/inputsTcpSplunkTcpToken:InputsTcpSplunkTcpToken", name, state, makeResourceOptions(options, id), false);
    }

    private static InputsTcpSplunkTcpTokenArgs makeArgs(@Nullable InputsTcpSplunkTcpTokenArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? InputsTcpSplunkTcpTokenArgs.Empty : args;
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<java.lang.String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .build();
        return com.pulumi.resources.CustomResourceOptions.merge(defaultOptions, options, id);
    }

    /**
     * Get an existing Host resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state
     * @param options Optional settings to control the behavior of the CustomResource.
     */
    public static InputsTcpSplunkTcpToken get(java.lang.String name, Output<java.lang.String> id, @Nullable InputsTcpSplunkTcpTokenState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new InputsTcpSplunkTcpToken(name, id, state, options);
    }
}
