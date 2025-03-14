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


public final class LookupTableFileState extends com.pulumi.resources.ResourceArgs {

    public static final LookupTableFileState Empty = new LookupTableFileState();

    /**
     * The app context for the resource.
     * 
     */
    @Import(name="app")
    private @Nullable Output<String> app;

    /**
     * @return The app context for the resource.
     * 
     */
    public Optional<Output<String>> app() {
        return Optional.ofNullable(this.app);
    }

    /**
     * The column header and row value contents for the lookup table file.
     * 
     */
    @Import(name="fileContents")
    private @Nullable Output<List<List<String>>> fileContents;

    /**
     * @return The column header and row value contents for the lookup table file.
     * 
     */
    public Optional<Output<List<List<String>>>> fileContents() {
        return Optional.ofNullable(this.fileContents);
    }

    /**
     * A name for the lookup table file. Generally ends with &#34;.csv&#34;
     * 
     */
    @Import(name="fileName")
    private @Nullable Output<String> fileName;

    /**
     * @return A name for the lookup table file. Generally ends with &#34;.csv&#34;
     * 
     */
    public Optional<Output<String>> fileName() {
        return Optional.ofNullable(this.fileName);
    }

    /**
     * User name of resource owner. Defaults to the resource creator. Required for updating any knowledge object ACL properties. nobody = All users may access the resource, but write access to the resource might be restricted.
     * 
     */
    @Import(name="owner")
    private @Nullable Output<String> owner;

    /**
     * @return User name of resource owner. Defaults to the resource creator. Required for updating any knowledge object ACL properties. nobody = All users may access the resource, but write access to the resource might be restricted.
     * 
     */
    public Optional<Output<String>> owner() {
        return Optional.ofNullable(this.owner);
    }

    private LookupTableFileState() {}

    private LookupTableFileState(LookupTableFileState $) {
        this.app = $.app;
        this.fileContents = $.fileContents;
        this.fileName = $.fileName;
        this.owner = $.owner;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(LookupTableFileState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private LookupTableFileState $;

        public Builder() {
            $ = new LookupTableFileState();
        }

        public Builder(LookupTableFileState defaults) {
            $ = new LookupTableFileState(Objects.requireNonNull(defaults));
        }

        /**
         * @param app The app context for the resource.
         * 
         * @return builder
         * 
         */
        public Builder app(@Nullable Output<String> app) {
            $.app = app;
            return this;
        }

        /**
         * @param app The app context for the resource.
         * 
         * @return builder
         * 
         */
        public Builder app(String app) {
            return app(Output.of(app));
        }

        /**
         * @param fileContents The column header and row value contents for the lookup table file.
         * 
         * @return builder
         * 
         */
        public Builder fileContents(@Nullable Output<List<List<String>>> fileContents) {
            $.fileContents = fileContents;
            return this;
        }

        /**
         * @param fileContents The column header and row value contents for the lookup table file.
         * 
         * @return builder
         * 
         */
        public Builder fileContents(List<List<String>> fileContents) {
            return fileContents(Output.of(fileContents));
        }

        /**
         * @param fileContents The column header and row value contents for the lookup table file.
         * 
         * @return builder
         * 
         */
        public Builder fileContents(List<String>... fileContents) {
            return fileContents(List.of(fileContents));
        }

        /**
         * @param fileName A name for the lookup table file. Generally ends with &#34;.csv&#34;
         * 
         * @return builder
         * 
         */
        public Builder fileName(@Nullable Output<String> fileName) {
            $.fileName = fileName;
            return this;
        }

        /**
         * @param fileName A name for the lookup table file. Generally ends with &#34;.csv&#34;
         * 
         * @return builder
         * 
         */
        public Builder fileName(String fileName) {
            return fileName(Output.of(fileName));
        }

        /**
         * @param owner User name of resource owner. Defaults to the resource creator. Required for updating any knowledge object ACL properties. nobody = All users may access the resource, but write access to the resource might be restricted.
         * 
         * @return builder
         * 
         */
        public Builder owner(@Nullable Output<String> owner) {
            $.owner = owner;
            return this;
        }

        /**
         * @param owner User name of resource owner. Defaults to the resource creator. Required for updating any knowledge object ACL properties. nobody = All users may access the resource, but write access to the resource might be restricted.
         * 
         * @return builder
         * 
         */
        public Builder owner(String owner) {
            return owner(Output.of(owner));
        }

        public LookupTableFileState build() {
            return $;
        }
    }

}
