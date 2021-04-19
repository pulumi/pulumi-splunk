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

__all__ = ['AppsLocalArgs', 'AppsLocal']

@pulumi.input_type
class AppsLocalArgs:
    def __init__(__self__, *,
                 acl: Optional[pulumi.Input['AppsLocalAclArgs']] = None,
                 auth: Optional[pulumi.Input[str]] = None,
                 author: Optional[pulumi.Input[str]] = None,
                 configured: Optional[pulumi.Input[bool]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 explicit_appname: Optional[pulumi.Input[str]] = None,
                 filename: Optional[pulumi.Input[bool]] = None,
                 label: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 session: Optional[pulumi.Input[str]] = None,
                 update: Optional[pulumi.Input[bool]] = None,
                 version: Optional[pulumi.Input[str]] = None,
                 visible: Optional[pulumi.Input[bool]] = None):
        """
        The set of arguments for constructing a AppsLocal resource.
        :param pulumi.Input['AppsLocalAclArgs'] acl: The app/user context that is the namespace for the resource
        :param pulumi.Input[str] auth: Splunkbase session token for operations like install and update that require login. Use auth or session when installing or updating an app through Splunkbase.
        :param pulumi.Input[str] author: For apps posted to Splunkbase, use your Splunk account username. For internal apps, include your name and contact information.
        :param pulumi.Input[bool] configured: Custom setup complete indication:
               <br>true = Custom app setup complete.
               <br>false = Custom app setup not complete.
        :param pulumi.Input[str] description: Short app description also displayed below the app title in Splunk Web Launcher.
        :param pulumi.Input[str] explicit_appname: Custom app name. Overrides name when installing an app from a file where filename is set to true. See also filename.
        :param pulumi.Input[bool] filename: Indicates whether to use the name value as the app source location.
               <br>true indicates that name is a path to a file to install.
               <br>false indicates that name is the literal app name and that the app is created from Splunkbase using a template.
        :param pulumi.Input[str] label: App name displayed in Splunk Web, from five to eighty characters excluding the prefix "Splunk for".
        :param pulumi.Input[str] name: Literal app name or path for the file to install, depending on the value of filename.
               <br>filename = false indicates that name is the literal app name and that the app is created from Splunkbase using a template.
               <br>filename = true indicates that name is the URL or path to the local .tar, .tgz or .spl file. If name is the Splunkbase URL, set auth or session to authenticate the request.
               The app folder name cannot include spaces or special characters.
        :param pulumi.Input[str] session: Login session token for installing or updating an app on Splunkbase. Alternatively, use auth.
        :param pulumi.Input[bool] update: File-based update indication:
               <br>true specifies that filename should be used to update an existing app. If not specified, update defaults to
               <br>false, which indicates that filename should not be used to update an existing app.
        :param pulumi.Input[str] version: App version.
        :param pulumi.Input[bool] visible: Indicates whether the app is visible and navigable from Splunk Web.
               <br>true = App is visible and navigable.
               <br>false = App is not visible or navigable.
        """
        if acl is not None:
            pulumi.set(__self__, "acl", acl)
        if auth is not None:
            pulumi.set(__self__, "auth", auth)
        if author is not None:
            pulumi.set(__self__, "author", author)
        if configured is not None:
            pulumi.set(__self__, "configured", configured)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if explicit_appname is not None:
            pulumi.set(__self__, "explicit_appname", explicit_appname)
        if filename is not None:
            pulumi.set(__self__, "filename", filename)
        if label is not None:
            pulumi.set(__self__, "label", label)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if session is not None:
            pulumi.set(__self__, "session", session)
        if update is not None:
            pulumi.set(__self__, "update", update)
        if version is not None:
            pulumi.set(__self__, "version", version)
        if visible is not None:
            pulumi.set(__self__, "visible", visible)

    @property
    @pulumi.getter
    def acl(self) -> Optional[pulumi.Input['AppsLocalAclArgs']]:
        """
        The app/user context that is the namespace for the resource
        """
        return pulumi.get(self, "acl")

    @acl.setter
    def acl(self, value: Optional[pulumi.Input['AppsLocalAclArgs']]):
        pulumi.set(self, "acl", value)

    @property
    @pulumi.getter
    def auth(self) -> Optional[pulumi.Input[str]]:
        """
        Splunkbase session token for operations like install and update that require login. Use auth or session when installing or updating an app through Splunkbase.
        """
        return pulumi.get(self, "auth")

    @auth.setter
    def auth(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "auth", value)

    @property
    @pulumi.getter
    def author(self) -> Optional[pulumi.Input[str]]:
        """
        For apps posted to Splunkbase, use your Splunk account username. For internal apps, include your name and contact information.
        """
        return pulumi.get(self, "author")

    @author.setter
    def author(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "author", value)

    @property
    @pulumi.getter
    def configured(self) -> Optional[pulumi.Input[bool]]:
        """
        Custom setup complete indication:
        <br>true = Custom app setup complete.
        <br>false = Custom app setup not complete.
        """
        return pulumi.get(self, "configured")

    @configured.setter
    def configured(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "configured", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        Short app description also displayed below the app title in Splunk Web Launcher.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="explicitAppname")
    def explicit_appname(self) -> Optional[pulumi.Input[str]]:
        """
        Custom app name. Overrides name when installing an app from a file where filename is set to true. See also filename.
        """
        return pulumi.get(self, "explicit_appname")

    @explicit_appname.setter
    def explicit_appname(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "explicit_appname", value)

    @property
    @pulumi.getter
    def filename(self) -> Optional[pulumi.Input[bool]]:
        """
        Indicates whether to use the name value as the app source location.
        <br>true indicates that name is a path to a file to install.
        <br>false indicates that name is the literal app name and that the app is created from Splunkbase using a template.
        """
        return pulumi.get(self, "filename")

    @filename.setter
    def filename(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "filename", value)

    @property
    @pulumi.getter
    def label(self) -> Optional[pulumi.Input[str]]:
        """
        App name displayed in Splunk Web, from five to eighty characters excluding the prefix "Splunk for".
        """
        return pulumi.get(self, "label")

    @label.setter
    def label(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "label", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Literal app name or path for the file to install, depending on the value of filename.
        <br>filename = false indicates that name is the literal app name and that the app is created from Splunkbase using a template.
        <br>filename = true indicates that name is the URL or path to the local .tar, .tgz or .spl file. If name is the Splunkbase URL, set auth or session to authenticate the request.
        The app folder name cannot include spaces or special characters.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def session(self) -> Optional[pulumi.Input[str]]:
        """
        Login session token for installing or updating an app on Splunkbase. Alternatively, use auth.
        """
        return pulumi.get(self, "session")

    @session.setter
    def session(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "session", value)

    @property
    @pulumi.getter
    def update(self) -> Optional[pulumi.Input[bool]]:
        """
        File-based update indication:
        <br>true specifies that filename should be used to update an existing app. If not specified, update defaults to
        <br>false, which indicates that filename should not be used to update an existing app.
        """
        return pulumi.get(self, "update")

    @update.setter
    def update(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "update", value)

    @property
    @pulumi.getter
    def version(self) -> Optional[pulumi.Input[str]]:
        """
        App version.
        """
        return pulumi.get(self, "version")

    @version.setter
    def version(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "version", value)

    @property
    @pulumi.getter
    def visible(self) -> Optional[pulumi.Input[bool]]:
        """
        Indicates whether the app is visible and navigable from Splunk Web.
        <br>true = App is visible and navigable.
        <br>false = App is not visible or navigable.
        """
        return pulumi.get(self, "visible")

    @visible.setter
    def visible(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "visible", value)


@pulumi.input_type
class _AppsLocalState:
    def __init__(__self__, *,
                 acl: Optional[pulumi.Input['AppsLocalAclArgs']] = None,
                 auth: Optional[pulumi.Input[str]] = None,
                 author: Optional[pulumi.Input[str]] = None,
                 configured: Optional[pulumi.Input[bool]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 explicit_appname: Optional[pulumi.Input[str]] = None,
                 filename: Optional[pulumi.Input[bool]] = None,
                 label: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 session: Optional[pulumi.Input[str]] = None,
                 update: Optional[pulumi.Input[bool]] = None,
                 version: Optional[pulumi.Input[str]] = None,
                 visible: Optional[pulumi.Input[bool]] = None):
        """
        Input properties used for looking up and filtering AppsLocal resources.
        :param pulumi.Input['AppsLocalAclArgs'] acl: The app/user context that is the namespace for the resource
        :param pulumi.Input[str] auth: Splunkbase session token for operations like install and update that require login. Use auth or session when installing or updating an app through Splunkbase.
        :param pulumi.Input[str] author: For apps posted to Splunkbase, use your Splunk account username. For internal apps, include your name and contact information.
        :param pulumi.Input[bool] configured: Custom setup complete indication:
               <br>true = Custom app setup complete.
               <br>false = Custom app setup not complete.
        :param pulumi.Input[str] description: Short app description also displayed below the app title in Splunk Web Launcher.
        :param pulumi.Input[str] explicit_appname: Custom app name. Overrides name when installing an app from a file where filename is set to true. See also filename.
        :param pulumi.Input[bool] filename: Indicates whether to use the name value as the app source location.
               <br>true indicates that name is a path to a file to install.
               <br>false indicates that name is the literal app name and that the app is created from Splunkbase using a template.
        :param pulumi.Input[str] label: App name displayed in Splunk Web, from five to eighty characters excluding the prefix "Splunk for".
        :param pulumi.Input[str] name: Literal app name or path for the file to install, depending on the value of filename.
               <br>filename = false indicates that name is the literal app name and that the app is created from Splunkbase using a template.
               <br>filename = true indicates that name is the URL or path to the local .tar, .tgz or .spl file. If name is the Splunkbase URL, set auth or session to authenticate the request.
               The app folder name cannot include spaces or special characters.
        :param pulumi.Input[str] session: Login session token for installing or updating an app on Splunkbase. Alternatively, use auth.
        :param pulumi.Input[bool] update: File-based update indication:
               <br>true specifies that filename should be used to update an existing app. If not specified, update defaults to
               <br>false, which indicates that filename should not be used to update an existing app.
        :param pulumi.Input[str] version: App version.
        :param pulumi.Input[bool] visible: Indicates whether the app is visible and navigable from Splunk Web.
               <br>true = App is visible and navigable.
               <br>false = App is not visible or navigable.
        """
        if acl is not None:
            pulumi.set(__self__, "acl", acl)
        if auth is not None:
            pulumi.set(__self__, "auth", auth)
        if author is not None:
            pulumi.set(__self__, "author", author)
        if configured is not None:
            pulumi.set(__self__, "configured", configured)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if explicit_appname is not None:
            pulumi.set(__self__, "explicit_appname", explicit_appname)
        if filename is not None:
            pulumi.set(__self__, "filename", filename)
        if label is not None:
            pulumi.set(__self__, "label", label)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if session is not None:
            pulumi.set(__self__, "session", session)
        if update is not None:
            pulumi.set(__self__, "update", update)
        if version is not None:
            pulumi.set(__self__, "version", version)
        if visible is not None:
            pulumi.set(__self__, "visible", visible)

    @property
    @pulumi.getter
    def acl(self) -> Optional[pulumi.Input['AppsLocalAclArgs']]:
        """
        The app/user context that is the namespace for the resource
        """
        return pulumi.get(self, "acl")

    @acl.setter
    def acl(self, value: Optional[pulumi.Input['AppsLocalAclArgs']]):
        pulumi.set(self, "acl", value)

    @property
    @pulumi.getter
    def auth(self) -> Optional[pulumi.Input[str]]:
        """
        Splunkbase session token for operations like install and update that require login. Use auth or session when installing or updating an app through Splunkbase.
        """
        return pulumi.get(self, "auth")

    @auth.setter
    def auth(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "auth", value)

    @property
    @pulumi.getter
    def author(self) -> Optional[pulumi.Input[str]]:
        """
        For apps posted to Splunkbase, use your Splunk account username. For internal apps, include your name and contact information.
        """
        return pulumi.get(self, "author")

    @author.setter
    def author(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "author", value)

    @property
    @pulumi.getter
    def configured(self) -> Optional[pulumi.Input[bool]]:
        """
        Custom setup complete indication:
        <br>true = Custom app setup complete.
        <br>false = Custom app setup not complete.
        """
        return pulumi.get(self, "configured")

    @configured.setter
    def configured(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "configured", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        Short app description also displayed below the app title in Splunk Web Launcher.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="explicitAppname")
    def explicit_appname(self) -> Optional[pulumi.Input[str]]:
        """
        Custom app name. Overrides name when installing an app from a file where filename is set to true. See also filename.
        """
        return pulumi.get(self, "explicit_appname")

    @explicit_appname.setter
    def explicit_appname(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "explicit_appname", value)

    @property
    @pulumi.getter
    def filename(self) -> Optional[pulumi.Input[bool]]:
        """
        Indicates whether to use the name value as the app source location.
        <br>true indicates that name is a path to a file to install.
        <br>false indicates that name is the literal app name and that the app is created from Splunkbase using a template.
        """
        return pulumi.get(self, "filename")

    @filename.setter
    def filename(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "filename", value)

    @property
    @pulumi.getter
    def label(self) -> Optional[pulumi.Input[str]]:
        """
        App name displayed in Splunk Web, from five to eighty characters excluding the prefix "Splunk for".
        """
        return pulumi.get(self, "label")

    @label.setter
    def label(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "label", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Literal app name or path for the file to install, depending on the value of filename.
        <br>filename = false indicates that name is the literal app name and that the app is created from Splunkbase using a template.
        <br>filename = true indicates that name is the URL or path to the local .tar, .tgz or .spl file. If name is the Splunkbase URL, set auth or session to authenticate the request.
        The app folder name cannot include spaces or special characters.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def session(self) -> Optional[pulumi.Input[str]]:
        """
        Login session token for installing or updating an app on Splunkbase. Alternatively, use auth.
        """
        return pulumi.get(self, "session")

    @session.setter
    def session(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "session", value)

    @property
    @pulumi.getter
    def update(self) -> Optional[pulumi.Input[bool]]:
        """
        File-based update indication:
        <br>true specifies that filename should be used to update an existing app. If not specified, update defaults to
        <br>false, which indicates that filename should not be used to update an existing app.
        """
        return pulumi.get(self, "update")

    @update.setter
    def update(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "update", value)

    @property
    @pulumi.getter
    def version(self) -> Optional[pulumi.Input[str]]:
        """
        App version.
        """
        return pulumi.get(self, "version")

    @version.setter
    def version(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "version", value)

    @property
    @pulumi.getter
    def visible(self) -> Optional[pulumi.Input[bool]]:
        """
        Indicates whether the app is visible and navigable from Splunk Web.
        <br>true = App is visible and navigable.
        <br>false = App is not visible or navigable.
        """
        return pulumi.get(self, "visible")

    @visible.setter
    def visible(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "visible", value)


class AppsLocal(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 acl: Optional[pulumi.Input[pulumi.InputType['AppsLocalAclArgs']]] = None,
                 auth: Optional[pulumi.Input[str]] = None,
                 author: Optional[pulumi.Input[str]] = None,
                 configured: Optional[pulumi.Input[bool]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 explicit_appname: Optional[pulumi.Input[str]] = None,
                 filename: Optional[pulumi.Input[bool]] = None,
                 label: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 session: Optional[pulumi.Input[str]] = None,
                 update: Optional[pulumi.Input[bool]] = None,
                 version: Optional[pulumi.Input[str]] = None,
                 visible: Optional[pulumi.Input[bool]] = None,
                 __props__=None):
        """
        ## # Resource: AppsLocal

        Create, install and manage apps on your Splunk instance

        ## Example Usage

        ```python
        import pulumi
        import pulumi_splunk as splunk

        amazon_connect_app = splunk.AppsLocal("amazonConnectApp",
            explicit_appname="amazon_connect_app_for_splunk",
            filename=True)
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['AppsLocalAclArgs']] acl: The app/user context that is the namespace for the resource
        :param pulumi.Input[str] auth: Splunkbase session token for operations like install and update that require login. Use auth or session when installing or updating an app through Splunkbase.
        :param pulumi.Input[str] author: For apps posted to Splunkbase, use your Splunk account username. For internal apps, include your name and contact information.
        :param pulumi.Input[bool] configured: Custom setup complete indication:
               <br>true = Custom app setup complete.
               <br>false = Custom app setup not complete.
        :param pulumi.Input[str] description: Short app description also displayed below the app title in Splunk Web Launcher.
        :param pulumi.Input[str] explicit_appname: Custom app name. Overrides name when installing an app from a file where filename is set to true. See also filename.
        :param pulumi.Input[bool] filename: Indicates whether to use the name value as the app source location.
               <br>true indicates that name is a path to a file to install.
               <br>false indicates that name is the literal app name and that the app is created from Splunkbase using a template.
        :param pulumi.Input[str] label: App name displayed in Splunk Web, from five to eighty characters excluding the prefix "Splunk for".
        :param pulumi.Input[str] name: Literal app name or path for the file to install, depending on the value of filename.
               <br>filename = false indicates that name is the literal app name and that the app is created from Splunkbase using a template.
               <br>filename = true indicates that name is the URL or path to the local .tar, .tgz or .spl file. If name is the Splunkbase URL, set auth or session to authenticate the request.
               The app folder name cannot include spaces or special characters.
        :param pulumi.Input[str] session: Login session token for installing or updating an app on Splunkbase. Alternatively, use auth.
        :param pulumi.Input[bool] update: File-based update indication:
               <br>true specifies that filename should be used to update an existing app. If not specified, update defaults to
               <br>false, which indicates that filename should not be used to update an existing app.
        :param pulumi.Input[str] version: App version.
        :param pulumi.Input[bool] visible: Indicates whether the app is visible and navigable from Splunk Web.
               <br>true = App is visible and navigable.
               <br>false = App is not visible or navigable.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[AppsLocalArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        ## # Resource: AppsLocal

        Create, install and manage apps on your Splunk instance

        ## Example Usage

        ```python
        import pulumi
        import pulumi_splunk as splunk

        amazon_connect_app = splunk.AppsLocal("amazonConnectApp",
            explicit_appname="amazon_connect_app_for_splunk",
            filename=True)
        ```

        :param str resource_name: The name of the resource.
        :param AppsLocalArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(AppsLocalArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 acl: Optional[pulumi.Input[pulumi.InputType['AppsLocalAclArgs']]] = None,
                 auth: Optional[pulumi.Input[str]] = None,
                 author: Optional[pulumi.Input[str]] = None,
                 configured: Optional[pulumi.Input[bool]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 explicit_appname: Optional[pulumi.Input[str]] = None,
                 filename: Optional[pulumi.Input[bool]] = None,
                 label: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 session: Optional[pulumi.Input[str]] = None,
                 update: Optional[pulumi.Input[bool]] = None,
                 version: Optional[pulumi.Input[str]] = None,
                 visible: Optional[pulumi.Input[bool]] = None,
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
            __props__ = AppsLocalArgs.__new__(AppsLocalArgs)

            __props__.__dict__["acl"] = acl
            __props__.__dict__["auth"] = auth
            __props__.__dict__["author"] = author
            __props__.__dict__["configured"] = configured
            __props__.__dict__["description"] = description
            __props__.__dict__["explicit_appname"] = explicit_appname
            __props__.__dict__["filename"] = filename
            __props__.__dict__["label"] = label
            __props__.__dict__["name"] = name
            __props__.__dict__["session"] = session
            __props__.__dict__["update"] = update
            __props__.__dict__["version"] = version
            __props__.__dict__["visible"] = visible
        super(AppsLocal, __self__).__init__(
            'splunk:index/appsLocal:AppsLocal',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            acl: Optional[pulumi.Input[pulumi.InputType['AppsLocalAclArgs']]] = None,
            auth: Optional[pulumi.Input[str]] = None,
            author: Optional[pulumi.Input[str]] = None,
            configured: Optional[pulumi.Input[bool]] = None,
            description: Optional[pulumi.Input[str]] = None,
            explicit_appname: Optional[pulumi.Input[str]] = None,
            filename: Optional[pulumi.Input[bool]] = None,
            label: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            session: Optional[pulumi.Input[str]] = None,
            update: Optional[pulumi.Input[bool]] = None,
            version: Optional[pulumi.Input[str]] = None,
            visible: Optional[pulumi.Input[bool]] = None) -> 'AppsLocal':
        """
        Get an existing AppsLocal resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['AppsLocalAclArgs']] acl: The app/user context that is the namespace for the resource
        :param pulumi.Input[str] auth: Splunkbase session token for operations like install and update that require login. Use auth or session when installing or updating an app through Splunkbase.
        :param pulumi.Input[str] author: For apps posted to Splunkbase, use your Splunk account username. For internal apps, include your name and contact information.
        :param pulumi.Input[bool] configured: Custom setup complete indication:
               <br>true = Custom app setup complete.
               <br>false = Custom app setup not complete.
        :param pulumi.Input[str] description: Short app description also displayed below the app title in Splunk Web Launcher.
        :param pulumi.Input[str] explicit_appname: Custom app name. Overrides name when installing an app from a file where filename is set to true. See also filename.
        :param pulumi.Input[bool] filename: Indicates whether to use the name value as the app source location.
               <br>true indicates that name is a path to a file to install.
               <br>false indicates that name is the literal app name and that the app is created from Splunkbase using a template.
        :param pulumi.Input[str] label: App name displayed in Splunk Web, from five to eighty characters excluding the prefix "Splunk for".
        :param pulumi.Input[str] name: Literal app name or path for the file to install, depending on the value of filename.
               <br>filename = false indicates that name is the literal app name and that the app is created from Splunkbase using a template.
               <br>filename = true indicates that name is the URL or path to the local .tar, .tgz or .spl file. If name is the Splunkbase URL, set auth or session to authenticate the request.
               The app folder name cannot include spaces or special characters.
        :param pulumi.Input[str] session: Login session token for installing or updating an app on Splunkbase. Alternatively, use auth.
        :param pulumi.Input[bool] update: File-based update indication:
               <br>true specifies that filename should be used to update an existing app. If not specified, update defaults to
               <br>false, which indicates that filename should not be used to update an existing app.
        :param pulumi.Input[str] version: App version.
        :param pulumi.Input[bool] visible: Indicates whether the app is visible and navigable from Splunk Web.
               <br>true = App is visible and navigable.
               <br>false = App is not visible or navigable.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _AppsLocalState.__new__(_AppsLocalState)

        __props__.__dict__["acl"] = acl
        __props__.__dict__["auth"] = auth
        __props__.__dict__["author"] = author
        __props__.__dict__["configured"] = configured
        __props__.__dict__["description"] = description
        __props__.__dict__["explicit_appname"] = explicit_appname
        __props__.__dict__["filename"] = filename
        __props__.__dict__["label"] = label
        __props__.__dict__["name"] = name
        __props__.__dict__["session"] = session
        __props__.__dict__["update"] = update
        __props__.__dict__["version"] = version
        __props__.__dict__["visible"] = visible
        return AppsLocal(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def acl(self) -> pulumi.Output['outputs.AppsLocalAcl']:
        """
        The app/user context that is the namespace for the resource
        """
        return pulumi.get(self, "acl")

    @property
    @pulumi.getter
    def auth(self) -> pulumi.Output[Optional[str]]:
        """
        Splunkbase session token for operations like install and update that require login. Use auth or session when installing or updating an app through Splunkbase.
        """
        return pulumi.get(self, "auth")

    @property
    @pulumi.getter
    def author(self) -> pulumi.Output[str]:
        """
        For apps posted to Splunkbase, use your Splunk account username. For internal apps, include your name and contact information.
        """
        return pulumi.get(self, "author")

    @property
    @pulumi.getter
    def configured(self) -> pulumi.Output[bool]:
        """
        Custom setup complete indication:
        <br>true = Custom app setup complete.
        <br>false = Custom app setup not complete.
        """
        return pulumi.get(self, "configured")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[str]:
        """
        Short app description also displayed below the app title in Splunk Web Launcher.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="explicitAppname")
    def explicit_appname(self) -> pulumi.Output[Optional[str]]:
        """
        Custom app name. Overrides name when installing an app from a file where filename is set to true. See also filename.
        """
        return pulumi.get(self, "explicit_appname")

    @property
    @pulumi.getter
    def filename(self) -> pulumi.Output[Optional[bool]]:
        """
        Indicates whether to use the name value as the app source location.
        <br>true indicates that name is a path to a file to install.
        <br>false indicates that name is the literal app name and that the app is created from Splunkbase using a template.
        """
        return pulumi.get(self, "filename")

    @property
    @pulumi.getter
    def label(self) -> pulumi.Output[str]:
        """
        App name displayed in Splunk Web, from five to eighty characters excluding the prefix "Splunk for".
        """
        return pulumi.get(self, "label")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Literal app name or path for the file to install, depending on the value of filename.
        <br>filename = false indicates that name is the literal app name and that the app is created from Splunkbase using a template.
        <br>filename = true indicates that name is the URL or path to the local .tar, .tgz or .spl file. If name is the Splunkbase URL, set auth or session to authenticate the request.
        The app folder name cannot include spaces or special characters.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def session(self) -> pulumi.Output[Optional[str]]:
        """
        Login session token for installing or updating an app on Splunkbase. Alternatively, use auth.
        """
        return pulumi.get(self, "session")

    @property
    @pulumi.getter
    def update(self) -> pulumi.Output[Optional[bool]]:
        """
        File-based update indication:
        <br>true specifies that filename should be used to update an existing app. If not specified, update defaults to
        <br>false, which indicates that filename should not be used to update an existing app.
        """
        return pulumi.get(self, "update")

    @property
    @pulumi.getter
    def version(self) -> pulumi.Output[str]:
        """
        App version.
        """
        return pulumi.get(self, "version")

    @property
    @pulumi.getter
    def visible(self) -> pulumi.Output[bool]:
        """
        Indicates whether the app is visible and navigable from Splunk Web.
        <br>true = App is visible and navigable.
        <br>false = App is not visible or navigable.
        """
        return pulumi.get(self, "visible")

