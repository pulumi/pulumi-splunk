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

__all__ = ['GenericAclArgs', 'GenericAcl']

@pulumi.input_type
class GenericAclArgs:
    def __init__(__self__, *,
                 path: pulumi.Input[str],
                 acl: Optional[pulumi.Input['GenericAclAclArgs']] = None):
        """
        The set of arguments for constructing a GenericAcl resource.
        :param pulumi.Input[str] path: REST API Endpoint path to the object, relative to servicesNS/<owner>/<app>
        :param pulumi.Input['GenericAclAclArgs'] acl: The ACL to apply to the object, including app/owner to properly identify the object.
               Though technically optional, it should be explicitly set for this resource to really be valid. Some objects, such as
               apps, require specific values for app and owner. Consult the REST API documentation regarding which values to use for
               app and owner for objects that don't fit in the normal namespace.
        """
        pulumi.set(__self__, "path", path)
        if acl is not None:
            pulumi.set(__self__, "acl", acl)

    @property
    @pulumi.getter
    def path(self) -> pulumi.Input[str]:
        """
        REST API Endpoint path to the object, relative to servicesNS/<owner>/<app>
        """
        return pulumi.get(self, "path")

    @path.setter
    def path(self, value: pulumi.Input[str]):
        pulumi.set(self, "path", value)

    @property
    @pulumi.getter
    def acl(self) -> Optional[pulumi.Input['GenericAclAclArgs']]:
        """
        The ACL to apply to the object, including app/owner to properly identify the object.
        Though technically optional, it should be explicitly set for this resource to really be valid. Some objects, such as
        apps, require specific values for app and owner. Consult the REST API documentation regarding which values to use for
        app and owner for objects that don't fit in the normal namespace.
        """
        return pulumi.get(self, "acl")

    @acl.setter
    def acl(self, value: Optional[pulumi.Input['GenericAclAclArgs']]):
        pulumi.set(self, "acl", value)


