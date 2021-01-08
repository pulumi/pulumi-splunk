module github.com/pulumi/pulumi-splunk/provider

go 1.15

replace github.com/Azure/go-autorest => github.com/Azure/go-autorest v12.4.3+incompatible

require (
	github.com/hashicorp/terraform-plugin-sdk v1.15.0
	github.com/pulumi/pulumi-terraform-bridge/v2 v2.16.0
	github.com/pulumi/pulumi/sdk/v2 v2.16.0
	github.com/splunk/terraform-provider-splunk v1.3.9
)
