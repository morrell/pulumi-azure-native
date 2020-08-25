// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Web.V20180201.Inputs
{

    /// <summary>
    /// Information about the formal API definition for the app.
    /// </summary>
    public sealed class ApiDefinitionInfoArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The URL of the API definition.
        /// </summary>
        [Input("url")]
        public Input<string>? Url { get; set; }

        public ApiDefinitionInfoArgs()
        {
        }
    }
}