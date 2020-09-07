// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.EventGrid.V20200601.Outputs
{

    [OutputType]
    public sealed class JsonFieldWithDefaultResponseResult
    {
        /// <summary>
        /// The default value to be used for mapping when a SourceField is not provided or if there's no property with the specified name in the published JSON event payload.
        /// </summary>
        public readonly string? DefaultValue;
        /// <summary>
        /// Name of a field in the input event schema that's to be used as the source of a mapping.
        /// </summary>
        public readonly string? SourceField;

        [OutputConstructor]
        private JsonFieldWithDefaultResponseResult(
            string? defaultValue,

            string? sourceField)
        {
            DefaultValue = defaultValue;
            SourceField = sourceField;
        }
    }
}