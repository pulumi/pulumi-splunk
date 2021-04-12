# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities, _tables
from . import outputs
from ._inputs import *

__all__ = ['DataUiViewsArgs', 'DataUiViews']

@pulumi.input_type
class DataUiViewsArgs:
    def __init__(__self__, *,
                 eai_data: pulumi.Input[str],
                 acl: Optional[pulumi.Input['DataUiViewsAclArgs']] = None,
                 name: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a DataUiViews resource.
        :param pulumi.Input[str] eai_data: Dashboard XML definition.
        :param pulumi.Input[str] name: Dashboard name.
               * `eai:data` - (Required) Dashboard XML definition.
        """
        pulumi.set(__self__, "eai_data", eai_data)
        if acl is not None:
            pulumi.set(__self__, "acl", acl)
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter(name="eaiData")
    def eai_data(self) -> pulumi.Input[str]:
        """
        Dashboard XML definition.
        """
        return pulumi.get(self, "eai_data")

    @eai_data.setter
    def eai_data(self, value: pulumi.Input[str]):
        pulumi.set(self, "eai_data", value)

    @property
    @pulumi.getter
    def acl(self) -> Optional[pulumi.Input['DataUiViewsAclArgs']]:
        return pulumi.get(self, "acl")

    @acl.setter
    def acl(self, value: Optional[pulumi.Input['DataUiViewsAclArgs']]):
        pulumi.set(self, "acl", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Dashboard name.
        * `eai:data` - (Required) Dashboard XML definition.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)


class DataUiViews(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 acl: Optional[pulumi.Input[pulumi.InputType['DataUiViewsAclArgs']]] = None,
                 eai_data: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        ## # Resource: DataUiViews

        Create and manage splunk dashboards/views.
        ## Example Usage

        ```python
        import pulumi
        import pulumi_splunk as splunk

        dashboard = splunk.DataUiViews("dashboard",
            acl=splunk.DataUiViewsAclArgs(
                app="search",
                owner="admin",
            ),
            eai_data=\"\"\"  <dashboard>
            <label> 
              Terraform Test Dashboard
            </label>
          </dashboard>
          
        \"\"\")
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] eai_data: Dashboard XML definition.
        :param pulumi.Input[str] name: Dashboard name.
               * `eai:data` - (Required) Dashboard XML definition.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: DataUiViewsArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        ## # Resource: DataUiViews

        Create and manage splunk dashboards/views.
        ## Example Usage

        ```python
        import pulumi
        import pulumi_splunk as splunk

        dashboard = splunk.DataUiViews("dashboard",
            acl=splunk.DataUiViewsAclArgs(
                app="search",
                owner="admin",
            ),
            eai_data=\"\"\"  <dashboard>
            <label> 
              Terraform Test Dashboard
            </label>
          </dashboard>
          
        \"\"\")
        ```

        :param str resource_name: The name of the resource.
        :param DataUiViewsArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(DataUiViewsArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 acl: Optional[pulumi.Input[pulumi.InputType['DataUiViewsAclArgs']]] = None,
                 eai_data: Optional[pulumi.Input[str]] = None,
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
            __props__ = dict()

            __props__['acl'] = acl
            if eai_data is None and not opts.urn:
                raise TypeError("Missing required property 'eai_data'")
            __props__['eai_data'] = eai_data
            __props__['name'] = name
        super(DataUiViews, __self__).__init__(
            'splunk:index/dataUiViews:DataUiViews',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            acl: Optional[pulumi.Input[pulumi.InputType['DataUiViewsAclArgs']]] = None,
            eai_data: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None) -> 'DataUiViews':
        """
        Get an existing DataUiViews resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] eai_data: Dashboard XML definition.
        :param pulumi.Input[str] name: Dashboard name.
               * `eai:data` - (Required) Dashboard XML definition.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["acl"] = acl
        __props__["eai_data"] = eai_data
        __props__["name"] = name
        return DataUiViews(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def acl(self) -> pulumi.Output['outputs.DataUiViewsAcl']:
        return pulumi.get(self, "acl")

    @property
    @pulumi.getter(name="eaiData")
    def eai_data(self) -> pulumi.Output[str]:
        """
        Dashboard XML definition.
        """
        return pulumi.get(self, "eai_data")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Dashboard name.
        * `eai:data` - (Required) Dashboard XML definition.
        """
        return pulumi.get(self, "name")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

