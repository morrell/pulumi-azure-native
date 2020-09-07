// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Web.Latest
{
    public static class GetSiteInstanceDeploymentSlot
    {
        public static Task<GetSiteInstanceDeploymentSlotResult> InvokeAsync(GetSiteInstanceDeploymentSlotArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetSiteInstanceDeploymentSlotResult>("azurerm:web/latest:getSiteInstanceDeploymentSlot", args ?? new GetSiteInstanceDeploymentSlotArgs(), options.WithVersion());
    }


    public sealed class GetSiteInstanceDeploymentSlotArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Id of the deployment
        /// </summary>
        [Input("id", required: true)]
        public string Id { get; set; } = null!;

        /// <summary>
        /// Id of web app instance
        /// </summary>
        [Input("instanceId", required: true)]
        public string InstanceId { get; set; } = null!;

        /// <summary>
        /// Name of web app
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// Name of resource group
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// Name of web app slot. If not specified then will default to production slot.
        /// </summary>
        [Input("slot", required: true)]
        public string Slot { get; set; } = null!;

        public GetSiteInstanceDeploymentSlotArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetSiteInstanceDeploymentSlotResult
    {
        /// <summary>
        /// Active
        /// </summary>
        public readonly bool? Active;
        /// <summary>
        /// Author
        /// </summary>
        public readonly string? Author;
        /// <summary>
        /// AuthorEmail
        /// </summary>
        public readonly string? AuthorEmail;
        /// <summary>
        /// Deployer
        /// </summary>
        public readonly string? Deployer;
        /// <summary>
        /// Detail
        /// </summary>
        public readonly string? Details;
        /// <summary>
        /// EndTime
        /// </summary>
        public readonly string? EndTime;
        /// <summary>
        /// Kind of resource
        /// </summary>
        public readonly string? Kind;
        /// <summary>
        /// Resource Location
        /// </summary>
        public readonly string Location;
        /// <summary>
        /// Message
        /// </summary>
        public readonly string? Message;
        /// <summary>
        /// Resource Name
        /// </summary>
        public readonly string? Name;
        /// <summary>
        /// StartTime
        /// </summary>
        public readonly string? StartTime;
        /// <summary>
        /// Status
        /// </summary>
        public readonly int? Status;
        /// <summary>
        /// Resource tags
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Tags;
        /// <summary>
        /// Resource type
        /// </summary>
        public readonly string? Type;

        [OutputConstructor]
        private GetSiteInstanceDeploymentSlotResult(
            bool? active,

            string? author,

            string? authorEmail,

            string? deployer,

            string? details,

            string? endTime,

            string? kind,

            string location,

            string? message,

            string? name,

            string? startTime,

            int? status,

            ImmutableDictionary<string, string>? tags,

            string? type)
        {
            Active = active;
            Author = author;
            AuthorEmail = authorEmail;
            Deployer = deployer;
            Details = details;
            EndTime = endTime;
            Kind = kind;
            Location = location;
            Message = message;
            Name = name;
            StartTime = startTime;
            Status = status;
            Tags = tags;
            Type = type;
        }
    }
}