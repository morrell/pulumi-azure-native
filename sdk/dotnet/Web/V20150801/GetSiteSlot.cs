// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Web.V20150801
{
    public static class GetSiteSlot
    {
        public static Task<GetSiteSlotResult> InvokeAsync(GetSiteSlotArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetSiteSlotResult>("azurerm:web/v20150801:getSiteSlot", args ?? new GetSiteSlotArgs(), options.WithVersion());
    }


    public sealed class GetSiteSlotArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Name of web app
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// Additional web app properties included in the response
        /// </summary>
        [Input("propertiesToInclude")]
        public string? PropertiesToInclude { get; set; }

        /// <summary>
        /// Name of resource group
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// Name of web app slot. If not specified then will default to production slot.
        /// </summary>
        [Input("slot", required: true)]
        public string Slot { get; set; } = null!;

        public GetSiteSlotArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetSiteSlotResult
    {
        /// <summary>
        /// Management information availability state for the web app. Possible values are Normal or Limited. 
        ///             Normal means that the site is running correctly and that management information for the site is available. 
        ///             Limited means that only partial management information for the site is available and that detailed site information is unavailable.
        /// </summary>
        public readonly string AvailabilityState;
        /// <summary>
        /// Specifies if the client affinity is enabled when load balancing http request for multiple instances of the web app
        /// </summary>
        public readonly bool? ClientAffinityEnabled;
        /// <summary>
        /// Specifies if the client certificate is enabled for the web app
        /// </summary>
        public readonly bool? ClientCertEnabled;
        /// <summary>
        /// This is only valid for web app creation. If specified, web app is cloned from 
        ///             a source web app
        /// </summary>
        public readonly Outputs.CloningInfoResponseResult? CloningInfo;
        /// <summary>
        /// Size of a function container
        /// </summary>
        public readonly int? ContainerSize;
        /// <summary>
        /// Default hostname of the web app
        /// </summary>
        public readonly string DefaultHostName;
        /// <summary>
        /// True if the site is enabled; otherwise, false. Setting this  value to false disables the site (takes the site off line).
        /// </summary>
        public readonly bool? Enabled;
        /// <summary>
        /// Hostnames for the web app that are enabled. Hostnames need to be assigned and enabled. If some hostnames are assigned but not enabled
        ///             the app is not served on those hostnames
        /// </summary>
        public readonly ImmutableArray<string> EnabledHostNames;
        /// <summary>
        /// Name of gateway app associated with web app
        /// </summary>
        public readonly string? GatewaySiteName;
        /// <summary>
        /// Hostname SSL states are  used to manage the SSL bindings for site's hostnames.
        /// </summary>
        public readonly ImmutableArray<Outputs.HostNameSslStateResponseResult> HostNameSslStates;
        /// <summary>
        /// Hostnames associated with web app
        /// </summary>
        public readonly ImmutableArray<string> HostNames;
        /// <summary>
        /// Specifies if the public hostnames are disabled the web app.
        ///             If set to true the app is only accessible via API Management process
        /// </summary>
        public readonly bool? HostNamesDisabled;
        /// <summary>
        /// Specification for the hosting environment (App Service Environment) to use for the web app
        /// </summary>
        public readonly Outputs.HostingEnvironmentProfileResponseResult? HostingEnvironmentProfile;
        /// <summary>
        /// Site is a default container
        /// </summary>
        public readonly bool IsDefaultContainer;
        /// <summary>
        /// Kind of resource
        /// </summary>
        public readonly string? Kind;
        /// <summary>
        /// Last time web app was modified in UTC
        /// </summary>
        public readonly string LastModifiedTimeUtc;
        /// <summary>
        /// Resource Location
        /// </summary>
        public readonly string Location;
        /// <summary>
        /// Maximum number of workers
        ///             This only applies to function container
        /// </summary>
        public readonly int? MaxNumberOfWorkers;
        public readonly string? MicroService;
        /// <summary>
        /// Resource Name
        /// </summary>
        public readonly string? Name;
        /// <summary>
        /// List of comma separated IP addresses that this web app uses for outbound connections. Those can be used when configuring firewall rules for databases accessed by this web app.
        /// </summary>
        public readonly string OutboundIpAddresses;
        /// <summary>
        /// If set indicates whether web app is deployed as a premium app
        /// </summary>
        public readonly bool PremiumAppDeployed;
        /// <summary>
        /// Name of repository site
        /// </summary>
        public readonly string RepositorySiteName;
        /// <summary>
        /// Resource group web app belongs to
        /// </summary>
        public readonly string ResourceGroup;
        /// <summary>
        /// If set indicates whether to stop SCM (KUDU) site when the web app is stopped. Default is false.
        /// </summary>
        public readonly bool? ScmSiteAlsoStopped;
        public readonly string? ServerFarmId;
        /// <summary>
        /// Configuration of web app
        /// </summary>
        public readonly Outputs.SiteConfigResponseResult? SiteConfig;
        /// <summary>
        /// State of the web app
        /// </summary>
        public readonly string State;
        /// <summary>
        /// Resource tags
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Tags;
        /// <summary>
        /// Read-only property that specifies which slot this app will swap into
        /// </summary>
        public readonly string TargetSwapSlot;
        /// <summary>
        /// Read-only list of Azure Traffic manager hostnames associated with web app
        /// </summary>
        public readonly ImmutableArray<string> TrafficManagerHostNames;
        /// <summary>
        /// Resource type
        /// </summary>
        public readonly string? Type;
        /// <summary>
        /// State indicating whether web app has exceeded its quota usage
        /// </summary>
        public readonly string UsageState;

        [OutputConstructor]
        private GetSiteSlotResult(
            string availabilityState,

            bool? clientAffinityEnabled,

            bool? clientCertEnabled,

            Outputs.CloningInfoResponseResult? cloningInfo,

            int? containerSize,

            string defaultHostName,

            bool? enabled,

            ImmutableArray<string> enabledHostNames,

            string? gatewaySiteName,

            ImmutableArray<Outputs.HostNameSslStateResponseResult> hostNameSslStates,

            ImmutableArray<string> hostNames,

            bool? hostNamesDisabled,

            Outputs.HostingEnvironmentProfileResponseResult? hostingEnvironmentProfile,

            bool isDefaultContainer,

            string? kind,

            string lastModifiedTimeUtc,

            string location,

            int? maxNumberOfWorkers,

            string? microService,

            string? name,

            string outboundIpAddresses,

            bool premiumAppDeployed,

            string repositorySiteName,

            string resourceGroup,

            bool? scmSiteAlsoStopped,

            string? serverFarmId,

            Outputs.SiteConfigResponseResult? siteConfig,

            string state,

            ImmutableDictionary<string, string>? tags,

            string targetSwapSlot,

            ImmutableArray<string> trafficManagerHostNames,

            string? type,

            string usageState)
        {
            AvailabilityState = availabilityState;
            ClientAffinityEnabled = clientAffinityEnabled;
            ClientCertEnabled = clientCertEnabled;
            CloningInfo = cloningInfo;
            ContainerSize = containerSize;
            DefaultHostName = defaultHostName;
            Enabled = enabled;
            EnabledHostNames = enabledHostNames;
            GatewaySiteName = gatewaySiteName;
            HostNameSslStates = hostNameSslStates;
            HostNames = hostNames;
            HostNamesDisabled = hostNamesDisabled;
            HostingEnvironmentProfile = hostingEnvironmentProfile;
            IsDefaultContainer = isDefaultContainer;
            Kind = kind;
            LastModifiedTimeUtc = lastModifiedTimeUtc;
            Location = location;
            MaxNumberOfWorkers = maxNumberOfWorkers;
            MicroService = microService;
            Name = name;
            OutboundIpAddresses = outboundIpAddresses;
            PremiumAppDeployed = premiumAppDeployed;
            RepositorySiteName = repositorySiteName;
            ResourceGroup = resourceGroup;
            ScmSiteAlsoStopped = scmSiteAlsoStopped;
            ServerFarmId = serverFarmId;
            SiteConfig = siteConfig;
            State = state;
            Tags = tags;
            TargetSwapSlot = targetSwapSlot;
            TrafficManagerHostNames = trafficManagerHostNames;
            Type = type;
            UsageState = usageState;
        }
    }
}