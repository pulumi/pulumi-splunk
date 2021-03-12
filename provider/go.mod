module github.com/pulumi/pulumi-splunk/provider

go 1.15

replace github.com/Azure/go-autorest => github.com/Azure/go-autorest v12.4.3+incompatible

require (
	github.com/hashicorp/terraform-plugin-sdk v1.15.0
	github.com/pulumi/pulumi-terraform-bridge/v2 v2.19.0
	github.com/pulumi/pulumi/sdk/v2 v2.20.1-0.20210212181059-f4b0fa86fedc
	github.com/splunk/terraform-provider-splunk v1.4.0
)
