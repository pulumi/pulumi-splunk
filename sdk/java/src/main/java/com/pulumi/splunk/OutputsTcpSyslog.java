// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.splunk;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.splunk.OutputsTcpSyslogArgs;
import com.pulumi.splunk.Utilities;
import com.pulumi.splunk.inputs.OutputsTcpSyslogState;
import com.pulumi.splunk.outputs.OutputsTcpSyslogAcl;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import javax.annotation.Nullable;

/**
 * ## # Resource: splunk.OutputsTcpSyslog
 * 
 * Access the configuration of a forwarded server configured to provide data in standard syslog format.
 * 
 * ## Example Usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.splunk.OutputsTcpSyslog;
 * import com.pulumi.splunk.OutputsTcpSyslogArgs;
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
 *         var tcpSyslog = new OutputsTcpSyslog(&#34;tcpSyslog&#34;, OutputsTcpSyslogArgs.builder()        
 *             .priority(5)
 *             .server(&#34;new-host-1:1234&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 */
@ResourceType(type="splunk:index/outputsTcpSyslog:OutputsTcpSyslog")
public class OutputsTcpSyslog extends com.pulumi.resources.CustomResource {
    /**
     * The app/user context that is the namespace for the resource
     * 
     */
    @Export(name="acl", refs={OutputsTcpSyslogAcl.class}, tree="[0]")
    private Output<OutputsTcpSyslogAcl> acl;

    /**
     * @return The app/user context that is the namespace for the resource
     * 
     */
    public Output<OutputsTcpSyslogAcl> acl() {
        return this.acl;
    }
    /**
     * If true, disables global syslog settings.
     * 
     */
    @Export(name="disabled", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> disabled;

    /**
     * @return If true, disables global syslog settings.
     * 
     */
    public Output<Boolean> disabled() {
        return this.disabled;
    }
    /**
     * Name of the syslog output group. This is name used when creating syslog configuration in outputs.conf.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return Name of the syslog output group. This is name used when creating syslog configuration in outputs.conf.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * Sets syslog priority value. The priority value should specified as an integer. See $SPLUNK_HOME/etc/system/README/outputs.conf.spec for details.
     * 
     */
    @Export(name="priority", refs={Integer.class}, tree="[0]")
    private Output<Integer> priority;

    /**
     * @return Sets syslog priority value. The priority value should specified as an integer. See $SPLUNK_HOME/etc/system/README/outputs.conf.spec for details.
     * 
     */
    public Output<Integer> priority() {
        return this.priority;
    }
    /**
     * host:port of the server where syslog data should be sent
     * 
     */
    @Export(name="server", refs={String.class}, tree="[0]")
    private Output<String> server;

    /**
     * @return host:port of the server where syslog data should be sent
     * 
     */
    public Output<String> server() {
        return this.server;
    }
    /**
     * Specifies a rule for handling data in addition to that provided by the &#34;syslog&#34; sourcetype. By default, there is no value for syslogSourceType.
     * &lt;br&gt;This string is used as a substring match against the sourcetype key. For example, if the string is set to &#39;syslog&#39;, then all source types containing the string &#34;syslog&#34; receives this special treatment.
     * To match a source type explicitly, use the pattern &#34;sourcetype::sourcetype_name.&#34; For example
     * syslogSourcetype = sourcetype::apache_common
     * Data that is &#34;syslog&#34; or matches this setting is assumed to already be in syslog format.
     * Data that does not match the rules has a header, potentially a timestamp, and a hostname added to the front of the event. This is how Splunk software causes arbitrary log data to match syslog expectations.
     * 
     */
    @Export(name="syslogSourcetype", refs={String.class}, tree="[0]")
    private Output<String> syslogSourcetype;

    /**
     * @return Specifies a rule for handling data in addition to that provided by the &#34;syslog&#34; sourcetype. By default, there is no value for syslogSourceType.
     * &lt;br&gt;This string is used as a substring match against the sourcetype key. For example, if the string is set to &#39;syslog&#39;, then all source types containing the string &#34;syslog&#34; receives this special treatment.
     * To match a source type explicitly, use the pattern &#34;sourcetype::sourcetype_name.&#34; For example
     * syslogSourcetype = sourcetype::apache_common
     * Data that is &#34;syslog&#34; or matches this setting is assumed to already be in syslog format.
     * Data that does not match the rules has a header, potentially a timestamp, and a hostname added to the front of the event. This is how Splunk software causes arbitrary log data to match syslog expectations.
     * 
     */
    public Output<String> syslogSourcetype() {
        return this.syslogSourcetype;
    }
    /**
     * Format of timestamp to add at start of the events to be forwarded.
     * The format is a strftime-style timestamp formatting string. See $SPLUNK_HOME/etc/system/README/outputs.conf.spec for details.
     * 
     */
    @Export(name="timestampFormat", refs={String.class}, tree="[0]")
    private Output<String> timestampFormat;

    /**
     * @return Format of timestamp to add at start of the events to be forwarded.
     * The format is a strftime-style timestamp formatting string. See $SPLUNK_HOME/etc/system/README/outputs.conf.spec for details.
     * 
     */
    public Output<String> timestampFormat() {
        return this.timestampFormat;
    }
    /**
     * Protocol to use to send syslog data. Valid values: (tcp | udp ).
     * 
     */
    @Export(name="type", refs={String.class}, tree="[0]")
    private Output<String> type;

    /**
     * @return Protocol to use to send syslog data. Valid values: (tcp | udp ).
     * 
     */
    public Output<String> type() {
        return this.type;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public OutputsTcpSyslog(String name) {
        this(name, OutputsTcpSyslogArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public OutputsTcpSyslog(String name, @Nullable OutputsTcpSyslogArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public OutputsTcpSyslog(String name, @Nullable OutputsTcpSyslogArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("splunk:index/outputsTcpSyslog:OutputsTcpSyslog", name, args == null ? OutputsTcpSyslogArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private OutputsTcpSyslog(String name, Output<String> id, @Nullable OutputsTcpSyslogState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("splunk:index/outputsTcpSyslog:OutputsTcpSyslog", name, state, makeResourceOptions(options, id));
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
    public static OutputsTcpSyslog get(String name, Output<String> id, @Nullable OutputsTcpSyslogState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new OutputsTcpSyslog(name, id, state, options);
    }
}
