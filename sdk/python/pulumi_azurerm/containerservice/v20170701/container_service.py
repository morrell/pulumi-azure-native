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

__all__ = ['ContainerService']


class ContainerService(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 agent_pool_profiles: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['ContainerServiceAgentPoolProfileArgs']]]]] = None,
                 custom_profile: Optional[pulumi.Input[pulumi.InputType['ContainerServiceCustomProfileArgs']]] = None,
                 diagnostics_profile: Optional[pulumi.Input[pulumi.InputType['ContainerServiceDiagnosticsProfileArgs']]] = None,
                 linux_profile: Optional[pulumi.Input[pulumi.InputType['ContainerServiceLinuxProfileArgs']]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 master_profile: Optional[pulumi.Input[pulumi.InputType['ContainerServiceMasterProfileArgs']]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 orchestrator_profile: Optional[pulumi.Input[pulumi.InputType['ContainerServiceOrchestratorProfileArgs']]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 service_principal_profile: Optional[pulumi.Input[pulumi.InputType['ContainerServiceServicePrincipalProfileArgs']]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 windows_profile: Optional[pulumi.Input[pulumi.InputType['ContainerServiceWindowsProfileArgs']]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Container service.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[List[pulumi.Input[pulumi.InputType['ContainerServiceAgentPoolProfileArgs']]]] agent_pool_profiles: Properties of the agent pool.
        :param pulumi.Input[pulumi.InputType['ContainerServiceCustomProfileArgs']] custom_profile: Properties to configure a custom container service cluster.
        :param pulumi.Input[pulumi.InputType['ContainerServiceDiagnosticsProfileArgs']] diagnostics_profile: Profile for diagnostics in the container service cluster.
        :param pulumi.Input[pulumi.InputType['ContainerServiceLinuxProfileArgs']] linux_profile: Profile for Linux VMs in the container service cluster.
        :param pulumi.Input[str] location: Resource location
        :param pulumi.Input[pulumi.InputType['ContainerServiceMasterProfileArgs']] master_profile: Profile for the container service master.
        :param pulumi.Input[str] name: The name of the container service in the specified subscription and resource group.
        :param pulumi.Input[pulumi.InputType['ContainerServiceOrchestratorProfileArgs']] orchestrator_profile: Profile for the container service orchestrator.
        :param pulumi.Input[str] resource_group_name: The name of the resource group.
        :param pulumi.Input[pulumi.InputType['ContainerServiceServicePrincipalProfileArgs']] service_principal_profile: Information about a service principal identity for the cluster to use for manipulating Azure APIs. Exact one of secret or keyVaultSecretRef need to be specified.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Resource tags
        :param pulumi.Input[pulumi.InputType['ContainerServiceWindowsProfileArgs']] windows_profile: Profile for Windows VMs in the container service cluster.
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

            __props__['agent_pool_profiles'] = agent_pool_profiles
            __props__['custom_profile'] = custom_profile
            __props__['diagnostics_profile'] = diagnostics_profile
            if linux_profile is None:
                raise TypeError("Missing required property 'linux_profile'")
            __props__['linux_profile'] = linux_profile
            if location is None:
                raise TypeError("Missing required property 'location'")
            __props__['location'] = location
            if master_profile is None:
                raise TypeError("Missing required property 'master_profile'")
            __props__['master_profile'] = master_profile
            if name is None:
                raise TypeError("Missing required property 'name'")
            __props__['name'] = name
            if orchestrator_profile is None:
                raise TypeError("Missing required property 'orchestrator_profile'")
            __props__['orchestrator_profile'] = orchestrator_profile
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['service_principal_profile'] = service_principal_profile
            __props__['tags'] = tags
            __props__['windows_profile'] = windows_profile
            __props__['provisioning_state'] = None
            __props__['type'] = None
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="azurerm:containerservice/v20160330:ContainerService"), pulumi.Alias(type_="azurerm:containerservice/v20160930:ContainerService"), pulumi.Alias(type_="azurerm:containerservice/v20170131:ContainerService")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(ContainerService, __self__).__init__(
            'azurerm:containerservice/v20170701:ContainerService',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'ContainerService':
        """
        Get an existing ContainerService resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return ContainerService(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="agentPoolProfiles")
    def agent_pool_profiles(self) -> Optional[List['outputs.ContainerServiceAgentPoolProfileResponse']]:
        """
        Properties of the agent pool.
        """
        return pulumi.get(self, "agent_pool_profiles")

    @property
    @pulumi.getter(name="customProfile")
    def custom_profile(self) -> Optional['outputs.ContainerServiceCustomProfileResponse']:
        """
        Properties to configure a custom container service cluster.
        """
        return pulumi.get(self, "custom_profile")

    @property
    @pulumi.getter(name="diagnosticsProfile")
    def diagnostics_profile(self) -> Optional['outputs.ContainerServiceDiagnosticsProfileResponse']:
        """
        Profile for diagnostics in the container service cluster.
        """
        return pulumi.get(self, "diagnostics_profile")

    @property
    @pulumi.getter(name="linuxProfile")
    def linux_profile(self) -> 'outputs.ContainerServiceLinuxProfileResponse':
        """
        Profile for Linux VMs in the container service cluster.
        """
        return pulumi.get(self, "linux_profile")

    @property
    @pulumi.getter
    def location(self) -> str:
        """
        Resource location
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter(name="masterProfile")
    def master_profile(self) -> 'outputs.ContainerServiceMasterProfileResponse':
        """
        Profile for the container service master.
        """
        return pulumi.get(self, "master_profile")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Resource name
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="orchestratorProfile")
    def orchestrator_profile(self) -> 'outputs.ContainerServiceOrchestratorProfileResponse':
        """
        Profile for the container service orchestrator.
        """
        return pulumi.get(self, "orchestrator_profile")

    @property
    @pulumi.getter(name="provisioningState")
    def provisioning_state(self) -> str:
        """
        The current deployment or provisioning state, which only appears in the response.
        """
        return pulumi.get(self, "provisioning_state")

    @property
    @pulumi.getter(name="servicePrincipalProfile")
    def service_principal_profile(self) -> Optional['outputs.ContainerServiceServicePrincipalProfileResponse']:
        """
        Information about a service principal identity for the cluster to use for manipulating Azure APIs. Exact one of secret or keyVaultSecretRef need to be specified.
        """
        return pulumi.get(self, "service_principal_profile")

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
        Resource type
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter(name="windowsProfile")
    def windows_profile(self) -> Optional['outputs.ContainerServiceWindowsProfileResponse']:
        """
        Profile for Windows VMs in the container service cluster.
        """
        return pulumi.get(self, "windows_profile")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