@pulumi.input_type
class _GenericAclState:
    def __init__(__self__, *,
                 acl: Optional[pulumi.Input['GenericAclAclArgs']] = None,
                 path: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering GenericAcl resources.
        :param pulumi.Input['GenericAclAclArgs'] acl: The ACL to apply to the object, including app/owner to properly identify the object.
               Though technically optional, it should be explicitly set for this resource to really be valid. Some objects, such as
               apps, require specific values for app and owner. Consult the REST API documentation regarding which values to use for
               app and owner for objects that don't fit in the normal namespace.
        :param pulumi.Input[str] path: REST API Endpoint path to the object, relative to servicesNS/<owner>/<app>
        """
        if acl is not None:
            pulumi.set(__self__, "acl", acl)
        if path is not None:
            pulumi.set(__self__, "path", path)

    @property
    @pulumi.getter
    def acl(self) -> Optional[pulumi.Input['GenericAclAclArgs']]:
        """
        The ACL to apply to the object, including app/owner to properly identify the object.
        Though technically optional, it should be explicitly set for this resource to really be valid. Some objects, such as
        apps, require specific values for app and owner. Consult the REST API documentation regarding which values to use for
        app and owner for objects that don't fit in the normal namespace.
        """
        return pulumi.get(self, "acl")

    @acl.setter
    def acl(self, value: Optional[pulumi.Input['GenericAclAclArgs']]):
        pulumi.set(self, "acl", value)

    @property
    @pulumi.getter
    def path(self) -> Optional[pulumi.Input[str]]:
        """
        REST API Endpoint path to the object, relative to servicesNS/<owner>/<app>
        """
        return pulumi.get(self, "path")

    @path.setter
    def path(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "path", value)


class GenericAcl(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 acl: Optional[pulumi.Input[pulumi.InputType['GenericAclAclArgs']]] = None,
                 path: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        ## Example Usage

        ```python
        import pulumi
        import pulumi_splunk as splunk

        my_app = splunk.GenericAcl("myApp",
            acl=splunk.GenericAclAclArgs(
                app="system",
                owner="nobody",
                reads=["*"],
                writes=[
                    "admin",
                    "power",
                ],
            ),
            path="apps/local/my_app")
        my_dashboard = splunk.GenericAcl("myDashboard",
            acl=splunk.GenericAclAclArgs(
                app="my_app",
                owner="joe_user",
                reads=["team_joe"],
                writes=["team_joe"],
            ),
            path="data/ui/views/my_dashboard")
        ```

        ## Import

        Generic ACL resources can be imported by specifying their owner, app, and path with a colon-delimited string as the ID

        ```sh
         $ pulumi import splunk:index/genericAcl:GenericAcl splunk_generic_acl <owner>:<app>:<path>
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['GenericAclAclArgs']] acl: The ACL to apply to the object, including app/owner to properly identify the object.
               Though technically optional, it should be explicitly set for this resource to really be valid. Some objects, such as
               apps, require specific values for app and owner. Consult the REST API documentation regarding which values to use for
               app and owner for objects that don't fit in the normal namespace.
        :param pulumi.Input[str] path: REST API Endpoint path to the object, relative to servicesNS/<owner>/<app>
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: GenericAclArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        ## Example Usage

        ```python
        import pulumi
        import pulumi_splunk as splunk

        my_app = splunk.GenericAcl("myApp",
            acl=splunk.GenericAclAclArgs(
                app="system",
                owner="nobody",
                reads=["*"],
                writes=[
                    "admin",
                    "power",
                ],
            ),
            path="apps/local/my_app")
        my_dashboard = splunk.GenericAcl("myDashboard",
            acl=splunk.GenericAclAclArgs(
                app="my_app",
                owner="joe_user",
                reads=["team_joe"],
                writes=["team_joe"],
            ),
            path="data/ui/views/my_dashboard")
        ```

        ## Import

        Generic ACL resources can be imported by specifying their owner, app, and path with a colon-delimited string as the ID

        ```sh
         $ pulumi import splunk:index/genericAcl:GenericAcl splunk_generic_acl <owner>:<app>:<path>
        ```

        :param str resource_name: The name of the resource.
        :param GenericAclArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(GenericAclArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 acl: Optional[pulumi.Input[pulumi.InputType['GenericAclAclArgs']]] = None,
                 path: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = GenericAclArgs.__new__(GenericAclArgs)

            __props__.__dict__["acl"] = acl
            if path is None and not opts.urn:
                raise TypeError("Missing required property 'path'")
            __props__.__dict__["path"] = path
        super(GenericAcl, __self__).__init__(
            'splunk:index/genericAcl:GenericAcl',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            acl: Optional[pulumi.Input[pulumi.InputType['GenericAclAclArgs']]] = None,
            path: Optional[pulumi.Input[str]] = None) -> 'GenericAcl':
        """
        Get an existing GenericAcl resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['GenericAclAclArgs']] acl: The ACL to apply to the object, including app/owner to properly identify the object.
               Though technically optional, it should be explicitly set for this resource to really be valid. Some objects, such as
               apps, require specific values for app and owner. Consult the REST API documentation regarding which values to use for
               app and owner for objects that don't fit in the normal namespace.
        :param pulumi.Input[str] path: REST API Endpoint path to the object, relative to servicesNS/<owner>/<app>
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _GenericAclState.__new__(_GenericAclState)

        __props__.__dict__["acl"] = acl
        __props__.__dict__["path"] = path
        return GenericAcl(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def acl(self) -> pulumi.Output['outputs.GenericAclAcl']:
        """
        The ACL to apply to the object, including app/owner to properly identify the object.
        Though technically optional, it should be explicitly set for this resource to really be valid. Some objects, such as
        apps, require specific values for app and owner. Consult the REST API documentation regarding which values to use for
        app and owner for objects that don't fit in the normal namespace.
        """
        return pulumi.get(self, "acl")

    @property
    @pulumi.getter
    def path(self) -> pulumi.Output[str]:
        """
        REST API Endpoint path to the object, relative to servicesNS/<owner>/<app>
        """
        return pulumi.get(self, "path")

