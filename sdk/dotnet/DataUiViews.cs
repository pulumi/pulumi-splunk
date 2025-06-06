// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Splunk
{
    /// <summary>
    /// ## # Resource: splunk.DataUiViews
    /// 
    /// Create and manage splunk dashboards/views.
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Splunk = Pulumi.Splunk;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var dashboard = new Splunk.DataUiViews("dashboard", new()
    ///     {
    ///         Name = "Terraform_Sample_Dashboard",
    ///         EaiData = "&lt;dashboard version=\"1.1\"&gt;&lt;label&gt;Terraform&lt;/label&gt;&lt;description&gt;Terraform operations&lt;/description&gt;&lt;row&gt;&lt;panel&gt;&lt;chart&gt;&lt;search&gt;&lt;query&gt;index=_internal sourcetype=splunkd_access useragent=\"splunk-simple-go-client\" | timechart fixedrange=f values(status) by uri_path&lt;/query&gt;&lt;earliest&gt;-24h@h&lt;/earliest&gt;&lt;latest&gt;now&lt;/latest&gt;&lt;sampleRatio&gt;1&lt;/sampleRatio&gt;&lt;/search&gt;&lt;option name=\"charting.axisLabelsX.majorLabelStyle.overflowMode\"&gt;ellipsisNone&lt;/option&gt;&lt;option name=\"charting.axisLabelsX.majorLabelStyle.rotation\"&gt;0&lt;/option&gt;&lt;option name=\"charting.axisTitleX.visibility\"&gt;collapsed&lt;/option&gt;&lt;option name=\"charting.axisTitleY.text\"&gt;HTTP status codes&lt;/option&gt;&lt;option name=\"charting.axisTitleY.visibility\"&gt;visible&lt;/option&gt;&lt;option name=\"charting.axisTitleY2.visibility\"&gt;visible&lt;/option&gt;&lt;option name=\"charting.axisX.abbreviation\"&gt;none&lt;/option&gt;&lt;option name=\"charting.axisX.scale\"&gt;linear&lt;/option&gt;&lt;option name=\"charting.axisY.abbreviation\"&gt;none&lt;/option&gt;&lt;option name=\"charting.axisY.scale\"&gt;linear&lt;/option&gt;&lt;option name=\"charting.axisY2.abbreviation\"&gt;none&lt;/option&gt;&lt;option name=\"charting.axisY2.enabled\"&gt;0&lt;/option&gt;&lt;option name=\"charting.axisY2.scale\"&gt;inherit&lt;/option&gt;&lt;option name=\"charting.chart\"&gt;column&lt;/option&gt;&lt;option name=\"charting.chart.bubbleMaximumSize\"&gt;50&lt;/option&gt;&lt;option name=\"charting.chart.bubbleMinimumSize\"&gt;10&lt;/option&gt;&lt;option name=\"charting.chart.bubbleSizeBy\"&gt;area&lt;/option&gt;&lt;option name=\"charting.chart.nullValueMode\"&gt;connect&lt;/option&gt;&lt;option name=\"charting.chart.showDataLabels\"&gt;none&lt;/option&gt;&lt;option name=\"charting.chart.sliceCollapsingThreshold\"&gt;0.01&lt;/option&gt;&lt;option name=\"charting.chart.stackMode\"&gt;default&lt;/option&gt;&lt;option name=\"charting.chart.style\"&gt;shiny&lt;/option&gt;&lt;option name=\"charting.drilldown\"&gt;none&lt;/option&gt;&lt;option name=\"charting.layout.splitSeries\"&gt;0&lt;/option&gt;&lt;option name=\"charting.layout.splitSeries.allowIndependentYRanges\"&gt;0&lt;/option&gt;&lt;option name=\"charting.legend.labelStyle.overflowMode\"&gt;ellipsisMiddle&lt;/option&gt;&lt;option name=\"charting.legend.mode\"&gt;standard&lt;/option&gt;&lt;option name=\"charting.legend.placement\"&gt;right&lt;/option&gt;&lt;option name=\"charting.lineWidth\"&gt;2&lt;/option&gt;&lt;option name=\"trellis.enabled\"&gt;0&lt;/option&gt;&lt;option name=\"trellis.scales.shared\"&gt;1&lt;/option&gt;&lt;option name=\"trellis.size\"&gt;small&lt;/option&gt;&lt;option name=\"trellis.splitBy\"&gt;_aggregation&lt;/option&gt;&lt;/chart&gt;&lt;/panel&gt;&lt;/row&gt;&lt;/dashboard&gt;",
    ///         Acl = new Splunk.Inputs.DataUiViewsAclArgs
    ///         {
    ///             Owner = "admin",
    ///             App = "search",
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// </summary>
    [SplunkResourceType("splunk:index/dataUiViews:DataUiViews")]
    public partial class DataUiViews : global::Pulumi.CustomResource
    {
        [Output("acl")]
        public Output<Outputs.DataUiViewsAcl> Acl { get; private set; } = null!;

        /// <summary>
        /// Dashboard XML definition.
        /// </summary>
        [Output("eaiData")]
        public Output<string> EaiData { get; private set; } = null!;

        /// <summary>
        /// Dashboard name.
        /// * `eai:data` - (Required) Dashboard XML definition.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;


        /// <summary>
        /// Create a DataUiViews resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public DataUiViews(string name, DataUiViewsArgs args, CustomResourceOptions? options = null)
            : base("splunk:index/dataUiViews:DataUiViews", name, args ?? new DataUiViewsArgs(), MakeResourceOptions(options, ""))
        {
        }

        private DataUiViews(string name, Input<string> id, DataUiViewsState? state = null, CustomResourceOptions? options = null)
            : base("splunk:index/dataUiViews:DataUiViews", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing DataUiViews resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static DataUiViews Get(string name, Input<string> id, DataUiViewsState? state = null, CustomResourceOptions? options = null)
        {
            return new DataUiViews(name, id, state, options);
        }
    }

    public sealed class DataUiViewsArgs : global::Pulumi.ResourceArgs
    {
        [Input("acl")]
        public Input<Inputs.DataUiViewsAclArgs>? Acl { get; set; }

        /// <summary>
        /// Dashboard XML definition.
        /// </summary>
        [Input("eaiData", required: true)]
        public Input<string> EaiData { get; set; } = null!;

        /// <summary>
        /// Dashboard name.
        /// * `eai:data` - (Required) Dashboard XML definition.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public DataUiViewsArgs()
        {
        }
        public static new DataUiViewsArgs Empty => new DataUiViewsArgs();
    }

    public sealed class DataUiViewsState : global::Pulumi.ResourceArgs
    {
        [Input("acl")]
        public Input<Inputs.DataUiViewsAclGetArgs>? Acl { get; set; }

        /// <summary>
        /// Dashboard XML definition.
        /// </summary>
        [Input("eaiData")]
        public Input<string>? EaiData { get; set; }

        /// <summary>
        /// Dashboard name.
        /// * `eai:data` - (Required) Dashboard XML definition.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public DataUiViewsState()
        {
        }
        public static new DataUiViewsState Empty => new DataUiViewsState();
    }
}
