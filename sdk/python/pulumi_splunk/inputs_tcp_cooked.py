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

__all__ = ['InputsTcpCookedArgs', 'InputsTcpCooked']

@pulumi.input_type
class InputsTcpCookedArgs:
    def __init__(__self__, *,
                 acl: Optional[pulumi.Input['InputsTcpCookedAclArgs']] = None,
                 connection_host: Optional[pulumi.Input[str]] = None,
                 disabled: Optional[pulumi.Input[bool]] = None,
                 host: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 restrict_to_host: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a InputsTcpCooked resource.
        :param pulumi.Input['InputsTcpCookedAclArgs'] acl: The app/user context that is the namespace for the resource
        :param pulumi.Input[str] connection_host: Valid values: (ip | dns | none)
               Set the host for the remote server that is sending data.
               ip sets the host to the IP address of the remote server sending data.
               dns sets the host to the reverse DNS entry for the IP address of the remote server sending data.
               none leaves the host as specified in inputs.conf, which is typically the Splunk system hostname.
               Default value is dns.
        :param pulumi.Input[bool] disabled: Indicates if input is disabled.
        :param pulumi.Input[str] host: Host from which the indexer gets data.
        :param pulumi.Input[str] name: The port number of this input.
        :param pulumi.Input[str] restrict_to_host: Restrict incoming connections on this port to the host specified here.
        """
        if acl is not None:
            pulumi.set(__self__, "acl", acl)
        if connection_host is not None:
            pulumi.set(__self__, "connection_host", connection_host)
        if disabled is not None:
            pulumi.set(__self__, "disabled", disabled)
        if host is not None:
            pulumi.set(__self__, "host", host)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if restrict_to_host is not None:
            pulumi.set(__self__, "restrict_to_host", restrict_to_host)

    @property
    @pulumi.getter
    def acl(self) -> Optional[pulumi.Input['InputsTcpCookedAclArgs']]:
        """
        The app/user context that is the namespace for the resource
        """
        return pulumi.get(self, "acl")

    @acl.setter
    def acl(self, value: Optional[pulumi.Input['InputsTcpCookedAclArgs']]):
        pulumi.set(self, "acl", value)

    @property
    @pulumi.getter(name="connectionHost")
    def connection_host(self) -> Optional[pulumi.Input[str]]:
        """
        Valid values: (ip | dns | none)
        Set the host for the remote server that is sending data.
        ip sets the host to the IP address of the remote server sending data.
        dns sets the host to the reverse DNS entry for the IP address of the remote server sending data.
        none leaves the host as specified in inputs.conf, which is typically the Splunk system hostname.
        Default value is dns.
        """
        return pulumi.get(self, "connection_host")

    @connection_host.setter
    def connection_host(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "connection_host", value)

    @property
    @pulumi.getter
    def disabled(self) -> Optional[pulumi.Input[bool]]:
        """
        Indicates if input is disabled.
        """
        return pulumi.get(self, "disabled")

    @disabled.setter
    def disabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "disabled", value)

    @property
    @pulumi.getter
    def host(self) -> Optional[pulumi.Input[str]]:
        """
        Host from which the indexer gets data.
        """
        return pulumi.get(self, "host")

    @host.setter
    def host(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "host", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The port number of this input.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="restrictToHost")
    def restrict_to_host(self) -> Optional[pulumi.Input[str]]:
        """
        Restrict incoming connections on this port to the host specified here.
        """
        return pulumi.get(self, "restrict_to_host")

    @restrict_to_host.setter
    def restrict_to_host(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "restrict_to_host", value)


@pulumi.input_type
class _InputsTcpCookedState:
    def __init__(__self__, *,
                 acl: Optional[pulumi.Input['InputsTcpCookedAclArgs']] = None,
                 connection_host: Optional[pulumi.Input[str]] = None,
                 disabled: Optional[pulumi.Input[bool]] = None,
                 host: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 restrict_to_host: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering InputsTcpCooked resources.
        :param pulumi.Input['InputsTcpCookedAclArgs'] acl: The app/user context that is the namespace for the resource
        :param pulumi.Input[str] connection_host: Valid values: (ip | dns | none)
               Set the host for the remote server that is sending data.
               ip sets the host to the IP address of the remote server sending data.
               dns sets the host to the reverse DNS entry for the IP address of the remote server sending data.
               none leaves the host as specified in inputs.conf, which is typically the Splunk system hostname.
               Default value is dns.
        :param pulumi.Input[bool] disabled: Indicates if input is disabled.
        :param pulumi.Input[str] host: Host from which the indexer gets data.
        :param pulumi.Input[str] name: The port number of this input.
        :param pulumi.Input[str] restrict_to_host: Restrict incoming connections on this port to the host specified here.
        """
        if acl is not None:
            pulumi.set(__self__, "acl", acl)
        if connection_host is not None:
            pulumi.set(__self__, "connection_host", connection_host)
        if disabled is not None:
            pulumi.set(__self__, "disabled", disabled)
        if host is not None:
            pulumi.set(__self__, "host", host)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if restrict_to_host is not None:
            pulumi.set(__self__, "restrict_to_host", restrict_to_host)

    @property
    @pulumi.getter
    def acl(self) -> Optional[pulumi.Input['InputsTcpCookedAclArgs']]:
        """
        The app/user context that is the namespace for the resource
        """
        return pulumi.get(self, "acl")

    @acl.setter
    def acl(self, value: Optional[pulumi.Input['InputsTcpCookedAclArgs']]):
        pulumi.set(self, "acl", value)

    @property
    @pulumi.getter(name="connectionHost")
    def connection_host(self) -> Optional[pulumi.Input[str]]:
        """
        Valid values: (ip | dns | none)
        Set the host for the remote server that is sending data.
        ip sets the host to the IP address of the remote server sending data.
        dns sets the host to the reverse DNS entry for the IP address of the remote server sending data.
        none leaves the host as specified in inputs.conf, which is typically the Splunk system hostname.
        Default value is dns.
        """
        return pulumi.get(self, "connection_host")

    @connection_host.setter
    def connection_host(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "connection_host", value)

    @property
    @pulumi.getter
    def disabled(self) -> Optional[pulumi.Input[bool]]:
        """
        Indicates if input is disabled.
        """
        return pulumi.get(self, "disabled")

    @disabled.setter
    def disabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "disabled", value)

    @property
    @pulumi.getter
    def host(self) -> Optional[pulumi.Input[str]]:
        """
        Host from which the indexer gets data.
        """
        return pulumi.get(self, "host")

    @host.setter
    def host(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "host", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The port number of this input.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="restrictToHost")
    def restrict_to_host(self) -> Optional[pulumi.Input[str]]:
        """
        Restrict incoming connections on this port to the host specified here.
        """
        return pulumi.get(self, "restrict_to_host")

    @restrict_to_host.setter
    def restrict_to_host(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "restrict_to_host", value)


