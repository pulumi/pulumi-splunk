# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import builtins
import copy
import warnings
import sys
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
if sys.version_info >= (3, 11):
    from typing import NotRequired, TypedDict, TypeAlias
else:
    from typing_extensions import NotRequired, TypedDict, TypeAlias
from . import _utilities

__all__ = ['AdminSamlGroupsArgs', 'AdminSamlGroups']

@pulumi.input_type
class AdminSamlGroupsArgs:
    def __init__(__self__, *,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 roles: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None):
        """
        The set of arguments for constructing a AdminSamlGroups resource.
        :param pulumi.Input[builtins.str] name: The name of the external group.
        :param pulumi.Input[Sequence[pulumi.Input[builtins.str]]] roles: List of internal roles assigned to the group.
        """
        if name is not None:
            pulumi.set(__self__, "name", name)
        if roles is not None:
            pulumi.set(__self__, "roles", roles)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The name of the external group.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def roles(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]]:
        """
        List of internal roles assigned to the group.
        """
        return pulumi.get(self, "roles")

    @roles.setter
    def roles(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]]):
        pulumi.set(self, "roles", value)


@pulumi.input_type
class _AdminSamlGroupsState:
    def __init__(__self__, *,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 roles: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None):
        """
        Input properties used for looking up and filtering AdminSamlGroups resources.
        :param pulumi.Input[builtins.str] name: The name of the external group.
        :param pulumi.Input[Sequence[pulumi.Input[builtins.str]]] roles: List of internal roles assigned to the group.
        """
        if name is not None:
            pulumi.set(__self__, "name", name)
        if roles is not None:
            pulumi.set(__self__, "roles", roles)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The name of the external group.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def roles(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]]:
        """
        List of internal roles assigned to the group.
        """
        return pulumi.get(self, "roles")

    @roles.setter
    def roles(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]]):
        pulumi.set(self, "roles", value)


class AdminSamlGroups(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 roles: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None,
                 __props__=None):
        """
        ## # Resource: AdminSamlGroups

        Manage external groups in an IdP response to internal Splunk roles.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_splunk as splunk

        saml_group = splunk.AdminSamlGroups("saml-group",
            name="mygroup",
            roles=[
                "admin",
                "power",
            ])
        ```

        ## Import

        SAML groups can be imported using the id, e.g.

        ```sh
        $ pulumi import splunk:index/adminSamlGroups:AdminSamlGroups saml-group mygroup
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[builtins.str] name: The name of the external group.
        :param pulumi.Input[Sequence[pulumi.Input[builtins.str]]] roles: List of internal roles assigned to the group.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[AdminSamlGroupsArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        ## # Resource: AdminSamlGroups

        Manage external groups in an IdP response to internal Splunk roles.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_splunk as splunk

        saml_group = splunk.AdminSamlGroups("saml-group",
            name="mygroup",
            roles=[
                "admin",
                "power",
            ])
        ```

        ## Import

        SAML groups can be imported using the id, e.g.

        ```sh
        $ pulumi import splunk:index/adminSamlGroups:AdminSamlGroups saml-group mygroup
        ```

        :param str resource_name: The name of the resource.
        :param AdminSamlGroupsArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(AdminSamlGroupsArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 roles: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = AdminSamlGroupsArgs.__new__(AdminSamlGroupsArgs)

            __props__.__dict__["name"] = name
            __props__.__dict__["roles"] = roles
        super(AdminSamlGroups, __self__).__init__(
            'splunk:index/adminSamlGroups:AdminSamlGroups',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            name: Optional[pulumi.Input[builtins.str]] = None,
            roles: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None) -> 'AdminSamlGroups':
        """
        Get an existing AdminSamlGroups resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[builtins.str] name: The name of the external group.
        :param pulumi.Input[Sequence[pulumi.Input[builtins.str]]] roles: List of internal roles assigned to the group.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _AdminSamlGroupsState.__new__(_AdminSamlGroupsState)

        __props__.__dict__["name"] = name
        __props__.__dict__["roles"] = roles
        return AdminSamlGroups(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[builtins.str]:
        """
        The name of the external group.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def roles(self) -> pulumi.Output[Sequence[builtins.str]]:
        """
        List of internal roles assigned to the group.
        """
        return pulumi.get(self, "roles")

