# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from ... import _utilities, _tables

__all__ = [
    'IPRuleArgs',
    'NetworkRuleSetArgs',
    'PoliciesArgs',
    'QuarantinePolicyArgs',
    'RetentionPolicyArgs',
    'SkuArgs',
    'StorageAccountPropertiesArgs',
    'TrustPolicyArgs',
    'VirtualNetworkRuleArgs',
]

@pulumi.input_type
class IPRuleArgs:
    def __init__(__self__, *,
                 i_p_address_or_range: pulumi.Input[str],
                 action: Optional[pulumi.Input[str]] = None):
        """
        IP rule with specific IP or IP range in CIDR format.
        :param pulumi.Input[str] i_p_address_or_range: Specifies the IP or IP range in CIDR format. Only IPV4 address is allowed.
        :param pulumi.Input[str] action: The action of IP ACL rule.
        """
        pulumi.set(__self__, "i_p_address_or_range", i_p_address_or_range)
        if action is not None:
            pulumi.set(__self__, "action", action)

    @property
    @pulumi.getter(name="iPAddressOrRange")
    def i_p_address_or_range(self) -> pulumi.Input[str]:
        """
        Specifies the IP or IP range in CIDR format. Only IPV4 address is allowed.
        """
        return pulumi.get(self, "i_p_address_or_range")

    @i_p_address_or_range.setter
    def i_p_address_or_range(self, value: pulumi.Input[str]):
        pulumi.set(self, "i_p_address_or_range", value)

    @property
    @pulumi.getter
    def action(self) -> Optional[pulumi.Input[str]]:
        """
        The action of IP ACL rule.
        """
        return pulumi.get(self, "action")

    @action.setter
    def action(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "action", value)


@pulumi.input_type
class NetworkRuleSetArgs:
    def __init__(__self__, *,
                 default_action: pulumi.Input[str],
                 ip_rules: Optional[pulumi.Input[List[pulumi.Input['IPRuleArgs']]]] = None,
                 virtual_network_rules: Optional[pulumi.Input[List[pulumi.Input['VirtualNetworkRuleArgs']]]] = None):
        """
        The network rule set for a container registry.
        :param pulumi.Input[str] default_action: The default action of allow or deny when no other rules match.
        :param pulumi.Input[List[pulumi.Input['IPRuleArgs']]] ip_rules: The IP ACL rules.
        :param pulumi.Input[List[pulumi.Input['VirtualNetworkRuleArgs']]] virtual_network_rules: The virtual network rules.
        """
        pulumi.set(__self__, "default_action", default_action)
        if ip_rules is not None:
            pulumi.set(__self__, "ip_rules", ip_rules)
        if virtual_network_rules is not None:
            pulumi.set(__self__, "virtual_network_rules", virtual_network_rules)

    @property
    @pulumi.getter(name="defaultAction")
    def default_action(self) -> pulumi.Input[str]:
        """
        The default action of allow or deny when no other rules match.
        """
        return pulumi.get(self, "default_action")

    @default_action.setter
    def default_action(self, value: pulumi.Input[str]):
        pulumi.set(self, "default_action", value)

    @property
    @pulumi.getter(name="ipRules")
    def ip_rules(self) -> Optional[pulumi.Input[List[pulumi.Input['IPRuleArgs']]]]:
        """
        The IP ACL rules.
        """
        return pulumi.get(self, "ip_rules")

    @ip_rules.setter
    def ip_rules(self, value: Optional[pulumi.Input[List[pulumi.Input['IPRuleArgs']]]]):
        pulumi.set(self, "ip_rules", value)

    @property
    @pulumi.getter(name="virtualNetworkRules")
    def virtual_network_rules(self) -> Optional[pulumi.Input[List[pulumi.Input['VirtualNetworkRuleArgs']]]]:
        """
        The virtual network rules.
        """
        return pulumi.get(self, "virtual_network_rules")

    @virtual_network_rules.setter
    def virtual_network_rules(self, value: Optional[pulumi.Input[List[pulumi.Input['VirtualNetworkRuleArgs']]]]):
        pulumi.set(self, "virtual_network_rules", value)


@pulumi.input_type
class PoliciesArgs:
    def __init__(__self__, *,
                 quarantine_policy: Optional[pulumi.Input['QuarantinePolicyArgs']] = None,
                 retention_policy: Optional[pulumi.Input['RetentionPolicyArgs']] = None,
                 trust_policy: Optional[pulumi.Input['TrustPolicyArgs']] = None):
        """
        The policies for a container registry.
        :param pulumi.Input['QuarantinePolicyArgs'] quarantine_policy: The quarantine policy for a container registry.
        :param pulumi.Input['RetentionPolicyArgs'] retention_policy: The retention policy for a container registry.
        :param pulumi.Input['TrustPolicyArgs'] trust_policy: The content trust policy for a container registry.
        """
        if quarantine_policy is not None:
            pulumi.set(__self__, "quarantine_policy", quarantine_policy)
        if retention_policy is not None:
            pulumi.set(__self__, "retention_policy", retention_policy)
        if trust_policy is not None:
            pulumi.set(__self__, "trust_policy", trust_policy)

    @property
    @pulumi.getter(name="quarantinePolicy")
    def quarantine_policy(self) -> Optional[pulumi.Input['QuarantinePolicyArgs']]:
        """
        The quarantine policy for a container registry.
        """
        return pulumi.get(self, "quarantine_policy")

    @quarantine_policy.setter
    def quarantine_policy(self, value: Optional[pulumi.Input['QuarantinePolicyArgs']]):
        pulumi.set(self, "quarantine_policy", value)

    @property
    @pulumi.getter(name="retentionPolicy")
    def retention_policy(self) -> Optional[pulumi.Input['RetentionPolicyArgs']]:
        """
        The retention policy for a container registry.
        """
        return pulumi.get(self, "retention_policy")

    @retention_policy.setter
    def retention_policy(self, value: Optional[pulumi.Input['RetentionPolicyArgs']]):
        pulumi.set(self, "retention_policy", value)

    @property
    @pulumi.getter(name="trustPolicy")
    def trust_policy(self) -> Optional[pulumi.Input['TrustPolicyArgs']]:
        """
        The content trust policy for a container registry.
        """
        return pulumi.get(self, "trust_policy")

    @trust_policy.setter
    def trust_policy(self, value: Optional[pulumi.Input['TrustPolicyArgs']]):
        pulumi.set(self, "trust_policy", value)