class InputsTcpCooked(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 acl: Optional[pulumi.Input[pulumi.InputType['InputsTcpCookedAclArgs']]] = None,
                 connection_host: Optional[pulumi.Input[str]] = None,
                 disabled: Optional[pulumi.Input[bool]] = None,
                 host: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 restrict_to_host: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        ## # Resource: InputsTcpCooked

        Create or update cooked TCP input information and create new containers for managing cooked data.

        ## Example Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import pulumi_splunk as splunk

        tcp_cooked = splunk.InputsTcpCooked("tcp_cooked",
            name="50000",
            disabled=False,
            connection_host="dns",
            restrict_to_host="splunk")
        ```
        <!--End PulumiCodeChooser -->

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['InputsTcpCookedAclArgs']] acl: The app/user context that is the namespace for the resource
        :param pulumi.Input[str] connection_host: Valid values: (ip | dns | none)
               Set the host for the remote server that is sending data.
               ip sets the host to the IP address of the remote server sending data.
               dns sets the host to the reverse DNS entry for the IP address of the remote server sending data.
               none leaves the host as specified in inputs.conf, which is typically the Splunk system hostname.
               Default value is dns.
        :param pulumi.Input[bool] disabled: Indicates if input is disabled.
        :param pulumi.Input[str] host: Host from which the indexer gets data.
        :param pulumi.Input[str] name: The port number of this input.
        :param pulumi.Input[str] restrict_to_host: Restrict incoming connections on this port to the host specified here.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[InputsTcpCookedArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        ## # Resource: InputsTcpCooked

        Create or update cooked TCP input information and create new containers for managing cooked data.

        ## Example Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import pulumi_splunk as splunk

        tcp_cooked = splunk.InputsTcpCooked("tcp_cooked",
            name="50000",
            disabled=False,
            connection_host="dns",
            restrict_to_host="splunk")
        ```
        <!--End PulumiCodeChooser -->

        :param str resource_name: The name of the resource.
        :param InputsTcpCookedArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(InputsTcpCookedArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 acl: Optional[pulumi.Input[pulumi.InputType['InputsTcpCookedAclArgs']]] = None,
                 connection_host: Optional[pulumi.Input[str]] = None,
                 disabled: Optional[pulumi.Input[bool]] = None,
                 host: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 restrict_to_host: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = InputsTcpCookedArgs.__new__(InputsTcpCookedArgs)

            __props__.__dict__["acl"] = acl
            __props__.__dict__["connection_host"] = connection_host
            __props__.__dict__["disabled"] = disabled
            __props__.__dict__["host"] = host
            __props__.__dict__["name"] = name
            __props__.__dict__["restrict_to_host"] = restrict_to_host
        super(InputsTcpCooked, __self__).__init__(
            'splunk:index/inputsTcpCooked:InputsTcpCooked',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            acl: Optional[pulumi.Input[pulumi.InputType['InputsTcpCookedAclArgs']]] = None,
            connection_host: Optional[pulumi.Input[str]] = None,
            disabled: Optional[pulumi.Input[bool]] = None,
            host: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            restrict_to_host: Optional[pulumi.Input[str]] = None) -> 'InputsTcpCooked':
        """
        Get an existing InputsTcpCooked resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['InputsTcpCookedAclArgs']] acl: The app/user context that is the namespace for the resource
        :param pulumi.Input[str] connection_host: Valid values: (ip | dns | none)
               Set the host for the remote server that is sending data.
               ip sets the host to the IP address of the remote server sending data.
               dns sets the host to the reverse DNS entry for the IP address of the remote server sending data.
               none leaves the host as specified in inputs.conf, which is typically the Splunk system hostname.
               Default value is dns.
        :param pulumi.Input[bool] disabled: Indicates if input is disabled.
        :param pulumi.Input[str] host: Host from which the indexer gets data.
        :param pulumi.Input[str] name: The port number of this input.
        :param pulumi.Input[str] restrict_to_host: Restrict incoming connections on this port to the host specified here.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _InputsTcpCookedState.__new__(_InputsTcpCookedState)

        __props__.__dict__["acl"] = acl
        __props__.__dict__["connection_host"] = connection_host
        __props__.__dict__["disabled"] = disabled
        __props__.__dict__["host"] = host
        __props__.__dict__["name"] = name
        __props__.__dict__["restrict_to_host"] = restrict_to_host
        return InputsTcpCooked(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def acl(self) -> pulumi.Output['outputs.InputsTcpCookedAcl']:
        """
        The app/user context that is the namespace for the resource
        """
        return pulumi.get(self, "acl")

    @property
    @pulumi.getter(name="connectionHost")
    def connection_host(self) -> pulumi.Output[str]:
        """
        Valid values: (ip | dns | none)
        Set the host for the remote server that is sending data.
        ip sets the host to the IP address of the remote server sending data.
        dns sets the host to the reverse DNS entry for the IP address of the remote server sending data.
        none leaves the host as specified in inputs.conf, which is typically the Splunk system hostname.
        Default value is dns.
        """
        return pulumi.get(self, "connection_host")

    @property
    @pulumi.getter
    def disabled(self) -> pulumi.Output[bool]:
        """
        Indicates if input is disabled.
        """
        return pulumi.get(self, "disabled")

    @property
    @pulumi.getter
    def host(self) -> pulumi.Output[str]:
        """
        Host from which the indexer gets data.
        """
        return pulumi.get(self, "host")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The port number of this input.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="restrictToHost")
    def restrict_to_host(self) -> pulumi.Output[str]:
        """
        Restrict incoming connections on this port to the host specified here.
        """
        return pulumi.get(self, "restrict_to_host")

