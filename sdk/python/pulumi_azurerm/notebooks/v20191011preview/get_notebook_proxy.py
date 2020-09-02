# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from ... import _utilities, _tables

__all__ = [
    'GetNotebookProxyResult',
    'AwaitableGetNotebookProxyResult',
    'get_notebook_proxy',
]

@pulumi.output_type
class GetNotebookProxyResult:
    """
    A NotebookProxy resource.
    """
    def __init__(__self__, hostname=None, name=None, resource_id=None, type=None):
        if hostname and not isinstance(hostname, str):
            raise TypeError("Expected argument 'hostname' to be a str")
        pulumi.set(__self__, "hostname", hostname)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if resource_id and not isinstance(resource_id, str):
            raise TypeError("Expected argument 'resource_id' to be a str")
        pulumi.set(__self__, "resource_id", resource_id)
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter
    def hostname(self) -> Optional[str]:
        """
        The friendly string identifier of the creator of the NotebookProxy resource.
        """
        return pulumi.get(self, "hostname")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The name of the resource
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="resourceId")
    def resource_id(self) -> str:
        """
        The unique identifier (a GUID) generated for every resource.
        """
        return pulumi.get(self, "resource_id")

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        The type of the resource. Ex- Microsoft.Storage/storageAccounts or Microsoft.Notebooks/notebookProxies.
        """
        return pulumi.get(self, "type")


class AwaitableGetNotebookProxyResult(GetNotebookProxyResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetNotebookProxyResult(
            hostname=self.hostname,
            name=self.name,
            resource_id=self.resource_id,
            type=self.type)


def get_notebook_proxy(resource_group_name: Optional[str] = None,
                       resource_name: Optional[str] = None,
                       opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetNotebookProxyResult:
    """
    Use this data source to access information about an existing resource.

    :param str resource_group_name: The name of the resource group. The name is case insensitive.
    :param str resource_name: The name of the resource.
    """
    __args__ = dict()
    __args__['resourceGroupName'] = resource_group_name
    __args__['resourceName'] = resource_name
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('azurerm:notebooks/v20191011preview:getNotebookProxy', __args__, opts=opts, typ=GetNotebookProxyResult).value

    return AwaitableGetNotebookProxyResult(
        hostname=__ret__.hostname,
        name=__ret__.name,
        resource_id=__ret__.resource_id,
        type=__ret__.type)