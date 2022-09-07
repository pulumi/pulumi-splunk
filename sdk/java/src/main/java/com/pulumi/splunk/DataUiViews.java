// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.splunk;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.splunk.DataUiViewsArgs;
import com.pulumi.splunk.Utilities;
import com.pulumi.splunk.inputs.DataUiViewsState;
import com.pulumi.splunk.outputs.DataUiViewsAcl;
import java.lang.String;
import javax.annotation.Nullable;

/**
 * ## # Resource: splunk.DataUiViews
 * 
 * Create and manage splunk dashboards/views.
 * ## Example Usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.splunk.DataUiViews;
 * import com.pulumi.splunk.DataUiViewsArgs;
 * import com.pulumi.splunk.inputs.DataUiViewsAclArgs;
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
 *         var dashboard = new DataUiViews(&#34;dashboard&#34;, DataUiViewsArgs.builder()        
 *             .acl(DataUiViewsAclArgs.builder()
 *                 .app(&#34;search&#34;)
 *                 .owner(&#34;admin&#34;)
 *                 .build())
 *             .eaiData(&#34;&#34;&#34;
 *   &lt;dashboard&gt;
 *     &lt;label&gt; 
 *       Terraform Test Dashboard
 *     &lt;/label&gt;
 *   &lt;/dashboard&gt;
 *   
 *             &#34;&#34;&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 */
@ResourceType(type="splunk:index/dataUiViews:DataUiViews")
public class DataUiViews extends com.pulumi.resources.CustomResource {
    @Export(name="acl", type=DataUiViewsAcl.class, parameters={})
    private Output<DataUiViewsAcl> acl;

    public Output<DataUiViewsAcl> acl() {
        return this.acl;
    }
    /**
     * Dashboard XML definition.
     * 
     */
    @Export(name="eaiData", type=String.class, parameters={})
    private Output<String> eaiData;

    /**
     * @return Dashboard XML definition.
     * 
     */
    public Output<String> eaiData() {
        return this.eaiData;
    }
    /**
     * Dashboard name.
     * * `eai:data` - (Required) Dashboard XML definition.
     * 
     */
    @Export(name="name", type=String.class, parameters={})
    private Output<String> name;

    /**
     * @return Dashboard name.
     * * `eai:data` - (Required) Dashboard XML definition.
     * 
     */
    public Output<String> name() {
        return this.name;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public DataUiViews(String name) {
        this(name, DataUiViewsArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public DataUiViews(String name, DataUiViewsArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public DataUiViews(String name, DataUiViewsArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("splunk:index/dataUiViews:DataUiViews", name, args == null ? DataUiViewsArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private DataUiViews(String name, Output<String> id, @Nullable DataUiViewsState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("splunk:index/dataUiViews:DataUiViews", name, state, makeResourceOptions(options, id));
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
    public static DataUiViews get(String name, Output<String> id, @Nullable DataUiViewsState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new DataUiViews(name, id, state, options);
    }
}
