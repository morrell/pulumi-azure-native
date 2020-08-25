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
    'GetPolicyDefinitionResult',
    'AwaitableGetPolicyDefinitionResult',
    'get_policy_definition',
]

@pulumi.output_type
class GetPolicyDefinitionResult:
    """
    Policy definition.
    """
    def __init__(__self__, name=None, properties=None):
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if properties and not isinstance(properties, dict):
            raise TypeError("Expected argument 'properties' to be a dict")
        pulumi.set(__self__, "properties", properties)

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        """
        Gets or sets the policy definition name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def properties(self) -> 'outputs.PolicyDefinitionPropertiesResponse':
        """
        Gets or sets the policy definition properties.
        """
        return pulumi.get(self, "properties")


class AwaitableGetPolicyDefinitionResult(GetPolicyDefinitionResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetPolicyDefinitionResult(
            name=self.name,
            properties=self.properties)


def get_policy_definition(name: Optional[str] = None,
                          opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetPolicyDefinitionResult:
    """
    Use this data source to access information about an existing resource.

    :param str name: The policy definition name.
    """
    __args__ = dict()
    __args__['name'] = name
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('azurerm:authorization/v20151101:getPolicyDefinition', __args__, opts=opts, typ=GetPolicyDefinitionResult).value

    return AwaitableGetPolicyDefinitionResult(
        name=__ret__.name,
        properties=__ret__.properties)