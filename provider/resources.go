// Copyright 2016-2018, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package splunk

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"unicode"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/splunk/terraform-provider-splunk/splunk"

	"github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfbridge"
	shimv1 "github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfshim/sdk-v1"
	"github.com/pulumi/pulumi/sdk/v3/go/common/tokens"

	"github.com/pulumi/pulumi-splunk/provider/pkg/version"
)

// all of the token components used below.
const (
	// packages:
	mainPkg = "splunk"
	// modules:
	mainMod = "index" // the main module
)

// makeMember manufactures a type token for the package and the given module and type.
func makeMember(mod string, mem string) tokens.ModuleMember {
	return tokens.ModuleMember(mainPkg + ":" + mod + ":" + mem)
}

// makeType manufactures a type token for the package and the given module and type.
func makeType(mod string, typ string) tokens.Type {
	return tokens.Type(makeMember(mod, typ))
}

// makeDataSource manufactures a standard resource token given a module and resource name.  It
// automatically uses the main package and names the file by simply lower casing the data source's
// first character.
//func makeDataSource(mod string, res string) tokens.ModuleMember {
//	fn := string(unicode.ToLower(rune(res[0]))) + res[1:]
//	return makeMember(mod+"/"+fn, res)
//}

// makeResource manufactures a standard resource token given a module and resource name.  It
// automatically uses the main package and names the file by simply lower casing the resource's
// first character.
func makeResource(mod string, res string) tokens.Type {
	fn := string(unicode.ToLower(rune(res[0]))) + res[1:]
	return makeType(mod+"/"+fn, res)
}

