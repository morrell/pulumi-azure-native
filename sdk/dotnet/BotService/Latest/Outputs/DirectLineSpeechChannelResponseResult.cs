// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.BotService.Latest.Outputs
{

    [OutputType]
    public sealed class DirectLineSpeechChannelResponseResult
    {
        /// <summary>
        /// The channel name
        /// </summary>
        public readonly string ChannelName;
        /// <summary>
        /// The set of properties specific to DirectLine Speech channel resource
        /// </summary>
        public readonly Outputs.DirectLineSpeechChannelPropertiesResponseResult? Properties;

        [OutputConstructor]
        private DirectLineSpeechChannelResponseResult(
            string channelName,

            Outputs.DirectLineSpeechChannelPropertiesResponseResult? properties)
        {
            ChannelName = channelName;
            Properties = properties;
        }
    }
}