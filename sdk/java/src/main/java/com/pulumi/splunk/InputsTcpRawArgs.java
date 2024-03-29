// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.splunk;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.splunk.inputs.InputsTcpRawAclArgs;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class InputsTcpRawArgs extends com.pulumi.resources.ResourceArgs {

    public static final InputsTcpRawArgs Empty = new InputsTcpRawArgs();

    /**
     * The app/user context that is the namespace for the resource
     * 
     */
    @Import(name="acl")
    private @Nullable Output<InputsTcpRawAclArgs> acl;

    /**
     * @return The app/user context that is the namespace for the resource
     * 
     */
    public Optional<Output<InputsTcpRawAclArgs>> acl() {
        return Optional.ofNullable(this.acl);
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
    @Import(name="connectionHost")
    private @Nullable Output<String> connectionHost;

    /**
     * @return Valid values: (ip | dns | none)
     * Set the host for the remote server that is sending data.
     * ip sets the host to the IP address of the remote server sending data.
     * dns sets the host to the reverse DNS entry for the IP address of the remote server sending data.
     * none leaves the host as specified in inputs.conf, which is typically the Splunk system hostname.
     * Default value is dns.
     * 
     */
    public Optional<Output<String>> connectionHost() {
        return Optional.ofNullable(this.connectionHost);
    }

    /**
     * Indicates if input is disabled.
     * 
     */
    @Import(name="disabled")
    private @Nullable Output<Boolean> disabled;

    /**
     * @return Indicates if input is disabled.
     * 
     */
    public Optional<Output<Boolean>> disabled() {
        return Optional.ofNullable(this.disabled);
    }

    /**
     * Host from which the indexer gets data.
     * 
     */
    @Import(name="host")
    private @Nullable Output<String> host;

    /**
     * @return Host from which the indexer gets data.
     * 
     */
    public Optional<Output<String>> host() {
        return Optional.ofNullable(this.host);
    }

    /**
     * Index to store generated events. Defaults to default.
     * 
     */
    @Import(name="index")
    private @Nullable Output<String> index;

    /**
     * @return Index to store generated events. Defaults to default.
     * 
     */
    public Optional<Output<String>> index() {
        return Optional.ofNullable(this.index);
    }

    /**
     * The input port which receives raw data.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return The input port which receives raw data.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * Valid values: (parsingQueue | indexQueue)
     * Specifies where the input processor should deposit the events it reads. Defaults to parsingQueue.
     * Set queue to parsingQueue to apply props.conf and other parsing rules to your data. For more information about props.conf and rules for timestamping and linebreaking, refer to props.conf and the online documentation at &#34;Monitor files and directories with inputs.conf&#34;
     * Set queue to indexQueue to send your data directly into the index.
     * 
     */
    @Import(name="queue")
    private @Nullable Output<String> queue;

    /**
     * @return Valid values: (parsingQueue | indexQueue)
     * Specifies where the input processor should deposit the events it reads. Defaults to parsingQueue.
     * Set queue to parsingQueue to apply props.conf and other parsing rules to your data. For more information about props.conf and rules for timestamping and linebreaking, refer to props.conf and the online documentation at &#34;Monitor files and directories with inputs.conf&#34;
     * Set queue to indexQueue to send your data directly into the index.
     * 
     */
    public Optional<Output<String>> queue() {
        return Optional.ofNullable(this.queue);
    }

    /**
     * Specifies in seconds the timeout value for adding a Done-key. Default value is 10 seconds.
     * If a connection over the port specified by name remains idle after receiving data for specified number of seconds, it adds a Done-key. This implies the last event is completely received.
     * 
     */
    @Import(name="rawTcpDoneTimeout")
    private @Nullable Output<Integer> rawTcpDoneTimeout;

    /**
     * @return Specifies in seconds the timeout value for adding a Done-key. Default value is 10 seconds.
     * If a connection over the port specified by name remains idle after receiving data for specified number of seconds, it adds a Done-key. This implies the last event is completely received.
     * 
     */
    public Optional<Output<Integer>> rawTcpDoneTimeout() {
        return Optional.ofNullable(this.rawTcpDoneTimeout);
    }

    /**
     * Allows for restricting this input to only accept data from the host specified here.
     * 
     */
    @Import(name="restrictToHost")
    private @Nullable Output<String> restrictToHost;

    /**
     * @return Allows for restricting this input to only accept data from the host specified here.
     * 
     */
    public Optional<Output<String>> restrictToHost() {
        return Optional.ofNullable(this.restrictToHost);
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
     * Set the source type for events from this input.
     * &#34;sourcetype=&#34; is automatically prepended to &lt;string&gt;.
     * Defaults to audittrail (if signedaudit=true) or fschange (if signedaudit=false).
     * 
     */
    @Import(name="sourcetype")
    private @Nullable Output<String> sourcetype;

    /**
     * @return Set the source type for events from this input.
     * &#34;sourcetype=&#34; is automatically prepended to &lt;string&gt;.
     * Defaults to audittrail (if signedaudit=true) or fschange (if signedaudit=false).
     * 
     */
    public Optional<Output<String>> sourcetype() {
        return Optional.ofNullable(this.sourcetype);
    }

    private InputsTcpRawArgs() {}

    private InputsTcpRawArgs(InputsTcpRawArgs $) {
        this.acl = $.acl;
        this.connectionHost = $.connectionHost;
        this.disabled = $.disabled;
        this.host = $.host;
        this.index = $.index;
        this.name = $.name;
        this.queue = $.queue;
        this.rawTcpDoneTimeout = $.rawTcpDoneTimeout;
        this.restrictToHost = $.restrictToHost;
        this.source = $.source;
        this.sourcetype = $.sourcetype;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(InputsTcpRawArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private InputsTcpRawArgs $;

        public Builder() {
            $ = new InputsTcpRawArgs();
        }

        public Builder(InputsTcpRawArgs defaults) {
            $ = new InputsTcpRawArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param acl The app/user context that is the namespace for the resource
         * 
         * @return builder
         * 
         */
        public Builder acl(@Nullable Output<InputsTcpRawAclArgs> acl) {
            $.acl = acl;
            return this;
        }

        /**
         * @param acl The app/user context that is the namespace for the resource
         * 
         * @return builder
         * 
         */
        public Builder acl(InputsTcpRawAclArgs acl) {
            return acl(Output.of(acl));
        }

        /**
         * @param connectionHost Valid values: (ip | dns | none)
         * Set the host for the remote server that is sending data.
         * ip sets the host to the IP address of the remote server sending data.
         * dns sets the host to the reverse DNS entry for the IP address of the remote server sending data.
         * none leaves the host as specified in inputs.conf, which is typically the Splunk system hostname.
         * Default value is dns.
         * 
         * @return builder
         * 
         */
        public Builder connectionHost(@Nullable Output<String> connectionHost) {
            $.connectionHost = connectionHost;
            return this;
        }

        /**
         * @param connectionHost Valid values: (ip | dns | none)
         * Set the host for the remote server that is sending data.
         * ip sets the host to the IP address of the remote server sending data.
         * dns sets the host to the reverse DNS entry for the IP address of the remote server sending data.
         * none leaves the host as specified in inputs.conf, which is typically the Splunk system hostname.
         * Default value is dns.
         * 
         * @return builder
         * 
         */
        public Builder connectionHost(String connectionHost) {
            return connectionHost(Output.of(connectionHost));
        }

        /**
         * @param disabled Indicates if input is disabled.
         * 
         * @return builder
         * 
         */
        public Builder disabled(@Nullable Output<Boolean> disabled) {
            $.disabled = disabled;
            return this;
        }

        /**
         * @param disabled Indicates if input is disabled.
         * 
         * @return builder
         * 
         */
        public Builder disabled(Boolean disabled) {
            return disabled(Output.of(disabled));
        }

        /**
         * @param host Host from which the indexer gets data.
         * 
         * @return builder
         * 
         */
        public Builder host(@Nullable Output<String> host) {
            $.host = host;
            return this;
        }

        /**
         * @param host Host from which the indexer gets data.
         * 
         * @return builder
         * 
         */
        public Builder host(String host) {
            return host(Output.of(host));
        }

        /**
         * @param index Index to store generated events. Defaults to default.
         * 
         * @return builder
         * 
         */
        public Builder index(@Nullable Output<String> index) {
            $.index = index;
            return this;
        }

        /**
         * @param index Index to store generated events. Defaults to default.
         * 
         * @return builder
         * 
         */
        public Builder index(String index) {
            return index(Output.of(index));
        }

        /**
         * @param name The input port which receives raw data.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name The input port which receives raw data.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param queue Valid values: (parsingQueue | indexQueue)
         * Specifies where the input processor should deposit the events it reads. Defaults to parsingQueue.
         * Set queue to parsingQueue to apply props.conf and other parsing rules to your data. For more information about props.conf and rules for timestamping and linebreaking, refer to props.conf and the online documentation at &#34;Monitor files and directories with inputs.conf&#34;
         * Set queue to indexQueue to send your data directly into the index.
         * 
         * @return builder
         * 
         */
        public Builder queue(@Nullable Output<String> queue) {
            $.queue = queue;
            return this;
        }

        /**
         * @param queue Valid values: (parsingQueue | indexQueue)
         * Specifies where the input processor should deposit the events it reads. Defaults to parsingQueue.
         * Set queue to parsingQueue to apply props.conf and other parsing rules to your data. For more information about props.conf and rules for timestamping and linebreaking, refer to props.conf and the online documentation at &#34;Monitor files and directories with inputs.conf&#34;
         * Set queue to indexQueue to send your data directly into the index.
         * 
         * @return builder
         * 
         */
        public Builder queue(String queue) {
            return queue(Output.of(queue));
        }

        /**
         * @param rawTcpDoneTimeout Specifies in seconds the timeout value for adding a Done-key. Default value is 10 seconds.
         * If a connection over the port specified by name remains idle after receiving data for specified number of seconds, it adds a Done-key. This implies the last event is completely received.
         * 
         * @return builder
         * 
         */
        public Builder rawTcpDoneTimeout(@Nullable Output<Integer> rawTcpDoneTimeout) {
            $.rawTcpDoneTimeout = rawTcpDoneTimeout;
            return this;
        }

        /**
         * @param rawTcpDoneTimeout Specifies in seconds the timeout value for adding a Done-key. Default value is 10 seconds.
         * If a connection over the port specified by name remains idle after receiving data for specified number of seconds, it adds a Done-key. This implies the last event is completely received.
         * 
         * @return builder
         * 
         */
        public Builder rawTcpDoneTimeout(Integer rawTcpDoneTimeout) {
            return rawTcpDoneTimeout(Output.of(rawTcpDoneTimeout));
        }

        /**
         * @param restrictToHost Allows for restricting this input to only accept data from the host specified here.
         * 
         * @return builder
         * 
         */
        public Builder restrictToHost(@Nullable Output<String> restrictToHost) {
            $.restrictToHost = restrictToHost;
            return this;
        }

        /**
         * @param restrictToHost Allows for restricting this input to only accept data from the host specified here.
         * 
         * @return builder
         * 
         */
        public Builder restrictToHost(String restrictToHost) {
            return restrictToHost(Output.of(restrictToHost));
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
         * @param sourcetype Set the source type for events from this input.
         * &#34;sourcetype=&#34; is automatically prepended to &lt;string&gt;.
         * Defaults to audittrail (if signedaudit=true) or fschange (if signedaudit=false).
         * 
         * @return builder
         * 
         */
        public Builder sourcetype(@Nullable Output<String> sourcetype) {
            $.sourcetype = sourcetype;
            return this;
        }

        /**
         * @param sourcetype Set the source type for events from this input.
         * &#34;sourcetype=&#34; is automatically prepended to &lt;string&gt;.
         * Defaults to audittrail (if signedaudit=true) or fschange (if signedaudit=false).
         * 
         * @return builder
         * 
         */
        public Builder sourcetype(String sourcetype) {
            return sourcetype(Output.of(sourcetype));
        }

        public InputsTcpRawArgs build() {
            return $;
        }
    }

}
