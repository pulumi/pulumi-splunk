# coding=utf-8
# *** WARNING: this file was generated by pulumi-language-python. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import builtins as _builtins
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
from . import outputs
from ._inputs import *

__all__ = ['InputsTcpSplunkTcpTokenArgs', 'InputsTcpSplunkTcpToken']

@pulumi.input_type
class InputsTcpSplunkTcpTokenArgs:
    def __init__(__self__, *,
                 acl: Optional[pulumi.Input['InputsTcpSplunkTcpTokenAclArgs']] = None,
                 name: Optional[pulumi.Input[_builtins.str]] = None,
                 token: Optional[pulumi.Input[_builtins.str]] = None):
        """
        The set of arguments for constructing a InputsTcpSplunkTcpToken resource.
        :param pulumi.Input['InputsTcpSplunkTcpTokenAclArgs'] acl: The app/user context that is the namespace for the resource
        :param pulumi.Input[_builtins.str] name: Required. Name for the token to create.
        :param pulumi.Input[_builtins.str] token: Optional. Token value to use. If unspecified, a token is generated automatically.
        """
        if acl is not None:
            pulumi.set(__self__, "acl", acl)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if token is not None:
            pulumi.set(__self__, "token", token)

    @_builtins.property
    @pulumi.getter
    def acl(self) -> Optional[pulumi.Input['InputsTcpSplunkTcpTokenAclArgs']]:
        """
        The app/user context that is the namespace for the resource
        """
        return pulumi.get(self, "acl")

    @acl.setter
    def acl(self, value: Optional[pulumi.Input['InputsTcpSplunkTcpTokenAclArgs']]):
        pulumi.set(self, "acl", value)

    @_builtins.property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[_builtins.str]]:
        """
        Required. Name for the token to create.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "name", value)

    @_builtins.property
    @pulumi.getter
    def token(self) -> Optional[pulumi.Input[_builtins.str]]:
        """
        Optional. Token value to use. If unspecified, a token is generated automatically.
        """
        return pulumi.get(self, "token")

    @token.setter
    def token(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "token", value)


@pulumi.input_type
class _InputsTcpSplunkTcpTokenState:
    def __init__(__self__, *,
                 acl: Optional[pulumi.Input['InputsTcpSplunkTcpTokenAclArgs']] = None,
                 name: Optional[pulumi.Input[_builtins.str]] = None,
                 token: Optional[pulumi.Input[_builtins.str]] = None):
        """
        Input properties used for looking up and filtering InputsTcpSplunkTcpToken resources.
        :param pulumi.Input['InputsTcpSplunkTcpTokenAclArgs'] acl: The app/user context that is the namespace for the resource
        :param pulumi.Input[_builtins.str] name: Required. Name for the token to create.
        :param pulumi.Input[_builtins.str] token: Optional. Token value to use. If unspecified, a token is generated automatically.
        """
        if acl is not None:
            pulumi.set(__self__, "acl", acl)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if token is not None:
            pulumi.set(__self__, "token", token)

    @_builtins.property
    @pulumi.getter
    def acl(self) -> Optional[pulumi.Input['InputsTcpSplunkTcpTokenAclArgs']]:
        """
        The app/user context that is the namespace for the resource
        """
        return pulumi.get(self, "acl")

    @acl.setter
    def acl(self, value: Optional[pulumi.Input['InputsTcpSplunkTcpTokenAclArgs']]):
        pulumi.set(self, "acl", value)

    @_builtins.property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[_builtins.str]]:
        """
        Required. Name for the token to create.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "name", value)

    @_builtins.property
    @pulumi.getter
    def token(self) -> Optional[pulumi.Input[_builtins.str]]:
        """
        Optional. Token value to use. If unspecified, a token is generated automatically.
        """
        return pulumi.get(self, "token")

    @token.setter
    def token(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "token", value)


