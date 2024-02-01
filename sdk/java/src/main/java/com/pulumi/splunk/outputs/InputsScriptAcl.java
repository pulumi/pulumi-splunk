// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.splunk.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class InputsScriptAcl {
    /**
     * @return The app context for the resource. Required for updating saved search ACL properties. Allowed values are:The name of an app and system
     * 
     */
    private @Nullable String app;
    /**
     * @return Indicates if the active user can change permissions for this object. Defaults to true.
     * 
     */
    private @Nullable Boolean canChangePerms;
    /**
     * @return Indicates if the active user can change sharing to app level. Defaults to true.
     * 
     */
    private @Nullable Boolean canShareApp;
    /**
     * @return Indicates if the active user can change sharing to system level. Defaults to true.
     * 
     */
    private @Nullable Boolean canShareGlobal;
    /**
     * @return Indicates if the active user can change sharing to user level. Defaults to true.
     * 
     */
    private @Nullable Boolean canShareUser;
    /**
     * @return Indicates if the active user can edit this object. Defaults to true.
     * 
     */
    private @Nullable Boolean canWrite;
    /**
     * @return User name of resource owner. Defaults to the resource creator. Required for updating any knowledge object ACL properties.nobody = All users may access the resource, but write access to the resource might be restricted.
     * 
     */
    private @Nullable String owner;
    /**
     * @return Properties that indicate resource read permissions.
     * 
     */
    private @Nullable List<String> reads;
    /**
     * @return Indicates whether an admin or user with sufficient permissions can delete the entity.
     * 
     */
    private @Nullable Boolean removable;
    /**
     * @return Indicates how the resource is shared. Required for updating any knowledge object ACL properties.app: Shared within a specific appglobal: (Default) Shared globally to all apps.user: Private to a user
     * 
     */
    private @Nullable String sharing;
    /**
     * @return Properties that indicate write permissions of the resource.
     * 
     */
    private @Nullable List<String> writes;

    private InputsScriptAcl() {}
    /**
     * @return The app context for the resource. Required for updating saved search ACL properties. Allowed values are:The name of an app and system
     * 
     */
    public Optional<String> app() {
        return Optional.ofNullable(this.app);
    }
    /**
     * @return Indicates if the active user can change permissions for this object. Defaults to true.
     * 
     */
    public Optional<Boolean> canChangePerms() {
        return Optional.ofNullable(this.canChangePerms);
    }
    /**
     * @return Indicates if the active user can change sharing to app level. Defaults to true.
     * 
     */
    public Optional<Boolean> canShareApp() {
        return Optional.ofNullable(this.canShareApp);
    }
    /**
     * @return Indicates if the active user can change sharing to system level. Defaults to true.
     * 
     */
    public Optional<Boolean> canShareGlobal() {
        return Optional.ofNullable(this.canShareGlobal);
    }
    /**
     * @return Indicates if the active user can change sharing to user level. Defaults to true.
     * 
     */
    public Optional<Boolean> canShareUser() {
        return Optional.ofNullable(this.canShareUser);
    }
    /**
     * @return Indicates if the active user can edit this object. Defaults to true.
     * 
     */
    public Optional<Boolean> canWrite() {
        return Optional.ofNullable(this.canWrite);
    }
    /**
     * @return User name of resource owner. Defaults to the resource creator. Required for updating any knowledge object ACL properties.nobody = All users may access the resource, but write access to the resource might be restricted.
     * 
     */
    public Optional<String> owner() {
        return Optional.ofNullable(this.owner);
    }
    /**
     * @return Properties that indicate resource read permissions.
     * 
     */
    public List<String> reads() {
        return this.reads == null ? List.of() : this.reads;
    }
    /**
     * @return Indicates whether an admin or user with sufficient permissions can delete the entity.
     * 
     */
    public Optional<Boolean> removable() {
        return Optional.ofNullable(this.removable);
    }
    /**
     * @return Indicates how the resource is shared. Required for updating any knowledge object ACL properties.app: Shared within a specific appglobal: (Default) Shared globally to all apps.user: Private to a user
     * 
     */
    public Optional<String> sharing() {
        return Optional.ofNullable(this.sharing);
    }
    /**
     * @return Properties that indicate write permissions of the resource.
     * 
     */
    public List<String> writes() {
        return this.writes == null ? List.of() : this.writes;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(InputsScriptAcl defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable String app;
        private @Nullable Boolean canChangePerms;
        private @Nullable Boolean canShareApp;
        private @Nullable Boolean canShareGlobal;
        private @Nullable Boolean canShareUser;
        private @Nullable Boolean canWrite;
        private @Nullable String owner;
        private @Nullable List<String> reads;
        private @Nullable Boolean removable;
        private @Nullable String sharing;
        private @Nullable List<String> writes;
        public Builder() {}
        public Builder(InputsScriptAcl defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.app = defaults.app;
    	      this.canChangePerms = defaults.canChangePerms;
    	      this.canShareApp = defaults.canShareApp;
    	      this.canShareGlobal = defaults.canShareGlobal;
    	      this.canShareUser = defaults.canShareUser;
    	      this.canWrite = defaults.canWrite;
    	      this.owner = defaults.owner;
    	      this.reads = defaults.reads;
    	      this.removable = defaults.removable;
    	      this.sharing = defaults.sharing;
    	      this.writes = defaults.writes;
        }

        @CustomType.Setter
        public Builder app(@Nullable String app) {

            this.app = app;
            return this;
        }
        @CustomType.Setter
        public Builder canChangePerms(@Nullable Boolean canChangePerms) {

            this.canChangePerms = canChangePerms;
            return this;
        }
        @CustomType.Setter
        public Builder canShareApp(@Nullable Boolean canShareApp) {

            this.canShareApp = canShareApp;
            return this;
        }
        @CustomType.Setter
        public Builder canShareGlobal(@Nullable Boolean canShareGlobal) {

            this.canShareGlobal = canShareGlobal;
            return this;
        }
        @CustomType.Setter
        public Builder canShareUser(@Nullable Boolean canShareUser) {

            this.canShareUser = canShareUser;
            return this;
        }
        @CustomType.Setter
        public Builder canWrite(@Nullable Boolean canWrite) {

            this.canWrite = canWrite;
            return this;
        }
        @CustomType.Setter
        public Builder owner(@Nullable String owner) {

            this.owner = owner;
            return this;
        }
        @CustomType.Setter
        public Builder reads(@Nullable List<String> reads) {

            this.reads = reads;
            return this;
        }
        public Builder reads(String... reads) {
            return reads(List.of(reads));
        }
        @CustomType.Setter
        public Builder removable(@Nullable Boolean removable) {

            this.removable = removable;
            return this;
        }
        @CustomType.Setter
        public Builder sharing(@Nullable String sharing) {

            this.sharing = sharing;
            return this;
        }
        @CustomType.Setter
        public Builder writes(@Nullable List<String> writes) {

            this.writes = writes;
            return this;
        }
        public Builder writes(String... writes) {
            return writes(List.of(writes));
        }
        public InputsScriptAcl build() {
            final var _resultValue = new InputsScriptAcl();
            _resultValue.app = app;
            _resultValue.canChangePerms = canChangePerms;
            _resultValue.canShareApp = canShareApp;
            _resultValue.canShareGlobal = canShareGlobal;
            _resultValue.canShareUser = canShareUser;
            _resultValue.canWrite = canWrite;
            _resultValue.owner = owner;
            _resultValue.reads = reads;
            _resultValue.removable = removable;
            _resultValue.sharing = sharing;
            _resultValue.writes = writes;
            return _resultValue;
        }
    }
}
