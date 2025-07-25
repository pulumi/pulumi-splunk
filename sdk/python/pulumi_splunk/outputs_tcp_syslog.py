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

__all__ = ['OutputsTcpSyslogArgs', 'OutputsTcpSyslog']

@pulumi.input_type
class OutputsTcpSyslogArgs:
    def __init__(__self__, *,
                 acl: Optional[pulumi.Input['OutputsTcpSyslogAclArgs']] = None,
                 disabled: Optional[pulumi.Input[_builtins.bool]] = None,
                 name: Optional[pulumi.Input[_builtins.str]] = None,
                 priority: Optional[pulumi.Input[_builtins.int]] = None,
                 server: Optional[pulumi.Input[_builtins.str]] = None,
                 syslog_sourcetype: Optional[pulumi.Input[_builtins.str]] = None,
                 timestamp_format: Optional[pulumi.Input[_builtins.str]] = None,
                 type: Optional[pulumi.Input[_builtins.str]] = None):
        """
        The set of arguments for constructing a OutputsTcpSyslog resource.
        :param pulumi.Input['OutputsTcpSyslogAclArgs'] acl: The app/user context that is the namespace for the resource
        :param pulumi.Input[_builtins.bool] disabled: If true, disables global syslog settings.
        :param pulumi.Input[_builtins.str] name: Name of the syslog output group. This is name used when creating syslog configuration in outputs.conf.
        :param pulumi.Input[_builtins.int] priority: Sets syslog priority value. The priority value should specified as an integer. See $SPLUNK_HOME/etc/system/README/outputs.conf.spec for details.
        :param pulumi.Input[_builtins.str] server: host:port of the server where syslog data should be sent
        :param pulumi.Input[_builtins.str] syslog_sourcetype: Specifies a rule for handling data in addition to that provided by the "syslog" sourcetype. By default, there is no value for syslogSourceType.
               <br>This string is used as a substring match against the sourcetype key. For example, if the string is set to 'syslog', then all source types containing the string "syslog" receives this special treatment.
               To match a source type explicitly, use the pattern "sourcetype::sourcetype_name." For example
               syslogSourcetype = sourcetype::apache_common
               Data that is "syslog" or matches this setting is assumed to already be in syslog format.
               Data that does not match the rules has a header, potentially a timestamp, and a hostname added to the front of the event. This is how Splunk software causes arbitrary log data to match syslog expectations.
        :param pulumi.Input[_builtins.str] timestamp_format: Format of timestamp to add at start of the events to be forwarded.
               The format is a strftime-style timestamp formatting string. See $SPLUNK_HOME/etc/system/README/outputs.conf.spec for details.
        :param pulumi.Input[_builtins.str] type: Protocol to use to send syslog data. Valid values: (tcp | udp ).
        """
        if acl is not None:
            pulumi.set(__self__, "acl", acl)
        if disabled is not None:
            pulumi.set(__self__, "disabled", disabled)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if priority is not None:
            pulumi.set(__self__, "priority", priority)
        if server is not None:
            pulumi.set(__self__, "server", server)
        if syslog_sourcetype is not None:
            pulumi.set(__self__, "syslog_sourcetype", syslog_sourcetype)
        if timestamp_format is not None:
            pulumi.set(__self__, "timestamp_format", timestamp_format)
        if type is not None:
            pulumi.set(__self__, "type", type)

    @_builtins.property
    @pulumi.getter
    def acl(self) -> Optional[pulumi.Input['OutputsTcpSyslogAclArgs']]:
        """
        The app/user context that is the namespace for the resource
        """
        return pulumi.get(self, "acl")

    @acl.setter
    def acl(self, value: Optional[pulumi.Input['OutputsTcpSyslogAclArgs']]):
        pulumi.set(self, "acl", value)

    @_builtins.property
    @pulumi.getter
    def disabled(self) -> Optional[pulumi.Input[_builtins.bool]]:
        """
        If true, disables global syslog settings.
        """
        return pulumi.get(self, "disabled")

    @disabled.setter
    def disabled(self, value: Optional[pulumi.Input[_builtins.bool]]):
        pulumi.set(self, "disabled", value)

    @_builtins.property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[_builtins.str]]:
        """
        Name of the syslog output group. This is name used when creating syslog configuration in outputs.conf.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "name", value)

    @_builtins.property
    @pulumi.getter
    def priority(self) -> Optional[pulumi.Input[_builtins.int]]:
        """
        Sets syslog priority value. The priority value should specified as an integer. See $SPLUNK_HOME/etc/system/README/outputs.conf.spec for details.
        """
        return pulumi.get(self, "priority")

    @priority.setter
    def priority(self, value: Optional[pulumi.Input[_builtins.int]]):
        pulumi.set(self, "priority", value)

    @_builtins.property
    @pulumi.getter
    def server(self) -> Optional[pulumi.Input[_builtins.str]]:
        """
        host:port of the server where syslog data should be sent
        """
        return pulumi.get(self, "server")

    @server.setter
    def server(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "server", value)

    @_builtins.property
    @pulumi.getter(name="syslogSourcetype")
    def syslog_sourcetype(self) -> Optional[pulumi.Input[_builtins.str]]:
        """
        Specifies a rule for handling data in addition to that provided by the "syslog" sourcetype. By default, there is no value for syslogSourceType.
        <br>This string is used as a substring match against the sourcetype key. For example, if the string is set to 'syslog', then all source types containing the string "syslog" receives this special treatment.
        To match a source type explicitly, use the pattern "sourcetype::sourcetype_name." For example
        syslogSourcetype = sourcetype::apache_common
        Data that is "syslog" or matches this setting is assumed to already be in syslog format.
        Data that does not match the rules has a header, potentially a timestamp, and a hostname added to the front of the event. This is how Splunk software causes arbitrary log data to match syslog expectations.
        """
        return pulumi.get(self, "syslog_sourcetype")

    @syslog_sourcetype.setter
    def syslog_sourcetype(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "syslog_sourcetype", value)

    @_builtins.property
    @pulumi.getter(name="timestampFormat")
    def timestamp_format(self) -> Optional[pulumi.Input[_builtins.str]]:
        """
        Format of timestamp to add at start of the events to be forwarded.
        The format is a strftime-style timestamp formatting string. See $SPLUNK_HOME/etc/system/README/outputs.conf.spec for details.
        """
        return pulumi.get(self, "timestamp_format")

    @timestamp_format.setter
    def timestamp_format(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "timestamp_format", value)

    @_builtins.property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input[_builtins.str]]:
        """
        Protocol to use to send syslog data. Valid values: (tcp | udp ).
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "type", value)


