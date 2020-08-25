# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from ... import _utilities, _tables

__all__ = ['ManagementLock']


class ManagementLock(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 level: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 notes: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Management lock information.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] level: The lock level of the management lock.
        :param pulumi.Input[str] name: The name of lock.
        :param pulumi.Input[str] notes: The notes of the management lock.
        """
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            __props__['level'] = level
            if name is None:
                raise TypeError("Missing required property 'name'")
            __props__['name'] = name
            __props__['notes'] = notes
            __props__['type'] = None
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="azurerm:authorization/v20160901:ManagementLock")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(ManagementLock, __self__).__init__(
            'azurerm:authorization/v20150101:ManagementLock',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'ManagementLock':
        """
        Get an existing ManagementLock resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return ManagementLock(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def level(self) -> Optional[str]:
        """
        The lock level of the management lock.
        """
        return pulumi.get(self, "level")

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        """
        The name of the lock.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def notes(self) -> Optional[str]:
        """
        The notes of the management lock.
        """
        return pulumi.get(self, "notes")

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        The type of the lock.
        """
        return pulumi.get(self, "type")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
