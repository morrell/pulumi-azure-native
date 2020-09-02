# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from ... import _utilities, _tables

__all__ = [
    'PlatformPropertiesArgs',
    'SourceControlAuthInfoArgs',
    'SourceRepositoryPropertiesArgs',
]

@pulumi.input_type
class PlatformPropertiesArgs:
    def __init__(__self__, *,
                 os_type: pulumi.Input[str],
                 cpu: Optional[pulumi.Input[float]] = None):
        """
        The platform properties against which the build has to happen.
        :param pulumi.Input[str] os_type: The operating system type required for the build.
        :param pulumi.Input[float] cpu: The CPU configuration in terms of number of cores required for the build.
        """
        pulumi.set(__self__, "os_type", os_type)
        if cpu is not None:
            pulumi.set(__self__, "cpu", cpu)

    @property
    @pulumi.getter(name="osType")
    def os_type(self) -> pulumi.Input[str]:
        """
        The operating system type required for the build.
        """
        return pulumi.get(self, "os_type")

    @os_type.setter
    def os_type(self, value: pulumi.Input[str]):
        pulumi.set(self, "os_type", value)

    @property
    @pulumi.getter
    def cpu(self) -> Optional[pulumi.Input[float]]:
        """
        The CPU configuration in terms of number of cores required for the build.
        """
        return pulumi.get(self, "cpu")

    @cpu.setter
    def cpu(self, value: Optional[pulumi.Input[float]]):
        pulumi.set(self, "cpu", value)


@pulumi.input_type
class SourceControlAuthInfoArgs:
    def __init__(__self__, *,
                 token: pulumi.Input[str],
                 expires_in: Optional[pulumi.Input[float]] = None,
                 refresh_token: Optional[pulumi.Input[str]] = None,
                 scope: Optional[pulumi.Input[str]] = None,
                 token_type: Optional[pulumi.Input[str]] = None):
        """
        The authorization properties for accessing the source code repository.
        :param pulumi.Input[str] token: The access token used to access the source control provider.
        :param pulumi.Input[float] expires_in: Time in seconds that the token remains valid
        :param pulumi.Input[str] refresh_token: The refresh token used to refresh the access token.
        :param pulumi.Input[str] scope: The scope of the access token.
        :param pulumi.Input[str] token_type: The type of Auth token.
        """
        pulumi.set(__self__, "token", token)
        if expires_in is not None:
            pulumi.set(__self__, "expires_in", expires_in)
        if refresh_token is not None:
            pulumi.set(__self__, "refresh_token", refresh_token)
        if scope is not None:
            pulumi.set(__self__, "scope", scope)
        if token_type is not None:
            pulumi.set(__self__, "token_type", token_type)

    @property
    @pulumi.getter
    def token(self) -> pulumi.Input[str]:
        """
        The access token used to access the source control provider.
        """
        return pulumi.get(self, "token")

    @token.setter
    def token(self, value: pulumi.Input[str]):
        pulumi.set(self, "token", value)

    @property
    @pulumi.getter(name="expiresIn")
    def expires_in(self) -> Optional[pulumi.Input[float]]:
        """
        Time in seconds that the token remains valid
        """
        return pulumi.get(self, "expires_in")

    @expires_in.setter
    def expires_in(self, value: Optional[pulumi.Input[float]]):
        pulumi.set(self, "expires_in", value)

    @property
    @pulumi.getter(name="refreshToken")
    def refresh_token(self) -> Optional[pulumi.Input[str]]:
        """
        The refresh token used to refresh the access token.
        """
        return pulumi.get(self, "refresh_token")

    @refresh_token.setter
    def refresh_token(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "refresh_token", value)

    @property
    @pulumi.getter
    def scope(self) -> Optional[pulumi.Input[str]]:
        """
        The scope of the access token.
        """
        return pulumi.get(self, "scope")

    @scope.setter
    def scope(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "scope", value)

    @property
    @pulumi.getter(name="tokenType")
    def token_type(self) -> Optional[pulumi.Input[str]]:
        """
        The type of Auth token.
        """
        return pulumi.get(self, "token_type")

    @token_type.setter
    def token_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "token_type", value)


@pulumi.input_type
class SourceRepositoryPropertiesArgs:
    def __init__(__self__, *,
                 repository_url: pulumi.Input[str],
                 source_control_type: pulumi.Input[str],
                 is_commit_trigger_enabled: Optional[pulumi.Input[bool]] = None,
                 source_control_auth_properties: Optional[pulumi.Input['SourceControlAuthInfoArgs']] = None):
        """
        The properties of the source code repository.
        :param pulumi.Input[str] repository_url: The full URL to the source code repository
        :param pulumi.Input[str] source_control_type: The type of source control service.
        :param pulumi.Input[bool] is_commit_trigger_enabled: The value of this property indicates whether the source control commit trigger is enabled or not.
        :param pulumi.Input['SourceControlAuthInfoArgs'] source_control_auth_properties: The authorization properties for accessing the source code repository.
        """
        pulumi.set(__self__, "repository_url", repository_url)
        pulumi.set(__self__, "source_control_type", source_control_type)
        if is_commit_trigger_enabled is not None:
            pulumi.set(__self__, "is_commit_trigger_enabled", is_commit_trigger_enabled)
        if source_control_auth_properties is not None:
            pulumi.set(__self__, "source_control_auth_properties", source_control_auth_properties)

    @property
    @pulumi.getter(name="repositoryUrl")
    def repository_url(self) -> pulumi.Input[str]:
        """
        The full URL to the source code repository
        """
        return pulumi.get(self, "repository_url")

    @repository_url.setter
    def repository_url(self, value: pulumi.Input[str]):
        pulumi.set(self, "repository_url", value)

    @property
    @pulumi.getter(name="sourceControlType")
    def source_control_type(self) -> pulumi.Input[str]:
        """
        The type of source control service.
        """
        return pulumi.get(self, "source_control_type")

    @source_control_type.setter
    def source_control_type(self, value: pulumi.Input[str]):
        pulumi.set(self, "source_control_type", value)

    @property
    @pulumi.getter(name="isCommitTriggerEnabled")
    def is_commit_trigger_enabled(self) -> Optional[pulumi.Input[bool]]:
        """
        The value of this property indicates whether the source control commit trigger is enabled or not.
        """
        return pulumi.get(self, "is_commit_trigger_enabled")

    @is_commit_trigger_enabled.setter
    def is_commit_trigger_enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "is_commit_trigger_enabled", value)

    @property
    @pulumi.getter(name="sourceControlAuthProperties")
    def source_control_auth_properties(self) -> Optional[pulumi.Input['SourceControlAuthInfoArgs']]:
        """
        The authorization properties for accessing the source code repository.
        """
        return pulumi.get(self, "source_control_auth_properties")

    @source_control_auth_properties.setter
    def source_control_auth_properties(self, value: Optional[pulumi.Input['SourceControlAuthInfoArgs']]):
        pulumi.set(self, "source_control_auth_properties", value)

