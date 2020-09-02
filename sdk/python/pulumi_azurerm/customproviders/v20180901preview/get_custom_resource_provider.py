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
    'GetCustomResourceProviderResult',
    'AwaitableGetCustomResourceProviderResult',
    'get_custom_resource_provider',
]

@pulumi.output_type
class GetCustomResourceProviderResult:
    """
    A manifest file that defines the custom resource provider resources.
    """
    def __init__(__self__, actions=None, location=None, name=None, provisioning_state=None, resource_types=None, tags=None, type=None, validations=None):
        if actions and not isinstance(actions, list):
            raise TypeError("Expected argument 'actions' to be a list")
        pulumi.set(__self__, "actions", actions)
        if location and not isinstance(location, str):
            raise TypeError("Expected argument 'location' to be a str")
        pulumi.set(__self__, "location", location)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if provisioning_state and not isinstance(provisioning_state, str):
            raise TypeError("Expected argument 'provisioning_state' to be a str")
        pulumi.set(__self__, "provisioning_state", provisioning_state)
        if resource_types and not isinstance(resource_types, list):
            raise TypeError("Expected argument 'resource_types' to be a list")
        pulumi.set(__self__, "resource_types", resource_types)
        if tags and not isinstance(tags, dict):
            raise TypeError("Expected argument 'tags' to be a dict")
        pulumi.set(__self__, "tags", tags)
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        pulumi.set(__self__, "type", type)
        if validations and not isinstance(validations, list):
            raise TypeError("Expected argument 'validations' to be a list")
        pulumi.set(__self__, "validations", validations)

    @property
    @pulumi.getter
    def actions(self) -> Optional[List['outputs.CustomRPActionRouteDefinitionResponse']]:
        """
        A list of actions that the custom resource provider implements.
        """
        return pulumi.get(self, "actions")

    @property
    @pulumi.getter
    def location(self) -> str:
        """
        Resource location
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Resource name
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="provisioningState")
    def provisioning_state(self) -> str:
        """
        The provisioning state of the resource provider.
        """
        return pulumi.get(self, "provisioning_state")

    @property
    @pulumi.getter(name="resourceTypes")
    def resource_types(self) -> Optional[List['outputs.CustomRPResourceTypeRouteDefinitionResponse']]:
        """
        A list of resource types that the custom resource provider implements.
        """
        return pulumi.get(self, "resource_types")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Mapping[str, str]]:
        """
        Resource tags
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        Resource type
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter
    def validations(self) -> Optional[List['outputs.CustomRPValidationsResponse']]:
        """
        A list of validations to run on the custom resource provider's requests.
        """
        return pulumi.get(self, "validations")


class AwaitableGetCustomResourceProviderResult(GetCustomResourceProviderResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetCustomResourceProviderResult(
            actions=self.actions,
            location=self.location,
            name=self.name,
            provisioning_state=self.provisioning_state,
            resource_types=self.resource_types,
            tags=self.tags,
            type=self.type,
            validations=self.validations)


def get_custom_resource_provider(resource_group_name: Optional[str] = None,
                                 resource_provider_name: Optional[str] = None,
                                 opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetCustomResourceProviderResult:
    """
    Use this data source to access information about an existing resource.

    :param str resource_group_name: The name of the resource group.
    :param str resource_provider_name: The name of the resource provider.
    """
    __args__ = dict()
    __args__['resourceGroupName'] = resource_group_name
    __args__['resourceProviderName'] = resource_provider_name
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('azurerm:customproviders/v20180901preview:getCustomResourceProvider', __args__, opts=opts, typ=GetCustomResourceProviderResult).value

    return AwaitableGetCustomResourceProviderResult(
        actions=__ret__.actions,
        location=__ret__.location,
        name=__ret__.name,
        provisioning_state=__ret__.provisioning_state,
        resource_types=__ret__.resource_types,
        tags=__ret__.tags,
        type=__ret__.type,
        validations=__ret__.validations)