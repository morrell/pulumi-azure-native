// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Insights.Latest
{
    public static class GetWorkbook
    {
        public static Task<GetWorkbookResult> InvokeAsync(GetWorkbookArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetWorkbookResult>("azurerm:insights/latest:getWorkbook", args ?? new GetWorkbookArgs(), options.WithVersion());
    }


    public sealed class GetWorkbookArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The name of the Application Insights component resource.
        /// </summary>
        [Input("resourceName", required: true)]
        public string ResourceName { get; set; } = null!;

        public GetWorkbookArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetWorkbookResult
    {
        /// <summary>
        /// Workbook category, as defined by the user at creation time.
        /// </summary>
        public readonly string Category;
        /// <summary>
        /// The kind of workbook. Choices are user and shared.
        /// </summary>
        public readonly string? Kind;
        /// <summary>
        /// Resource location
        /// </summary>
        public readonly string? Location;
        /// <summary>
        /// Azure resource name
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Configuration of this particular workbook. Configuration data is a string containing valid JSON
        /// </summary>
        public readonly string SerializedData;
        /// <summary>
        /// Enum indicating if this workbook definition is owned by a specific user or is shared between all users with access to the Application Insights component.
        /// </summary>
        public readonly string SharedTypeKind;
        /// <summary>
        /// Optional resourceId for a source resource.
        /// </summary>
        public readonly string? SourceResourceId;
        /// <summary>
        /// Resource tags
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Tags;
        /// <summary>
        /// Date and time in UTC of the last modification that was made to this workbook definition.
        /// </summary>
        public readonly string TimeModified;
        /// <summary>
        /// Azure resource type
        /// </summary>
        public readonly string Type;
        /// <summary>
        /// Unique user id of the specific user that owns this workbook.
        /// </summary>
        public readonly string UserId;
        /// <summary>
        /// This instance's version of the data model. This can change as new features are added that can be marked workbook.
        /// </summary>
        public readonly string? Version;
        /// <summary>
        /// Internally assigned unique id of the workbook definition.
        /// </summary>
        public readonly string WorkbookId;

        [OutputConstructor]
        private GetWorkbookResult(
            string category,

            string? kind,

            string? location,

            string name,

            string serializedData,

            string sharedTypeKind,

            string? sourceResourceId,

            ImmutableDictionary<string, string>? tags,

            string timeModified,

            string type,

            string userId,

            string? version,

            string workbookId)
        {
            Category = category;
            Kind = kind;
            Location = location;
            Name = name;
            SerializedData = serializedData;
            SharedTypeKind = sharedTypeKind;
            SourceResourceId = sourceResourceId;
            Tags = tags;
            TimeModified = timeModified;
            Type = type;
            UserId = userId;
            Version = version;
            WorkbookId = workbookId;
        }
    }
}