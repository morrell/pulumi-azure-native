# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

from enum import Enum

__all__ = [
    'MaintenanceScope',
    'RebootOptions',
    'TaskScope',
    'Visibility',
]


class MaintenanceScope(str, Enum):
    """
    Gets or sets maintenanceScope of the configuration
    """
    HOST = "Host"
    OS_IMAGE = "OSImage"
    EXTENSION = "Extension"
    IN_GUEST_PATCH = "InGuestPatch"
    SQLDB = "SQLDB"
    SQL_MANAGED_INSTANCE = "SQLManagedInstance"


class RebootOptions(str, Enum):
    """
    Possible reboot preference as defined by the user based on which it would be decided to reboot the machine or not after the patch operation is completed.
    """
    NEVER_REBOOT = "NeverReboot"
    REBOOT_IF_REQUIRED = "RebootIfRequired"
    ALWAYS_REBOOT = "AlwaysReboot"


class TaskScope(str, Enum):
    """
    Global Task execute once when schedule trigger. Resource task execute for each VM.
    """
    GLOBAL_ = "Global"
    RESOURCE = "Resource"


class Visibility(str, Enum):
    """
    Gets or sets the visibility of the configuration
    """
    CUSTOM = "Custom"
    PUBLIC = "Public"
