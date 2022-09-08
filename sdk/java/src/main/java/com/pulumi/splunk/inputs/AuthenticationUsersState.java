// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.splunk.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class AuthenticationUsersState extends com.pulumi.resources.ResourceArgs {

    public static final AuthenticationUsersState Empty = new AuthenticationUsersState();

    /**
     * User default app. Overrides the default app inherited from the user roles.
     * 
     */
    @Import(name="defaultApp")
    private @Nullable Output<String> defaultApp;

    /**
     * @return User default app. Overrides the default app inherited from the user roles.
     * 
     */
    public Optional<Output<String>> defaultApp() {
        return Optional.ofNullable(this.defaultApp);
    }

    /**
     * User email address.
     * 
     */
    @Import(name="email")
    private @Nullable Output<String> email;

    /**
     * @return User email address.
     * 
     */
    public Optional<Output<String>> email() {
        return Optional.ofNullable(this.email);
    }

    /**
     * Force user to change password indication
     * 
     */
    @Import(name="forceChangePass")
    private @Nullable Output<Boolean> forceChangePass;

    /**
     * @return Force user to change password indication
     * 
     */
    public Optional<Output<Boolean>> forceChangePass() {
        return Optional.ofNullable(this.forceChangePass);
    }

    /**
     * Unique user login name.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return Unique user login name.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * User login password.
     * 
     */
    @Import(name="password")
    private @Nullable Output<String> password;

    /**
     * @return User login password.
     * 
     */
    public Optional<Output<String>> password() {
        return Optional.ofNullable(this.password);
    }

    /**
     * Full user name.
     * 
     */
    @Import(name="realname")
    private @Nullable Output<String> realname;

    /**
     * @return Full user name.
     * 
     */
    public Optional<Output<String>> realname() {
        return Optional.ofNullable(this.realname);
    }

    /**
     * Restart background search job that has not completed when Splunk restarts indication.
     * 
     */
    @Import(name="restartBackgroundJobs")
    private @Nullable Output<Boolean> restartBackgroundJobs;

    /**
     * @return Restart background search job that has not completed when Splunk restarts indication.
     * 
     */
    public Optional<Output<Boolean>> restartBackgroundJobs() {
        return Optional.ofNullable(this.restartBackgroundJobs);
    }

    /**
     * Role to assign to this user. At least one existing role is required.
     * 
     */
    @Import(name="roles")
    private @Nullable Output<List<String>> roles;

    /**
     * @return Role to assign to this user. At least one existing role is required.
     * 
     */
    public Optional<Output<List<String>>> roles() {
        return Optional.ofNullable(this.roles);
    }

    /**
     * User timezone.
     * 
     */
    @Import(name="tz")
    private @Nullable Output<String> tz;

    /**
     * @return User timezone.
     * 
     */
    public Optional<Output<String>> tz() {
        return Optional.ofNullable(this.tz);
    }

    private AuthenticationUsersState() {}

    private AuthenticationUsersState(AuthenticationUsersState $) {
        this.defaultApp = $.defaultApp;
        this.email = $.email;
        this.forceChangePass = $.forceChangePass;
        this.name = $.name;
        this.password = $.password;
        this.realname = $.realname;
        this.restartBackgroundJobs = $.restartBackgroundJobs;
        this.roles = $.roles;
        this.tz = $.tz;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(AuthenticationUsersState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private AuthenticationUsersState $;

        public Builder() {
            $ = new AuthenticationUsersState();
        }

        public Builder(AuthenticationUsersState defaults) {
            $ = new AuthenticationUsersState(Objects.requireNonNull(defaults));
        }

        /**
         * @param defaultApp User default app. Overrides the default app inherited from the user roles.
         * 
         * @return builder
         * 
         */
        public Builder defaultApp(@Nullable Output<String> defaultApp) {
            $.defaultApp = defaultApp;
            return this;
        }

        /**
         * @param defaultApp User default app. Overrides the default app inherited from the user roles.
         * 
         * @return builder
         * 
         */
        public Builder defaultApp(String defaultApp) {
            return defaultApp(Output.of(defaultApp));
        }

        /**
         * @param email User email address.
         * 
         * @return builder
         * 
         */
        public Builder email(@Nullable Output<String> email) {
            $.email = email;
            return this;
        }

        /**
         * @param email User email address.
         * 
         * @return builder
         * 
         */
        public Builder email(String email) {
            return email(Output.of(email));
        }

        /**
         * @param forceChangePass Force user to change password indication
         * 
         * @return builder
         * 
         */
        public Builder forceChangePass(@Nullable Output<Boolean> forceChangePass) {
            $.forceChangePass = forceChangePass;
            return this;
        }

        /**
         * @param forceChangePass Force user to change password indication
         * 
         * @return builder
         * 
         */
        public Builder forceChangePass(Boolean forceChangePass) {
            return forceChangePass(Output.of(forceChangePass));
        }

        /**
         * @param name Unique user login name.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name Unique user login name.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param password User login password.
         * 
         * @return builder
         * 
         */
        public Builder password(@Nullable Output<String> password) {
            $.password = password;
            return this;
        }

        /**
         * @param password User login password.
         * 
         * @return builder
         * 
         */
        public Builder password(String password) {
            return password(Output.of(password));
        }

        /**
         * @param realname Full user name.
         * 
         * @return builder
         * 
         */
        public Builder realname(@Nullable Output<String> realname) {
            $.realname = realname;
            return this;
        }

        /**
         * @param realname Full user name.
         * 
         * @return builder
         * 
         */
        public Builder realname(String realname) {
            return realname(Output.of(realname));
        }

        /**
         * @param restartBackgroundJobs Restart background search job that has not completed when Splunk restarts indication.
         * 
         * @return builder
         * 
         */
        public Builder restartBackgroundJobs(@Nullable Output<Boolean> restartBackgroundJobs) {
            $.restartBackgroundJobs = restartBackgroundJobs;
            return this;
        }

        /**
         * @param restartBackgroundJobs Restart background search job that has not completed when Splunk restarts indication.
         * 
         * @return builder
         * 
         */
        public Builder restartBackgroundJobs(Boolean restartBackgroundJobs) {
            return restartBackgroundJobs(Output.of(restartBackgroundJobs));
        }

        /**
         * @param roles Role to assign to this user. At least one existing role is required.
         * 
         * @return builder
         * 
         */
        public Builder roles(@Nullable Output<List<String>> roles) {
            $.roles = roles;
            return this;
        }

        /**
         * @param roles Role to assign to this user. At least one existing role is required.
         * 
         * @return builder
         * 
         */
        public Builder roles(List<String> roles) {
            return roles(Output.of(roles));
        }

        /**
         * @param roles Role to assign to this user. At least one existing role is required.
         * 
         * @return builder
         * 
         */
        public Builder roles(String... roles) {
            return roles(List.of(roles));
        }

        /**
         * @param tz User timezone.
         * 
         * @return builder
         * 
         */
        public Builder tz(@Nullable Output<String> tz) {
            $.tz = tz;
            return this;
        }

        /**
         * @param tz User timezone.
         * 
         * @return builder
         * 
         */
        public Builder tz(String tz) {
            return tz(Output.of(tz));
        }

        public AuthenticationUsersState build() {
            return $;
        }
    }

}
