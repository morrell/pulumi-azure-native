// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.NetApp.V20200601
{
    /// <summary>
    /// Snapshot policy information
    /// </summary>
    public partial class SnapshotPolicy : Pulumi.CustomResource
    {
        /// <summary>
        /// Schedule for daily snapshots
        /// </summary>
        [Output("dailySchedule")]
        public Output<ImmutableDictionary<string, object>?> DailySchedule { get; private set; } = null!;

        /// <summary>
        /// The property to decide policy is enabled or not
        /// </summary>
        [Output("enabled")]
        public Output<bool?> Enabled { get; private set; } = null!;

        /// <summary>
        /// Schedule for hourly snapshots
        /// </summary>
        [Output("hourlySchedule")]
        public Output<ImmutableDictionary<string, object>?> HourlySchedule { get; private set; } = null!;

        /// <summary>
        /// Resource location
        /// </summary>
        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// Schedule for monthly snapshots
        /// </summary>
        [Output("monthlySchedule")]
        public Output<ImmutableDictionary<string, object>?> MonthlySchedule { get; private set; } = null!;

        /// <summary>
        /// Resource name
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Resource tags
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// Resource type
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;

        /// <summary>
        /// Schedule for weekly snapshots
        /// </summary>
        [Output("weeklySchedule")]
        public Output<ImmutableDictionary<string, object>?> WeeklySchedule { get; private set; } = null!;


        /// <summary>
        /// Create a SnapshotPolicy resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public SnapshotPolicy(string name, SnapshotPolicyArgs args, CustomResourceOptions? options = null)
            : base("azurerm:netapp/v20200601:snapshotPolicy", name, args ?? new SnapshotPolicyArgs(), MakeResourceOptions(options, ""))
        {
        }

        private SnapshotPolicy(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azurerm:netapp/v20200601:snapshotPolicy", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new Pulumi.Alias { Type = "azurerm:netapp/latest:snapshotPolicy"},
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing SnapshotPolicy resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static SnapshotPolicy Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new SnapshotPolicy(name, id, options);
        }
    }

    public sealed class SnapshotPolicyArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the NetApp account
        /// </summary>
        [Input("accountName", required: true)]
        public Input<string> AccountName { get; set; } = null!;

        [Input("dailySchedule")]
        private InputMap<object>? _dailySchedule;

        /// <summary>
        /// Schedule for daily snapshots
        /// </summary>
        public InputMap<object> DailySchedule
        {
            get => _dailySchedule ?? (_dailySchedule = new InputMap<object>());
            set => _dailySchedule = value;
        }

        /// <summary>
        /// The property to decide policy is enabled or not
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        [Input("hourlySchedule")]
        private InputMap<object>? _hourlySchedule;

        /// <summary>
        /// Schedule for hourly snapshots
        /// </summary>
        public InputMap<object> HourlySchedule
        {
            get => _hourlySchedule ?? (_hourlySchedule = new InputMap<object>());
            set => _hourlySchedule = value;
        }

        /// <summary>
        /// Resource location
        /// </summary>
        [Input("location", required: true)]
        public Input<string> Location { get; set; } = null!;

        [Input("monthlySchedule")]
        private InputMap<object>? _monthlySchedule;

        /// <summary>
        /// Schedule for monthly snapshots
        /// </summary>
        public InputMap<object> MonthlySchedule
        {
            get => _monthlySchedule ?? (_monthlySchedule = new InputMap<object>());
            set => _monthlySchedule = value;
        }

        /// <summary>
        /// The name of the resource group.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The name of the snapshot policy target
        /// </summary>
        [Input("snapshotPolicyName", required: true)]
        public Input<string> SnapshotPolicyName { get; set; } = null!;

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// Resource tags
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        [Input("weeklySchedule")]
        private InputMap<object>? _weeklySchedule;

        /// <summary>
        /// Schedule for weekly snapshots
        /// </summary>
        public InputMap<object> WeeklySchedule
        {
            get => _weeklySchedule ?? (_weeklySchedule = new InputMap<object>());
            set => _weeklySchedule = value;
        }

        public SnapshotPolicyArgs()
        {
        }
    }
}