@pulumi.input_type
class _OutputsTcpSyslogState:
    def __init__(__self__, *,
                 acl: Optional[pulumi.Input['OutputsTcpSyslogAclArgs']] = None,
                 disabled: Optional[pulumi.Input[_builtins.bool]] = None,
                 name: Optional[pulumi.Input[_builtins.str]] = None,
                 priority: Optional[pulumi.Input[_builtins.int]] = None,
                 server: Optional[pulumi.Input[_builtins.str]] = None,
                 syslog_sourcetype: Optional[pulumi.Input[_builtins.str]] = None,
                 timestamp_format: Optional[pulumi.Input[_builtins.str]] = None,
                 type: Optional[pulumi.Input[_builtins.str]] = None):
        """
        Input properties used for looking up and filtering OutputsTcpSyslog resources.
        :param pulumi.Input['OutputsTcpSyslogAclArgs'] acl: The app/user context that is the namespace for the resource
        :param pulumi.Input[_builtins.bool] disabled: If true, disables global syslog settings.
        :param pulumi.Input[_builtins.str] name: Name of the syslog output group. This is name used when creating syslog configuration in outputs.conf.
        :param pulumi.Input[_builtins.int] priority: Sets syslog priority value. The priority value should specified as an integer. See $SPLUNK_HOME/etc/system/README/outputs.conf.spec for details.
        :param pulumi.Input[_builtins.str] server: host:port of the server where syslog data should be sent
        :param pulumi.Input[_builtins.str] syslog_sourcetype: Specifies a rule for handling data in addition to that provided by the "syslog" sourcetype. By default, there is no value for syslogSourceType.
               <br>This string is used as a substring match against the sourcetype key. For example, if the string is set to 'syslog', then all source types containing the string "syslog" receives this special treatment.
               To match a source type explicitly, use the pattern "sourcetype::sourcetype_name." For example
               syslogSourcetype = sourcetype::apache_common
               Data that is "syslog" or matches this setting is assumed to already be in syslog format.
               Data that does not match the rules has a header, potentially a timestamp, and a hostname added to the front of the event. This is how Splunk software causes arbitrary log data to match syslog expectations.
        :param pulumi.Input[_builtins.str] timestamp_format: Format of timestamp to add at start of the events to be forwarded.
               The format is a strftime-style timestamp formatting string. See $SPLUNK_HOME/etc/system/README/outputs.conf.spec for details.
        :param pulumi.Input[_builtins.str] type: Protocol to use to send syslog data. Valid values: (tcp | udp ).
        """
        if acl is not None:
            pulumi.set(__self__, "acl", acl)
        if disabled is not None:
            pulumi.set(__self__, "disabled", disabled)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if priority is not None:
            pulumi.set(__self__, "priority", priority)
        if server is not None:
            pulumi.set(__self__, "server", server)
        if syslog_sourcetype is not None:
            pulumi.set(__self__, "syslog_sourcetype", syslog_sourcetype)
        if timestamp_format is not None:
            pulumi.set(__self__, "timestamp_format", timestamp_format)
        if type is not None:
            pulumi.set(__self__, "type", type)

    @_builtins.property
    @pulumi.getter
    def acl(self) -> Optional[pulumi.Input['OutputsTcpSyslogAclArgs']]:
        """
        The app/user context that is the namespace for the resource
        """
        return pulumi.get(self, "acl")

    @acl.setter
    def acl(self, value: Optional[pulumi.Input['OutputsTcpSyslogAclArgs']]):
        pulumi.set(self, "acl", value)

    @_builtins.property
    @pulumi.getter
    def disabled(self) -> Optional[pulumi.Input[_builtins.bool]]:
        """
        If true, disables global syslog settings.
        """
        return pulumi.get(self, "disabled")

    @disabled.setter
    def disabled(self, value: Optional[pulumi.Input[_builtins.bool]]):
        pulumi.set(self, "disabled", value)

    @_builtins.property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[_builtins.str]]:
        """
        Name of the syslog output group. This is name used when creating syslog configuration in outputs.conf.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "name", value)

    @_builtins.property
    @pulumi.getter
    def priority(self) -> Optional[pulumi.Input[_builtins.int]]:
        """
        Sets syslog priority value. The priority value should specified as an integer. See $SPLUNK_HOME/etc/system/README/outputs.conf.spec for details.
        """
        return pulumi.get(self, "priority")

    @priority.setter
    def priority(self, value: Optional[pulumi.Input[_builtins.int]]):
        pulumi.set(self, "priority", value)

    @_builtins.property
    @pulumi.getter
    def server(self) -> Optional[pulumi.Input[_builtins.str]]:
        """
        host:port of the server where syslog data should be sent
        """
        return pulumi.get(self, "server")

    @server.setter
    def server(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "server", value)

    @_builtins.property
    @pulumi.getter(name="syslogSourcetype")
    def syslog_sourcetype(self) -> Optional[pulumi.Input[_builtins.str]]:
        """
        Specifies a rule for handling data in addition to that provided by the "syslog" sourcetype. By default, there is no value for syslogSourceType.
        <br>This string is used as a substring match against the sourcetype key. For example, if the string is set to 'syslog', then all source types containing the string "syslog" receives this special treatment.
        To match a source type explicitly, use the pattern "sourcetype::sourcetype_name." For example
        syslogSourcetype = sourcetype::apache_common
        Data that is "syslog" or matches this setting is assumed to already be in syslog format.
        Data that does not match the rules has a header, potentially a timestamp, and a hostname added to the front of the event. This is how Splunk software causes arbitrary log data to match syslog expectations.
        """
        return pulumi.get(self, "syslog_sourcetype")

    @syslog_sourcetype.setter
    def syslog_sourcetype(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "syslog_sourcetype", value)

    @_builtins.property
    @pulumi.getter(name="timestampFormat")
    def timestamp_format(self) -> Optional[pulumi.Input[_builtins.str]]:
        """
        Format of timestamp to add at start of the events to be forwarded.
        The format is a strftime-style timestamp formatting string. See $SPLUNK_HOME/etc/system/README/outputs.conf.spec for details.
        """
        return pulumi.get(self, "timestamp_format")

    @timestamp_format.setter
    def timestamp_format(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "timestamp_format", value)

    @_builtins.property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input[_builtins.str]]:
        """
        Protocol to use to send syslog data. Valid values: (tcp | udp ).
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "type", value)


