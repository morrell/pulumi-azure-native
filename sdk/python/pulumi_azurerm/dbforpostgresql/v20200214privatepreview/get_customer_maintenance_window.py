# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from ... import _utilities, _tables

__all__ = [
    'GetCustomerMaintenanceWindowResult',
    'AwaitableGetCustomerMaintenanceWindowResult',
    'get_customer_maintenance_window',
]

@pulumi.output_type
class GetCustomerMaintenanceWindowResult:
    """
    Represents a server firewall rule.
    """
    def __init__(__self__, day_of_week=None, duration_in_minutes=None, name=None, start_hour=None, start_minute=None, type=None):
        if day_of_week and not isinstance(day_of_week, float):
            raise TypeError("Expected argument 'day_of_week' to be a float")
        pulumi.set(__self__, "day_of_week", day_of_week)
        if duration_in_minutes and not isinstance(duration_in_minutes, float):
            raise TypeError("Expected argument 'duration_in_minutes' to be a float")
        pulumi.set(__self__, "duration_in_minutes", duration_in_minutes)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if start_hour and not isinstance(start_hour, float):
            raise TypeError("Expected argument 'start_hour' to be a float")
        pulumi.set(__self__, "start_hour", start_hour)
        if start_minute and not isinstance(start_minute, float):
            raise TypeError("Expected argument 'start_minute' to be a float")
        pulumi.set(__self__, "start_minute", start_minute)
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter(name="dayOfWeek")
    def day_of_week(self) -> float:
        """
        The day of week of the customer maintenance window to start
        """
        return pulumi.get(self, "day_of_week")

    @property
    @pulumi.getter(name="durationInMinutes")
    def duration_in_minutes(self) -> Optional[float]:
        """
        The duration of the customer maintenance window to run.
        """
        return pulumi.get(self, "duration_in_minutes")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The name of the resource
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="startHour")
    def start_hour(self) -> float:
        """
        The starting hour of the customer maintenance window.
        """
        return pulumi.get(self, "start_hour")

    @property
    @pulumi.getter(name="startMinute")
    def start_minute(self) -> float:
        """
        The starting minutes of the customer maintenance window.
        """
        return pulumi.get(self, "start_minute")

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
        """
        return pulumi.get(self, "type")


class AwaitableGetCustomerMaintenanceWindowResult(GetCustomerMaintenanceWindowResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetCustomerMaintenanceWindowResult(
            day_of_week=self.day_of_week,
            duration_in_minutes=self.duration_in_minutes,
            name=self.name,
            start_hour=self.start_hour,
            start_minute=self.start_minute,
            type=self.type)


def get_customer_maintenance_window(maintenance_window_name: Optional[str] = None,
                                    resource_group_name: Optional[str] = None,
                                    server_name: Optional[str] = None,
                                    opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetCustomerMaintenanceWindowResult:
    """
    Use this data source to access information about an existing resource.

    :param str maintenance_window_name: The name of the maintenance window.
    :param str resource_group_name: The name of the resource group. The name is case insensitive.
    :param str server_name: The name of the server.
    """
    __args__ = dict()
    __args__['maintenanceWindowName'] = maintenance_window_name
    __args__['resourceGroupName'] = resource_group_name
    __args__['serverName'] = server_name
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('azurerm:dbforpostgresql/v20200214privatepreview:getCustomerMaintenanceWindow', __args__, opts=opts, typ=GetCustomerMaintenanceWindowResult).value

    return AwaitableGetCustomerMaintenanceWindowResult(
        day_of_week=__ret__.day_of_week,
        duration_in_minutes=__ret__.duration_in_minutes,
        name=__ret__.name,
        start_hour=__ret__.start_hour,
        start_minute=__ret__.start_minute,
        type=__ret__.type)