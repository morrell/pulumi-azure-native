# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from ... import _utilities, _tables
from . import outputs

__all__ = ['IpAllocation']


class IpAllocation(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 allocation_tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 id: Optional[pulumi.Input[str]] = None,
                 ipam_allocation_id: Optional[pulumi.Input[str]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 prefix: Optional[pulumi.Input[str]] = None,
                 prefix_length: Optional[pulumi.Input[float]] = None,
                 prefix_type: Optional[pulumi.Input[str]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        IpAllocation resource.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] allocation_tags: IpAllocation tags.
        :param pulumi.Input[str] id: Resource ID.
        :param pulumi.Input[str] ipam_allocation_id: The IPAM allocation ID.
        :param pulumi.Input[str] location: Resource location.
        :param pulumi.Input[str] name: The name of the IpAllocation.
        :param pulumi.Input[str] prefix: The address prefix for the IpAllocation.
        :param pulumi.Input[float] prefix_length: The address prefix length for the IpAllocation.
        :param pulumi.Input[str] prefix_type: The address prefix Type for the IpAllocation.
        :param pulumi.Input[str] resource_group_name: The name of the resource group.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Resource tags.
        :param pulumi.Input[str] type: The type for the IpAllocation.
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

            __props__['allocation_tags'] = allocation_tags
            __props__['id'] = id
            __props__['ipam_allocation_id'] = ipam_allocation_id
            __props__['location'] = location
            if name is None:
                raise TypeError("Missing required property 'name'")
            __props__['name'] = name
            __props__['prefix'] = prefix
            __props__['prefix_length'] = prefix_length
            __props__['prefix_type'] = prefix_type
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['tags'] = tags
            __props__['type'] = type
            __props__['etag'] = None
            __props__['subnet'] = None
            __props__['virtual_network'] = None
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="azurerm:network/v20200301:IpAllocation"), pulumi.Alias(type_="azurerm:network/v20200501:IpAllocation"), pulumi.Alias(type_="azurerm:network/v20200601:IpAllocation")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(IpAllocation, __self__).__init__(
            'azurerm:network/v20200401:IpAllocation',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'IpAllocation':
        """
        Get an existing IpAllocation resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return IpAllocation(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="allocationTags")
    def allocation_tags(self) -> Optional[Mapping[str, str]]:
        """
        IpAllocation tags.
        """
        return pulumi.get(self, "allocation_tags")

    @property
    @pulumi.getter
    def etag(self) -> str:
        """
        A unique read-only string that changes whenever the resource is updated.
        """
        return pulumi.get(self, "etag")

    @property
    @pulumi.getter(name="ipamAllocationId")
    def ipam_allocation_id(self) -> Optional[str]:
        """
        The IPAM allocation ID.
        """
        return pulumi.get(self, "ipam_allocation_id")

    @property
    @pulumi.getter
    def location(self) -> Optional[str]:
        """
        Resource location.
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Resource name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def prefix(self) -> Optional[str]:
        """
        The address prefix for the IpAllocation.
        """
        return pulumi.get(self, "prefix")

    @property
    @pulumi.getter(name="prefixLength")
    def prefix_length(self) -> Optional[float]:
        """
        The address prefix length for the IpAllocation.
        """
        return pulumi.get(self, "prefix_length")

    @property
    @pulumi.getter(name="prefixType")
    def prefix_type(self) -> Optional[str]:
        """
        The address prefix Type for the IpAllocation.
        """
        return pulumi.get(self, "prefix_type")

    @property
    @pulumi.getter
    def subnet(self) -> 'outputs.SubResourceResponse':
        """
        The Subnet that using the prefix of this IpAllocation resource.
        """
        return pulumi.get(self, "subnet")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Mapping[str, str]]:
        """
        Resource tags.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        Resource type.
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter(name="virtualNetwork")
    def virtual_network(self) -> 'outputs.SubResourceResponse':
        """
        The VirtualNetwork that using the prefix of this IpAllocation resource.
        """
        return pulumi.get(self, "virtual_network")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
