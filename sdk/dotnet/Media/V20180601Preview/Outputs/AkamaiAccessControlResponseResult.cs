// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Media.V20180601Preview.Outputs
{

    [OutputType]
    public sealed class AkamaiAccessControlResponseResult
    {
        /// <summary>
        /// authentication key list
        /// </summary>
        public readonly ImmutableArray<Outputs.AkamaiSignatureHeaderAuthenticationKeyResponseResult> AkamaiSignatureHeaderAuthenticationKeyList;

        [OutputConstructor]
        private AkamaiAccessControlResponseResult(ImmutableArray<Outputs.AkamaiSignatureHeaderAuthenticationKeyResponseResult> akamaiSignatureHeaderAuthenticationKeyList)
        {
            AkamaiSignatureHeaderAuthenticationKeyList = akamaiSignatureHeaderAuthenticationKeyList;
        }
    }
}