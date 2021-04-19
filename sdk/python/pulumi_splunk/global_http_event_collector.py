# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['GlobalHttpEventCollectorArgs', 'GlobalHttpEventCollector']

@pulumi.input_type
class GlobalHttpEventCollectorArgs:
    def __init__(__self__, *,
                 dedicated_io_threads: Optional[pulumi.Input[int]] = None,
                 disabled: Optional[pulumi.Input[bool]] = None,
                 enable_ssl: Optional[pulumi.Input[bool]] = None,
                 max_sockets: Optional[pulumi.Input[int]] = None,
                 max_threads: Optional[pulumi.Input[int]] = None,
                 port: Optional[pulumi.Input[int]] = None,
                 use_deployment_server: Optional[pulumi.Input[int]] = None):
        """
        The set of arguments for constructing a GlobalHttpEventCollector resource.
        :param pulumi.Input[int] dedicated_io_threads: Number of threads used by HTTP Input server.
        :param pulumi.Input[bool] disabled: Input disabled indicator.
        :param pulumi.Input[bool] enable_ssl: Enable SSL protocol for HTTP data input. `true` = SSL enabled, `false` = SSL disabled.
        :param pulumi.Input[int] max_sockets: Maximum number of simultaneous HTTP connections accepted. Adjusting this value may cause server performance issues and is not generally recommended. Possible values for this setting vary by OS.
        :param pulumi.Input[int] max_threads: Maximum number of threads that can be used by active HTTP transactions. Adjusting this value may cause server performance issues and is not generally recommended. Possible values for this setting vary by OS.
        :param pulumi.Input[int] port: HTTP data input IP port.
        :param pulumi.Input[int] use_deployment_server: Indicates whether the event collector input writes its configuration to a deployment server repository. When this setting is set to 1 (enabled), the input writes its configuration to the directory specified as repositoryLocation in serverclass.conf.
               Copy the full contents of the splunk_httpinput app directory to this directory for the configuration to work. When enabled, only the tokens defined in the splunk_httpinput app in this repository are viewable and editable on the API and the Data Inputs page in Splunk Web. When disabled, the input writes its configuration to $SPLUNK_HOME/etc/apps by default. Defaults to 0 (disabled).
        """
        if dedicated_io_threads is not None:
            pulumi.set(__self__, "dedicated_io_threads", dedicated_io_threads)
        if disabled is not None:
            pulumi.set(__self__, "disabled", disabled)
        if enable_ssl is not None:
            pulumi.set(__self__, "enable_ssl", enable_ssl)
        if max_sockets is not None:
            pulumi.set(__self__, "max_sockets", max_sockets)
        if max_threads is not None:
            pulumi.set(__self__, "max_threads", max_threads)
        if port is not None:
            pulumi.set(__self__, "port", port)
        if use_deployment_server is not None:
            pulumi.set(__self__, "use_deployment_server", use_deployment_server)

    @property
    @pulumi.getter(name="dedicatedIoThreads")
    def dedicated_io_threads(self) -> Optional[pulumi.Input[int]]:
        """
        Number of threads used by HTTP Input server.
        """
        return pulumi.get(self, "dedicated_io_threads")

    @dedicated_io_threads.setter
    def dedicated_io_threads(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "dedicated_io_threads", value)

    @property
    @pulumi.getter
    def disabled(self) -> Optional[pulumi.Input[bool]]:
        """
        Input disabled indicator.
        """
        return pulumi.get(self, "disabled")

    @disabled.setter
    def disabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "disabled", value)

    @property
    @pulumi.getter(name="enableSsl")
    def enable_ssl(self) -> Optional[pulumi.Input[bool]]:
        """
        Enable SSL protocol for HTTP data input. `true` = SSL enabled, `false` = SSL disabled.
        """
        return pulumi.get(self, "enable_ssl")

    @enable_ssl.setter
    def enable_ssl(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enable_ssl", value)

    @property
    @pulumi.getter(name="maxSockets")
    def max_sockets(self) -> Optional[pulumi.Input[int]]:
        """
        Maximum number of simultaneous HTTP connections accepted. Adjusting this value may cause server performance issues and is not generally recommended. Possible values for this setting vary by OS.
        """
        return pulumi.get(self, "max_sockets")

    @max_sockets.setter
    def max_sockets(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "max_sockets", value)

    @property
    @pulumi.getter(name="maxThreads")
    def max_threads(self) -> Optional[pulumi.Input[int]]:
        """
        Maximum number of threads that can be used by active HTTP transactions. Adjusting this value may cause server performance issues and is not generally recommended. Possible values for this setting vary by OS.
        """
        return pulumi.get(self, "max_threads")

    @max_threads.setter
    def max_threads(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "max_threads", value)

    @property
    @pulumi.getter
    def port(self) -> Optional[pulumi.Input[int]]:
        """
        HTTP data input IP port.
        """
        return pulumi.get(self, "port")

    @port.setter
    def port(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "port", value)

    @property
    @pulumi.getter(name="useDeploymentServer")
    def use_deployment_server(self) -> Optional[pulumi.Input[int]]:
        """
        Indicates whether the event collector input writes its configuration to a deployment server repository. When this setting is set to 1 (enabled), the input writes its configuration to the directory specified as repositoryLocation in serverclass.conf.
        Copy the full contents of the splunk_httpinput app directory to this directory for the configuration to work. When enabled, only the tokens defined in the splunk_httpinput app in this repository are viewable and editable on the API and the Data Inputs page in Splunk Web. When disabled, the input writes its configuration to $SPLUNK_HOME/etc/apps by default. Defaults to 0 (disabled).
        """
        return pulumi.get(self, "use_deployment_server")

    @use_deployment_server.setter
    def use_deployment_server(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "use_deployment_server", value)


@pulumi.input_type
class _GlobalHttpEventCollectorState:
    def __init__(__self__, *,
                 dedicated_io_threads: Optional[pulumi.Input[int]] = None,
                 disabled: Optional[pulumi.Input[bool]] = None,
                 enable_ssl: Optional[pulumi.Input[bool]] = None,
                 max_sockets: Optional[pulumi.Input[int]] = None,
                 max_threads: Optional[pulumi.Input[int]] = None,
                 port: Optional[pulumi.Input[int]] = None,
                 use_deployment_server: Optional[pulumi.Input[int]] = None):
        """
        Input properties used for looking up and filtering GlobalHttpEventCollector resources.
        :param pulumi.Input[int] dedicated_io_threads: Number of threads used by HTTP Input server.
        :param pulumi.Input[bool] disabled: Input disabled indicator.
        :param pulumi.Input[bool] enable_ssl: Enable SSL protocol for HTTP data input. `true` = SSL enabled, `false` = SSL disabled.
        :param pulumi.Input[int] max_sockets: Maximum number of simultaneous HTTP connections accepted. Adjusting this value may cause server performance issues and is not generally recommended. Possible values for this setting vary by OS.
        :param pulumi.Input[int] max_threads: Maximum number of threads that can be used by active HTTP transactions. Adjusting this value may cause server performance issues and is not generally recommended. Possible values for this setting vary by OS.
        :param pulumi.Input[int] port: HTTP data input IP port.
        :param pulumi.Input[int] use_deployment_server: Indicates whether the event collector input writes its configuration to a deployment server repository. When this setting is set to 1 (enabled), the input writes its configuration to the directory specified as repositoryLocation in serverclass.conf.
               Copy the full contents of the splunk_httpinput app directory to this directory for the configuration to work. When enabled, only the tokens defined in the splunk_httpinput app in this repository are viewable and editable on the API and the Data Inputs page in Splunk Web. When disabled, the input writes its configuration to $SPLUNK_HOME/etc/apps by default. Defaults to 0 (disabled).
        """
        if dedicated_io_threads is not None:
            pulumi.set(__self__, "dedicated_io_threads", dedicated_io_threads)
        if disabled is not None:
            pulumi.set(__self__, "disabled", disabled)
        if enable_ssl is not None:
            pulumi.set(__self__, "enable_ssl", enable_ssl)
        if max_sockets is not None:
            pulumi.set(__self__, "max_sockets", max_sockets)
        if max_threads is not None:
            pulumi.set(__self__, "max_threads", max_threads)
        if port is not None:
            pulumi.set(__self__, "port", port)
        if use_deployment_server is not None:
            pulumi.set(__self__, "use_deployment_server", use_deployment_server)

    @property
    @pulumi.getter(name="dedicatedIoThreads")
    def dedicated_io_threads(self) -> Optional[pulumi.Input[int]]:
        """
        Number of threads used by HTTP Input server.
        """
        return pulumi.get(self, "dedicated_io_threads")

    @dedicated_io_threads.setter
    def dedicated_io_threads(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "dedicated_io_threads", value)

    @property
    @pulumi.getter
    def disabled(self) -> Optional[pulumi.Input[bool]]:
        """
        Input disabled indicator.
        """
        return pulumi.get(self, "disabled")

    @disabled.setter
    def disabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "disabled", value)

    @property
    @pulumi.getter(name="enableSsl")
    def enable_ssl(self) -> Optional[pulumi.Input[bool]]:
        """
        Enable SSL protocol for HTTP data input. `true` = SSL enabled, `false` = SSL disabled.
        """
        return pulumi.get(self, "enable_ssl")

    @enable_ssl.setter
    def enable_ssl(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enable_ssl", value)

    @property
    @pulumi.getter(name="maxSockets")
    def max_sockets(self) -> Optional[pulumi.Input[int]]:
        """
        Maximum number of simultaneous HTTP connections accepted. Adjusting this value may cause server performance issues and is not generally recommended. Possible values for this setting vary by OS.
        """
        return pulumi.get(self, "max_sockets")

    @max_sockets.setter
    def max_sockets(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "max_sockets", value)

    @property
    @pulumi.getter(name="maxThreads")
    def max_threads(self) -> Optional[pulumi.Input[int]]:
        """
        Maximum number of threads that can be used by active HTTP transactions. Adjusting this value may cause server performance issues and is not generally recommended. Possible values for this setting vary by OS.
        """
        return pulumi.get(self, "max_threads")

    @max_threads.setter
    def max_threads(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "max_threads", value)

    @property
    @pulumi.getter
    def port(self) -> Optional[pulumi.Input[int]]:
        """
        HTTP data input IP port.
        """
        return pulumi.get(self, "port")

    @port.setter
    def port(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "port", value)

    @property
    @pulumi.getter(name="useDeploymentServer")
    def use_deployment_server(self) -> Optional[pulumi.Input[int]]:
        """
        Indicates whether the event collector input writes its configuration to a deployment server repository. When this setting is set to 1 (enabled), the input writes its configuration to the directory specified as repositoryLocation in serverclass.conf.
        Copy the full contents of the splunk_httpinput app directory to this directory for the configuration to work. When enabled, only the tokens defined in the splunk_httpinput app in this repository are viewable and editable on the API and the Data Inputs page in Splunk Web. When disabled, the input writes its configuration to $SPLUNK_HOME/etc/apps by default. Defaults to 0 (disabled).
        """
        return pulumi.get(self, "use_deployment_server")

    @use_deployment_server.setter
    def use_deployment_server(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "use_deployment_server", value)


class GlobalHttpEventCollector(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 dedicated_io_threads: Optional[pulumi.Input[int]] = None,
                 disabled: Optional[pulumi.Input[bool]] = None,
                 enable_ssl: Optional[pulumi.Input[bool]] = None,
                 max_sockets: Optional[pulumi.Input[int]] = None,
                 max_threads: Optional[pulumi.Input[int]] = None,
                 port: Optional[pulumi.Input[int]] = None,
                 use_deployment_server: Optional[pulumi.Input[int]] = None,
                 __props__=None):
        """
        ## # Resource: GlobalHttpEventCollector

        Update Global HTTP Event Collector input configuration.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_splunk as splunk

        http = splunk.GlobalHttpEventCollector("http",
            disabled=False,
            enable_ssl=True,
            port=8088)
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[int] dedicated_io_threads: Number of threads used by HTTP Input server.
        :param pulumi.Input[bool] disabled: Input disabled indicator.
        :param pulumi.Input[bool] enable_ssl: Enable SSL protocol for HTTP data input. `true` = SSL enabled, `false` = SSL disabled.
        :param pulumi.Input[int] max_sockets: Maximum number of simultaneous HTTP connections accepted. Adjusting this value may cause server performance issues and is not generally recommended. Possible values for this setting vary by OS.
        :param pulumi.Input[int] max_threads: Maximum number of threads that can be used by active HTTP transactions. Adjusting this value may cause server performance issues and is not generally recommended. Possible values for this setting vary by OS.
        :param pulumi.Input[int] port: HTTP data input IP port.
        :param pulumi.Input[int] use_deployment_server: Indicates whether the event collector input writes its configuration to a deployment server repository. When this setting is set to 1 (enabled), the input writes its configuration to the directory specified as repositoryLocation in serverclass.conf.
               Copy the full contents of the splunk_httpinput app directory to this directory for the configuration to work. When enabled, only the tokens defined in the splunk_httpinput app in this repository are viewable and editable on the API and the Data Inputs page in Splunk Web. When disabled, the input writes its configuration to $SPLUNK_HOME/etc/apps by default. Defaults to 0 (disabled).
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[GlobalHttpEventCollectorArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        ## # Resource: GlobalHttpEventCollector

        Update Global HTTP Event Collector input configuration.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_splunk as splunk

        http = splunk.GlobalHttpEventCollector("http",
            disabled=False,
            enable_ssl=True,
            port=8088)
        ```

        :param str resource_name: The name of the resource.
        :param GlobalHttpEventCollectorArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(GlobalHttpEventCollectorArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 dedicated_io_threads: Optional[pulumi.Input[int]] = None,
                 disabled: Optional[pulumi.Input[bool]] = None,
                 enable_ssl: Optional[pulumi.Input[bool]] = None,
                 max_sockets: Optional[pulumi.Input[int]] = None,
                 max_threads: Optional[pulumi.Input[int]] = None,
                 port: Optional[pulumi.Input[int]] = None,
                 use_deployment_server: Optional[pulumi.Input[int]] = None,
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
            __props__ = GlobalHttpEventCollectorArgs.__new__(GlobalHttpEventCollectorArgs)

            __props__.__dict__["dedicated_io_threads"] = dedicated_io_threads
            __props__.__dict__["disabled"] = disabled
            __props__.__dict__["enable_ssl"] = enable_ssl
            __props__.__dict__["max_sockets"] = max_sockets
            __props__.__dict__["max_threads"] = max_threads
            __props__.__dict__["port"] = port
            __props__.__dict__["use_deployment_server"] = use_deployment_server
        super(GlobalHttpEventCollector, __self__).__init__(
            'splunk:index/globalHttpEventCollector:GlobalHttpEventCollector',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            dedicated_io_threads: Optional[pulumi.Input[int]] = None,
            disabled: Optional[pulumi.Input[bool]] = None,
            enable_ssl: Optional[pulumi.Input[bool]] = None,
            max_sockets: Optional[pulumi.Input[int]] = None,
            max_threads: Optional[pulumi.Input[int]] = None,
            port: Optional[pulumi.Input[int]] = None,
            use_deployment_server: Optional[pulumi.Input[int]] = None) -> 'GlobalHttpEventCollector':
        """
        Get an existing GlobalHttpEventCollector resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[int] dedicated_io_threads: Number of threads used by HTTP Input server.
        :param pulumi.Input[bool] disabled: Input disabled indicator.
        :param pulumi.Input[bool] enable_ssl: Enable SSL protocol for HTTP data input. `true` = SSL enabled, `false` = SSL disabled.
        :param pulumi.Input[int] max_sockets: Maximum number of simultaneous HTTP connections accepted. Adjusting this value may cause server performance issues and is not generally recommended. Possible values for this setting vary by OS.
        :param pulumi.Input[int] max_threads: Maximum number of threads that can be used by active HTTP transactions. Adjusting this value may cause server performance issues and is not generally recommended. Possible values for this setting vary by OS.
        :param pulumi.Input[int] port: HTTP data input IP port.
        :param pulumi.Input[int] use_deployment_server: Indicates whether the event collector input writes its configuration to a deployment server repository. When this setting is set to 1 (enabled), the input writes its configuration to the directory specified as repositoryLocation in serverclass.conf.
               Copy the full contents of the splunk_httpinput app directory to this directory for the configuration to work. When enabled, only the tokens defined in the splunk_httpinput app in this repository are viewable and editable on the API and the Data Inputs page in Splunk Web. When disabled, the input writes its configuration to $SPLUNK_HOME/etc/apps by default. Defaults to 0 (disabled).
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _GlobalHttpEventCollectorState.__new__(_GlobalHttpEventCollectorState)

        __props__.__dict__["dedicated_io_threads"] = dedicated_io_threads
        __props__.__dict__["disabled"] = disabled
        __props__.__dict__["enable_ssl"] = enable_ssl
        __props__.__dict__["max_sockets"] = max_sockets
        __props__.__dict__["max_threads"] = max_threads
        __props__.__dict__["port"] = port
        __props__.__dict__["use_deployment_server"] = use_deployment_server
        return GlobalHttpEventCollector(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="dedicatedIoThreads")
    def dedicated_io_threads(self) -> pulumi.Output[int]:
        """
        Number of threads used by HTTP Input server.
        """
        return pulumi.get(self, "dedicated_io_threads")

    @property
    @pulumi.getter
    def disabled(self) -> pulumi.Output[bool]:
        """
        Input disabled indicator.
        """
        return pulumi.get(self, "disabled")

    @property
    @pulumi.getter(name="enableSsl")
    def enable_ssl(self) -> pulumi.Output[bool]:
        """
        Enable SSL protocol for HTTP data input. `true` = SSL enabled, `false` = SSL disabled.
        """
        return pulumi.get(self, "enable_ssl")

    @property
    @pulumi.getter(name="maxSockets")
    def max_sockets(self) -> pulumi.Output[int]:
        """
        Maximum number of simultaneous HTTP connections accepted. Adjusting this value may cause server performance issues and is not generally recommended. Possible values for this setting vary by OS.
        """
        return pulumi.get(self, "max_sockets")

    @property
    @pulumi.getter(name="maxThreads")
    def max_threads(self) -> pulumi.Output[int]:
        """
        Maximum number of threads that can be used by active HTTP transactions. Adjusting this value may cause server performance issues and is not generally recommended. Possible values for this setting vary by OS.
        """
        return pulumi.get(self, "max_threads")

    @property
    @pulumi.getter
    def port(self) -> pulumi.Output[int]:
        """
        HTTP data input IP port.
        """
        return pulumi.get(self, "port")

    @property
    @pulumi.getter(name="useDeploymentServer")
    def use_deployment_server(self) -> pulumi.Output[int]:
        """
        Indicates whether the event collector input writes its configuration to a deployment server repository. When this setting is set to 1 (enabled), the input writes its configuration to the directory specified as repositoryLocation in serverclass.conf.
        Copy the full contents of the splunk_httpinput app directory to this directory for the configuration to work. When enabled, only the tokens defined in the splunk_httpinput app in this repository are viewable and editable on the API and the Data Inputs page in Splunk Web. When disabled, the input writes its configuration to $SPLUNK_HOME/etc/apps by default. Defaults to 0 (disabled).
        """
        return pulumi.get(self, "use_deployment_server")

