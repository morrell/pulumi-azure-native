// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.RecoveryServices.V20180110.Inputs
{

    /// <summary>
    /// A2A container mapping input.
    /// </summary>
    public sealed class A2AContainerMappingInputArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// A value indicating whether the auto update is enabled.
        /// </summary>
        [Input("agentAutoUpdateStatus")]
        public Input<string>? AgentAutoUpdateStatus { get; set; }

        /// <summary>
        /// The automation account arm id.
        /// </summary>
        [Input("automationAccountArmId")]
        public Input<string>? AutomationAccountArmId { get; set; }

        /// <summary>
        /// The class type.
        /// </summary>
        [Input("instanceType")]
        public Input<string>? InstanceType { get; set; }

        public A2AContainerMappingInputArgs()
        {
        }
    }
}