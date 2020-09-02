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
    'GetIntegrationAccountMapResult',
    'AwaitableGetIntegrationAccountMapResult',
    'get_integration_account_map',
]

@pulumi.output_type
class GetIntegrationAccountMapResult:
    def __init__(__self__, changed_time=None, content=None, content_link=None, content_type=None, created_time=None, location=None, map_type=None, metadata=None, name=None, tags=None, type=None):
        if changed_time and not isinstance(changed_time, str):
            raise TypeError("Expected argument 'changed_time' to be a str")
        pulumi.set(__self__, "changed_time", changed_time)
        if content and not isinstance(content, dict):
            raise TypeError("Expected argument 'content' to be a dict")
        pulumi.set(__self__, "content", content)
        if content_link and not isinstance(content_link, dict):
            raise TypeError("Expected argument 'content_link' to be a dict")
        pulumi.set(__self__, "content_link", content_link)
        if content_type and not isinstance(content_type, str):
            raise TypeError("Expected argument 'content_type' to be a str")
        pulumi.set(__self__, "content_type", content_type)
        if created_time and not isinstance(created_time, str):
            raise TypeError("Expected argument 'created_time' to be a str")
        pulumi.set(__self__, "created_time", created_time)
        if location and not isinstance(location, str):
            raise TypeError("Expected argument 'location' to be a str")
        pulumi.set(__self__, "location", location)
        if map_type and not isinstance(map_type, str):
            raise TypeError("Expected argument 'map_type' to be a str")
        pulumi.set(__self__, "map_type", map_type)
        if metadata and not isinstance(metadata, dict):
            raise TypeError("Expected argument 'metadata' to be a dict")
        pulumi.set(__self__, "metadata", metadata)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if tags and not isinstance(tags, dict):
            raise TypeError("Expected argument 'tags' to be a dict")
        pulumi.set(__self__, "tags", tags)
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter(name="changedTime")
    def changed_time(self) -> str:
        """
        The changed time.
        """
        return pulumi.get(self, "changed_time")

    @property
    @pulumi.getter
    def content(self) -> Optional[Mapping[str, Any]]:
        """
        The content.
        """
        return pulumi.get(self, "content")

    @property
    @pulumi.getter(name="contentLink")
    def content_link(self) -> 'outputs.IntegrationAccountContentLinkResponse':
        """
        The content link.
        """
        return pulumi.get(self, "content_link")

    @property
    @pulumi.getter(name="contentType")
    def content_type(self) -> Optional[str]:
        """
        The content type.
        """
        return pulumi.get(self, "content_type")

    @property
    @pulumi.getter(name="createdTime")
    def created_time(self) -> str:
        """
        The created time.
        """
        return pulumi.get(self, "created_time")

    @property
    @pulumi.getter
    def location(self) -> Optional[str]:
        """
        The resource location.
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter(name="mapType")
    def map_type(self) -> Optional[str]:
        """
        The map type.
        """
        return pulumi.get(self, "map_type")

    @property
    @pulumi.getter
    def metadata(self) -> Optional[Mapping[str, Any]]:
        """
        The metadata.
        """
        return pulumi.get(self, "metadata")

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        """
        The resource name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Mapping[str, str]]:
        """
        The resource tags.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def type(self) -> Optional[str]:
        """
        The resource type.
        """
        return pulumi.get(self, "type")


class AwaitableGetIntegrationAccountMapResult(GetIntegrationAccountMapResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetIntegrationAccountMapResult(
            changed_time=self.changed_time,
            content=self.content,
            content_link=self.content_link,
            content_type=self.content_type,
            created_time=self.created_time,
            location=self.location,
            map_type=self.map_type,
            metadata=self.metadata,
            name=self.name,
            tags=self.tags,
            type=self.type)


def get_integration_account_map(integration_account_name: Optional[str] = None,
                                map_name: Optional[str] = None,
                                resource_group_name: Optional[str] = None,
                                opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetIntegrationAccountMapResult:
    """
    Use this data source to access information about an existing resource.

    :param str integration_account_name: The integration account name.
    :param str map_name: The integration account map name.
    :param str resource_group_name: The resource group name.
    """
    __args__ = dict()
    __args__['integrationAccountName'] = integration_account_name
    __args__['mapName'] = map_name
    __args__['resourceGroupName'] = resource_group_name
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('azurerm:logic/v20150801preview:getIntegrationAccountMap', __args__, opts=opts, typ=GetIntegrationAccountMapResult).value

    return AwaitableGetIntegrationAccountMapResult(
        changed_time=__ret__.changed_time,
        content=__ret__.content,
        content_link=__ret__.content_link,
        content_type=__ret__.content_type,
        created_time=__ret__.created_time,
        location=__ret__.location,
        map_type=__ret__.map_type,
        metadata=__ret__.metadata,
        name=__ret__.name,
        tags=__ret__.tags,
        type=__ret__.type)