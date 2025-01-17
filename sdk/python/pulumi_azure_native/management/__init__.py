# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

from .. import _utilities
import typing
# Export this package's modules as members:
from .get_entity import *
from .get_hierarchy_setting import *
from .get_management_group import *
from .get_management_group_subscription import *
from .hierarchy_setting import *
from .management_group import *
from .management_group_subscription import *
from ._inputs import *
from . import outputs

# Make subpackages available:
if typing.TYPE_CHECKING:
    import pulumi_azure_native.management.v20171101preview as v20171101preview
    import pulumi_azure_native.management.v20180101preview as v20180101preview
    import pulumi_azure_native.management.v20180301preview as v20180301preview
    import pulumi_azure_native.management.v20191101 as v20191101
    import pulumi_azure_native.management.v20200201 as v20200201
    import pulumi_azure_native.management.v20200501 as v20200501
    import pulumi_azure_native.management.v20201001 as v20201001
    import pulumi_azure_native.management.v20210401 as v20210401
else:
    v20171101preview = _utilities.lazy_import('pulumi_azure_native.management.v20171101preview')
    v20180101preview = _utilities.lazy_import('pulumi_azure_native.management.v20180101preview')
    v20180301preview = _utilities.lazy_import('pulumi_azure_native.management.v20180301preview')
    v20191101 = _utilities.lazy_import('pulumi_azure_native.management.v20191101')
    v20200201 = _utilities.lazy_import('pulumi_azure_native.management.v20200201')
    v20200501 = _utilities.lazy_import('pulumi_azure_native.management.v20200501')
    v20201001 = _utilities.lazy_import('pulumi_azure_native.management.v20201001')
    v20210401 = _utilities.lazy_import('pulumi_azure_native.management.v20210401')

