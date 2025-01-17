// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.ComponentModel;
using Pulumi;

namespace Pulumi.AzureNative.ApiManagement.V20161010
{
    [EnumType]
    public readonly struct ApiProtocolContract : IEquatable<ApiProtocolContract>
    {
        private readonly string _value;

        private ApiProtocolContract(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static ApiProtocolContract Http { get; } = new ApiProtocolContract("Http");
        public static ApiProtocolContract Https { get; } = new ApiProtocolContract("Https");

        public static bool operator ==(ApiProtocolContract left, ApiProtocolContract right) => left.Equals(right);
        public static bool operator !=(ApiProtocolContract left, ApiProtocolContract right) => !left.Equals(right);

        public static explicit operator string(ApiProtocolContract value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is ApiProtocolContract other && Equals(other);
        public bool Equals(ApiProtocolContract other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// Type of API.
    /// </summary>
    [EnumType]
    public readonly struct ApiTypeContract : IEquatable<ApiTypeContract>
    {
        private readonly string _value;

        private ApiTypeContract(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static ApiTypeContract Http { get; } = new ApiTypeContract("Http");
        public static ApiTypeContract Soap { get; } = new ApiTypeContract("Soap");

        public static bool operator ==(ApiTypeContract left, ApiTypeContract right) => left.Equals(right);
        public static bool operator !=(ApiTypeContract left, ApiTypeContract right) => !left.Equals(right);

        public static explicit operator string(ApiTypeContract value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is ApiTypeContract other && Equals(other);
        public bool Equals(ApiTypeContract other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// Backend communication protocol.
    /// </summary>
    [EnumType]
    public readonly struct BackendProtocol : IEquatable<BackendProtocol>
    {
        private readonly string _value;

        private BackendProtocol(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static BackendProtocol Http { get; } = new BackendProtocol("http");
        public static BackendProtocol Soap { get; } = new BackendProtocol("soap");

        public static bool operator ==(BackendProtocol left, BackendProtocol right) => left.Equals(right);
        public static bool operator !=(BackendProtocol left, BackendProtocol right) => !left.Equals(right);

        public static explicit operator string(BackendProtocol value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is BackendProtocol other && Equals(other);
        public bool Equals(BackendProtocol other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    [EnumType]
    public readonly struct BearerTokenSendingMethodsContract : IEquatable<BearerTokenSendingMethodsContract>
    {
        private readonly string _value;

        private BearerTokenSendingMethodsContract(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static BearerTokenSendingMethodsContract AuthorizationHeader { get; } = new BearerTokenSendingMethodsContract("authorizationHeader");
        public static BearerTokenSendingMethodsContract Query { get; } = new BearerTokenSendingMethodsContract("query");

        public static bool operator ==(BearerTokenSendingMethodsContract left, BearerTokenSendingMethodsContract right) => left.Equals(right);
        public static bool operator !=(BearerTokenSendingMethodsContract left, BearerTokenSendingMethodsContract right) => !left.Equals(right);

        public static explicit operator string(BearerTokenSendingMethodsContract value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is BearerTokenSendingMethodsContract other && Equals(other);
        public bool Equals(BearerTokenSendingMethodsContract other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    [EnumType]
    public readonly struct ClientAuthenticationMethodContract : IEquatable<ClientAuthenticationMethodContract>
    {
        private readonly string _value;

        private ClientAuthenticationMethodContract(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static ClientAuthenticationMethodContract Basic { get; } = new ClientAuthenticationMethodContract("Basic");
        public static ClientAuthenticationMethodContract Body { get; } = new ClientAuthenticationMethodContract("Body");

        public static bool operator ==(ClientAuthenticationMethodContract left, ClientAuthenticationMethodContract right) => left.Equals(right);
        public static bool operator !=(ClientAuthenticationMethodContract left, ClientAuthenticationMethodContract right) => !left.Equals(right);

        public static explicit operator string(ClientAuthenticationMethodContract value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is ClientAuthenticationMethodContract other && Equals(other);
        public bool Equals(ClientAuthenticationMethodContract other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    [EnumType]
    public readonly struct GrantTypesContract : IEquatable<GrantTypesContract>
    {
        private readonly string _value;

        private GrantTypesContract(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static GrantTypesContract AuthorizationCode { get; } = new GrantTypesContract("authorizationCode");
        public static GrantTypesContract @Implicit { get; } = new GrantTypesContract("implicit");
        public static GrantTypesContract ResourceOwnerPassword { get; } = new GrantTypesContract("resourceOwnerPassword");
        public static GrantTypesContract ClientCredentials { get; } = new GrantTypesContract("clientCredentials");

        public static bool operator ==(GrantTypesContract left, GrantTypesContract right) => left.Equals(right);
        public static bool operator !=(GrantTypesContract left, GrantTypesContract right) => !left.Equals(right);

        public static explicit operator string(GrantTypesContract value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is GrantTypesContract other && Equals(other);
        public bool Equals(GrantTypesContract other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// Group type.
    /// </summary>
    [EnumType]
    public readonly struct GroupTypeContract : IEquatable<GroupTypeContract>
    {
        private readonly string _value;

        private GroupTypeContract(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static GroupTypeContract Custom { get; } = new GroupTypeContract("Custom");
        public static GroupTypeContract System { get; } = new GroupTypeContract("System");
        public static GroupTypeContract External { get; } = new GroupTypeContract("External");

        public static bool operator ==(GroupTypeContract left, GroupTypeContract right) => left.Equals(right);
        public static bool operator !=(GroupTypeContract left, GroupTypeContract right) => !left.Equals(right);

        public static explicit operator string(GroupTypeContract value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is GroupTypeContract other && Equals(other);
        public bool Equals(GroupTypeContract other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// Hostname type.
    /// </summary>
    [EnumType]
    public readonly struct HostnameType : IEquatable<HostnameType>
    {
        private readonly string _value;

        private HostnameType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static HostnameType Proxy { get; } = new HostnameType("Proxy");
        public static HostnameType Portal { get; } = new HostnameType("Portal");
        public static HostnameType Management { get; } = new HostnameType("Management");
        public static HostnameType Scm { get; } = new HostnameType("Scm");

        public static bool operator ==(HostnameType left, HostnameType right) => left.Equals(right);
        public static bool operator !=(HostnameType left, HostnameType right) => !left.Equals(right);

        public static explicit operator string(HostnameType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is HostnameType other && Equals(other);
        public bool Equals(HostnameType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// Identity Provider Type identifier.
    /// </summary>
    [EnumType]
    public readonly struct IdentityProviderNameType : IEquatable<IdentityProviderNameType>
    {
        private readonly string _value;

        private IdentityProviderNameType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static IdentityProviderNameType Facebook { get; } = new IdentityProviderNameType("facebook");
        public static IdentityProviderNameType Google { get; } = new IdentityProviderNameType("google");
        public static IdentityProviderNameType Microsoft { get; } = new IdentityProviderNameType("microsoft");
        public static IdentityProviderNameType Twitter { get; } = new IdentityProviderNameType("twitter");
        public static IdentityProviderNameType Aad { get; } = new IdentityProviderNameType("aad");
        public static IdentityProviderNameType AadB2C { get; } = new IdentityProviderNameType("aadB2C");

        public static bool operator ==(IdentityProviderNameType left, IdentityProviderNameType right) => left.Equals(right);
        public static bool operator !=(IdentityProviderNameType left, IdentityProviderNameType right) => !left.Equals(right);

        public static explicit operator string(IdentityProviderNameType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is IdentityProviderNameType other && Equals(other);
        public bool Equals(IdentityProviderNameType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// The Key to be used to generate token for user.
    /// </summary>
    [EnumType]
    public readonly struct KeyTypeContract : IEquatable<KeyTypeContract>
    {
        private readonly string _value;

        private KeyTypeContract(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static KeyTypeContract Primary { get; } = new KeyTypeContract("primary");
        public static KeyTypeContract Secondary { get; } = new KeyTypeContract("secondary");

        public static bool operator ==(KeyTypeContract left, KeyTypeContract right) => left.Equals(right);
        public static bool operator !=(KeyTypeContract left, KeyTypeContract right) => !left.Equals(right);

        public static explicit operator string(KeyTypeContract value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is KeyTypeContract other && Equals(other);
        public bool Equals(KeyTypeContract other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// Logger type.
    /// </summary>
    [EnumType]
    public readonly struct LoggerTypeContract : IEquatable<LoggerTypeContract>
    {
        private readonly string _value;

        private LoggerTypeContract(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static LoggerTypeContract AzureEventHub { get; } = new LoggerTypeContract("AzureEventHub");

        public static bool operator ==(LoggerTypeContract left, LoggerTypeContract right) => left.Equals(right);
        public static bool operator !=(LoggerTypeContract left, LoggerTypeContract right) => !left.Equals(right);

        public static explicit operator string(LoggerTypeContract value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is LoggerTypeContract other && Equals(other);
        public bool Equals(LoggerTypeContract other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    [EnumType]
    public readonly struct MethodContract : IEquatable<MethodContract>
    {
        private readonly string _value;

        private MethodContract(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static MethodContract HEAD { get; } = new MethodContract("HEAD");
        public static MethodContract OPTIONS { get; } = new MethodContract("OPTIONS");
        public static MethodContract TRACE { get; } = new MethodContract("TRACE");
        public static MethodContract GET { get; } = new MethodContract("GET");
        public static MethodContract POST { get; } = new MethodContract("POST");
        public static MethodContract PUT { get; } = new MethodContract("PUT");
        public static MethodContract PATCH { get; } = new MethodContract("PATCH");
        public static MethodContract DELETE { get; } = new MethodContract("DELETE");

        public static bool operator ==(MethodContract left, MethodContract right) => left.Equals(right);
        public static bool operator !=(MethodContract left, MethodContract right) => !left.Equals(right);

        public static explicit operator string(MethodContract value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is MethodContract other && Equals(other);
        public bool Equals(MethodContract other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// whether product is published or not. Published products are discoverable by users of developer portal. Non published products are visible only to administrators. Default state of Product is NotPublished.
    /// </summary>
    [EnumType]
    public readonly struct ProductStateContract : IEquatable<ProductStateContract>
    {
        private readonly string _value;

        private ProductStateContract(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static ProductStateContract NotPublished { get; } = new ProductStateContract("NotPublished");
        public static ProductStateContract Published { get; } = new ProductStateContract("Published");

        public static bool operator ==(ProductStateContract left, ProductStateContract right) => left.Equals(right);
        public static bool operator !=(ProductStateContract left, ProductStateContract right) => !left.Equals(right);

        public static explicit operator string(ProductStateContract value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is ProductStateContract other && Equals(other);
        public bool Equals(ProductStateContract other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// Name of the Sku.
    /// </summary>
    [EnumType]
    public readonly struct SkuType : IEquatable<SkuType>
    {
        private readonly string _value;

        private SkuType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static SkuType Developer { get; } = new SkuType("Developer");
        public static SkuType Standard { get; } = new SkuType("Standard");
        public static SkuType Premium { get; } = new SkuType("Premium");

        public static bool operator ==(SkuType left, SkuType right) => left.Equals(right);
        public static bool operator !=(SkuType left, SkuType right) => !left.Equals(right);

        public static explicit operator string(SkuType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is SkuType other && Equals(other);
        public bool Equals(SkuType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// Initial subscription state. If no value is specified, subscription is created with Submitted state. Possible states are * active – the subscription is active, * suspended – the subscription is blocked, and the subscriber cannot call any APIs of the product, * submitted – the subscription request has been made by the developer, but has not yet been approved or rejected, * rejected – the subscription request has been denied by an administrator, * cancelled – the subscription has been cancelled by the developer or administrator, * expired – the subscription reached its expiration date and was deactivated.
    /// </summary>
    [EnumType]
    public readonly struct SubscriptionStateContract : IEquatable<SubscriptionStateContract>
    {
        private readonly string _value;

        private SubscriptionStateContract(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static SubscriptionStateContract Suspended { get; } = new SubscriptionStateContract("Suspended");
        public static SubscriptionStateContract Active { get; } = new SubscriptionStateContract("Active");
        public static SubscriptionStateContract Expired { get; } = new SubscriptionStateContract("Expired");
        public static SubscriptionStateContract Submitted { get; } = new SubscriptionStateContract("Submitted");
        public static SubscriptionStateContract Rejected { get; } = new SubscriptionStateContract("Rejected");
        public static SubscriptionStateContract Cancelled { get; } = new SubscriptionStateContract("Cancelled");

        public static bool operator ==(SubscriptionStateContract left, SubscriptionStateContract right) => left.Equals(right);
        public static bool operator !=(SubscriptionStateContract left, SubscriptionStateContract right) => !left.Equals(right);

        public static explicit operator string(SubscriptionStateContract value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is SubscriptionStateContract other && Equals(other);
        public bool Equals(SubscriptionStateContract other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// Account state. Specifies whether the user is active or not. Blocked users are unable to sign into the developer portal or call any APIs of subscribed products. Default state is Active.
    /// </summary>
    [EnumType]
    public readonly struct UserStateContract : IEquatable<UserStateContract>
    {
        private readonly string _value;

        private UserStateContract(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static UserStateContract Active { get; } = new UserStateContract("Active");
        public static UserStateContract Blocked { get; } = new UserStateContract("Blocked");

        public static bool operator ==(UserStateContract left, UserStateContract right) => left.Equals(right);
        public static bool operator !=(UserStateContract left, UserStateContract right) => !left.Equals(right);

        public static explicit operator string(UserStateContract value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is UserStateContract other && Equals(other);
        public bool Equals(UserStateContract other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// The type of VPN in which API Management service needs to be configured in. None (Default Value) means the API Management service is not part of any Virtual Network, External means the API Management deployment is set up inside a Virtual Network having an Internet Facing Endpoint, and Internal means that API Management deployment is setup inside a Virtual Network having an Intranet Facing Endpoint only.
    /// </summary>
    [EnumType]
    public readonly struct VirtualNetworkType : IEquatable<VirtualNetworkType>
    {
        private readonly string _value;

        private VirtualNetworkType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static VirtualNetworkType None { get; } = new VirtualNetworkType("None");
        public static VirtualNetworkType External { get; } = new VirtualNetworkType("External");
        public static VirtualNetworkType Internal { get; } = new VirtualNetworkType("Internal");

        public static bool operator ==(VirtualNetworkType left, VirtualNetworkType right) => left.Equals(right);
        public static bool operator !=(VirtualNetworkType left, VirtualNetworkType right) => !left.Equals(right);

        public static explicit operator string(VirtualNetworkType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is VirtualNetworkType other && Equals(other);
        public bool Equals(VirtualNetworkType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }
}
