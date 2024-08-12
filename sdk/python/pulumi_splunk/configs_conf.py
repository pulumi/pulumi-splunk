# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities
from . import outputs
from ._inputs import *

__all__ = ['ConfigsConfArgs', 'ConfigsConf']

@pulumi.input_type
class ConfigsConfArgs:
    def __init__(__self__, *,
                 acl: Optional[pulumi.Input['ConfigsConfAclArgs']] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 variables: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None):
        """
        The set of arguments for constructing a ConfigsConf resource.
        :param pulumi.Input[str] name: A '/' separated string consisting of {conf_file_name}/{stanza_name} ex. props/custom_stanza
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] variables: A map of key value pairs for a stanza.
        """
        if acl is not None:
            pulumi.set(__self__, "acl", acl)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if variables is not None:
            pulumi.set(__self__, "variables", variables)

    @property
    @pulumi.getter
    def acl(self) -> Optional[pulumi.Input['ConfigsConfAclArgs']]:
        return pulumi.get(self, "acl")

    @acl.setter
    def acl(self, value: Optional[pulumi.Input['ConfigsConfAclArgs']]):
        pulumi.set(self, "acl", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        A '/' separated string consisting of {conf_file_name}/{stanza_name} ex. props/custom_stanza
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def variables(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        A map of key value pairs for a stanza.
        """
        return pulumi.get(self, "variables")

    @variables.setter
    def variables(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "variables", value)


@pulumi.input_type
class _ConfigsConfState:
    def __init__(__self__, *,
                 acl: Optional[pulumi.Input['ConfigsConfAclArgs']] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 variables: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None):
        """
        Input properties used for looking up and filtering ConfigsConf resources.
        :param pulumi.Input[str] name: A '/' separated string consisting of {conf_file_name}/{stanza_name} ex. props/custom_stanza
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] variables: A map of key value pairs for a stanza.
        """
        if acl is not None:
            pulumi.set(__self__, "acl", acl)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if variables is not None:
            pulumi.set(__self__, "variables", variables)

    @property
    @pulumi.getter
    def acl(self) -> Optional[pulumi.Input['ConfigsConfAclArgs']]:
        return pulumi.get(self, "acl")

    @acl.setter
    def acl(self, value: Optional[pulumi.Input['ConfigsConfAclArgs']]):
        pulumi.set(self, "acl", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        A '/' separated string consisting of {conf_file_name}/{stanza_name} ex. props/custom_stanza
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def variables(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        A map of key value pairs for a stanza.
        """
        return pulumi.get(self, "variables")

    @variables.setter
    def variables(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "variables", value)


class ConfigsConf(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 acl: Optional[pulumi.Input[Union['ConfigsConfAclArgs', 'ConfigsConfAclArgsDict']]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 variables: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 __props__=None):
        """
        ## # Resource: ConfigsConf

        Create and manage configuration file stanzas.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_splunk as splunk

        new_conf_stanza = splunk.ConfigsConf("new-conf-stanza",
            name="custom-conf/custom",
            variables={
                "disabled": "false",
                "custom_key": "value",
            })
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] name: A '/' separated string consisting of {conf_file_name}/{stanza_name} ex. props/custom_stanza
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] variables: A map of key value pairs for a stanza.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[ConfigsConfArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        ## # Resource: ConfigsConf

        Create and manage configuration file stanzas.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_splunk as splunk

        new_conf_stanza = splunk.ConfigsConf("new-conf-stanza",
            name="custom-conf/custom",
            variables={
                "disabled": "false",
                "custom_key": "value",
            })
        ```

        :param str resource_name: The name of the resource.
        :param ConfigsConfArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ConfigsConfArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 acl: Optional[pulumi.Input[Union['ConfigsConfAclArgs', 'ConfigsConfAclArgsDict']]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 variables: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ConfigsConfArgs.__new__(ConfigsConfArgs)

            __props__.__dict__["acl"] = acl
            __props__.__dict__["name"] = name
            __props__.__dict__["variables"] = variables
        super(ConfigsConf, __self__).__init__(
            'splunk:index/configsConf:ConfigsConf',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            acl: Optional[pulumi.Input[Union['ConfigsConfAclArgs', 'ConfigsConfAclArgsDict']]] = None,
            name: Optional[pulumi.Input[str]] = None,
            variables: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None) -> 'ConfigsConf':
        """
        Get an existing ConfigsConf resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] name: A '/' separated string consisting of {conf_file_name}/{stanza_name} ex. props/custom_stanza
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] variables: A map of key value pairs for a stanza.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ConfigsConfState.__new__(_ConfigsConfState)

        __props__.__dict__["acl"] = acl
        __props__.__dict__["name"] = name
        __props__.__dict__["variables"] = variables
        return ConfigsConf(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def acl(self) -> pulumi.Output['outputs.ConfigsConfAcl']:
        return pulumi.get(self, "acl")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        A '/' separated string consisting of {conf_file_name}/{stanza_name} ex. props/custom_stanza
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def variables(self) -> pulumi.Output[Mapping[str, str]]:
        """
        A map of key value pairs for a stanza.
        """
        return pulumi.get(self, "variables")

