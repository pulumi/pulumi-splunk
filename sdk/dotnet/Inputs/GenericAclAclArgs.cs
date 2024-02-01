// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Splunk.Inputs
{

    public sealed class GenericAclAclArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The app context for the resource. Required for updating saved search ACL properties. Allowed values are:The name of an app and system
        /// </summary>
        [Input("app")]
        public Input<string>? App { get; set; }

        /// <summary>
        /// Indicates if the active user can change permissions for this object. Defaults to true.
        /// </summary>
        [Input("canChangePerms")]
        public Input<bool>? CanChangePerms { get; set; }

        /// <summary>
        /// Indicates if the active user can change sharing to app level. Defaults to true.
        /// </summary>
        [Input("canShareApp")]
        public Input<bool>? CanShareApp { get; set; }

        /// <summary>
        /// Indicates if the active user can change sharing to system level. Defaults to true.
        /// </summary>
        [Input("canShareGlobal")]
        public Input<bool>? CanShareGlobal { get; set; }

        /// <summary>
        /// Indicates if the active user can change sharing to user level. Defaults to true.
        /// </summary>
        [Input("canShareUser")]
        public Input<bool>? CanShareUser { get; set; }

        /// <summary>
        /// Indicates if the active user can edit this object. Defaults to true.
        /// </summary>
        [Input("canWrite")]
        public Input<bool>? CanWrite { get; set; }

        /// <summary>
        /// User name of resource owner. Defaults to the resource creator. Required for updating any knowledge object ACL properties.nobody = All users may access the resource, but write access to the resource might be restricted.
        /// </summary>
        [Input("owner")]
        public Input<string>? Owner { get; set; }

        [Input("reads")]
        private InputList<string>? _reads;

        /// <summary>
        /// Properties that indicate resource read permissions.
        /// </summary>
        public InputList<string> Reads
        {
            get => _reads ?? (_reads = new InputList<string>());
            set => _reads = value;
        }

        /// <summary>
        /// Indicates whether an admin or user with sufficient permissions can delete the entity.
        /// </summary>
        [Input("removable")]
        public Input<bool>? Removable { get; set; }

        /// <summary>
        /// Indicates how the resource is shared. Required for updating any knowledge object ACL properties.app: Shared within a specific appglobal: (Default) Shared globally to all apps.user: Private to a user
        /// </summary>
        [Input("sharing")]
        public Input<string>? Sharing { get; set; }

        [Input("writes")]
        private InputList<string>? _writes;

        /// <summary>
        /// Properties that indicate write permissions of the resource.
        /// </summary>
        public InputList<string> Writes
        {
            get => _writes ?? (_writes = new InputList<string>());
            set => _writes = value;
        }

        public GenericAclAclArgs()
        {
        }
        public static new GenericAclAclArgs Empty => new GenericAclAclArgs();
    }
}
