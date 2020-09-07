// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Network.V20160330.Outputs
{

    [OutputType]
    public sealed class ApplicationGatewayRequestRoutingRuleResponseResult
    {
        /// <summary>
        /// Gets or sets backend address pool resource of application gateway 
        /// </summary>
        public readonly Outputs.SubResourceResponseResult? BackendAddressPool;
        /// <summary>
        /// Gets or sets frontend port resource of application gateway 
        /// </summary>
        public readonly Outputs.SubResourceResponseResult? BackendHttpSettings;
        /// <summary>
        /// A unique read-only string that changes whenever the resource is updated
        /// </summary>
        public readonly string? Etag;
        /// <summary>
        /// Gets or sets http listener resource of application gateway 
        /// </summary>
        public readonly Outputs.SubResourceResponseResult? HttpListener;
        /// <summary>
        /// Resource Id
        /// </summary>
        public readonly string? Id;
        /// <summary>
        /// Gets name of the resource that is unique within a resource group. This name can be used to access the resource
        /// </summary>
        public readonly string? Name;
        /// <summary>
        /// Gets or sets Provisioning state of the request routing rule resource Updating/Deleting/Failed
        /// </summary>
        public readonly string? ProvisioningState;
        /// <summary>
        /// Gets or sets the rule type
        /// </summary>
        public readonly string? RuleType;
        /// <summary>
        /// Gets or sets url path map resource of application gateway 
        /// </summary>
        public readonly Outputs.SubResourceResponseResult? UrlPathMap;

        [OutputConstructor]
        private ApplicationGatewayRequestRoutingRuleResponseResult(
            Outputs.SubResourceResponseResult? backendAddressPool,

            Outputs.SubResourceResponseResult? backendHttpSettings,

            string? etag,

            Outputs.SubResourceResponseResult? httpListener,

            string? id,

            string? name,

            string? provisioningState,

            string? ruleType,

            Outputs.SubResourceResponseResult? urlPathMap)
        {
            BackendAddressPool = backendAddressPool;
            BackendHttpSettings = backendHttpSettings;
            Etag = etag;
            HttpListener = httpListener;
            Id = id;
            Name = name;
            ProvisioningState = provisioningState;
            RuleType = ruleType;
            UrlPathMap = urlPathMap;
        }
    }
}