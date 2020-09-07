// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Compute.V20190301
{
    public static class GetVirtualMachineScaleSet
    {
        public static Task<GetVirtualMachineScaleSetResult> InvokeAsync(GetVirtualMachineScaleSetArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetVirtualMachineScaleSetResult>("azurerm:compute/v20190301:getVirtualMachineScaleSet", args ?? new GetVirtualMachineScaleSetArgs(), options.WithVersion());
    }


    public sealed class GetVirtualMachineScaleSetArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the resource group.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The name of the VM scale set.
        /// </summary>
        [Input("vmScaleSetName", required: true)]
        public string VmScaleSetName { get; set; } = null!;

        public GetVirtualMachineScaleSetArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetVirtualMachineScaleSetResult
    {
        /// <summary>
        /// Specifies additional capabilities enabled or disabled on the Virtual Machines in the Virtual Machine Scale Set. For instance: whether the Virtual Machines have the capability to support attaching managed data disks with UltraSSD_LRS storage account type.
        /// </summary>
        public readonly Outputs.AdditionalCapabilitiesResponseResult? AdditionalCapabilities;
        /// <summary>
        /// Policy for automatic repairs.
        /// </summary>
        public readonly Outputs.AutomaticRepairsPolicyResponseResult? AutomaticRepairsPolicy;
        /// <summary>
        /// When Overprovision is enabled, extensions are launched only on the requested number of VMs which are finally kept. This property will hence ensure that the extensions do not run on the extra overprovisioned VMs.
        /// </summary>
        public readonly bool? DoNotRunExtensionsOnOverprovisionedVMs;
        /// <summary>
        /// The identity of the virtual machine scale set, if configured.
        /// </summary>
        public readonly Outputs.VirtualMachineScaleSetIdentityResponseResult? Identity;
        /// <summary>
        /// Resource location
        /// </summary>
        public readonly string Location;
        /// <summary>
        /// Resource name
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Specifies whether the Virtual Machine Scale Set should be overprovisioned.
        /// </summary>
        public readonly bool? Overprovision;
        /// <summary>
        /// Specifies information about the marketplace image used to create the virtual machine. This element is only used for marketplace images. Before you can use a marketplace image from an API, you must enable the image for programmatic use.  In the Azure portal, find the marketplace image that you want to use and then click **Want to deploy programmatically, Get Started -&gt;**. Enter any required information and then click **Save**.
        /// </summary>
        public readonly Outputs.PlanResponseResult? Plan;
        /// <summary>
        /// Fault Domain count for each placement group.
        /// </summary>
        public readonly int? PlatformFaultDomainCount;
        /// <summary>
        /// The provisioning state, which only appears in the response.
        /// </summary>
        public readonly string ProvisioningState;
        /// <summary>
        /// Specifies information about the proximity placement group that the virtual machine scale set should be assigned to. &lt;br&gt;&lt;br&gt;Minimum api-version: 2018-04-01.
        /// </summary>
        public readonly Outputs.SubResourceResponseResult? ProximityPlacementGroup;
        /// <summary>
        /// Specifies the scale-in policy that decides which virtual machines are chosen for removal when a Virtual Machine Scale Set is scaled-in.
        /// </summary>
        public readonly Outputs.ScaleInPolicyResponseResult? ScaleInPolicy;
        /// <summary>
        /// When true this limits the scale set to a single placement group, of max size 100 virtual machines.
        /// </summary>
        public readonly bool? SinglePlacementGroup;
        /// <summary>
        /// The virtual machine scale set sku.
        /// </summary>
        public readonly Outputs.SkuResponseResult? Sku;
        /// <summary>
        /// Resource tags
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Tags;
        /// <summary>
        /// Resource type
        /// </summary>
        public readonly string Type;
        /// <summary>
        /// Specifies the ID which uniquely identifies a Virtual Machine Scale Set.
        /// </summary>
        public readonly string UniqueId;
        /// <summary>
        /// The upgrade policy.
        /// </summary>
        public readonly Outputs.UpgradePolicyResponseResult? UpgradePolicy;
        /// <summary>
        /// The virtual machine profile.
        /// </summary>
        public readonly Outputs.VirtualMachineScaleSetVMProfileResponseResult? VirtualMachineProfile;
        /// <summary>
        /// Whether to force strictly even Virtual Machine distribution cross x-zones in case there is zone outage.
        /// </summary>
        public readonly bool? ZoneBalance;
        /// <summary>
        /// The virtual machine scale set zones. NOTE: Availability zones can only be set when you create the scale set.
        /// </summary>
        public readonly ImmutableArray<string> Zones;

        [OutputConstructor]
        private GetVirtualMachineScaleSetResult(
            Outputs.AdditionalCapabilitiesResponseResult? additionalCapabilities,

            Outputs.AutomaticRepairsPolicyResponseResult? automaticRepairsPolicy,

            bool? doNotRunExtensionsOnOverprovisionedVMs,

            Outputs.VirtualMachineScaleSetIdentityResponseResult? identity,

            string location,

            string name,

            bool? overprovision,

            Outputs.PlanResponseResult? plan,

            int? platformFaultDomainCount,

            string provisioningState,

            Outputs.SubResourceResponseResult? proximityPlacementGroup,

            Outputs.ScaleInPolicyResponseResult? scaleInPolicy,

            bool? singlePlacementGroup,

            Outputs.SkuResponseResult? sku,

            ImmutableDictionary<string, string>? tags,

            string type,

            string uniqueId,

            Outputs.UpgradePolicyResponseResult? upgradePolicy,

            Outputs.VirtualMachineScaleSetVMProfileResponseResult? virtualMachineProfile,

            bool? zoneBalance,

            ImmutableArray<string> zones)
        {
            AdditionalCapabilities = additionalCapabilities;
            AutomaticRepairsPolicy = automaticRepairsPolicy;
            DoNotRunExtensionsOnOverprovisionedVMs = doNotRunExtensionsOnOverprovisionedVMs;
            Identity = identity;
            Location = location;
            Name = name;
            Overprovision = overprovision;
            Plan = plan;
            PlatformFaultDomainCount = platformFaultDomainCount;
            ProvisioningState = provisioningState;
            ProximityPlacementGroup = proximityPlacementGroup;
            ScaleInPolicy = scaleInPolicy;
            SinglePlacementGroup = singlePlacementGroup;
            Sku = sku;
            Tags = tags;
            Type = type;
            UniqueId = uniqueId;
            UpgradePolicy = upgradePolicy;
            VirtualMachineProfile = virtualMachineProfile;
            ZoneBalance = zoneBalance;
            Zones = zones;
        }
    }
}