// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.MachineLearningCompute.V20170801Preview.Outputs
{

    [OutputType]
    public sealed class KubernetesClusterPropertiesResponseResult
    {
        /// <summary>
        /// The Azure Service Principal used by Kubernetes
        /// </summary>
        public readonly Outputs.ServicePrincipalPropertiesResponseResult? ServicePrincipal;

        [OutputConstructor]
        private KubernetesClusterPropertiesResponseResult(Outputs.ServicePrincipalPropertiesResponseResult? servicePrincipal)
        {
            ServicePrincipal = servicePrincipal;
        }
    }
}