# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from ... import _utilities, _tables

__all__ = ['ManagerExtendedInfo']


class ManagerExtendedInfo(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 algorithm: Optional[pulumi.Input[str]] = None,
                 encryption_key: Optional[pulumi.Input[str]] = None,
                 encryption_key_thumbprint: Optional[pulumi.Input[str]] = None,
                 etag: Optional[pulumi.Input[str]] = None,
                 integrity_key: Optional[pulumi.Input[str]] = None,
                 kind: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 portal_certificate_thumbprint: Optional[pulumi.Input[str]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 version: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        The extended info of the manager.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] algorithm: Represents the encryption algorithm used to encrypt the keys. None - if Key is saved in plain text format. Algorithm name - if key is encrypted
        :param pulumi.Input[str] encryption_key: Represents the CEK of the resource.
        :param pulumi.Input[str] encryption_key_thumbprint: Represents the Cert thumbprint that was used to encrypt the CEK.
        :param pulumi.Input[str] etag: The etag of the resource.
        :param pulumi.Input[str] integrity_key: Represents the CIK of the resource.
        :param pulumi.Input[str] kind: The Kind of the object. Currently only Series8000 is supported
        :param pulumi.Input[str] name: The manager name
        :param pulumi.Input[str] portal_certificate_thumbprint: Represents the portal thumbprint which can be used optionally to encrypt the entire data before storing it.
        :param pulumi.Input[str] resource_group_name: The resource group name
        :param pulumi.Input[str] version: The version of the extended info being persisted.
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

            if algorithm is None:
                raise TypeError("Missing required property 'algorithm'")
            __props__['algorithm'] = algorithm
            __props__['encryption_key'] = encryption_key
            __props__['encryption_key_thumbprint'] = encryption_key_thumbprint
            __props__['etag'] = etag
            if integrity_key is None:
                raise TypeError("Missing required property 'integrity_key'")
            __props__['integrity_key'] = integrity_key
            __props__['kind'] = kind
            if name is None:
                raise TypeError("Missing required property 'name'")
            __props__['name'] = name
            __props__['portal_certificate_thumbprint'] = portal_certificate_thumbprint
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['version'] = version
            __props__['type'] = None
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="azurerm:storsimple/v20161001:ManagerExtendedInfo")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(ManagerExtendedInfo, __self__).__init__(
            'azurerm:storsimple/v20170601:ManagerExtendedInfo',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'ManagerExtendedInfo':
        """
        Get an existing ManagerExtendedInfo resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return ManagerExtendedInfo(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def algorithm(self) -> str:
        """
        Represents the encryption algorithm used to encrypt the keys. None - if Key is saved in plain text format. Algorithm name - if key is encrypted
        """
        return pulumi.get(self, "algorithm")

    @property
    @pulumi.getter(name="encryptionKey")
    def encryption_key(self) -> Optional[str]:
        """
        Represents the CEK of the resource.
        """
        return pulumi.get(self, "encryption_key")

    @property
    @pulumi.getter(name="encryptionKeyThumbprint")
    def encryption_key_thumbprint(self) -> Optional[str]:
        """
        Represents the Cert thumbprint that was used to encrypt the CEK.
        """
        return pulumi.get(self, "encryption_key_thumbprint")

    @property
    @pulumi.getter
    def etag(self) -> Optional[str]:
        """
        The etag of the resource.
        """
        return pulumi.get(self, "etag")

    @property
    @pulumi.getter(name="integrityKey")
    def integrity_key(self) -> str:
        """
        Represents the CIK of the resource.
        """
        return pulumi.get(self, "integrity_key")

    @property
    @pulumi.getter
    def kind(self) -> Optional[str]:
        """
        The Kind of the object. Currently only Series8000 is supported
        """
        return pulumi.get(self, "kind")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The name of the object.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="portalCertificateThumbprint")
    def portal_certificate_thumbprint(self) -> Optional[str]:
        """
        Represents the portal thumbprint which can be used optionally to encrypt the entire data before storing it.
        """
        return pulumi.get(self, "portal_certificate_thumbprint")

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        The hierarchical type of the object.
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter
    def version(self) -> Optional[str]:
        """
        The version of the extended info being persisted.
        """
        return pulumi.get(self, "version")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
