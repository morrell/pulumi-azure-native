# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs

__all__ = [
    'ListNamespaceKeysResult',
    'AwaitableListNamespaceKeysResult',
    'list_namespace_keys',
]

@pulumi.output_type
class ListNamespaceKeysResult:
    """
    The response of the List Namespace operation.
    """
    def __init__(__self__, next_link=None, value=None):
        if next_link and not isinstance(next_link, str):
            raise TypeError("Expected argument 'next_link' to be a str")
        pulumi.set(__self__, "next_link", next_link)
        if value and not isinstance(value, list):
            raise TypeError("Expected argument 'value' to be a list")
        pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter(name="nextLink")
    def next_link(self) -> Optional[str]:
        """
        Link to the next set of results. Not empty if Value contains incomplete list of AuthorizationRules
        """
        return pulumi.get(self, "next_link")

    @property
    @pulumi.getter
    def value(self) -> Optional[Sequence['outputs.SharedAccessAuthorizationRuleResourceResponse']]:
        """
        Result of the List AuthorizationRules operation.
        """
        return pulumi.get(self, "value")


class AwaitableListNamespaceKeysResult(ListNamespaceKeysResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return ListNamespaceKeysResult(
            next_link=self.next_link,
            value=self.value)


def list_namespace_keys(authorization_rule_name: Optional[str] = None,
                        namespace_name: Optional[str] = None,
                        resource_group_name: Optional[str] = None,
                        opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableListNamespaceKeysResult:
    """
    The response of the List Namespace operation.
    API Version: 2017-04-01.


    :param str authorization_rule_name: The connection string of the namespace for the specified authorizationRule.
    :param str namespace_name: The namespace name.
    :param str resource_group_name: The name of the resource group.
    """
    __args__ = dict()
    __args__['authorizationRuleName'] = authorization_rule_name
    __args__['namespaceName'] = namespace_name
    __args__['resourceGroupName'] = resource_group_name
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('azure-native:notificationhubs:listNamespaceKeys', __args__, opts=opts, typ=ListNamespaceKeysResult).value

    return AwaitableListNamespaceKeysResult(
        next_link=__ret__.next_link,
        value=__ret__.value)
