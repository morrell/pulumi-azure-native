# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

from enum import Enum

__all__ = [
    'MediaGraphRtspTransport',
]


class MediaGraphRtspTransport(str, Enum):
    """
    Underlying RTSP transport. This can be used to enable or disable HTTP tunneling.
    """
    HTTP = "Http"
    """HTTP/HTTPS transport. This should be used when HTTP tunneling is desired."""
    TCP = "Tcp"
    """TCP transport. This should be used when HTTP tunneling is not desired."""
