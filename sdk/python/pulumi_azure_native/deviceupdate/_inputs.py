# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from ._enums import *

__all__ = [
    'IdentityArgs',
    'IotHubSettingsArgs',
]

@pulumi.input_type
class IdentityArgs:
    def __init__(__self__, *,
                 type: Optional[pulumi.Input['ResourceIdentityType']] = None):
        """
        Identity for the resource.
        :param pulumi.Input['ResourceIdentityType'] type: The identity type.
        """
        if type is not None:
            pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input['ResourceIdentityType']]:
        """
        The identity type.
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[pulumi.Input['ResourceIdentityType']]):
        pulumi.set(self, "type", value)


@pulumi.input_type
class IotHubSettingsArgs:
    def __init__(__self__, *,
                 resource_id: pulumi.Input[str],
                 event_hub_connection_string: Optional[pulumi.Input[str]] = None,
                 io_t_hub_connection_string: Optional[pulumi.Input[str]] = None):
        """
        Device Update account integration with IoT Hub settings.
        :param pulumi.Input[str] resource_id: IoTHub resource ID
        :param pulumi.Input[str] event_hub_connection_string: EventHub connection string.
        :param pulumi.Input[str] io_t_hub_connection_string: IoTHub connection string.
        """
        pulumi.set(__self__, "resource_id", resource_id)
        if event_hub_connection_string is not None:
            pulumi.set(__self__, "event_hub_connection_string", event_hub_connection_string)
        if io_t_hub_connection_string is not None:
            pulumi.set(__self__, "io_t_hub_connection_string", io_t_hub_connection_string)

    @property
    @pulumi.getter(name="resourceId")
    def resource_id(self) -> pulumi.Input[str]:
        """
        IoTHub resource ID
        """
        return pulumi.get(self, "resource_id")

    @resource_id.setter
    def resource_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "resource_id", value)

    @property
    @pulumi.getter(name="eventHubConnectionString")
    def event_hub_connection_string(self) -> Optional[pulumi.Input[str]]:
        """
        EventHub connection string.
        """
        return pulumi.get(self, "event_hub_connection_string")

    @event_hub_connection_string.setter
    def event_hub_connection_string(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "event_hub_connection_string", value)

    @property
    @pulumi.getter(name="ioTHubConnectionString")
    def io_t_hub_connection_string(self) -> Optional[pulumi.Input[str]]:
        """
        IoTHub connection string.
        """
        return pulumi.get(self, "io_t_hub_connection_string")

    @io_t_hub_connection_string.setter
    def io_t_hub_connection_string(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "io_t_hub_connection_string", value)


