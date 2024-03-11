// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.splunk;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.splunk.ShIndexesManagerArgs;
import com.pulumi.splunk.Utilities;
import com.pulumi.splunk.inputs.ShIndexesManagerState;
import com.pulumi.splunk.outputs.ShIndexesManagerAcl;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * ## # Resource: splunk.ShIndexesManager
 * 
 * Create indexes on Splunk Cloud instances. [BETA]
 * 
 * ## Authorization and authentication
 * 
 * As of now there is no support to create indexes in user-specified workspaces on Splunk Cloud.
 * 
 * ## Example Usage
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.splunk.ShIndexesManager;
 * import com.pulumi.splunk.ShIndexesManagerArgs;
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
 *         var tf_index = new ShIndexesManager(&#34;tf-index&#34;, ShIndexesManagerArgs.builder()        
 *             .datatype(&#34;event&#34;)
 *             .frozenTimePeriodInSecs(&#34;94608000&#34;)
 *             .maxGlobalRawDataSizeMb(&#34;100&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 */
@ResourceType(type="splunk:index/shIndexesManager:ShIndexesManager")
public class ShIndexesManager extends com.pulumi.resources.CustomResource {
    @Export(name="acl", refs={ShIndexesManagerAcl.class}, tree="[0]")
    private Output<ShIndexesManagerAcl> acl;

    public Output<ShIndexesManagerAcl> acl() {
        return this.acl;
    }
    /**
     * Valid values: (event | metric). Specifies the type of index.
     * 
     */
    @Export(name="datatype", refs={String.class}, tree="[0]")
    private Output<String> datatype;

    /**
     * @return Valid values: (event | metric). Specifies the type of index.
     * 
     */
    public Output<String> datatype() {
        return this.datatype;
    }
    /**
     * Number of seconds after which indexed data rolls to frozen.
     * Defaults to 94608000 (3 years).Freezing data means it is removed from the index. If you need to archive your data, refer to coldToFrozenDir and coldToFrozenScript parameter documentation.
     * 
     */
    @Export(name="frozenTimePeriodInSecs", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> frozenTimePeriodInSecs;

    /**
     * @return Number of seconds after which indexed data rolls to frozen.
     * Defaults to 94608000 (3 years).Freezing data means it is removed from the index. If you need to archive your data, refer to coldToFrozenDir and coldToFrozenScript parameter documentation.
     * 
     */
    public Output<Optional<String>> frozenTimePeriodInSecs() {
        return Codegen.optional(this.frozenTimePeriodInSecs);
    }
    /**
     * The maximum size of an index (in MB). If an index grows larger than the maximum size, the oldest data is frozen.
     * Defaults to 100 MB.
     * 
     */
    @Export(name="maxGlobalRawDataSizeMb", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> maxGlobalRawDataSizeMb;

    /**
     * @return The maximum size of an index (in MB). If an index grows larger than the maximum size, the oldest data is frozen.
     * Defaults to 100 MB.
     * 
     */
    public Output<Optional<String>> maxGlobalRawDataSizeMb() {
        return Codegen.optional(this.maxGlobalRawDataSizeMb);
    }
    /**
     * The name of the index to create.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return The name of the index to create.
     * 
     */
    public Output<String> name() {
        return this.name;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public ShIndexesManager(String name) {
        this(name, ShIndexesManagerArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public ShIndexesManager(String name, @Nullable ShIndexesManagerArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public ShIndexesManager(String name, @Nullable ShIndexesManagerArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("splunk:index/shIndexesManager:ShIndexesManager", name, args == null ? ShIndexesManagerArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private ShIndexesManager(String name, Output<String> id, @Nullable ShIndexesManagerState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("splunk:index/shIndexesManager:ShIndexesManager", name, state, makeResourceOptions(options, id));
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
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
    public static ShIndexesManager get(String name, Output<String> id, @Nullable ShIndexesManagerState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new ShIndexesManager(name, id, state, options);
    }
}
