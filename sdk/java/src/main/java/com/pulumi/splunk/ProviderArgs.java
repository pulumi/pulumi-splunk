// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.splunk;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ProviderArgs extends com.pulumi.resources.ResourceArgs {

    public static final ProviderArgs Empty = new ProviderArgs();

    /**
     * Authentication tokens, also known as JSON Web Tokens (JWT), are a method for authenticating Splunk platform users into
     * the Splunk platform
     * 
     */
    @Import(name="authToken")
    private @Nullable Output<String> authToken;

    /**
     * @return Authentication tokens, also known as JSON Web Tokens (JWT), are a method for authenticating Splunk platform users into
     * the Splunk platform
     * 
     */
    public Optional<Output<String>> authToken() {
        return Optional.ofNullable(this.authToken);
    }

    /**
     * insecure skip verification flag
     * 
     */
    @Import(name="insecureSkipVerify", json=true)
    private @Nullable Output<Boolean> insecureSkipVerify;

    /**
     * @return insecure skip verification flag
     * 
     */
    public Optional<Output<Boolean>> insecureSkipVerify() {
        return Optional.ofNullable(this.insecureSkipVerify);
    }

    /**
     * Splunk instance password
     * 
     */
    @Import(name="password")
    private @Nullable Output<String> password;

    /**
     * @return Splunk instance password
     * 
     */
    public Optional<Output<String>> password() {
        return Optional.ofNullable(this.password);
    }

    /**
     * Timeout when making calls to Splunk server. Defaults to 60 seconds
     * 
     */
    @Import(name="timeout", json=true)
    private @Nullable Output<Integer> timeout;

    /**
     * @return Timeout when making calls to Splunk server. Defaults to 60 seconds
     * 
     */
    public Optional<Output<Integer>> timeout() {
        return Optional.ofNullable(this.timeout);
    }

    /**
     * Splunk instance URL
     * 
     */
    @Import(name="url", required=true)
    private Output<String> url;

    /**
     * @return Splunk instance URL
     * 
     */
    public Output<String> url() {
        return this.url;
    }

    /**
     * Splunk instance admin username
     * 
     */
    @Import(name="username")
    private @Nullable Output<String> username;

    /**
     * @return Splunk instance admin username
     * 
     */
    public Optional<Output<String>> username() {
        return Optional.ofNullable(this.username);
    }

    private ProviderArgs() {}

    private ProviderArgs(ProviderArgs $) {
        this.authToken = $.authToken;
        this.insecureSkipVerify = $.insecureSkipVerify;
        this.password = $.password;
        this.timeout = $.timeout;
        this.url = $.url;
        this.username = $.username;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ProviderArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ProviderArgs $;

        public Builder() {
            $ = new ProviderArgs();
        }

        public Builder(ProviderArgs defaults) {
            $ = new ProviderArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param authToken Authentication tokens, also known as JSON Web Tokens (JWT), are a method for authenticating Splunk platform users into
         * the Splunk platform
         * 
         * @return builder
         * 
         */
        public Builder authToken(@Nullable Output<String> authToken) {
            $.authToken = authToken;
            return this;
        }

        /**
         * @param authToken Authentication tokens, also known as JSON Web Tokens (JWT), are a method for authenticating Splunk platform users into
         * the Splunk platform
         * 
         * @return builder
         * 
         */
        public Builder authToken(String authToken) {
            return authToken(Output.of(authToken));
        }

        /**
         * @param insecureSkipVerify insecure skip verification flag
         * 
         * @return builder
         * 
         */
        public Builder insecureSkipVerify(@Nullable Output<Boolean> insecureSkipVerify) {
            $.insecureSkipVerify = insecureSkipVerify;
            return this;
        }

        /**
         * @param insecureSkipVerify insecure skip verification flag
         * 
         * @return builder
         * 
         */
        public Builder insecureSkipVerify(Boolean insecureSkipVerify) {
            return insecureSkipVerify(Output.of(insecureSkipVerify));
        }

        /**
         * @param password Splunk instance password
         * 
         * @return builder
         * 
         */
        public Builder password(@Nullable Output<String> password) {
            $.password = password;
            return this;
        }

        /**
         * @param password Splunk instance password
         * 
         * @return builder
         * 
         */
        public Builder password(String password) {
            return password(Output.of(password));
        }

        /**
         * @param timeout Timeout when making calls to Splunk server. Defaults to 60 seconds
         * 
         * @return builder
         * 
         */
        public Builder timeout(@Nullable Output<Integer> timeout) {
            $.timeout = timeout;
            return this;
        }

        /**
         * @param timeout Timeout when making calls to Splunk server. Defaults to 60 seconds
         * 
         * @return builder
         * 
         */
        public Builder timeout(Integer timeout) {
            return timeout(Output.of(timeout));
        }

        /**
         * @param url Splunk instance URL
         * 
         * @return builder
         * 
         */
        public Builder url(Output<String> url) {
            $.url = url;
            return this;
        }

        /**
         * @param url Splunk instance URL
         * 
         * @return builder
         * 
         */
        public Builder url(String url) {
            return url(Output.of(url));
        }

        /**
         * @param username Splunk instance admin username
         * 
         * @return builder
         * 
         */
        public Builder username(@Nullable Output<String> username) {
            $.username = username;
            return this;
        }

        /**
         * @param username Splunk instance admin username
         * 
         * @return builder
         * 
         */
        public Builder username(String username) {
            return username(Output.of(username));
        }

        public ProviderArgs build() {
            if ($.url == null) {
                throw new MissingRequiredPropertyException("ProviderArgs", "url");
            }
            return $;
        }
    }

}
