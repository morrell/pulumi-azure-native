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
    'AutomaticResolutionPropertiesResponse',
    'IdentityResponse',
    'JobStatusResponse',
    'ManualResolutionPropertiesResponse',
    'MoveCollectionPropertiesResponse',
    'MoveResourceDependencyOverrideResponse',
    'MoveResourceDependencyResponse',
    'MoveResourceErrorBodyResponse',
    'MoveResourceErrorResponse',
    'MoveResourcePropertiesResponse',
    'MoveResourcePropertiesResponseErrors',
    'MoveResourcePropertiesResponseMoveStatus',
    'MoveResourcePropertiesResponseSourceResourceSettings',
    'ResourceSettingsResponse',
]

@pulumi.output_type
class AutomaticResolutionPropertiesResponse(dict):
    """
    Defines the properties for automatic resolution.
    """
    def __init__(__self__, *,
                 move_resource_id: Optional[str] = None):
        """
        Defines the properties for automatic resolution.
        :param str move_resource_id: Gets the MoveResource ARM ID of
               the dependent resource if the resolution type is Automatic.
        """
        if move_resource_id is not None:
            pulumi.set(__self__, "move_resource_id", move_resource_id)

    @property
    @pulumi.getter(name="moveResourceId")
    def move_resource_id(self) -> Optional[str]:
        """
        Gets the MoveResource ARM ID of
        the dependent resource if the resolution type is Automatic.
        """
        return pulumi.get(self, "move_resource_id")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class IdentityResponse(dict):
    """
    Defines the MSI properties of the Move Collection.
    """
    def __init__(__self__, *,
                 principal_id: Optional[str] = None,
                 tenant_id: Optional[str] = None,
                 type: Optional[str] = None):
        """
        Defines the MSI properties of the Move Collection.
        :param str principal_id: Gets or sets the principal id.
        :param str tenant_id: Gets or sets the tenant id.
        :param str type: The type of identity used for the region move service.
        """
        if principal_id is not None:
            pulumi.set(__self__, "principal_id", principal_id)
        if tenant_id is not None:
            pulumi.set(__self__, "tenant_id", tenant_id)
        if type is not None:
            pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter(name="principalId")
    def principal_id(self) -> Optional[str]:
        """
        Gets or sets the principal id.
        """
        return pulumi.get(self, "principal_id")

    @property
    @pulumi.getter(name="tenantId")
    def tenant_id(self) -> Optional[str]:
        """
        Gets or sets the tenant id.
        """
        return pulumi.get(self, "tenant_id")

    @property
    @pulumi.getter
    def type(self) -> Optional[str]:
        """
        The type of identity used for the region move service.
        """
        return pulumi.get(self, "type")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class JobStatusResponse(dict):
    """
    Defines the job status.
    """
    def __init__(__self__, *,
                 job_progress: str,
                 job_name: Optional[str] = None):
        """
        Defines the job status.
        :param str job_progress: Gets or sets the monitoring job percentage.
        :param str job_name: Defines the job name.
        """
        pulumi.set(__self__, "job_progress", job_progress)
        if job_name is not None:
            pulumi.set(__self__, "job_name", job_name)

    @property
    @pulumi.getter(name="jobProgress")
    def job_progress(self) -> str:
        """
        Gets or sets the monitoring job percentage.
        """
        return pulumi.get(self, "job_progress")

    @property
    @pulumi.getter(name="jobName")
    def job_name(self) -> Optional[str]:
        """
        Defines the job name.
        """
        return pulumi.get(self, "job_name")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class ManualResolutionPropertiesResponse(dict):
    """
    Defines the properties for manual resolution.
    """
    def __init__(__self__, *,
                 target_id: Optional[str] = None):
        """
        Defines the properties for manual resolution.
        :param str target_id: Gets or sets the target resource ARM ID of the dependent resource if the resource type is Manual.
        """
        if target_id is not None:
            pulumi.set(__self__, "target_id", target_id)

    @property
    @pulumi.getter(name="targetId")
    def target_id(self) -> Optional[str]:
        """
        Gets or sets the target resource ARM ID of the dependent resource if the resource type is Manual.
        """
        return pulumi.get(self, "target_id")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class MoveCollectionPropertiesResponse(dict):
    """
    Defines the move collection properties.
    """
    def __init__(__self__, *,
                 source_region: str,
                 target_region: str,
                 provisioning_state: Optional[str] = None):
        """
        Defines the move collection properties.
        :param str source_region: Gets or sets the source region.
        :param str target_region: Gets or sets the target region.
        :param str provisioning_state: Defines the provisioning states.
        """
        pulumi.set(__self__, "source_region", source_region)
        pulumi.set(__self__, "target_region", target_region)
        if provisioning_state is not None:
            pulumi.set(__self__, "provisioning_state", provisioning_state)

    @property
    @pulumi.getter(name="sourceRegion")
    def source_region(self) -> str:
        """
        Gets or sets the source region.
        """
        return pulumi.get(self, "source_region")

    @property
    @pulumi.getter(name="targetRegion")
    def target_region(self) -> str:
        """
        Gets or sets the target region.
        """
        return pulumi.get(self, "target_region")

    @property
    @pulumi.getter(name="provisioningState")
    def provisioning_state(self) -> Optional[str]:
        """
        Defines the provisioning states.
        """
        return pulumi.get(self, "provisioning_state")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class MoveResourceDependencyOverrideResponse(dict):
    """
    Defines the dependency override of the move resource.
    """
    def __init__(__self__, *,
                 id: Optional[str] = None,
                 target_id: Optional[str] = None):
        """
        Defines the dependency override of the move resource.
        :param str id: Gets or sets the ARM ID of the dependent resource.
        :param str target_id: Gets or sets the resource ARM id of either the MoveResource or the resource ARM ID of
               the dependent resource.
        """
        if id is not None:
            pulumi.set(__self__, "id", id)
        if target_id is not None:
            pulumi.set(__self__, "target_id", target_id)

    @property
    @pulumi.getter
    def id(self) -> Optional[str]:
        """
        Gets or sets the ARM ID of the dependent resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="targetId")
    def target_id(self) -> Optional[str]:
        """
        Gets or sets the resource ARM id of either the MoveResource or the resource ARM ID of
        the dependent resource.
        """
        return pulumi.get(self, "target_id")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class MoveResourceDependencyResponse(dict):
    """
    Defines the dependency of the move resource.
    """
    def __init__(__self__, *,
                 automatic_resolution: Optional['outputs.AutomaticResolutionPropertiesResponse'] = None,
                 dependency_type: Optional[str] = None,
                 id: Optional[str] = None,
                 is_optional: Optional[str] = None,
                 manual_resolution: Optional['outputs.ManualResolutionPropertiesResponse'] = None,
                 resolution_status: Optional[str] = None,
                 resolution_type: Optional[str] = None):
        """
        Defines the dependency of the move resource.
        :param 'AutomaticResolutionPropertiesResponseArgs' automatic_resolution: Defines the properties for automatic resolution.
        :param str dependency_type: Defines the dependency type.
        :param str id: Gets the source ARM ID of the dependent resource.
        :param str is_optional: Gets or sets a value indicating whether the dependency is optional.
        :param 'ManualResolutionPropertiesResponseArgs' manual_resolution: Defines the properties for manual resolution.
        :param str resolution_status: Gets the dependency resolution status.
        :param str resolution_type: Defines the resolution type.
        """
        if automatic_resolution is not None:
            pulumi.set(__self__, "automatic_resolution", automatic_resolution)
        if dependency_type is not None:
            pulumi.set(__self__, "dependency_type", dependency_type)
        if id is not None:
            pulumi.set(__self__, "id", id)
        if is_optional is not None:
            pulumi.set(__self__, "is_optional", is_optional)
        if manual_resolution is not None:
            pulumi.set(__self__, "manual_resolution", manual_resolution)
        if resolution_status is not None:
            pulumi.set(__self__, "resolution_status", resolution_status)
        if resolution_type is not None:
            pulumi.set(__self__, "resolution_type", resolution_type)

    @property
    @pulumi.getter(name="automaticResolution")
    def automatic_resolution(self) -> Optional['outputs.AutomaticResolutionPropertiesResponse']:
        """
        Defines the properties for automatic resolution.
        """
        return pulumi.get(self, "automatic_resolution")

    @property
    @pulumi.getter(name="dependencyType")
    def dependency_type(self) -> Optional[str]:
        """
        Defines the dependency type.
        """
        return pulumi.get(self, "dependency_type")

    @property
    @pulumi.getter
    def id(self) -> Optional[str]:
        """
        Gets the source ARM ID of the dependent resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="isOptional")
    def is_optional(self) -> Optional[str]:
        """
        Gets or sets a value indicating whether the dependency is optional.
        """
        return pulumi.get(self, "is_optional")

    @property
    @pulumi.getter(name="manualResolution")
    def manual_resolution(self) -> Optional['outputs.ManualResolutionPropertiesResponse']:
        """
        Defines the properties for manual resolution.
        """
        return pulumi.get(self, "manual_resolution")

    @property
    @pulumi.getter(name="resolutionStatus")
    def resolution_status(self) -> Optional[str]:
        """
        Gets the dependency resolution status.
        """
        return pulumi.get(self, "resolution_status")

    @property
    @pulumi.getter(name="resolutionType")
    def resolution_type(self) -> Optional[str]:
        """
        Defines the resolution type.
        """
        return pulumi.get(self, "resolution_type")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class MoveResourceErrorBodyResponse(dict):
    """
    An error response from the Azure Migrate service.
    """
    def __init__(__self__, *,
                 code: str,
                 details: List['outputs.MoveResourceErrorBodyResponse'],
                 message: str,
                 target: str):
        """
        An error response from the Azure Migrate service.
        :param str code: An identifier for the error. Codes are invariant and are intended to be consumed programmatically.
        :param List['MoveResourceErrorBodyResponseArgs'] details: A list of additional details about the error.
        :param str message: A message describing the error, intended to be suitable for display in a user interface.
        :param str target: The target of the particular error. For example, the name of the property in error.
        """
        pulumi.set(__self__, "code", code)
        pulumi.set(__self__, "details", details)
        pulumi.set(__self__, "message", message)
        pulumi.set(__self__, "target", target)

    @property
    @pulumi.getter
    def code(self) -> str:
        """
        An identifier for the error. Codes are invariant and are intended to be consumed programmatically.
        """
        return pulumi.get(self, "code")

    @property
    @pulumi.getter
    def details(self) -> List['outputs.MoveResourceErrorBodyResponse']:
        """
        A list of additional details about the error.
        """
        return pulumi.get(self, "details")

    @property
    @pulumi.getter
    def message(self) -> str:
        """
        A message describing the error, intended to be suitable for display in a user interface.
        """
        return pulumi.get(self, "message")

    @property
    @pulumi.getter
    def target(self) -> str:
        """
        The target of the particular error. For example, the name of the property in error.
        """
        return pulumi.get(self, "target")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class MoveResourceErrorResponse(dict):
    """
    An error response from the azure region move service.
    """
    def __init__(__self__, *,
                 properties: Optional['outputs.MoveResourceErrorBodyResponse'] = None):
        """
        An error response from the azure region move service.
        :param 'MoveResourceErrorBodyResponseArgs' properties: The move resource error body.
        """
        if properties is not None:
            pulumi.set(__self__, "properties", properties)

    @property
    @pulumi.getter
    def properties(self) -> Optional['outputs.MoveResourceErrorBodyResponse']:
        """
        The move resource error body.
        """
        return pulumi.get(self, "properties")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class MoveResourcePropertiesResponse(dict):
    """
    Defines the move resource properties.
    """
    def __init__(__self__, *,
                 depends_on: List['outputs.MoveResourceDependencyResponse'],
                 errors: 'outputs.MoveResourcePropertiesResponseErrors',
                 move_status: 'outputs.MoveResourcePropertiesResponseMoveStatus',
                 source_id: str,
                 source_resource_settings: 'outputs.MoveResourcePropertiesResponseSourceResourceSettings',
                 target_id: str,
                 depends_on_overrides: Optional[List['outputs.MoveResourceDependencyOverrideResponse']] = None,
                 existing_target_id: Optional[str] = None,
                 provisioning_state: Optional[str] = None,
                 resource_settings: Optional['outputs.ResourceSettingsResponse'] = None):
        """
        Defines the move resource properties.
        :param List['MoveResourceDependencyResponseArgs'] depends_on: Gets or sets the move resource dependencies.
        :param 'MoveResourcePropertiesResponseErrorsArgs' errors: Defines the move resource errors.
        :param 'MoveResourcePropertiesResponseMoveStatusArgs' move_status: Defines the move resource status.
        :param str source_id: Gets or sets the Source ARM Id of the resource.
        :param 'MoveResourcePropertiesResponseSourceResourceSettingsArgs' source_resource_settings: Gets or sets the source resource settings.
        :param str target_id: Gets or sets the Target ARM Id of the resource.
        :param List['MoveResourceDependencyOverrideResponseArgs'] depends_on_overrides: Gets or sets the move resource dependencies overrides.
        :param str existing_target_id: Gets or sets the existing target ARM Id of the resource.
        :param str provisioning_state: Defines the provisioning states.
        :param 'ResourceSettingsResponseArgs' resource_settings: Gets or sets the resource settings.
        """
        pulumi.set(__self__, "depends_on", depends_on)
        pulumi.set(__self__, "errors", errors)
        pulumi.set(__self__, "move_status", move_status)
        pulumi.set(__self__, "source_id", source_id)
        pulumi.set(__self__, "source_resource_settings", source_resource_settings)
        pulumi.set(__self__, "target_id", target_id)
        if depends_on_overrides is not None:
            pulumi.set(__self__, "depends_on_overrides", depends_on_overrides)
        if existing_target_id is not None:
            pulumi.set(__self__, "existing_target_id", existing_target_id)
        if provisioning_state is not None:
            pulumi.set(__self__, "provisioning_state", provisioning_state)
        if resource_settings is not None:
            pulumi.set(__self__, "resource_settings", resource_settings)

    @property
    @pulumi.getter(name="dependsOn")
    def depends_on(self) -> List['outputs.MoveResourceDependencyResponse']:
        """
        Gets or sets the move resource dependencies.
        """
        return pulumi.get(self, "depends_on")

    @property
    @pulumi.getter
    def errors(self) -> 'outputs.MoveResourcePropertiesResponseErrors':
        """
        Defines the move resource errors.
        """
        return pulumi.get(self, "errors")

    @property
    @pulumi.getter(name="moveStatus")
    def move_status(self) -> 'outputs.MoveResourcePropertiesResponseMoveStatus':
        """
        Defines the move resource status.
        """
        return pulumi.get(self, "move_status")

    @property
    @pulumi.getter(name="sourceId")
    def source_id(self) -> str:
        """
        Gets or sets the Source ARM Id of the resource.
        """
        return pulumi.get(self, "source_id")

    @property
    @pulumi.getter(name="sourceResourceSettings")
    def source_resource_settings(self) -> 'outputs.MoveResourcePropertiesResponseSourceResourceSettings':
        """
        Gets or sets the source resource settings.
        """
        return pulumi.get(self, "source_resource_settings")

    @property
    @pulumi.getter(name="targetId")
    def target_id(self) -> str:
        """
        Gets or sets the Target ARM Id of the resource.
        """
        return pulumi.get(self, "target_id")

    @property
    @pulumi.getter(name="dependsOnOverrides")
    def depends_on_overrides(self) -> Optional[List['outputs.MoveResourceDependencyOverrideResponse']]:
        """
        Gets or sets the move resource dependencies overrides.
        """
        return pulumi.get(self, "depends_on_overrides")

    @property
    @pulumi.getter(name="existingTargetId")
    def existing_target_id(self) -> Optional[str]:
        """
        Gets or sets the existing target ARM Id of the resource.
        """
        return pulumi.get(self, "existing_target_id")

    @property
    @pulumi.getter(name="provisioningState")
    def provisioning_state(self) -> Optional[str]:
        """
        Defines the provisioning states.
        """
        return pulumi.get(self, "provisioning_state")

    @property
    @pulumi.getter(name="resourceSettings")
    def resource_settings(self) -> Optional['outputs.ResourceSettingsResponse']:
        """
        Gets or sets the resource settings.
        """
        return pulumi.get(self, "resource_settings")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class MoveResourcePropertiesResponseErrors(dict):
    """
    Defines the move resource errors.
    """
    def __init__(__self__, *,
                 properties: Optional['outputs.MoveResourceErrorBodyResponse'] = None):
        """
        Defines the move resource errors.
        :param 'MoveResourceErrorBodyResponseArgs' properties: The move resource error body.
        """
        if properties is not None:
            pulumi.set(__self__, "properties", properties)

    @property
    @pulumi.getter
    def properties(self) -> Optional['outputs.MoveResourceErrorBodyResponse']:
        """
        The move resource error body.
        """
        return pulumi.get(self, "properties")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class MoveResourcePropertiesResponseMoveStatus(dict):
    """
    Defines the move resource status.
    """
    def __init__(__self__, *,
                 target_id: str,
                 errors: Optional['outputs.MoveResourceErrorResponse'] = None,
                 job_status: Optional['outputs.JobStatusResponse'] = None,
                 move_state: Optional[str] = None):
        """
        Defines the move resource status.
        :param str target_id: Gets the Target ARM Id of the resource.
        :param 'MoveResourceErrorResponseArgs' errors: An error response from the azure region move service.
        :param 'JobStatusResponseArgs' job_status: Defines the job status.
        :param str move_state: Defines the MoveResource states.
        """
        pulumi.set(__self__, "target_id", target_id)
        if errors is not None:
            pulumi.set(__self__, "errors", errors)
        if job_status is not None:
            pulumi.set(__self__, "job_status", job_status)
        if move_state is not None:
            pulumi.set(__self__, "move_state", move_state)

    @property
    @pulumi.getter(name="targetId")
    def target_id(self) -> str:
        """
        Gets the Target ARM Id of the resource.
        """
        return pulumi.get(self, "target_id")

    @property
    @pulumi.getter
    def errors(self) -> Optional['outputs.MoveResourceErrorResponse']:
        """
        An error response from the azure region move service.
        """
        return pulumi.get(self, "errors")

    @property
    @pulumi.getter(name="jobStatus")
    def job_status(self) -> Optional['outputs.JobStatusResponse']:
        """
        Defines the job status.
        """
        return pulumi.get(self, "job_status")

    @property
    @pulumi.getter(name="moveState")
    def move_state(self) -> Optional[str]:
        """
        Defines the MoveResource states.
        """
        return pulumi.get(self, "move_state")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class MoveResourcePropertiesResponseSourceResourceSettings(dict):
    """
    Gets or sets the source resource settings.
    """
    def __init__(__self__, *,
                 resource_type: str,
                 target_resource_name: str):
        """
        Gets or sets the source resource settings.
        :param str resource_type: The resource type. For example, the value can be Microsoft.Compute/virtualMachines.
        :param str target_resource_name: Gets or sets the target Resource name.
        """
        pulumi.set(__self__, "resource_type", resource_type)
        pulumi.set(__self__, "target_resource_name", target_resource_name)

    @property
    @pulumi.getter(name="resourceType")
    def resource_type(self) -> str:
        """
        The resource type. For example, the value can be Microsoft.Compute/virtualMachines.
        """
        return pulumi.get(self, "resource_type")

    @property
    @pulumi.getter(name="targetResourceName")
    def target_resource_name(self) -> str:
        """
        Gets or sets the target Resource name.
        """
        return pulumi.get(self, "target_resource_name")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class ResourceSettingsResponse(dict):
    """
    Gets or sets the resource settings.
    """
    def __init__(__self__, *,
                 resource_type: str,
                 target_resource_name: str):
        """
        Gets or sets the resource settings.
        :param str resource_type: The resource type. For example, the value can be Microsoft.Compute/virtualMachines.
        :param str target_resource_name: Gets or sets the target Resource name.
        """
        pulumi.set(__self__, "resource_type", resource_type)
        pulumi.set(__self__, "target_resource_name", target_resource_name)

    @property
    @pulumi.getter(name="resourceType")
    def resource_type(self) -> str:
        """
        The resource type. For example, the value can be Microsoft.Compute/virtualMachines.
        """
        return pulumi.get(self, "resource_type")

    @property
    @pulumi.getter(name="targetResourceName")
    def target_resource_name(self) -> str:
        """
        Gets or sets the target Resource name.
        """
        return pulumi.get(self, "target_resource_name")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

