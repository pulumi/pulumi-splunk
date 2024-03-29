// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.splunk;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.splunk.inputs.AppsLocalAclArgs;
import java.lang.Boolean;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class AppsLocalArgs extends com.pulumi.resources.ResourceArgs {

    public static final AppsLocalArgs Empty = new AppsLocalArgs();

    /**
     * The app/user context that is the namespace for the resource
     * 
     */
    @Import(name="acl")
    private @Nullable Output<AppsLocalAclArgs> acl;

    /**
     * @return The app/user context that is the namespace for the resource
     * 
     */
    public Optional<Output<AppsLocalAclArgs>> acl() {
        return Optional.ofNullable(this.acl);
    }

    /**
     * Splunkbase session token for operations like install and update that require login. Use auth or session when installing or updating an app through Splunkbase.
     * 
     */
    @Import(name="auth")
    private @Nullable Output<String> auth;

    /**
     * @return Splunkbase session token for operations like install and update that require login. Use auth or session when installing or updating an app through Splunkbase.
     * 
     */
    public Optional<Output<String>> auth() {
        return Optional.ofNullable(this.auth);
    }

    /**
     * For apps posted to Splunkbase, use your Splunk account username. For internal apps, include your name and contact information.
     * 
     */
    @Import(name="author")
    private @Nullable Output<String> author;

    /**
     * @return For apps posted to Splunkbase, use your Splunk account username. For internal apps, include your name and contact information.
     * 
     */
    public Optional<Output<String>> author() {
        return Optional.ofNullable(this.author);
    }

    /**
     * Custom setup complete indication:
     * &lt;br&gt;true = Custom app setup complete.
     * &lt;br&gt;false = Custom app setup not complete.
     * 
     */
    @Import(name="configured")
    private @Nullable Output<Boolean> configured;

    /**
     * @return Custom setup complete indication:
     * &lt;br&gt;true = Custom app setup complete.
     * &lt;br&gt;false = Custom app setup not complete.
     * 
     */
    public Optional<Output<Boolean>> configured() {
        return Optional.ofNullable(this.configured);
    }

    /**
     * Short app description also displayed below the app title in Splunk Web Launcher.
     * 
     */
    @Import(name="description")
    private @Nullable Output<String> description;

    /**
     * @return Short app description also displayed below the app title in Splunk Web Launcher.
     * 
     */
    public Optional<Output<String>> description() {
        return Optional.ofNullable(this.description);
    }

    /**
     * Custom app name. Overrides name when installing an app from a file where filename is set to true. See also filename.
     * 
     */
    @Import(name="explicitAppname")
    private @Nullable Output<String> explicitAppname;

    /**
     * @return Custom app name. Overrides name when installing an app from a file where filename is set to true. See also filename.
     * 
     */
    public Optional<Output<String>> explicitAppname() {
        return Optional.ofNullable(this.explicitAppname);
    }

    /**
     * Indicates whether to use the name value as the app source location.
     * &lt;br&gt;true indicates that name is a path to a file to install.
     * &lt;br&gt;false indicates that name is the literal app name and that the app is created from Splunkbase using a template.
     * 
     */
    @Import(name="filename")
    private @Nullable Output<Boolean> filename;

    /**
     * @return Indicates whether to use the name value as the app source location.
     * &lt;br&gt;true indicates that name is a path to a file to install.
     * &lt;br&gt;false indicates that name is the literal app name and that the app is created from Splunkbase using a template.
     * 
     */
    public Optional<Output<Boolean>> filename() {
        return Optional.ofNullable(this.filename);
    }

    /**
     * App name displayed in Splunk Web, from five to eighty characters excluding the prefix &#34;Splunk for&#34;.
     * 
     */
    @Import(name="label")
    private @Nullable Output<String> label;

    /**
     * @return App name displayed in Splunk Web, from five to eighty characters excluding the prefix &#34;Splunk for&#34;.
     * 
     */
    public Optional<Output<String>> label() {
        return Optional.ofNullable(this.label);
    }

    /**
     * Literal app name or path for the file to install, depending on the value of filename.
     * &lt;br&gt;filename = false indicates that name is the literal app name and that the app is created from Splunkbase using a template.
     * &lt;br&gt;filename = true indicates that name is the URL or path to the local .tar, .tgz or .spl file. If name is the Splunkbase URL, set auth or session to authenticate the request.
     * The app folder name cannot include spaces or special characters.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return Literal app name or path for the file to install, depending on the value of filename.
     * &lt;br&gt;filename = false indicates that name is the literal app name and that the app is created from Splunkbase using a template.
     * &lt;br&gt;filename = true indicates that name is the URL or path to the local .tar, .tgz or .spl file. If name is the Splunkbase URL, set auth or session to authenticate the request.
     * The app folder name cannot include spaces or special characters.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * Login session token for installing or updating an app on Splunkbase. Alternatively, use auth.
     * 
     */
    @Import(name="session")
    private @Nullable Output<String> session;

    /**
     * @return Login session token for installing or updating an app on Splunkbase. Alternatively, use auth.
     * 
     */
    public Optional<Output<String>> session() {
        return Optional.ofNullable(this.session);
    }

    /**
     * File-based update indication:
     * &lt;br&gt;true specifies that filename should be used to update an existing app. If not specified, update defaults to
     * &lt;br&gt;false, which indicates that filename should not be used to update an existing app.
     * 
     */
    @Import(name="update")
    private @Nullable Output<Boolean> update;

    /**
     * @return File-based update indication:
     * &lt;br&gt;true specifies that filename should be used to update an existing app. If not specified, update defaults to
     * &lt;br&gt;false, which indicates that filename should not be used to update an existing app.
     * 
     */
    public Optional<Output<Boolean>> update() {
        return Optional.ofNullable(this.update);
    }

    /**
     * App version.
     * 
     */
    @Import(name="version")
    private @Nullable Output<String> version;

    /**
     * @return App version.
     * 
     */
    public Optional<Output<String>> version() {
        return Optional.ofNullable(this.version);
    }

    /**
     * Indicates whether the app is visible and navigable from Splunk Web.
     * &lt;br&gt;true = App is visible and navigable.
     * &lt;br&gt;false = App is not visible or navigable.
     * 
     */
    @Import(name="visible")
    private @Nullable Output<Boolean> visible;

    /**
     * @return Indicates whether the app is visible and navigable from Splunk Web.
     * &lt;br&gt;true = App is visible and navigable.
     * &lt;br&gt;false = App is not visible or navigable.
     * 
     */
    public Optional<Output<Boolean>> visible() {
        return Optional.ofNullable(this.visible);
    }

    private AppsLocalArgs() {}

    private AppsLocalArgs(AppsLocalArgs $) {
        this.acl = $.acl;
        this.auth = $.auth;
        this.author = $.author;
        this.configured = $.configured;
        this.description = $.description;
        this.explicitAppname = $.explicitAppname;
        this.filename = $.filename;
        this.label = $.label;
        this.name = $.name;
        this.session = $.session;
        this.update = $.update;
        this.version = $.version;
        this.visible = $.visible;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(AppsLocalArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private AppsLocalArgs $;

        public Builder() {
            $ = new AppsLocalArgs();
        }

        public Builder(AppsLocalArgs defaults) {
            $ = new AppsLocalArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param acl The app/user context that is the namespace for the resource
         * 
         * @return builder
         * 
         */
        public Builder acl(@Nullable Output<AppsLocalAclArgs> acl) {
            $.acl = acl;
            return this;
        }

        /**
         * @param acl The app/user context that is the namespace for the resource
         * 
         * @return builder
         * 
         */
        public Builder acl(AppsLocalAclArgs acl) {
            return acl(Output.of(acl));
        }

        /**
         * @param auth Splunkbase session token for operations like install and update that require login. Use auth or session when installing or updating an app through Splunkbase.
         * 
         * @return builder
         * 
         */
        public Builder auth(@Nullable Output<String> auth) {
            $.auth = auth;
            return this;
        }

        /**
         * @param auth Splunkbase session token for operations like install and update that require login. Use auth or session when installing or updating an app through Splunkbase.
         * 
         * @return builder
         * 
         */
        public Builder auth(String auth) {
            return auth(Output.of(auth));
        }

        /**
         * @param author For apps posted to Splunkbase, use your Splunk account username. For internal apps, include your name and contact information.
         * 
         * @return builder
         * 
         */
        public Builder author(@Nullable Output<String> author) {
            $.author = author;
            return this;
        }

        /**
         * @param author For apps posted to Splunkbase, use your Splunk account username. For internal apps, include your name and contact information.
         * 
         * @return builder
         * 
         */
        public Builder author(String author) {
            return author(Output.of(author));
        }

        /**
         * @param configured Custom setup complete indication:
         * &lt;br&gt;true = Custom app setup complete.
         * &lt;br&gt;false = Custom app setup not complete.
         * 
         * @return builder
         * 
         */
        public Builder configured(@Nullable Output<Boolean> configured) {
            $.configured = configured;
            return this;
        }

        /**
         * @param configured Custom setup complete indication:
         * &lt;br&gt;true = Custom app setup complete.
         * &lt;br&gt;false = Custom app setup not complete.
         * 
         * @return builder
         * 
         */
        public Builder configured(Boolean configured) {
            return configured(Output.of(configured));
        }

        /**
         * @param description Short app description also displayed below the app title in Splunk Web Launcher.
         * 
         * @return builder
         * 
         */
        public Builder description(@Nullable Output<String> description) {
            $.description = description;
            return this;
        }

        /**
         * @param description Short app description also displayed below the app title in Splunk Web Launcher.
         * 
         * @return builder
         * 
         */
        public Builder description(String description) {
            return description(Output.of(description));
        }

        /**
         * @param explicitAppname Custom app name. Overrides name when installing an app from a file where filename is set to true. See also filename.
         * 
         * @return builder
         * 
         */
        public Builder explicitAppname(@Nullable Output<String> explicitAppname) {
            $.explicitAppname = explicitAppname;
            return this;
        }

        /**
         * @param explicitAppname Custom app name. Overrides name when installing an app from a file where filename is set to true. See also filename.
         * 
         * @return builder
         * 
         */
        public Builder explicitAppname(String explicitAppname) {
            return explicitAppname(Output.of(explicitAppname));
        }

        /**
         * @param filename Indicates whether to use the name value as the app source location.
         * &lt;br&gt;true indicates that name is a path to a file to install.
         * &lt;br&gt;false indicates that name is the literal app name and that the app is created from Splunkbase using a template.
         * 
         * @return builder
         * 
         */
        public Builder filename(@Nullable Output<Boolean> filename) {
            $.filename = filename;
            return this;
        }

        /**
         * @param filename Indicates whether to use the name value as the app source location.
         * &lt;br&gt;true indicates that name is a path to a file to install.
         * &lt;br&gt;false indicates that name is the literal app name and that the app is created from Splunkbase using a template.
         * 
         * @return builder
         * 
         */
        public Builder filename(Boolean filename) {
            return filename(Output.of(filename));
        }

        /**
         * @param label App name displayed in Splunk Web, from five to eighty characters excluding the prefix &#34;Splunk for&#34;.
         * 
         * @return builder
         * 
         */
        public Builder label(@Nullable Output<String> label) {
            $.label = label;
            return this;
        }

        /**
         * @param label App name displayed in Splunk Web, from five to eighty characters excluding the prefix &#34;Splunk for&#34;.
         * 
         * @return builder
         * 
         */
        public Builder label(String label) {
            return label(Output.of(label));
        }

        /**
         * @param name Literal app name or path for the file to install, depending on the value of filename.
         * &lt;br&gt;filename = false indicates that name is the literal app name and that the app is created from Splunkbase using a template.
         * &lt;br&gt;filename = true indicates that name is the URL or path to the local .tar, .tgz or .spl file. If name is the Splunkbase URL, set auth or session to authenticate the request.
         * The app folder name cannot include spaces or special characters.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name Literal app name or path for the file to install, depending on the value of filename.
         * &lt;br&gt;filename = false indicates that name is the literal app name and that the app is created from Splunkbase using a template.
         * &lt;br&gt;filename = true indicates that name is the URL or path to the local .tar, .tgz or .spl file. If name is the Splunkbase URL, set auth or session to authenticate the request.
         * The app folder name cannot include spaces or special characters.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param session Login session token for installing or updating an app on Splunkbase. Alternatively, use auth.
         * 
         * @return builder
         * 
         */
        public Builder session(@Nullable Output<String> session) {
            $.session = session;
            return this;
        }

        /**
         * @param session Login session token for installing or updating an app on Splunkbase. Alternatively, use auth.
         * 
         * @return builder
         * 
         */
        public Builder session(String session) {
            return session(Output.of(session));
        }

        /**
         * @param update File-based update indication:
         * &lt;br&gt;true specifies that filename should be used to update an existing app. If not specified, update defaults to
         * &lt;br&gt;false, which indicates that filename should not be used to update an existing app.
         * 
         * @return builder
         * 
         */
        public Builder update(@Nullable Output<Boolean> update) {
            $.update = update;
            return this;
        }

        /**
         * @param update File-based update indication:
         * &lt;br&gt;true specifies that filename should be used to update an existing app. If not specified, update defaults to
         * &lt;br&gt;false, which indicates that filename should not be used to update an existing app.
         * 
         * @return builder
         * 
         */
        public Builder update(Boolean update) {
            return update(Output.of(update));
        }

        /**
         * @param version App version.
         * 
         * @return builder
         * 
         */
        public Builder version(@Nullable Output<String> version) {
            $.version = version;
            return this;
        }

        /**
         * @param version App version.
         * 
         * @return builder
         * 
         */
        public Builder version(String version) {
            return version(Output.of(version));
        }

        /**
         * @param visible Indicates whether the app is visible and navigable from Splunk Web.
         * &lt;br&gt;true = App is visible and navigable.
         * &lt;br&gt;false = App is not visible or navigable.
         * 
         * @return builder
         * 
         */
        public Builder visible(@Nullable Output<Boolean> visible) {
            $.visible = visible;
            return this;
        }

        /**
         * @param visible Indicates whether the app is visible and navigable from Splunk Web.
         * &lt;br&gt;true = App is visible and navigable.
         * &lt;br&gt;false = App is not visible or navigable.
         * 
         * @return builder
         * 
         */
        public Builder visible(Boolean visible) {
            return visible(Output.of(visible));
        }

        public AppsLocalArgs build() {
            return $;
        }
    }

}
