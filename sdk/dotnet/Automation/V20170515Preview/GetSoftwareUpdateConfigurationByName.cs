// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Automation.V20170515Preview
{
    public static class GetSoftwareUpdateConfigurationByName
    {
        public static Task<GetSoftwareUpdateConfigurationByNameResult> InvokeAsync(GetSoftwareUpdateConfigurationByNameArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetSoftwareUpdateConfigurationByNameResult>("azurerm:automation/v20170515preview:getSoftwareUpdateConfigurationByName", args ?? new GetSoftwareUpdateConfigurationByNameArgs(), options.WithVersion());
    }


    public sealed class GetSoftwareUpdateConfigurationByNameArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the automation account.
        /// </summary>
        [Input("automationAccountName", required: true)]
        public string AutomationAccountName { get; set; } = null!;

        /// <summary>
        /// Name of an Azure Resource group.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The name of the software update configuration to be created.
        /// </summary>
        [Input("softwareUpdateConfigurationName", required: true)]
        public string SoftwareUpdateConfigurationName { get; set; } = null!;

        public GetSoftwareUpdateConfigurationByNameArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetSoftwareUpdateConfigurationByNameResult
    {
        /// <summary>
        /// CreatedBy property, which only appears in the response.
        /// </summary>
        public readonly string CreatedBy;
        /// <summary>
        /// Creation time of the resource, which only appears in the response.
        /// </summary>
        public readonly string CreationTime;
        /// <summary>
        /// Details of provisioning error
        /// </summary>
        public readonly Outputs.ErrorResponseResponseResult? Error;
        /// <summary>
        /// LastModifiedBy property, which only appears in the response.
        /// </summary>
        public readonly string LastModifiedBy;
        /// <summary>
        /// Last time resource was modified, which only appears in the response.
        /// </summary>
        public readonly string LastModifiedTime;
        /// <summary>
        /// Resource name.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Provisioning state for the software update configuration, which only appears in the response.
        /// </summary>
        public readonly string ProvisioningState;
        /// <summary>
        /// Schedule information for the Software update configuration
        /// </summary>
        public readonly Outputs.SchedulePropertiesResponseResult ScheduleInfo;
        /// <summary>
        /// Tasks information for the Software update configuration.
        /// </summary>
        public readonly Outputs.SoftwareUpdateConfigurationTasksResponseResult? Tasks;
        /// <summary>
        /// Resource type
        /// </summary>
        public readonly string Type;
        /// <summary>
        /// update specific properties for the Software update configuration
        /// </summary>
        public readonly Outputs.UpdateConfigurationResponseResult UpdateConfiguration;

        [OutputConstructor]
        private GetSoftwareUpdateConfigurationByNameResult(
            string createdBy,

            string creationTime,

            Outputs.ErrorResponseResponseResult? error,

            string lastModifiedBy,

            string lastModifiedTime,

            string name,

            string provisioningState,

            Outputs.SchedulePropertiesResponseResult scheduleInfo,

            Outputs.SoftwareUpdateConfigurationTasksResponseResult? tasks,

            string type,

            Outputs.UpdateConfigurationResponseResult updateConfiguration)
        {
            CreatedBy = createdBy;
            CreationTime = creationTime;
            Error = error;
            LastModifiedBy = lastModifiedBy;
            LastModifiedTime = lastModifiedTime;
            Name = name;
            ProvisioningState = provisioningState;
            ScheduleInfo = scheduleInfo;
            Tasks = tasks;
            Type = type;
            UpdateConfiguration = updateConfiguration;
        }
    }
}