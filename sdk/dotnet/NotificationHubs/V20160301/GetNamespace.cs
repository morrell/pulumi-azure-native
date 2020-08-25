// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.NotificationHubs.V20160301
{
    public static class GetNamespace
    {
        public static Task<GetNamespaceResult> InvokeAsync(GetNamespaceArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetNamespaceResult>("azurerm:notificationhubs/v20160301:getNamespace", args ?? new GetNamespaceArgs(), options.WithVersion());
    }


    public sealed class GetNamespaceArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The namespace name.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// The name of the resource group.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetNamespaceArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetNamespaceResult
    {
        /// <summary>
        /// The time the namespace was created.
        /// </summary>
        public readonly string? CreatedAt;
        /// <summary>
        /// Whether or not the namespace is set as Critical.
        /// </summary>
        public readonly bool? Critical;
        /// <summary>
        /// Whether or not the namespace is currently enabled.
        /// </summary>
        public readonly bool? Enabled;
        /// <summary>
        /// Resource location
        /// </summary>
        public readonly string Location;
        /// <summary>
        /// Resource name
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The namespace type.
        /// </summary>
        public readonly string? NamespaceType;
        /// <summary>
        /// Provisioning state of the Namespace.
        /// </summary>
        public readonly string? ProvisioningState;
        /// <summary>
        /// Specifies the targeted region in which the namespace should be created. It can be any of the following values: Australia East, Australia Southeast, Central US, East US, East US 2, West US, North Central US, South Central US, East Asia, Southeast Asia, Brazil South, Japan East, Japan West, North Europe, West Europe
        /// </summary>
        public readonly string? Region;
        /// <summary>
        /// ScaleUnit where the namespace gets created
        /// </summary>
        public readonly string? ScaleUnit;
        /// <summary>
        /// Endpoint you can use to perform NotificationHub operations.
        /// </summary>
        public readonly string? ServiceBusEndpoint;
        /// <summary>
        /// The sku of the created namespace
        /// </summary>
        public readonly Outputs.SkuResponseResult? Sku;
        /// <summary>
        /// Status of the namespace. It can be any of these values:1 = Created/Active2 = Creating3 = Suspended4 = Deleting
        /// </summary>
        public readonly string? Status;
        /// <summary>
        /// The Id of the Azure subscription associated with the namespace.
        /// </summary>
        public readonly string? SubscriptionId;
        /// <summary>
        /// Resource tags
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Tags;
        /// <summary>
        /// Resource type
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetNamespaceResult(
            string? createdAt,

            bool? critical,

            bool? enabled,

            string location,

            string name,

            string? namespaceType,

            string? provisioningState,

            string? region,

            string? scaleUnit,

            string? serviceBusEndpoint,

            Outputs.SkuResponseResult? sku,

            string? status,

            string? subscriptionId,

            ImmutableDictionary<string, string>? tags,

            string type)
        {
            CreatedAt = createdAt;
            Critical = critical;
            Enabled = enabled;
            Location = location;
            Name = name;
            NamespaceType = namespaceType;
            ProvisioningState = provisioningState;
            Region = region;
            ScaleUnit = scaleUnit;
            ServiceBusEndpoint = serviceBusEndpoint;
            Sku = sku;
            Status = status;
            SubscriptionId = subscriptionId;
            Tags = tags;
            Type = type;
        }
    }
}