// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.splunk;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.splunk.InputsMonitorArgs;
import com.pulumi.splunk.Utilities;
import com.pulumi.splunk.inputs.InputsMonitorState;
import com.pulumi.splunk.outputs.InputsMonitorAcl;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import javax.annotation.Nullable;

/**
 * # Resource: splunk.InputsMonitor
 * Create or update a new file or directory monitor input.
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
 * import com.pulumi.splunk.InputsMonitor;
 * import com.pulumi.splunk.InputsMonitorArgs;
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
 *         var monitor = new InputsMonitor("monitor", InputsMonitorArgs.builder()
 *             .name("opt/splunk/var/log/splunk/health.log")
 *             .recursive(true)
 *             .sourcetype("text")
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 */
@ResourceType(type="splunk:index/inputsMonitor:InputsMonitor")
public class InputsMonitor extends com.pulumi.resources.CustomResource {
    /**
     * The app/user context that is the namespace for the resource
     * 
     */
    @Export(name="acl", refs={InputsMonitorAcl.class}, tree="[0]")
    private Output<InputsMonitorAcl> acl;

    /**
     * @return The app/user context that is the namespace for the resource
     * 
     */
    public Output<InputsMonitorAcl> acl() {
        return this.acl;
    }
    /**
     * Specify a regular expression for a file path. The file path that matches this regular expression is not indexed.
     * 
     */
    @Export(name="blacklist", refs={String.class}, tree="[0]")
    private Output<String> blacklist;

    /**
     * @return Specify a regular expression for a file path. The file path that matches this regular expression is not indexed.
     * 
     */
    public Output<String> blacklist() {
        return this.blacklist;
    }
    /**
     * A string that modifies the file tracking identity for files in this input. The magic value &lt;SOURCE&gt; invokes special behavior.
     * 
     */
    @Export(name="crcSalt", refs={String.class}, tree="[0]")
    private Output<String> crcSalt;

