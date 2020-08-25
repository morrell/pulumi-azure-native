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
    'GetResourceResult',
    'AwaitableGetResourceResult',
    'get_resource',
]

@pulumi.output_type
class GetResourceResult:
    """
    Resource information.
    """
    def __init__(__self__, identity=None, kind=None, location=None, managed_by=None, name=None, plan=None, properties=None, sku=None, tags=None, type=None):
        if identity and not isinstance(identity, dict):
            raise TypeError("Expected argument 'identity' to be a dict")
        pulumi.set(__self__, "identity", identity)
        if kind and not isinstance(kind, str):
            raise TypeError("Expected argument 'kind' to be a str")
        pulumi.set(__self__, "kind", kind)
        if location and not isinstance(location, str):
            raise TypeError("Expected argument 'location' to be a str")
        pulumi.set(__self__, "location", location)
        if managed_by and not isinstance(managed_by, str):
            raise TypeError("Expected argument 'managed_by' to be a str")
        pulumi.set(__self__, "managed_by", managed_by)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if plan and not isinstance(plan, dict):
            raise TypeError("Expected argument 'plan' to be a dict")
        pulumi.set(__self__, "plan", plan)
        if properties and not isinstance(properties, dict):
            raise TypeError("Expected argument 'properties' to be a dict")
        pulumi.set(__self__, "properties", properties)
        if sku and not isinstance(sku, dict):
            raise TypeError("Expected argument 'sku' to be a dict")
        pulumi.set(__self__, "sku", sku)
        if tags and not isinstance(tags, dict):
            raise TypeError("Expected argument 'tags' to be a dict")
        pulumi.set(__self__, "tags", tags)
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter
    def identity(self) -> Optional['outputs.IdentityResponse']:
        """
        The identity of the resource.
        """
        return pulumi.get(self, "identity")

    @property
    @pulumi.getter
    def kind(self) -> Optional[str]:
        """
        The kind of the resource.
        """
        return pulumi.get(self, "kind")

    @property
    @pulumi.getter
    def location(self) -> Optional[str]:
        """
        Resource location
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter(name="managedBy")
    def managed_by(self) -> Optional[str]:
        """
        ID of the resource that manages this resource.
        """
        return pulumi.get(self, "managed_by")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Resource name
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def plan(self) -> Optional['outputs.PlanResponse']:
        """
        The plan of the resource.
        """
        return pulumi.get(self, "plan")

    @property
    @pulumi.getter
    def properties(self) -> Mapping[str, Any]:
        """
        The resource properties.
        """
        return pulumi.get(self, "properties")

    @property
    @pulumi.getter
    def sku(self) -> Optional['outputs.SkuResponse']:
        """
        The SKU of the resource.
        """
        return pulumi.get(self, "sku")

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


class AwaitableGetResourceResult(GetResourceResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetResourceResult(
            identity=self.identity,
            kind=self.kind,
            location=self.location,
            managed_by=self.managed_by,
            name=self.name,
            plan=self.plan,
            properties=self.properties,
            sku=self.sku,
            tags=self.tags,
            type=self.type)


def get_resource(name: Optional[str] = None,
                 parent_resource_path: Optional[str] = None,
                 resource_group_name: Optional[str] = None,
                 resource_provider_namespace: Optional[str] = None,
                 resource_type: Optional[str] = None,
                 opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetResourceResult:
    """
    Use this data source to access information about an existing resource.

    :param str name: The name of the resource to get.
    :param str parent_resource_path: The parent resource identity.
    :param str resource_group_name: The name of the resource group containing the resource to get. The name is case insensitive.
    :param str resource_provider_namespace: The namespace of the resource provider.
    :param str resource_type: The resource type of the resource.
    """
    __args__ = dict()
    __args__['name'] = name
    __args__['parentResourcePath'] = parent_resource_path
    __args__['resourceGroupName'] = resource_group_name
    __args__['resourceProviderNamespace'] = resource_provider_namespace
    __args__['resourceType'] = resource_type
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('azurerm:resources/v20190301:getResource', __args__, opts=opts, typ=GetResourceResult).value

    return AwaitableGetResourceResult(
        identity=__ret__.identity,
        kind=__ret__.kind,
        location=__ret__.location,
        managed_by=__ret__.managed_by,
        name=__ret__.name,
        plan=__ret__.plan,
        properties=__ret__.properties,
        sku=__ret__.sku,
        tags=__ret__.tags,
        type=__ret__.type)