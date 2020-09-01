// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.SoftwarePlan.Latest
{
    /// <summary>
    /// Response on GET of a hybrid use benefit
    /// </summary>
    public partial class HybridUseBenefit : Pulumi.CustomResource
    {
        /// <summary>
        /// Created date
        /// </summary>
        [Output("createdDate")]
        public Output<string> CreatedDate { get; private set; } = null!;

        /// <summary>
        /// Indicates the revision of the hybrid use benefit
        /// </summary>
        [Output("etag")]
        public Output<int> Etag { get; private set; } = null!;

        /// <summary>
        /// Last updated date
        /// </summary>
        [Output("lastUpdatedDate")]
        public Output<string> LastUpdatedDate { get; private set; } = null!;

        /// <summary>
        /// The name of the resource
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Provisioning state
        /// </summary>
        [Output("provisioningState")]
        public Output<string> ProvisioningState { get; private set; } = null!;

        /// <summary>
        /// Hybrid use benefit SKU
        /// </summary>
        [Output("sku")]
        public Output<Outputs.SkuResponseResult> Sku { get; private set; } = null!;

        /// <summary>
        /// The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a HybridUseBenefit resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public HybridUseBenefit(string name, HybridUseBenefitArgs args, CustomResourceOptions? options = null)
            : base("azurerm:softwareplan/latest:HybridUseBenefit", name, args ?? new HybridUseBenefitArgs(), MakeResourceOptions(options, ""))
        {
        }

        private HybridUseBenefit(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azurerm:softwareplan/latest:HybridUseBenefit", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new Pulumi.Alias { Type = "azurerm:softwareplan/v20191201:HybridUseBenefit"},
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing HybridUseBenefit resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static HybridUseBenefit Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new HybridUseBenefit(name, id, options);
        }
    }

    public sealed class HybridUseBenefitArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// This is a unique identifier for a plan. Should be a guid.
        /// </summary>
        [Input("planId", required: true)]
        public Input<string> PlanId { get; set; } = null!;

        /// <summary>
        /// The scope at which the operation is performed. This is limited to Microsoft.Compute/virtualMachines and Microsoft.Compute/hostGroups/hosts for now
        /// </summary>
        [Input("scope", required: true)]
        public Input<string> Scope { get; set; } = null!;

        /// <summary>
        /// Hybrid use benefit SKU
        /// </summary>
        [Input("sku", required: true)]
        public Input<Inputs.SkuArgs> Sku { get; set; } = null!;

        public HybridUseBenefitArgs()
        {
        }
    }
}