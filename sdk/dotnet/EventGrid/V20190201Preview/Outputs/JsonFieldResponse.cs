// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.EventGrid.V20190201Preview.Outputs
{

    [OutputType]
    public sealed class JsonFieldResponse
    {
        /// <summary>
        /// Name of a field in the input event schema that's to be used as the source of a mapping.
        /// </summary>
        public readonly string? SourceField;

        [OutputConstructor]
        private JsonFieldResponse(string? sourceField)
        {
            SourceField = sourceField;
        }
    }
}
