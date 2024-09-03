// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.splunk;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.splunk.InputsTcpSslArgs;
import com.pulumi.splunk.Utilities;
import com.pulumi.splunk.inputs.InputsTcpSslState;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import javax.annotation.Nullable;

/**
 * # Resource: splunk.InputsTcpSsl
 * Access or update the SSL configuration for the host.
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
 * import com.pulumi.splunk.InputsTcpSsl;
 * import com.pulumi.splunk.InputsTcpSslArgs;
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
 *         var test = new InputsTcpSsl("test", InputsTcpSslArgs.builder()
 *             .disabled(false)
 *             .requireClientCert(true)
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 */
@ResourceType(type="splunk:index/inputsTcpSsl:InputsTcpSsl")
public class InputsTcpSsl extends com.pulumi.resources.CustomResource {
    /**
     * Indicates if input is disabled.
     * 
     */
    @Export(name="disabled", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> disabled;

    /**
     * @return Indicates if input is disabled.
     * 
     */
    public Output<Boolean> disabled() {
        return this.disabled;
    }
    /**
     * Server certificate password, if any.
     * 
     */
    @Export(name="password", refs={String.class}, tree="[0]")
    private Output<String> password;

    /**
     * @return Server certificate password, if any.
     * 
     */
    public Output<String> password() {
        return this.password;
    }
    /**
     * Determines whether a client must authenticate.
     * 
     */
    @Export(name="requireClientCert", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> requireClientCert;

    /**
     * @return Determines whether a client must authenticate.
     * 
     */
    public Output<Boolean> requireClientCert() {
        return this.requireClientCert;
    }
    /**
     * Certificate authority list (root file)
     * 
     */
    @Export(name="rootCa", refs={String.class}, tree="[0]")
    private Output<String> rootCa;

    /**
     * @return Certificate authority list (root file)
     * 
     */
    public Output<String> rootCa() {
        return this.rootCa;
    }
    /**
     * Full path to the server certificate.
     * 
     */
    @Export(name="serverCert", refs={String.class}, tree="[0]")
    private Output<String> serverCert;

    /**
     * @return Full path to the server certificate.
     * 
     */
    public Output<String> serverCert() {
        return this.serverCert;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public InputsTcpSsl(java.lang.String name) {
        this(name, InputsTcpSslArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public InputsTcpSsl(java.lang.String name, @Nullable InputsTcpSslArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public InputsTcpSsl(java.lang.String name, @Nullable InputsTcpSslArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("splunk:index/inputsTcpSsl:InputsTcpSsl", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private InputsTcpSsl(java.lang.String name, Output<java.lang.String> id, @Nullable InputsTcpSslState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("splunk:index/inputsTcpSsl:InputsTcpSsl", name, state, makeResourceOptions(options, id), false);
    }

    private static InputsTcpSslArgs makeArgs(@Nullable InputsTcpSslArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? InputsTcpSslArgs.Empty : args;
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<java.lang.String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .additionalSecretOutputs(List.of(
                "password"
            ))
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
    public static InputsTcpSsl get(java.lang.String name, Output<java.lang.String> id, @Nullable InputsTcpSslState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new InputsTcpSsl(name, id, state, options);
    }
}
