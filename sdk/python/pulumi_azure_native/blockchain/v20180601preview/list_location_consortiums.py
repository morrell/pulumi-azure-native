# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from ... import _utilities
from . import outputs

__all__ = [
    'ListLocationConsortiumsResult',
    'AwaitableListLocationConsortiumsResult',
    'list_location_consortiums',
]

@pulumi.output_type
class ListLocationConsortiumsResult:
    """
    Collection of the consortium payload.
    """
    def __init__(__self__, value=None):
        if value and not isinstance(value, list):
            raise TypeError("Expected argument 'value' to be a list")
        pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter
    def value(self) -> Optional[Sequence['outputs.ConsortiumResponse']]:
        """
        Gets or sets the collection of consortiums.
        """
        return pulumi.get(self, "value")


class AwaitableListLocationConsortiumsResult(ListLocationConsortiumsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return ListLocationConsortiumsResult(
            value=self.value)


def list_location_consortiums(location_name: Optional[str] = None,
                              opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableListLocationConsortiumsResult:
    """
    Collection of the consortium payload.


    :param str location_name: Location Name.
    """
    __args__ = dict()
    __args__['locationName'] = location_name
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('azure-native:blockchain/v20180601preview:listLocationConsortiums', __args__, opts=opts, typ=ListLocationConsortiumsResult).value

    return AwaitableListLocationConsortiumsResult(
        value=__ret__.value)
