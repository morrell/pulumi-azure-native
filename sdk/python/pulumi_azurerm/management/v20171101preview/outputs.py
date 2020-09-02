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
    'ManagementGroupChildInfoResponse',
    'ManagementGroupDetailsResponse',
    'ParentGroupInfoResponse',
]

@pulumi.output_type
class ManagementGroupChildInfoResponse(dict):
    """
    The child information of a management group.
    """
    def __init__(__self__, *,
                 child_id: Optional[str] = None,
                 child_type: Optional[str] = None,
                 children: Optional[List['outputs.ManagementGroupChildInfoResponse']] = None,
                 display_name: Optional[str] = None):
        """
        The child information of a management group.
        :param str child_id: The fully qualified ID for the child resource (management group or subscription).  For example, /providers/Microsoft.Management/managementGroups/0000000-0000-0000-0000-000000000000
        :param str child_type: The type of child resource.
        :param List['ManagementGroupChildInfoResponseArgs'] children: The list of children.
        :param str display_name: The friendly name of the child resource.
        """
        if child_id is not None:
            pulumi.set(__self__, "child_id", child_id)
        if child_type is not None:
            pulumi.set(__self__, "child_type", child_type)
        if children is not None:
            pulumi.set(__self__, "children", children)
        if display_name is not None:
            pulumi.set(__self__, "display_name", display_name)

    @property
    @pulumi.getter(name="childId")
    def child_id(self) -> Optional[str]:
        """
        The fully qualified ID for the child resource (management group or subscription).  For example, /providers/Microsoft.Management/managementGroups/0000000-0000-0000-0000-000000000000
        """
        return pulumi.get(self, "child_id")

    @property
    @pulumi.getter(name="childType")
    def child_type(self) -> Optional[str]:
        """
        The type of child resource.
        """
        return pulumi.get(self, "child_type")

    @property
    @pulumi.getter
    def children(self) -> Optional[List['outputs.ManagementGroupChildInfoResponse']]:
        """
        The list of children.
        """
        return pulumi.get(self, "children")

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> Optional[str]:
        """
        The friendly name of the child resource.
        """
        return pulumi.get(self, "display_name")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class ManagementGroupDetailsResponse(dict):
    """
    The details of a management group.
    """
    def __init__(__self__, *,
                 parent: Optional['outputs.ParentGroupInfoResponse'] = None,
                 updated_by: Optional[str] = None,
                 updated_time: Optional[str] = None,
                 version: Optional[float] = None):
        """
        The details of a management group.
        :param 'ParentGroupInfoResponseArgs' parent: (Optional) The ID of the parent management group.
        :param str updated_by: The identity of the principal or process that updated the object.
        :param str updated_time: The date and time when this object was last updated.
        :param float version: The version number of the object.
        """
        if parent is not None:
            pulumi.set(__self__, "parent", parent)
        if updated_by is not None:
            pulumi.set(__self__, "updated_by", updated_by)
        if updated_time is not None:
            pulumi.set(__self__, "updated_time", updated_time)
        if version is not None:
            pulumi.set(__self__, "version", version)

    @property
    @pulumi.getter
    def parent(self) -> Optional['outputs.ParentGroupInfoResponse']:
        """
        (Optional) The ID of the parent management group.
        """
        return pulumi.get(self, "parent")

    @property
    @pulumi.getter(name="updatedBy")
    def updated_by(self) -> Optional[str]:
        """
        The identity of the principal or process that updated the object.
        """
        return pulumi.get(self, "updated_by")

    @property
    @pulumi.getter(name="updatedTime")
    def updated_time(self) -> Optional[str]:
        """
        The date and time when this object was last updated.
        """
        return pulumi.get(self, "updated_time")

    @property
    @pulumi.getter
    def version(self) -> Optional[float]:
        """
        The version number of the object.
        """
        return pulumi.get(self, "version")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class ParentGroupInfoResponse(dict):
    """
    (Optional) The ID of the parent management group.
    """
    def __init__(__self__, *,
                 display_name: Optional[str] = None,
                 parent_id: Optional[str] = None):
        """
        (Optional) The ID of the parent management group.
        :param str display_name: The friendly name of the parent management group.
        :param str parent_id: The fully qualified ID for the parent management group.  For example, /providers/Microsoft.Management/managementGroups/0000000-0000-0000-0000-000000000000
        """
        if display_name is not None:
            pulumi.set(__self__, "display_name", display_name)
        if parent_id is not None:
            pulumi.set(__self__, "parent_id", parent_id)

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> Optional[str]:
        """
        The friendly name of the parent management group.
        """
        return pulumi.get(self, "display_name")

    @property
    @pulumi.getter(name="parentId")
    def parent_id(self) -> Optional[str]:
        """
        The fully qualified ID for the parent management group.  For example, /providers/Microsoft.Management/managementGroups/0000000-0000-0000-0000-000000000000
        """
        return pulumi.get(self, "parent_id")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

