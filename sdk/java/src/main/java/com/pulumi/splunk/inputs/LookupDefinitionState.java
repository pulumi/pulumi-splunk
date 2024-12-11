// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.splunk.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.splunk.inputs.LookupDefinitionAclArgs;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class LookupDefinitionState extends com.pulumi.resources.ResourceArgs {

    public static final LookupDefinitionState Empty = new LookupDefinitionState();

    /**
     * Defines the access control list (ACL) for the lookup definition. See acl.md for more details.
     * 
     */
    @Import(name="acl")
    private @Nullable Output<LookupDefinitionAclArgs> acl;

    /**
     * @return Defines the access control list (ACL) for the lookup definition. See acl.md for more details.
     * 
     */
    public Optional<Output<LookupDefinitionAclArgs>> acl() {
        return Optional.ofNullable(this.acl);
    }

    /**
     * The filename for the lookup table, usually ending in `.csv`.
     * 
     */
    @Import(name="filename")
    private @Nullable Output<String> filename;

    /**
     * @return The filename for the lookup table, usually ending in `.csv`.
     * 
     */
    public Optional<Output<String>> filename() {
        return Optional.ofNullable(this.filename);
    }

    /**
     * A unique name for the lookup definition within the app context.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return A unique name for the lookup definition within the app context.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    private LookupDefinitionState() {}

    private LookupDefinitionState(LookupDefinitionState $) {
        this.acl = $.acl;
        this.filename = $.filename;
        this.name = $.name;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(LookupDefinitionState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private LookupDefinitionState $;

        public Builder() {
            $ = new LookupDefinitionState();
        }

        public Builder(LookupDefinitionState defaults) {
            $ = new LookupDefinitionState(Objects.requireNonNull(defaults));
        }

        /**
         * @param acl Defines the access control list (ACL) for the lookup definition. See acl.md for more details.
         * 
         * @return builder
         * 
         */
        public Builder acl(@Nullable Output<LookupDefinitionAclArgs> acl) {
            $.acl = acl;
            return this;
        }

        /**
         * @param acl Defines the access control list (ACL) for the lookup definition. See acl.md for more details.
         * 
         * @return builder
         * 
         */
        public Builder acl(LookupDefinitionAclArgs acl) {
            return acl(Output.of(acl));
        }

        /**
         * @param filename The filename for the lookup table, usually ending in `.csv`.
         * 
         * @return builder
         * 
         */
        public Builder filename(@Nullable Output<String> filename) {
            $.filename = filename;
            return this;
        }

        /**
         * @param filename The filename for the lookup table, usually ending in `.csv`.
         * 
         * @return builder
         * 
         */
        public Builder filename(String filename) {
            return filename(Output.of(filename));
        }

        /**
         * @param name A unique name for the lookup definition within the app context.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name A unique name for the lookup definition within the app context.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        public LookupDefinitionState build() {
            return $;
        }
    }

}