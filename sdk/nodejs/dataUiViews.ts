// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * ## # Resource: splunk.DataUiViews
 *
 * Create and manage splunk dashboards/views.
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as splunk from "@pulumi/splunk";
 *
 * const dashboard = new splunk.DataUiViews("dashboard", {
 *     acl: {
 *         app: "search",
 *         owner: "admin",
 *     },
 *     eaiData: "<dashboard version=\"1.1\"><label>Terraform</label><description>Terraform operations</description><row><panel><chart><search><query>index=_internal sourcetype=splunkd_access useragent=\"splunk-simple-go-client\" | timechart fixedrange=f values(status) by uri_path</query><earliest>-24h@h</earliest><latest>now</latest><sampleRatio>1</sampleRatio></search><option name=\"charting.axisLabelsX.majorLabelStyle.overflowMode\">ellipsisNone</option><option name=\"charting.axisLabelsX.majorLabelStyle.rotation\">0</option><option name=\"charting.axisTitleX.visibility\">collapsed</option><option name=\"charting.axisTitleY.text\">HTTP status codes</option><option name=\"charting.axisTitleY.visibility\">visible</option><option name=\"charting.axisTitleY2.visibility\">visible</option><option name=\"charting.axisX.abbreviation\">none</option><option name=\"charting.axisX.scale\">linear</option><option name=\"charting.axisY.abbreviation\">none</option><option name=\"charting.axisY.scale\">linear</option><option name=\"charting.axisY2.abbreviation\">none</option><option name=\"charting.axisY2.enabled\">0</option><option name=\"charting.axisY2.scale\">inherit</option><option name=\"charting.chart\">column</option><option name=\"charting.chart.bubbleMaximumSize\">50</option><option name=\"charting.chart.bubbleMinimumSize\">10</option><option name=\"charting.chart.bubbleSizeBy\">area</option><option name=\"charting.chart.nullValueMode\">connect</option><option name=\"charting.chart.showDataLabels\">none</option><option name=\"charting.chart.sliceCollapsingThreshold\">0.01</option><option name=\"charting.chart.stackMode\">default</option><option name=\"charting.chart.style\">shiny</option><option name=\"charting.drilldown\">none</option><option name=\"charting.layout.splitSeries\">0</option><option name=\"charting.layout.splitSeries.allowIndependentYRanges\">0</option><option name=\"charting.legend.labelStyle.overflowMode\">ellipsisMiddle</option><option name=\"charting.legend.mode\">standard</option><option name=\"charting.legend.placement\">right</option><option name=\"charting.lineWidth\">2</option><option name=\"trellis.enabled\">0</option><option name=\"trellis.scales.shared\">1</option><option name=\"trellis.size\">small</option><option name=\"trellis.splitBy\">_aggregation</option></chart></panel></row></dashboard>",
 * });
 * ```
 */
export class DataUiViews extends pulumi.CustomResource {
    /**
     * Get an existing DataUiViews resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: DataUiViewsState, opts?: pulumi.CustomResourceOptions): DataUiViews {
        return new DataUiViews(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'splunk:index/dataUiViews:DataUiViews';

    /**
     * Returns true if the given object is an instance of DataUiViews.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is DataUiViews {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === DataUiViews.__pulumiType;
    }

    public readonly acl!: pulumi.Output<outputs.DataUiViewsAcl>;
    /**
     * Dashboard XML definition.
     */
    public readonly eaiData!: pulumi.Output<string>;
    /**
     * Dashboard name.
     */
    public readonly name!: pulumi.Output<string>;

    /**
     * Create a DataUiViews resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DataUiViewsArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: DataUiViewsArgs | DataUiViewsState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as DataUiViewsState | undefined;
            resourceInputs["acl"] = state ? state.acl : undefined;
            resourceInputs["eaiData"] = state ? state.eaiData : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
        } else {
            const args = argsOrState as DataUiViewsArgs | undefined;
            if ((!args || args.eaiData === undefined) && !opts.urn) {
                throw new Error("Missing required property 'eaiData'");
            }
            resourceInputs["acl"] = args ? args.acl : undefined;
            resourceInputs["eaiData"] = args ? args.eaiData : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(DataUiViews.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering DataUiViews resources.
 */
export interface DataUiViewsState {
    acl?: pulumi.Input<inputs.DataUiViewsAcl>;
    /**
     * Dashboard XML definition.
     */
    eaiData?: pulumi.Input<string>;
    /**
     * Dashboard name.
     */
    name?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a DataUiViews resource.
 */
export interface DataUiViewsArgs {
    acl?: pulumi.Input<inputs.DataUiViewsAcl>;
    /**
     * Dashboard XML definition.
     */
    eaiData: pulumi.Input<string>;
    /**
     * Dashboard name.
     */
    name?: pulumi.Input<string>;
}
