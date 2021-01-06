# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

# Export this package's modules as members:
from .admin_saml_groups import *
from .apps_local import *
from .authentication_users import *
from .authorization_roles import *
from .configs_conf import *
from .data_ui_views import *
from .global_http_event_collector import *
from .indexes import *
from .inputs_http_event_collector import *
from .inputs_monitor import *
from .inputs_script import *
from .inputs_tcp_cooked import *
from .inputs_tcp_raw import *
from .inputs_tcp_splunk_tcp_token import *
from .inputs_tcp_ssl import *
from .inputs_udp import *
from .outputs_tcp_default import *
from .outputs_tcp_group import *
from .outputs_tcp_server import *
from .outputs_tcp_syslog import *
from .provider import *
from .saved_searches import *
from ._inputs import *
from . import outputs

# Make subpackages available:
from . import (
    config,
)

def _register_module():
    import pulumi
    from . import _utilities


    class Module(pulumi.runtime.ResourceModule):
        _version = _utilities.get_semver_version()

        def version(self):
            return Module._version

        def construct(self, name: str, typ: str, urn: str) -> pulumi.Resource:
            if typ == "splunk:index/adminSamlGroups:AdminSamlGroups":
                return AdminSamlGroups(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "splunk:index/appsLocal:AppsLocal":
                return AppsLocal(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "splunk:index/authenticationUsers:AuthenticationUsers":
                return AuthenticationUsers(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "splunk:index/authorizationRoles:AuthorizationRoles":
                return AuthorizationRoles(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "splunk:index/configsConf:ConfigsConf":
                return ConfigsConf(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "splunk:index/dataUiViews:DataUiViews":
                return DataUiViews(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "splunk:index/globalHttpEventCollector:GlobalHttpEventCollector":
                return GlobalHttpEventCollector(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "splunk:index/indexes:Indexes":
                return Indexes(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "splunk:index/inputsHttpEventCollector:InputsHttpEventCollector":
                return InputsHttpEventCollector(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "splunk:index/inputsMonitor:InputsMonitor":
                return InputsMonitor(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "splunk:index/inputsScript:InputsScript":
                return InputsScript(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "splunk:index/inputsTcpCooked:InputsTcpCooked":
                return InputsTcpCooked(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "splunk:index/inputsTcpRaw:InputsTcpRaw":
                return InputsTcpRaw(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "splunk:index/inputsTcpSplunkTcpToken:InputsTcpSplunkTcpToken":
                return InputsTcpSplunkTcpToken(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "splunk:index/inputsTcpSsl:InputsTcpSsl":
                return InputsTcpSsl(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "splunk:index/inputsUdp:InputsUdp":
                return InputsUdp(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "splunk:index/outputsTcpDefault:OutputsTcpDefault":
                return OutputsTcpDefault(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "splunk:index/outputsTcpGroup:OutputsTcpGroup":
                return OutputsTcpGroup(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "splunk:index/outputsTcpServer:OutputsTcpServer":
                return OutputsTcpServer(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "splunk:index/outputsTcpSyslog:OutputsTcpSyslog":
                return OutputsTcpSyslog(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "splunk:index/savedSearches:SavedSearches":
                return SavedSearches(name, pulumi.ResourceOptions(urn=urn))
            else:
                raise Exception(f"unknown resource type {typ}")


    _module_instance = Module()
    pulumi.runtime.register_resource_module("splunk", "index/adminSamlGroups", _module_instance)
    pulumi.runtime.register_resource_module("splunk", "index/appsLocal", _module_instance)
    pulumi.runtime.register_resource_module("splunk", "index/authenticationUsers", _module_instance)
    pulumi.runtime.register_resource_module("splunk", "index/authorizationRoles", _module_instance)
    pulumi.runtime.register_resource_module("splunk", "index/configsConf", _module_instance)
    pulumi.runtime.register_resource_module("splunk", "index/dataUiViews", _module_instance)
    pulumi.runtime.register_resource_module("splunk", "index/globalHttpEventCollector", _module_instance)
    pulumi.runtime.register_resource_module("splunk", "index/indexes", _module_instance)
    pulumi.runtime.register_resource_module("splunk", "index/inputsHttpEventCollector", _module_instance)
    pulumi.runtime.register_resource_module("splunk", "index/inputsMonitor", _module_instance)
    pulumi.runtime.register_resource_module("splunk", "index/inputsScript", _module_instance)
    pulumi.runtime.register_resource_module("splunk", "index/inputsTcpCooked", _module_instance)
    pulumi.runtime.register_resource_module("splunk", "index/inputsTcpRaw", _module_instance)
    pulumi.runtime.register_resource_module("splunk", "index/inputsTcpSplunkTcpToken", _module_instance)
    pulumi.runtime.register_resource_module("splunk", "index/inputsTcpSsl", _module_instance)
    pulumi.runtime.register_resource_module("splunk", "index/inputsUdp", _module_instance)
    pulumi.runtime.register_resource_module("splunk", "index/outputsTcpDefault", _module_instance)
    pulumi.runtime.register_resource_module("splunk", "index/outputsTcpGroup", _module_instance)
    pulumi.runtime.register_resource_module("splunk", "index/outputsTcpServer", _module_instance)
    pulumi.runtime.register_resource_module("splunk", "index/outputsTcpSyslog", _module_instance)
    pulumi.runtime.register_resource_module("splunk", "index/savedSearches", _module_instance)


    class Package(pulumi.runtime.ResourcePackage):
        _version = _utilities.get_semver_version()

        def version(self):
            return Package._version

        def construct_provider(self, name: str, typ: str, urn: str) -> pulumi.ProviderResource:
            if typ != "pulumi:providers:splunk":
                raise Exception(f"unknown provider type {typ}")
            return Provider(name, pulumi.ResourceOptions(urn=urn))


    pulumi.runtime.register_resource_package("splunk", Package())

_register_module()