@pulumi.input_type
class QuarantinePolicyArgs:
    def __init__(__self__, *,
                 status: Optional[pulumi.Input[str]] = None):
        """
        The quarantine policy for a container registry.
        :param pulumi.Input[str] status: The value that indicates whether the policy is enabled or not.
        """
        if status is not None:
            pulumi.set(__self__, "status", status)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[str]]:
        """
        The value that indicates whether the policy is enabled or not.
        """
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "status", value)


@pulumi.input_type
class RetentionPolicyArgs:
    def __init__(__self__, *,
                 days: Optional[pulumi.Input[float]] = None,
                 status: Optional[pulumi.Input[str]] = None):
        """
        The retention policy for a container registry.
        :param pulumi.Input[float] days: The number of days to retain an untagged manifest after which it gets purged.
        :param pulumi.Input[str] status: The value that indicates whether the policy is enabled or not.
        """
        if days is not None:
            pulumi.set(__self__, "days", days)
        if status is not None:
            pulumi.set(__self__, "status", status)

    @property
    @pulumi.getter
    def days(self) -> Optional[pulumi.Input[float]]:
        """
        The number of days to retain an untagged manifest after which it gets purged.
        """
        return pulumi.get(self, "days")

    @days.setter
    def days(self, value: Optional[pulumi.Input[float]]):
        pulumi.set(self, "days", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[str]]:
        """
        The value that indicates whether the policy is enabled or not.
        """
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "status", value)


@pulumi.input_type
class SkuArgs:
    def __init__(__self__, *,
                 name: pulumi.Input[str]):
        """
        The SKU of a container registry.
        :param pulumi.Input[str] name: The SKU name of the container registry. Required for registry creation.
        """
        pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def name(self) -> pulumi.Input[str]:
        """
        The SKU name of the container registry. Required for registry creation.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: pulumi.Input[str]):
        pulumi.set(self, "name", value)


@pulumi.input_type
class StorageAccountPropertiesArgs:
    def __init__(__self__, *,
                 id: pulumi.Input[str]):
        """
        The properties of a storage account for a container registry. Only applicable to Classic SKU.
        :param pulumi.Input[str] id: The resource ID of the storage account.
        """
        pulumi.set(__self__, "id", id)

    @property
    @pulumi.getter
    def id(self) -> pulumi.Input[str]:
        """
        The resource ID of the storage account.
        """
        return pulumi.get(self, "id")

    @id.setter
    def id(self, value: pulumi.Input[str]):
        pulumi.set(self, "id", value)


@pulumi.input_type
class TrustPolicyArgs:
    def __init__(__self__, *,
                 status: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None):
        """
        The content trust policy for a container registry.
        :param pulumi.Input[str] status: The value that indicates whether the policy is enabled or not.
        :param pulumi.Input[str] type: The type of trust policy.
        """
        if status is not None:
            pulumi.set(__self__, "status", status)
        if type is not None:
            pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[str]]:
        """
        The value that indicates whether the policy is enabled or not.
        """
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "status", value)

    @property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input[str]]:
        """
        The type of trust policy.
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "type", value)


@pulumi.input_type
class VirtualNetworkRuleArgs:
    def __init__(__self__, *,
                 virtual_network_resource_id: pulumi.Input[str],
                 action: Optional[pulumi.Input[str]] = None):
        """
        Virtual network rule.
        :param pulumi.Input[str] virtual_network_resource_id: Resource ID of a subnet, for example: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworks/{vnetName}/subnets/{subnetName}.
        :param pulumi.Input[str] action: The action of virtual network rule.
        """
        pulumi.set(__self__, "virtual_network_resource_id", virtual_network_resource_id)
        if action is not None:
            pulumi.set(__self__, "action", action)

    @property
    @pulumi.getter(name="virtualNetworkResourceId")
    def virtual_network_resource_id(self) -> pulumi.Input[str]:
        """
        Resource ID of a subnet, for example: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworks/{vnetName}/subnets/{subnetName}.
        """
        return pulumi.get(self, "virtual_network_resource_id")

    @virtual_network_resource_id.setter
    def virtual_network_resource_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "virtual_network_resource_id", value)

    @property
    @pulumi.getter
    def action(self) -> Optional[pulumi.Input[str]]:
        """
        The action of virtual network rule.
        """
        return pulumi.get(self, "action")

    @action.setter
    def action(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "action", value)

