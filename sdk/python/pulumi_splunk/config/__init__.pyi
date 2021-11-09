# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

authToken: Optional[str]
"""
Authentication tokens, also known as JSON Web Tokens (JWT), are a method for authenticating Splunk platform users into
the Splunk platform
"""

insecureSkipVerify: Optional[bool]
"""
insecure skip verification flag
"""

password: Optional[str]
"""
Splunk instance password
"""

timeout: Optional[int]
"""
Timeout when making calls to Splunk server. Defaults to 60 seconds
"""

url: Optional[str]
"""
Splunk instance URL
"""

username: Optional[str]
"""
Splunk instance admin username
"""

