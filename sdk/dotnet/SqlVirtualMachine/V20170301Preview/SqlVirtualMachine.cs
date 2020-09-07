// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.SqlVirtualMachine.V20170301Preview
{
    /// <summary>
    /// A SQL virtual machine.
    /// </summary>
    public partial class SqlVirtualMachine : Pulumi.CustomResource
    {
        /// <summary>
        /// Auto backup settings for SQL Server.
        /// </summary>
        [Output("autoBackupSettings")]
        public Output<Outputs.AutoBackupSettingsResponseResult?> AutoBackupSettings { get; private set; } = null!;

        /// <summary>
        /// Auto patching settings for applying critical security updates to SQL virtual machine.
        /// </summary>
        [Output("autoPatchingSettings")]
        public Output<Outputs.AutoPatchingSettingsResponseResult?> AutoPatchingSettings { get; private set; } = null!;

        /// <summary>
        /// Azure Active Directory identity of the server.
        /// </summary>
        [Output("identity")]
        public Output<Outputs.ResourceIdentityResponseResult?> Identity { get; private set; } = null!;

        /// <summary>
        /// Key vault credential settings.
        /// </summary>
        [Output("keyVaultCredentialSettings")]
        public Output<Outputs.KeyVaultCredentialSettingsResponseResult?> KeyVaultCredentialSettings { get; private set; } = null!;

        /// <summary>
        /// Resource location.
        /// </summary>
        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// Resource name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Provisioning state to track the async operation status.
        /// </summary>
        [Output("provisioningState")]
        public Output<string> ProvisioningState { get; private set; } = null!;

        /// <summary>
        /// SQL Server configuration management settings.
        /// </summary>
        [Output("serverConfigurationsManagementSettings")]
        public Output<Outputs.ServerConfigurationsManagementSettingsResponseResult?> ServerConfigurationsManagementSettings { get; private set; } = null!;

        /// <summary>
        /// SQL image offer. Examples include SQL2016-WS2016, SQL2017-WS2016.
        /// </summary>
        [Output("sqlImageOffer")]
        public Output<string?> SqlImageOffer { get; private set; } = null!;

        /// <summary>
        /// SQL Server edition type.
        /// </summary>
        [Output("sqlImageSku")]
        public Output<string?> SqlImageSku { get; private set; } = null!;

        /// <summary>
        /// SQL Server Management type.
        /// </summary>
        [Output("sqlManagement")]
        public Output<string?> SqlManagement { get; private set; } = null!;

        /// <summary>
        /// SQL Server license type.
        /// </summary>
        [Output("sqlServerLicenseType")]
        public Output<string?> SqlServerLicenseType { get; private set; } = null!;

        /// <summary>
        /// ARM resource id of the SQL virtual machine group this SQL virtual machine is or will be part of.
        /// </summary>
        [Output("sqlVirtualMachineGroupResourceId")]
        public Output<string?> SqlVirtualMachineGroupResourceId { get; private set; } = null!;

        /// <summary>
        /// Storage Configuration Settings.
        /// </summary>
        [Output("storageConfigurationSettings")]
        public Output<Outputs.StorageConfigurationSettingsResponseResult?> StorageConfigurationSettings { get; private set; } = null!;

        /// <summary>
        /// Resource tags.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// Resource type.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;

        /// <summary>
        /// ARM Resource id of underlying virtual machine created from SQL marketplace image.
        /// </summary>
        [Output("virtualMachineResourceId")]
        public Output<string?> VirtualMachineResourceId { get; private set; } = null!;

        /// <summary>
        /// Domain credentials for setting up Windows Server Failover Cluster for SQL availability group.
        /// </summary>
        [Output("wsfcDomainCredentials")]
        public Output<Outputs.WsfcDomainCredentialsResponseResult?> WsfcDomainCredentials { get; private set; } = null!;


        /// <summary>
        /// Create a SqlVirtualMachine resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public SqlVirtualMachine(string name, SqlVirtualMachineArgs args, CustomResourceOptions? options = null)
            : base("azurerm:sqlvirtualmachine/v20170301preview:SqlVirtualMachine", name, args ?? new SqlVirtualMachineArgs(), MakeResourceOptions(options, ""))
        {
        }

        private SqlVirtualMachine(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azurerm:sqlvirtualmachine/v20170301preview:SqlVirtualMachine", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing SqlVirtualMachine resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static SqlVirtualMachine Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new SqlVirtualMachine(name, id, options);
        }
    }

    public sealed class SqlVirtualMachineArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Auto backup settings for SQL Server.
        /// </summary>
        [Input("autoBackupSettings")]
        public Input<Inputs.AutoBackupSettingsArgs>? AutoBackupSettings { get; set; }

        /// <summary>
        /// Auto patching settings for applying critical security updates to SQL virtual machine.
        /// </summary>
        [Input("autoPatchingSettings")]
        public Input<Inputs.AutoPatchingSettingsArgs>? AutoPatchingSettings { get; set; }

        /// <summary>
        /// Azure Active Directory identity of the server.
        /// </summary>
        [Input("identity")]
        public Input<Inputs.ResourceIdentityArgs>? Identity { get; set; }

        /// <summary>
        /// Key vault credential settings.
        /// </summary>
        [Input("keyVaultCredentialSettings")]
        public Input<Inputs.KeyVaultCredentialSettingsArgs>? KeyVaultCredentialSettings { get; set; }

        /// <summary>
        /// Resource location.
        /// </summary>
        [Input("location", required: true)]
        public Input<string> Location { get; set; } = null!;

        /// <summary>
        /// Name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// SQL Server configuration management settings.
        /// </summary>
        [Input("serverConfigurationsManagementSettings")]
        public Input<Inputs.ServerConfigurationsManagementSettingsArgs>? ServerConfigurationsManagementSettings { get; set; }

        /// <summary>
        /// SQL image offer. Examples include SQL2016-WS2016, SQL2017-WS2016.
        /// </summary>
        [Input("sqlImageOffer")]
        public Input<string>? SqlImageOffer { get; set; }

        /// <summary>
        /// SQL Server edition type.
        /// </summary>
        [Input("sqlImageSku")]
        public Input<string>? SqlImageSku { get; set; }

        /// <summary>
        /// SQL Server Management type.
        /// </summary>
        [Input("sqlManagement")]
        public Input<string>? SqlManagement { get; set; }

        /// <summary>
        /// SQL Server license type.
        /// </summary>
        [Input("sqlServerLicenseType")]
        public Input<string>? SqlServerLicenseType { get; set; }

        /// <summary>
        /// ARM resource id of the SQL virtual machine group this SQL virtual machine is or will be part of.
        /// </summary>
        [Input("sqlVirtualMachineGroupResourceId")]
        public Input<string>? SqlVirtualMachineGroupResourceId { get; set; }

        /// <summary>
        /// Name of the SQL virtual machine.
        /// </summary>
        [Input("sqlVirtualMachineName", required: true)]
        public Input<string> SqlVirtualMachineName { get; set; } = null!;

        /// <summary>
        /// Storage Configuration Settings.
        /// </summary>
        [Input("storageConfigurationSettings")]
        public Input<Inputs.StorageConfigurationSettingsArgs>? StorageConfigurationSettings { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// Resource tags.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        /// <summary>
        /// ARM Resource id of underlying virtual machine created from SQL marketplace image.
        /// </summary>
        [Input("virtualMachineResourceId")]
        public Input<string>? VirtualMachineResourceId { get; set; }

        /// <summary>
        /// Domain credentials for setting up Windows Server Failover Cluster for SQL availability group.
        /// </summary>
        [Input("wsfcDomainCredentials")]
        public Input<Inputs.WsfcDomainCredentialsArgs>? WsfcDomainCredentials { get; set; }

        public SqlVirtualMachineArgs()
        {
        }
    }
}