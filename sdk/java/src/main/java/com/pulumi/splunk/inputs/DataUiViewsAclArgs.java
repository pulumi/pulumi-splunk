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


public final class DataUiViewsAclArgs extends com.pulumi.resources.ResourceArgs {

    public static final DataUiViewsAclArgs Empty = new DataUiViewsAclArgs();

    /**
     * The app context for the resource. Required for updating saved search ACL properties. Allowed values are:The name of an app and system
     * 
     */
    @Import(name="app")
    private @Nullable Output<String> app;

    /**
     * @return The app context for the resource. Required for updating saved search ACL properties. Allowed values are:The name of an app and system
     * 
     */
    public Optional<Output<String>> app() {
        return Optional.ofNullable(this.app);
    }

    /**
     * Indicates if the active user can change permissions for this object. Defaults to true.
     * 
     */
    @Import(name="canChangePerms")
    private @Nullable Output<Boolean> canChangePerms;

    /**
     * @return Indicates if the active user can change permissions for this object. Defaults to true.
     * 
     */
    public Optional<Output<Boolean>> canChangePerms() {
        return Optional.ofNullable(this.canChangePerms);
    }

    /**
     * Indicates if the active user can change sharing to app level. Defaults to true.
     * 
     */
    @Import(name="canShareApp")
    private @Nullable Output<Boolean> canShareApp;

    /**
     * @return Indicates if the active user can change sharing to app level. Defaults to true.
     * 
     */
    public Optional<Output<Boolean>> canShareApp() {
        return Optional.ofNullable(this.canShareApp);
    }

    /**
     * Indicates if the active user can change sharing to system level. Defaults to true.
     * 
     */
    @Import(name="canShareGlobal")
    private @Nullable Output<Boolean> canShareGlobal;

    /**
     * @return Indicates if the active user can change sharing to system level. Defaults to true.
     * 
     */
    public Optional<Output<Boolean>> canShareGlobal() {
        return Optional.ofNullable(this.canShareGlobal);
    }

    /**
     * Indicates if the active user can change sharing to user level. Defaults to true.
     * 
     */
    @Import(name="canShareUser")
    private @Nullable Output<Boolean> canShareUser;

    /**
     * @return Indicates if the active user can change sharing to user level. Defaults to true.
     * 
     */
    public Optional<Output<Boolean>> canShareUser() {
        return Optional.ofNullable(this.canShareUser);
    }

    /**
     * Indicates if the active user can edit this object. Defaults to true.
     * 
     */
    @Import(name="canWrite")
    private @Nullable Output<Boolean> canWrite;

    /**
     * @return Indicates if the active user can edit this object. Defaults to true.
     * 
     */
    public Optional<Output<Boolean>> canWrite() {
        return Optional.ofNullable(this.canWrite);
    }

    /**
     * User name of resource owner. Defaults to the resource creator. Required for updating any knowledge object ACL properties.nobody = All users may access the resource, but write access to the resource might be restricted.
     * 
     */
    @Import(name="owner")
    private @Nullable Output<String> owner;

    /**
     * @return User name of resource owner. Defaults to the resource creator. Required for updating any knowledge object ACL properties.nobody = All users may access the resource, but write access to the resource might be restricted.
     * 
     */
    public Optional<Output<String>> owner() {
        return Optional.ofNullable(this.owner);
    }

    /**
     * Properties that indicate resource read permissions.
     * 
     */
    @Import(name="reads")
    private @Nullable Output<List<String>> reads;

    /**
     * @return Properties that indicate resource read permissions.
     * 
     */
    public Optional<Output<List<String>>> reads() {
        return Optional.ofNullable(this.reads);
    }

    /**
     * Indicates whether an admin or user with sufficient permissions can delete the entity.
     * 
     */
    @Import(name="removable")
    private @Nullable Output<Boolean> removable;

    /**
     * @return Indicates whether an admin or user with sufficient permissions can delete the entity.
     * 
     */
    public Optional<Output<Boolean>> removable() {
        return Optional.ofNullable(this.removable);
    }

