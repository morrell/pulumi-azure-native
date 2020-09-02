# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from ... import _utilities, _tables

__all__ = [
    'ActivityLogAlertActionGroupArgs',
    'ActivityLogAlertActionListArgs',
    'ActivityLogAlertAllOfConditionArgs',
    'ActivityLogAlertLeafConditionArgs',
]

@pulumi.input_type
class ActivityLogAlertActionGroupArgs:
    def __init__(__self__, *,
                 action_group_id: pulumi.Input[str],
                 webhook_properties: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None):
        """
        A pointer to an Azure Action Group.
        :param pulumi.Input[str] action_group_id: The resourceId of the action group. This cannot be null or empty.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] webhook_properties: The dictionary of custom properties to include with the post operation. These data are appended to the webhook payload.
        """
        pulumi.set(__self__, "action_group_id", action_group_id)
        if webhook_properties is not None:
            pulumi.set(__self__, "webhook_properties", webhook_properties)

    @property
    @pulumi.getter(name="actionGroupId")
    def action_group_id(self) -> pulumi.Input[str]:
        """
        The resourceId of the action group. This cannot be null or empty.
        """
        return pulumi.get(self, "action_group_id")

    @action_group_id.setter
    def action_group_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "action_group_id", value)

    @property
    @pulumi.getter(name="webhookProperties")
    def webhook_properties(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        The dictionary of custom properties to include with the post operation. These data are appended to the webhook payload.
        """
        return pulumi.get(self, "webhook_properties")

    @webhook_properties.setter
    def webhook_properties(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "webhook_properties", value)


@pulumi.input_type
class ActivityLogAlertActionListArgs:
    def __init__(__self__, *,
                 action_groups: Optional[pulumi.Input[List[pulumi.Input['ActivityLogAlertActionGroupArgs']]]] = None):
        """
        A list of activity log alert actions.
        :param pulumi.Input[List[pulumi.Input['ActivityLogAlertActionGroupArgs']]] action_groups: The list of activity log alerts.
        """
        if action_groups is not None:
            pulumi.set(__self__, "action_groups", action_groups)

    @property
    @pulumi.getter(name="actionGroups")
    def action_groups(self) -> Optional[pulumi.Input[List[pulumi.Input['ActivityLogAlertActionGroupArgs']]]]:
        """
        The list of activity log alerts.
        """
        return pulumi.get(self, "action_groups")

    @action_groups.setter
    def action_groups(self, value: Optional[pulumi.Input[List[pulumi.Input['ActivityLogAlertActionGroupArgs']]]]):
        pulumi.set(self, "action_groups", value)


@pulumi.input_type
class ActivityLogAlertAllOfConditionArgs:
    def __init__(__self__, *,
                 all_of: pulumi.Input[List[pulumi.Input['ActivityLogAlertLeafConditionArgs']]]):
        """
        An Activity Log alert condition that is met when all its member conditions are met.
        :param pulumi.Input[List[pulumi.Input['ActivityLogAlertLeafConditionArgs']]] all_of: The list of activity log alert conditions.
        """
        pulumi.set(__self__, "all_of", all_of)

    @property
    @pulumi.getter(name="allOf")
    def all_of(self) -> pulumi.Input[List[pulumi.Input['ActivityLogAlertLeafConditionArgs']]]:
        """
        The list of activity log alert conditions.
        """
        return pulumi.get(self, "all_of")

    @all_of.setter
    def all_of(self, value: pulumi.Input[List[pulumi.Input['ActivityLogAlertLeafConditionArgs']]]):
        pulumi.set(self, "all_of", value)


@pulumi.input_type
class ActivityLogAlertLeafConditionArgs:
    def __init__(__self__, *,
                 equals: pulumi.Input[str],
                 field: pulumi.Input[str]):
        """
        An Activity Log alert condition that is met by comparing an activity log field and value.
        :param pulumi.Input[str] equals: The field value will be compared to this value (case-insensitive) to determine if the condition is met.
        :param pulumi.Input[str] field: The name of the field that this condition will examine. The possible values for this field are (case-insensitive): 'resourceId', 'category', 'caller', 'level', 'operationName', 'resourceGroup', 'resourceProvider', 'status', 'subStatus', 'resourceType', or anything beginning with 'properties.'.
        """
        pulumi.set(__self__, "equals", equals)
        pulumi.set(__self__, "field", field)

    @property
    @pulumi.getter
    def equals(self) -> pulumi.Input[str]:
        """
        The field value will be compared to this value (case-insensitive) to determine if the condition is met.
        """
        return pulumi.get(self, "equals")

    @equals.setter
    def equals(self, value: pulumi.Input[str]):
        pulumi.set(self, "equals", value)

    @property
    @pulumi.getter
    def field(self) -> pulumi.Input[str]:
        """
        The name of the field that this condition will examine. The possible values for this field are (case-insensitive): 'resourceId', 'category', 'caller', 'level', 'operationName', 'resourceGroup', 'resourceProvider', 'status', 'subStatus', 'resourceType', or anything beginning with 'properties.'.
        """
        return pulumi.get(self, "field")

    @field.setter
    def field(self, value: pulumi.Input[str]):
        pulumi.set(self, "field", value)

