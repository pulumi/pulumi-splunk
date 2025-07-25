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

__all__ = ['AuthenticationUsersArgs', 'AuthenticationUsers']

@pulumi.input_type
class AuthenticationUsersArgs:
    def __init__(__self__, *,
                 default_app: Optional[pulumi.Input[_builtins.str]] = None,
                 email: Optional[pulumi.Input[_builtins.str]] = None,
                 force_change_pass: Optional[pulumi.Input[_builtins.bool]] = None,
                 name: Optional[pulumi.Input[_builtins.str]] = None,
                 password: Optional[pulumi.Input[_builtins.str]] = None,
                 realname: Optional[pulumi.Input[_builtins.str]] = None,
                 restart_background_jobs: Optional[pulumi.Input[_builtins.bool]] = None,
                 roles: Optional[pulumi.Input[Sequence[pulumi.Input[_builtins.str]]]] = None,
                 tz: Optional[pulumi.Input[_builtins.str]] = None):
        """
        The set of arguments for constructing a AuthenticationUsers resource.
        :param pulumi.Input[_builtins.str] default_app: User default app. Overrides the default app inherited from the user roles.
        :param pulumi.Input[_builtins.str] email: User email address.
        :param pulumi.Input[_builtins.bool] force_change_pass: Force user to change password indication
        :param pulumi.Input[_builtins.str] name: Unique user login name.
        :param pulumi.Input[_builtins.str] password: User login password.
        :param pulumi.Input[_builtins.str] realname: Full user name.
        :param pulumi.Input[_builtins.bool] restart_background_jobs: Restart background search job that has not completed when Splunk restarts indication.
        :param pulumi.Input[Sequence[pulumi.Input[_builtins.str]]] roles: Role to assign to this user. At least one existing role is required.
        :param pulumi.Input[_builtins.str] tz: User timezone.
        """
        if default_app is not None:
            pulumi.set(__self__, "default_app", default_app)
        if email is not None:
            pulumi.set(__self__, "email", email)
        if force_change_pass is not None:
            pulumi.set(__self__, "force_change_pass", force_change_pass)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if password is not None:
            pulumi.set(__self__, "password", password)
        if realname is not None:
            pulumi.set(__self__, "realname", realname)
        if restart_background_jobs is not None:
            pulumi.set(__self__, "restart_background_jobs", restart_background_jobs)
        if roles is not None:
            pulumi.set(__self__, "roles", roles)
        if tz is not None:
            pulumi.set(__self__, "tz", tz)

    @_builtins.property
    @pulumi.getter(name="defaultApp")
    def default_app(self) -> Optional[pulumi.Input[_builtins.str]]:
        """
        User default app. Overrides the default app inherited from the user roles.
        """
        return pulumi.get(self, "default_app")

    @default_app.setter
    def default_app(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "default_app", value)

    @_builtins.property
    @pulumi.getter
    def email(self) -> Optional[pulumi.Input[_builtins.str]]:
        """
        User email address.
        """
        return pulumi.get(self, "email")

    @email.setter
    def email(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "email", value)

    @_builtins.property
    @pulumi.getter(name="forceChangePass")
    def force_change_pass(self) -> Optional[pulumi.Input[_builtins.bool]]:
        """
        Force user to change password indication
        """
        return pulumi.get(self, "force_change_pass")

    @force_change_pass.setter
    def force_change_pass(self, value: Optional[pulumi.Input[_builtins.bool]]):
        pulumi.set(self, "force_change_pass", value)

    @_builtins.property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[_builtins.str]]:
        """
        Unique user login name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "name", value)

    @_builtins.property
    @pulumi.getter
    def password(self) -> Optional[pulumi.Input[_builtins.str]]:
        """
        User login password.
        """
        return pulumi.get(self, "password")

    @password.setter
    def password(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "password", value)

    @_builtins.property
    @pulumi.getter
    def realname(self) -> Optional[pulumi.Input[_builtins.str]]:
        """
        Full user name.
        """
        return pulumi.get(self, "realname")

    @realname.setter
    def realname(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "realname", value)

    @_builtins.property
    @pulumi.getter(name="restartBackgroundJobs")
    def restart_background_jobs(self) -> Optional[pulumi.Input[_builtins.bool]]:
        """
        Restart background search job that has not completed when Splunk restarts indication.
        """
        return pulumi.get(self, "restart_background_jobs")

    @restart_background_jobs.setter
    def restart_background_jobs(self, value: Optional[pulumi.Input[_builtins.bool]]):
        pulumi.set(self, "restart_background_jobs", value)

    @_builtins.property
    @pulumi.getter
    def roles(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[_builtins.str]]]]:
        """
        Role to assign to this user. At least one existing role is required.
        """
        return pulumi.get(self, "roles")

    @roles.setter
    def roles(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[_builtins.str]]]]):
        pulumi.set(self, "roles", value)

    @_builtins.property
    @pulumi.getter
    def tz(self) -> Optional[pulumi.Input[_builtins.str]]:
        """
        User timezone.
        """
        return pulumi.get(self, "tz")

    @tz.setter
    def tz(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "tz", value)


@pulumi.input_type
class _AuthenticationUsersState:
    def __init__(__self__, *,
                 default_app: Optional[pulumi.Input[_builtins.str]] = None,
                 email: Optional[pulumi.Input[_builtins.str]] = None,
                 force_change_pass: Optional[pulumi.Input[_builtins.bool]] = None,
                 name: Optional[pulumi.Input[_builtins.str]] = None,
                 password: Optional[pulumi.Input[_builtins.str]] = None,
                 realname: Optional[pulumi.Input[_builtins.str]] = None,
                 restart_background_jobs: Optional[pulumi.Input[_builtins.bool]] = None,
                 roles: Optional[pulumi.Input[Sequence[pulumi.Input[_builtins.str]]]] = None,
                 tz: Optional[pulumi.Input[_builtins.str]] = None):
        """
        Input properties used for looking up and filtering AuthenticationUsers resources.
        :param pulumi.Input[_builtins.str] default_app: User default app. Overrides the default app inherited from the user roles.
        :param pulumi.Input[_builtins.str] email: User email address.
        :param pulumi.Input[_builtins.bool] force_change_pass: Force user to change password indication
        :param pulumi.Input[_builtins.str] name: Unique user login name.
        :param pulumi.Input[_builtins.str] password: User login password.
        :param pulumi.Input[_builtins.str] realname: Full user name.
        :param pulumi.Input[_builtins.bool] restart_background_jobs: Restart background search job that has not completed when Splunk restarts indication.
        :param pulumi.Input[Sequence[pulumi.Input[_builtins.str]]] roles: Role to assign to this user. At least one existing role is required.
        :param pulumi.Input[_builtins.str] tz: User timezone.
        """
        if default_app is not None:
            pulumi.set(__self__, "default_app", default_app)
        if email is not None:
            pulumi.set(__self__, "email", email)
        if force_change_pass is not None:
            pulumi.set(__self__, "force_change_pass", force_change_pass)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if password is not None:
            pulumi.set(__self__, "password", password)
        if realname is not None:
            pulumi.set(__self__, "realname", realname)
        if restart_background_jobs is not None:
            pulumi.set(__self__, "restart_background_jobs", restart_background_jobs)
        if roles is not None:
            pulumi.set(__self__, "roles", roles)
        if tz is not None:
            pulumi.set(__self__, "tz", tz)

    @_builtins.property
    @pulumi.getter(name="defaultApp")
    def default_app(self) -> Optional[pulumi.Input[_builtins.str]]:
        """
        User default app. Overrides the default app inherited from the user roles.
        """
        return pulumi.get(self, "default_app")

    @default_app.setter
    def default_app(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "default_app", value)

    @_builtins.property
    @pulumi.getter
    def email(self) -> Optional[pulumi.Input[_builtins.str]]:
        """
        User email address.
        """
        return pulumi.get(self, "email")

    @email.setter
    def email(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "email", value)

    @_builtins.property
    @pulumi.getter(name="forceChangePass")
    def force_change_pass(self) -> Optional[pulumi.Input[_builtins.bool]]:
        """
        Force user to change password indication
        """
        return pulumi.get(self, "force_change_pass")

    @force_change_pass.setter
    def force_change_pass(self, value: Optional[pulumi.Input[_builtins.bool]]):
        pulumi.set(self, "force_change_pass", value)

    @_builtins.property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[_builtins.str]]:
        """
        Unique user login name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "name", value)

    @_builtins.property
    @pulumi.getter
    def password(self) -> Optional[pulumi.Input[_builtins.str]]:
        """
        User login password.
        """
        return pulumi.get(self, "password")

    @password.setter
    def password(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "password", value)

    @_builtins.property
    @pulumi.getter
    def realname(self) -> Optional[pulumi.Input[_builtins.str]]:
        """
        Full user name.
        """
        return pulumi.get(self, "realname")

    @realname.setter
    def realname(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "realname", value)

    @_builtins.property
    @pulumi.getter(name="restartBackgroundJobs")
    def restart_background_jobs(self) -> Optional[pulumi.Input[_builtins.bool]]:
        """
        Restart background search job that has not completed when Splunk restarts indication.
        """
        return pulumi.get(self, "restart_background_jobs")

    @restart_background_jobs.setter
    def restart_background_jobs(self, value: Optional[pulumi.Input[_builtins.bool]]):
        pulumi.set(self, "restart_background_jobs", value)

    @_builtins.property
    @pulumi.getter
    def roles(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[_builtins.str]]]]:
        """
        Role to assign to this user. At least one existing role is required.
        """
        return pulumi.get(self, "roles")

    @roles.setter
    def roles(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[_builtins.str]]]]):
        pulumi.set(self, "roles", value)

    @_builtins.property
    @pulumi.getter
    def tz(self) -> Optional[pulumi.Input[_builtins.str]]:
        """
        User timezone.
        """
        return pulumi.get(self, "tz")

    @tz.setter
    def tz(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "tz", value)


