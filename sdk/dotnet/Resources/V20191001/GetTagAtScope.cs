// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Resources.V20191001
{
    public static class GetTagAtScope
    {
        public static Task<GetTagAtScopeResult> InvokeAsync(GetTagAtScopeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetTagAtScopeResult>("azurerm:resources/v20191001:getTagAtScope", args ?? new GetTagAtScopeArgs(), options.WithVersion());
    }


    public sealed class GetTagAtScopeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The resource scope.
        /// </summary>
        [Input("scope", required: true)]
        public string Scope { get; set; } = null!;

        public GetTagAtScopeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetTagAtScopeResult
    {
        /// <summary>
        /// The name of the tags wrapper resource.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The set of tags.
        /// </summary>
        public readonly Outputs.TagsResponseResult Properties;
        /// <summary>
        /// The type of the tags wrapper resource.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetTagAtScopeResult(
            string name,

            Outputs.TagsResponseResult properties,

            string type)
        {
            Name = name;
            Properties = properties;
            Type = type;
        }
    }
}