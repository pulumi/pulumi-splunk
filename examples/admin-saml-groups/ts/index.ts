import * as pulumi from "@pulumi/pulumi";
import * as splunk from "@pulumi/splunk";

const adminSamlGroup = new splunk.AdminSamlGroups("demo-ts-group", {
    roles: ["admin", "power"]
});

export const id = adminSamlGroup.id;
