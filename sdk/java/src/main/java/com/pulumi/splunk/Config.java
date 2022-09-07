// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.splunk;

import com.pulumi.core.internal.Codegen;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.Optional;

public final class Config {

    private static final com.pulumi.Config config = com.pulumi.Config.of("splunk");
/**
 * Authentication tokens, also known as JSON Web Tokens (JWT), are a method for authenticating Splunk platform users into
 * the Splunk platform
 * 
 */
    public Optional<String> authToken() {
        return Codegen.stringProp("authToken").config(config).get();
    }
/**
 * insecure skip verification flag
 * 
 */
    public Optional<Boolean> insecureSkipVerify() {
        return Codegen.booleanProp("insecureSkipVerify").config(config).get();
    }
/**
 * Splunk instance password
 * 
 */
    public Optional<String> password() {
        return Codegen.stringProp("password").config(config).get();
    }
/**
 * Timeout when making calls to Splunk server. Defaults to 60 seconds
 * 
 */
    public Optional<Integer> timeout() {
        return Codegen.integerProp("timeout").config(config).get();
    }
/**
 * Splunk instance URL
 * 
 */
    public String url() {
        return Codegen.stringProp("url").config(config).require();
    }
/**
 * Splunk instance admin username
 * 
 */
    public Optional<String> username() {
        return Codegen.stringProp("username").config(config).get();
    }
}
