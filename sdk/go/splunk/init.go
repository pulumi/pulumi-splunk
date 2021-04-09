// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package splunk

import (
	"fmt"

	"github.com/blang/semver"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

type module struct {
	version semver.Version
}

func (m *module) Version() semver.Version {
	return m.version
}

func (m *module) Construct(ctx *pulumi.Context, name, typ, urn string) (r pulumi.Resource, err error) {
	switch typ {
	case "splunk:index/adminSamlGroups:AdminSamlGroups":
		r, err = NewAdminSamlGroups(ctx, name, nil, pulumi.URN_(urn))
	case "splunk:index/appsLocal:AppsLocal":
		r, err = NewAppsLocal(ctx, name, nil, pulumi.URN_(urn))
	case "splunk:index/authenticationUsers:AuthenticationUsers":
		r, err = NewAuthenticationUsers(ctx, name, nil, pulumi.URN_(urn))
	case "splunk:index/authorizationRoles:AuthorizationRoles":
		r, err = NewAuthorizationRoles(ctx, name, nil, pulumi.URN_(urn))
	case "splunk:index/configsConf:ConfigsConf":
		r, err = NewConfigsConf(ctx, name, nil, pulumi.URN_(urn))
	case "splunk:index/dataUiViews:DataUiViews":
		r, err = NewDataUiViews(ctx, name, nil, pulumi.URN_(urn))
	case "splunk:index/globalHttpEventCollector:GlobalHttpEventCollector":
		r, err = NewGlobalHttpEventCollector(ctx, name, nil, pulumi.URN_(urn))
	case "splunk:index/indexes:Indexes":
		r, err = NewIndexes(ctx, name, nil, pulumi.URN_(urn))
	case "splunk:index/inputsHttpEventCollector:InputsHttpEventCollector":
		r, err = NewInputsHttpEventCollector(ctx, name, nil, pulumi.URN_(urn))
	case "splunk:index/inputsMonitor:InputsMonitor":
		r, err = NewInputsMonitor(ctx, name, nil, pulumi.URN_(urn))
	case "splunk:index/inputsScript:InputsScript":
		r, err = NewInputsScript(ctx, name, nil, pulumi.URN_(urn))
	case "splunk:index/inputsTcpCooked:InputsTcpCooked":
		r, err = NewInputsTcpCooked(ctx, name, nil, pulumi.URN_(urn))
	case "splunk:index/inputsTcpRaw:InputsTcpRaw":
		r, err = NewInputsTcpRaw(ctx, name, nil, pulumi.URN_(urn))
	case "splunk:index/inputsTcpSplunkTcpToken:InputsTcpSplunkTcpToken":
		r, err = NewInputsTcpSplunkTcpToken(ctx, name, nil, pulumi.URN_(urn))
	case "splunk:index/inputsTcpSsl:InputsTcpSsl":
		r, err = NewInputsTcpSsl(ctx, name, nil, pulumi.URN_(urn))
	case "splunk:index/inputsUdp:InputsUdp":
		r, err = NewInputsUdp(ctx, name, nil, pulumi.URN_(urn))
	case "splunk:index/outputsTcpDefault:OutputsTcpDefault":
		r, err = NewOutputsTcpDefault(ctx, name, nil, pulumi.URN_(urn))
	case "splunk:index/outputsTcpGroup:OutputsTcpGroup":
		r, err = NewOutputsTcpGroup(ctx, name, nil, pulumi.URN_(urn))
	case "splunk:index/outputsTcpServer:OutputsTcpServer":
		r, err = NewOutputsTcpServer(ctx, name, nil, pulumi.URN_(urn))
	case "splunk:index/outputsTcpSyslog:OutputsTcpSyslog":
		r, err = NewOutputsTcpSyslog(ctx, name, nil, pulumi.URN_(urn))
	case "splunk:index/savedSearches:SavedSearches":
		r, err = NewSavedSearches(ctx, name, nil, pulumi.URN_(urn))
	case "splunk:index/shIndexesManager:ShIndexesManager":
		r, err = NewShIndexesManager(ctx, name, nil, pulumi.URN_(urn))
	default:
		return nil, fmt.Errorf("unknown resource type: %s", typ)
	}

	return
}

type pkg struct {
	version semver.Version
}

func (p *pkg) Version() semver.Version {
	return p.version
}

func (p *pkg) ConstructProvider(ctx *pulumi.Context, name, typ, urn string) (pulumi.ProviderResource, error) {
	if typ != "pulumi:providers:splunk" {
		return nil, fmt.Errorf("unknown provider type: %s", typ)
	}

	return NewProvider(ctx, name, nil, pulumi.URN_(urn))
}

func init() {
	version, err := PkgVersion()
	if err != nil {
		fmt.Println("failed to determine package version. defaulting to v1: %v", err)
	}
	pulumi.RegisterResourceModule(
		"splunk",
		"index/adminSamlGroups",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"splunk",
		"index/appsLocal",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"splunk",
		"index/authenticationUsers",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"splunk",
		"index/authorizationRoles",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"splunk",
		"index/configsConf",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"splunk",
		"index/dataUiViews",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"splunk",
		"index/globalHttpEventCollector",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"splunk",
		"index/indexes",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"splunk",
		"index/inputsHttpEventCollector",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"splunk",
		"index/inputsMonitor",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"splunk",
		"index/inputsScript",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"splunk",
		"index/inputsTcpCooked",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"splunk",
		"index/inputsTcpRaw",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"splunk",
		"index/inputsTcpSplunkTcpToken",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"splunk",
		"index/inputsTcpSsl",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"splunk",
		"index/inputsUdp",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"splunk",
		"index/outputsTcpDefault",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"splunk",
		"index/outputsTcpGroup",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"splunk",
		"index/outputsTcpServer",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"splunk",
		"index/outputsTcpSyslog",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"splunk",
		"index/savedSearches",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"splunk",
		"index/shIndexesManager",
		&module{version},
	)
	pulumi.RegisterResourcePackage(
		"splunk",
		&pkg{version},
	)
}