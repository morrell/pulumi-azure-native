# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from ... import _utilities, _tables

__all__ = [
    'CustomRPActionRouteDefinitionResponse',
    'CustomRPResourceTypeRouteDefinitionResponse',
    'CustomRPValidationsResponse',
]

@pulumi.output_type
class CustomRPActionRouteDefinitionResponse(dict):
    """
    The route definition for an action implemented by the custom resource provider.
    """
    def __init__(__self__, *,
                 endpoint: str,
                 name: str,
                 routing_type: Optional[str] = None):
        """
        The route definition for an action implemented by the custom resource provider.
        :param str endpoint: The route definition endpoint URI that the custom resource provider will proxy requests to. This can be in the form of a flat URI (e.g. 'https://testendpoint/') or can specify to route via a path (e.g. 'https://testendpoint/{requestPath}')
        :param str name: The name of the route definition. This becomes the name for the ARM extension (e.g. '/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CustomProviders/resourceProviders/{resourceProviderName}/{name}')
        :param str routing_type: The routing types that are supported for action requests.
        """
        pulumi.set(__self__, "endpoint", endpoint)
        pulumi.set(__self__, "name", name)
        if routing_type is not None:
            pulumi.set(__self__, "routing_type", routing_type)

    @property
    @pulumi.getter
    def endpoint(self) -> str:
        """
        The route definition endpoint URI that the custom resource provider will proxy requests to. This can be in the form of a flat URI (e.g. 'https://testendpoint/') or can specify to route via a path (e.g. 'https://testendpoint/{requestPath}')
        """
        return pulumi.get(self, "endpoint")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The name of the route definition. This becomes the name for the ARM extension (e.g. '/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CustomProviders/resourceProviders/{resourceProviderName}/{name}')
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="routingType")
    def routing_type(self) -> Optional[str]:
        """
        The routing types that are supported for action requests.
        """
        return pulumi.get(self, "routing_type")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class CustomRPResourceTypeRouteDefinitionResponse(dict):
    """
    The route definition for a resource implemented by the custom resource provider.
    """
    def __init__(__self__, *,
                 endpoint: str,
                 name: str,
                 routing_type: Optional[str] = None):
        """
        The route definition for a resource implemented by the custom resource provider.
        :param str endpoint: The route definition endpoint URI that the custom resource provider will proxy requests to. This can be in the form of a flat URI (e.g. 'https://testendpoint/') or can specify to route via a path (e.g. 'https://testendpoint/{requestPath}')
        :param str name: The name of the route definition. This becomes the name for the ARM extension (e.g. '/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CustomProviders/resourceProviders/{resourceProviderName}/{name}')
        :param str routing_type: The routing types that are supported for resource requests.
        """
        pulumi.set(__self__, "endpoint", endpoint)
        pulumi.set(__self__, "name", name)
        if routing_type is not None:
            pulumi.set(__self__, "routing_type", routing_type)

    @property
    @pulumi.getter
    def endpoint(self) -> str:
        """
        The route definition endpoint URI that the custom resource provider will proxy requests to. This can be in the form of a flat URI (e.g. 'https://testendpoint/') or can specify to route via a path (e.g. 'https://testendpoint/{requestPath}')
        """
        return pulumi.get(self, "endpoint")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The name of the route definition. This becomes the name for the ARM extension (e.g. '/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CustomProviders/resourceProviders/{resourceProviderName}/{name}')
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="routingType")
    def routing_type(self) -> Optional[str]:
        """
        The routing types that are supported for resource requests.
        """
        return pulumi.get(self, "routing_type")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class CustomRPValidationsResponse(dict):
    """
    A validation to apply on custom resource provider requests.
    """
    def __init__(__self__, *,
                 specification: str,
                 validation_type: Optional[str] = None):
        """
        A validation to apply on custom resource provider requests.
        :param str specification: A link to the validation specification. The specification must be hosted on raw.githubusercontent.com.
        :param str validation_type: The type of validation to run against a matching request.
        """
        pulumi.set(__self__, "specification", specification)
        if validation_type is not None:
            pulumi.set(__self__, "validation_type", validation_type)

    @property
    @pulumi.getter
    def specification(self) -> str:
        """
        A link to the validation specification. The specification must be hosted on raw.githubusercontent.com.
        """
        return pulumi.get(self, "specification")

    @property
    @pulumi.getter(name="validationType")
    def validation_type(self) -> Optional[str]:
        """
        The type of validation to run against a matching request.
        """
        return pulumi.get(self, "validation_type")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

