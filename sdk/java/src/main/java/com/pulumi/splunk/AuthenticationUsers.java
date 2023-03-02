// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.splunk;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.splunk.AuthenticationUsersArgs;
import com.pulumi.splunk.Utilities;
import com.pulumi.splunk.inputs.AuthenticationUsersState;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * ## # Resource: splunk.AuthenticationUsers
 * 
 * Create and update user information or delete the user.
 * 
 * ## Example Usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.splunk.AuthenticationUsers;
 * import com.pulumi.splunk.AuthenticationUsersArgs;
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
 *         var user01 = new AuthenticationUsers(&#34;user01&#34;, AuthenticationUsersArgs.builder()        
 *             .email(&#34;user01@example.com&#34;)
 *             .forceChangePass(false)
 *             .password(&#34;password01&#34;)
 *             .roles(&#34;terraform-user01-role&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 */
@ResourceType(type="splunk:index/authenticationUsers:AuthenticationUsers")
public class AuthenticationUsers extends com.pulumi.resources.CustomResource {
    /**
     * User default app. Overrides the default app inherited from the user roles.
     * 
     */
    @Export(name="defaultApp", type=String.class, parameters={})
    private Output<String> defaultApp;

    /**
     * @return User default app. Overrides the default app inherited from the user roles.
     * 
     */
    public Output<String> defaultApp() {
        return this.defaultApp;
    }
    /**
     * User email address.
     * 
     */
    @Export(name="email", type=String.class, parameters={})
    private Output<String> email;

    /**
     * @return User email address.
     * 
     */
    public Output<String> email() {
        return this.email;
    }
    /**
     * Force user to change password indication
     * 
     */
    @Export(name="forceChangePass", type=Boolean.class, parameters={})
    private Output</* @Nullable */ Boolean> forceChangePass;

    /**
     * @return Force user to change password indication
     * 
     */
    public Output<Optional<Boolean>> forceChangePass() {
        return Codegen.optional(this.forceChangePass);
    }
    /**
     * Unique user login name.
     * 
     */
    @Export(name="name", type=String.class, parameters={})
    private Output<String> name;

    /**
     * @return Unique user login name.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * User login password.
     * 
     */
    @Export(name="password", type=String.class, parameters={})
    private Output</* @Nullable */ String> password;

    /**
     * @return User login password.
     * 
     */
    public Output<Optional<String>> password() {
        return Codegen.optional(this.password);
    }
    /**
     * Full user name.
     * 
     */
    @Export(name="realname", type=String.class, parameters={})
    private Output<String> realname;

    /**
     * @return Full user name.
     * 
     */
    public Output<String> realname() {
        return this.realname;
    }
    /**
     * Restart background search job that has not completed when Splunk restarts indication.
     * 
     */
    @Export(name="restartBackgroundJobs", type=Boolean.class, parameters={})
    private Output<Boolean> restartBackgroundJobs;

    /**
     * @return Restart background search job that has not completed when Splunk restarts indication.
     * 
     */
    public Output<Boolean> restartBackgroundJobs() {
        return this.restartBackgroundJobs;
    }
    /**
     * Role to assign to this user. At least one existing role is required.
     * 
     */
    @Export(name="roles", type=List.class, parameters={String.class})
    private Output<List<String>> roles;

    /**
     * @return Role to assign to this user. At least one existing role is required.
     * 
     */
    public Output<List<String>> roles() {
        return this.roles;
    }
    /**
     * User timezone.
     * 
     */
    @Export(name="tz", type=String.class, parameters={})
    private Output<String> tz;

    /**
     * @return User timezone.
     * 
     */
    public Output<String> tz() {
        return this.tz;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public AuthenticationUsers(String name) {
        this(name, AuthenticationUsersArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public AuthenticationUsers(String name, @Nullable AuthenticationUsersArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public AuthenticationUsers(String name, @Nullable AuthenticationUsersArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("splunk:index/authenticationUsers:AuthenticationUsers", name, args == null ? AuthenticationUsersArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private AuthenticationUsers(String name, Output<String> id, @Nullable AuthenticationUsersState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("splunk:index/authenticationUsers:AuthenticationUsers", name, state, makeResourceOptions(options, id));
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
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
    public static AuthenticationUsers get(String name, Output<String> id, @Nullable AuthenticationUsersState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new AuthenticationUsers(name, id, state, options);
    }
}