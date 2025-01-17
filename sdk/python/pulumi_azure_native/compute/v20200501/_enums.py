# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

from enum import Enum

__all__ = [
    'DiskCreateOption',
    'DiskEncryptionSetIdentityType',
    'DiskStorageAccountTypes',
    'EncryptionType',
    'HyperVGeneration',
    'NetworkAccessPolicy',
    'OperatingSystemTypes',
    'SnapshotStorageAccountTypes',
]


class DiskCreateOption(str, Enum):
    """
    This enumerates the possible sources of a disk's creation.
    """
    EMPTY = "Empty"
    """Create an empty data disk of a size given by diskSizeGB."""
    ATTACH = "Attach"
    """Disk will be attached to a VM."""
    FROM_IMAGE = "FromImage"
    """Create a new disk from a platform image specified by the given imageReference or galleryImageReference."""
    IMPORT_ = "Import"
    """Create a disk by importing from a blob specified by a sourceUri in a storage account specified by storageAccountId."""
    COPY = "Copy"
    """Create a new disk or snapshot by copying from a disk or snapshot specified by the given sourceResourceId."""
    RESTORE = "Restore"
    """Create a new disk by copying from a backup recovery point."""
    UPLOAD = "Upload"
    """Create a new disk by obtaining a write token and using it to directly upload the contents of the disk."""


class DiskEncryptionSetIdentityType(str, Enum):
    """
    The type of Managed Identity used by the DiskEncryptionSet. Only SystemAssigned is supported.
    """
    SYSTEM_ASSIGNED = "SystemAssigned"


class DiskStorageAccountTypes(str, Enum):
    """
    The sku name.
    """
    STANDARD_LRS = "Standard_LRS"
    """Standard HDD locally redundant storage. Best for backup, non-critical, and infrequent access."""
    PREMIUM_LRS = "Premium_LRS"
    """Premium SSD locally redundant storage. Best for production and performance sensitive workloads."""
    STANDARD_SS_D_LRS = "StandardSSD_LRS"
    """Standard SSD locally redundant storage. Best for web servers, lightly used enterprise applications and dev/test."""
    ULTRA_SS_D_LRS = "UltraSSD_LRS"
    """Ultra SSD locally redundant storage. Best for IO-intensive workloads such as SAP HANA, top tier databases (for example, SQL, Oracle), and other transaction-heavy workloads."""


class EncryptionType(str, Enum):
    """
    The type of key used to encrypt the data of the disk.
    """
    ENCRYPTION_AT_REST_WITH_PLATFORM_KEY = "EncryptionAtRestWithPlatformKey"
    """Disk is encrypted at rest with Platform managed key. It is the default encryption type. This is not a valid encryption type for disk encryption sets."""
    ENCRYPTION_AT_REST_WITH_CUSTOMER_KEY = "EncryptionAtRestWithCustomerKey"
    """Disk is encrypted at rest with Customer managed key that can be changed and revoked by a customer."""
    ENCRYPTION_AT_REST_WITH_PLATFORM_AND_CUSTOMER_KEYS = "EncryptionAtRestWithPlatformAndCustomerKeys"
    """Disk is encrypted at rest with 2 layers of encryption. One of the keys is Customer managed and the other key is Platform managed."""


class HyperVGeneration(str, Enum):
    """
    The hypervisor generation of the Virtual Machine. Applicable to OS disks only.
    """
    V1 = "V1"
    V2 = "V2"


class NetworkAccessPolicy(str, Enum):
    """
    Policy for accessing the disk via network.
    """
    ALLOW_ALL = "AllowAll"
    """The disk can be exported or uploaded to from any network."""
    ALLOW_PRIVATE = "AllowPrivate"
    """The disk can be exported or uploaded to using a DiskAccess resource's private endpoints."""
    DENY_ALL = "DenyAll"
    """The disk cannot be exported."""


class OperatingSystemTypes(str, Enum):
    """
    The Operating System type.
    """
    WINDOWS = "Windows"
    LINUX = "Linux"


class SnapshotStorageAccountTypes(str, Enum):
    """
    The sku name.
    """
    STANDARD_LRS = "Standard_LRS"
    """Standard HDD locally redundant storage"""
    PREMIUM_LRS = "Premium_LRS"
    """Premium SSD locally redundant storage"""
    STANDARD_ZRS = "Standard_ZRS"
    """Standard zone redundant storage"""
