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

__all__ = ['InputsScriptArgs', 'InputsScript']

@pulumi.input_type
class InputsScriptArgs:
    def __init__(__self__, *,
                 interval: pulumi.Input[int],
                 acl: Optional[pulumi.Input['InputsScriptAclArgs']] = None,
                 disabled: Optional[pulumi.Input[bool]] = None,
                 host: Optional[pulumi.Input[str]] = None,
                 index: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 passauth: Optional[pulumi.Input[str]] = None,
                 rename_source: Optional[pulumi.Input[str]] = None,
                 source: Optional[pulumi.Input[str]] = None,
                 sourcetype: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a InputsScript resource.
        :param pulumi.Input[int] interval: Specify an integer or cron schedule. This parameter specifies how often to execute the specified script, in seconds or a valid cron schedule. If you specify a cron schedule, the script is not executed on start-up.
        :param pulumi.Input['InputsScriptAclArgs'] acl: The app/user context that is the namespace for the resource
        :param pulumi.Input[bool] disabled: Specifies whether the input script is disabled.
        :param pulumi.Input[str] host: Sets the host for events from this input. Defaults to whatever host sent the event.
        :param pulumi.Input[str] index: Sets the index for events from this input. Defaults to the main index.
        :param pulumi.Input[str] name: Specify the name of the scripted input.
        :param pulumi.Input[str] passauth: User to run the script as. If you provide a username, Splunk software generates an auth token for that user and passes it to the script.
        :param pulumi.Input[str] rename_source: Specify a new name for the source field for the script.
        :param pulumi.Input[str] source: Sets the source key/field for events from this input. Defaults to the input file path.
               Sets the source key initial value. The key is used during parsing/indexing, in particular to set the source field during indexing. It is also the source field used at search time. As a convenience, the chosen string is prepended with 'source::'.
        :param pulumi.Input[str] sourcetype: Sets the sourcetype key/field for events from this input. If unset, Splunk software picks a source type based on various aspects of the data. As a convenience, the chosen string is prepended with 'sourcetype::'. There is no hard-coded default.
               Sets the sourcetype key initial value. The key is used during parsing/indexing, in particular to set the source type field during indexing. It is also the source type field used at search time.
        """
        pulumi.set(__self__, "interval", interval)
        if acl is not None:
            pulumi.set(__self__, "acl", acl)
        if disabled is not None:
            pulumi.set(__self__, "disabled", disabled)
        if host is not None:
            pulumi.set(__self__, "host", host)
        if index is not None:
            pulumi.set(__self__, "index", index)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if passauth is not None:
            pulumi.set(__self__, "passauth", passauth)
        if rename_source is not None:
            pulumi.set(__self__, "rename_source", rename_source)
        if source is not None:
            pulumi.set(__self__, "source", source)
        if sourcetype is not None:
            pulumi.set(__self__, "sourcetype", sourcetype)

    @property
    @pulumi.getter
    def interval(self) -> pulumi.Input[int]:
        """
        Specify an integer or cron schedule. This parameter specifies how often to execute the specified script, in seconds or a valid cron schedule. If you specify a cron schedule, the script is not executed on start-up.
        """
        return pulumi.get(self, "interval")

    @interval.setter
    def interval(self, value: pulumi.Input[int]):
        pulumi.set(self, "interval", value)

    @property
    @pulumi.getter
    def acl(self) -> Optional[pulumi.Input['InputsScriptAclArgs']]:
        """
        The app/user context that is the namespace for the resource
        """
        return pulumi.get(self, "acl")

    @acl.setter
    def acl(self, value: Optional[pulumi.Input['InputsScriptAclArgs']]):
        pulumi.set(self, "acl", value)

    @property
    @pulumi.getter
    def disabled(self) -> Optional[pulumi.Input[bool]]:
        """
        Specifies whether the input script is disabled.
        """
        return pulumi.get(self, "disabled")

    @disabled.setter
    def disabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "disabled", value)

    @property
    @pulumi.getter
    def host(self) -> Optional[pulumi.Input[str]]:
        """
        Sets the host for events from this input. Defaults to whatever host sent the event.
        """
        return pulumi.get(self, "host")

    @host.setter
    def host(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "host", value)

    @property
    @pulumi.getter
    def index(self) -> Optional[pulumi.Input[str]]:
        """
        Sets the index for events from this input. Defaults to the main index.
        """
        return pulumi.get(self, "index")

    @index.setter
    def index(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "index", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Specify the name of the scripted input.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def passauth(self) -> Optional[pulumi.Input[str]]:
        """
        User to run the script as. If you provide a username, Splunk software generates an auth token for that user and passes it to the script.
        """
        return pulumi.get(self, "passauth")

    @passauth.setter
    def passauth(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "passauth", value)

    @property
    @pulumi.getter(name="renameSource")
    def rename_source(self) -> Optional[pulumi.Input[str]]:
        """
        Specify a new name for the source field for the script.
        """
        return pulumi.get(self, "rename_source")

    @rename_source.setter
    def rename_source(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "rename_source", value)

    @property
    @pulumi.getter
    def source(self) -> Optional[pulumi.Input[str]]:
        """
        Sets the source key/field for events from this input. Defaults to the input file path.
        Sets the source key initial value. The key is used during parsing/indexing, in particular to set the source field during indexing. It is also the source field used at search time. As a convenience, the chosen string is prepended with 'source::'.
        """
        return pulumi.get(self, "source")

    @source.setter
    def source(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "source", value)

    @property
    @pulumi.getter
    def sourcetype(self) -> Optional[pulumi.Input[str]]:
        """
        Sets the sourcetype key/field for events from this input. If unset, Splunk software picks a source type based on various aspects of the data. As a convenience, the chosen string is prepended with 'sourcetype::'. There is no hard-coded default.
        Sets the sourcetype key initial value. The key is used during parsing/indexing, in particular to set the source type field during indexing. It is also the source type field used at search time.
        """
        return pulumi.get(self, "sourcetype")

    @sourcetype.setter
    def sourcetype(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "sourcetype", value)


@pulumi.input_type
class _InputsScriptState:
    def __init__(__self__, *,
                 acl: Optional[pulumi.Input['InputsScriptAclArgs']] = None,
                 disabled: Optional[pulumi.Input[bool]] = None,
                 host: Optional[pulumi.Input[str]] = None,
                 index: Optional[pulumi.Input[str]] = None,
                 interval: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 passauth: Optional[pulumi.Input[str]] = None,
                 rename_source: Optional[pulumi.Input[str]] = None,
                 source: Optional[pulumi.Input[str]] = None,
                 sourcetype: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering InputsScript resources.
        :param pulumi.Input['InputsScriptAclArgs'] acl: The app/user context that is the namespace for the resource
        :param pulumi.Input[bool] disabled: Specifies whether the input script is disabled.
        :param pulumi.Input[str] host: Sets the host for events from this input. Defaults to whatever host sent the event.
        :param pulumi.Input[str] index: Sets the index for events from this input. Defaults to the main index.
        :param pulumi.Input[int] interval: Specify an integer or cron schedule. This parameter specifies how often to execute the specified script, in seconds or a valid cron schedule. If you specify a cron schedule, the script is not executed on start-up.
        :param pulumi.Input[str] name: Specify the name of the scripted input.
        :param pulumi.Input[str] passauth: User to run the script as. If you provide a username, Splunk software generates an auth token for that user and passes it to the script.
        :param pulumi.Input[str] rename_source: Specify a new name for the source field for the script.
        :param pulumi.Input[str] source: Sets the source key/field for events from this input. Defaults to the input file path.
               Sets the source key initial value. The key is used during parsing/indexing, in particular to set the source field during indexing. It is also the source field used at search time. As a convenience, the chosen string is prepended with 'source::'.
        :param pulumi.Input[str] sourcetype: Sets the sourcetype key/field for events from this input. If unset, Splunk software picks a source type based on various aspects of the data. As a convenience, the chosen string is prepended with 'sourcetype::'. There is no hard-coded default.
               Sets the sourcetype key initial value. The key is used during parsing/indexing, in particular to set the source type field during indexing. It is also the source type field used at search time.
        """
        if acl is not None:
            pulumi.set(__self__, "acl", acl)
        if disabled is not None:
            pulumi.set(__self__, "disabled", disabled)
        if host is not None:
            pulumi.set(__self__, "host", host)
        if index is not None:
            pulumi.set(__self__, "index", index)
        if interval is not None:
            pulumi.set(__self__, "interval", interval)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if passauth is not None:
            pulumi.set(__self__, "passauth", passauth)
        if rename_source is not None:
            pulumi.set(__self__, "rename_source", rename_source)
        if source is not None:
            pulumi.set(__self__, "source", source)
        if sourcetype is not None:
            pulumi.set(__self__, "sourcetype", sourcetype)

    @property
    @pulumi.getter
    def acl(self) -> Optional[pulumi.Input['InputsScriptAclArgs']]:
        """
        The app/user context that is the namespace for the resource
        """
        return pulumi.get(self, "acl")

    @acl.setter
    def acl(self, value: Optional[pulumi.Input['InputsScriptAclArgs']]):
        pulumi.set(self, "acl", value)

    @property
    @pulumi.getter
    def disabled(self) -> Optional[pulumi.Input[bool]]:
        """
        Specifies whether the input script is disabled.
        """
        return pulumi.get(self, "disabled")

    @disabled.setter
    def disabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "disabled", value)

    @property
    @pulumi.getter
    def host(self) -> Optional[pulumi.Input[str]]:
        """
        Sets the host for events from this input. Defaults to whatever host sent the event.
        """
        return pulumi.get(self, "host")

    @host.setter
    def host(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "host", value)

    @property
    @pulumi.getter
    def index(self) -> Optional[pulumi.Input[str]]:
        """
        Sets the index for events from this input. Defaults to the main index.
        """
        return pulumi.get(self, "index")

    @index.setter
    def index(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "index", value)

    @property
    @pulumi.getter
    def interval(self) -> Optional[pulumi.Input[int]]:
        """
        Specify an integer or cron schedule. This parameter specifies how often to execute the specified script, in seconds or a valid cron schedule. If you specify a cron schedule, the script is not executed on start-up.
        """
        return pulumi.get(self, "interval")

    @interval.setter
    def interval(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "interval", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Specify the name of the scripted input.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def passauth(self) -> Optional[pulumi.Input[str]]:
        """
        User to run the script as. If you provide a username, Splunk software generates an auth token for that user and passes it to the script.
        """
        return pulumi.get(self, "passauth")

    @passauth.setter
    def passauth(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "passauth", value)

    @property
    @pulumi.getter(name="renameSource")
    def rename_source(self) -> Optional[pulumi.Input[str]]:
        """
        Specify a new name for the source field for the script.
        """
        return pulumi.get(self, "rename_source")

    @rename_source.setter
    def rename_source(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "rename_source", value)

    @property
    @pulumi.getter
    def source(self) -> Optional[pulumi.Input[str]]:
        """
        Sets the source key/field for events from this input. Defaults to the input file path.
        Sets the source key initial value. The key is used during parsing/indexing, in particular to set the source field during indexing. It is also the source field used at search time. As a convenience, the chosen string is prepended with 'source::'.
        """
        return pulumi.get(self, "source")

    @source.setter
    def source(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "source", value)

    @property
    @pulumi.getter
    def sourcetype(self) -> Optional[pulumi.Input[str]]:
        """
        Sets the sourcetype key/field for events from this input. If unset, Splunk software picks a source type based on various aspects of the data. As a convenience, the chosen string is prepended with 'sourcetype::'. There is no hard-coded default.
        Sets the sourcetype key initial value. The key is used during parsing/indexing, in particular to set the source type field during indexing. It is also the source type field used at search time.
        """
        return pulumi.get(self, "sourcetype")

    @sourcetype.setter
    def sourcetype(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "sourcetype", value)


class InputsScript(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 acl: Optional[pulumi.Input[pulumi.InputType['InputsScriptAclArgs']]] = None,
                 disabled: Optional[pulumi.Input[bool]] = None,
                 host: Optional[pulumi.Input[str]] = None,
                 index: Optional[pulumi.Input[str]] = None,
                 interval: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 passauth: Optional[pulumi.Input[str]] = None,
                 rename_source: Optional[pulumi.Input[str]] = None,
                 source: Optional[pulumi.Input[str]] = None,
                 sourcetype: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        ## # Resource: InputsScript

        Create or update scripted inputs.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['InputsScriptAclArgs']] acl: The app/user context that is the namespace for the resource
        :param pulumi.Input[bool] disabled: Specifies whether the input script is disabled.
        :param pulumi.Input[str] host: Sets the host for events from this input. Defaults to whatever host sent the event.
        :param pulumi.Input[str] index: Sets the index for events from this input. Defaults to the main index.
        :param pulumi.Input[int] interval: Specify an integer or cron schedule. This parameter specifies how often to execute the specified script, in seconds or a valid cron schedule. If you specify a cron schedule, the script is not executed on start-up.
        :param pulumi.Input[str] name: Specify the name of the scripted input.
        :param pulumi.Input[str] passauth: User to run the script as. If you provide a username, Splunk software generates an auth token for that user and passes it to the script.
        :param pulumi.Input[str] rename_source: Specify a new name for the source field for the script.
        :param pulumi.Input[str] source: Sets the source key/field for events from this input. Defaults to the input file path.
               Sets the source key initial value. The key is used during parsing/indexing, in particular to set the source field during indexing. It is also the source field used at search time. As a convenience, the chosen string is prepended with 'source::'.
        :param pulumi.Input[str] sourcetype: Sets the sourcetype key/field for events from this input. If unset, Splunk software picks a source type based on various aspects of the data. As a convenience, the chosen string is prepended with 'sourcetype::'. There is no hard-coded default.
               Sets the sourcetype key initial value. The key is used during parsing/indexing, in particular to set the source type field during indexing. It is also the source type field used at search time.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: InputsScriptArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        ## # Resource: InputsScript

        Create or update scripted inputs.

        :param str resource_name: The name of the resource.
        :param InputsScriptArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(InputsScriptArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 acl: Optional[pulumi.Input[pulumi.InputType['InputsScriptAclArgs']]] = None,
                 disabled: Optional[pulumi.Input[bool]] = None,
                 host: Optional[pulumi.Input[str]] = None,
                 index: Optional[pulumi.Input[str]] = None,
                 interval: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 passauth: Optional[pulumi.Input[str]] = None,
                 rename_source: Optional[pulumi.Input[str]] = None,
                 source: Optional[pulumi.Input[str]] = None,
                 sourcetype: Optional[pulumi.Input[str]] = None,
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
            __props__ = InputsScriptArgs.__new__(InputsScriptArgs)

            __props__.__dict__["acl"] = acl
            __props__.__dict__["disabled"] = disabled
            __props__.__dict__["host"] = host
            __props__.__dict__["index"] = index
            if interval is None and not opts.urn:
                raise TypeError("Missing required property 'interval'")
            __props__.__dict__["interval"] = interval
            __props__.__dict__["name"] = name
            __props__.__dict__["passauth"] = passauth
            __props__.__dict__["rename_source"] = rename_source
            __props__.__dict__["source"] = source
            __props__.__dict__["sourcetype"] = sourcetype
        super(InputsScript, __self__).__init__(
            'splunk:index/inputsScript:InputsScript',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            acl: Optional[pulumi.Input[pulumi.InputType['InputsScriptAclArgs']]] = None,
            disabled: Optional[pulumi.Input[bool]] = None,
            host: Optional[pulumi.Input[str]] = None,
            index: Optional[pulumi.Input[str]] = None,
            interval: Optional[pulumi.Input[int]] = None,
            name: Optional[pulumi.Input[str]] = None,
            passauth: Optional[pulumi.Input[str]] = None,
            rename_source: Optional[pulumi.Input[str]] = None,
            source: Optional[pulumi.Input[str]] = None,
            sourcetype: Optional[pulumi.Input[str]] = None) -> 'InputsScript':
        """
        Get an existing InputsScript resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['InputsScriptAclArgs']] acl: The app/user context that is the namespace for the resource
        :param pulumi.Input[bool] disabled: Specifies whether the input script is disabled.
        :param pulumi.Input[str] host: Sets the host for events from this input. Defaults to whatever host sent the event.
        :param pulumi.Input[str] index: Sets the index for events from this input. Defaults to the main index.
        :param pulumi.Input[int] interval: Specify an integer or cron schedule. This parameter specifies how often to execute the specified script, in seconds or a valid cron schedule. If you specify a cron schedule, the script is not executed on start-up.
        :param pulumi.Input[str] name: Specify the name of the scripted input.
        :param pulumi.Input[str] passauth: User to run the script as. If you provide a username, Splunk software generates an auth token for that user and passes it to the script.
        :param pulumi.Input[str] rename_source: Specify a new name for the source field for the script.
        :param pulumi.Input[str] source: Sets the source key/field for events from this input. Defaults to the input file path.
               Sets the source key initial value. The key is used during parsing/indexing, in particular to set the source field during indexing. It is also the source field used at search time. As a convenience, the chosen string is prepended with 'source::'.
        :param pulumi.Input[str] sourcetype: Sets the sourcetype key/field for events from this input. If unset, Splunk software picks a source type based on various aspects of the data. As a convenience, the chosen string is prepended with 'sourcetype::'. There is no hard-coded default.
               Sets the sourcetype key initial value. The key is used during parsing/indexing, in particular to set the source type field during indexing. It is also the source type field used at search time.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _InputsScriptState.__new__(_InputsScriptState)

        __props__.__dict__["acl"] = acl
        __props__.__dict__["disabled"] = disabled
        __props__.__dict__["host"] = host
        __props__.__dict__["index"] = index
        __props__.__dict__["interval"] = interval
        __props__.__dict__["name"] = name
        __props__.__dict__["passauth"] = passauth
        __props__.__dict__["rename_source"] = rename_source
        __props__.__dict__["source"] = source
        __props__.__dict__["sourcetype"] = sourcetype
        return InputsScript(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def acl(self) -> pulumi.Output['outputs.InputsScriptAcl']:
        """
        The app/user context that is the namespace for the resource
        """
        return pulumi.get(self, "acl")

    @property
    @pulumi.getter
    def disabled(self) -> pulumi.Output[bool]:
        """
        Specifies whether the input script is disabled.
        """
        return pulumi.get(self, "disabled")

    @property
    @pulumi.getter
    def host(self) -> pulumi.Output[str]:
        """
        Sets the host for events from this input. Defaults to whatever host sent the event.
        """
        return pulumi.get(self, "host")

    @property
    @pulumi.getter
    def index(self) -> pulumi.Output[str]:
        """
        Sets the index for events from this input. Defaults to the main index.
        """
        return pulumi.get(self, "index")

    @property
    @pulumi.getter
    def interval(self) -> pulumi.Output[int]:
        """
        Specify an integer or cron schedule. This parameter specifies how often to execute the specified script, in seconds or a valid cron schedule. If you specify a cron schedule, the script is not executed on start-up.
        """
        return pulumi.get(self, "interval")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Specify the name of the scripted input.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def passauth(self) -> pulumi.Output[str]:
        """
        User to run the script as. If you provide a username, Splunk software generates an auth token for that user and passes it to the script.
        """
        return pulumi.get(self, "passauth")

    @property
    @pulumi.getter(name="renameSource")
    def rename_source(self) -> pulumi.Output[str]:
        """
        Specify a new name for the source field for the script.
        """
        return pulumi.get(self, "rename_source")

    @property
    @pulumi.getter
    def source(self) -> pulumi.Output[str]:
        """
        Sets the source key/field for events from this input. Defaults to the input file path.
        Sets the source key initial value. The key is used during parsing/indexing, in particular to set the source field during indexing. It is also the source field used at search time. As a convenience, the chosen string is prepended with 'source::'.
        """
        return pulumi.get(self, "source")

    @property
    @pulumi.getter
    def sourcetype(self) -> pulumi.Output[str]:
        """
        Sets the sourcetype key/field for events from this input. If unset, Splunk software picks a source type based on various aspects of the data. As a convenience, the chosen string is prepended with 'sourcetype::'. There is no hard-coded default.
        Sets the sourcetype key initial value. The key is used during parsing/indexing, in particular to set the source type field during indexing. It is also the source type field used at search time.
        """
        return pulumi.get(self, "sourcetype")

