# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from ... import _utilities, _tables
from . import outputs

__all__ = [
    'GetVpnConnectionResult',
    'AwaitableGetVpnConnectionResult',
    'get_vpn_connection',
]

@pulumi.output_type
class GetVpnConnectionResult:
    """
    VpnConnection Resource.
    """
    def __init__(__self__, connection_bandwidth_in_mbps=None, connection_status=None, egress_bytes_transferred=None, enable_bgp=None, etag=None, ingress_bytes_transferred=None, ipsec_policies=None, name=None, provisioning_state=None, remote_vpn_site=None, routing_weight=None, shared_key=None):
        if connection_bandwidth_in_mbps and not isinstance(connection_bandwidth_in_mbps, float):
            raise TypeError("Expected argument 'connection_bandwidth_in_mbps' to be a float")
        pulumi.set(__self__, "connection_bandwidth_in_mbps", connection_bandwidth_in_mbps)
        if connection_status and not isinstance(connection_status, str):
            raise TypeError("Expected argument 'connection_status' to be a str")
        pulumi.set(__self__, "connection_status", connection_status)
        if egress_bytes_transferred and not isinstance(egress_bytes_transferred, float):
            raise TypeError("Expected argument 'egress_bytes_transferred' to be a float")
        pulumi.set(__self__, "egress_bytes_transferred", egress_bytes_transferred)
        if enable_bgp and not isinstance(enable_bgp, bool):
            raise TypeError("Expected argument 'enable_bgp' to be a bool")
        pulumi.set(__self__, "enable_bgp", enable_bgp)
        if etag and not isinstance(etag, str):
            raise TypeError("Expected argument 'etag' to be a str")
        pulumi.set(__self__, "etag", etag)
        if ingress_bytes_transferred and not isinstance(ingress_bytes_transferred, float):
            raise TypeError("Expected argument 'ingress_bytes_transferred' to be a float")
        pulumi.set(__self__, "ingress_bytes_transferred", ingress_bytes_transferred)
        if ipsec_policies and not isinstance(ipsec_policies, list):
            raise TypeError("Expected argument 'ipsec_policies' to be a list")
        pulumi.set(__self__, "ipsec_policies", ipsec_policies)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if provisioning_state and not isinstance(provisioning_state, str):
            raise TypeError("Expected argument 'provisioning_state' to be a str")
        pulumi.set(__self__, "provisioning_state", provisioning_state)
        if remote_vpn_site and not isinstance(remote_vpn_site, dict):
            raise TypeError("Expected argument 'remote_vpn_site' to be a dict")
        pulumi.set(__self__, "remote_vpn_site", remote_vpn_site)
        if routing_weight and not isinstance(routing_weight, float):
            raise TypeError("Expected argument 'routing_weight' to be a float")
        pulumi.set(__self__, "routing_weight", routing_weight)
        if shared_key and not isinstance(shared_key, str):
            raise TypeError("Expected argument 'shared_key' to be a str")
        pulumi.set(__self__, "shared_key", shared_key)

    @property
    @pulumi.getter(name="connectionBandwidthInMbps")
    def connection_bandwidth_in_mbps(self) -> float:
        """
        Expected bandwidth in MBPS.
        """
        return pulumi.get(self, "connection_bandwidth_in_mbps")

    @property
    @pulumi.getter(name="connectionStatus")
    def connection_status(self) -> Optional[str]:
        """
        The connection status.
        """
        return pulumi.get(self, "connection_status")

    @property
    @pulumi.getter(name="egressBytesTransferred")
    def egress_bytes_transferred(self) -> float:
        """
        Egress bytes transferred.
        """
        return pulumi.get(self, "egress_bytes_transferred")

    @property
    @pulumi.getter(name="enableBgp")
    def enable_bgp(self) -> Optional[bool]:
        """
        EnableBgp flag
        """
        return pulumi.get(self, "enable_bgp")

    @property
    @pulumi.getter
    def etag(self) -> str:
        """
        Gets a unique read-only string that changes whenever the resource is updated.
        """
        return pulumi.get(self, "etag")

    @property
    @pulumi.getter(name="ingressBytesTransferred")
    def ingress_bytes_transferred(self) -> float:
        """
        Ingress bytes transferred.
        """
        return pulumi.get(self, "ingress_bytes_transferred")

    @property
    @pulumi.getter(name="ipsecPolicies")
    def ipsec_policies(self) -> Optional[List['outputs.IpsecPolicyResponse']]:
        """
        The IPSec Policies to be considered by this connection.
        """
        return pulumi.get(self, "ipsec_policies")

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        """
        The name of the resource that is unique within a resource group. This name can be used to access the resource.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="provisioningState")
    def provisioning_state(self) -> Optional[str]:
        """
        The provisioning state of the resource.
        """
        return pulumi.get(self, "provisioning_state")

    @property
    @pulumi.getter(name="remoteVpnSite")
    def remote_vpn_site(self) -> Optional['outputs.SubResourceResponse']:
        """
        Id of the connected vpn site.
        """
        return pulumi.get(self, "remote_vpn_site")

    @property
    @pulumi.getter(name="routingWeight")
    def routing_weight(self) -> Optional[float]:
        """
        routing weight for vpn connection.
        """
        return pulumi.get(self, "routing_weight")

    @property
    @pulumi.getter(name="sharedKey")
    def shared_key(self) -> Optional[str]:
        """
        SharedKey for the vpn connection.
        """
        return pulumi.get(self, "shared_key")


class AwaitableGetVpnConnectionResult(GetVpnConnectionResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetVpnConnectionResult(
            connection_bandwidth_in_mbps=self.connection_bandwidth_in_mbps,
            connection_status=self.connection_status,
            egress_bytes_transferred=self.egress_bytes_transferred,
            enable_bgp=self.enable_bgp,
            etag=self.etag,
            ingress_bytes_transferred=self.ingress_bytes_transferred,
            ipsec_policies=self.ipsec_policies,
            name=self.name,
            provisioning_state=self.provisioning_state,
            remote_vpn_site=self.remote_vpn_site,
            routing_weight=self.routing_weight,
            shared_key=self.shared_key)


def get_vpn_connection(gateway_name: Optional[str] = None,
                       name: Optional[str] = None,
                       resource_group_name: Optional[str] = None,
                       opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetVpnConnectionResult:
    """
    Use this data source to access information about an existing resource.

    :param str gateway_name: The name of the gateway.
    :param str name: The name of the vpn connection.
    :param str resource_group_name: The resource group name of the VpnGateway.
    """
    __args__ = dict()
    __args__['gatewayName'] = gateway_name
    __args__['name'] = name
    __args__['resourceGroupName'] = resource_group_name
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('azurerm:network/v20180701:getVpnConnection', __args__, opts=opts, typ=GetVpnConnectionResult).value

    return AwaitableGetVpnConnectionResult(
        connection_bandwidth_in_mbps=__ret__.connection_bandwidth_in_mbps,
        connection_status=__ret__.connection_status,
        egress_bytes_transferred=__ret__.egress_bytes_transferred,
        enable_bgp=__ret__.enable_bgp,
        etag=__ret__.etag,
        ingress_bytes_transferred=__ret__.ingress_bytes_transferred,
        ipsec_policies=__ret__.ipsec_policies,
        name=__ret__.name,
        provisioning_state=__ret__.provisioning_state,
        remote_vpn_site=__ret__.remote_vpn_site,
        routing_weight=__ret__.routing_weight,
        shared_key=__ret__.shared_key)