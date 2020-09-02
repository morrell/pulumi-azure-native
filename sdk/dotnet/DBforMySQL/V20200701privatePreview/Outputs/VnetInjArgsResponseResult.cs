// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.DBForMySql.V20200701privatePreview.Outputs
{

    [OutputType]
    public sealed class VnetInjArgsResponseResult
    {
        /// <summary>
        /// delegated subnet name
        /// </summary>
        public readonly string? DelegatedSubnetName;
        /// <summary>
        /// delegated vNet ID
        /// </summary>
        public readonly string? DelegatedVnetID;
        /// <summary>
        /// delegated vNet name
        /// </summary>
        public readonly string? DelegatedVnetName;

        [OutputConstructor]
        private VnetInjArgsResponseResult(
            string? delegatedSubnetName,

            string? delegatedVnetID,

            string? delegatedVnetName)
        {
            DelegatedSubnetName = delegatedSubnetName;
            DelegatedVnetID = delegatedVnetID;
            DelegatedVnetName = delegatedVnetName;
        }
    }
}