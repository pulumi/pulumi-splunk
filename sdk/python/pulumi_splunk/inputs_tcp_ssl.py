# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['InputsTcpSslArgs', 'InputsTcpSsl']

@pulumi.input_type
class InputsTcpSslArgs:
    def __init__(__self__, *,
                 disabled: Optional[pulumi.Input[bool]] = None,
                 password: Optional[pulumi.Input[str]] = None,
                 require_client_cert: Optional[pulumi.Input[bool]] = None,
                 root_ca: Optional[pulumi.Input[str]] = None,
                 server_cert: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a InputsTcpSsl resource.
        :param pulumi.Input[bool] disabled: Indicates if input is disabled.
        :param pulumi.Input[str] password: Server certificate password, if any.
        :param pulumi.Input[bool] require_client_cert: Determines whether a client must authenticate.
        :param pulumi.Input[str] root_ca: Certificate authority list (root file)
        :param pulumi.Input[str] server_cert: Full path to the server certificate.
        """
        if disabled is not None:
            pulumi.set(__self__, "disabled", disabled)
        if password is not None:
            pulumi.set(__self__, "password", password)
        if require_client_cert is not None:
            pulumi.set(__self__, "require_client_cert", require_client_cert)
        if root_ca is not None:
            pulumi.set(__self__, "root_ca", root_ca)
        if server_cert is not None:
            pulumi.set(__self__, "server_cert", server_cert)

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
    def password(self) -> Optional[pulumi.Input[str]]:
        """
        Server certificate password, if any.
        """
        return pulumi.get(self, "password")

    @password.setter
    def password(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "password", value)

    @property
    @pulumi.getter(name="requireClientCert")
    def require_client_cert(self) -> Optional[pulumi.Input[bool]]:
        """
        Determines whether a client must authenticate.
        """
        return pulumi.get(self, "require_client_cert")

    @require_client_cert.setter
    def require_client_cert(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "require_client_cert", value)

    @property
    @pulumi.getter(name="rootCa")
    def root_ca(self) -> Optional[pulumi.Input[str]]:
        """
        Certificate authority list (root file)
        """
        return pulumi.get(self, "root_ca")

    @root_ca.setter
    def root_ca(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "root_ca", value)

    @property
    @pulumi.getter(name="serverCert")
    def server_cert(self) -> Optional[pulumi.Input[str]]:
        """
        Full path to the server certificate.
        """
        return pulumi.get(self, "server_cert")

    @server_cert.setter
    def server_cert(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "server_cert", value)


@pulumi.input_type
class _InputsTcpSslState:
    def __init__(__self__, *,
                 disabled: Optional[pulumi.Input[bool]] = None,
                 password: Optional[pulumi.Input[str]] = None,
                 require_client_cert: Optional[pulumi.Input[bool]] = None,
                 root_ca: Optional[pulumi.Input[str]] = None,
                 server_cert: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering InputsTcpSsl resources.
        :param pulumi.Input[bool] disabled: Indicates if input is disabled.
        :param pulumi.Input[str] password: Server certificate password, if any.
        :param pulumi.Input[bool] require_client_cert: Determines whether a client must authenticate.
        :param pulumi.Input[str] root_ca: Certificate authority list (root file)
        :param pulumi.Input[str] server_cert: Full path to the server certificate.
        """
        if disabled is not None:
            pulumi.set(__self__, "disabled", disabled)
        if password is not None:
            pulumi.set(__self__, "password", password)
        if require_client_cert is not None:
            pulumi.set(__self__, "require_client_cert", require_client_cert)
        if root_ca is not None:
            pulumi.set(__self__, "root_ca", root_ca)
        if server_cert is not None:
            pulumi.set(__self__, "server_cert", server_cert)

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
    def password(self) -> Optional[pulumi.Input[str]]:
        """
        Server certificate password, if any.
        """
        return pulumi.get(self, "password")

    @password.setter
    def password(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "password", value)

    @property
    @pulumi.getter(name="requireClientCert")
    def require_client_cert(self) -> Optional[pulumi.Input[bool]]:
        """
        Determines whether a client must authenticate.
        """
        return pulumi.get(self, "require_client_cert")

    @require_client_cert.setter
    def require_client_cert(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "require_client_cert", value)

    @property
    @pulumi.getter(name="rootCa")
    def root_ca(self) -> Optional[pulumi.Input[str]]:
        """
        Certificate authority list (root file)
        """
        return pulumi.get(self, "root_ca")

    @root_ca.setter
    def root_ca(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "root_ca", value)

    @property
    @pulumi.getter(name="serverCert")
    def server_cert(self) -> Optional[pulumi.Input[str]]:
        """
        Full path to the server certificate.
        """
        return pulumi.get(self, "server_cert")

    @server_cert.setter
    def server_cert(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "server_cert", value)


class InputsTcpSsl(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 disabled: Optional[pulumi.Input[bool]] = None,
                 password: Optional[pulumi.Input[str]] = None,
                 require_client_cert: Optional[pulumi.Input[bool]] = None,
                 root_ca: Optional[pulumi.Input[str]] = None,
                 server_cert: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        ## # Resource: InputsTcpSsl

        Access or update the SSL configuration for the host.

        ## Example Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import pulumi_splunk as splunk

        test = splunk.InputsTcpSsl("test",
            disabled=False,
            require_client_cert=True)
        ```
        <!--End PulumiCodeChooser -->

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] disabled: Indicates if input is disabled.
        :param pulumi.Input[str] password: Server certificate password, if any.
        :param pulumi.Input[bool] require_client_cert: Determines whether a client must authenticate.
        :param pulumi.Input[str] root_ca: Certificate authority list (root file)
        :param pulumi.Input[str] server_cert: Full path to the server certificate.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[InputsTcpSslArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        ## # Resource: InputsTcpSsl

        Access or update the SSL configuration for the host.

        ## Example Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import pulumi_splunk as splunk

        test = splunk.InputsTcpSsl("test",
            disabled=False,
            require_client_cert=True)
        ```
        <!--End PulumiCodeChooser -->

        :param str resource_name: The name of the resource.
        :param InputsTcpSslArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(InputsTcpSslArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 disabled: Optional[pulumi.Input[bool]] = None,
                 password: Optional[pulumi.Input[str]] = None,
                 require_client_cert: Optional[pulumi.Input[bool]] = None,
                 root_ca: Optional[pulumi.Input[str]] = None,
                 server_cert: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = InputsTcpSslArgs.__new__(InputsTcpSslArgs)

            __props__.__dict__["disabled"] = disabled
            __props__.__dict__["password"] = None if password is None else pulumi.Output.secret(password)
            __props__.__dict__["require_client_cert"] = require_client_cert
            __props__.__dict__["root_ca"] = root_ca
            __props__.__dict__["server_cert"] = server_cert
        secret_opts = pulumi.ResourceOptions(additional_secret_outputs=["password"])
        opts = pulumi.ResourceOptions.merge(opts, secret_opts)
        super(InputsTcpSsl, __self__).__init__(
            'splunk:index/inputsTcpSsl:InputsTcpSsl',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            disabled: Optional[pulumi.Input[bool]] = None,
            password: Optional[pulumi.Input[str]] = None,
            require_client_cert: Optional[pulumi.Input[bool]] = None,
            root_ca: Optional[pulumi.Input[str]] = None,
            server_cert: Optional[pulumi.Input[str]] = None) -> 'InputsTcpSsl':
        """
        Get an existing InputsTcpSsl resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] disabled: Indicates if input is disabled.
        :param pulumi.Input[str] password: Server certificate password, if any.
        :param pulumi.Input[bool] require_client_cert: Determines whether a client must authenticate.
        :param pulumi.Input[str] root_ca: Certificate authority list (root file)
        :param pulumi.Input[str] server_cert: Full path to the server certificate.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _InputsTcpSslState.__new__(_InputsTcpSslState)

        __props__.__dict__["disabled"] = disabled
        __props__.__dict__["password"] = password
        __props__.__dict__["require_client_cert"] = require_client_cert
        __props__.__dict__["root_ca"] = root_ca
        __props__.__dict__["server_cert"] = server_cert
        return InputsTcpSsl(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def disabled(self) -> pulumi.Output[bool]:
        """
        Indicates if input is disabled.
        """
        return pulumi.get(self, "disabled")

    @property
    @pulumi.getter
    def password(self) -> pulumi.Output[str]:
        """
        Server certificate password, if any.
        """
        return pulumi.get(self, "password")

    @property
    @pulumi.getter(name="requireClientCert")
    def require_client_cert(self) -> pulumi.Output[bool]:
        """
        Determines whether a client must authenticate.
        """
        return pulumi.get(self, "require_client_cert")

    @property
    @pulumi.getter(name="rootCa")
    def root_ca(self) -> pulumi.Output[str]:
        """
        Certificate authority list (root file)
        """
        return pulumi.get(self, "root_ca")

    @property
    @pulumi.getter(name="serverCert")
    def server_cert(self) -> pulumi.Output[str]:
        """
        Full path to the server certificate.
        """
        return pulumi.get(self, "server_cert")

