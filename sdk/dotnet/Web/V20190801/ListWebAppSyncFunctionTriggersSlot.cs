// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Web.V20190801
{
    public static class ListWebAppSyncFunctionTriggersSlot
    {
        public static Task<ListWebAppSyncFunctionTriggersSlotResult> InvokeAsync(ListWebAppSyncFunctionTriggersSlotArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<ListWebAppSyncFunctionTriggersSlotResult>("azurerm:web/v20190801:listWebAppSyncFunctionTriggersSlot", args ?? new ListWebAppSyncFunctionTriggersSlotArgs(), options.WithVersion());
    }


    public sealed class ListWebAppSyncFunctionTriggersSlotArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Name of the deployment slot.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// Name of the resource group to which the resource belongs.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public ListWebAppSyncFunctionTriggersSlotArgs()
        {
        }
    }


    [OutputType]
    public sealed class ListWebAppSyncFunctionTriggersSlotResult
    {
        /// <summary>
        /// Secret key.
        /// </summary>
        public readonly string? Key;
        /// <summary>
        /// Kind of resource.
        /// </summary>
        public readonly string? Kind;
        /// <summary>
        /// Resource Name.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Trigger URL.
        /// </summary>
        public readonly string? TriggerUrl;
        /// <summary>
        /// Resource type.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private ListWebAppSyncFunctionTriggersSlotResult(
            string? key,

            string? kind,

            string name,

            string? triggerUrl,

            string type)
        {
            Key = key;
            Kind = kind;
            Name = name;
            TriggerUrl = triggerUrl;
            Type = type;
        }
    }
}