    /**
     * Indicates how the resource is shared. Required for updating any knowledge object ACL properties.app: Shared within a specific appglobal: (Default) Shared globally to all apps.user: Private to a user
     * 
     */
    @Import(name="sharing")
    private @Nullable Output<String> sharing;

    /**
     * @return Indicates how the resource is shared. Required for updating any knowledge object ACL properties.app: Shared within a specific appglobal: (Default) Shared globally to all apps.user: Private to a user
     * 
     */
    public Optional<Output<String>> sharing() {
        return Optional.ofNullable(this.sharing);
    }

    /**
     * Properties that indicate resource write permissions.
     * 
     */
    @Import(name="writes")
    private @Nullable Output<List<String>> writes;

    /**
     * @return Properties that indicate resource write permissions.
     * 
     */
    public Optional<Output<List<String>>> writes() {
        return Optional.ofNullable(this.writes);
    }

    private DataUiViewsAclArgs() {}

    private DataUiViewsAclArgs(DataUiViewsAclArgs $) {
        this.app = $.app;
        this.canChangePerms = $.canChangePerms;
        this.canShareApp = $.canShareApp;
        this.canShareGlobal = $.canShareGlobal;
        this.canShareUser = $.canShareUser;
        this.canWrite = $.canWrite;
        this.owner = $.owner;
        this.reads = $.reads;
        this.removable = $.removable;
        this.sharing = $.sharing;
        this.writes = $.writes;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(DataUiViewsAclArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private DataUiViewsAclArgs $;

        public Builder() {
            $ = new DataUiViewsAclArgs();
        }

        public Builder(DataUiViewsAclArgs defaults) {
            $ = new DataUiViewsAclArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param app The app context for the resource. Required for updating saved search ACL properties. Allowed values are:The name of an app and system
         * 
         * @return builder
         * 
         */
        public Builder app(@Nullable Output<String> app) {
            $.app = app;
            return this;
        }

        /**
         * @param app The app context for the resource. Required for updating saved search ACL properties. Allowed values are:The name of an app and system
         * 
         * @return builder
         * 
         */
        public Builder app(String app) {
            return app(Output.of(app));
        }

        /**
         * @param canChangePerms Indicates if the active user can change permissions for this object. Defaults to true.
         * 
         * @return builder
         * 
         */
        public Builder canChangePerms(@Nullable Output<Boolean> canChangePerms) {
            $.canChangePerms = canChangePerms;
            return this;
        }

        /**
         * @param canChangePerms Indicates if the active user can change permissions for this object. Defaults to true.
         * 
         * @return builder
         * 
         */
        public Builder canChangePerms(Boolean canChangePerms) {
            return canChangePerms(Output.of(canChangePerms));
        }

        /**
         * @param canShareApp Indicates if the active user can change sharing to app level. Defaults to true.
         * 
         * @return builder
         * 
         */
        public Builder canShareApp(@Nullable Output<Boolean> canShareApp) {
            $.canShareApp = canShareApp;
            return this;
        }

        /**
         * @param canShareApp Indicates if the active user can change sharing to app level. Defaults to true.
         * 
         * @return builder
         * 
         */
        public Builder canShareApp(Boolean canShareApp) {
            return canShareApp(Output.of(canShareApp));
        }

        /**
         * @param canShareGlobal Indicates if the active user can change sharing to system level. Defaults to true.
         * 
         * @return builder
         * 
         */
        public Builder canShareGlobal(@Nullable Output<Boolean> canShareGlobal) {
            $.canShareGlobal = canShareGlobal;
            return this;
        }

        /**
         * @param canShareGlobal Indicates if the active user can change sharing to system level. Defaults to true.
         * 
         * @return builder
         * 
         */
        public Builder canShareGlobal(Boolean canShareGlobal) {
            return canShareGlobal(Output.of(canShareGlobal));
        }

        /**
         * @param canShareUser Indicates if the active user can change sharing to user level. Defaults to true.
         * 
         * @return builder
         * 
         */
        public Builder canShareUser(@Nullable Output<Boolean> canShareUser) {
            $.canShareUser = canShareUser;
            return this;
        }

        /**
         * @param canShareUser Indicates if the active user can change sharing to user level. Defaults to true.
         * 
         * @return builder
         * 
         */
        public Builder canShareUser(Boolean canShareUser) {
            return canShareUser(Output.of(canShareUser));
        }

        /**
         * @param canWrite Indicates if the active user can edit this object. Defaults to true.
         * 
         * @return builder
         * 
         */
        public Builder canWrite(@Nullable Output<Boolean> canWrite) {
            $.canWrite = canWrite;
            return this;
        }

        /**
         * @param canWrite Indicates if the active user can edit this object. Defaults to true.
         * 
         * @return builder
         * 
         */
        public Builder canWrite(Boolean canWrite) {
            return canWrite(Output.of(canWrite));
        }

        /**
         * @param owner User name of resource owner. Defaults to the resource creator. Required for updating any knowledge object ACL properties.nobody = All users may access the resource, but write access to the resource might be restricted.
         * 
         * @return builder
         * 
         */
        public Builder owner(@Nullable Output<String> owner) {
            $.owner = owner;
            return this;
        }

        /**
         * @param owner User name of resource owner. Defaults to the resource creator. Required for updating any knowledge object ACL properties.nobody = All users may access the resource, but write access to the resource might be restricted.
         * 
         * @return builder
         * 
         */
        public Builder owner(String owner) {
            return owner(Output.of(owner));
        }

        /**
         * @param reads Properties that indicate resource read permissions.
         * 
         * @return builder
         * 
         */
        public Builder reads(@Nullable Output<List<String>> reads) {
            $.reads = reads;
            return this;
        }

        /**
         * @param reads Properties that indicate resource read permissions.
         * 
         * @return builder
         * 
         */
        public Builder reads(List<String> reads) {
            return reads(Output.of(reads));
        }

        /**
         * @param reads Properties that indicate resource read permissions.
         * 
         * @return builder
         * 
         */
        public Builder reads(String... reads) {
            return reads(List.of(reads));
        }

        /**
         * @param removable Indicates whether an admin or user with sufficient permissions can delete the entity.
         * 
         * @return builder
         * 
         */
        public Builder removable(@Nullable Output<Boolean> removable) {
            $.removable = removable;
            return this;
        }

        /**
         * @param removable Indicates whether an admin or user with sufficient permissions can delete the entity.
         * 
         * @return builder
         * 
         */
        public Builder removable(Boolean removable) {
            return removable(Output.of(removable));
        }

        /**
         * @param sharing Indicates how the resource is shared. Required for updating any knowledge object ACL properties.app: Shared within a specific appglobal: (Default) Shared globally to all apps.user: Private to a user
         * 
         * @return builder
         * 
         */
        public Builder sharing(@Nullable Output<String> sharing) {
            $.sharing = sharing;
            return this;
        }

        /**
         * @param sharing Indicates how the resource is shared. Required for updating any knowledge object ACL properties.app: Shared within a specific appglobal: (Default) Shared globally to all apps.user: Private to a user
         * 
         * @return builder
         * 
         */
        public Builder sharing(String sharing) {
            return sharing(Output.of(sharing));
        }

        /**
         * @param writes Properties that indicate resource write permissions.
         * 
         * @return builder
         * 
         */
        public Builder writes(@Nullable Output<List<String>> writes) {
            $.writes = writes;
            return this;
        }

        /**
         * @param writes Properties that indicate resource write permissions.
         * 
         * @return builder
         * 
         */
        public Builder writes(List<String> writes) {
            return writes(Output.of(writes));
        }

        /**
         * @param writes Properties that indicate resource write permissions.
         * 
         * @return builder
         * 
         */
        public Builder writes(String... writes) {
            return writes(List.of(writes));
        }

        public DataUiViewsAclArgs build() {
            return $;
        }
    }

}
