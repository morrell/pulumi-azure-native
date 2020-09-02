// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.SqlVirtualMachine.V20170301Preview.Inputs
{

    /// <summary>
    /// Configure your SQL virtual machine to be able to connect to the Azure Key Vault service.
    /// </summary>
    public sealed class KeyVaultCredentialSettingsArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Azure Key Vault url.
        /// </summary>
        [Input("azureKeyVaultUrl")]
        public Input<string>? AzureKeyVaultUrl { get; set; }

        /// <summary>
        /// Credential name.
        /// </summary>
        [Input("credentialName")]
        public Input<string>? CredentialName { get; set; }

        /// <summary>
        /// Enable or disable key vault credential setting.
        /// </summary>
        [Input("enable")]
        public Input<bool>? Enable { get; set; }

        /// <summary>
        /// Service principal name to access key vault.
        /// </summary>
        [Input("servicePrincipalName")]
        public Input<string>? ServicePrincipalName { get; set; }

        /// <summary>
        /// Service principal name secret to access key vault.
        /// </summary>
        [Input("servicePrincipalSecret")]
        public Input<string>? ServicePrincipalSecret { get; set; }

        public KeyVaultCredentialSettingsArgs()
        {
        }
    }
}