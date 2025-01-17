# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

from enum import Enum

__all__ = [
    'PublicNetworkAccessType',
    'StatusLevelTypes',
]


class PublicNetworkAccessType(str, Enum):
    """
    Indicates whether machines associated with the private link scope can also use public Azure Arc service endpoints.
    """
    ENABLED = "Enabled"
    """Allows Azure Arc agents to communicate with Azure Arc services over both public (internet) and private endpoints."""
    DISABLED = "Disabled"
    """Does not allow Azure Arc agents to communicate with Azure Arc services over public (internet) endpoints. The agents must use the private link."""


class StatusLevelTypes(str, Enum):
    """
    The level code.
    """
    INFO = "Info"
    WARNING = "Warning"
    ERROR = "Error"
