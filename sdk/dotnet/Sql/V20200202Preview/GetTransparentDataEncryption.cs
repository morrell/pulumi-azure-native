// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Sql.V20200202Preview
{
    public static class GetTransparentDataEncryption
    {
        /// <summary>
        /// A logical database transparent data encryption state.
        /// </summary>
        public static Task<GetTransparentDataEncryptionResult> InvokeAsync(GetTransparentDataEncryptionArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetTransparentDataEncryptionResult>("azure-native:sql/v20200202preview:getTransparentDataEncryption", args ?? new GetTransparentDataEncryptionArgs(), options.WithVersion());
    }


    public sealed class GetTransparentDataEncryptionArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the logical database for which the transparent data encryption is defined.
        /// </summary>
        [Input("databaseName", required: true)]
        public string DatabaseName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The name of the server.
        /// </summary>
        [Input("serverName", required: true)]
        public string ServerName { get; set; } = null!;

        /// <summary>
        /// The name of the transparent data encryption configuration.
        /// </summary>
        [Input("tdeName", required: true)]
        public string TdeName { get; set; } = null!;

        public GetTransparentDataEncryptionArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetTransparentDataEncryptionResult
    {
        /// <summary>
        /// Resource ID.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Resource name.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Specifies the state of the transparent data encryption.
        /// </summary>
        public readonly string State;
        /// <summary>
        /// Resource type.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetTransparentDataEncryptionResult(
            string id,

            string name,

            string state,

            string type)
        {
            Id = id;
            Name = name;
            State = state;
            Type = type;
        }
    }
}
