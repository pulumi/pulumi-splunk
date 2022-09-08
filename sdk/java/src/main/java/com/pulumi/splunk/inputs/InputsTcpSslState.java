// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.splunk.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class InputsTcpSslState extends com.pulumi.resources.ResourceArgs {

    public static final InputsTcpSslState Empty = new InputsTcpSslState();

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
     * Server certificate password, if any.
     * 
     */
    @Import(name="password")
    private @Nullable Output<String> password;

    /**
     * @return Server certificate password, if any.
     * 
     */
    public Optional<Output<String>> password() {
        return Optional.ofNullable(this.password);
    }

    /**
     * Determines whether a client must authenticate.
     * 
     */
    @Import(name="requireClientCert")
    private @Nullable Output<Boolean> requireClientCert;

    /**
     * @return Determines whether a client must authenticate.
     * 
     */
    public Optional<Output<Boolean>> requireClientCert() {
        return Optional.ofNullable(this.requireClientCert);
    }

    /**
     * Certificate authority list (root file)
     * 
     */
    @Import(name="rootCa")
    private @Nullable Output<String> rootCa;

    /**
     * @return Certificate authority list (root file)
     * 
     */
    public Optional<Output<String>> rootCa() {
        return Optional.ofNullable(this.rootCa);
    }

    /**
     * Full path to the server certificate.
     * 
     */
    @Import(name="serverCert")
    private @Nullable Output<String> serverCert;

    /**
     * @return Full path to the server certificate.
     * 
     */
    public Optional<Output<String>> serverCert() {
        return Optional.ofNullable(this.serverCert);
    }

    private InputsTcpSslState() {}

    private InputsTcpSslState(InputsTcpSslState $) {
        this.disabled = $.disabled;
        this.password = $.password;
        this.requireClientCert = $.requireClientCert;
        this.rootCa = $.rootCa;
        this.serverCert = $.serverCert;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(InputsTcpSslState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private InputsTcpSslState $;

        public Builder() {
            $ = new InputsTcpSslState();
        }

        public Builder(InputsTcpSslState defaults) {
            $ = new InputsTcpSslState(Objects.requireNonNull(defaults));
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
         * @param password Server certificate password, if any.
         * 
         * @return builder
         * 
         */
        public Builder password(@Nullable Output<String> password) {
            $.password = password;
            return this;
        }

        /**
         * @param password Server certificate password, if any.
         * 
         * @return builder
         * 
         */
        public Builder password(String password) {
            return password(Output.of(password));
        }

        /**
         * @param requireClientCert Determines whether a client must authenticate.
         * 
         * @return builder
         * 
         */
        public Builder requireClientCert(@Nullable Output<Boolean> requireClientCert) {
            $.requireClientCert = requireClientCert;
            return this;
        }

        /**
         * @param requireClientCert Determines whether a client must authenticate.
         * 
         * @return builder
         * 
         */
        public Builder requireClientCert(Boolean requireClientCert) {
            return requireClientCert(Output.of(requireClientCert));
        }

        /**
         * @param rootCa Certificate authority list (root file)
         * 
         * @return builder
         * 
         */
        public Builder rootCa(@Nullable Output<String> rootCa) {
            $.rootCa = rootCa;
            return this;
        }

        /**
         * @param rootCa Certificate authority list (root file)
         * 
         * @return builder
         * 
         */
        public Builder rootCa(String rootCa) {
            return rootCa(Output.of(rootCa));
        }

        /**
         * @param serverCert Full path to the server certificate.
         * 
         * @return builder
         * 
         */
        public Builder serverCert(@Nullable Output<String> serverCert) {
            $.serverCert = serverCert;
            return this;
        }

        /**
         * @param serverCert Full path to the server certificate.
         * 
         * @return builder
         * 
         */
        public Builder serverCert(String serverCert) {
            return serverCert(Output.of(serverCert));
        }

        public InputsTcpSslState build() {
            return $;
        }
    }

}
