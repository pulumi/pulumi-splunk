// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.splunk;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.splunk.InputsUdpArgs;
import com.pulumi.splunk.Utilities;
import com.pulumi.splunk.inputs.InputsUdpState;
import com.pulumi.splunk.outputs.InputsUdpAcl;
import java.lang.Boolean;
import java.lang.String;
import javax.annotation.Nullable;

/**
 * ## # Resource: splunk.InputsTcpRaw
 * 
 * Create and manage UDP data inputs.
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
 * import com.pulumi.splunk.InputsUdp;
 * import com.pulumi.splunk.InputsUdpArgs;
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
 *         var udp = new InputsUdp(&#34;udp&#34;, InputsUdpArgs.builder()        
 *             .name(&#34;41000&#34;)
 *             .index(&#34;main&#34;)
 *             .source(&#34;new&#34;)
 *             .sourcetype(&#34;new&#34;)
 *             .disabled(false)
 *             .build());
 * 
 *     }
 * }
 * ```
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 */
@ResourceType(type="splunk:index/inputsUdp:InputsUdp")
public class InputsUdp extends com.pulumi.resources.CustomResource {
    /**
     * The app/user context that is the namespace for the resource
     * 
     */
    @Export(name="acl", refs={InputsUdpAcl.class}, tree="[0]")
    private Output<InputsUdpAcl> acl;

    /**
     * @return The app/user context that is the namespace for the resource
     * 
     */
    public Output<InputsUdpAcl> acl() {
        return this.acl;
    }
    /**
     * Valid values: (ip | dns | none)
     * Set the host for the remote server that is sending data.
     * ip sets the host to the IP address of the remote server sending data.
     * dns sets the host to the reverse DNS entry for the IP address of the remote server sending data.
     * none leaves the host as specified in inputs.conf, which is typically the Splunk system hostname.
     * Default value is dns.
     * 
     */
    @Export(name="connectionHost", refs={String.class}, tree="[0]")
    private Output<String> connectionHost;

    /**
     * @return Valid values: (ip | dns | none)
     * Set the host for the remote server that is sending data.
     * ip sets the host to the IP address of the remote server sending data.
     * dns sets the host to the reverse DNS entry for the IP address of the remote server sending data.
     * none leaves the host as specified in inputs.conf, which is typically the Splunk system hostname.
     * Default value is dns.
     * 
     */
    public Output<String> connectionHost() {
        return this.connectionHost;
    }
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
     * The value to populate in the host field for incoming events. This is used during parsing/indexing, in particular to set the host field. It is also the host field used at search time.
     * 
     */
    @Export(name="host", refs={String.class}, tree="[0]")
    private Output<String> host;

    /**
     * @return The value to populate in the host field for incoming events. This is used during parsing/indexing, in particular to set the host field. It is also the host field used at search time.
     * 
     */
    public Output<String> host() {
        return this.host;
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
     * The UDP port that this input should listen on.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return The UDP port that this input should listen on.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * If set to true, prevents Splunk software from prepending a timestamp and hostname to incoming events.
     * 
     */
    @Export(name="noAppendingTimestamp", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> noAppendingTimestamp;

    /**
     * @return If set to true, prevents Splunk software from prepending a timestamp and hostname to incoming events.
     * 
     */
    public Output<Boolean> noAppendingTimestamp() {
        return this.noAppendingTimestamp;
    }
    /**
     * If set to true, Splunk software does not remove the priority field from incoming syslog events.
     * 
     */
    @Export(name="noPriorityStripping", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> noPriorityStripping;

    /**
     * @return If set to true, Splunk software does not remove the priority field from incoming syslog events.
     * 
     */
    public Output<Boolean> noPriorityStripping() {
        return this.noPriorityStripping;
    }
    /**
     * Which queue events from this input should be sent to. Generally this does not need to be changed.
     * 
     */
    @Export(name="queue", refs={String.class}, tree="[0]")
    private Output<String> queue;

    /**
     * @return Which queue events from this input should be sent to. Generally this does not need to be changed.
     * 
     */
    public Output<String> queue() {
        return this.queue;
    }
    /**
     * Restrict incoming connections on this port to the host specified here.
     * If this is not set, the value specified in [udp://&lt;remote server&gt;:&lt;port&gt;] in inputs.conf is used.
     * 
     */
    @Export(name="restrictToHost", refs={String.class}, tree="[0]")
    private Output<String> restrictToHost;

    /**
     * @return Restrict incoming connections on this port to the host specified here.
     * If this is not set, the value specified in [udp://&lt;remote server&gt;:&lt;port&gt;] in inputs.conf is used.
     * 
     */
    public Output<String> restrictToHost() {
        return this.restrictToHost;
    }
    /**
     * The value to populate in the source field for incoming events. The same source should not be used for multiple data inputs.
     * 
     */
    @Export(name="source", refs={String.class}, tree="[0]")
    private Output<String> source;

    /**
     * @return The value to populate in the source field for incoming events. The same source should not be used for multiple data inputs.
     * 
     */
    public Output<String> source() {
        return this.source;
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
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public InputsUdp(String name) {
        this(name, InputsUdpArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public InputsUdp(String name, @Nullable InputsUdpArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public InputsUdp(String name, @Nullable InputsUdpArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("splunk:index/inputsUdp:InputsUdp", name, args == null ? InputsUdpArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private InputsUdp(String name, Output<String> id, @Nullable InputsUdpState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("splunk:index/inputsUdp:InputsUdp", name, state, makeResourceOptions(options, id));
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
    public static InputsUdp get(String name, Output<String> id, @Nullable InputsUdpState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new InputsUdp(name, id, state, options);
    }
}
