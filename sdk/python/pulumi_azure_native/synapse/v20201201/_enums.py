# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

from enum import Enum

__all__ = [
    'DataFlowComputeType',
    'IntegrationRuntimeEdition',
    'IntegrationRuntimeEntityReferenceType',
    'IntegrationRuntimeLicenseType',
    'IntegrationRuntimeSsisCatalogPricingTier',
    'IntegrationRuntimeType',
    'NodeSize',
    'NodeSizeFamily',
    'ResourceIdentityType',
    'SensitivityLabelRank',
    'StorageAccountType',
    'TransparentDataEncryptionStatus',
]


class DataFlowComputeType(str, Enum):
    """
    Compute type of the cluster which will execute data flow job.
    """
    GENERAL = "General"
    MEMORY_OPTIMIZED = "MemoryOptimized"
    COMPUTE_OPTIMIZED = "ComputeOptimized"


class IntegrationRuntimeEdition(str, Enum):
    """
    The edition for the SSIS Integration Runtime
    """
    STANDARD = "Standard"
    ENTERPRISE = "Enterprise"


class IntegrationRuntimeEntityReferenceType(str, Enum):
    """
    The type of this referenced entity.
    """
    INTEGRATION_RUNTIME_REFERENCE = "IntegrationRuntimeReference"
    LINKED_SERVICE_REFERENCE = "LinkedServiceReference"


class IntegrationRuntimeLicenseType(str, Enum):
    """
    License type for bringing your own license scenario.
    """
    BASE_PRICE = "BasePrice"
    LICENSE_INCLUDED = "LicenseIncluded"


class IntegrationRuntimeSsisCatalogPricingTier(str, Enum):
    """
    The pricing tier for the catalog database. The valid values could be found in https://azure.microsoft.com/en-us/pricing/details/sql-database/
    """
    BASIC = "Basic"
    STANDARD = "Standard"
    PREMIUM = "Premium"
    PREMIUM_RS = "PremiumRS"


class IntegrationRuntimeType(str, Enum):
    """
    Type of integration runtime.
    """
    MANAGED = "Managed"
    SELF_HOSTED = "SelfHosted"


class NodeSize(str, Enum):
    """
    The level of compute power that each node in the Big Data pool has.
    """
    NONE = "None"
    SMALL = "Small"
    MEDIUM = "Medium"
    LARGE = "Large"
    X_LARGE = "XLarge"
    XX_LARGE = "XXLarge"
    XXX_LARGE = "XXXLarge"


class NodeSizeFamily(str, Enum):
    """
    The kind of nodes that the Big Data pool provides.
    """
    NONE = "None"
    MEMORY_OPTIMIZED = "MemoryOptimized"
    HARDWARE_ACCELERATED_FPGA = "HardwareAcceleratedFPGA"
    HARDWARE_ACCELERATED_GPU = "HardwareAcceleratedGPU"


class ResourceIdentityType(str, Enum):
    """
    The type of managed identity for the workspace
    """
    NONE = "None"
    SYSTEM_ASSIGNED = "SystemAssigned"


class SensitivityLabelRank(str, Enum):
    NONE = "None"
    LOW = "Low"
    MEDIUM = "Medium"
    HIGH = "High"
    CRITICAL = "Critical"


class StorageAccountType(str, Enum):
    """
    The storage account type used to store backups for this sql pool.
    """
    GRS = "GRS"
    LRS = "LRS"
    ZRS = "ZRS"


class TransparentDataEncryptionStatus(str, Enum):
    """
    The status of the database transparent data encryption.
    """
    ENABLED = "Enabled"
    DISABLED = "Disabled"
