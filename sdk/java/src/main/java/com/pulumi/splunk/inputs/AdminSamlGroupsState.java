// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.splunk.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class AdminSamlGroupsState extends com.pulumi.resources.ResourceArgs {

    public static final AdminSamlGroupsState Empty = new AdminSamlGroupsState();

    /**
     * The name of the external group.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return The name of the external group.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * List of internal roles assigned to the group.
     * 
     */
    @Import(name="roles")
    private @Nullable Output<List<String>> roles;

    /**
     * @return List of internal roles assigned to the group.
     * 
     */
    public Optional<Output<List<String>>> roles() {
        return Optional.ofNullable(this.roles);
    }

    private AdminSamlGroupsState() {}

    private AdminSamlGroupsState(AdminSamlGroupsState $) {
        this.name = $.name;
        this.roles = $.roles;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(AdminSamlGroupsState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private AdminSamlGroupsState $;

        public Builder() {
            $ = new AdminSamlGroupsState();
        }

        public Builder(AdminSamlGroupsState defaults) {
            $ = new AdminSamlGroupsState(Objects.requireNonNull(defaults));
        }

        /**
         * @param name The name of the external group.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name The name of the external group.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param roles List of internal roles assigned to the group.
         * 
         * @return builder
         * 
         */
        public Builder roles(@Nullable Output<List<String>> roles) {
            $.roles = roles;
            return this;
        }

        /**
         * @param roles List of internal roles assigned to the group.
         * 
         * @return builder
         * 
         */
        public Builder roles(List<String> roles) {
            return roles(Output.of(roles));
        }

        /**
         * @param roles List of internal roles assigned to the group.
         * 
         * @return builder
         * 
         */
        public Builder roles(String... roles) {
            return roles(List.of(roles));
        }

        public AdminSamlGroupsState build() {
            return $;
        }
    }

}
