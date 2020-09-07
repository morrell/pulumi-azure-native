// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Insights.V20191017Preview.Outputs
{

    [OutputType]
    public sealed class WorkbookTemplateLocalizedGalleryResponseResult
    {
        /// <summary>
        /// Workbook galleries supported by the template.
        /// </summary>
        public readonly ImmutableArray<Outputs.WorkbookTemplateGalleryResponseResult> Galleries;
        /// <summary>
        /// Valid JSON object containing workbook template payload.
        /// </summary>
        public readonly ImmutableDictionary<string, object>? TemplateData;

        [OutputConstructor]
        private WorkbookTemplateLocalizedGalleryResponseResult(
            ImmutableArray<Outputs.WorkbookTemplateGalleryResponseResult> galleries,

            ImmutableDictionary<string, object>? templateData)
        {
            Galleries = galleries;
            TemplateData = templateData;
        }
    }
}