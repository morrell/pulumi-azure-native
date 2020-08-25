// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.ApiManagement.V20191201
{
    public static class ListTenantAccessGitSecrets
    {
        public static Task<ListTenantAccessGitSecretsResult> InvokeAsync(ListTenantAccessGitSecretsArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<ListTenantAccessGitSecretsResult>("azurerm:apimanagement/v20191201:listTenantAccessGitSecrets", args ?? new ListTenantAccessGitSecretsArgs(), options.WithVersion());
    }


    public sealed class ListTenantAccessGitSecretsArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The identifier of the Access configuration.
        /// </summary>
        [Input("accessName", required: true)]
        public string AccessName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The name of the API Management service.
        /// </summary>
        [Input("serviceName", required: true)]
        public string ServiceName { get; set; } = null!;

        public ListTenantAccessGitSecretsArgs()
        {
        }
    }


    [OutputType]
    public sealed class ListTenantAccessGitSecretsResult
    {
        /// <summary>
        /// Determines whether direct access is enabled.
        /// </summary>
        public readonly bool? Enabled;
        /// <summary>
        /// Primary access key. This property will not be filled on 'GET' operations! Use '/listSecrets' POST request to get the value.
        /// </summary>
        public readonly string? PrimaryKey;
        /// <summary>
        /// Secondary access key. This property will not be filled on 'GET' operations! Use '/listSecrets' POST request to get the value.
        /// </summary>
        public readonly string? SecondaryKey;

        [OutputConstructor]
        private ListTenantAccessGitSecretsResult(
            bool? enabled,

            string? primaryKey,

            string? secondaryKey)
        {
            Enabled = enabled;
            PrimaryKey = primaryKey;
            SecondaryKey = secondaryKey;
        }
    }
}