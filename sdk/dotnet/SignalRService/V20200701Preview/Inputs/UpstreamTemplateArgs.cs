// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.SignalRService.V20200701Preview.Inputs
{

    /// <summary>
    /// Upstream template item settings. It defines the Upstream URL of the incoming requests.
    /// The template defines the pattern of the event, the hub or the category of the incoming request that matches current URL template.
    /// </summary>
    public sealed class UpstreamTemplateArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Gets or sets the auth settings for an upstream. If not set, no auth is used for upstream messages.
        /// </summary>
        [Input("auth")]
        public Input<Inputs.UpstreamAuthSettingsArgs>? Auth { get; set; }

        /// <summary>
        /// Gets or sets the matching pattern for category names. If not set, it matches any category.
        /// There are 3 kind of patterns supported:
        ///     1. "*", it to matches any category name
        ///     2. Combine multiple categories with ",", for example "connections,messages", it matches category "connections" and "messages"
        ///     3. The single category name, for example, "connections", it matches the category "connections"
        /// </summary>
        [Input("categoryPattern")]
        public Input<string>? CategoryPattern { get; set; }

        /// <summary>
        /// Gets or sets the matching pattern for event names. If not set, it matches any event.
        /// There are 3 kind of patterns supported:
        ///     1. "*", it to matches any event name
        ///     2. Combine multiple events with ",", for example "connect,disconnect", it matches event "connect" and "disconnect"
        ///     3. The single event name, for example, "connect", it matches "connect"
        /// </summary>
        [Input("eventPattern")]
        public Input<string>? EventPattern { get; set; }

        /// <summary>
        /// Gets or sets the matching pattern for hub names. If not set, it matches any hub.
        /// There are 3 kind of patterns supported:
        ///     1. "*", it to matches any hub name
        ///     2. Combine multiple hubs with ",", for example "hub1,hub2", it matches "hub1" and "hub2"
        ///     3. The single hub name, for example, "hub1", it matches "hub1"
        /// </summary>
        [Input("hubPattern")]
        public Input<string>? HubPattern { get; set; }

        /// <summary>
        /// Gets or sets the Upstream URL template. You can use 3 predefined parameters {hub}, {category} {event} inside the template, the value of the Upstream URL is dynamically calculated when the client request comes in.
        /// For example, if the urlTemplate is `http://example.com/{hub}/api/{event}`, with a client request from hub `chat` connects, it will first POST to this URL: `http://example.com/chat/api/connect`.
        /// </summary>
        [Input("urlTemplate", required: true)]
        public Input<string> UrlTemplate { get; set; } = null!;

        public UpstreamTemplateArgs()
        {
        }
    }
}