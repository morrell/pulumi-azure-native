// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.DesktopVirtualization.V20210513Preview.Inputs
{

    /// <summary>
    /// Secondary maintenance windows is a list of days at one specific hour. Maintenance windows are 2 hours long. We try to exercise this only when the primary window update fails.
    /// </summary>
    public sealed class SecondaryWindowPropertiesArgs : Pulumi.ResourceArgs
    {
        [Input("daysOfWeek")]
        private InputList<string>? _daysOfWeek;

        /// <summary>
        /// Set of days of the week on which this schedule is active.
        /// </summary>
        public InputList<string> DaysOfWeek
        {
            get => _daysOfWeek ?? (_daysOfWeek = new InputList<string>());
            set => _daysOfWeek = value;
        }

        /// <summary>
        /// The update start hour of the day. (0 - 23)
        /// </summary>
        [Input("hour")]
        public Input<int>? Hour { get; set; }

        public SecondaryWindowPropertiesArgs()
        {
        }
    }
}
