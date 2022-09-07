// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.splunk;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.splunk.inputs.OutputsTcpServerAclArgs;
import java.lang.Boolean;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class OutputsTcpServerArgs extends com.pulumi.resources.ResourceArgs {

    public static final OutputsTcpServerArgs Empty = new OutputsTcpServerArgs();

    /**
     * The app/user context that is the namespace for the resource
     * 
     */
    @Import(name="acl")
    private @Nullable Output<OutputsTcpServerAclArgs> acl;

    /**
     * @return The app/user context that is the namespace for the resource
     * 
     */
    public Optional<Output<OutputsTcpServerAclArgs>> acl() {
        return Optional.ofNullable(this.acl);
    }

    /**
     * If true, disables the group.
     * 
     */
    @Import(name="disabled")
    private @Nullable Output<Boolean> disabled;

    /**
     * @return If true, disables the group.
     * 
     */
    public Optional<Output<Boolean>> disabled() {
        return Optional.ofNullable(this.disabled);
    }

    /**
     * Valid values: (clone | balance | autobalance)
     * The data distribution method used when two or more servers exist in the same forwarder group.
     * 
     */
    @Import(name="method")
    private @Nullable Output<String> method;

    /**
     * @return Valid values: (clone | balance | autobalance)
     * The data distribution method used when two or more servers exist in the same forwarder group.
     * 
     */
    public Optional<Output<String>> method() {
        return Optional.ofNullable(this.method);
    }

    /**
     * &lt;host&gt;:&lt;port&gt; of the Splunk receiver. &lt;host&gt; can be either an ip address or server name. &lt;port&gt; is the that port that the Splunk receiver is listening on.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return &lt;host&gt;:&lt;port&gt; of the Splunk receiver. &lt;host&gt; can be either an ip address or server name. &lt;port&gt; is the that port that the Splunk receiver is listening on.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * The alternate name to match in the remote server&#39;s SSL certificate.
     * 
     */
    @Import(name="sslAltNameToCheck")
    private @Nullable Output<String> sslAltNameToCheck;

    /**
     * @return The alternate name to match in the remote server&#39;s SSL certificate.
     * 
     */
    public Optional<Output<String>> sslAltNameToCheck() {
        return Optional.ofNullable(this.sslAltNameToCheck);
    }

    /**
     * Path to the client certificate. If specified, connection uses SSL.
     * 
     */
    @Import(name="sslCertPath")
    private @Nullable Output<String> sslCertPath;

    /**
     * @return Path to the client certificate. If specified, connection uses SSL.
     * 
     */
    public Optional<Output<String>> sslCertPath() {
        return Optional.ofNullable(this.sslCertPath);
    }

    /**
     * SSL Cipher in the form ALL:!aNULL:!eNULL:!LOW:!EXP:RC4+RSA:+HIGH:+MEDIUM
     * 
     */
    @Import(name="sslCipher")
    private @Nullable Output<String> sslCipher;

    /**
     * @return SSL Cipher in the form ALL:!aNULL:!eNULL:!LOW:!EXP:RC4+RSA:+HIGH:+MEDIUM
     * 
     */
    public Optional<Output<String>> sslCipher() {
        return Optional.ofNullable(this.sslCipher);
    }

    /**
     * Check the common name of the server&#39;s certificate against this name.
     * If there is no match, assume that Splunk Enterprise is not authenticated against this server. You must specify this setting if sslVerifyServerCert is true.
     * 
     */
    @Import(name="sslCommonNameToCheck")
    private @Nullable Output<String> sslCommonNameToCheck;

    /**
     * @return Check the common name of the server&#39;s certificate against this name.
     * If there is no match, assume that Splunk Enterprise is not authenticated against this server. You must specify this setting if sslVerifyServerCert is true.
     * 
     */
    public Optional<Output<String>> sslCommonNameToCheck() {
        return Optional.ofNullable(this.sslCommonNameToCheck);
    }

    /**
     * The password associated with the CAcert.
     * The default Splunk Enterprise CAcert uses the password &#34;password.&#34;
     * 
     */
    @Import(name="sslPassword")
    private @Nullable Output<String> sslPassword;

    /**
     * @return The password associated with the CAcert.
     * The default Splunk Enterprise CAcert uses the password &#34;password.&#34;
     * 
     */
    public Optional<Output<String>> sslPassword() {
        return Optional.ofNullable(this.sslPassword);
    }

    /**
     * The path to the root certificate authority file.
     * 
     */
    @Import(name="sslRootCaPath")
    private @Nullable Output<String> sslRootCaPath;

    /**
     * @return The path to the root certificate authority file.
     * 
     */
    public Optional<Output<String>> sslRootCaPath() {
        return Optional.ofNullable(this.sslRootCaPath);
    }

    /**
     * If true, make sure that the server you are connecting to is a valid one (authenticated). Both the common name and the alternate name of the server are then checked for a match.
     * 
     */
    @Import(name="sslVerifyServerCert")
    private @Nullable Output<Boolean> sslVerifyServerCert;

    /**
     * @return If true, make sure that the server you are connecting to is a valid one (authenticated). Both the common name and the alternate name of the server are then checked for a match.
     * 
     */
    public Optional<Output<Boolean>> sslVerifyServerCert() {
        return Optional.ofNullable(this.sslVerifyServerCert);
    }

    private OutputsTcpServerArgs() {}

    private OutputsTcpServerArgs(OutputsTcpServerArgs $) {
        this.acl = $.acl;
        this.disabled = $.disabled;
        this.method = $.method;
        this.name = $.name;
        this.sslAltNameToCheck = $.sslAltNameToCheck;
        this.sslCertPath = $.sslCertPath;
        this.sslCipher = $.sslCipher;
        this.sslCommonNameToCheck = $.sslCommonNameToCheck;
        this.sslPassword = $.sslPassword;
        this.sslRootCaPath = $.sslRootCaPath;
        this.sslVerifyServerCert = $.sslVerifyServerCert;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(OutputsTcpServerArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private OutputsTcpServerArgs $;

        public Builder() {
            $ = new OutputsTcpServerArgs();
        }

        public Builder(OutputsTcpServerArgs defaults) {
            $ = new OutputsTcpServerArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param acl The app/user context that is the namespace for the resource
         * 
         * @return builder
         * 
         */
        public Builder acl(@Nullable Output<OutputsTcpServerAclArgs> acl) {
            $.acl = acl;
            return this;
        }

        /**
         * @param acl The app/user context that is the namespace for the resource
         * 
         * @return builder
         * 
         */
        public Builder acl(OutputsTcpServerAclArgs acl) {
            return acl(Output.of(acl));
        }

        /**
         * @param disabled If true, disables the group.
         * 
         * @return builder
         * 
         */
        public Builder disabled(@Nullable Output<Boolean> disabled) {
            $.disabled = disabled;
            return this;
        }

        /**
         * @param disabled If true, disables the group.
         * 
         * @return builder
         * 
         */
        public Builder disabled(Boolean disabled) {
            return disabled(Output.of(disabled));
        }

        /**
         * @param method Valid values: (clone | balance | autobalance)
         * The data distribution method used when two or more servers exist in the same forwarder group.
         * 
         * @return builder
         * 
         */
        public Builder method(@Nullable Output<String> method) {
            $.method = method;
            return this;
        }

        /**
         * @param method Valid values: (clone | balance | autobalance)
         * The data distribution method used when two or more servers exist in the same forwarder group.
         * 
         * @return builder
         * 
         */
        public Builder method(String method) {
            return method(Output.of(method));
        }

        /**
         * @param name &lt;host&gt;:&lt;port&gt; of the Splunk receiver. &lt;host&gt; can be either an ip address or server name. &lt;port&gt; is the that port that the Splunk receiver is listening on.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name &lt;host&gt;:&lt;port&gt; of the Splunk receiver. &lt;host&gt; can be either an ip address or server name. &lt;port&gt; is the that port that the Splunk receiver is listening on.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param sslAltNameToCheck The alternate name to match in the remote server&#39;s SSL certificate.
         * 
         * @return builder
         * 
         */
        public Builder sslAltNameToCheck(@Nullable Output<String> sslAltNameToCheck) {
            $.sslAltNameToCheck = sslAltNameToCheck;
            return this;
        }

        /**
         * @param sslAltNameToCheck The alternate name to match in the remote server&#39;s SSL certificate.
         * 
         * @return builder
         * 
         */
        public Builder sslAltNameToCheck(String sslAltNameToCheck) {
            return sslAltNameToCheck(Output.of(sslAltNameToCheck));
        }

        /**
         * @param sslCertPath Path to the client certificate. If specified, connection uses SSL.
         * 
         * @return builder
         * 
         */
        public Builder sslCertPath(@Nullable Output<String> sslCertPath) {
            $.sslCertPath = sslCertPath;
            return this;
        }

        /**
         * @param sslCertPath Path to the client certificate. If specified, connection uses SSL.
         * 
         * @return builder
         * 
         */
        public Builder sslCertPath(String sslCertPath) {
            return sslCertPath(Output.of(sslCertPath));
        }

        /**
         * @param sslCipher SSL Cipher in the form ALL:!aNULL:!eNULL:!LOW:!EXP:RC4+RSA:+HIGH:+MEDIUM
         * 
         * @return builder
         * 
         */
        public Builder sslCipher(@Nullable Output<String> sslCipher) {
            $.sslCipher = sslCipher;
            return this;
        }

        /**
         * @param sslCipher SSL Cipher in the form ALL:!aNULL:!eNULL:!LOW:!EXP:RC4+RSA:+HIGH:+MEDIUM
         * 
         * @return builder
         * 
         */
        public Builder sslCipher(String sslCipher) {
            return sslCipher(Output.of(sslCipher));
        }

        /**
         * @param sslCommonNameToCheck Check the common name of the server&#39;s certificate against this name.
         * If there is no match, assume that Splunk Enterprise is not authenticated against this server. You must specify this setting if sslVerifyServerCert is true.
         * 
         * @return builder
         * 
         */
        public Builder sslCommonNameToCheck(@Nullable Output<String> sslCommonNameToCheck) {
            $.sslCommonNameToCheck = sslCommonNameToCheck;
            return this;
        }

        /**
         * @param sslCommonNameToCheck Check the common name of the server&#39;s certificate against this name.
         * If there is no match, assume that Splunk Enterprise is not authenticated against this server. You must specify this setting if sslVerifyServerCert is true.
         * 
         * @return builder
         * 
         */
        public Builder sslCommonNameToCheck(String sslCommonNameToCheck) {
            return sslCommonNameToCheck(Output.of(sslCommonNameToCheck));
        }

        /**
         * @param sslPassword The password associated with the CAcert.
         * The default Splunk Enterprise CAcert uses the password &#34;password.&#34;
         * 
         * @return builder
         * 
         */
        public Builder sslPassword(@Nullable Output<String> sslPassword) {
            $.sslPassword = sslPassword;
            return this;
        }

        /**
         * @param sslPassword The password associated with the CAcert.
         * The default Splunk Enterprise CAcert uses the password &#34;password.&#34;
         * 
         * @return builder
         * 
         */
        public Builder sslPassword(String sslPassword) {
            return sslPassword(Output.of(sslPassword));
        }

        /**
         * @param sslRootCaPath The path to the root certificate authority file.
         * 
         * @return builder
         * 
         */
        public Builder sslRootCaPath(@Nullable Output<String> sslRootCaPath) {
            $.sslRootCaPath = sslRootCaPath;
            return this;
        }

        /**
         * @param sslRootCaPath The path to the root certificate authority file.
         * 
         * @return builder
         * 
         */
        public Builder sslRootCaPath(String sslRootCaPath) {
            return sslRootCaPath(Output.of(sslRootCaPath));
        }

        /**
         * @param sslVerifyServerCert If true, make sure that the server you are connecting to is a valid one (authenticated). Both the common name and the alternate name of the server are then checked for a match.
         * 
         * @return builder
         * 
         */
        public Builder sslVerifyServerCert(@Nullable Output<Boolean> sslVerifyServerCert) {
            $.sslVerifyServerCert = sslVerifyServerCert;
            return this;
        }

        /**
         * @param sslVerifyServerCert If true, make sure that the server you are connecting to is a valid one (authenticated). Both the common name and the alternate name of the server are then checked for a match.
         * 
         * @return builder
         * 
         */
        public Builder sslVerifyServerCert(Boolean sslVerifyServerCert) {
            return sslVerifyServerCert(Output.of(sslVerifyServerCert));
        }

        public OutputsTcpServerArgs build() {
            return $;
        }
    }

}
