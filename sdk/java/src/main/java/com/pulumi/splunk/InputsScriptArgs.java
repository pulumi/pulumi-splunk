// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.splunk;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.splunk.inputs.InputsScriptAclArgs;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class InputsScriptArgs extends com.pulumi.resources.ResourceArgs {

    public static final InputsScriptArgs Empty = new InputsScriptArgs();

    /**
     * The app/user context that is the namespace for the resource
     * 
     */
    @Import(name="acl")
    private @Nullable Output<InputsScriptAclArgs> acl;

    /**
     * @return The app/user context that is the namespace for the resource
     * 
     */
    public Optional<Output<InputsScriptAclArgs>> acl() {
        return Optional.ofNullable(this.acl);
    }

    /**
     * Specifies whether the input script is disabled.
     * 
     */
    @Import(name="disabled")
    private @Nullable Output<Boolean> disabled;

    /**
     * @return Specifies whether the input script is disabled.
     * 
     */
    public Optional<Output<Boolean>> disabled() {
        return Optional.ofNullable(this.disabled);
    }

    /**
     * Sets the host for events from this input. Defaults to whatever host sent the event.
     * 
     */
    @Import(name="host")
    private @Nullable Output<String> host;

    /**
     * @return Sets the host for events from this input. Defaults to whatever host sent the event.
     * 
     */
    public Optional<Output<String>> host() {
        return Optional.ofNullable(this.host);
    }

    /**
     * Sets the index for events from this input. Defaults to the main index.
     * 
     */
    @Import(name="index")
    private @Nullable Output<String> index;

    /**
     * @return Sets the index for events from this input. Defaults to the main index.
     * 
     */
    public Optional<Output<String>> index() {
        return Optional.ofNullable(this.index);
    }

    /**
     * Specify an integer or cron schedule. This parameter specifies how often to execute the specified script, in seconds or a valid cron schedule. If you specify a cron schedule, the script is not executed on start-up.
     * 
     */
    @Import(name="interval", required=true)
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
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return Specify the name of the scripted input.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * User to run the script as. If you provide a username, Splunk software generates an auth token for that user and passes it to the script.
     * 
     */
    @Import(name="passauth")
    private @Nullable Output<String> passauth;

    /**
     * @return User to run the script as. If you provide a username, Splunk software generates an auth token for that user and passes it to the script.
     * 
     */
    public Optional<Output<String>> passauth() {
        return Optional.ofNullable(this.passauth);
    }

    /**
     * Specify a new name for the source field for the script.
     * 
     */
    @Import(name="renameSource")
    private @Nullable Output<String> renameSource;

    /**
     * @return Specify a new name for the source field for the script.
     * 
     */
    public Optional<Output<String>> renameSource() {
        return Optional.ofNullable(this.renameSource);
    }

    /**
     * Sets the source key/field for events from this input. Defaults to the input file path.
     * Sets the source key initial value. The key is used during parsing/indexing, in particular to set the source field during indexing. It is also the source field used at search time. As a convenience, the chosen string is prepended with &#39;source::&#39;.
     * 
     */
    @Import(name="source")
    private @Nullable Output<String> source;

    /**
     * @return Sets the source key/field for events from this input. Defaults to the input file path.
     * Sets the source key initial value. The key is used during parsing/indexing, in particular to set the source field during indexing. It is also the source field used at search time. As a convenience, the chosen string is prepended with &#39;source::&#39;.
     * 
     */
    public Optional<Output<String>> source() {
        return Optional.ofNullable(this.source);
    }

    /**
     * Sets the sourcetype key/field for events from this input. If unset, Splunk software picks a source type based on various aspects of the data. As a convenience, the chosen string is prepended with &#39;sourcetype::&#39;. There is no hard-coded default.
     * Sets the sourcetype key initial value. The key is used during parsing/indexing, in particular to set the source type field during indexing. It is also the source type field used at search time.
     * 
     */
    @Import(name="sourcetype")
    private @Nullable Output<String> sourcetype;

    /**
     * @return Sets the sourcetype key/field for events from this input. If unset, Splunk software picks a source type based on various aspects of the data. As a convenience, the chosen string is prepended with &#39;sourcetype::&#39;. There is no hard-coded default.
     * Sets the sourcetype key initial value. The key is used during parsing/indexing, in particular to set the source type field during indexing. It is also the source type field used at search time.
     * 
     */
    public Optional<Output<String>> sourcetype() {
        return Optional.ofNullable(this.sourcetype);
    }

    private InputsScriptArgs() {}

    private InputsScriptArgs(InputsScriptArgs $) {
        this.acl = $.acl;
        this.disabled = $.disabled;
        this.host = $.host;
        this.index = $.index;
        this.interval = $.interval;
        this.name = $.name;
        this.passauth = $.passauth;
        this.renameSource = $.renameSource;
        this.source = $.source;
        this.sourcetype = $.sourcetype;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(InputsScriptArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private InputsScriptArgs $;

        public Builder() {
            $ = new InputsScriptArgs();
        }

        public Builder(InputsScriptArgs defaults) {
            $ = new InputsScriptArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param acl The app/user context that is the namespace for the resource
         * 
         * @return builder
         * 
         */
        public Builder acl(@Nullable Output<InputsScriptAclArgs> acl) {
            $.acl = acl;
            return this;
        }

        /**
         * @param acl The app/user context that is the namespace for the resource
         * 
         * @return builder
         * 
         */
        public Builder acl(InputsScriptAclArgs acl) {
            return acl(Output.of(acl));
        }

        /**
         * @param disabled Specifies whether the input script is disabled.
         * 
         * @return builder
         * 
         */
        public Builder disabled(@Nullable Output<Boolean> disabled) {
            $.disabled = disabled;
            return this;
        }

        /**
         * @param disabled Specifies whether the input script is disabled.
         * 
         * @return builder
         * 
         */
        public Builder disabled(Boolean disabled) {
            return disabled(Output.of(disabled));
        }

        /**
         * @param host Sets the host for events from this input. Defaults to whatever host sent the event.
         * 
         * @return builder
         * 
         */
        public Builder host(@Nullable Output<String> host) {
            $.host = host;
            return this;
        }

        /**
         * @param host Sets the host for events from this input. Defaults to whatever host sent the event.
         * 
         * @return builder
         * 
         */
        public Builder host(String host) {
            return host(Output.of(host));
        }

        /**
         * @param index Sets the index for events from this input. Defaults to the main index.
         * 
         * @return builder
         * 
         */
        public Builder index(@Nullable Output<String> index) {
            $.index = index;
            return this;
        }

        /**
         * @param index Sets the index for events from this input. Defaults to the main index.
         * 
         * @return builder
         * 
         */
        public Builder index(String index) {
            return index(Output.of(index));
        }

        /**
         * @param interval Specify an integer or cron schedule. This parameter specifies how often to execute the specified script, in seconds or a valid cron schedule. If you specify a cron schedule, the script is not executed on start-up.
         * 
         * @return builder
         * 
         */
        public Builder interval(Output<Integer> interval) {
            $.interval = interval;
            return this;
        }

        /**
         * @param interval Specify an integer or cron schedule. This parameter specifies how often to execute the specified script, in seconds or a valid cron schedule. If you specify a cron schedule, the script is not executed on start-up.
         * 
         * @return builder
         * 
         */
        public Builder interval(Integer interval) {
            return interval(Output.of(interval));
        }

        /**
         * @param name Specify the name of the scripted input.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name Specify the name of the scripted input.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param passauth User to run the script as. If you provide a username, Splunk software generates an auth token for that user and passes it to the script.
         * 
         * @return builder
         * 
         */
        public Builder passauth(@Nullable Output<String> passauth) {
            $.passauth = passauth;
            return this;
        }

        /**
         * @param passauth User to run the script as. If you provide a username, Splunk software generates an auth token for that user and passes it to the script.
         * 
         * @return builder
         * 
         */
        public Builder passauth(String passauth) {
            return passauth(Output.of(passauth));
        }

        /**
         * @param renameSource Specify a new name for the source field for the script.
         * 
         * @return builder
         * 
         */
        public Builder renameSource(@Nullable Output<String> renameSource) {
            $.renameSource = renameSource;
            return this;
        }

        /**
         * @param renameSource Specify a new name for the source field for the script.
         * 
         * @return builder
         * 
         */
        public Builder renameSource(String renameSource) {
            return renameSource(Output.of(renameSource));
        }

        /**
         * @param source Sets the source key/field for events from this input. Defaults to the input file path.
         * Sets the source key initial value. The key is used during parsing/indexing, in particular to set the source field during indexing. It is also the source field used at search time. As a convenience, the chosen string is prepended with &#39;source::&#39;.
         * 
         * @return builder
         * 
         */
        public Builder source(@Nullable Output<String> source) {
            $.source = source;
            return this;
        }

        /**
         * @param source Sets the source key/field for events from this input. Defaults to the input file path.
         * Sets the source key initial value. The key is used during parsing/indexing, in particular to set the source field during indexing. It is also the source field used at search time. As a convenience, the chosen string is prepended with &#39;source::&#39;.
         * 
         * @return builder
         * 
         */
        public Builder source(String source) {
            return source(Output.of(source));
        }

        /**
         * @param sourcetype Sets the sourcetype key/field for events from this input. If unset, Splunk software picks a source type based on various aspects of the data. As a convenience, the chosen string is prepended with &#39;sourcetype::&#39;. There is no hard-coded default.
         * Sets the sourcetype key initial value. The key is used during parsing/indexing, in particular to set the source type field during indexing. It is also the source type field used at search time.
         * 
         * @return builder
         * 
         */
        public Builder sourcetype(@Nullable Output<String> sourcetype) {
            $.sourcetype = sourcetype;
            return this;
        }

        /**
         * @param sourcetype Sets the sourcetype key/field for events from this input. If unset, Splunk software picks a source type based on various aspects of the data. As a convenience, the chosen string is prepended with &#39;sourcetype::&#39;. There is no hard-coded default.
         * Sets the sourcetype key initial value. The key is used during parsing/indexing, in particular to set the source type field during indexing. It is also the source type field used at search time.
         * 
         * @return builder
         * 
         */
        public Builder sourcetype(String sourcetype) {
            return sourcetype(Output.of(sourcetype));
        }

        public InputsScriptArgs build() {
            $.interval = Objects.requireNonNull($.interval, "expected parameter 'interval' to be non-null");
            return $;
        }
    }

}