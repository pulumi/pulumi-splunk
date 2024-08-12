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

__all__ = ['InputsUdpArgs', 'InputsUdp']

@pulumi.input_type
class InputsUdpArgs:
    def __init__(__self__, *,
                 acl: Optional[pulumi.Input['InputsUdpAclArgs']] = None,
                 connection_host: Optional[pulumi.Input[str]] = None,
                 disabled: Optional[pulumi.Input[bool]] = None,
                 host: Optional[pulumi.Input[str]] = None,
                 index: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 no_appending_timestamp: Optional[pulumi.Input[bool]] = None,
                 no_priority_stripping: Optional[pulumi.Input[bool]] = None,
                 queue: Optional[pulumi.Input[str]] = None,
                 restrict_to_host: Optional[pulumi.Input[str]] = None,
                 source: Optional[pulumi.Input[str]] = None,
                 sourcetype: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a InputsUdp resource.
        :param pulumi.Input['InputsUdpAclArgs'] acl: The app/user context that is the namespace for the resource
        :param pulumi.Input[str] connection_host: Valid values: (ip | dns | none)
               Set the host for the remote server that is sending data.
               ip sets the host to the IP address of the remote server sending data.
               dns sets the host to the reverse DNS entry for the IP address of the remote server sending data.
               none leaves the host as specified in inputs.conf, which is typically the Splunk system hostname.
               Default value is dns.
        :param pulumi.Input[bool] disabled: Indicates if input is disabled.
        :param pulumi.Input[str] host: The value to populate in the host field for incoming events. This is used during parsing/indexing, in particular to set the host field. It is also the host field used at search time.
        :param pulumi.Input[str] index: Which index events from this input should be stored in. Defaults to default.
        :param pulumi.Input[str] name: The UDP port that this input should listen on.
        :param pulumi.Input[bool] no_appending_timestamp: If set to true, prevents Splunk software from prepending a timestamp and hostname to incoming events.
        :param pulumi.Input[bool] no_priority_stripping: If set to true, Splunk software does not remove the priority field from incoming syslog events.
        :param pulumi.Input[str] queue: Which queue events from this input should be sent to. Generally this does not need to be changed.
        :param pulumi.Input[str] restrict_to_host: Restrict incoming connections on this port to the host specified here.
               If this is not set, the value specified in [udp://<remote server>:<port>] in inputs.conf is used.
        :param pulumi.Input[str] source: The value to populate in the source field for incoming events. The same source should not be used for multiple data inputs.
        :param pulumi.Input[str] sourcetype: The value to populate in the sourcetype field for incoming events.
        """
        if acl is not None:
            pulumi.set(__self__, "acl", acl)
        if connection_host is not None:
            pulumi.set(__self__, "connection_host", connection_host)
        if disabled is not None:
            pulumi.set(__self__, "disabled", disabled)
        if host is not None:
            pulumi.set(__self__, "host", host)
        if index is not None:
            pulumi.set(__self__, "index", index)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if no_appending_timestamp is not None:
            pulumi.set(__self__, "no_appending_timestamp", no_appending_timestamp)
        if no_priority_stripping is not None:
            pulumi.set(__self__, "no_priority_stripping", no_priority_stripping)
        if queue is not None:
            pulumi.set(__self__, "queue", queue)
        if restrict_to_host is not None:
            pulumi.set(__self__, "restrict_to_host", restrict_to_host)
        if source is not None:
            pulumi.set(__self__, "source", source)
        if sourcetype is not None:
            pulumi.set(__self__, "sourcetype", sourcetype)

    @property
    @pulumi.getter
    def acl(self) -> Optional[pulumi.Input['InputsUdpAclArgs']]:
        """
        The app/user context that is the namespace for the resource
        """
        return pulumi.get(self, "acl")

    @acl.setter
    def acl(self, value: Optional[pulumi.Input['InputsUdpAclArgs']]):
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
        The value to populate in the host field for incoming events. This is used during parsing/indexing, in particular to set the host field. It is also the host field used at search time.
        """
        return pulumi.get(self, "host")

    @host.setter
    def host(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "host", value)

    @property
    @pulumi.getter
    def index(self) -> Optional[pulumi.Input[str]]:
        """
        Which index events from this input should be stored in. Defaults to default.
        """
        return pulumi.get(self, "index")

    @index.setter
    def index(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "index", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The UDP port that this input should listen on.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="noAppendingTimestamp")
    def no_appending_timestamp(self) -> Optional[pulumi.Input[bool]]:
        """
        If set to true, prevents Splunk software from prepending a timestamp and hostname to incoming events.
        """
        return pulumi.get(self, "no_appending_timestamp")

    @no_appending_timestamp.setter
    def no_appending_timestamp(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "no_appending_timestamp", value)

    @property
    @pulumi.getter(name="noPriorityStripping")
    def no_priority_stripping(self) -> Optional[pulumi.Input[bool]]:
        """
        If set to true, Splunk software does not remove the priority field from incoming syslog events.
        """
        return pulumi.get(self, "no_priority_stripping")

    @no_priority_stripping.setter
    def no_priority_stripping(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "no_priority_stripping", value)

    @property
    @pulumi.getter
    def queue(self) -> Optional[pulumi.Input[str]]:
        """
        Which queue events from this input should be sent to. Generally this does not need to be changed.
        """
        return pulumi.get(self, "queue")

    @queue.setter
    def queue(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "queue", value)

    @property
    @pulumi.getter(name="restrictToHost")
    def restrict_to_host(self) -> Optional[pulumi.Input[str]]:
        """
        Restrict incoming connections on this port to the host specified here.
        If this is not set, the value specified in [udp://<remote server>:<port>] in inputs.conf is used.
        """
        return pulumi.get(self, "restrict_to_host")

    @restrict_to_host.setter
    def restrict_to_host(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "restrict_to_host", value)

    @property
    @pulumi.getter
    def source(self) -> Optional[pulumi.Input[str]]:
        """
        The value to populate in the source field for incoming events. The same source should not be used for multiple data inputs.
        """
        return pulumi.get(self, "source")

    @source.setter
    def source(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "source", value)

    @property
    @pulumi.getter
    def sourcetype(self) -> Optional[pulumi.Input[str]]:
        """
        The value to populate in the sourcetype field for incoming events.
        """
        return pulumi.get(self, "sourcetype")

    @sourcetype.setter
    def sourcetype(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "sourcetype", value)


@pulumi.input_type
class _InputsUdpState:
    def __init__(__self__, *,
                 acl: Optional[pulumi.Input['InputsUdpAclArgs']] = None,
                 connection_host: Optional[pulumi.Input[str]] = None,
                 disabled: Optional[pulumi.Input[bool]] = None,
                 host: Optional[pulumi.Input[str]] = None,
                 index: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 no_appending_timestamp: Optional[pulumi.Input[bool]] = None,
                 no_priority_stripping: Optional[pulumi.Input[bool]] = None,
                 queue: Optional[pulumi.Input[str]] = None,
                 restrict_to_host: Optional[pulumi.Input[str]] = None,
                 source: Optional[pulumi.Input[str]] = None,
                 sourcetype: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering InputsUdp resources.
        :param pulumi.Input['InputsUdpAclArgs'] acl: The app/user context that is the namespace for the resource
        :param pulumi.Input[str] connection_host: Valid values: (ip | dns | none)
               Set the host for the remote server that is sending data.
               ip sets the host to the IP address of the remote server sending data.
               dns sets the host to the reverse DNS entry for the IP address of the remote server sending data.
               none leaves the host as specified in inputs.conf, which is typically the Splunk system hostname.
               Default value is dns.
        :param pulumi.Input[bool] disabled: Indicates if input is disabled.
        :param pulumi.Input[str] host: The value to populate in the host field for incoming events. This is used during parsing/indexing, in particular to set the host field. It is also the host field used at search time.
        :param pulumi.Input[str] index: Which index events from this input should be stored in. Defaults to default.
        :param pulumi.Input[str] name: The UDP port that this input should listen on.
        :param pulumi.Input[bool] no_appending_timestamp: If set to true, prevents Splunk software from prepending a timestamp and hostname to incoming events.
        :param pulumi.Input[bool] no_priority_stripping: If set to true, Splunk software does not remove the priority field from incoming syslog events.
        :param pulumi.Input[str] queue: Which queue events from this input should be sent to. Generally this does not need to be changed.
        :param pulumi.Input[str] restrict_to_host: Restrict incoming connections on this port to the host specified here.
               If this is not set, the value specified in [udp://<remote server>:<port>] in inputs.conf is used.
        :param pulumi.Input[str] source: The value to populate in the source field for incoming events. The same source should not be used for multiple data inputs.
        :param pulumi.Input[str] sourcetype: The value to populate in the sourcetype field for incoming events.
        """
        if acl is not None:
            pulumi.set(__self__, "acl", acl)
        if connection_host is not None:
            pulumi.set(__self__, "connection_host", connection_host)
        if disabled is not None:
            pulumi.set(__self__, "disabled", disabled)
        if host is not None:
            pulumi.set(__self__, "host", host)
        if index is not None:
            pulumi.set(__self__, "index", index)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if no_appending_timestamp is not None:
            pulumi.set(__self__, "no_appending_timestamp", no_appending_timestamp)
        if no_priority_stripping is not None:
            pulumi.set(__self__, "no_priority_stripping", no_priority_stripping)
        if queue is not None:
            pulumi.set(__self__, "queue", queue)
        if restrict_to_host is not None:
            pulumi.set(__self__, "restrict_to_host", restrict_to_host)
        if source is not None:
            pulumi.set(__self__, "source", source)
        if sourcetype is not None:
            pulumi.set(__self__, "sourcetype", sourcetype)

    @property
    @pulumi.getter
    def acl(self) -> Optional[pulumi.Input['InputsUdpAclArgs']]:
        """
        The app/user context that is the namespace for the resource
        """
        return pulumi.get(self, "acl")

    @acl.setter
    def acl(self, value: Optional[pulumi.Input['InputsUdpAclArgs']]):
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
        The value to populate in the host field for incoming events. This is used during parsing/indexing, in particular to set the host field. It is also the host field used at search time.
        """
        return pulumi.get(self, "host")

    @host.setter
    def host(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "host", value)

    @property
    @pulumi.getter
    def index(self) -> Optional[pulumi.Input[str]]:
        """
        Which index events from this input should be stored in. Defaults to default.
        """
        return pulumi.get(self, "index")

    @index.setter
    def index(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "index", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The UDP port that this input should listen on.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="noAppendingTimestamp")
    def no_appending_timestamp(self) -> Optional[pulumi.Input[bool]]:
        """
        If set to true, prevents Splunk software from prepending a timestamp and hostname to incoming events.
        """
        return pulumi.get(self, "no_appending_timestamp")

    @no_appending_timestamp.setter
    def no_appending_timestamp(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "no_appending_timestamp", value)

    @property
    @pulumi.getter(name="noPriorityStripping")
    def no_priority_stripping(self) -> Optional[pulumi.Input[bool]]:
        """
        If set to true, Splunk software does not remove the priority field from incoming syslog events.
        """
        return pulumi.get(self, "no_priority_stripping")

    @no_priority_stripping.setter
    def no_priority_stripping(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "no_priority_stripping", value)

    @property
    @pulumi.getter
    def queue(self) -> Optional[pulumi.Input[str]]:
        """
        Which queue events from this input should be sent to. Generally this does not need to be changed.
        """
        return pulumi.get(self, "queue")

    @queue.setter
    def queue(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "queue", value)

    @property
    @pulumi.getter(name="restrictToHost")
    def restrict_to_host(self) -> Optional[pulumi.Input[str]]:
        """
        Restrict incoming connections on this port to the host specified here.
        If this is not set, the value specified in [udp://<remote server>:<port>] in inputs.conf is used.
        """
        return pulumi.get(self, "restrict_to_host")

    @restrict_to_host.setter
    def restrict_to_host(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "restrict_to_host", value)

    @property
    @pulumi.getter
    def source(self) -> Optional[pulumi.Input[str]]:
        """
        The value to populate in the source field for incoming events. The same source should not be used for multiple data inputs.
        """
        return pulumi.get(self, "source")

    @source.setter
    def source(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "source", value)

    @property
    @pulumi.getter
    def sourcetype(self) -> Optional[pulumi.Input[str]]:
        """
        The value to populate in the sourcetype field for incoming events.
        """
        return pulumi.get(self, "sourcetype")

    @sourcetype.setter
    def sourcetype(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "sourcetype", value)


class InputsUdp(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 acl: Optional[pulumi.Input[Union['InputsUdpAclArgs', 'InputsUdpAclArgsDict']]] = None,
                 connection_host: Optional[pulumi.Input[str]] = None,
                 disabled: Optional[pulumi.Input[bool]] = None,
                 host: Optional[pulumi.Input[str]] = None,
                 index: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 no_appending_timestamp: Optional[pulumi.Input[bool]] = None,
                 no_priority_stripping: Optional[pulumi.Input[bool]] = None,
                 queue: Optional[pulumi.Input[str]] = None,
                 restrict_to_host: Optional[pulumi.Input[str]] = None,
                 source: Optional[pulumi.Input[str]] = None,
                 sourcetype: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        ## # Resource: InputsTcpRaw

        Create and manage UDP data inputs.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_splunk as splunk

        udp = splunk.InputsUdp("udp",
            name="41000",
            index="main",
            source="new",
            sourcetype="new",
            disabled=False)
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Union['InputsUdpAclArgs', 'InputsUdpAclArgsDict']] acl: The app/user context that is the namespace for the resource
        :param pulumi.Input[str] connection_host: Valid values: (ip | dns | none)
               Set the host for the remote server that is sending data.
               ip sets the host to the IP address of the remote server sending data.
               dns sets the host to the reverse DNS entry for the IP address of the remote server sending data.
               none leaves the host as specified in inputs.conf, which is typically the Splunk system hostname.
               Default value is dns.
        :param pulumi.Input[bool] disabled: Indicates if input is disabled.
        :param pulumi.Input[str] host: The value to populate in the host field for incoming events. This is used during parsing/indexing, in particular to set the host field. It is also the host field used at search time.
        :param pulumi.Input[str] index: Which index events from this input should be stored in. Defaults to default.
        :param pulumi.Input[str] name: The UDP port that this input should listen on.
        :param pulumi.Input[bool] no_appending_timestamp: If set to true, prevents Splunk software from prepending a timestamp and hostname to incoming events.
        :param pulumi.Input[bool] no_priority_stripping: If set to true, Splunk software does not remove the priority field from incoming syslog events.
        :param pulumi.Input[str] queue: Which queue events from this input should be sent to. Generally this does not need to be changed.
        :param pulumi.Input[str] restrict_to_host: Restrict incoming connections on this port to the host specified here.
               If this is not set, the value specified in [udp://<remote server>:<port>] in inputs.conf is used.
        :param pulumi.Input[str] source: The value to populate in the source field for incoming events. The same source should not be used for multiple data inputs.
        :param pulumi.Input[str] sourcetype: The value to populate in the sourcetype field for incoming events.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[InputsUdpArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        ## # Resource: InputsTcpRaw

        Create and manage UDP data inputs.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_splunk as splunk

        udp = splunk.InputsUdp("udp",
            name="41000",
            index="main",
            source="new",
            sourcetype="new",
            disabled=False)
        ```

        :param str resource_name: The name of the resource.
        :param InputsUdpArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(InputsUdpArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 acl: Optional[pulumi.Input[Union['InputsUdpAclArgs', 'InputsUdpAclArgsDict']]] = None,
                 connection_host: Optional[pulumi.Input[str]] = None,
                 disabled: Optional[pulumi.Input[bool]] = None,
                 host: Optional[pulumi.Input[str]] = None,
                 index: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 no_appending_timestamp: Optional[pulumi.Input[bool]] = None,
                 no_priority_stripping: Optional[pulumi.Input[bool]] = None,
                 queue: Optional[pulumi.Input[str]] = None,
                 restrict_to_host: Optional[pulumi.Input[str]] = None,
                 source: Optional[pulumi.Input[str]] = None,
                 sourcetype: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = InputsUdpArgs.__new__(InputsUdpArgs)

            __props__.__dict__["acl"] = acl
            __props__.__dict__["connection_host"] = connection_host
            __props__.__dict__["disabled"] = disabled
            __props__.__dict__["host"] = host
            __props__.__dict__["index"] = index
            __props__.__dict__["name"] = name
            __props__.__dict__["no_appending_timestamp"] = no_appending_timestamp
            __props__.__dict__["no_priority_stripping"] = no_priority_stripping
            __props__.__dict__["queue"] = queue
            __props__.__dict__["restrict_to_host"] = restrict_to_host
            __props__.__dict__["source"] = source
            __props__.__dict__["sourcetype"] = sourcetype
        super(InputsUdp, __self__).__init__(
            'splunk:index/inputsUdp:InputsUdp',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            acl: Optional[pulumi.Input[Union['InputsUdpAclArgs', 'InputsUdpAclArgsDict']]] = None,
            connection_host: Optional[pulumi.Input[str]] = None,
            disabled: Optional[pulumi.Input[bool]] = None,
            host: Optional[pulumi.Input[str]] = None,
            index: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            no_appending_timestamp: Optional[pulumi.Input[bool]] = None,
            no_priority_stripping: Optional[pulumi.Input[bool]] = None,
            queue: Optional[pulumi.Input[str]] = None,
            restrict_to_host: Optional[pulumi.Input[str]] = None,
            source: Optional[pulumi.Input[str]] = None,
            sourcetype: Optional[pulumi.Input[str]] = None) -> 'InputsUdp':
        """
        Get an existing InputsUdp resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Union['InputsUdpAclArgs', 'InputsUdpAclArgsDict']] acl: The app/user context that is the namespace for the resource
        :param pulumi.Input[str] connection_host: Valid values: (ip | dns | none)
               Set the host for the remote server that is sending data.
               ip sets the host to the IP address of the remote server sending data.
               dns sets the host to the reverse DNS entry for the IP address of the remote server sending data.
               none leaves the host as specified in inputs.conf, which is typically the Splunk system hostname.
               Default value is dns.
        :param pulumi.Input[bool] disabled: Indicates if input is disabled.
        :param pulumi.Input[str] host: The value to populate in the host field for incoming events. This is used during parsing/indexing, in particular to set the host field. It is also the host field used at search time.
        :param pulumi.Input[str] index: Which index events from this input should be stored in. Defaults to default.
        :param pulumi.Input[str] name: The UDP port that this input should listen on.
        :param pulumi.Input[bool] no_appending_timestamp: If set to true, prevents Splunk software from prepending a timestamp and hostname to incoming events.
        :param pulumi.Input[bool] no_priority_stripping: If set to true, Splunk software does not remove the priority field from incoming syslog events.
        :param pulumi.Input[str] queue: Which queue events from this input should be sent to. Generally this does not need to be changed.
        :param pulumi.Input[str] restrict_to_host: Restrict incoming connections on this port to the host specified here.
               If this is not set, the value specified in [udp://<remote server>:<port>] in inputs.conf is used.
        :param pulumi.Input[str] source: The value to populate in the source field for incoming events. The same source should not be used for multiple data inputs.
        :param pulumi.Input[str] sourcetype: The value to populate in the sourcetype field for incoming events.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _InputsUdpState.__new__(_InputsUdpState)

        __props__.__dict__["acl"] = acl
        __props__.__dict__["connection_host"] = connection_host
        __props__.__dict__["disabled"] = disabled
        __props__.__dict__["host"] = host
        __props__.__dict__["index"] = index
        __props__.__dict__["name"] = name
        __props__.__dict__["no_appending_timestamp"] = no_appending_timestamp
        __props__.__dict__["no_priority_stripping"] = no_priority_stripping
        __props__.__dict__["queue"] = queue
        __props__.__dict__["restrict_to_host"] = restrict_to_host
        __props__.__dict__["source"] = source
        __props__.__dict__["sourcetype"] = sourcetype
        return InputsUdp(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def acl(self) -> pulumi.Output['outputs.InputsUdpAcl']:
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
        The value to populate in the host field for incoming events. This is used during parsing/indexing, in particular to set the host field. It is also the host field used at search time.
        """
        return pulumi.get(self, "host")

    @property
    @pulumi.getter
    def index(self) -> pulumi.Output[str]:
        """
        Which index events from this input should be stored in. Defaults to default.
        """
        return pulumi.get(self, "index")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The UDP port that this input should listen on.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="noAppendingTimestamp")
    def no_appending_timestamp(self) -> pulumi.Output[bool]:
        """
        If set to true, prevents Splunk software from prepending a timestamp and hostname to incoming events.
        """
        return pulumi.get(self, "no_appending_timestamp")

    @property
    @pulumi.getter(name="noPriorityStripping")
    def no_priority_stripping(self) -> pulumi.Output[bool]:
        """
        If set to true, Splunk software does not remove the priority field from incoming syslog events.
        """
        return pulumi.get(self, "no_priority_stripping")

    @property
    @pulumi.getter
    def queue(self) -> pulumi.Output[str]:
        """
        Which queue events from this input should be sent to. Generally this does not need to be changed.
        """
        return pulumi.get(self, "queue")

    @property
    @pulumi.getter(name="restrictToHost")
    def restrict_to_host(self) -> pulumi.Output[str]:
        """
        Restrict incoming connections on this port to the host specified here.
        If this is not set, the value specified in [udp://<remote server>:<port>] in inputs.conf is used.
        """
        return pulumi.get(self, "restrict_to_host")

    @property
    @pulumi.getter
    def source(self) -> pulumi.Output[str]:
        """
        The value to populate in the source field for incoming events. The same source should not be used for multiple data inputs.
        """
        return pulumi.get(self, "source")

    @property
    @pulumi.getter
    def sourcetype(self) -> pulumi.Output[str]:
        """
        The value to populate in the sourcetype field for incoming events.
        """
        return pulumi.get(self, "sourcetype")

