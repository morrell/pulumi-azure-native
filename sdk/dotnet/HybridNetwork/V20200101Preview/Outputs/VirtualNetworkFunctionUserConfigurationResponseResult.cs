// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.HybridNetwork.V20200101Preview.Outputs
{

    [OutputType]
    public sealed class VirtualNetworkFunctionUserConfigurationResponseResult
    {
        /// <summary>
        /// The network interface configuration.
        /// </summary>
        public readonly ImmutableArray<Outputs.NetworkInterfaceResponseResult> NetworkInterfaces;
        /// <summary>
        /// The name of the virtual network function role.
        /// </summary>
        public readonly string? RoleName;
        /// <summary>
        /// The user data parameters from the customer.
        /// </summary>
        public readonly ImmutableDictionary<string, object>? UserDataParameters;

        [OutputConstructor]
        private VirtualNetworkFunctionUserConfigurationResponseResult(
            ImmutableArray<Outputs.NetworkInterfaceResponseResult> networkInterfaces,

            string? roleName,

            ImmutableDictionary<string, object>? userDataParameters)
        {
            NetworkInterfaces = networkInterfaces;
            RoleName = roleName;
            UserDataParameters = userDataParameters;
        }
    }
}