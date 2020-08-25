// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Portal.V20181001
{
    public static class GetUserSettings
    {
        public static Task<GetUserSettingsResult> InvokeAsync(GetUserSettingsArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetUserSettingsResult>("azurerm:portal/v20181001:getUserSettings", args ?? new GetUserSettingsArgs(), options.WithVersion());
    }


    public sealed class GetUserSettingsArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the user settings
        /// </summary>
        [Input("userSettingsName", required: true)]
        public string UserSettingsName { get; set; } = null!;

        public GetUserSettingsArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetUserSettingsResult
    {
        /// <summary>
        /// The cloud shell user settings properties.
        /// </summary>
        public readonly Outputs.UserPropertiesResponseResult Properties;

        [OutputConstructor]
        private GetUserSettingsResult(Outputs.UserPropertiesResponseResult properties)
        {
            Properties = properties;
        }
    }
}