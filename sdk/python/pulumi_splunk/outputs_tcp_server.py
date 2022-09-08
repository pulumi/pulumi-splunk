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

__all__ = ['OutputsTcpServerArgs', 'OutputsTcpServer']

@pulumi.input_type
class OutputsTcpServerArgs:
    def __init__(__self__, *,
                 acl: Optional[pulumi.Input['OutputsTcpServerAclArgs']] = None,
                 disabled: Optional[pulumi.Input[bool]] = None,
                 method: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 ssl_alt_name_to_check: Optional[pulumi.Input[str]] = None,
                 ssl_cert_path: Optional[pulumi.Input[str]] = None,
                 ssl_cipher: Optional[pulumi.Input[str]] = None,
                 ssl_common_name_to_check: Optional[pulumi.Input[str]] = None,
                 ssl_password: Optional[pulumi.Input[str]] = None,
                 ssl_root_ca_path: Optional[pulumi.Input[str]] = None,
                 ssl_verify_server_cert: Optional[pulumi.Input[bool]] = None):
        """
        The set of arguments for constructing a OutputsTcpServer resource.
        :param pulumi.Input['OutputsTcpServerAclArgs'] acl: The app/user context that is the namespace for the resource
        :param pulumi.Input[bool] disabled: If true, disables the group.
        :param pulumi.Input[str] method: Valid values: (clone | balance | autobalance)
               The data distribution method used when two or more servers exist in the same forwarder group.
        :param pulumi.Input[str] name: <host>:<port> of the Splunk receiver. <host> can be either an ip address or server name. <port> is the that port that the Splunk receiver is listening on.
        :param pulumi.Input[str] ssl_alt_name_to_check: The alternate name to match in the remote server's SSL certificate.
        :param pulumi.Input[str] ssl_cert_path: Path to the client certificate. If specified, connection uses SSL.
        :param pulumi.Input[str] ssl_cipher: SSL Cipher in the form ALL:!aNULL:!eNULL:!LOW:!EXP:RC4+RSA:+HIGH:+MEDIUM
        :param pulumi.Input[str] ssl_common_name_to_check: Check the common name of the server's certificate against this name.
               If there is no match, assume that Splunk Enterprise is not authenticated against this server. You must specify this setting if sslVerifyServerCert is true.
        :param pulumi.Input[str] ssl_password: The password associated with the CAcert.
               The default Splunk Enterprise CAcert uses the password "password."
        :param pulumi.Input[str] ssl_root_ca_path: The path to the root certificate authority file.
        :param pulumi.Input[bool] ssl_verify_server_cert: If true, make sure that the server you are connecting to is a valid one (authenticated). Both the common name and the alternate name of the server are then checked for a match.
        """
        if acl is not None:
            pulumi.set(__self__, "acl", acl)
        if disabled is not None:
            pulumi.set(__self__, "disabled", disabled)
        if method is not None:
            pulumi.set(__self__, "method", method)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if ssl_alt_name_to_check is not None:
            pulumi.set(__self__, "ssl_alt_name_to_check", ssl_alt_name_to_check)
        if ssl_cert_path is not None:
            pulumi.set(__self__, "ssl_cert_path", ssl_cert_path)
        if ssl_cipher is not None:
            pulumi.set(__self__, "ssl_cipher", ssl_cipher)
        if ssl_common_name_to_check is not None:
            pulumi.set(__self__, "ssl_common_name_to_check", ssl_common_name_to_check)
        if ssl_password is not None:
            pulumi.set(__self__, "ssl_password", ssl_password)
        if ssl_root_ca_path is not None:
            pulumi.set(__self__, "ssl_root_ca_path", ssl_root_ca_path)
        if ssl_verify_server_cert is not None:
            pulumi.set(__self__, "ssl_verify_server_cert", ssl_verify_server_cert)

    @property
    @pulumi.getter
    def acl(self) -> Optional[pulumi.Input['OutputsTcpServerAclArgs']]:
        """
        The app/user context that is the namespace for the resource
        """
        return pulumi.get(self, "acl")

    @acl.setter
    def acl(self, value: Optional[pulumi.Input['OutputsTcpServerAclArgs']]):
        pulumi.set(self, "acl", value)

    @property
    @pulumi.getter
    def disabled(self) -> Optional[pulumi.Input[bool]]:
        """
        If true, disables the group.
        """
        return pulumi.get(self, "disabled")

    @disabled.setter
    def disabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "disabled", value)

    @property
    @pulumi.getter
    def method(self) -> Optional[pulumi.Input[str]]:
        """
        Valid values: (clone | balance | autobalance)
        The data distribution method used when two or more servers exist in the same forwarder group.
        """
        return pulumi.get(self, "method")

    @method.setter
    def method(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "method", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        <host>:<port> of the Splunk receiver. <host> can be either an ip address or server name. <port> is the that port that the Splunk receiver is listening on.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="sslAltNameToCheck")
    def ssl_alt_name_to_check(self) -> Optional[pulumi.Input[str]]:
        """
        The alternate name to match in the remote server's SSL certificate.
        """
        return pulumi.get(self, "ssl_alt_name_to_check")

    @ssl_alt_name_to_check.setter
    def ssl_alt_name_to_check(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ssl_alt_name_to_check", value)

    @property
    @pulumi.getter(name="sslCertPath")
    def ssl_cert_path(self) -> Optional[pulumi.Input[str]]:
        """
        Path to the client certificate. If specified, connection uses SSL.
        """
        return pulumi.get(self, "ssl_cert_path")

    @ssl_cert_path.setter
    def ssl_cert_path(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ssl_cert_path", value)

    @property
    @pulumi.getter(name="sslCipher")
    def ssl_cipher(self) -> Optional[pulumi.Input[str]]:
        """
        SSL Cipher in the form ALL:!aNULL:!eNULL:!LOW:!EXP:RC4+RSA:+HIGH:+MEDIUM
        """
        return pulumi.get(self, "ssl_cipher")

    @ssl_cipher.setter
    def ssl_cipher(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ssl_cipher", value)

    @property
    @pulumi.getter(name="sslCommonNameToCheck")
    def ssl_common_name_to_check(self) -> Optional[pulumi.Input[str]]:
        """
        Check the common name of the server's certificate against this name.
        If there is no match, assume that Splunk Enterprise is not authenticated against this server. You must specify this setting if sslVerifyServerCert is true.
        """
        return pulumi.get(self, "ssl_common_name_to_check")

    @ssl_common_name_to_check.setter
    def ssl_common_name_to_check(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ssl_common_name_to_check", value)

    @property
    @pulumi.getter(name="sslPassword")
    def ssl_password(self) -> Optional[pulumi.Input[str]]:
        """
        The password associated with the CAcert.
        The default Splunk Enterprise CAcert uses the password "password."
        """
        return pulumi.get(self, "ssl_password")

    @ssl_password.setter
    def ssl_password(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ssl_password", value)

    @property
    @pulumi.getter(name="sslRootCaPath")
    def ssl_root_ca_path(self) -> Optional[pulumi.Input[str]]:
        """
        The path to the root certificate authority file.
        """
        return pulumi.get(self, "ssl_root_ca_path")

    @ssl_root_ca_path.setter
    def ssl_root_ca_path(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ssl_root_ca_path", value)

    @property
    @pulumi.getter(name="sslVerifyServerCert")
    def ssl_verify_server_cert(self) -> Optional[pulumi.Input[bool]]:
        """
        If true, make sure that the server you are connecting to is a valid one (authenticated). Both the common name and the alternate name of the server are then checked for a match.
        """
        return pulumi.get(self, "ssl_verify_server_cert")

    @ssl_verify_server_cert.setter
    def ssl_verify_server_cert(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "ssl_verify_server_cert", value)


@pulumi.input_type
class _OutputsTcpServerState:
    def __init__(__self__, *,
                 acl: Optional[pulumi.Input['OutputsTcpServerAclArgs']] = None,
                 disabled: Optional[pulumi.Input[bool]] = None,
                 method: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 ssl_alt_name_to_check: Optional[pulumi.Input[str]] = None,
                 ssl_cert_path: Optional[pulumi.Input[str]] = None,
                 ssl_cipher: Optional[pulumi.Input[str]] = None,
                 ssl_common_name_to_check: Optional[pulumi.Input[str]] = None,
                 ssl_password: Optional[pulumi.Input[str]] = None,
                 ssl_root_ca_path: Optional[pulumi.Input[str]] = None,
                 ssl_verify_server_cert: Optional[pulumi.Input[bool]] = None):
        """
        Input properties used for looking up and filtering OutputsTcpServer resources.
        :param pulumi.Input['OutputsTcpServerAclArgs'] acl: The app/user context that is the namespace for the resource
        :param pulumi.Input[bool] disabled: If true, disables the group.
        :param pulumi.Input[str] method: Valid values: (clone | balance | autobalance)
               The data distribution method used when two or more servers exist in the same forwarder group.
        :param pulumi.Input[str] name: <host>:<port> of the Splunk receiver. <host> can be either an ip address or server name. <port> is the that port that the Splunk receiver is listening on.
        :param pulumi.Input[str] ssl_alt_name_to_check: The alternate name to match in the remote server's SSL certificate.
        :param pulumi.Input[str] ssl_cert_path: Path to the client certificate. If specified, connection uses SSL.
        :param pulumi.Input[str] ssl_cipher: SSL Cipher in the form ALL:!aNULL:!eNULL:!LOW:!EXP:RC4+RSA:+HIGH:+MEDIUM
        :param pulumi.Input[str] ssl_common_name_to_check: Check the common name of the server's certificate against this name.
               If there is no match, assume that Splunk Enterprise is not authenticated against this server. You must specify this setting if sslVerifyServerCert is true.
        :param pulumi.Input[str] ssl_password: The password associated with the CAcert.
               The default Splunk Enterprise CAcert uses the password "password."
        :param pulumi.Input[str] ssl_root_ca_path: The path to the root certificate authority file.
        :param pulumi.Input[bool] ssl_verify_server_cert: If true, make sure that the server you are connecting to is a valid one (authenticated). Both the common name and the alternate name of the server are then checked for a match.
        """
        if acl is not None:
            pulumi.set(__self__, "acl", acl)
        if disabled is not None:
            pulumi.set(__self__, "disabled", disabled)
        if method is not None:
            pulumi.set(__self__, "method", method)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if ssl_alt_name_to_check is not None:
            pulumi.set(__self__, "ssl_alt_name_to_check", ssl_alt_name_to_check)
        if ssl_cert_path is not None:
            pulumi.set(__self__, "ssl_cert_path", ssl_cert_path)
        if ssl_cipher is not None:
            pulumi.set(__self__, "ssl_cipher", ssl_cipher)
        if ssl_common_name_to_check is not None:
            pulumi.set(__self__, "ssl_common_name_to_check", ssl_common_name_to_check)
        if ssl_password is not None:
            pulumi.set(__self__, "ssl_password", ssl_password)
        if ssl_root_ca_path is not None:
            pulumi.set(__self__, "ssl_root_ca_path", ssl_root_ca_path)
        if ssl_verify_server_cert is not None:
            pulumi.set(__self__, "ssl_verify_server_cert", ssl_verify_server_cert)

    @property
    @pulumi.getter
    def acl(self) -> Optional[pulumi.Input['OutputsTcpServerAclArgs']]:
        """
        The app/user context that is the namespace for the resource
        """
        return pulumi.get(self, "acl")

    @acl.setter
    def acl(self, value: Optional[pulumi.Input['OutputsTcpServerAclArgs']]):
        pulumi.set(self, "acl", value)

    @property
    @pulumi.getter
    def disabled(self) -> Optional[pulumi.Input[bool]]:
        """
        If true, disables the group.
        """
        return pulumi.get(self, "disabled")

    @disabled.setter
    def disabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "disabled", value)

    @property
    @pulumi.getter
    def method(self) -> Optional[pulumi.Input[str]]:
        """
        Valid values: (clone | balance | autobalance)
        The data distribution method used when two or more servers exist in the same forwarder group.
        """
        return pulumi.get(self, "method")

    @method.setter
    def method(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "method", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        <host>:<port> of the Splunk receiver. <host> can be either an ip address or server name. <port> is the that port that the Splunk receiver is listening on.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="sslAltNameToCheck")
    def ssl_alt_name_to_check(self) -> Optional[pulumi.Input[str]]:
        """
        The alternate name to match in the remote server's SSL certificate.
        """
        return pulumi.get(self, "ssl_alt_name_to_check")

    @ssl_alt_name_to_check.setter
    def ssl_alt_name_to_check(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ssl_alt_name_to_check", value)

    @property
    @pulumi.getter(name="sslCertPath")
    def ssl_cert_path(self) -> Optional[pulumi.Input[str]]:
        """
        Path to the client certificate. If specified, connection uses SSL.
        """
        return pulumi.get(self, "ssl_cert_path")

    @ssl_cert_path.setter
    def ssl_cert_path(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ssl_cert_path", value)

    @property
    @pulumi.getter(name="sslCipher")
    def ssl_cipher(self) -> Optional[pulumi.Input[str]]:
        """
        SSL Cipher in the form ALL:!aNULL:!eNULL:!LOW:!EXP:RC4+RSA:+HIGH:+MEDIUM
        """
        return pulumi.get(self, "ssl_cipher")

    @ssl_cipher.setter
    def ssl_cipher(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ssl_cipher", value)

    @property
    @pulumi.getter(name="sslCommonNameToCheck")
    def ssl_common_name_to_check(self) -> Optional[pulumi.Input[str]]:
        """
        Check the common name of the server's certificate against this name.
        If there is no match, assume that Splunk Enterprise is not authenticated against this server. You must specify this setting if sslVerifyServerCert is true.
        """
        return pulumi.get(self, "ssl_common_name_to_check")

    @ssl_common_name_to_check.setter
    def ssl_common_name_to_check(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ssl_common_name_to_check", value)

    @property
    @pulumi.getter(name="sslPassword")
    def ssl_password(self) -> Optional[pulumi.Input[str]]:
        """
        The password associated with the CAcert.
        The default Splunk Enterprise CAcert uses the password "password."
        """
        return pulumi.get(self, "ssl_password")

    @ssl_password.setter
    def ssl_password(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ssl_password", value)

    @property
    @pulumi.getter(name="sslRootCaPath")
    def ssl_root_ca_path(self) -> Optional[pulumi.Input[str]]:
        """
        The path to the root certificate authority file.
        """
        return pulumi.get(self, "ssl_root_ca_path")

    @ssl_root_ca_path.setter
    def ssl_root_ca_path(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ssl_root_ca_path", value)

    @property
    @pulumi.getter(name="sslVerifyServerCert")
    def ssl_verify_server_cert(self) -> Optional[pulumi.Input[bool]]:
        """
        If true, make sure that the server you are connecting to is a valid one (authenticated). Both the common name and the alternate name of the server are then checked for a match.
        """
        return pulumi.get(self, "ssl_verify_server_cert")

    @ssl_verify_server_cert.setter
    def ssl_verify_server_cert(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "ssl_verify_server_cert", value)


class OutputsTcpServer(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 acl: Optional[pulumi.Input[pulumi.InputType['OutputsTcpServerAclArgs']]] = None,
                 disabled: Optional[pulumi.Input[bool]] = None,
                 method: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 ssl_alt_name_to_check: Optional[pulumi.Input[str]] = None,
                 ssl_cert_path: Optional[pulumi.Input[str]] = None,
                 ssl_cipher: Optional[pulumi.Input[str]] = None,
                 ssl_common_name_to_check: Optional[pulumi.Input[str]] = None,
                 ssl_password: Optional[pulumi.Input[str]] = None,
                 ssl_root_ca_path: Optional[pulumi.Input[str]] = None,
                 ssl_verify_server_cert: Optional[pulumi.Input[bool]] = None,
                 __props__=None):
        """
        ## # Resource: OutputsTcpServer

        Access data forwarding configurations.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_splunk as splunk

        tcp_server = splunk.OutputsTcpServer("tcpServer", ssl_alt_name_to_check="old-host")
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['OutputsTcpServerAclArgs']] acl: The app/user context that is the namespace for the resource
        :param pulumi.Input[bool] disabled: If true, disables the group.
        :param pulumi.Input[str] method: Valid values: (clone | balance | autobalance)
               The data distribution method used when two or more servers exist in the same forwarder group.
        :param pulumi.Input[str] name: <host>:<port> of the Splunk receiver. <host> can be either an ip address or server name. <port> is the that port that the Splunk receiver is listening on.
        :param pulumi.Input[str] ssl_alt_name_to_check: The alternate name to match in the remote server's SSL certificate.
        :param pulumi.Input[str] ssl_cert_path: Path to the client certificate. If specified, connection uses SSL.
        :param pulumi.Input[str] ssl_cipher: SSL Cipher in the form ALL:!aNULL:!eNULL:!LOW:!EXP:RC4+RSA:+HIGH:+MEDIUM
        :param pulumi.Input[str] ssl_common_name_to_check: Check the common name of the server's certificate against this name.
               If there is no match, assume that Splunk Enterprise is not authenticated against this server. You must specify this setting if sslVerifyServerCert is true.
        :param pulumi.Input[str] ssl_password: The password associated with the CAcert.
               The default Splunk Enterprise CAcert uses the password "password."
        :param pulumi.Input[str] ssl_root_ca_path: The path to the root certificate authority file.
        :param pulumi.Input[bool] ssl_verify_server_cert: If true, make sure that the server you are connecting to is a valid one (authenticated). Both the common name and the alternate name of the server are then checked for a match.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[OutputsTcpServerArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        ## # Resource: OutputsTcpServer

        Access data forwarding configurations.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_splunk as splunk

        tcp_server = splunk.OutputsTcpServer("tcpServer", ssl_alt_name_to_check="old-host")
        ```

        :param str resource_name: The name of the resource.
        :param OutputsTcpServerArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(OutputsTcpServerArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 acl: Optional[pulumi.Input[pulumi.InputType['OutputsTcpServerAclArgs']]] = None,
                 disabled: Optional[pulumi.Input[bool]] = None,
                 method: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 ssl_alt_name_to_check: Optional[pulumi.Input[str]] = None,
                 ssl_cert_path: Optional[pulumi.Input[str]] = None,
                 ssl_cipher: Optional[pulumi.Input[str]] = None,
                 ssl_common_name_to_check: Optional[pulumi.Input[str]] = None,
                 ssl_password: Optional[pulumi.Input[str]] = None,
                 ssl_root_ca_path: Optional[pulumi.Input[str]] = None,
                 ssl_verify_server_cert: Optional[pulumi.Input[bool]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = OutputsTcpServerArgs.__new__(OutputsTcpServerArgs)

            __props__.__dict__["acl"] = acl
            __props__.__dict__["disabled"] = disabled
            __props__.__dict__["method"] = method
            __props__.__dict__["name"] = name
            __props__.__dict__["ssl_alt_name_to_check"] = ssl_alt_name_to_check
            __props__.__dict__["ssl_cert_path"] = ssl_cert_path
            __props__.__dict__["ssl_cipher"] = ssl_cipher
            __props__.__dict__["ssl_common_name_to_check"] = ssl_common_name_to_check
            __props__.__dict__["ssl_password"] = ssl_password
            __props__.__dict__["ssl_root_ca_path"] = ssl_root_ca_path
            __props__.__dict__["ssl_verify_server_cert"] = ssl_verify_server_cert
        super(OutputsTcpServer, __self__).__init__(
            'splunk:index/outputsTcpServer:OutputsTcpServer',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            acl: Optional[pulumi.Input[pulumi.InputType['OutputsTcpServerAclArgs']]] = None,
            disabled: Optional[pulumi.Input[bool]] = None,
            method: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            ssl_alt_name_to_check: Optional[pulumi.Input[str]] = None,
            ssl_cert_path: Optional[pulumi.Input[str]] = None,
            ssl_cipher: Optional[pulumi.Input[str]] = None,
            ssl_common_name_to_check: Optional[pulumi.Input[str]] = None,
            ssl_password: Optional[pulumi.Input[str]] = None,
            ssl_root_ca_path: Optional[pulumi.Input[str]] = None,
            ssl_verify_server_cert: Optional[pulumi.Input[bool]] = None) -> 'OutputsTcpServer':
        """
        Get an existing OutputsTcpServer resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['OutputsTcpServerAclArgs']] acl: The app/user context that is the namespace for the resource
        :param pulumi.Input[bool] disabled: If true, disables the group.
        :param pulumi.Input[str] method: Valid values: (clone | balance | autobalance)
               The data distribution method used when two or more servers exist in the same forwarder group.
        :param pulumi.Input[str] name: <host>:<port> of the Splunk receiver. <host> can be either an ip address or server name. <port> is the that port that the Splunk receiver is listening on.
        :param pulumi.Input[str] ssl_alt_name_to_check: The alternate name to match in the remote server's SSL certificate.
        :param pulumi.Input[str] ssl_cert_path: Path to the client certificate. If specified, connection uses SSL.
        :param pulumi.Input[str] ssl_cipher: SSL Cipher in the form ALL:!aNULL:!eNULL:!LOW:!EXP:RC4+RSA:+HIGH:+MEDIUM
        :param pulumi.Input[str] ssl_common_name_to_check: Check the common name of the server's certificate against this name.
               If there is no match, assume that Splunk Enterprise is not authenticated against this server. You must specify this setting if sslVerifyServerCert is true.
        :param pulumi.Input[str] ssl_password: The password associated with the CAcert.
               The default Splunk Enterprise CAcert uses the password "password."
        :param pulumi.Input[str] ssl_root_ca_path: The path to the root certificate authority file.
        :param pulumi.Input[bool] ssl_verify_server_cert: If true, make sure that the server you are connecting to is a valid one (authenticated). Both the common name and the alternate name of the server are then checked for a match.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _OutputsTcpServerState.__new__(_OutputsTcpServerState)

        __props__.__dict__["acl"] = acl
        __props__.__dict__["disabled"] = disabled
        __props__.__dict__["method"] = method
        __props__.__dict__["name"] = name
        __props__.__dict__["ssl_alt_name_to_check"] = ssl_alt_name_to_check
        __props__.__dict__["ssl_cert_path"] = ssl_cert_path
        __props__.__dict__["ssl_cipher"] = ssl_cipher
        __props__.__dict__["ssl_common_name_to_check"] = ssl_common_name_to_check
        __props__.__dict__["ssl_password"] = ssl_password
        __props__.__dict__["ssl_root_ca_path"] = ssl_root_ca_path
        __props__.__dict__["ssl_verify_server_cert"] = ssl_verify_server_cert
        return OutputsTcpServer(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def acl(self) -> pulumi.Output['outputs.OutputsTcpServerAcl']:
        """
        The app/user context that is the namespace for the resource
        """
        return pulumi.get(self, "acl")

    @property
    @pulumi.getter
    def disabled(self) -> pulumi.Output[bool]:
        """
        If true, disables the group.
        """
        return pulumi.get(self, "disabled")

    @property
    @pulumi.getter
    def method(self) -> pulumi.Output[str]:
        """
        Valid values: (clone | balance | autobalance)
        The data distribution method used when two or more servers exist in the same forwarder group.
        """
        return pulumi.get(self, "method")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        <host>:<port> of the Splunk receiver. <host> can be either an ip address or server name. <port> is the that port that the Splunk receiver is listening on.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="sslAltNameToCheck")
    def ssl_alt_name_to_check(self) -> pulumi.Output[str]:
        """
        The alternate name to match in the remote server's SSL certificate.
        """
        return pulumi.get(self, "ssl_alt_name_to_check")

    @property
    @pulumi.getter(name="sslCertPath")
    def ssl_cert_path(self) -> pulumi.Output[str]:
        """
        Path to the client certificate. If specified, connection uses SSL.
        """
        return pulumi.get(self, "ssl_cert_path")

    @property
    @pulumi.getter(name="sslCipher")
    def ssl_cipher(self) -> pulumi.Output[str]:
        """
        SSL Cipher in the form ALL:!aNULL:!eNULL:!LOW:!EXP:RC4+RSA:+HIGH:+MEDIUM
        """
        return pulumi.get(self, "ssl_cipher")

    @property
    @pulumi.getter(name="sslCommonNameToCheck")
    def ssl_common_name_to_check(self) -> pulumi.Output[str]:
        """
        Check the common name of the server's certificate against this name.
        If there is no match, assume that Splunk Enterprise is not authenticated against this server. You must specify this setting if sslVerifyServerCert is true.
        """
        return pulumi.get(self, "ssl_common_name_to_check")

    @property
    @pulumi.getter(name="sslPassword")
    def ssl_password(self) -> pulumi.Output[str]:
        """
        The password associated with the CAcert.
        The default Splunk Enterprise CAcert uses the password "password."
        """
        return pulumi.get(self, "ssl_password")

    @property
    @pulumi.getter(name="sslRootCaPath")
    def ssl_root_ca_path(self) -> pulumi.Output[str]:
        """
        The path to the root certificate authority file.
        """
        return pulumi.get(self, "ssl_root_ca_path")

    @property
    @pulumi.getter(name="sslVerifyServerCert")
    def ssl_verify_server_cert(self) -> pulumi.Output[bool]:
        """
        If true, make sure that the server you are connecting to is a valid one (authenticated). Both the common name and the alternate name of the server are then checked for a match.
        """
        return pulumi.get(self, "ssl_verify_server_cert")

