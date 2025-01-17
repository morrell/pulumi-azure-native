# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from ... import _utilities

__all__ = [
    'GetIotSensorResult',
    'AwaitableGetIotSensorResult',
    'get_iot_sensor',
]

@pulumi.output_type
class GetIotSensorResult:
    """
    IoT sensor model
    """
    def __init__(__self__, connectivity_time=None, creation_time=None, dynamic_learning=None, id=None, learning_mode=None, name=None, sensor_status=None, sensor_type=None, sensor_version=None, ti_automatic_updates=None, ti_status=None, ti_version=None, type=None, zone=None):
        if connectivity_time and not isinstance(connectivity_time, str):
            raise TypeError("Expected argument 'connectivity_time' to be a str")
        pulumi.set(__self__, "connectivity_time", connectivity_time)
        if creation_time and not isinstance(creation_time, str):
            raise TypeError("Expected argument 'creation_time' to be a str")
        pulumi.set(__self__, "creation_time", creation_time)
        if dynamic_learning and not isinstance(dynamic_learning, bool):
            raise TypeError("Expected argument 'dynamic_learning' to be a bool")
        pulumi.set(__self__, "dynamic_learning", dynamic_learning)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if learning_mode and not isinstance(learning_mode, bool):
            raise TypeError("Expected argument 'learning_mode' to be a bool")
        pulumi.set(__self__, "learning_mode", learning_mode)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if sensor_status and not isinstance(sensor_status, str):
            raise TypeError("Expected argument 'sensor_status' to be a str")
        pulumi.set(__self__, "sensor_status", sensor_status)
        if sensor_type and not isinstance(sensor_type, str):
            raise TypeError("Expected argument 'sensor_type' to be a str")
        pulumi.set(__self__, "sensor_type", sensor_type)
        if sensor_version and not isinstance(sensor_version, str):
            raise TypeError("Expected argument 'sensor_version' to be a str")
        pulumi.set(__self__, "sensor_version", sensor_version)
        if ti_automatic_updates and not isinstance(ti_automatic_updates, bool):
            raise TypeError("Expected argument 'ti_automatic_updates' to be a bool")
        pulumi.set(__self__, "ti_automatic_updates", ti_automatic_updates)
        if ti_status and not isinstance(ti_status, str):
            raise TypeError("Expected argument 'ti_status' to be a str")
        pulumi.set(__self__, "ti_status", ti_status)
        if ti_version and not isinstance(ti_version, str):
            raise TypeError("Expected argument 'ti_version' to be a str")
        pulumi.set(__self__, "ti_version", ti_version)
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        pulumi.set(__self__, "type", type)
        if zone and not isinstance(zone, str):
            raise TypeError("Expected argument 'zone' to be a str")
        pulumi.set(__self__, "zone", zone)

    @property
    @pulumi.getter(name="connectivityTime")
    def connectivity_time(self) -> str:
        """
        Last connectivity time of the IoT sensor
        """
        return pulumi.get(self, "connectivity_time")

    @property
    @pulumi.getter(name="creationTime")
    def creation_time(self) -> str:
        """
        Creation time of the IoT sensor
        """
        return pulumi.get(self, "creation_time")

    @property
    @pulumi.getter(name="dynamicLearning")
    def dynamic_learning(self) -> bool:
        """
        Dynamic mode status of the IoT sensor
        """
        return pulumi.get(self, "dynamic_learning")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        Resource Id
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="learningMode")
    def learning_mode(self) -> bool:
        """
        Learning mode status of the IoT sensor
        """
        return pulumi.get(self, "learning_mode")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Resource name
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="sensorStatus")
    def sensor_status(self) -> str:
        """
        Status of the IoT sensor
        """
        return pulumi.get(self, "sensor_status")

    @property
    @pulumi.getter(name="sensorType")
    def sensor_type(self) -> Optional[str]:
        """
        Type of sensor
        """
        return pulumi.get(self, "sensor_type")

    @property
    @pulumi.getter(name="sensorVersion")
    def sensor_version(self) -> str:
        """
        Version of the IoT sensor
        """
        return pulumi.get(self, "sensor_version")

    @property
    @pulumi.getter(name="tiAutomaticUpdates")
    def ti_automatic_updates(self) -> Optional[bool]:
        """
        TI Automatic mode status of the IoT sensor
        """
        return pulumi.get(self, "ti_automatic_updates")

    @property
    @pulumi.getter(name="tiStatus")
    def ti_status(self) -> str:
        """
        TI Status of the IoT sensor
        """
        return pulumi.get(self, "ti_status")

    @property
    @pulumi.getter(name="tiVersion")
    def ti_version(self) -> str:
        """
        TI Version of the IoT sensor
        """
        return pulumi.get(self, "ti_version")

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        Resource type
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter
    def zone(self) -> Optional[str]:
        """
        Zone of the IoT sensor
        """
        return pulumi.get(self, "zone")


class AwaitableGetIotSensorResult(GetIotSensorResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetIotSensorResult(
            connectivity_time=self.connectivity_time,
            creation_time=self.creation_time,
            dynamic_learning=self.dynamic_learning,
            id=self.id,
            learning_mode=self.learning_mode,
            name=self.name,
            sensor_status=self.sensor_status,
            sensor_type=self.sensor_type,
            sensor_version=self.sensor_version,
            ti_automatic_updates=self.ti_automatic_updates,
            ti_status=self.ti_status,
            ti_version=self.ti_version,
            type=self.type,
            zone=self.zone)


def get_iot_sensor(iot_sensor_name: Optional[str] = None,
                   scope: Optional[str] = None,
                   opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetIotSensorResult:
    """
    IoT sensor model


    :param str iot_sensor_name: Name of the IoT sensor
    :param str scope: Scope of the query (IoT Hub, /providers/Microsoft.Devices/iotHubs/myHub)
    """
    __args__ = dict()
    __args__['iotSensorName'] = iot_sensor_name
    __args__['scope'] = scope
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('azure-native:security/v20200806preview:getIotSensor', __args__, opts=opts, typ=GetIotSensorResult).value

    return AwaitableGetIotSensorResult(
        connectivity_time=__ret__.connectivity_time,
        creation_time=__ret__.creation_time,
        dynamic_learning=__ret__.dynamic_learning,
        id=__ret__.id,
        learning_mode=__ret__.learning_mode,
        name=__ret__.name,
        sensor_status=__ret__.sensor_status,
        sensor_type=__ret__.sensor_type,
        sensor_version=__ret__.sensor_version,
        ti_automatic_updates=__ret__.ti_automatic_updates,
        ti_status=__ret__.ti_status,
        ti_version=__ret__.ti_version,
        type=__ret__.type,
        zone=__ret__.zone)
