// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Media.Latest.Outputs
{

    [OutputType]
    public sealed class AbsoluteClipTimeResponseResult
    {
        /// <summary>
        /// The discriminator for derived types.
        /// </summary>
        public readonly string OdataType;
        /// <summary>
        /// The time position on the timeline of the input media. It is usually specified as an ISO8601 period. e.g PT30S for 30 seconds.
        /// </summary>
        public readonly string Time;

        [OutputConstructor]
        private AbsoluteClipTimeResponseResult(
            string odataType,

            string time)
        {
            OdataType = odataType;
            Time = time;
        }
    }
}