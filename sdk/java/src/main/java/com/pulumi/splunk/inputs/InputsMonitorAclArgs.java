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


public final class InputsMonitorAclArgs extends com.pulumi.resources.ResourceArgs {

    public static final InputsMonitorAclArgs Empty = new InputsMonitorAclArgs();

    @Import(name="app")
    private @Nullable Output<String> app;

    public Optional<Output<String>> app() {
        return Optional.ofNullable(this.app);
    }

    @Import(name="canChangePerms")
    private @Nullable Output<Boolean> canChangePerms;

    public Optional<Output<Boolean>> canChangePerms() {
        return Optional.ofNullable(this.canChangePerms);
    }

    @Import(name="canShareApp")
    private @Nullable Output<Boolean> canShareApp;

    public Optional<Output<Boolean>> canShareApp() {
        return Optional.ofNullable(this.canShareApp);
    }

    @Import(name="canShareGlobal")
    private @Nullable Output<Boolean> canShareGlobal;

    public Optional<Output<Boolean>> canShareGlobal() {
        return Optional.ofNullable(this.canShareGlobal);
    }

    @Import(name="canShareUser")
    private @Nullable Output<Boolean> canShareUser;

    public Optional<Output<Boolean>> canShareUser() {
        return Optional.ofNullable(this.canShareUser);
    }

    @Import(name="canWrite")
    private @Nullable Output<Boolean> canWrite;

    public Optional<Output<Boolean>> canWrite() {
        return Optional.ofNullable(this.canWrite);
    }

    @Import(name="owner")
    private @Nullable Output<String> owner;

    public Optional<Output<String>> owner() {
        return Optional.ofNullable(this.owner);
    }

    @Import(name="reads")
    private @Nullable Output<List<String>> reads;

    public Optional<Output<List<String>>> reads() {
        return Optional.ofNullable(this.reads);
    }

    @Import(name="removable")
    private @Nullable Output<Boolean> removable;

    public Optional<Output<Boolean>> removable() {
        return Optional.ofNullable(this.removable);
    }

    @Import(name="sharing")
    private @Nullable Output<String> sharing;

    public Optional<Output<String>> sharing() {
        return Optional.ofNullable(this.sharing);
    }

    @Import(name="writes")
    private @Nullable Output<List<String>> writes;

    public Optional<Output<List<String>>> writes() {
        return Optional.ofNullable(this.writes);
    }

    private InputsMonitorAclArgs() {}

    private InputsMonitorAclArgs(InputsMonitorAclArgs $) {
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
    public static Builder builder(InputsMonitorAclArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private InputsMonitorAclArgs $;

        public Builder() {
            $ = new InputsMonitorAclArgs();
        }

        public Builder(InputsMonitorAclArgs defaults) {
            $ = new InputsMonitorAclArgs(Objects.requireNonNull(defaults));
        }

        public Builder app(@Nullable Output<String> app) {
            $.app = app;
            return this;
        }

        public Builder app(String app) {
            return app(Output.of(app));
        }

        public Builder canChangePerms(@Nullable Output<Boolean> canChangePerms) {
            $.canChangePerms = canChangePerms;
            return this;
        }

        public Builder canChangePerms(Boolean canChangePerms) {
            return canChangePerms(Output.of(canChangePerms));
        }

        public Builder canShareApp(@Nullable Output<Boolean> canShareApp) {
            $.canShareApp = canShareApp;
            return this;
        }

        public Builder canShareApp(Boolean canShareApp) {
            return canShareApp(Output.of(canShareApp));
        }

        public Builder canShareGlobal(@Nullable Output<Boolean> canShareGlobal) {
            $.canShareGlobal = canShareGlobal;
            return this;
        }

        public Builder canShareGlobal(Boolean canShareGlobal) {
            return canShareGlobal(Output.of(canShareGlobal));
        }

        public Builder canShareUser(@Nullable Output<Boolean> canShareUser) {
            $.canShareUser = canShareUser;
            return this;
        }

        public Builder canShareUser(Boolean canShareUser) {
            return canShareUser(Output.of(canShareUser));
        }

        public Builder canWrite(@Nullable Output<Boolean> canWrite) {
            $.canWrite = canWrite;
            return this;
        }

        public Builder canWrite(Boolean canWrite) {
            return canWrite(Output.of(canWrite));
        }

        public Builder owner(@Nullable Output<String> owner) {
            $.owner = owner;
            return this;
        }

        public Builder owner(String owner) {
            return owner(Output.of(owner));
        }

        public Builder reads(@Nullable Output<List<String>> reads) {
            $.reads = reads;
            return this;
        }

        public Builder reads(List<String> reads) {
            return reads(Output.of(reads));
        }

        public Builder reads(String... reads) {
            return reads(List.of(reads));
        }

        public Builder removable(@Nullable Output<Boolean> removable) {
            $.removable = removable;
            return this;
        }

        public Builder removable(Boolean removable) {
            return removable(Output.of(removable));
        }

        public Builder sharing(@Nullable Output<String> sharing) {
            $.sharing = sharing;
            return this;
        }

        public Builder sharing(String sharing) {
            return sharing(Output.of(sharing));
        }

        public Builder writes(@Nullable Output<List<String>> writes) {
            $.writes = writes;
            return this;
        }

        public Builder writes(List<String> writes) {
            return writes(Output.of(writes));
        }

        public Builder writes(String... writes) {
            return writes(List.of(writes));
        }

        public InputsMonitorAclArgs build() {
            return $;
        }
    }

}
