# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from ... import _utilities, _tables

__all__ = [
    'CacheHealthResponse',
    'CacheResponseSku',
    'CacheUpgradeStatusResponse',
    'ClfsTargetResponse',
    'NamespaceJunctionResponse',
    'Nfs3TargetResponse',
    'UnknownTargetResponse',
]

@pulumi.output_type
class CacheHealthResponse(dict):
    """
    An indication of Cache health. Gives more information about health than just that related to provisioning.
    """
    def __init__(__self__, *,
                 state: Optional[str] = None,
                 status_description: Optional[str] = None):
        """
        An indication of Cache health. Gives more information about health than just that related to provisioning.
        :param str state: List of Cache health states.
        :param str status_description: Describes explanation of state.
        """
        if state is not None:
            pulumi.set(__self__, "state", state)
        if status_description is not None:
            pulumi.set(__self__, "status_description", status_description)

    @property
    @pulumi.getter
    def state(self) -> Optional[str]:
        """
        List of Cache health states.
        """
        return pulumi.get(self, "state")

    @property
    @pulumi.getter(name="statusDescription")
    def status_description(self) -> Optional[str]:
        """
        Describes explanation of state.
        """
        return pulumi.get(self, "status_description")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class CacheResponseSku(dict):
    """
    SKU for the Cache.
    """
    def __init__(__self__, *,
                 name: Optional[str] = None):
        """
        SKU for the Cache.
        :param str name: SKU name for this Cache.
        """
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        """
        SKU name for this Cache.
        """
        return pulumi.get(self, "name")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class CacheUpgradeStatusResponse(dict):
    """
    Properties describing the software upgrade state of the Cache.
    """
    def __init__(__self__, *,
                 current_firmware_version: str,
                 firmware_update_deadline: str,
                 firmware_update_status: str,
                 last_firmware_update: str,
                 pending_firmware_version: str):
        """
        Properties describing the software upgrade state of the Cache.
        :param str current_firmware_version: Version string of the firmware currently installed on this Cache.
        :param str firmware_update_deadline: Time at which the pending firmware update will automatically be installed on the Cache.
        :param str firmware_update_status: True if there is a firmware update ready to install on this Cache. The firmware will automatically be installed after firmwareUpdateDeadline if not triggered earlier via the upgrade operation.
        :param str last_firmware_update: Time of the last successful firmware update.
        :param str pending_firmware_version: When firmwareUpdateAvailable is true, this field holds the version string for the update.
        """
        pulumi.set(__self__, "current_firmware_version", current_firmware_version)
        pulumi.set(__self__, "firmware_update_deadline", firmware_update_deadline)
        pulumi.set(__self__, "firmware_update_status", firmware_update_status)
        pulumi.set(__self__, "last_firmware_update", last_firmware_update)
        pulumi.set(__self__, "pending_firmware_version", pending_firmware_version)

    @property
    @pulumi.getter(name="currentFirmwareVersion")
    def current_firmware_version(self) -> str:
        """
        Version string of the firmware currently installed on this Cache.
        """
        return pulumi.get(self, "current_firmware_version")

    @property
    @pulumi.getter(name="firmwareUpdateDeadline")
    def firmware_update_deadline(self) -> str:
        """
        Time at which the pending firmware update will automatically be installed on the Cache.
        """
        return pulumi.get(self, "firmware_update_deadline")

    @property
    @pulumi.getter(name="firmwareUpdateStatus")
    def firmware_update_status(self) -> str:
        """
        True if there is a firmware update ready to install on this Cache. The firmware will automatically be installed after firmwareUpdateDeadline if not triggered earlier via the upgrade operation.
        """
        return pulumi.get(self, "firmware_update_status")

    @property
    @pulumi.getter(name="lastFirmwareUpdate")
    def last_firmware_update(self) -> str:
        """
        Time of the last successful firmware update.
        """
        return pulumi.get(self, "last_firmware_update")

    @property
    @pulumi.getter(name="pendingFirmwareVersion")
    def pending_firmware_version(self) -> str:
        """
        When firmwareUpdateAvailable is true, this field holds the version string for the update.
        """
        return pulumi.get(self, "pending_firmware_version")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class ClfsTargetResponse(dict):
    """
    Storage container for use as a CLFS Storage Target.
    """
    def __init__(__self__, *,
                 target: Optional[str] = None):
        """
        Storage container for use as a CLFS Storage Target.
        :param str target: Resource ID of storage container.
        """
        if target is not None:
            pulumi.set(__self__, "target", target)

    @property
    @pulumi.getter
    def target(self) -> Optional[str]:
        """
        Resource ID of storage container.
        """
        return pulumi.get(self, "target")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class NamespaceJunctionResponse(dict):
    """
    A namespace junction.
    """
    def __init__(__self__, *,
                 namespace_path: Optional[str] = None,
                 nfs_export: Optional[str] = None,
                 target_path: Optional[str] = None):
        """
        A namespace junction.
        :param str namespace_path: Namespace path on a Cache for a Storage Target.
        :param str nfs_export: NFS export where targetPath exists.
        :param str target_path: Path in Storage Target to which namespacePath points.
        """
        if namespace_path is not None:
            pulumi.set(__self__, "namespace_path", namespace_path)
        if nfs_export is not None:
            pulumi.set(__self__, "nfs_export", nfs_export)
        if target_path is not None:
            pulumi.set(__self__, "target_path", target_path)

    @property
    @pulumi.getter(name="namespacePath")
    def namespace_path(self) -> Optional[str]:
        """
        Namespace path on a Cache for a Storage Target.
        """
        return pulumi.get(self, "namespace_path")

    @property
    @pulumi.getter(name="nfsExport")
    def nfs_export(self) -> Optional[str]:
        """
        NFS export where targetPath exists.
        """
        return pulumi.get(self, "nfs_export")

    @property
    @pulumi.getter(name="targetPath")
    def target_path(self) -> Optional[str]:
        """
        Path in Storage Target to which namespacePath points.
        """
        return pulumi.get(self, "target_path")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class Nfs3TargetResponse(dict):
    """
    An NFSv3 mount point for use as a Storage Target.
    """
    def __init__(__self__, *,
                 target: Optional[str] = None,
                 usage_model: Optional[str] = None):
        """
        An NFSv3 mount point for use as a Storage Target.
        :param str target: IP address or host name of an NFSv3 host (e.g., 10.0.44.44).
        :param str usage_model: Identifies the primary usage model to be used for this Storage Target. Get choices from .../usageModels
        """
        if target is not None:
            pulumi.set(__self__, "target", target)
        if usage_model is not None:
            pulumi.set(__self__, "usage_model", usage_model)

    @property
    @pulumi.getter
    def target(self) -> Optional[str]:
        """
        IP address or host name of an NFSv3 host (e.g., 10.0.44.44).
        """
        return pulumi.get(self, "target")

    @property
    @pulumi.getter(name="usageModel")
    def usage_model(self) -> Optional[str]:
        """
        Identifies the primary usage model to be used for this Storage Target. Get choices from .../usageModels
        """
        return pulumi.get(self, "usage_model")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class UnknownTargetResponse(dict):
    """
    Storage container for use as an Unknown Storage Target.
    """
    def __init__(__self__, *,
                 unknown_map: Optional[Mapping[str, str]] = None):
        """
        Storage container for use as an Unknown Storage Target.
        :param Mapping[str, str] unknown_map: Dictionary of string->string pairs containing information about the Storage Target.
        """
        if unknown_map is not None:
            pulumi.set(__self__, "unknown_map", unknown_map)

    @property
    @pulumi.getter(name="unknownMap")
    def unknown_map(self) -> Optional[Mapping[str, str]]:
        """
        Dictionary of string->string pairs containing information about the Storage Target.
        """
        return pulumi.get(self, "unknown_map")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

