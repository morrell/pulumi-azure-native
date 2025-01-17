// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.ComponentModel;
using Pulumi;

namespace Pulumi.AzureNative.RecoveryServices.V20180110
{
    /// <summary>
    /// A value indicating whether the auto update is enabled.
    /// </summary>
    [EnumType]
    public readonly struct AgentAutoUpdateStatus : IEquatable<AgentAutoUpdateStatus>
    {
        private readonly string _value;

        private AgentAutoUpdateStatus(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static AgentAutoUpdateStatus Disabled { get; } = new AgentAutoUpdateStatus("Disabled");
        public static AgentAutoUpdateStatus Enabled { get; } = new AgentAutoUpdateStatus("Enabled");

        public static bool operator ==(AgentAutoUpdateStatus left, AgentAutoUpdateStatus right) => left.Equals(right);
        public static bool operator !=(AgentAutoUpdateStatus left, AgentAutoUpdateStatus right) => !left.Equals(right);

        public static explicit operator string(AgentAutoUpdateStatus value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is AgentAutoUpdateStatus other && Equals(other);
        public bool Equals(AgentAutoUpdateStatus other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// The disk type.
    /// </summary>
    [EnumType]
    public readonly struct DiskAccountType : IEquatable<DiskAccountType>
    {
        private readonly string _value;

        private DiskAccountType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static DiskAccountType Standard_LRS { get; } = new DiskAccountType("Standard_LRS");
        public static DiskAccountType Premium_LRS { get; } = new DiskAccountType("Premium_LRS");
        public static DiskAccountType StandardSSD_LRS { get; } = new DiskAccountType("StandardSSD_LRS");

        public static bool operator ==(DiskAccountType left, DiskAccountType right) => left.Equals(right);
        public static bool operator !=(DiskAccountType left, DiskAccountType right) => !left.Equals(right);

        public static explicit operator string(DiskAccountType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is DiskAccountType other && Equals(other);
        public bool Equals(DiskAccountType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// The failover deployment model.
    /// </summary>
    [EnumType]
    public readonly struct FailoverDeploymentModel : IEquatable<FailoverDeploymentModel>
    {
        private readonly string _value;

        private FailoverDeploymentModel(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static FailoverDeploymentModel NotApplicable { get; } = new FailoverDeploymentModel("NotApplicable");
        public static FailoverDeploymentModel Classic { get; } = new FailoverDeploymentModel("Classic");
        public static FailoverDeploymentModel ResourceManager { get; } = new FailoverDeploymentModel("ResourceManager");

        public static bool operator ==(FailoverDeploymentModel left, FailoverDeploymentModel right) => left.Equals(right);
        public static bool operator !=(FailoverDeploymentModel left, FailoverDeploymentModel right) => !left.Equals(right);

        public static explicit operator string(FailoverDeploymentModel value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is FailoverDeploymentModel other && Equals(other);
        public bool Equals(FailoverDeploymentModel other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// License type.
    /// </summary>
    [EnumType]
    public readonly struct LicenseType : IEquatable<LicenseType>
    {
        private readonly string _value;

        private LicenseType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static LicenseType NotSpecified { get; } = new LicenseType("NotSpecified");
        public static LicenseType NoLicenseType { get; } = new LicenseType("NoLicenseType");
        public static LicenseType WindowsServer { get; } = new LicenseType("WindowsServer");

        public static bool operator ==(LicenseType left, LicenseType right) => left.Equals(right);
        public static bool operator !=(LicenseType left, LicenseType right) => !left.Equals(right);

        public static explicit operator string(LicenseType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is LicenseType other && Equals(other);
        public bool Equals(LicenseType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    [EnumType]
    public readonly struct PossibleOperationsDirections : IEquatable<PossibleOperationsDirections>
    {
        private readonly string _value;

        private PossibleOperationsDirections(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static PossibleOperationsDirections PrimaryToRecovery { get; } = new PossibleOperationsDirections("PrimaryToRecovery");
        public static PossibleOperationsDirections RecoveryToPrimary { get; } = new PossibleOperationsDirections("RecoveryToPrimary");

        public static bool operator ==(PossibleOperationsDirections left, PossibleOperationsDirections right) => left.Equals(right);
        public static bool operator !=(PossibleOperationsDirections left, PossibleOperationsDirections right) => !left.Equals(right);

        public static explicit operator string(PossibleOperationsDirections value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is PossibleOperationsDirections other && Equals(other);
        public bool Equals(PossibleOperationsDirections other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// The group type.
    /// </summary>
    [EnumType]
    public readonly struct RecoveryPlanGroupType : IEquatable<RecoveryPlanGroupType>
    {
        private readonly string _value;

        private RecoveryPlanGroupType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static RecoveryPlanGroupType Shutdown { get; } = new RecoveryPlanGroupType("Shutdown");
        public static RecoveryPlanGroupType Boot { get; } = new RecoveryPlanGroupType("Boot");
        public static RecoveryPlanGroupType Failover { get; } = new RecoveryPlanGroupType("Failover");

        public static bool operator ==(RecoveryPlanGroupType left, RecoveryPlanGroupType right) => left.Equals(right);
        public static bool operator !=(RecoveryPlanGroupType left, RecoveryPlanGroupType right) => !left.Equals(right);

        public static explicit operator string(RecoveryPlanGroupType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is RecoveryPlanGroupType other && Equals(other);
        public bool Equals(RecoveryPlanGroupType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    [EnumType]
    public readonly struct ReplicationProtectedItemOperation : IEquatable<ReplicationProtectedItemOperation>
    {
        private readonly string _value;

        private ReplicationProtectedItemOperation(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static ReplicationProtectedItemOperation ReverseReplicate { get; } = new ReplicationProtectedItemOperation("ReverseReplicate");
        public static ReplicationProtectedItemOperation Commit { get; } = new ReplicationProtectedItemOperation("Commit");
        public static ReplicationProtectedItemOperation PlannedFailover { get; } = new ReplicationProtectedItemOperation("PlannedFailover");
        public static ReplicationProtectedItemOperation UnplannedFailover { get; } = new ReplicationProtectedItemOperation("UnplannedFailover");
        public static ReplicationProtectedItemOperation DisableProtection { get; } = new ReplicationProtectedItemOperation("DisableProtection");
        public static ReplicationProtectedItemOperation TestFailover { get; } = new ReplicationProtectedItemOperation("TestFailover");
        public static ReplicationProtectedItemOperation TestFailoverCleanup { get; } = new ReplicationProtectedItemOperation("TestFailoverCleanup");
        public static ReplicationProtectedItemOperation Failback { get; } = new ReplicationProtectedItemOperation("Failback");
        public static ReplicationProtectedItemOperation FinalizeFailback { get; } = new ReplicationProtectedItemOperation("FinalizeFailback");
        public static ReplicationProtectedItemOperation ChangePit { get; } = new ReplicationProtectedItemOperation("ChangePit");
        public static ReplicationProtectedItemOperation RepairReplication { get; } = new ReplicationProtectedItemOperation("RepairReplication");
        public static ReplicationProtectedItemOperation SwitchProtection { get; } = new ReplicationProtectedItemOperation("SwitchProtection");
        public static ReplicationProtectedItemOperation CompleteMigration { get; } = new ReplicationProtectedItemOperation("CompleteMigration");

        public static bool operator ==(ReplicationProtectedItemOperation left, ReplicationProtectedItemOperation right) => left.Equals(right);
        public static bool operator !=(ReplicationProtectedItemOperation left, ReplicationProtectedItemOperation right) => !left.Equals(right);

        public static explicit operator string(ReplicationProtectedItemOperation value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is ReplicationProtectedItemOperation other && Equals(other);
        public bool Equals(ReplicationProtectedItemOperation other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// A value indicating whether multi-VM sync has to be enabled. Value should be 'Enabled' or 'Disabled'.
    /// </summary>
    [EnumType]
    public readonly struct SetMultiVmSyncStatus : IEquatable<SetMultiVmSyncStatus>
    {
        private readonly string _value;

        private SetMultiVmSyncStatus(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static SetMultiVmSyncStatus Enable { get; } = new SetMultiVmSyncStatus("Enable");
        public static SetMultiVmSyncStatus Disable { get; } = new SetMultiVmSyncStatus("Disable");

        public static bool operator ==(SetMultiVmSyncStatus left, SetMultiVmSyncStatus right) => left.Equals(right);
        public static bool operator !=(SetMultiVmSyncStatus left, SetMultiVmSyncStatus right) => !left.Equals(right);

        public static explicit operator string(SetMultiVmSyncStatus value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is SetMultiVmSyncStatus other && Equals(other);
        public bool Equals(SetMultiVmSyncStatus other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }
}
