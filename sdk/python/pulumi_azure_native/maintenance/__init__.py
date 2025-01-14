# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

from .. import _utilities
import typing
# Export this package's modules as members:
from ._enums import *
from .configuration_assignment import *
from .configuration_assignment_parent import *
from .get_configuration_assignment import *
from .get_configuration_assignment_parent import *
from .get_maintenance_configuration import *
from .maintenance_configuration import *
from . import outputs

# Make subpackages available:
if typing.TYPE_CHECKING:
    import pulumi_azure_native.maintenance.v20180601preview as v20180601preview
    import pulumi_azure_native.maintenance.v20200401 as v20200401
    import pulumi_azure_native.maintenance.v20200701preview as v20200701preview
    import pulumi_azure_native.maintenance.v20210401preview as v20210401preview
    import pulumi_azure_native.maintenance.v20210501 as v20210501
else:
    v20180601preview = _utilities.lazy_import('pulumi_azure_native.maintenance.v20180601preview')
    v20200401 = _utilities.lazy_import('pulumi_azure_native.maintenance.v20200401')
    v20200701preview = _utilities.lazy_import('pulumi_azure_native.maintenance.v20200701preview')
    v20210401preview = _utilities.lazy_import('pulumi_azure_native.maintenance.v20210401preview')
    v20210501 = _utilities.lazy_import('pulumi_azure_native.maintenance.v20210501')

