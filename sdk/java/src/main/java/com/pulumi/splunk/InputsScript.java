// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.splunk;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.splunk.InputsScriptArgs;
import com.pulumi.splunk.Utilities;
import com.pulumi.splunk.inputs.InputsScriptState;
import com.pulumi.splunk.outputs.InputsScriptAcl;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import javax.annotation.Nullable;

/**
 * ## # Resource: splunk.InputsScript
 * 
 * Create or update scripted inputs.
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
 * import com.pulumi.splunk.InputsScript;
 * import com.pulumi.splunk.InputsScriptArgs;
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
 *         var script = new InputsScript(&#34;script&#34;, InputsScriptArgs.builder()        
 *             .interval(360)
 *             .build());
 * 
 *     }
 * }
 * ```
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 */
@ResourceType(type="splunk:index/inputsScript:InputsScript")
public class InputsScript extends com.pulumi.resources.CustomResource {
    /**
     * The app/user context that is the namespace for the resource
     * 
     */
    @Export(name="acl", refs={InputsScriptAcl.class}, tree="[0]")
    private Output<InputsScriptAcl> acl;

    /**
     * @return The app/user context that is the namespace for the resource
     * 
     */
    public Output<InputsScriptAcl> acl() {
        return this.acl;
    }
    /**
     * Specifies whether the input script is disabled.
     * 
     */
    @Export(name="disabled", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> disabled;

    /**
     * @return Specifies whether the input script is disabled.
     * 
     */
    public Output<Boolean> disabled() {
        return this.disabled;
    }
    /**
     * Sets the host for events from this input. Defaults to whatever host sent the event.
     * 
     */
    @Export(name="host", refs={String.class}, tree="[0]")
    private Output<String> host;

    /**
     * @return Sets the host for events from this input. Defaults to whatever host sent the event.
     * 
     */
    public Output<String> host() {
        return this.host;
    }
    /**
     * Sets the index for events from this input. Defaults to the main index.
     * 
     */
    @Export(name="index", refs={String.class}, tree="[0]")
    private Output<String> index;

    /**
     * @return Sets the index for events from this input. Defaults to the main index.
     * 
     */
    public Output<String> index() {
        return this.index;
    }
    /**
     * Specify an integer or cron schedule. This parameter specifies how often to execute the specified script, in seconds or a valid cron schedule. If you specify a cron schedule, the script is not executed on start-up.
     * 
     */
    @Export(name="interval", refs={Integer.class}, tree="[0]")
    private Output<Integer> interval;

    /**
     * @return Specify an integer or cron schedule. This parameter specifies how often to execute the specified script, in seconds or a valid cron schedule. If you specify a cron schedule, the script is not executed on start-up.
     * 
     */
    public Output<Integer> interval() {
        return this.interval;
    }
    /**
     * Specify the name of the scripted input.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return Specify the name of the scripted input.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * User to run the script as. If you provide a username, Splunk software generates an auth token for that user and passes it to the script.
     * 
     */
    @Export(name="passauth", refs={String.class}, tree="[0]")
    private Output<String> passauth;

    /**
     * @return User to run the script as. If you provide a username, Splunk software generates an auth token for that user and passes it to the script.
     * 
     */
    public Output<String> passauth() {
        return this.passauth;
    }
    /**
     * Specify a new name for the source field for the script.
     * 
     */
    @Export(name="renameSource", refs={String.class}, tree="[0]")
    private Output<String> renameSource;

    /**
     * @return Specify a new name for the source field for the script.
     * 
     */
    public Output<String> renameSource() {
        return this.renameSource;
    }
    /**
     * Sets the source key/field for events from this input. Defaults to the input file path.
     * Sets the source key initial value. The key is used during parsing/indexing, in particular to set the source field during indexing. It is also the source field used at search time. As a convenience, the chosen string is prepended with &#39;source::&#39;.
     * 
     */
    @Export(name="source", refs={String.class}, tree="[0]")
    private Output<String> source;

    /**
     * @return Sets the source key/field for events from this input. Defaults to the input file path.
     * Sets the source key initial value. The key is used during parsing/indexing, in particular to set the source field during indexing. It is also the source field used at search time. As a convenience, the chosen string is prepended with &#39;source::&#39;.
     * 
     */
    public Output<String> source() {
        return this.source;
    }
    /**
     * Sets the sourcetype key/field for events from this input. If unset, Splunk software picks a source type based on various aspects of the data. As a convenience, the chosen string is prepended with &#39;sourcetype::&#39;. There is no hard-coded default.
     * Sets the sourcetype key initial value. The key is used during parsing/indexing, in particular to set the source type field during indexing. It is also the source type field used at search time.
     * 
     */
    @Export(name="sourcetype", refs={String.class}, tree="[0]")
    private Output<String> sourcetype;

    /**
     * @return Sets the sourcetype key/field for events from this input. If unset, Splunk software picks a source type based on various aspects of the data. As a convenience, the chosen string is prepended with &#39;sourcetype::&#39;. There is no hard-coded default.
     * Sets the sourcetype key initial value. The key is used during parsing/indexing, in particular to set the source type field during indexing. It is also the source type field used at search time.
     * 
     */
    public Output<String> sourcetype() {
        return this.sourcetype;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public InputsScript(String name) {
        this(name, InputsScriptArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public InputsScript(String name, InputsScriptArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public InputsScript(String name, InputsScriptArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("splunk:index/inputsScript:InputsScript", name, args == null ? InputsScriptArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private InputsScript(String name, Output<String> id, @Nullable InputsScriptState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("splunk:index/inputsScript:InputsScript", name, state, makeResourceOptions(options, id));
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
    public static InputsScript get(String name, Output<String> id, @Nullable InputsScriptState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new InputsScript(name, id, state, options);
    }
}
