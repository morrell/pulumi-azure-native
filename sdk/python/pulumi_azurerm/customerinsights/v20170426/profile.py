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

__all__ = ['Profile']


class Profile(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 api_entity_set_name: Optional[pulumi.Input[str]] = None,
                 attributes: Optional[pulumi.Input[Mapping[str, pulumi.Input[List[pulumi.Input[str]]]]]] = None,
                 description: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 display_name: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 entity_type: Optional[pulumi.Input[str]] = None,
                 fields: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['PropertyDefinitionArgs']]]]] = None,
                 hub_name: Optional[pulumi.Input[str]] = None,
                 instances_count: Optional[pulumi.Input[float]] = None,
                 large_image: Optional[pulumi.Input[str]] = None,
                 localized_attributes: Optional[pulumi.Input[Mapping[str, pulumi.Input[Mapping[str, pulumi.Input[str]]]]]] = None,
                 medium_image: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 schema_item_type_link: Optional[pulumi.Input[str]] = None,
                 small_image: Optional[pulumi.Input[str]] = None,
                 strong_ids: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['StrongIdArgs']]]]] = None,
                 timestamp_field_name: Optional[pulumi.Input[str]] = None,
                 type_name: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        The profile resource format.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] api_entity_set_name: The api entity set name. This becomes the odata entity set name for the entity Type being referred in this object.
        :param pulumi.Input[Mapping[str, pulumi.Input[List[pulumi.Input[str]]]]] attributes: The attributes for the Type.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] description: Localized descriptions for the property.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] display_name: Localized display names for the property.
        :param pulumi.Input[str] entity_type: Type of entity.
        :param pulumi.Input[List[pulumi.Input[pulumi.InputType['PropertyDefinitionArgs']]]] fields: The properties of the Profile.
        :param pulumi.Input[str] hub_name: The name of the hub.
        :param pulumi.Input[float] instances_count: The instance count.
        :param pulumi.Input[str] large_image: Large Image associated with the Property or EntityType.
        :param pulumi.Input[Mapping[str, pulumi.Input[Mapping[str, pulumi.Input[str]]]]] localized_attributes: Any custom localized attributes for the Type.
        :param pulumi.Input[str] medium_image: Medium Image associated with the Property or EntityType.
        :param pulumi.Input[str] name: The name of the profile.
        :param pulumi.Input[str] resource_group_name: The name of the resource group.
        :param pulumi.Input[str] schema_item_type_link: The schema org link. This helps ACI identify and suggest semantic models.
        :param pulumi.Input[str] small_image: Small Image associated with the Property or EntityType.
        :param pulumi.Input[List[pulumi.Input[pulumi.InputType['StrongIdArgs']]]] strong_ids: The strong IDs.
        :param pulumi.Input[str] timestamp_field_name: The timestamp property name. Represents the time when the interaction or profile update happened.
        :param pulumi.Input[str] type_name: The name of the entity.
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

            __props__['api_entity_set_name'] = api_entity_set_name
            __props__['attributes'] = attributes
            __props__['description'] = description
            __props__['display_name'] = display_name
            __props__['entity_type'] = entity_type
            __props__['fields'] = fields
            if hub_name is None:
                raise TypeError("Missing required property 'hub_name'")
            __props__['hub_name'] = hub_name
            __props__['instances_count'] = instances_count
            __props__['large_image'] = large_image
            __props__['localized_attributes'] = localized_attributes
            __props__['medium_image'] = medium_image
            if name is None:
                raise TypeError("Missing required property 'name'")
            __props__['name'] = name
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['schema_item_type_link'] = schema_item_type_link
            __props__['small_image'] = small_image
            __props__['strong_ids'] = strong_ids
            __props__['timestamp_field_name'] = timestamp_field_name
            __props__['type_name'] = type_name
            __props__['last_changed_utc'] = None
            __props__['provisioning_state'] = None
            __props__['tenant_id'] = None
            __props__['type'] = None
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="azurerm:customerinsights/v20170101:Profile")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(Profile, __self__).__init__(
            'azurerm:customerinsights/v20170426:Profile',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'Profile':
        """
        Get an existing Profile resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return Profile(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="apiEntitySetName")
    def api_entity_set_name(self) -> Optional[str]:
        """
        The api entity set name. This becomes the odata entity set name for the entity Type being referred in this object.
        """
        return pulumi.get(self, "api_entity_set_name")

    @property
    @pulumi.getter
    def attributes(self) -> Optional[Mapping[str, List[str]]]:
        """
        The attributes for the Type.
        """
        return pulumi.get(self, "attributes")

    @property
    @pulumi.getter
    def description(self) -> Optional[Mapping[str, str]]:
        """
        Localized descriptions for the property.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> Optional[Mapping[str, str]]:
        """
        Localized display names for the property.
        """
        return pulumi.get(self, "display_name")

    @property
    @pulumi.getter(name="entityType")
    def entity_type(self) -> Optional[str]:
        """
        Type of entity.
        """
        return pulumi.get(self, "entity_type")

    @property
    @pulumi.getter
    def fields(self) -> Optional[List['outputs.PropertyDefinitionResponse']]:
        """
        The properties of the Profile.
        """
        return pulumi.get(self, "fields")

    @property
    @pulumi.getter(name="instancesCount")
    def instances_count(self) -> Optional[float]:
        """
        The instance count.
        """
        return pulumi.get(self, "instances_count")

    @property
    @pulumi.getter(name="largeImage")
    def large_image(self) -> Optional[str]:
        """
        Large Image associated with the Property or EntityType.
        """
        return pulumi.get(self, "large_image")

    @property
    @pulumi.getter(name="lastChangedUtc")
    def last_changed_utc(self) -> str:
        """
        The last changed time for the type definition.
        """
        return pulumi.get(self, "last_changed_utc")

    @property
    @pulumi.getter(name="localizedAttributes")
    def localized_attributes(self) -> Optional[Mapping[str, Mapping[str, str]]]:
        """
        Any custom localized attributes for the Type.
        """
        return pulumi.get(self, "localized_attributes")

    @property
    @pulumi.getter(name="mediumImage")
    def medium_image(self) -> Optional[str]:
        """
        Medium Image associated with the Property or EntityType.
        """
        return pulumi.get(self, "medium_image")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Resource name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="provisioningState")
    def provisioning_state(self) -> str:
        """
        Provisioning state.
        """
        return pulumi.get(self, "provisioning_state")

    @property
    @pulumi.getter(name="schemaItemTypeLink")
    def schema_item_type_link(self) -> Optional[str]:
        """
        The schema org link. This helps ACI identify and suggest semantic models.
        """
        return pulumi.get(self, "schema_item_type_link")

    @property
    @pulumi.getter(name="smallImage")
    def small_image(self) -> Optional[str]:
        """
        Small Image associated with the Property or EntityType.
        """
        return pulumi.get(self, "small_image")

    @property
    @pulumi.getter(name="strongIds")
    def strong_ids(self) -> Optional[List['outputs.StrongIdResponse']]:
        """
        The strong IDs.
        """
        return pulumi.get(self, "strong_ids")

    @property
    @pulumi.getter(name="tenantId")
    def tenant_id(self) -> str:
        """
        The hub name.
        """
        return pulumi.get(self, "tenant_id")

    @property
    @pulumi.getter(name="timestampFieldName")
    def timestamp_field_name(self) -> Optional[str]:
        """
        The timestamp property name. Represents the time when the interaction or profile update happened.
        """
        return pulumi.get(self, "timestamp_field_name")

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        Resource type.
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter(name="typeName")
    def type_name(self) -> Optional[str]:
        """
        The name of the entity.
        """
        return pulumi.get(self, "type_name")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
