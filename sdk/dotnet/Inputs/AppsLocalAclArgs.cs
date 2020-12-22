// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Splunk.Inputs
{

    public sealed class AppsLocalAclArgs : Pulumi.ResourceArgs
    {
        [Input("app")]
        public Input<string>? App { get; set; }

        [Input("canChangePerms")]
        public Input<bool>? CanChangePerms { get; set; }

        [Input("canShareApp")]
        public Input<bool>? CanShareApp { get; set; }

        [Input("canShareGlobal")]
        public Input<bool>? CanShareGlobal { get; set; }

        [Input("canShareUser")]
        public Input<bool>? CanShareUser { get; set; }

        [Input("canWrite")]
        public Input<bool>? CanWrite { get; set; }

        [Input("owner")]
        public Input<string>? Owner { get; set; }

        [Input("reads")]
        private InputList<string>? _reads;
        public InputList<string> Reads
        {
            get => _reads ?? (_reads = new InputList<string>());
            set => _reads = value;
        }

        [Input("removable")]
        public Input<bool>? Removable { get; set; }

        [Input("sharing")]
        public Input<string>? Sharing { get; set; }

        [Input("writes")]
        private InputList<string>? _writes;
        public InputList<string> Writes
        {
            get => _writes ?? (_writes = new InputList<string>());
            set => _writes = value;
        }

        public AppsLocalAclArgs()
        {
        }
    }
}
