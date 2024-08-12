// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.splunk;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.splunk.ConfigsConfArgs;
import com.pulumi.splunk.Utilities;
import com.pulumi.splunk.inputs.ConfigsConfState;
import com.pulumi.splunk.outputs.ConfigsConfAcl;
import java.lang.String;
import java.util.Map;
import javax.annotation.Nullable;

/**
 * ## # Resource: splunk.ConfigsConf
 * 
 * Create and manage configuration file stanzas.
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
 * import com.pulumi.splunk.ConfigsConf;
 * import com.pulumi.splunk.ConfigsConfArgs;
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
 *         var new_conf_stanza = new ConfigsConf("new-conf-stanza", ConfigsConfArgs.builder()
 *             .name("custom-conf/custom")
 *             .variables(Map.ofEntries(
 *                 Map.entry("disabled", "false"),
 *                 Map.entry("custom_key", "value")
 *             ))
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 */
@ResourceType(type="splunk:index/configsConf:ConfigsConf")
public class ConfigsConf extends com.pulumi.resources.CustomResource {
    @Export(name="acl", refs={ConfigsConfAcl.class}, tree="[0]")
    private Output<ConfigsConfAcl> acl;

    public Output<ConfigsConfAcl> acl() {
        return this.acl;
    }
    /**
     * A &#39;/&#39; separated string consisting of {conf_file_name}/{stanza_name} ex. props/custom_stanza
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return A &#39;/&#39; separated string consisting of {conf_file_name}/{stanza_name} ex. props/custom_stanza
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * A map of key value pairs for a stanza.
     * 
     */
    @Export(name="variables", refs={Map.class,String.class}, tree="[0,1,1]")
    private Output<Map<String,String>> variables;

    /**
     * @return A map of key value pairs for a stanza.
     * 
     */
    public Output<Map<String,String>> variables() {
        return this.variables;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public ConfigsConf(java.lang.String name) {
        this(name, ConfigsConfArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public ConfigsConf(java.lang.String name, @Nullable ConfigsConfArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public ConfigsConf(java.lang.String name, @Nullable ConfigsConfArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("splunk:index/configsConf:ConfigsConf", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private ConfigsConf(java.lang.String name, Output<java.lang.String> id, @Nullable ConfigsConfState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("splunk:index/configsConf:ConfigsConf", name, state, makeResourceOptions(options, id), false);
    }

    private static ConfigsConfArgs makeArgs(@Nullable ConfigsConfArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? ConfigsConfArgs.Empty : args;
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
    public static ConfigsConf get(java.lang.String name, Output<java.lang.String> id, @Nullable ConfigsConfState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new ConfigsConf(name, id, state, options);
    }
}
