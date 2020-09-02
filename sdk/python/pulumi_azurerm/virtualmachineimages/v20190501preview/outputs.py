# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from ... import _utilities, _tables
from . import outputs

__all__ = [
    'ImageTemplateCustomizerResponse',
    'ImageTemplateDistributorResponse',
    'ImageTemplateIdentityResponse',
    'ImageTemplateIdentityResponseUserAssignedIdentities',
    'ImageTemplateLastRunStatusResponse',
    'ImageTemplateSourceResponse',
    'ImageTemplateVmProfileResponse',
    'ProvisioningErrorResponse',
]

@pulumi.output_type
class ImageTemplateCustomizerResponse(dict):
    """
    Describes a unit of image customization
    """
    def __init__(__self__, *,
                 type: str,
                 name: Optional[str] = None):
        """
        Describes a unit of image customization
        :param str type: The type of customization tool you want to use on the Image. For example, "Shell" can be shell customizer
        :param str name: Friendly Name to provide context on what this customization step does
        """
        pulumi.set(__self__, "type", type)
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        The type of customization tool you want to use on the Image. For example, "Shell" can be shell customizer
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        """
        Friendly Name to provide context on what this customization step does
        """
        return pulumi.get(self, "name")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class ImageTemplateDistributorResponse(dict):
    """
    Generic distribution object
    """
    def __init__(__self__, *,
                 run_output_name: str,
                 type: str,
                 artifact_tags: Optional[Mapping[str, str]] = None):
        """
        Generic distribution object
        :param str run_output_name: The name to be used for the associated RunOutput.
        :param str type: Type of distribution.
        :param Mapping[str, str] artifact_tags: Tags that will be applied to the artifact once it has been created/updated by the distributor.
        """
        pulumi.set(__self__, "run_output_name", run_output_name)
        pulumi.set(__self__, "type", type)
        if artifact_tags is not None:
            pulumi.set(__self__, "artifact_tags", artifact_tags)

    @property
    @pulumi.getter(name="runOutputName")
    def run_output_name(self) -> str:
        """
        The name to be used for the associated RunOutput.
        """
        return pulumi.get(self, "run_output_name")

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        Type of distribution.
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter(name="artifactTags")
    def artifact_tags(self) -> Optional[Mapping[str, str]]:
        """
        Tags that will be applied to the artifact once it has been created/updated by the distributor.
        """
        return pulumi.get(self, "artifact_tags")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class ImageTemplateIdentityResponse(dict):
    """
    Identity for the image template.
    """
    def __init__(__self__, *,
                 type: Optional[str] = None,
                 user_assigned_identities: Optional[Mapping[str, 'outputs.ImageTemplateIdentityResponseUserAssignedIdentities']] = None):
        """
        Identity for the image template.
        :param str type: The type of identity used for the image template. The type 'None' will remove any identities from the image template.
        :param Mapping[str, 'ImageTemplateIdentityResponseUserAssignedIdentitiesArgs'] user_assigned_identities: The list of user identities associated with the image template. The user identity dictionary key references will be ARM resource ids in the form: '/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ManagedIdentity/userAssignedIdentities/{identityName}'.
        """
        if type is not None:
            pulumi.set(__self__, "type", type)
        if user_assigned_identities is not None:
            pulumi.set(__self__, "user_assigned_identities", user_assigned_identities)

    @property
    @pulumi.getter
    def type(self) -> Optional[str]:
        """
        The type of identity used for the image template. The type 'None' will remove any identities from the image template.
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter(name="userAssignedIdentities")
    def user_assigned_identities(self) -> Optional[Mapping[str, 'outputs.ImageTemplateIdentityResponseUserAssignedIdentities']]:
        """
        The list of user identities associated with the image template. The user identity dictionary key references will be ARM resource ids in the form: '/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ManagedIdentity/userAssignedIdentities/{identityName}'.
        """
        return pulumi.get(self, "user_assigned_identities")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class ImageTemplateIdentityResponseUserAssignedIdentities(dict):
    def __init__(__self__, *,
                 client_id: str,
                 principal_id: str):
        """
        :param str client_id: The client id of user assigned identity.
        :param str principal_id: The principal id of user assigned identity.
        """
        pulumi.set(__self__, "client_id", client_id)
        pulumi.set(__self__, "principal_id", principal_id)

    @property
    @pulumi.getter(name="clientId")
    def client_id(self) -> str:
        """
        The client id of user assigned identity.
        """
        return pulumi.get(self, "client_id")

    @property
    @pulumi.getter(name="principalId")
    def principal_id(self) -> str:
        """
        The principal id of user assigned identity.
        """
        return pulumi.get(self, "principal_id")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class ImageTemplateLastRunStatusResponse(dict):
    """
    Describes the latest status of running an image template
    """
    def __init__(__self__, *,
                 end_time: Optional[str] = None,
                 message: Optional[str] = None,
                 run_state: Optional[str] = None,
                 run_sub_state: Optional[str] = None,
                 start_time: Optional[str] = None):
        """
        Describes the latest status of running an image template
        :param str end_time: End time of the last run (UTC)
        :param str message: Verbose information about the last run state
        :param str run_state: State of the last run
        :param str run_sub_state: Sub-state of the last run
        :param str start_time: Start time of the last run (UTC)
        """
        if end_time is not None:
            pulumi.set(__self__, "end_time", end_time)
        if message is not None:
            pulumi.set(__self__, "message", message)
        if run_state is not None:
            pulumi.set(__self__, "run_state", run_state)
        if run_sub_state is not None:
            pulumi.set(__self__, "run_sub_state", run_sub_state)
        if start_time is not None:
            pulumi.set(__self__, "start_time", start_time)

    @property
    @pulumi.getter(name="endTime")
    def end_time(self) -> Optional[str]:
        """
        End time of the last run (UTC)
        """
        return pulumi.get(self, "end_time")

    @property
    @pulumi.getter
    def message(self) -> Optional[str]:
        """
        Verbose information about the last run state
        """
        return pulumi.get(self, "message")

    @property
    @pulumi.getter(name="runState")
    def run_state(self) -> Optional[str]:
        """
        State of the last run
        """
        return pulumi.get(self, "run_state")

    @property
    @pulumi.getter(name="runSubState")
    def run_sub_state(self) -> Optional[str]:
        """
        Sub-state of the last run
        """
        return pulumi.get(self, "run_sub_state")

    @property
    @pulumi.getter(name="startTime")
    def start_time(self) -> Optional[str]:
        """
        Start time of the last run (UTC)
        """
        return pulumi.get(self, "start_time")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class ImageTemplateSourceResponse(dict):
    """
    Describes a virtual machine image source for building, customizing and distributing
    """
    def __init__(__self__, *,
                 type: str):
        """
        Describes a virtual machine image source for building, customizing and distributing
        :param str type: Specifies the type of source image you want to start with.
        """
        pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        Specifies the type of source image you want to start with.
        """
        return pulumi.get(self, "type")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class ImageTemplateVmProfileResponse(dict):
    """
    Describes the virtual machine used to build, customize and capture images
    """
    def __init__(__self__, *,
                 vm_size: Optional[str] = None):
        """
        Describes the virtual machine used to build, customize and capture images
        :param str vm_size: Size of the virtual machine used to build, customize and capture images. Omit or specify empty string to use the default (Standard_D1_v2).
        """
        if vm_size is not None:
            pulumi.set(__self__, "vm_size", vm_size)

    @property
    @pulumi.getter(name="vmSize")
    def vm_size(self) -> Optional[str]:
        """
        Size of the virtual machine used to build, customize and capture images. Omit or specify empty string to use the default (Standard_D1_v2).
        """
        return pulumi.get(self, "vm_size")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class ProvisioningErrorResponse(dict):
    """
    Describes the error happened when create or update an image template
    """
    def __init__(__self__, *,
                 message: Optional[str] = None,
                 provisioning_error_code: Optional[str] = None):
        """
        Describes the error happened when create or update an image template
        :param str message: Verbose error message about the provisioning failure
        :param str provisioning_error_code: Error code of the provisioning failure
        """
        if message is not None:
            pulumi.set(__self__, "message", message)
        if provisioning_error_code is not None:
            pulumi.set(__self__, "provisioning_error_code", provisioning_error_code)

    @property
    @pulumi.getter
    def message(self) -> Optional[str]:
        """
        Verbose error message about the provisioning failure
        """
        return pulumi.get(self, "message")

    @property
    @pulumi.getter(name="provisioningErrorCode")
    def provisioning_error_code(self) -> Optional[str]:
        """
        Error code of the provisioning failure
        """
        return pulumi.get(self, "provisioning_error_code")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

