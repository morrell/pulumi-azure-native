// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.DomainRegistration.V20150801
{
    public static class ListTopLevelDomainAgreements
    {
        public static Task<ListTopLevelDomainAgreementsResult> InvokeAsync(ListTopLevelDomainAgreementsArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<ListTopLevelDomainAgreementsResult>("azurerm:domainregistration/v20150801:listTopLevelDomainAgreements", args ?? new ListTopLevelDomainAgreementsArgs(), options.WithVersion());
    }


    public sealed class ListTopLevelDomainAgreementsArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// If true then the list of agreements will include agreements for domain privacy as well.
        /// </summary>
        [Input("includePrivacy")]
        public bool? IncludePrivacy { get; set; }

        /// <summary>
        /// Name of the top level domain
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        public ListTopLevelDomainAgreementsArgs()
        {
        }
    }


    [OutputType]
    public sealed class ListTopLevelDomainAgreementsResult
    {
        /// <summary>
        /// Link to next page of resources
        /// </summary>
        public readonly string? NextLink;
        /// <summary>
        /// Collection of resources
        /// </summary>
        public readonly ImmutableArray<Outputs.TldLegalAgreementResponseResult> Value;

        [OutputConstructor]
        private ListTopLevelDomainAgreementsResult(
            string? nextLink,

            ImmutableArray<Outputs.TldLegalAgreementResponseResult> value)
        {
            NextLink = nextLink;
            Value = value;
        }
    }
}