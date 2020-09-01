// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Automation.Latest
{
    public static class GetConnection
    {
        public static Task<GetConnectionResult> InvokeAsync(GetConnectionArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetConnectionResult>("azurerm:automation/latest:getConnection", args ?? new GetConnectionArgs(), options.WithVersion());
    }


    public sealed class GetConnectionArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the automation account.
        /// </summary>
        [Input("automationAccountName", required: true)]
        public string AutomationAccountName { get; set; } = null!;

        /// <summary>
        /// The name of connection.
        /// </summary>
        [Input("connectionName", required: true)]
        public string ConnectionName { get; set; } = null!;

        /// <summary>
        /// Name of an Azure Resource group.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetConnectionArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetConnectionResult
    {
        /// <summary>
        /// Gets or sets the connectionType of the connection.
        /// </summary>
        public readonly Outputs.ConnectionTypeAssociationPropertyResponseResult? ConnectionType;
        /// <summary>
        /// Gets the creation time.
        /// </summary>
        public readonly string CreationTime;
        /// <summary>
        /// Gets or sets the description.
        /// </summary>
        public readonly string? Description;
        /// <summary>
        /// Gets the field definition values of the connection.
        /// </summary>
        public readonly ImmutableDictionary<string, string> FieldDefinitionValues;
        /// <summary>
        /// Gets the last modified time.
        /// </summary>
        public readonly string LastModifiedTime;
        /// <summary>
        /// The name of the resource
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The type of the resource.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetConnectionResult(
            Outputs.ConnectionTypeAssociationPropertyResponseResult? connectionType,

            string creationTime,

            string? description,

            ImmutableDictionary<string, string> fieldDefinitionValues,

            string lastModifiedTime,

            string name,

            string type)
        {
            ConnectionType = connectionType;
            CreationTime = creationTime;
            Description = description;
            FieldDefinitionValues = fieldDefinitionValues;
            LastModifiedTime = lastModifiedTime;
            Name = name;
            Type = type;
        }
    }
}