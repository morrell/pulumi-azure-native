# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from ... import _utilities, _tables
from . import outputs
from ._inputs import *

__all__ = ['ActionGroup']


class ActionGroup(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 arm_role_receivers: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['ArmRoleReceiverArgs']]]]] = None,
                 automation_runbook_receivers: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['AutomationRunbookReceiverArgs']]]]] = None,
                 azure_app_push_receivers: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['AzureAppPushReceiverArgs']]]]] = None,
                 azure_function_receivers: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['AzureFunctionReceiverArgs']]]]] = None,
                 email_receivers: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['EmailReceiverArgs']]]]] = None,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 group_short_name: Optional[pulumi.Input[str]] = None,
                 itsm_receivers: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['ItsmReceiverArgs']]]]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 logic_app_receivers: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['LogicAppReceiverArgs']]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 sms_receivers: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['SmsReceiverArgs']]]]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 voice_receivers: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['VoiceReceiverArgs']]]]] = None,
                 webhook_receivers: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['WebhookReceiverArgs']]]]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        An action group resource.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[List[pulumi.Input[pulumi.InputType['ArmRoleReceiverArgs']]]] arm_role_receivers: The list of ARM role receivers that are part of this action group. Roles are Azure RBAC roles and only built-in roles are supported.
        :param pulumi.Input[List[pulumi.Input[pulumi.InputType['AutomationRunbookReceiverArgs']]]] automation_runbook_receivers: The list of AutomationRunbook receivers that are part of this action group.
        :param pulumi.Input[List[pulumi.Input[pulumi.InputType['AzureAppPushReceiverArgs']]]] azure_app_push_receivers: The list of AzureAppPush receivers that are part of this action group.
        :param pulumi.Input[List[pulumi.Input[pulumi.InputType['AzureFunctionReceiverArgs']]]] azure_function_receivers: The list of azure function receivers that are part of this action group.
        :param pulumi.Input[List[pulumi.Input[pulumi.InputType['EmailReceiverArgs']]]] email_receivers: The list of email receivers that are part of this action group.
        :param pulumi.Input[bool] enabled: Indicates whether this action group is enabled. If an action group is not enabled, then none of its receivers will receive communications.
        :param pulumi.Input[str] group_short_name: The short name of the action group. This will be used in SMS messages.
        :param pulumi.Input[List[pulumi.Input[pulumi.InputType['ItsmReceiverArgs']]]] itsm_receivers: The list of ITSM receivers that are part of this action group.
        :param pulumi.Input[str] location: Resource location
        :param pulumi.Input[List[pulumi.Input[pulumi.InputType['LogicAppReceiverArgs']]]] logic_app_receivers: The list of logic app receivers that are part of this action group.
        :param pulumi.Input[str] name: The name of the action group.
        :param pulumi.Input[str] resource_group_name: The name of the resource group.
        :param pulumi.Input[List[pulumi.Input[pulumi.InputType['SmsReceiverArgs']]]] sms_receivers: The list of SMS receivers that are part of this action group.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Resource tags
        :param pulumi.Input[List[pulumi.Input[pulumi.InputType['VoiceReceiverArgs']]]] voice_receivers: The list of voice receivers that are part of this action group.
        :param pulumi.Input[List[pulumi.Input[pulumi.InputType['WebhookReceiverArgs']]]] webhook_receivers: The list of webhook receivers that are part of this action group.
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

            __props__['arm_role_receivers'] = arm_role_receivers
            __props__['automation_runbook_receivers'] = automation_runbook_receivers
            __props__['azure_app_push_receivers'] = azure_app_push_receivers
            __props__['azure_function_receivers'] = azure_function_receivers
            __props__['email_receivers'] = email_receivers
            if enabled is None:
                raise TypeError("Missing required property 'enabled'")
            __props__['enabled'] = enabled
            if group_short_name is None:
                raise TypeError("Missing required property 'group_short_name'")
            __props__['group_short_name'] = group_short_name
            __props__['itsm_receivers'] = itsm_receivers
            if location is None:
                raise TypeError("Missing required property 'location'")
            __props__['location'] = location
            __props__['logic_app_receivers'] = logic_app_receivers
            if name is None:
                raise TypeError("Missing required property 'name'")
            __props__['name'] = name
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['sms_receivers'] = sms_receivers
            __props__['tags'] = tags
            __props__['voice_receivers'] = voice_receivers
            __props__['webhook_receivers'] = webhook_receivers
            __props__['type'] = None
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="azurerm:insights/v20170401:ActionGroup"), pulumi.Alias(type_="azurerm:insights/v20180301:ActionGroup"), pulumi.Alias(type_="azurerm:insights/v20180901:ActionGroup"), pulumi.Alias(type_="azurerm:insights/v20190301:ActionGroup")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(ActionGroup, __self__).__init__(
            'azurerm:insights/v20190601:ActionGroup',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'ActionGroup':
        """
        Get an existing ActionGroup resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return ActionGroup(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="armRoleReceivers")
    def arm_role_receivers(self) -> Optional[List['outputs.ArmRoleReceiverResponse']]:
        """
        The list of ARM role receivers that are part of this action group. Roles are Azure RBAC roles and only built-in roles are supported.
        """
        return pulumi.get(self, "arm_role_receivers")

    @property
    @pulumi.getter(name="automationRunbookReceivers")
    def automation_runbook_receivers(self) -> Optional[List['outputs.AutomationRunbookReceiverResponse']]:
        """
        The list of AutomationRunbook receivers that are part of this action group.
        """
        return pulumi.get(self, "automation_runbook_receivers")

    @property
    @pulumi.getter(name="azureAppPushReceivers")
    def azure_app_push_receivers(self) -> Optional[List['outputs.AzureAppPushReceiverResponse']]:
        """
        The list of AzureAppPush receivers that are part of this action group.
        """
        return pulumi.get(self, "azure_app_push_receivers")

    @property
    @pulumi.getter(name="azureFunctionReceivers")
    def azure_function_receivers(self) -> Optional[List['outputs.AzureFunctionReceiverResponse']]:
        """
        The list of azure function receivers that are part of this action group.
        """
        return pulumi.get(self, "azure_function_receivers")

    @property
    @pulumi.getter(name="emailReceivers")
    def email_receivers(self) -> Optional[List['outputs.EmailReceiverResponse']]:
        """
        The list of email receivers that are part of this action group.
        """
        return pulumi.get(self, "email_receivers")

    @property
    @pulumi.getter
    def enabled(self) -> bool:
        """
        Indicates whether this action group is enabled. If an action group is not enabled, then none of its receivers will receive communications.
        """
        return pulumi.get(self, "enabled")

    @property
    @pulumi.getter(name="groupShortName")
    def group_short_name(self) -> str:
        """
        The short name of the action group. This will be used in SMS messages.
        """
        return pulumi.get(self, "group_short_name")

    @property
    @pulumi.getter(name="itsmReceivers")
    def itsm_receivers(self) -> Optional[List['outputs.ItsmReceiverResponse']]:
        """
        The list of ITSM receivers that are part of this action group.
        """
        return pulumi.get(self, "itsm_receivers")

    @property
    @pulumi.getter
    def location(self) -> str:
        """
        Resource location
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter(name="logicAppReceivers")
    def logic_app_receivers(self) -> Optional[List['outputs.LogicAppReceiverResponse']]:
        """
        The list of logic app receivers that are part of this action group.
        """
        return pulumi.get(self, "logic_app_receivers")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Azure resource name
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="smsReceivers")
    def sms_receivers(self) -> Optional[List['outputs.SmsReceiverResponse']]:
        """
        The list of SMS receivers that are part of this action group.
        """
        return pulumi.get(self, "sms_receivers")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Mapping[str, str]]:
        """
        Resource tags
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        Azure resource type
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter(name="voiceReceivers")
    def voice_receivers(self) -> Optional[List['outputs.VoiceReceiverResponse']]:
        """
        The list of voice receivers that are part of this action group.
        """
        return pulumi.get(self, "voice_receivers")

    @property
    @pulumi.getter(name="webhookReceivers")
    def webhook_receivers(self) -> Optional[List['outputs.WebhookReceiverResponse']]:
        """
        The list of webhook receivers that are part of this action group.
        """
        return pulumi.get(self, "webhook_receivers")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
