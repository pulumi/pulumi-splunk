[![Actions Status](https://github.com/pulumi/pulumi-splunk/workflows/master/badge.svg)](https://github.com/pulumi/pulumi-splunk/actions)
[![Slack](http://www.pulumi.com/images/docs/badges/slack.svg)](https://slack.pulumi.com)
[![NPM version](https://badge.fury.io/js/%40pulumi%2Fsplunk.svg)](https://www.npmjs.com/package/@pulumi/splunk)
[![Python version](https://badge.fury.io/py/pulumi-splunk.svg)](https://pypi.org/project/pulumi-splunk)
[![NuGet version](https://badge.fury.io/nu/pulumi.splunk.svg)](https://badge.fury.io/nu/pulumi.splunk)
[![PkgGoDev](https://pkg.go.dev/badge/github.com/pulumi/pulumi-splunk/sdk/go)](https://pkg.go.dev/github.com/pulumi/pulumi-splunk/sdk/go)
[![License](https://img.shields.io/npm/l/%40pulumi%2Fpulumi.svg)](https://github.com/pulumi/pulumi-splunk/blob/master/LICENSE)

# Splunk Resource Provider

The Splunk Resource Provider lets you manage Splunk resources.

## Installing

This package is available in many languages in the standard packaging formats.

### Node.js (Java/TypeScript)

To use from JavaScript or TypeScript in Node.js, install using either `npm`:

    $ npm install @pulumi/splunk

or `yarn`:

    $ yarn add @pulumi/splunk

### Python

To use from Python, install using `pip`:

    $ pip install pulumi_splunk

### Go

To use from Go, use `go get` to grab the latest version of the library

    $ go get github.com/pulumi/pulumi-splunk/sdk

### .NET

To use from .NET, install using `dotnet add package`:

    $ dotnet add package Pulumi.Splunk

## Configuration

The following configuration points are available:

- `splunk:url` - (Required) The URL for the Splunk instance to be configured. (The provider uses `https` as the default schema as prefix to the URL)
  It can also be sourced from the `SPLUNK_URL` environment variable.
- `splunk:username` - (Optional) The username to access the Splunk instance to be configured. It can also be sourced
  from the `SPLUNK_USERNAME` environment variable.
- `splunk:password` - (Optional) The password to access the Splunk instance to be configured. It can also be sourced 
  from the `SPLUNK_PASSWORD` environment variable.
- `splunk:authToken` - (Optional) Use auth token instead of username and password to configure Splunk instance. If specified, auth token takes priority over username/password.
  It can also be sourced  from the `SPLUNK_AUTH_TOKEN` environment variable.
- `splunk:insecureSkipVerify` -(Optional) Insecure skip verification flag (Defaults to `true`)
  It can also be sourced  from the `SPLUNK_INSECURE_SKIP_VERIFY` environment variable.
- `splunk:timeout` - (Optional) Timeout when making calls to Splunk server. (Defaults to `60 seconds`)
  It can also be sourced  from the `SPLUNK_TIMEOUT` environment variable.

## Reference

For further information, please visit [the Splunk provider docs](https://www.pulumi.com/docs/intro/cloud-providers/splunk)
or for detailed reference documentation, please visit [the API docs](https://www.pulumi.com/docs/reference/pkg/splunk).
