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

__all__ = ['InputsHttpEventCollectorArgs', 'InputsHttpEventCollector']

@pulumi.input_type
class InputsHttpEventCollectorArgs:
    def __init__(__self__, *,
                 acl: Optional[pulumi.Input['InputsHttpEventCollectorAclArgs']] = None,
                 disabled: Optional[pulumi.Input[bool]] = None,
                 host: Optional[pulumi.Input[str]] = None,
                 index: Optional[pulumi.Input[str]] = None,
                 indexes: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 source: Optional[pulumi.Input[str]] = None,
                 sourcetype: Optional[pulumi.Input[str]] = None,
                 token: Optional[pulumi.Input[str]] = None,
                 use_ack: Optional[pulumi.Input[int]] = None):
        """
        The set of arguments for constructing a InputsHttpEventCollector resource.
        :param pulumi.Input['InputsHttpEventCollectorAclArgs'] acl: The app/user context that is the namespace for the resource
        :param pulumi.Input[bool] disabled: Input disabled indicator
        :param pulumi.Input[str] host: Default host value for events with this token
        :param pulumi.Input[str] index: Index to store generated events
        :param pulumi.Input[Sequence[pulumi.Input[str]]] indexes: Set of indexes allowed for events with this token
        :param pulumi.Input[str] name: Token name (inputs.conf key)
        :param pulumi.Input[str] source: Default source for events with this token
        :param pulumi.Input[str] sourcetype: Default source type for events with this token
        :param pulumi.Input[str] token: Token value for sending data to collector/event endpoint
        :param pulumi.Input[int] use_ack: Indexer acknowledgement for this token
        """
        if acl is not None:
            pulumi.set(__self__, "acl", acl)
        if disabled is not None:
            pulumi.set(__self__, "disabled", disabled)
        if host is not None:
            pulumi.set(__self__, "host", host)
        if index is not None:
            pulumi.set(__self__, "index", index)
        if indexes is not None:
            pulumi.set(__self__, "indexes", indexes)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if source is not None:
            pulumi.set(__self__, "source", source)
        if sourcetype is not None:
            pulumi.set(__self__, "sourcetype", sourcetype)
        if token is not None:
            pulumi.set(__self__, "token", token)
        if use_ack is not None:
            pulumi.set(__self__, "use_ack", use_ack)

    @property
    @pulumi.getter
    def acl(self) -> Optional[pulumi.Input['InputsHttpEventCollectorAclArgs']]:
        """
        The app/user context that is the namespace for the resource
        """
        return pulumi.get(self, "acl")

    @acl.setter
    def acl(self, value: Optional[pulumi.Input['InputsHttpEventCollectorAclArgs']]):
        pulumi.set(self, "acl", value)

    @property
    @pulumi.getter
    def disabled(self) -> Optional[pulumi.Input[bool]]:
        """
        Input disabled indicator
        """
        return pulumi.get(self, "disabled")

    @disabled.setter
    def disabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "disabled", value)

    @property
    @pulumi.getter
    def host(self) -> Optional[pulumi.Input[str]]:
        """
        Default host value for events with this token
        """
        return pulumi.get(self, "host")

    @host.setter
    def host(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "host", value)

    @property
    @pulumi.getter
    def index(self) -> Optional[pulumi.Input[str]]:
        """
        Index to store generated events
        """
        return pulumi.get(self, "index")

    @index.setter
    def index(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "index", value)

    @property
    @pulumi.getter
    def indexes(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Set of indexes allowed for events with this token
        """
        return pulumi.get(self, "indexes")

    @indexes.setter
    def indexes(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "indexes", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Token name (inputs.conf key)
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def source(self) -> Optional[pulumi.Input[str]]:
        """
        Default source for events with this token
        """
        return pulumi.get(self, "source")

    @source.setter
    def source(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "source", value)

    @property
    @pulumi.getter
    def sourcetype(self) -> Optional[pulumi.Input[str]]:
        """
        Default source type for events with this token
        """
        return pulumi.get(self, "sourcetype")

    @sourcetype.setter
    def sourcetype(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "sourcetype", value)

    @property
    @pulumi.getter
    def token(self) -> Optional[pulumi.Input[str]]:
        """
        Token value for sending data to collector/event endpoint
        """
        return pulumi.get(self, "token")

    @token.setter
    def token(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "token", value)

    @property
    @pulumi.getter(name="useAck")
    def use_ack(self) -> Optional[pulumi.Input[int]]:
        """
        Indexer acknowledgement for this token
        """
        return pulumi.get(self, "use_ack")

    @use_ack.setter
    def use_ack(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "use_ack", value)


@pulumi.input_type
class _InputsHttpEventCollectorState:
    def __init__(__self__, *,
                 acl: Optional[pulumi.Input['InputsHttpEventCollectorAclArgs']] = None,
                 disabled: Optional[pulumi.Input[bool]] = None,
                 host: Optional[pulumi.Input[str]] = None,
                 index: Optional[pulumi.Input[str]] = None,
                 indexes: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 source: Optional[pulumi.Input[str]] = None,
                 sourcetype: Optional[pulumi.Input[str]] = None,
                 token: Optional[pulumi.Input[str]] = None,
                 use_ack: Optional[pulumi.Input[int]] = None):
        """
        Input properties used for looking up and filtering InputsHttpEventCollector resources.
        :param pulumi.Input['InputsHttpEventCollectorAclArgs'] acl: The app/user context that is the namespace for the resource
        :param pulumi.Input[bool] disabled: Input disabled indicator
        :param pulumi.Input[str] host: Default host value for events with this token
        :param pulumi.Input[str] index: Index to store generated events
        :param pulumi.Input[Sequence[pulumi.Input[str]]] indexes: Set of indexes allowed for events with this token
        :param pulumi.Input[str] name: Token name (inputs.conf key)
        :param pulumi.Input[str] source: Default source for events with this token
        :param pulumi.Input[str] sourcetype: Default source type for events with this token
        :param pulumi.Input[str] token: Token value for sending data to collector/event endpoint
        :param pulumi.Input[int] use_ack: Indexer acknowledgement for this token
        """
        if acl is not None:
            pulumi.set(__self__, "acl", acl)
        if disabled is not None:
            pulumi.set(__self__, "disabled", disabled)
        if host is not None:
            pulumi.set(__self__, "host", host)
        if index is not None:
            pulumi.set(__self__, "index", index)
        if indexes is not None:
            pulumi.set(__self__, "indexes", indexes)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if source is not None:
            pulumi.set(__self__, "source", source)
        if sourcetype is not None:
            pulumi.set(__self__, "sourcetype", sourcetype)
        if token is not None:
            pulumi.set(__self__, "token", token)
        if use_ack is not None:
            pulumi.set(__self__, "use_ack", use_ack)

    @property
    @pulumi.getter
    def acl(self) -> Optional[pulumi.Input['InputsHttpEventCollectorAclArgs']]:
        """
        The app/user context that is the namespace for the resource
        """
        return pulumi.get(self, "acl")

    @acl.setter
    def acl(self, value: Optional[pulumi.Input['InputsHttpEventCollectorAclArgs']]):
        pulumi.set(self, "acl", value)

    @property
    @pulumi.getter
    def disabled(self) -> Optional[pulumi.Input[bool]]:
        """
        Input disabled indicator
        """
        return pulumi.get(self, "disabled")

    @disabled.setter
    def disabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "disabled", value)

    @property
    @pulumi.getter
    def host(self) -> Optional[pulumi.Input[str]]:
        """
        Default host value for events with this token
        """
        return pulumi.get(self, "host")

    @host.setter
    def host(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "host", value)

    @property
    @pulumi.getter
    def index(self) -> Optional[pulumi.Input[str]]:
        """
        Index to store generated events
        """
        return pulumi.get(self, "index")

    @index.setter
    def index(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "index", value)

    @property
    @pulumi.getter
    def indexes(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Set of indexes allowed for events with this token
        """
        return pulumi.get(self, "indexes")

    @indexes.setter
    def indexes(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "indexes", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Token name (inputs.conf key)
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def source(self) -> Optional[pulumi.Input[str]]:
        """
        Default source for events with this token
        """
        return pulumi.get(self, "source")

    @source.setter
    def source(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "source", value)

    @property
    @pulumi.getter
    def sourcetype(self) -> Optional[pulumi.Input[str]]:
        """
        Default source type for events with this token
        """
        return pulumi.get(self, "sourcetype")

    @sourcetype.setter
    def sourcetype(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "sourcetype", value)

    @property
    @pulumi.getter
    def token(self) -> Optional[pulumi.Input[str]]:
        """
        Token value for sending data to collector/event endpoint
        """
        return pulumi.get(self, "token")

    @token.setter
    def token(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "token", value)

    @property
    @pulumi.getter(name="useAck")
    def use_ack(self) -> Optional[pulumi.Input[int]]:
        """
        Indexer acknowledgement for this token
        """
        return pulumi.get(self, "use_ack")

    @use_ack.setter
    def use_ack(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "use_ack", value)


class InputsHttpEventCollector(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 acl: Optional[pulumi.Input[pulumi.InputType['InputsHttpEventCollectorAclArgs']]] = None,
                 disabled: Optional[pulumi.Input[bool]] = None,
                 host: Optional[pulumi.Input[str]] = None,
                 index: Optional[pulumi.Input[str]] = None,
                 indexes: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 source: Optional[pulumi.Input[str]] = None,
                 sourcetype: Optional[pulumi.Input[str]] = None,
                 token: Optional[pulumi.Input[str]] = None,
                 use_ack: Optional[pulumi.Input[int]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        ## # Resource: InputsHttpEventCollector

        Create or update HTTP Event Collector input configuration tokens.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['InputsHttpEventCollectorAclArgs']] acl: The app/user context that is the namespace for the resource
        :param pulumi.Input[bool] disabled: Input disabled indicator
        :param pulumi.Input[str] host: Default host value for events with this token
        :param pulumi.Input[str] index: Index to store generated events
        :param pulumi.Input[Sequence[pulumi.Input[str]]] indexes: Set of indexes allowed for events with this token
        :param pulumi.Input[str] name: Token name (inputs.conf key)
        :param pulumi.Input[str] source: Default source for events with this token
        :param pulumi.Input[str] sourcetype: Default source type for events with this token
        :param pulumi.Input[str] token: Token value for sending data to collector/event endpoint
        :param pulumi.Input[int] use_ack: Indexer acknowledgement for this token
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[InputsHttpEventCollectorArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        ## # Resource: InputsHttpEventCollector

        Create or update HTTP Event Collector input configuration tokens.

        :param str resource_name: The name of the resource.
        :param InputsHttpEventCollectorArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(InputsHttpEventCollectorArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 acl: Optional[pulumi.Input[pulumi.InputType['InputsHttpEventCollectorAclArgs']]] = None,
                 disabled: Optional[pulumi.Input[bool]] = None,
                 host: Optional[pulumi.Input[str]] = None,
                 index: Optional[pulumi.Input[str]] = None,
                 indexes: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 source: Optional[pulumi.Input[str]] = None,
                 sourcetype: Optional[pulumi.Input[str]] = None,
                 token: Optional[pulumi.Input[str]] = None,
                 use_ack: Optional[pulumi.Input[int]] = None,
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
            __props__ = InputsHttpEventCollectorArgs.__new__(InputsHttpEventCollectorArgs)

            __props__.__dict__["acl"] = acl
            __props__.__dict__["disabled"] = disabled
            __props__.__dict__["host"] = host
            __props__.__dict__["index"] = index
            __props__.__dict__["indexes"] = indexes
            __props__.__dict__["name"] = name
            __props__.__dict__["source"] = source
            __props__.__dict__["sourcetype"] = sourcetype
            __props__.__dict__["token"] = token
            __props__.__dict__["use_ack"] = use_ack
        super(InputsHttpEventCollector, __self__).__init__(
            'splunk:index/inputsHttpEventCollector:InputsHttpEventCollector',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            acl: Optional[pulumi.Input[pulumi.InputType['InputsHttpEventCollectorAclArgs']]] = None,
            disabled: Optional[pulumi.Input[bool]] = None,
            host: Optional[pulumi.Input[str]] = None,
            index: Optional[pulumi.Input[str]] = None,
            indexes: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            name: Optional[pulumi.Input[str]] = None,
            source: Optional[pulumi.Input[str]] = None,
            sourcetype: Optional[pulumi.Input[str]] = None,
            token: Optional[pulumi.Input[str]] = None,
            use_ack: Optional[pulumi.Input[int]] = None) -> 'InputsHttpEventCollector':
        """
        Get an existing InputsHttpEventCollector resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['InputsHttpEventCollectorAclArgs']] acl: The app/user context that is the namespace for the resource
        :param pulumi.Input[bool] disabled: Input disabled indicator
        :param pulumi.Input[str] host: Default host value for events with this token
        :param pulumi.Input[str] index: Index to store generated events
        :param pulumi.Input[Sequence[pulumi.Input[str]]] indexes: Set of indexes allowed for events with this token
        :param pulumi.Input[str] name: Token name (inputs.conf key)
        :param pulumi.Input[str] source: Default source for events with this token
        :param pulumi.Input[str] sourcetype: Default source type for events with this token
        :param pulumi.Input[str] token: Token value for sending data to collector/event endpoint
        :param pulumi.Input[int] use_ack: Indexer acknowledgement for this token
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _InputsHttpEventCollectorState.__new__(_InputsHttpEventCollectorState)

        __props__.__dict__["acl"] = acl
        __props__.__dict__["disabled"] = disabled
        __props__.__dict__["host"] = host
        __props__.__dict__["index"] = index
        __props__.__dict__["indexes"] = indexes
        __props__.__dict__["name"] = name
        __props__.__dict__["source"] = source
        __props__.__dict__["sourcetype"] = sourcetype
        __props__.__dict__["token"] = token
        __props__.__dict__["use_ack"] = use_ack
        return InputsHttpEventCollector(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def acl(self) -> pulumi.Output['outputs.InputsHttpEventCollectorAcl']:
        """
        The app/user context that is the namespace for the resource
        """
        return pulumi.get(self, "acl")

    @property
    @pulumi.getter
    def disabled(self) -> pulumi.Output[bool]:
        """
        Input disabled indicator
        """
        return pulumi.get(self, "disabled")

    @property
    @pulumi.getter
    def host(self) -> pulumi.Output[str]:
        """
        Default host value for events with this token
        """
        return pulumi.get(self, "host")

    @property
    @pulumi.getter
    def index(self) -> pulumi.Output[str]:
        """
        Index to store generated events
        """
        return pulumi.get(self, "index")

    @property
    @pulumi.getter
    def indexes(self) -> pulumi.Output[Sequence[str]]:
        """
        Set of indexes allowed for events with this token
        """
        return pulumi.get(self, "indexes")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Token name (inputs.conf key)
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def source(self) -> pulumi.Output[str]:
        """
        Default source for events with this token
        """
        return pulumi.get(self, "source")

    @property
    @pulumi.getter
    def sourcetype(self) -> pulumi.Output[str]:
        """
        Default source type for events with this token
        """
        return pulumi.get(self, "sourcetype")

    @property
    @pulumi.getter
    def token(self) -> pulumi.Output[str]:
        """
        Token value for sending data to collector/event endpoint
        """
        return pulumi.get(self, "token")

    @property
    @pulumi.getter(name="useAck")
    def use_ack(self) -> pulumi.Output[int]:
        """
        Indexer acknowledgement for this token
        """
        return pulumi.get(self, "use_ack")

