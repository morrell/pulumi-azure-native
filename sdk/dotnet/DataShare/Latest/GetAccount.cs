// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.DataShare.Latest
{
    public static class GetAccount
    {
        public static Task<GetAccountResult> InvokeAsync(GetAccountArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetAccountResult>("azurerm:datashare/latest:getAccount", args ?? new GetAccountArgs(), options.WithVersion());
    }


    public sealed class GetAccountArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the share account.
        /// </summary>
        [Input("accountName", required: true)]
        public string AccountName { get; set; } = null!;

        /// <summary>
        /// The resource group name.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetAccountArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetAccountResult
    {
        /// <summary>
        /// Time at which the account was created.
        /// </summary>
        public readonly string CreatedAt;
        /// <summary>
        /// Identity Info on the Account
        /// </summary>
        public readonly Outputs.IdentityResponseResult Identity;
        /// <summary>
        /// Location of the azure resource.
        /// </summary>
        public readonly string? Location;
        /// <summary>
        /// Name of the azure resource
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Provisioning state of the Account
        /// </summary>
        public readonly string ProvisioningState;
        /// <summary>
        /// Tags on the azure resource.
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Tags;
        /// <summary>
        /// Type of the azure resource
        /// </summary>
        public readonly string Type;
        /// <summary>
        /// Email of the user who created the resource
        /// </summary>
        public readonly string UserEmail;
        /// <summary>
        /// Name of the user who created the resource
        /// </summary>
        public readonly string UserName;

        [OutputConstructor]
        private GetAccountResult(
            string createdAt,

            Outputs.IdentityResponseResult identity,

            string? location,

            string name,

            string provisioningState,

            ImmutableDictionary<string, string>? tags,

            string type,

            string userEmail,

            string userName)
        {
            CreatedAt = createdAt;
            Identity = identity;
            Location = location;
            Name = name;
            ProvisioningState = provisioningState;
            Tags = tags;
            Type = type;
            UserEmail = userEmail;
            UserName = userName;
        }
    }
}