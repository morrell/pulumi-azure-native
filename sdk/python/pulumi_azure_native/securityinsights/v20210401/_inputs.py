# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from ... import _utilities
from ._enums import *

__all__ = [
    'IncidentLabelArgs',
    'IncidentOwnerInfoArgs',
    'WatchlistUserInfoArgs',
]

@pulumi.input_type
class IncidentLabelArgs:
    def __init__(__self__, *,
                 label_name: pulumi.Input[str]):
        """
        Represents an incident label
        :param pulumi.Input[str] label_name: The name of the label
        """
        pulumi.set(__self__, "label_name", label_name)

    @property
    @pulumi.getter(name="labelName")
    def label_name(self) -> pulumi.Input[str]:
        """
        The name of the label
        """
        return pulumi.get(self, "label_name")

    @label_name.setter
    def label_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "label_name", value)


@pulumi.input_type
class IncidentOwnerInfoArgs:
    def __init__(__self__, *,
                 assigned_to: Optional[pulumi.Input[str]] = None,
                 email: Optional[pulumi.Input[str]] = None,
                 object_id: Optional[pulumi.Input[str]] = None,
                 user_principal_name: Optional[pulumi.Input[str]] = None):
        """
        Information on the user an incident is assigned to
        :param pulumi.Input[str] assigned_to: The name of the user the incident is assigned to.
        :param pulumi.Input[str] email: The email of the user the incident is assigned to.
        :param pulumi.Input[str] object_id: The object id of the user the incident is assigned to.
        :param pulumi.Input[str] user_principal_name: The user principal name of the user the incident is assigned to.
        """
        if assigned_to is not None:
            pulumi.set(__self__, "assigned_to", assigned_to)
        if email is not None:
            pulumi.set(__self__, "email", email)
        if object_id is not None:
            pulumi.set(__self__, "object_id", object_id)
        if user_principal_name is not None:
            pulumi.set(__self__, "user_principal_name", user_principal_name)

    @property
    @pulumi.getter(name="assignedTo")
    def assigned_to(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the user the incident is assigned to.
        """
        return pulumi.get(self, "assigned_to")

    @assigned_to.setter
    def assigned_to(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "assigned_to", value)

    @property
    @pulumi.getter
    def email(self) -> Optional[pulumi.Input[str]]:
        """
        The email of the user the incident is assigned to.
        """
        return pulumi.get(self, "email")

    @email.setter
    def email(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "email", value)

    @property
    @pulumi.getter(name="objectId")
    def object_id(self) -> Optional[pulumi.Input[str]]:
        """
        The object id of the user the incident is assigned to.
        """
        return pulumi.get(self, "object_id")

    @object_id.setter
    def object_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "object_id", value)

    @property
    @pulumi.getter(name="userPrincipalName")
    def user_principal_name(self) -> Optional[pulumi.Input[str]]:
        """
        The user principal name of the user the incident is assigned to.
        """
        return pulumi.get(self, "user_principal_name")

    @user_principal_name.setter
    def user_principal_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "user_principal_name", value)


@pulumi.input_type
class WatchlistUserInfoArgs:
    def __init__(__self__, *,
                 object_id: Optional[pulumi.Input[str]] = None):
        """
        User information that made some action
        :param pulumi.Input[str] object_id: The object id of the user.
        """
        if object_id is not None:
            pulumi.set(__self__, "object_id", object_id)

    @property
    @pulumi.getter(name="objectId")
    def object_id(self) -> Optional[pulumi.Input[str]]:
        """
        The object id of the user.
        """
        return pulumi.get(self, "object_id")

    @object_id.setter
    def object_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "object_id", value)


