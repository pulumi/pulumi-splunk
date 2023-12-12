// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.splunk;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.splunk.InputsHttpEventCollectorArgs;
import com.pulumi.splunk.Utilities;
import com.pulumi.splunk.inputs.InputsHttpEventCollectorState;
import com.pulumi.splunk.outputs.InputsHttpEventCollectorAcl;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import javax.annotation.Nullable;

/**
 * ## # Resource: splunk.InputsHttpEventCollector
 * 
 * Create or update HTTP Event Collector input configuration tokens.
 * 
 */
@ResourceType(type="splunk:index/inputsHttpEventCollector:InputsHttpEventCollector")
public class InputsHttpEventCollector extends com.pulumi.resources.CustomResource {
    /**
     * The app/user context that is the namespace for the resource
     * 
     */
    @Export(name="acl", refs={InputsHttpEventCollectorAcl.class}, tree="[0]")
    private Output<InputsHttpEventCollectorAcl> acl;

    /**
     * @return The app/user context that is the namespace for the resource
     * 
     */
    public Output<InputsHttpEventCollectorAcl> acl() {
        return this.acl;
    }
    /**
     * Input disabled indicator
     * 
     */
    @Export(name="disabled", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> disabled;

    /**
     * @return Input disabled indicator
     * 
     */
    public Output<Boolean> disabled() {
        return this.disabled;
    }
    /**
     * Default host value for events with this token
     * 
     */
    @Export(name="host", refs={String.class}, tree="[0]")
    private Output<String> host;

    /**
     * @return Default host value for events with this token
     * 
     */
    public Output<String> host() {
        return this.host;
    }
    /**
     * Index to store generated events
     * 
     */
    @Export(name="index", refs={String.class}, tree="[0]")
    private Output<String> index;

    /**
     * @return Index to store generated events
     * 
     */
    public Output<String> index() {
        return this.index;
    }
    /**
     * Set of indexes allowed for events with this token
     * 
     */
    @Export(name="indexes", refs={List.class,String.class}, tree="[0,1]")
    private Output<List<String>> indexes;

    /**
     * @return Set of indexes allowed for events with this token
     * 
     */
    public Output<List<String>> indexes() {
        return this.indexes;
    }
    /**
     * Token name (inputs.conf key)
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return Token name (inputs.conf key)
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * Default source for events with this token
     * 
     */
    @Export(name="source", refs={String.class}, tree="[0]")
    private Output<String> source;

    /**
     * @return Default source for events with this token
     * 
     */
    public Output<String> source() {
        return this.source;
    }
    /**
     * Default source type for events with this token
     * 
     */
    @Export(name="sourcetype", refs={String.class}, tree="[0]")
    private Output<String> sourcetype;

    /**
     * @return Default source type for events with this token
     * 
     */
    public Output<String> sourcetype() {
        return this.sourcetype;
    }
    /**
     * Token value for sending data to collector/event endpoint
     * 
     */
    @Export(name="token", refs={String.class}, tree="[0]")
    private Output<String> token;

    /**
     * @return Token value for sending data to collector/event endpoint
     * 
     */
    public Output<String> token() {
        return this.token;
    }
    /**
     * Indexer acknowledgement for this token
     * 
     */
    @Export(name="useAck", refs={Integer.class}, tree="[0]")
    private Output<Integer> useAck;

    /**
     * @return Indexer acknowledgement for this token
     * 
     */
    public Output<Integer> useAck() {
        return this.useAck;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public InputsHttpEventCollector(String name) {
        this(name, InputsHttpEventCollectorArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public InputsHttpEventCollector(String name, @Nullable InputsHttpEventCollectorArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public InputsHttpEventCollector(String name, @Nullable InputsHttpEventCollectorArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("splunk:index/inputsHttpEventCollector:InputsHttpEventCollector", name, args == null ? InputsHttpEventCollectorArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private InputsHttpEventCollector(String name, Output<String> id, @Nullable InputsHttpEventCollectorState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("splunk:index/inputsHttpEventCollector:InputsHttpEventCollector", name, state, makeResourceOptions(options, id));
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
    public static InputsHttpEventCollector get(String name, Output<String> id, @Nullable InputsHttpEventCollectorState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new InputsHttpEventCollector(name, id, state, options);
    }
}
