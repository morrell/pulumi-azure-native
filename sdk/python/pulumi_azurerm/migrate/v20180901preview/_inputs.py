# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from ... import _utilities, _tables

__all__ = [
    'MigrateProjectPropertiesArgs',
    'MigrateProjectTagsArgs',
    'SolutionDetailsArgs',
    'SolutionPropertiesArgs',
]

@pulumi.input_type
class MigrateProjectPropertiesArgs:
    def __init__(__self__, *,
                 provisioning_state: Optional[pulumi.Input[str]] = None,
                 registered_tools: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None):
        """
        Class for migrate project properties.
        :param pulumi.Input[str] provisioning_state: Provisioning state of the migrate project.
        :param pulumi.Input[List[pulumi.Input[str]]] registered_tools: Gets or sets the list of tools registered with the migrate project.
        """
        if provisioning_state is not None:
            pulumi.set(__self__, "provisioning_state", provisioning_state)
        if registered_tools is not None:
            pulumi.set(__self__, "registered_tools", registered_tools)

    @property
    @pulumi.getter(name="provisioningState")
    def provisioning_state(self) -> Optional[pulumi.Input[str]]:
        """
        Provisioning state of the migrate project.
        """
        return pulumi.get(self, "provisioning_state")

    @provisioning_state.setter
    def provisioning_state(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "provisioning_state", value)

    @property
    @pulumi.getter(name="registeredTools")
    def registered_tools(self) -> Optional[pulumi.Input[List[pulumi.Input[str]]]]:
        """
        Gets or sets the list of tools registered with the migrate project.
        """
        return pulumi.get(self, "registered_tools")

    @registered_tools.setter
    def registered_tools(self, value: Optional[pulumi.Input[List[pulumi.Input[str]]]]):
        pulumi.set(self, "registered_tools", value)


@pulumi.input_type
class MigrateProjectTagsArgs:
    def __init__(__self__, *,
                 additional_properties: Optional[pulumi.Input[str]] = None):
        """
        Gets or sets the tags.
        """
        if additional_properties is not None:
            pulumi.set(__self__, "additional_properties", additional_properties)

    @property
    @pulumi.getter(name="additionalProperties")
    def additional_properties(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "additional_properties")

    @additional_properties.setter
    def additional_properties(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "additional_properties", value)


@pulumi.input_type
class SolutionDetailsArgs:
    def __init__(__self__, *,
                 assessment_count: Optional[pulumi.Input[float]] = None,
                 extended_details: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 group_count: Optional[pulumi.Input[float]] = None):
        """
        Class representing the details of the solution.
        :param pulumi.Input[float] assessment_count: Gets or sets the count of assessments reported by the solution.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] extended_details: Gets or sets the extended details reported by the solution.
        :param pulumi.Input[float] group_count: Gets or sets the count of groups reported by the solution.
        """
        if assessment_count is not None:
            pulumi.set(__self__, "assessment_count", assessment_count)
        if extended_details is not None:
            pulumi.set(__self__, "extended_details", extended_details)
        if group_count is not None:
            pulumi.set(__self__, "group_count", group_count)

    @property
    @pulumi.getter(name="assessmentCount")
    def assessment_count(self) -> Optional[pulumi.Input[float]]:
        """
        Gets or sets the count of assessments reported by the solution.
        """
        return pulumi.get(self, "assessment_count")

    @assessment_count.setter
    def assessment_count(self, value: Optional[pulumi.Input[float]]):
        pulumi.set(self, "assessment_count", value)

    @property
    @pulumi.getter(name="extendedDetails")
    def extended_details(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        Gets or sets the extended details reported by the solution.
        """
        return pulumi.get(self, "extended_details")

    @extended_details.setter
    def extended_details(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "extended_details", value)

    @property
    @pulumi.getter(name="groupCount")
    def group_count(self) -> Optional[pulumi.Input[float]]:
        """
        Gets or sets the count of groups reported by the solution.
        """
        return pulumi.get(self, "group_count")

    @group_count.setter
    def group_count(self, value: Optional[pulumi.Input[float]]):
        pulumi.set(self, "group_count", value)


@pulumi.input_type
class SolutionPropertiesArgs:
    def __init__(__self__, *,
                 cleanup_state: Optional[pulumi.Input[str]] = None,
                 details: Optional[pulumi.Input['SolutionDetailsArgs']] = None,
                 goal: Optional[pulumi.Input[str]] = None,
                 purpose: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 tool: Optional[pulumi.Input[str]] = None):
        """
        Class for solution properties.
        :param pulumi.Input[str] cleanup_state: Gets or sets the cleanup state of the solution.
        :param pulumi.Input['SolutionDetailsArgs'] details: Gets or sets the details of the solution.
        :param pulumi.Input[str] goal: Gets or sets the goal of the solution.
        :param pulumi.Input[str] purpose: Gets or sets the purpose of the solution.
        :param pulumi.Input[str] status: Gets or sets the current status of the solution.
        :param pulumi.Input[str] tool: Gets or sets the tool being used in the solution.
        """
        if cleanup_state is not None:
            pulumi.set(__self__, "cleanup_state", cleanup_state)
        if details is not None:
            pulumi.set(__self__, "details", details)
        if goal is not None:
            pulumi.set(__self__, "goal", goal)
        if purpose is not None:
            pulumi.set(__self__, "purpose", purpose)
        if status is not None:
            pulumi.set(__self__, "status", status)
        if tool is not None:
            pulumi.set(__self__, "tool", tool)

    @property
    @pulumi.getter(name="cleanupState")
    def cleanup_state(self) -> Optional[pulumi.Input[str]]:
        """
        Gets or sets the cleanup state of the solution.
        """
        return pulumi.get(self, "cleanup_state")

    @cleanup_state.setter
    def cleanup_state(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "cleanup_state", value)

    @property
    @pulumi.getter
    def details(self) -> Optional[pulumi.Input['SolutionDetailsArgs']]:
        """
        Gets or sets the details of the solution.
        """
        return pulumi.get(self, "details")

    @details.setter
    def details(self, value: Optional[pulumi.Input['SolutionDetailsArgs']]):
        pulumi.set(self, "details", value)

    @property
    @pulumi.getter
    def goal(self) -> Optional[pulumi.Input[str]]:
        """
        Gets or sets the goal of the solution.
        """
        return pulumi.get(self, "goal")

    @goal.setter
    def goal(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "goal", value)

    @property
    @pulumi.getter
    def purpose(self) -> Optional[pulumi.Input[str]]:
        """
        Gets or sets the purpose of the solution.
        """
        return pulumi.get(self, "purpose")

    @purpose.setter
    def purpose(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "purpose", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[str]]:
        """
        Gets or sets the current status of the solution.
        """
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "status", value)

    @property
    @pulumi.getter
    def tool(self) -> Optional[pulumi.Input[str]]:
        """
        Gets or sets the tool being used in the solution.
        """
        return pulumi.get(self, "tool")

    @tool.setter
    def tool(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "tool", value)

