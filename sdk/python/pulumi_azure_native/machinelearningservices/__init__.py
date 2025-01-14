# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

from .. import _utilities
import typing
# Export this package's modules as members:
from ._enums import *
from .aci_service import *
from .aks_service import *
from .batch_deployment import *
from .batch_endpoint import *
from .code_container import *
from .code_version import *
from .data_container import *
from .data_version import *
from .endpoint_variant import *
from .environment_container import *
from .environment_specification_version import *
from .get_aci_service import *
from .get_aks_service import *
from .get_batch_deployment import *
from .get_batch_endpoint import *
from .get_code_container import *
from .get_code_version import *
from .get_data_container import *
from .get_data_version import *
from .get_endpoint_variant import *
from .get_environment_container import *
from .get_environment_specification_version import *
from .get_job import *
from .get_labeling_job import *
from .get_linked_service import *
from .get_linked_workspace import *
from .get_machine_learning_compute import *
from .get_machine_learning_dataset import *
from .get_machine_learning_datastore import *
from .get_machine_learning_service import *
from .get_model_container import *
from .get_model_version import *
from .get_online_deployment import *
from .get_online_deployment_logs import *
from .get_online_endpoint import *
from .get_online_endpoint_token import *
from .get_private_endpoint_connection import *
from .get_workspace import *
from .get_workspace_connection import *
from .job import *
from .labeling_job import *
from .linked_service import *
from .linked_workspace import *
from .list_batch_endpoint_keys import *
from .list_datastore_secrets import *
from .list_machine_learning_compute_keys import *
from .list_machine_learning_compute_nodes import *
from .list_notebook_keys import *
from .list_online_endpoint_keys import *
from .list_storage_account_keys import *
from .list_workspace_keys import *
from .list_workspace_notebook_access_token import *
from .machine_learning_compute import *
from .machine_learning_dataset import *
from .machine_learning_datastore import *
from .machine_learning_service import *
from .model_container import *
from .model_version import *
from .online_deployment import *
from .online_endpoint import *
from .private_endpoint_connection import *
from .workspace import *
from .workspace_connection import *
from ._inputs import *
from . import outputs

# Make subpackages available:
if typing.TYPE_CHECKING:
    import pulumi_azure_native.machinelearningservices.v20180301preview as v20180301preview
    import pulumi_azure_native.machinelearningservices.v20181119 as v20181119
    import pulumi_azure_native.machinelearningservices.v20190501 as v20190501
    import pulumi_azure_native.machinelearningservices.v20190601 as v20190601
    import pulumi_azure_native.machinelearningservices.v20191101 as v20191101
    import pulumi_azure_native.machinelearningservices.v20200101 as v20200101
    import pulumi_azure_native.machinelearningservices.v20200218preview as v20200218preview
    import pulumi_azure_native.machinelearningservices.v20200301 as v20200301
    import pulumi_azure_native.machinelearningservices.v20200401 as v20200401
    import pulumi_azure_native.machinelearningservices.v20200501preview as v20200501preview
    import pulumi_azure_native.machinelearningservices.v20200515preview as v20200515preview
    import pulumi_azure_native.machinelearningservices.v20200601 as v20200601
    import pulumi_azure_native.machinelearningservices.v20200801 as v20200801
    import pulumi_azure_native.machinelearningservices.v20200901preview as v20200901preview
    import pulumi_azure_native.machinelearningservices.v20210101 as v20210101
    import pulumi_azure_native.machinelearningservices.v20210301preview as v20210301preview
    import pulumi_azure_native.machinelearningservices.v20210401 as v20210401
else:
    v20180301preview = _utilities.lazy_import('pulumi_azure_native.machinelearningservices.v20180301preview')
    v20181119 = _utilities.lazy_import('pulumi_azure_native.machinelearningservices.v20181119')
    v20190501 = _utilities.lazy_import('pulumi_azure_native.machinelearningservices.v20190501')
    v20190601 = _utilities.lazy_import('pulumi_azure_native.machinelearningservices.v20190601')
    v20191101 = _utilities.lazy_import('pulumi_azure_native.machinelearningservices.v20191101')
    v20200101 = _utilities.lazy_import('pulumi_azure_native.machinelearningservices.v20200101')
    v20200218preview = _utilities.lazy_import('pulumi_azure_native.machinelearningservices.v20200218preview')
    v20200301 = _utilities.lazy_import('pulumi_azure_native.machinelearningservices.v20200301')
    v20200401 = _utilities.lazy_import('pulumi_azure_native.machinelearningservices.v20200401')
    v20200501preview = _utilities.lazy_import('pulumi_azure_native.machinelearningservices.v20200501preview')
    v20200515preview = _utilities.lazy_import('pulumi_azure_native.machinelearningservices.v20200515preview')
    v20200601 = _utilities.lazy_import('pulumi_azure_native.machinelearningservices.v20200601')
    v20200801 = _utilities.lazy_import('pulumi_azure_native.machinelearningservices.v20200801')
    v20200901preview = _utilities.lazy_import('pulumi_azure_native.machinelearningservices.v20200901preview')
    v20210101 = _utilities.lazy_import('pulumi_azure_native.machinelearningservices.v20210101')
    v20210301preview = _utilities.lazy_import('pulumi_azure_native.machinelearningservices.v20210301preview')
    v20210401 = _utilities.lazy_import('pulumi_azure_native.machinelearningservices.v20210401')

