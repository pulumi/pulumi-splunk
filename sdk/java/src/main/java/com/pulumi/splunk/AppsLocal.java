// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.splunk;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.splunk.AppsLocalArgs;
import com.pulumi.splunk.Utilities;
import com.pulumi.splunk.inputs.AppsLocalState;
import com.pulumi.splunk.outputs.AppsLocalAcl;
import java.lang.Boolean;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * ## # Resource: splunk.AppsLocal
 * 
 * Create, install and manage apps on your Splunk instance
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
 * import com.pulumi.splunk.AppsLocal;
 * import com.pulumi.splunk.AppsLocalArgs;
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
 *         var amazonConnectApp = new AppsLocal("amazonConnectApp", AppsLocalArgs.builder()
 *             .filename(true)
 *             .name("/usr/home/amazon_connect_app_for_splunk-0.0.1.tar.gz")
 *             .explicitAppname("amazon_connect_app_for_splunk")
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 */
@ResourceType(type="splunk:index/appsLocal:AppsLocal")
public class AppsLocal extends com.pulumi.resources.CustomResource {
    /**
     * The app/user context that is the namespace for the resource
     * 
     */
    @Export(name="acl", refs={AppsLocalAcl.class}, tree="[0]")
    private Output<AppsLocalAcl> acl;

    /**
     * @return The app/user context that is the namespace for the resource
     * 
     */
    public Output<AppsLocalAcl> acl() {
        return this.acl;
    }
    /**
     * Splunkbase session token for operations like install and update that require login. Use auth or session when installing or updating an app through Splunkbase.
     * 
     */
    @Export(name="auth", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> auth;

    /**
     * @return Splunkbase session token for operations like install and update that require login. Use auth or session when installing or updating an app through Splunkbase.
     * 
     */
    public Output<Optional<String>> auth() {
        return Codegen.optional(this.auth);
    }
    /**
     * For apps posted to Splunkbase, use your Splunk account username. For internal apps, include your name and contact information.
     * 
     */
    @Export(name="author", refs={String.class}, tree="[0]")
    private Output<String> author;

    /**
     * @return For apps posted to Splunkbase, use your Splunk account username. For internal apps, include your name and contact information.
     * 
     */
    public Output<String> author() {
        return this.author;
    }
    /**
     * Custom setup complete indication:
     * &lt;br&gt;true = Custom app setup complete.
     * &lt;br&gt;false = Custom app setup not complete.
     * 
     */
    @Export(name="configured", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> configured;

    /**
     * @return Custom setup complete indication:
     * &lt;br&gt;true = Custom app setup complete.
     * &lt;br&gt;false = Custom app setup not complete.
     * 
     */
    public Output<Boolean> configured() {
        return this.configured;
    }
    /**
     * Short app description also displayed below the app title in Splunk Web Launcher.
     * 
     */
    @Export(name="description", refs={String.class}, tree="[0]")
    private Output<String> description;

    /**
     * @return Short app description also displayed below the app title in Splunk Web Launcher.
     * 
     */
    public Output<String> description() {
        return this.description;
    }
    /**
     * Custom app name. Overrides name when installing an app from a file where filename is set to true. See also filename.
     * 
     */
    @Export(name="explicitAppname", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> explicitAppname;

    /**
     * @return Custom app name. Overrides name when installing an app from a file where filename is set to true. See also filename.
     * 
     */
    public Output<Optional<String>> explicitAppname() {
        return Codegen.optional(this.explicitAppname);
    }
    /**
     * Indicates whether to use the name value as the app source location.
     * &lt;br&gt;true indicates that name is a path to a file to install.
     * &lt;br&gt;false indicates that name is the literal app name and that the app is created from Splunkbase using a template.
     * 
     */
    @Export(name="filename", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> filename;

    /**
     * @return Indicates whether to use the name value as the app source location.
     * &lt;br&gt;true indicates that name is a path to a file to install.
     * &lt;br&gt;false indicates that name is the literal app name and that the app is created from Splunkbase using a template.
     * 
     */
    public Output<Optional<Boolean>> filename() {
        return Codegen.optional(this.filename);
    }
    /**
     * App name displayed in Splunk Web, from five to eighty characters excluding the prefix &#34;Splunk for&#34;.
     * 
     */
    @Export(name="label", refs={String.class}, tree="[0]")
    private Output<String> label;

    /**
     * @return App name displayed in Splunk Web, from five to eighty characters excluding the prefix &#34;Splunk for&#34;.
     * 
     */
    public Output<String> label() {
        return this.label;
    }
    /**
     * Literal app name or path for the file to install, depending on the value of filename.
     * &lt;br&gt;filename = false indicates that name is the literal app name and that the app is created from Splunkbase using a template.
     * &lt;br&gt;filename = true indicates that name is the URL or path to the local .tar, .tgz or .spl file. If name is the Splunkbase URL, set auth or session to authenticate the request.
     * The app folder name cannot include spaces or special characters.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return Literal app name or path for the file to install, depending on the value of filename.
     * &lt;br&gt;filename = false indicates that name is the literal app name and that the app is created from Splunkbase using a template.
     * &lt;br&gt;filename = true indicates that name is the URL or path to the local .tar, .tgz or .spl file. If name is the Splunkbase URL, set auth or session to authenticate the request.
     * The app folder name cannot include spaces or special characters.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * Login session token for installing or updating an app on Splunkbase. Alternatively, use auth.
     * 
     */
    @Export(name="session", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> session;

    /**
     * @return Login session token for installing or updating an app on Splunkbase. Alternatively, use auth.
     * 
     */
    public Output<Optional<String>> session() {
        return Codegen.optional(this.session);
    }
    /**
     * File-based update indication:
     * &lt;br&gt;true specifies that filename should be used to update an existing app. If not specified, update defaults to
     * &lt;br&gt;false, which indicates that filename should not be used to update an existing app.
     * 
     */
    @Export(name="update", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> update;

    /**
     * @return File-based update indication:
     * &lt;br&gt;true specifies that filename should be used to update an existing app. If not specified, update defaults to
     * &lt;br&gt;false, which indicates that filename should not be used to update an existing app.
     * 
     */
    public Output<Optional<Boolean>> update() {
        return Codegen.optional(this.update);
    }
    /**
     * App version.
     * 
     */
    @Export(name="version", refs={String.class}, tree="[0]")
    private Output<String> version;

    /**
     * @return App version.
     * 
     */
    public Output<String> version() {
        return this.version;
    }
    /**
     * Indicates whether the app is visible and navigable from Splunk Web.
     * &lt;br&gt;true = App is visible and navigable.
     * &lt;br&gt;false = App is not visible or navigable.
     * 
     */
    @Export(name="visible", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> visible;

    /**
     * @return Indicates whether the app is visible and navigable from Splunk Web.
     * &lt;br&gt;true = App is visible and navigable.
     * &lt;br&gt;false = App is not visible or navigable.
     * 
     */
    public Output<Boolean> visible() {
        return this.visible;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public AppsLocal(java.lang.String name) {
        this(name, AppsLocalArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public AppsLocal(java.lang.String name, @Nullable AppsLocalArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public AppsLocal(java.lang.String name, @Nullable AppsLocalArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("splunk:index/appsLocal:AppsLocal", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private AppsLocal(java.lang.String name, Output<java.lang.String> id, @Nullable AppsLocalState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("splunk:index/appsLocal:AppsLocal", name, state, makeResourceOptions(options, id), false);
    }

    private static AppsLocalArgs makeArgs(@Nullable AppsLocalArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? AppsLocalArgs.Empty : args;
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
    public static AppsLocal get(java.lang.String name, Output<java.lang.String> id, @Nullable AppsLocalState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new AppsLocal(name, id, state, options);
    }
}
