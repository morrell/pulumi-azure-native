// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.AppPlatform.V20190501Preview.Outputs
{

    [OutputType]
    public sealed class ClusterResourcePropertiesResponseResult
    {
        /// <summary>
        /// Config server git properties of the Service
        /// </summary>
        public readonly Outputs.ConfigServerPropertiesResponseResult? ConfigServerProperties;
        /// <summary>
        /// Network profile of the Service
        /// </summary>
        public readonly Outputs.NetworkProfileResponseResult? NetworkProfile;
        /// <summary>
        /// Provisioning state of the Service
        /// </summary>
        public readonly string ProvisioningState;
        /// <summary>
        /// ServiceInstanceEntity GUID which uniquely identifies a created resource
        /// </summary>
        public readonly string ServiceId;
        /// <summary>
        /// Trace properties of the Service
        /// </summary>
        public readonly Outputs.TracePropertiesResponseResult? Trace;
        /// <summary>
        /// Version of the Service
        /// </summary>
        public readonly int Version;

        [OutputConstructor]
        private ClusterResourcePropertiesResponseResult(
            Outputs.ConfigServerPropertiesResponseResult? configServerProperties,

            Outputs.NetworkProfileResponseResult? networkProfile,

            string provisioningState,

            string serviceId,

            Outputs.TracePropertiesResponseResult? trace,

            int version)
        {
            ConfigServerProperties = configServerProperties;
            NetworkProfile = networkProfile;
            ProvisioningState = provisioningState;
            ServiceId = serviceId;
            Trace = trace;
            Version = version;
        }
    }
}