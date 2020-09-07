// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Blueprint.V20171111Preview.Inputs
{

    /// <summary>
    /// Managed Service Identity
    /// </summary>
    public sealed class ManagedServiceIdentityArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Azure Active Directory principal ID associated with this Identity.
        /// </summary>
        [Input("principalId")]
        public Input<string>? PrincipalId { get; set; }

        /// <summary>
        /// ID of the Azure Active Directory.
        /// </summary>
        [Input("tenantId")]
        public Input<string>? TenantId { get; set; }

        /// <summary>
        /// Type of the Managed Service Identity.
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        public ManagedServiceIdentityArgs()
        {
        }
    }
}