@pulumi.type_token("splunk:index/outputsTcpSyslog:OutputsTcpSyslog")
class OutputsTcpSyslog(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 acl: Optional[pulumi.Input[Union['OutputsTcpSyslogAclArgs', 'OutputsTcpSyslogAclArgsDict']]] = None,
                 disabled: Optional[pulumi.Input[_builtins.bool]] = None,
                 name: Optional[pulumi.Input[_builtins.str]] = None,
                 priority: Optional[pulumi.Input[_builtins.int]] = None,
                 server: Optional[pulumi.Input[_builtins.str]] = None,
                 syslog_sourcetype: Optional[pulumi.Input[_builtins.str]] = None,
                 timestamp_format: Optional[pulumi.Input[_builtins.str]] = None,
                 type: Optional[pulumi.Input[_builtins.str]] = None,
                 __props__=None):
        """
        ## # Resource: OutputsTcpSyslog

        Access the configuration of a forwarded server configured to provide data in standard syslog format.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_splunk as splunk

        tcp_syslog = splunk.OutputsTcpSyslog("tcp_syslog",
            name="new-syslog",
            server="new-host-1:1234",
            priority=5)
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Union['OutputsTcpSyslogAclArgs', 'OutputsTcpSyslogAclArgsDict']] acl: The app/user context that is the namespace for the resource
        :param pulumi.Input[_builtins.bool] disabled: If true, disables global syslog settings.
        :param pulumi.Input[_builtins.str] name: Name of the syslog output group. This is name used when creating syslog configuration in outputs.conf.
        :param pulumi.Input[_builtins.int] priority: Sets syslog priority value. The priority value should specified as an integer. See $SPLUNK_HOME/etc/system/README/outputs.conf.spec for details.
        :param pulumi.Input[_builtins.str] server: host:port of the server where syslog data should be sent
        :param pulumi.Input[_builtins.str] syslog_sourcetype: Specifies a rule for handling data in addition to that provided by the "syslog" sourcetype. By default, there is no value for syslogSourceType.
               <br>This string is used as a substring match against the sourcetype key. For example, if the string is set to 'syslog', then all source types containing the string "syslog" receives this special treatment.
               To match a source type explicitly, use the pattern "sourcetype::sourcetype_name." For example
               syslogSourcetype = sourcetype::apache_common
               Data that is "syslog" or matches this setting is assumed to already be in syslog format.
               Data that does not match the rules has a header, potentially a timestamp, and a hostname added to the front of the event. This is how Splunk software causes arbitrary log data to match syslog expectations.
        :param pulumi.Input[_builtins.str] timestamp_format: Format of timestamp to add at start of the events to be forwarded.
               The format is a strftime-style timestamp formatting string. See $SPLUNK_HOME/etc/system/README/outputs.conf.spec for details.
        :param pulumi.Input[_builtins.str] type: Protocol to use to send syslog data. Valid values: (tcp | udp ).
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[OutputsTcpSyslogArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        ## # Resource: OutputsTcpSyslog

        Access the configuration of a forwarded server configured to provide data in standard syslog format.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_splunk as splunk

        tcp_syslog = splunk.OutputsTcpSyslog("tcp_syslog",
            name="new-syslog",
            server="new-host-1:1234",
            priority=5)
        ```

        :param str resource_name: The name of the resource.
        :param OutputsTcpSyslogArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(OutputsTcpSyslogArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 acl: Optional[pulumi.Input[Union['OutputsTcpSyslogAclArgs', 'OutputsTcpSyslogAclArgsDict']]] = None,
                 disabled: Optional[pulumi.Input[_builtins.bool]] = None,
                 name: Optional[pulumi.Input[_builtins.str]] = None,
                 priority: Optional[pulumi.Input[_builtins.int]] = None,
                 server: Optional[pulumi.Input[_builtins.str]] = None,
                 syslog_sourcetype: Optional[pulumi.Input[_builtins.str]] = None,
                 timestamp_format: Optional[pulumi.Input[_builtins.str]] = None,
                 type: Optional[pulumi.Input[_builtins.str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = OutputsTcpSyslogArgs.__new__(OutputsTcpSyslogArgs)

            __props__.__dict__["acl"] = acl
            __props__.__dict__["disabled"] = disabled
            __props__.__dict__["name"] = name
            __props__.__dict__["priority"] = priority
            __props__.__dict__["server"] = server
            __props__.__dict__["syslog_sourcetype"] = syslog_sourcetype
            __props__.__dict__["timestamp_format"] = timestamp_format
            __props__.__dict__["type"] = type
        super(OutputsTcpSyslog, __self__).__init__(
            'splunk:index/outputsTcpSyslog:OutputsTcpSyslog',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            acl: Optional[pulumi.Input[Union['OutputsTcpSyslogAclArgs', 'OutputsTcpSyslogAclArgsDict']]] = None,
            disabled: Optional[pulumi.Input[_builtins.bool]] = None,
            name: Optional[pulumi.Input[_builtins.str]] = None,
            priority: Optional[pulumi.Input[_builtins.int]] = None,
            server: Optional[pulumi.Input[_builtins.str]] = None,
            syslog_sourcetype: Optional[pulumi.Input[_builtins.str]] = None,
            timestamp_format: Optional[pulumi.Input[_builtins.str]] = None,
            type: Optional[pulumi.Input[_builtins.str]] = None) -> 'OutputsTcpSyslog':
        """
        Get an existing OutputsTcpSyslog resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Union['OutputsTcpSyslogAclArgs', 'OutputsTcpSyslogAclArgsDict']] acl: The app/user context that is the namespace for the resource
        :param pulumi.Input[_builtins.bool] disabled: If true, disables global syslog settings.
        :param pulumi.Input[_builtins.str] name: Name of the syslog output group. This is name used when creating syslog configuration in outputs.conf.
        :param pulumi.Input[_builtins.int] priority: Sets syslog priority value. The priority value should specified as an integer. See $SPLUNK_HOME/etc/system/README/outputs.conf.spec for details.
        :param pulumi.Input[_builtins.str] server: host:port of the server where syslog data should be sent
        :param pulumi.Input[_builtins.str] syslog_sourcetype: Specifies a rule for handling data in addition to that provided by the "syslog" sourcetype. By default, there is no value for syslogSourceType.
               <br>This string is used as a substring match against the sourcetype key. For example, if the string is set to 'syslog', then all source types containing the string "syslog" receives this special treatment.
               To match a source type explicitly, use the pattern "sourcetype::sourcetype_name." For example
               syslogSourcetype = sourcetype::apache_common
               Data that is "syslog" or matches this setting is assumed to already be in syslog format.
               Data that does not match the rules has a header, potentially a timestamp, and a hostname added to the front of the event. This is how Splunk software causes arbitrary log data to match syslog expectations.
        :param pulumi.Input[_builtins.str] timestamp_format: Format of timestamp to add at start of the events to be forwarded.
               The format is a strftime-style timestamp formatting string. See $SPLUNK_HOME/etc/system/README/outputs.conf.spec for details.
        :param pulumi.Input[_builtins.str] type: Protocol to use to send syslog data. Valid values: (tcp | udp ).
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _OutputsTcpSyslogState.__new__(_OutputsTcpSyslogState)

        __props__.__dict__["acl"] = acl
        __props__.__dict__["disabled"] = disabled
        __props__.__dict__["name"] = name
        __props__.__dict__["priority"] = priority
        __props__.__dict__["server"] = server
        __props__.__dict__["syslog_sourcetype"] = syslog_sourcetype
        __props__.__dict__["timestamp_format"] = timestamp_format
        __props__.__dict__["type"] = type
        return OutputsTcpSyslog(resource_name, opts=opts, __props__=__props__)

    @_builtins.property
    @pulumi.getter
    def acl(self) -> pulumi.Output['outputs.OutputsTcpSyslogAcl']:
        """
        The app/user context that is the namespace for the resource
        """
        return pulumi.get(self, "acl")

    @_builtins.property
    @pulumi.getter
    def disabled(self) -> pulumi.Output[_builtins.bool]:
        """
        If true, disables global syslog settings.
        """
        return pulumi.get(self, "disabled")

    @_builtins.property
    @pulumi.getter
    def name(self) -> pulumi.Output[_builtins.str]:
        """
        Name of the syslog output group. This is name used when creating syslog configuration in outputs.conf.
        """
        return pulumi.get(self, "name")

    @_builtins.property
    @pulumi.getter
    def priority(self) -> pulumi.Output[_builtins.int]:
        """
        Sets syslog priority value. The priority value should specified as an integer. See $SPLUNK_HOME/etc/system/README/outputs.conf.spec for details.
        """
        return pulumi.get(self, "priority")

    @_builtins.property
    @pulumi.getter
    def server(self) -> pulumi.Output[_builtins.str]:
        """
        host:port of the server where syslog data should be sent
        """
        return pulumi.get(self, "server")

    @_builtins.property
    @pulumi.getter(name="syslogSourcetype")
    def syslog_sourcetype(self) -> pulumi.Output[_builtins.str]:
        """
        Specifies a rule for handling data in addition to that provided by the "syslog" sourcetype. By default, there is no value for syslogSourceType.
        <br>This string is used as a substring match against the sourcetype key. For example, if the string is set to 'syslog', then all source types containing the string "syslog" receives this special treatment.
        To match a source type explicitly, use the pattern "sourcetype::sourcetype_name." For example
        syslogSourcetype = sourcetype::apache_common
        Data that is "syslog" or matches this setting is assumed to already be in syslog format.
        Data that does not match the rules has a header, potentially a timestamp, and a hostname added to the front of the event. This is how Splunk software causes arbitrary log data to match syslog expectations.
        """
        return pulumi.get(self, "syslog_sourcetype")

    @_builtins.property
    @pulumi.getter(name="timestampFormat")
    def timestamp_format(self) -> pulumi.Output[_builtins.str]:
        """
        Format of timestamp to add at start of the events to be forwarded.
        The format is a strftime-style timestamp formatting string. See $SPLUNK_HOME/etc/system/README/outputs.conf.spec for details.
        """
        return pulumi.get(self, "timestamp_format")

    @_builtins.property
    @pulumi.getter
    def type(self) -> pulumi.Output[_builtins.str]:
        """
        Protocol to use to send syslog data. Valid values: (tcp | udp ).
        """
        return pulumi.get(self, "type")

