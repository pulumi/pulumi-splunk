"""A Python Pulumi program"""

import pulumi
import pulumi_splunk as splunk

saml_group = splunk.AdminSamlGroups("demo-py-group", roles=[
    "admin",
    "power",
])