@pulumi.type_token("splunk:index/authenticationUsers:AuthenticationUsers")
class AuthenticationUsers(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 default_app: Optional[pulumi.Input[_builtins.str]] = None,
                 email: Optional[pulumi.Input[_builtins.str]] = None,
                 force_change_pass: Optional[pulumi.Input[_builtins.bool]] = None,
                 name: Optional[pulumi.Input[_builtins.str]] = None,
                 password: Optional[pulumi.Input[_builtins.str]] = None,
                 realname: Optional[pulumi.Input[_builtins.str]] = None,
                 restart_background_jobs: Optional[pulumi.Input[_builtins.bool]] = None,
                 roles: Optional[pulumi.Input[Sequence[pulumi.Input[_builtins.str]]]] = None,
                 tz: Optional[pulumi.Input[_builtins.str]] = None,
                 __props__=None):
        """
        ## # Resource: AuthenticationUsers

        Create and update user information or delete the user.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_splunk as splunk

        user01 = splunk.AuthenticationUsers("user01",
            name="user01",
            email="user01@example.com",
            password="password01",
            force_change_pass=False,
            roles=["terraform-user01-role"])
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[_builtins.str] default_app: User default app. Overrides the default app inherited from the user roles.
        :param pulumi.Input[_builtins.str] email: User email address.
        :param pulumi.Input[_builtins.bool] force_change_pass: Force user to change password indication
        :param pulumi.Input[_builtins.str] name: Unique user login name.
        :param pulumi.Input[_builtins.str] password: User login password.
        :param pulumi.Input[_builtins.str] realname: Full user name.
        :param pulumi.Input[_builtins.bool] restart_background_jobs: Restart background search job that has not completed when Splunk restarts indication.
        :param pulumi.Input[Sequence[pulumi.Input[_builtins.str]]] roles: Role to assign to this user. At least one existing role is required.
        :param pulumi.Input[_builtins.str] tz: User timezone.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[AuthenticationUsersArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        ## # Resource: AuthenticationUsers

        Create and update user information or delete the user.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_splunk as splunk

        user01 = splunk.AuthenticationUsers("user01",
            name="user01",
            email="user01@example.com",
            password="password01",
            force_change_pass=False,
            roles=["terraform-user01-role"])
        ```

        :param str resource_name: The name of the resource.
        :param AuthenticationUsersArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(AuthenticationUsersArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 default_app: Optional[pulumi.Input[_builtins.str]] = None,
                 email: Optional[pulumi.Input[_builtins.str]] = None,
                 force_change_pass: Optional[pulumi.Input[_builtins.bool]] = None,
                 name: Optional[pulumi.Input[_builtins.str]] = None,
                 password: Optional[pulumi.Input[_builtins.str]] = None,
                 realname: Optional[pulumi.Input[_builtins.str]] = None,
                 restart_background_jobs: Optional[pulumi.Input[_builtins.bool]] = None,
                 roles: Optional[pulumi.Input[Sequence[pulumi.Input[_builtins.str]]]] = None,
                 tz: Optional[pulumi.Input[_builtins.str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = AuthenticationUsersArgs.__new__(AuthenticationUsersArgs)

            __props__.__dict__["default_app"] = default_app
            __props__.__dict__["email"] = email
            __props__.__dict__["force_change_pass"] = force_change_pass
            __props__.__dict__["name"] = name
            __props__.__dict__["password"] = None if password is None else pulumi.Output.secret(password)
            __props__.__dict__["realname"] = realname
            __props__.__dict__["restart_background_jobs"] = restart_background_jobs
            __props__.__dict__["roles"] = roles
            __props__.__dict__["tz"] = tz
        secret_opts = pulumi.ResourceOptions(additional_secret_outputs=["password"])
        opts = pulumi.ResourceOptions.merge(opts, secret_opts)
        super(AuthenticationUsers, __self__).__init__(
            'splunk:index/authenticationUsers:AuthenticationUsers',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            default_app: Optional[pulumi.Input[_builtins.str]] = None,
            email: Optional[pulumi.Input[_builtins.str]] = None,
            force_change_pass: Optional[pulumi.Input[_builtins.bool]] = None,
            name: Optional[pulumi.Input[_builtins.str]] = None,
            password: Optional[pulumi.Input[_builtins.str]] = None,
            realname: Optional[pulumi.Input[_builtins.str]] = None,
            restart_background_jobs: Optional[pulumi.Input[_builtins.bool]] = None,
            roles: Optional[pulumi.Input[Sequence[pulumi.Input[_builtins.str]]]] = None,
            tz: Optional[pulumi.Input[_builtins.str]] = None) -> 'AuthenticationUsers':
        """
        Get an existing AuthenticationUsers resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[_builtins.str] default_app: User default app. Overrides the default app inherited from the user roles.
        :param pulumi.Input[_builtins.str] email: User email address.
        :param pulumi.Input[_builtins.bool] force_change_pass: Force user to change password indication
        :param pulumi.Input[_builtins.str] name: Unique user login name.
        :param pulumi.Input[_builtins.str] password: User login password.
        :param pulumi.Input[_builtins.str] realname: Full user name.
        :param pulumi.Input[_builtins.bool] restart_background_jobs: Restart background search job that has not completed when Splunk restarts indication.
        :param pulumi.Input[Sequence[pulumi.Input[_builtins.str]]] roles: Role to assign to this user. At least one existing role is required.
        :param pulumi.Input[_builtins.str] tz: User timezone.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _AuthenticationUsersState.__new__(_AuthenticationUsersState)

        __props__.__dict__["default_app"] = default_app
        __props__.__dict__["email"] = email
        __props__.__dict__["force_change_pass"] = force_change_pass
        __props__.__dict__["name"] = name
        __props__.__dict__["password"] = password
        __props__.__dict__["realname"] = realname
        __props__.__dict__["restart_background_jobs"] = restart_background_jobs
        __props__.__dict__["roles"] = roles
        __props__.__dict__["tz"] = tz
        return AuthenticationUsers(resource_name, opts=opts, __props__=__props__)

    @_builtins.property
    @pulumi.getter(name="defaultApp")
    def default_app(self) -> pulumi.Output[_builtins.str]:
        """
        User default app. Overrides the default app inherited from the user roles.
        """
        return pulumi.get(self, "default_app")

    @_builtins.property
    @pulumi.getter
    def email(self) -> pulumi.Output[_builtins.str]:
        """
        User email address.
        """
        return pulumi.get(self, "email")

    @_builtins.property
    @pulumi.getter(name="forceChangePass")
    def force_change_pass(self) -> pulumi.Output[Optional[_builtins.bool]]:
        """
        Force user to change password indication
        """
        return pulumi.get(self, "force_change_pass")

    @_builtins.property
    @pulumi.getter
    def name(self) -> pulumi.Output[_builtins.str]:
        """
        Unique user login name.
        """
        return pulumi.get(self, "name")

    @_builtins.property
    @pulumi.getter
    def password(self) -> pulumi.Output[Optional[_builtins.str]]:
        """
        User login password.
        """
        return pulumi.get(self, "password")

    @_builtins.property
    @pulumi.getter
    def realname(self) -> pulumi.Output[_builtins.str]:
        """
        Full user name.
        """
        return pulumi.get(self, "realname")

    @_builtins.property
    @pulumi.getter(name="restartBackgroundJobs")
    def restart_background_jobs(self) -> pulumi.Output[_builtins.bool]:
        """
        Restart background search job that has not completed when Splunk restarts indication.
        """
        return pulumi.get(self, "restart_background_jobs")

    @_builtins.property
    @pulumi.getter
    def roles(self) -> pulumi.Output[Sequence[_builtins.str]]:
        """
        Role to assign to this user. At least one existing role is required.
        """
        return pulumi.get(self, "roles")

    @_builtins.property
    @pulumi.getter
    def tz(self) -> pulumi.Output[_builtins.str]:
        """
        User timezone.
        """
        return pulumi.get(self, "tz")

