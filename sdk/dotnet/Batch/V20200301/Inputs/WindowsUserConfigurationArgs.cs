// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Batch.V20200301.Inputs
{

    public sealed class WindowsUserConfigurationArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies login mode for the user. The default value for VirtualMachineConfiguration pools is interactive mode and for CloudServiceConfiguration pools is batch mode.
        /// </summary>
        [Input("loginMode")]
        public Input<string>? LoginMode { get; set; }

        public WindowsUserConfigurationArgs()
        {
        }
    }
}