@pulumi.type_token("splunk:index/inputsTcpSplunkTcpToken:InputsTcpSplunkTcpToken")
class InputsTcpSplunkTcpToken(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 acl: Optional[pulumi.Input[Union['InputsTcpSplunkTcpTokenAclArgs', 'InputsTcpSplunkTcpTokenAclArgsDict']]] = None,
                 name: Optional[pulumi.Input[_builtins.str]] = None,
                 token: Optional[pulumi.Input[_builtins.str]] = None,
                 __props__=None):
        """
        ## # Resource: InputsTcpSplunkTcpToken

        Manage receiver access using tokens.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_splunk as splunk

        tcp_splunk_tcp_token = splunk.InputsTcpSplunkTcpToken("tcp_splunk_tcp_token",
            name="new-splunk-tcp-token",
            token="D66C45B3-7C28-48A1-A13A-027914146501")
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Union['InputsTcpSplunkTcpTokenAclArgs', 'InputsTcpSplunkTcpTokenAclArgsDict']] acl: The app/user context that is the namespace for the resource
        :param pulumi.Input[_builtins.str] name: Required. Name for the token to create.
        :param pulumi.Input[_builtins.str] token: Optional. Token value to use. If unspecified, a token is generated automatically.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[InputsTcpSplunkTcpTokenArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        ## # Resource: InputsTcpSplunkTcpToken

        Manage receiver access using tokens.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_splunk as splunk

        tcp_splunk_tcp_token = splunk.InputsTcpSplunkTcpToken("tcp_splunk_tcp_token",
            name="new-splunk-tcp-token",
            token="D66C45B3-7C28-48A1-A13A-027914146501")
        ```

        :param str resource_name: The name of the resource.
        :param InputsTcpSplunkTcpTokenArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(InputsTcpSplunkTcpTokenArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 acl: Optional[pulumi.Input[Union['InputsTcpSplunkTcpTokenAclArgs', 'InputsTcpSplunkTcpTokenAclArgsDict']]] = None,
                 name: Optional[pulumi.Input[_builtins.str]] = None,
                 token: Optional[pulumi.Input[_builtins.str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = InputsTcpSplunkTcpTokenArgs.__new__(InputsTcpSplunkTcpTokenArgs)

            __props__.__dict__["acl"] = acl
            __props__.__dict__["name"] = name
            __props__.__dict__["token"] = token
        super(InputsTcpSplunkTcpToken, __self__).__init__(
            'splunk:index/inputsTcpSplunkTcpToken:InputsTcpSplunkTcpToken',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            acl: Optional[pulumi.Input[Union['InputsTcpSplunkTcpTokenAclArgs', 'InputsTcpSplunkTcpTokenAclArgsDict']]] = None,
            name: Optional[pulumi.Input[_builtins.str]] = None,
            token: Optional[pulumi.Input[_builtins.str]] = None) -> 'InputsTcpSplunkTcpToken':
        """
        Get an existing InputsTcpSplunkTcpToken resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Union['InputsTcpSplunkTcpTokenAclArgs', 'InputsTcpSplunkTcpTokenAclArgsDict']] acl: The app/user context that is the namespace for the resource
        :param pulumi.Input[_builtins.str] name: Required. Name for the token to create.
        :param pulumi.Input[_builtins.str] token: Optional. Token value to use. If unspecified, a token is generated automatically.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _InputsTcpSplunkTcpTokenState.__new__(_InputsTcpSplunkTcpTokenState)

        __props__.__dict__["acl"] = acl
        __props__.__dict__["name"] = name
        __props__.__dict__["token"] = token
        return InputsTcpSplunkTcpToken(resource_name, opts=opts, __props__=__props__)

    @_builtins.property
    @pulumi.getter
    def acl(self) -> pulumi.Output['outputs.InputsTcpSplunkTcpTokenAcl']:
        """
        The app/user context that is the namespace for the resource
        """
        return pulumi.get(self, "acl")

    @_builtins.property
    @pulumi.getter
    def name(self) -> pulumi.Output[_builtins.str]:
        """
        Required. Name for the token to create.
        """
        return pulumi.get(self, "name")

    @_builtins.property
    @pulumi.getter
    def token(self) -> pulumi.Output[_builtins.str]:
        """
        Optional. Token value to use. If unspecified, a token is generated automatically.
        """
        return pulumi.get(self, "token")

