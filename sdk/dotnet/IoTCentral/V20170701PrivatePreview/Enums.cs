// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.ComponentModel;
using Pulumi;

namespace Pulumi.AzureNative.IoTCentral.V20170701PrivatePreview
{
    /// <summary>
    /// The name of the SKU.
    /// </summary>
    [EnumType]
    public readonly struct AppSku : IEquatable<AppSku>
    {
        private readonly string _value;

        private AppSku(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static AppSku F1 { get; } = new AppSku("F1");
        public static AppSku S1 { get; } = new AppSku("S1");

        public static bool operator ==(AppSku left, AppSku right) => left.Equals(right);
        public static bool operator !=(AppSku left, AppSku right) => !left.Equals(right);

        public static explicit operator string(AppSku value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is AppSku other && Equals(other);
        public bool Equals(AppSku other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }
}
