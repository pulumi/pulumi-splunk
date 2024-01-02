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
public final class DataUiViewsAcl {
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

    private DataUiViewsAcl() {}
    public Optional<String> app() {
        return Optional.ofNullable(this.app);
    }
    public Optional<Boolean> canChangePerms() {
        return Optional.ofNullable(this.canChangePerms);
    }
    public Optional<Boolean> canShareApp() {
        return Optional.ofNullable(this.canShareApp);
    }
    public Optional<Boolean> canShareGlobal() {
        return Optional.ofNullable(this.canShareGlobal);
    }
    public Optional<Boolean> canShareUser() {
        return Optional.ofNullable(this.canShareUser);
    }
    public Optional<Boolean> canWrite() {
        return Optional.ofNullable(this.canWrite);
    }
    public Optional<String> owner() {
        return Optional.ofNullable(this.owner);
    }
    public List<String> reads() {
        return this.reads == null ? List.of() : this.reads;
    }
    public Optional<Boolean> removable() {
        return Optional.ofNullable(this.removable);
    }
    public Optional<String> sharing() {
        return Optional.ofNullable(this.sharing);
    }
    public List<String> writes() {
        return this.writes == null ? List.of() : this.writes;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(DataUiViewsAcl defaults) {
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
        public Builder(DataUiViewsAcl defaults) {
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
        public DataUiViewsAcl build() {
            final var _resultValue = new DataUiViewsAcl();
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
