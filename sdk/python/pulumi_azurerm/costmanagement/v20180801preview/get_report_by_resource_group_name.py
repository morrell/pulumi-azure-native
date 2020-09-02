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
    'GetReportByResourceGroupNameResult',
    'AwaitableGetReportByResourceGroupNameResult',
    'get_report_by_resource_group_name',
]

@pulumi.output_type
class GetReportByResourceGroupNameResult:
    """
    A report resource.
    """
    def __init__(__self__, definition=None, delivery_info=None, format=None, name=None, schedule=None, tags=None, type=None):
        if definition and not isinstance(definition, dict):
            raise TypeError("Expected argument 'definition' to be a dict")
        pulumi.set(__self__, "definition", definition)
        if delivery_info and not isinstance(delivery_info, dict):
            raise TypeError("Expected argument 'delivery_info' to be a dict")
        pulumi.set(__self__, "delivery_info", delivery_info)
        if format and not isinstance(format, str):
            raise TypeError("Expected argument 'format' to be a str")
        pulumi.set(__self__, "format", format)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if schedule and not isinstance(schedule, dict):
            raise TypeError("Expected argument 'schedule' to be a dict")
        pulumi.set(__self__, "schedule", schedule)
        if tags and not isinstance(tags, dict):
            raise TypeError("Expected argument 'tags' to be a dict")
        pulumi.set(__self__, "tags", tags)
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter
    def definition(self) -> 'outputs.ReportDefinitionResponse':
        """
        Has definition for the report.
        """
        return pulumi.get(self, "definition")

    @property
    @pulumi.getter(name="deliveryInfo")
    def delivery_info(self) -> 'outputs.ReportDeliveryInfoResponse':
        """
        Has delivery information for the report.
        """
        return pulumi.get(self, "delivery_info")

    @property
    @pulumi.getter
    def format(self) -> Optional[str]:
        """
        The format of the report being delivered.
        """
        return pulumi.get(self, "format")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Resource name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def schedule(self) -> Optional['outputs.ReportScheduleResponse']:
        """
        Has schedule information for the report.
        """
        return pulumi.get(self, "schedule")

    @property
    @pulumi.getter
    def tags(self) -> Mapping[str, str]:
        """
        Resource tags.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        Resource type.
        """
        return pulumi.get(self, "type")


class AwaitableGetReportByResourceGroupNameResult(GetReportByResourceGroupNameResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetReportByResourceGroupNameResult(
            definition=self.definition,
            delivery_info=self.delivery_info,
            format=self.format,
            name=self.name,
            schedule=self.schedule,
            tags=self.tags,
            type=self.type)


def get_report_by_resource_group_name(report_name: Optional[str] = None,
                                      resource_group_name: Optional[str] = None,
                                      opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetReportByResourceGroupNameResult:
    """
    Use this data source to access information about an existing resource.

    :param str report_name: Report Name.
    :param str resource_group_name: Azure Resource Group Name.
    """
    __args__ = dict()
    __args__['reportName'] = report_name
    __args__['resourceGroupName'] = resource_group_name
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('azurerm:costmanagement/v20180801preview:getReportByResourceGroupName', __args__, opts=opts, typ=GetReportByResourceGroupNameResult).value

    return AwaitableGetReportByResourceGroupNameResult(
        definition=__ret__.definition,
        delivery_info=__ret__.delivery_info,
        format=__ret__.format,
        name=__ret__.name,
        schedule=__ret__.schedule,
        tags=__ret__.tags,
        type=__ret__.type)