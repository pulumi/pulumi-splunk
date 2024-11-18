// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Splunk.Outputs
{

    [OutputType]
    public sealed class AppsLocalAcl
    {
        /// <summary>
        /// The app context for the resource. Required for updating saved search ACL properties. Allowed values are:The name of an app and system
        /// </summary>
        public readonly string? App;
        /// <summary>
        /// Indicates if the active user can change permissions for this object. Defaults to true.
        /// </summary>
        public readonly bool? CanChangePerms;
        /// <summary>
        /// Indicates if the active user can change sharing to app level. Defaults to true.
        /// </summary>
        public readonly bool? CanShareApp;
        /// <summary>
        /// Indicates if the active user can change sharing to system level. Defaults to true.
        /// </summary>
        public readonly bool? CanShareGlobal;
        /// <summary>
        /// Indicates if the active user can change sharing to user level. Defaults to true.
        /// </summary>
        public readonly bool? CanShareUser;
        /// <summary>
        /// Indicates if the active user can edit this object. Defaults to true.
        /// </summary>
        public readonly bool? CanWrite;
        /// <summary>
        /// User name of resource owner. Defaults to the resource creator. Required for updating any knowledge object ACL properties.nobody = All users may access the resource, but write access to the resource might be restricted.
        /// </summary>
        public readonly string? Owner;
        /// <summary>
        /// Properties that indicate resource read permissions.
        /// </summary>
        public readonly ImmutableArray<string> Reads;
        /// <summary>
        /// Indicates whether an admin or user with sufficient permissions can delete the entity.
        /// </summary>
        public readonly bool? Removable;
        /// <summary>
        /// Indicates how the resource is shared. Required for updating any knowledge object ACL properties.app: Shared within a specific appglobal: (Default) Shared globally to all apps.user: Private to a user
        /// </summary>
        public readonly string? Sharing;
        /// <summary>
        /// Properties that indicate resource write permissions.
        /// </summary>
        public readonly ImmutableArray<string> Writes;

        [OutputConstructor]
        private AppsLocalAcl(
            string? app,

            bool? canChangePerms,

            bool? canShareApp,

            bool? canShareGlobal,

            bool? canShareUser,

            bool? canWrite,

            string? owner,

            ImmutableArray<string> reads,

            bool? removable,

            string? sharing,

            ImmutableArray<string> writes)
        {
            App = app;
            CanChangePerms = canChangePerms;
            CanShareApp = canShareApp;
            CanShareGlobal = canShareGlobal;
            CanShareUser = canShareUser;
            CanWrite = canWrite;
            Owner = owner;
            Reads = reads;
            Removable = removable;
            Sharing = sharing;
            Writes = writes;
        }
    }
}
