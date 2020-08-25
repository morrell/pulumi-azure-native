// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Logic.V20160601.Outputs
{

    [OutputType]
    public sealed class EdifactProcessingSettingsResponseResult
    {
        /// <summary>
        /// The value indicating whether to create empty xml tags for trailing separators.
        /// </summary>
        public readonly bool CreateEmptyXmlTagsForTrailingSeparators;
        /// <summary>
        /// The value indicating whether to mask security information.
        /// </summary>
        public readonly bool MaskSecurityInfo;
        /// <summary>
        /// The value indicating whether to preserve interchange.
        /// </summary>
        public readonly bool PreserveInterchange;
        /// <summary>
        /// The value indicating whether to suspend interchange on error.
        /// </summary>
        public readonly bool SuspendInterchangeOnError;
        /// <summary>
        /// The value indicating whether to use dot as decimal separator.
        /// </summary>
        public readonly bool UseDotAsDecimalSeparator;

        [OutputConstructor]
        private EdifactProcessingSettingsResponseResult(
            bool createEmptyXmlTagsForTrailingSeparators,

            bool maskSecurityInfo,

            bool preserveInterchange,

            bool suspendInterchangeOnError,

            bool useDotAsDecimalSeparator)
        {
            CreateEmptyXmlTagsForTrailingSeparators = createEmptyXmlTagsForTrailingSeparators;
            MaskSecurityInfo = maskSecurityInfo;
            PreserveInterchange = preserveInterchange;
            SuspendInterchangeOnError = suspendInterchangeOnError;
            UseDotAsDecimalSeparator = useDotAsDecimalSeparator;
        }
    }
}