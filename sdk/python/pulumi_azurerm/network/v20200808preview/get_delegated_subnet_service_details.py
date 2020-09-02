# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from ... import _utilities, _tables

__all__ = [
    'GetDelegatedSubnetServiceDetailsResult',
    'AwaitableGetDelegatedSubnetServiceDetailsResult',
    'get_delegated_subnet_service_details',
]

@pulumi.output_type
class GetDelegatedSubnetServiceDetailsResult:
    """
    Delegated subnet details
    """
    def __init__(__self__, location=None, name=None, state=None, type=None):
        if location and not isinstance(location, str):
            raise TypeError("Expected argument 'location' to be a str")
        pulumi.set(__self__, "location", location)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if state and not isinstance(state, str):
            raise TypeError("Expected argument 'state' to be a str")
        pulumi.set(__self__, "state", state)
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter
    def location(self) -> Optional[str]:
        """
        Location of the DelegatedSubnet resource.
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The name of the DelegatedSubnet resource.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def state(self) -> str:
        """
        The current state of delegated subnet resource.
        """
        return pulumi.get(self, "state")

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        The type of the DelegatedSubnet  resource.(Microsoft.DelegatedNetwork/delegatedSubnet)
        """
        return pulumi.get(self, "type")


class AwaitableGetDelegatedSubnetServiceDetailsResult(GetDelegatedSubnetServiceDetailsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetDelegatedSubnetServiceDetailsResult(
            location=self.location,
            name=self.name,
            state=self.state,
            type=self.type)


def get_delegated_subnet_service_details(resource_group_name: Optional[str] = None,
                                         resource_name: Optional[str] = None,
                                         subnet_name: Optional[str] = None,
                                         vnet_name: Optional[str] = None,
                                         opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetDelegatedSubnetServiceDetailsResult:
    """
    Use this data source to access information about an existing resource.

    :param str resource_group_name: The name of the Azure Resource group of which a given DelegatedNetwork resource is part. This name must be at least 1 character in length, and no more than 90.
    :param str resource_name: The name of the resource. It must be a minimum of 3 characters, and a maximum of 63.
    :param str subnet_name: The name of the delegated subnet. This name must be at least 1 character in length, and no more than 90.
    :param str vnet_name: The name of the virtual network. This name must be at least 1 character in length, and no more than 90.
    """
    __args__ = dict()
    __args__['resourceGroupName'] = resource_group_name
    __args__['resourceName'] = resource_name
    __args__['subnetName'] = subnet_name
    __args__['vnetName'] = vnet_name
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('azurerm:network/v20200808preview:getDelegatedSubnetServiceDetails', __args__, opts=opts, typ=GetDelegatedSubnetServiceDetailsResult).value

    return AwaitableGetDelegatedSubnetServiceDetailsResult(
        location=__ret__.location,
        name=__ret__.name,
        state=__ret__.state,
        type=__ret__.type)