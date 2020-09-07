// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.BotService.V20171201.Outputs
{

    [OutputType]
    public sealed class WebChatChannelResponseResult
    {
        /// <summary>
        /// The channel name
        /// </summary>
        public readonly string ChannelName;
        /// <summary>
        /// The set of properties specific to Web Chat channel resource
        /// </summary>
        public readonly Outputs.WebChatChannelPropertiesResponseResult? Properties;

        [OutputConstructor]
        private WebChatChannelResponseResult(
            string channelName,

            Outputs.WebChatChannelPropertiesResponseResult? properties)
        {
            ChannelName = channelName;
            Properties = properties;
        }
    }
}