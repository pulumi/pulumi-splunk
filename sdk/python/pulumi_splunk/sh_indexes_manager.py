# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities
from . import outputs
from ._inputs import *

__all__ = ['ShIndexesManagerArgs', 'ShIndexesManager']

@pulumi.input_type
class ShIndexesManagerArgs:
    def __init__(__self__, *,
                 acl: Optional[pulumi.Input['ShIndexesManagerAclArgs']] = None,
                 datatype: Optional[pulumi.Input[str]] = None,
                 frozen_time_period_in_secs: Optional[pulumi.Input[str]] = None,
                 max_global_raw_data_size_mb: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a ShIndexesManager resource.
        :param pulumi.Input[str] datatype: Valid values: (event | metric). Specifies the type of index.
        :param pulumi.Input[str] frozen_time_period_in_secs: Number of seconds after which indexed data rolls to frozen.
               Defaults to 94608000 (3 years).Freezing data means it is removed from the index. If you need to archive your data, refer to coldToFrozenDir and coldToFrozenScript parameter documentation.
        :param pulumi.Input[str] max_global_raw_data_size_mb: The maximum size of an index (in MB). If an index grows larger than the maximum size, the oldest data is frozen.
               Defaults to 100 MB.
        :param pulumi.Input[str] name: The name of the index to create.
        """
        if acl is not None:
            pulumi.set(__self__, "acl", acl)
        if datatype is not None:
            pulumi.set(__self__, "datatype", datatype)
        if frozen_time_period_in_secs is not None:
            pulumi.set(__self__, "frozen_time_period_in_secs", frozen_time_period_in_secs)
        if max_global_raw_data_size_mb is not None:
            pulumi.set(__self__, "max_global_raw_data_size_mb", max_global_raw_data_size_mb)
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def acl(self) -> Optional[pulumi.Input['ShIndexesManagerAclArgs']]:
        return pulumi.get(self, "acl")

    @acl.setter
    def acl(self, value: Optional[pulumi.Input['ShIndexesManagerAclArgs']]):
        pulumi.set(self, "acl", value)

    @property
    @pulumi.getter
    def datatype(self) -> Optional[pulumi.Input[str]]:
        """
        Valid values: (event | metric). Specifies the type of index.
        """
        return pulumi.get(self, "datatype")

    @datatype.setter
    def datatype(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "datatype", value)

    @property
    @pulumi.getter(name="frozenTimePeriodInSecs")
    def frozen_time_period_in_secs(self) -> Optional[pulumi.Input[str]]:
        """
        Number of seconds after which indexed data rolls to frozen.
        Defaults to 94608000 (3 years).Freezing data means it is removed from the index. If you need to archive your data, refer to coldToFrozenDir and coldToFrozenScript parameter documentation.
        """
        return pulumi.get(self, "frozen_time_period_in_secs")

    @frozen_time_period_in_secs.setter
    def frozen_time_period_in_secs(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "frozen_time_period_in_secs", value)

    @property
    @pulumi.getter(name="maxGlobalRawDataSizeMb")
    def max_global_raw_data_size_mb(self) -> Optional[pulumi.Input[str]]:
        """
        The maximum size of an index (in MB). If an index grows larger than the maximum size, the oldest data is frozen.
        Defaults to 100 MB.
        """
        return pulumi.get(self, "max_global_raw_data_size_mb")

    @max_global_raw_data_size_mb.setter
    def max_global_raw_data_size_mb(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "max_global_raw_data_size_mb", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the index to create.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)


@pulumi.input_type
class _ShIndexesManagerState:
    def __init__(__self__, *,
                 acl: Optional[pulumi.Input['ShIndexesManagerAclArgs']] = None,
                 datatype: Optional[pulumi.Input[str]] = None,
                 frozen_time_period_in_secs: Optional[pulumi.Input[str]] = None,
                 max_global_raw_data_size_mb: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering ShIndexesManager resources.
        :param pulumi.Input[str] datatype: Valid values: (event | metric). Specifies the type of index.
        :param pulumi.Input[str] frozen_time_period_in_secs: Number of seconds after which indexed data rolls to frozen.
               Defaults to 94608000 (3 years).Freezing data means it is removed from the index. If you need to archive your data, refer to coldToFrozenDir and coldToFrozenScript parameter documentation.
        :param pulumi.Input[str] max_global_raw_data_size_mb: The maximum size of an index (in MB). If an index grows larger than the maximum size, the oldest data is frozen.
               Defaults to 100 MB.
        :param pulumi.Input[str] name: The name of the index to create.
        """
        if acl is not None:
            pulumi.set(__self__, "acl", acl)
        if datatype is not None:
            pulumi.set(__self__, "datatype", datatype)
        if frozen_time_period_in_secs is not None:
            pulumi.set(__self__, "frozen_time_period_in_secs", frozen_time_period_in_secs)
        if max_global_raw_data_size_mb is not None:
            pulumi.set(__self__, "max_global_raw_data_size_mb", max_global_raw_data_size_mb)
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def acl(self) -> Optional[pulumi.Input['ShIndexesManagerAclArgs']]:
        return pulumi.get(self, "acl")

    @acl.setter
    def acl(self, value: Optional[pulumi.Input['ShIndexesManagerAclArgs']]):
        pulumi.set(self, "acl", value)

    @property
    @pulumi.getter
    def datatype(self) -> Optional[pulumi.Input[str]]:
        """
        Valid values: (event | metric). Specifies the type of index.
        """
        return pulumi.get(self, "datatype")

    @datatype.setter
    def datatype(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "datatype", value)

    @property
    @pulumi.getter(name="frozenTimePeriodInSecs")
    def frozen_time_period_in_secs(self) -> Optional[pulumi.Input[str]]:
        """
        Number of seconds after which indexed data rolls to frozen.
        Defaults to 94608000 (3 years).Freezing data means it is removed from the index. If you need to archive your data, refer to coldToFrozenDir and coldToFrozenScript parameter documentation.
        """
        return pulumi.get(self, "frozen_time_period_in_secs")

    @frozen_time_period_in_secs.setter
    def frozen_time_period_in_secs(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "frozen_time_period_in_secs", value)

    @property
    @pulumi.getter(name="maxGlobalRawDataSizeMb")
    def max_global_raw_data_size_mb(self) -> Optional[pulumi.Input[str]]:
        """
        The maximum size of an index (in MB). If an index grows larger than the maximum size, the oldest data is frozen.
        Defaults to 100 MB.
        """
        return pulumi.get(self, "max_global_raw_data_size_mb")

    @max_global_raw_data_size_mb.setter
    def max_global_raw_data_size_mb(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "max_global_raw_data_size_mb", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the index to create.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)


class ShIndexesManager(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 acl: Optional[pulumi.Input[pulumi.InputType['ShIndexesManagerAclArgs']]] = None,
                 datatype: Optional[pulumi.Input[str]] = None,
                 frozen_time_period_in_secs: Optional[pulumi.Input[str]] = None,
                 max_global_raw_data_size_mb: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        ## # Resource: ShIndexesManager

        Create indexes on Splunk Cloud instances. [BETA]

        ## Authorization and authentication

        As of now there is no support to create indexes in user-specified workspaces on Splunk Cloud.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_splunk as splunk

        tf_index = splunk.ShIndexesManager("tf-index",
            datatype="event",
            frozen_time_period_in_secs="94608000",
            max_global_raw_data_size_mb="100")
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] datatype: Valid values: (event | metric). Specifies the type of index.
        :param pulumi.Input[str] frozen_time_period_in_secs: Number of seconds after which indexed data rolls to frozen.
               Defaults to 94608000 (3 years).Freezing data means it is removed from the index. If you need to archive your data, refer to coldToFrozenDir and coldToFrozenScript parameter documentation.
        :param pulumi.Input[str] max_global_raw_data_size_mb: The maximum size of an index (in MB). If an index grows larger than the maximum size, the oldest data is frozen.
               Defaults to 100 MB.
        :param pulumi.Input[str] name: The name of the index to create.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[ShIndexesManagerArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        ## # Resource: ShIndexesManager

        Create indexes on Splunk Cloud instances. [BETA]

        ## Authorization and authentication

        As of now there is no support to create indexes in user-specified workspaces on Splunk Cloud.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_splunk as splunk

        tf_index = splunk.ShIndexesManager("tf-index",
            datatype="event",
            frozen_time_period_in_secs="94608000",
            max_global_raw_data_size_mb="100")
        ```

        :param str resource_name: The name of the resource.
        :param ShIndexesManagerArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ShIndexesManagerArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 acl: Optional[pulumi.Input[pulumi.InputType['ShIndexesManagerAclArgs']]] = None,
                 datatype: Optional[pulumi.Input[str]] = None,
                 frozen_time_period_in_secs: Optional[pulumi.Input[str]] = None,
                 max_global_raw_data_size_mb: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
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
            __props__ = ShIndexesManagerArgs.__new__(ShIndexesManagerArgs)

            __props__.__dict__["acl"] = acl
            __props__.__dict__["datatype"] = datatype
            __props__.__dict__["frozen_time_period_in_secs"] = frozen_time_period_in_secs
            __props__.__dict__["max_global_raw_data_size_mb"] = max_global_raw_data_size_mb
            __props__.__dict__["name"] = name
        super(ShIndexesManager, __self__).__init__(
            'splunk:index/shIndexesManager:ShIndexesManager',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            acl: Optional[pulumi.Input[pulumi.InputType['ShIndexesManagerAclArgs']]] = None,
            datatype: Optional[pulumi.Input[str]] = None,
            frozen_time_period_in_secs: Optional[pulumi.Input[str]] = None,
            max_global_raw_data_size_mb: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None) -> 'ShIndexesManager':
        """
        Get an existing ShIndexesManager resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] datatype: Valid values: (event | metric). Specifies the type of index.
        :param pulumi.Input[str] frozen_time_period_in_secs: Number of seconds after which indexed data rolls to frozen.
               Defaults to 94608000 (3 years).Freezing data means it is removed from the index. If you need to archive your data, refer to coldToFrozenDir and coldToFrozenScript parameter documentation.
        :param pulumi.Input[str] max_global_raw_data_size_mb: The maximum size of an index (in MB). If an index grows larger than the maximum size, the oldest data is frozen.
               Defaults to 100 MB.
        :param pulumi.Input[str] name: The name of the index to create.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ShIndexesManagerState.__new__(_ShIndexesManagerState)

        __props__.__dict__["acl"] = acl
        __props__.__dict__["datatype"] = datatype
        __props__.__dict__["frozen_time_period_in_secs"] = frozen_time_period_in_secs
        __props__.__dict__["max_global_raw_data_size_mb"] = max_global_raw_data_size_mb
        __props__.__dict__["name"] = name
        return ShIndexesManager(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def acl(self) -> pulumi.Output['outputs.ShIndexesManagerAcl']:
        return pulumi.get(self, "acl")

    @property
    @pulumi.getter
    def datatype(self) -> pulumi.Output[str]:
        """
        Valid values: (event | metric). Specifies the type of index.
        """
        return pulumi.get(self, "datatype")

    @property
    @pulumi.getter(name="frozenTimePeriodInSecs")
    def frozen_time_period_in_secs(self) -> pulumi.Output[Optional[str]]:
        """
        Number of seconds after which indexed data rolls to frozen.
        Defaults to 94608000 (3 years).Freezing data means it is removed from the index. If you need to archive your data, refer to coldToFrozenDir and coldToFrozenScript parameter documentation.
        """
        return pulumi.get(self, "frozen_time_period_in_secs")

    @property
    @pulumi.getter(name="maxGlobalRawDataSizeMb")
    def max_global_raw_data_size_mb(self) -> pulumi.Output[Optional[str]]:
        """
        The maximum size of an index (in MB). If an index grows larger than the maximum size, the oldest data is frozen.
        Defaults to 100 MB.
        """
        return pulumi.get(self, "max_global_raw_data_size_mb")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the index to create.
        """
        return pulumi.get(self, "name")

