// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Logic.V20160601.Outputs
{

    [OutputType]
    public sealed class RosettaNetPipRoleSettingsResponseResult
    {
        /// <summary>
        /// The action name.
        /// </summary>
        public readonly string Action;
        /// <summary>
        /// The RosettaNet ProcessConfiguration business document.
        /// </summary>
        public readonly Outputs.RosettaNetPipBusinessDocumentResponseResult BusinessDocument;
        /// <summary>
        /// The description.
        /// </summary>
        public readonly string? Description;
        /// <summary>
        /// The role name.
        /// </summary>
        public readonly string Role;
        /// <summary>
        /// The RosettaNet ProcessConfiguration role type.
        /// </summary>
        public readonly string RoleType;
        /// <summary>
        /// The service name.
        /// </summary>
        public readonly string Service;
        /// <summary>
        /// The service classification name.
        /// </summary>
        public readonly string ServiceClassification;

        [OutputConstructor]
        private RosettaNetPipRoleSettingsResponseResult(
            string action,

            Outputs.RosettaNetPipBusinessDocumentResponseResult businessDocument,

            string? description,

            string role,

            string roleType,

            string service,

            string serviceClassification)
        {
            Action = action;
            BusinessDocument = businessDocument;
            Description = description;
            Role = role;
            RoleType = roleType;
            Service = service;
            ServiceClassification = serviceClassification;
        }
    }
}