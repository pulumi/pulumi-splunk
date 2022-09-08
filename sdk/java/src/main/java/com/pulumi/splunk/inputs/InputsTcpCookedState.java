// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.splunk.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.splunk.inputs.InputsTcpCookedAclArgs;
import java.lang.Boolean;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class InputsTcpCookedState extends com.pulumi.resources.ResourceArgs {

    public static final InputsTcpCookedState Empty = new InputsTcpCookedState();

    /**
     * The app/user context that is the namespace for the resource
     * 
     */
    @Import(name="acl")
    private @Nullable Output<InputsTcpCookedAclArgs> acl;

    /**
     * @return The app/user context that is the namespace for the resource
     * 
     */
    public Optional<Output<InputsTcpCookedAclArgs>> acl() {
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
     * The port number of this input.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return The port number of this input.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * Restrict incoming connections on this port to the host specified here.
     * 
     */
    @Import(name="restrictToHost")
    private @Nullable Output<String> restrictToHost;

    /**
     * @return Restrict incoming connections on this port to the host specified here.
     * 
     */
    public Optional<Output<String>> restrictToHost() {
        return Optional.ofNullable(this.restrictToHost);
    }

    private InputsTcpCookedState() {}

    private InputsTcpCookedState(InputsTcpCookedState $) {
        this.acl = $.acl;
        this.connectionHost = $.connectionHost;
        this.disabled = $.disabled;
        this.host = $.host;
        this.name = $.name;
        this.restrictToHost = $.restrictToHost;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(InputsTcpCookedState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private InputsTcpCookedState $;

        public Builder() {
            $ = new InputsTcpCookedState();
        }

        public Builder(InputsTcpCookedState defaults) {
            $ = new InputsTcpCookedState(Objects.requireNonNull(defaults));
        }

        /**
         * @param acl The app/user context that is the namespace for the resource
         * 
         * @return builder
         * 
         */
        public Builder acl(@Nullable Output<InputsTcpCookedAclArgs> acl) {
            $.acl = acl;
            return this;
        }

        /**
         * @param acl The app/user context that is the namespace for the resource
         * 
         * @return builder
         * 
         */
        public Builder acl(InputsTcpCookedAclArgs acl) {
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
         * @param name The port number of this input.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name The port number of this input.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param restrictToHost Restrict incoming connections on this port to the host specified here.
         * 
         * @return builder
         * 
         */
        public Builder restrictToHost(@Nullable Output<String> restrictToHost) {
            $.restrictToHost = restrictToHost;
            return this;
        }

        /**
         * @param restrictToHost Restrict incoming connections on this port to the host specified here.
         * 
         * @return builder
         * 
         */
        public Builder restrictToHost(String restrictToHost) {
            return restrictToHost(Output.of(restrictToHost));
        }

        public InputsTcpCookedState build() {
            return $;
        }
    }

}
