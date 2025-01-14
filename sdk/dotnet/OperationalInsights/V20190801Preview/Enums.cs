// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.ComponentModel;
using Pulumi;

namespace Pulumi.AzureNative.OperationalInsights.V20190801Preview
{
    /// <summary>
    /// The identity type.
    /// </summary>
    [EnumType]
    public readonly struct IdentityType : IEquatable<IdentityType>
    {
        private readonly string _value;

        private IdentityType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static IdentityType SystemAssigned { get; } = new IdentityType("SystemAssigned");
        public static IdentityType None { get; } = new IdentityType("None");

        public static bool operator ==(IdentityType left, IdentityType right) => left.Equals(right);
        public static bool operator !=(IdentityType left, IdentityType right) => !left.Equals(right);

        public static explicit operator string(IdentityType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is IdentityType other && Equals(other);
        public bool Equals(IdentityType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// The name of the SKU.
    /// </summary>
    [EnumType]
    public readonly struct SkuNameEnum : IEquatable<SkuNameEnum>
    {
        private readonly string _value;

        private SkuNameEnum(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static SkuNameEnum CapacityReservation { get; } = new SkuNameEnum("CapacityReservation");

        public static bool operator ==(SkuNameEnum left, SkuNameEnum right) => left.Equals(right);
        public static bool operator !=(SkuNameEnum left, SkuNameEnum right) => !left.Equals(right);

        public static explicit operator string(SkuNameEnum value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is SkuNameEnum other && Equals(other);
        public bool Equals(SkuNameEnum other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }
}
