# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from ... import _utilities, _tables

__all__ = ['PolicyAssignment']


class PolicyAssignment(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 parameters: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 policy_definition_id: Optional[pulumi.Input[str]] = None,
                 scope: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        The policy assignment.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: This message will be part of response in case of policy violation.
        :param pulumi.Input[str] display_name: The display name of the policy assignment.
        :param pulumi.Input[str] name: The name of the policy assignment.
        :param pulumi.Input[Mapping[str, Any]] parameters: Required if a parameter is used in policy rule.
        :param pulumi.Input[str] policy_definition_id: The ID of the policy definition.
        :param pulumi.Input[str] scope: The scope for the policy assignment.
        :param pulumi.Input[str] type: The type of the policy assignment.
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

            __props__['description'] = description
            __props__['display_name'] = display_name
            if name is None:
                raise TypeError("Missing required property 'name'")
            __props__['name'] = name
            __props__['parameters'] = parameters
            __props__['policy_definition_id'] = policy_definition_id
            if scope is None:
                raise TypeError("Missing required property 'scope'")
            __props__['scope'] = scope
            __props__['type'] = type
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="azurerm:authorization/v20151101:PolicyAssignment"), pulumi.Alias(type_="azurerm:authorization/v20180301:PolicyAssignment"), pulumi.Alias(type_="azurerm:authorization/v20180501:PolicyAssignment"), pulumi.Alias(type_="azurerm:authorization/v20190101:PolicyAssignment"), pulumi.Alias(type_="azurerm:authorization/v20190601:PolicyAssignment"), pulumi.Alias(type_="azurerm:authorization/v20190901:PolicyAssignment")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(PolicyAssignment, __self__).__init__(
            'azurerm:authorization/v20161201:PolicyAssignment',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'PolicyAssignment':
        """
        Get an existing PolicyAssignment resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return PolicyAssignment(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def description(self) -> Optional[str]:
        """
        This message will be part of response in case of policy violation.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> Optional[str]:
        """
        The display name of the policy assignment.
        """
        return pulumi.get(self, "display_name")

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        """
        The name of the policy assignment.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def parameters(self) -> Optional[Mapping[str, Any]]:
        """
        Required if a parameter is used in policy rule.
        """
        return pulumi.get(self, "parameters")

    @property
    @pulumi.getter(name="policyDefinitionId")
    def policy_definition_id(self) -> Optional[str]:
        """
        The ID of the policy definition.
        """
        return pulumi.get(self, "policy_definition_id")

    @property
    @pulumi.getter
    def scope(self) -> Optional[str]:
        """
        The scope for the policy assignment.
        """
        return pulumi.get(self, "scope")

    @property
    @pulumi.getter
    def type(self) -> Optional[str]:
        """
        The type of the policy assignment.
        """
        return pulumi.get(self, "type")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