    /**
     * @return A string that modifies the file tracking identity for files in this input. The magic value &lt;SOURCE&gt; invokes special behavior.
     * 
     */
    public Output<String> crcSalt() {
        return this.crcSalt;
    }
    /**
     * Indicates if input monitoring is disabled.
     * 
     */
    @Export(name="disabled", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> disabled;

    /**
     * @return Indicates if input monitoring is disabled.
     * 
     */
    public Output<Boolean> disabled() {
        return this.disabled;
    }
    /**
     * If set to true, files that are seen for the first time is read from the end.
     * 
     */
    @Export(name="followTail", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> followTail;

    /**
     * @return If set to true, files that are seen for the first time is read from the end.
     * 
     */
    public Output<Boolean> followTail() {
        return this.followTail;
    }
    /**
     * The value to populate in the host field for events from this data input.
     * 
     */
    @Export(name="host", refs={String.class}, tree="[0]")
    private Output<String> host;

    /**
     * @return The value to populate in the host field for events from this data input.
     * 
     */
    public Output<String> host() {
        return this.host;
    }
    /**
     * Specify a regular expression for a file path. If the path for a file matches this regular expression, the captured value is used to populate the host field for events from this data input. The regular expression must have one capture group.
     * 
     */
    @Export(name="hostRegex", refs={String.class}, tree="[0]")
    private Output<String> hostRegex;

    /**
     * @return Specify a regular expression for a file path. If the path for a file matches this regular expression, the captured value is used to populate the host field for events from this data input. The regular expression must have one capture group.
     * 
     */
    public Output<String> hostRegex() {
        return this.hostRegex;
    }
    /**
     * Use the specified slash-separate segment of the filepath as the host field value.
     * 
     */
    @Export(name="hostSegment", refs={Integer.class}, tree="[0]")
    private Output<Integer> hostSegment;

    /**
     * @return Use the specified slash-separate segment of the filepath as the host field value.
     * 
     */
    public Output<Integer> hostSegment() {
        return this.hostSegment;
    }
    /**
     * Specify a time value. If the modification time of a file being monitored falls outside of this rolling time window, the file is no longer being monitored.
     * 
     */
    @Export(name="ignoreOlderThan", refs={String.class}, tree="[0]")
    private Output<String> ignoreOlderThan;

    /**
     * @return Specify a time value. If the modification time of a file being monitored falls outside of this rolling time window, the file is no longer being monitored.
     * 
     */
    public Output<String> ignoreOlderThan() {
        return this.ignoreOlderThan;
    }
    /**
     * Which index events from this input should be stored in. Defaults to default.
     * 
     */
    @Export(name="index", refs={String.class}, tree="[0]")
    private Output<String> index;

    /**
     * @return Which index events from this input should be stored in. Defaults to default.
     * 
     */
    public Output<String> index() {
        return this.index;
    }
    /**
     * The file or directory path to monitor on the system.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return The file or directory path to monitor on the system.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * Setting this to false prevents monitoring of any subdirectories encountered within this data input.
     * 
     */
    @Export(name="recursive", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> recursive;

    /**
     * @return Setting this to false prevents monitoring of any subdirectories encountered within this data input.
     * 
     */
    public Output<Boolean> recursive() {
        return this.recursive;
    }
    /**
     * The value to populate in the source field for events from this data input. The same source should not be used for multiple data inputs.
     * 
     */
    @Export(name="renameSource", refs={String.class}, tree="[0]")
    private Output<String> renameSource;

    /**
     * @return The value to populate in the source field for events from this data input. The same source should not be used for multiple data inputs.
     * 
     */
    public Output<String> renameSource() {
        return this.renameSource;
    }
    /**
     * The value to populate in the sourcetype field for incoming events.
     * 
     */
    @Export(name="sourcetype", refs={String.class}, tree="[0]")
    private Output<String> sourcetype;

    /**
     * @return The value to populate in the sourcetype field for incoming events.
     * 
     */
    public Output<String> sourcetype() {
        return this.sourcetype;
    }
    /**
     * When Splunk software reaches the end of a file that is being read, the file is kept open for a minimum of the number of seconds specified in this value. After this period has elapsed, the file is checked again for more data.
     * 
     */
    @Export(name="timeBeforeClose", refs={Integer.class}, tree="[0]")
    private Output<Integer> timeBeforeClose;

    /**
     * @return When Splunk software reaches the end of a file that is being read, the file is kept open for a minimum of the number of seconds specified in this value. After this period has elapsed, the file is checked again for more data.
     * 
     */
    public Output<Integer> timeBeforeClose() {
        return this.timeBeforeClose;
    }
    /**
     * Specify a regular expression for a file path. Only file paths that match this regular expression are indexed.
     * 
     */
    @Export(name="whitelist", refs={String.class}, tree="[0]")
    private Output<String> whitelist;

    /**
     * @return Specify a regular expression for a file path. Only file paths that match this regular expression are indexed.
     * 
     */
    public Output<String> whitelist() {
        return this.whitelist;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public InputsMonitor(java.lang.String name) {
        this(name, InputsMonitorArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public InputsMonitor(java.lang.String name, @Nullable InputsMonitorArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public InputsMonitor(java.lang.String name, @Nullable InputsMonitorArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("splunk:index/inputsMonitor:InputsMonitor", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private InputsMonitor(java.lang.String name, Output<java.lang.String> id, @Nullable InputsMonitorState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("splunk:index/inputsMonitor:InputsMonitor", name, state, makeResourceOptions(options, id), false);
    }

    private static InputsMonitorArgs makeArgs(@Nullable InputsMonitorArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? InputsMonitorArgs.Empty : args;
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
    public static InputsMonitor get(java.lang.String name, Output<java.lang.String> id, @Nullable InputsMonitorState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new InputsMonitor(name, id, state, options);
    }
}