// Provider returns additional overlaid schema and metadata associated with the provider..
func Provider() tfbridge.ProviderInfo {
	p := shimv1.NewProvider(splunk.Provider().(*schema.Provider))

	prov := tfbridge.ProviderInfo{
		P:           p,
		Name:        "splunk",
		Description: "A Pulumi package for creating and managing splunk cloud resources.",
		Keywords:    []string{"pulumi", "splunk"},
		License:     "Apache-2.0",
		Homepage:    "https://pulumi.io",
		Repository:  "https://github.com/pulumi/pulumi-splunk",
		GitHubOrg:   "splunk",
		Config:      map[string]*tfbridge.SchemaInfo{},
		DocRules:    &tfbridge.DocRuleInfo{EditRules: docEditRules},
		Resources: map[string]*tfbridge.ResourceInfo{
			"splunk_admin_saml_groups": {
				Tok: makeResource(mainMod, "AdminSamlGroups"),
				Docs: &tfbridge.DocInfo{
					Source: "admin_saml_groups.md",
				},
			},
			"splunk_apps_local": {
				Tok: makeResource(mainMod, "AppsLocal"),
				Docs: &tfbridge.DocInfo{
					Source: "apps_local.md",
				},
			},
			"splunk_authentication_users": {
				Tok: makeResource(mainMod, "AuthenticationUsers"),
				Docs: &tfbridge.DocInfo{
					Source: "authentication_users.md",
				},
			},
			"splunk_authorization_roles": {
				Tok: makeResource(mainMod, "AuthorizationRoles"),
				Docs: &tfbridge.DocInfo{
					Source: "authorization_roles.md",
				},
			},
			"splunk_global_http_event_collector": {
				Tok: makeResource(mainMod, "GlobalHttpEventCollector"),
				Docs: &tfbridge.DocInfo{
					Source: "global_http_event_collector.md",
				},
			},
			"splunk_inputs_http_event_collector": {
				Tok: makeResource(mainMod, "InputsHttpEventCollector"),
				Docs: &tfbridge.DocInfo{
					Source: "inputs_http_event_collector.md",
				},
			},
			"splunk_inputs_script": {
				Tok: makeResource(mainMod, "InputsScript"),
				Docs: &tfbridge.DocInfo{
					Source: "inputs_script.md",
				},
			},
			"splunk_inputs_monitor": {
				Tok: makeResource(mainMod, "InputsMonitor"),
				Docs: &tfbridge.DocInfo{
					Source: "inputs_monitor.md",
				},
			},
			"splunk_inputs_udp": {
				Tok: makeResource(mainMod, "InputsUdp"),
				Docs: &tfbridge.DocInfo{
					Source: "inputs_udp.md",
				},
			},
			"splunk_inputs_tcp_raw": {
				Tok: makeResource(mainMod, "InputsTcpRaw"),
				Docs: &tfbridge.DocInfo{
					Source: "inputs_tcp_raw.md",
				},
			},
			"splunk_inputs_tcp_cooked": {
				Tok: makeResource(mainMod, "InputsTcpCooked"),
				Docs: &tfbridge.DocInfo{
					Source: "inputs_tcp_cooked.md",
				},
			},
			"splunk_inputs_tcp_splunk_tcp_token": {
				Tok: makeResource(mainMod, "InputsTcpSplunkTcpToken"),
				Docs: &tfbridge.DocInfo{
					Source: "inputs_tcp_splunktcptoken.md",
				},
			},
			"splunk_inputs_tcp_ssl": {
				Tok: makeResource(mainMod, "InputsTcpSsl"),
				Docs: &tfbridge.DocInfo{
					Source: "inputs_tcp_ssl.md",
				},
			},
			"splunk_outputs_tcp_default": {
				Tok: makeResource(mainMod, "OutputsTcpDefault"),
				Docs: &tfbridge.DocInfo{
					Source: "outputs_tcp_default.md",
				},
			},
			"splunk_outputs_tcp_server": {
				Tok: makeResource(mainMod, "OutputsTcpServer"),
				Docs: &tfbridge.DocInfo{
					Source: "outputs_tcp_server.md",
				},
			},
			"splunk_outputs_tcp_group": {
				Tok: makeResource(mainMod, "OutputsTcpGroup"),
				Docs: &tfbridge.DocInfo{
					Source: "outputs_tcp_group.md",
				},
			},
			"splunk_outputs_tcp_syslog": {
				Tok: makeResource(mainMod, "OutputsTcpSyslog"),
				Docs: &tfbridge.DocInfo{
					Source: "outputs_tcp_syslog.md",
				},
			},
			"splunk_saved_searches": {
				Tok: makeResource(mainMod, "SavedSearches"),
				Docs: &tfbridge.DocInfo{
					Source: "saved_searches.md",
				},
			},
			"splunk_indexes": {
				Tok: makeResource(mainMod, "Indexes"),
				Docs: &tfbridge.DocInfo{
					Source: "indexes.md",
				},
			},
			"splunk_configs_conf": {
				Tok: makeResource(mainMod, "ConfigsConf"),
				Docs: &tfbridge.DocInfo{
					Source: "configs_conf.md",
				},
			},
			"splunk_data_ui_views":      {Tok: makeResource(mainMod, "DataUiViews")},
			"splunk_sh_indexes_manager": {Tok: makeResource(mainMod, "ShIndexesManager")},
			"splunk_generic_acl":        {Tok: makeResource(mainMod, "GenericAcl")},
		},
		DataSources: map[string]*tfbridge.DataSourceInfo{},
		JavaScript: &tfbridge.JavaScriptInfo{
			Dependencies: map[string]string{
				"@pulumi/pulumi": "^3.0.0-alpha.0",
			},
			DevDependencies: map[string]string{
				"@types/node": "^10.0.0", // so we can access strongly typed node definitions.
				"@types/mime": "^2.0.0",
			},
			RespectSchemaVersion: true,
		},
		Python: (func() *tfbridge.PythonInfo {
			i := &tfbridge.PythonInfo{
				RespectSchemaVersion: true,
				Requires: map[string]string{
					"pulumi": ">=3.0.0a1,<4.0.0",
				},
			}
			i.PyProject.Enabled = true
			return i
		})(),

		Golang: &tfbridge.GolangInfo{
			ImportBasePath: filepath.Join(
				fmt.Sprintf("github.com/pulumi/pulumi-%[1]s/sdk/", mainPkg),
				tfbridge.GetModuleMajorVersion(version.Version),
				"go",
				mainPkg,
			),
			GenerateResourceContainerTypes: true,
			RespectSchemaVersion:           true,
		},
		CSharp: &tfbridge.CSharpInfo{
			RespectSchemaVersion: true,
			PackageReferences: map[string]string{
				"Pulumi": "3.*-*",
			},
		},
	}

	prov.SetAutonaming(255, "-")

	return prov
}
func docEditRules(defaults []tfbridge.DocsEdit) []tfbridge.DocsEdit {
	return append(
		defaults,
		cleanUpExample,
	)
}

var cleanUpExample = tfbridge.DocsEdit{
	Path: "index.md",
	Edit: func(_ string, content []byte) ([]byte, error) {
		replacesDir := "docs/index-md-replaces/"

		input, err := os.ReadFile(replacesDir + "examples-input.md")
		if err != nil {
			return nil, err
		}
		desired, err := os.ReadFile(replacesDir + "examples-desired.md")
		if err != nil {
			return nil, err
		}
		if bytes.Contains(content, input) {
			content = bytes.ReplaceAll(
				content,
				input,
				desired)
		} else {
			// Hard error to ensure we keep this content up to date
			return nil, fmt.Errorf("could not find text in upstream index.md, "+
				"please verify file content at %s\n*****\n%s\n*****\n", replacesDir+"overview-input.md", string(input))
		}
		return content, nil
	},
}
