# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from ... import _utilities, _tables

__all__ = [
    'GetIpFirewallRuleResult',
    'AwaitableGetIpFirewallRuleResult',
    'get_ip_firewall_rule',
]

@pulumi.output_type
class GetIpFirewallRuleResult:
    """
    IP firewall rule
    """
    def __init__(__self__, end_ip_address=None, name=None, provisioning_state=None, start_ip_address=None, type=None):
        if end_ip_address and not isinstance(end_ip_address, str):
            raise TypeError("Expected argument 'end_ip_address' to be a str")
        pulumi.set(__self__, "end_ip_address", end_ip_address)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if provisioning_state and not isinstance(provisioning_state, str):
            raise TypeError("Expected argument 'provisioning_state' to be a str")
        pulumi.set(__self__, "provisioning_state", provisioning_state)
        if start_ip_address and not isinstance(start_ip_address, str):
            raise TypeError("Expected argument 'start_ip_address' to be a str")
        pulumi.set(__self__, "start_ip_address", start_ip_address)
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter(name="endIpAddress")
    def end_ip_address(self) -> Optional[str]:
        """
        The end IP address of the firewall rule. Must be IPv4 format. Must be greater than or equal to startIpAddress
        """
        return pulumi.get(self, "end_ip_address")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The name of the resource
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="provisioningState")
    def provisioning_state(self) -> str:
        """
        Resource provisioning state
        """
        return pulumi.get(self, "provisioning_state")

    @property
    @pulumi.getter(name="startIpAddress")
    def start_ip_address(self) -> Optional[str]:
        """
        The start IP address of the firewall rule. Must be IPv4 format
        """
        return pulumi.get(self, "start_ip_address")

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
        """
        return pulumi.get(self, "type")


class AwaitableGetIpFirewallRuleResult(GetIpFirewallRuleResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetIpFirewallRuleResult(
            end_ip_address=self.end_ip_address,
            name=self.name,
            provisioning_state=self.provisioning_state,
            start_ip_address=self.start_ip_address,
            type=self.type)


def get_ip_firewall_rule(resource_group_name: Optional[str] = None,
                         rule_name: Optional[str] = None,
                         workspace_name: Optional[str] = None,
                         opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetIpFirewallRuleResult:
    """
    Use this data source to access information about an existing resource.

    :param str resource_group_name: The name of the resource group. The name is case insensitive.
    :param str rule_name: The IP firewall rule name
    :param str workspace_name: The name of the workspace
    """
    __args__ = dict()
    __args__['resourceGroupName'] = resource_group_name
    __args__['ruleName'] = rule_name
    __args__['workspaceName'] = workspace_name
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('azurerm:synapse/v20190601preview:getIpFirewallRule', __args__, opts=opts, typ=GetIpFirewallRuleResult).value

    return AwaitableGetIpFirewallRuleResult(
        end_ip_address=__ret__.end_ip_address,
        name=__ret__.name,
        provisioning_state=__ret__.provisioning_state,
        start_ip_address=__ret__.start_ip_address,
        type=__ret__.type)