/*
Copyright The Helm Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package config

import (
	"github.com/urfave/cli"
)

type (
	configVar struct {
		Type    configVarType
		Default interface{}
		CLIFlag cli.Flag
	}

	configVarType string
)

// Will be populated in init() below
var CLIFlags []cli.Flag

var (
	stringType configVarType = "string"
	intType    configVarType = "int"
	boolType   configVarType = "bool"
)

var configVars = map[string]configVar{
	"genindex": {
		Type:    boolType,
		Default: false,
		CLIFlag: cli.BoolFlag{
			Name:   "gen-index",
			Usage:  "generate index.yaml, print to stdout and exit",
			EnvVar: "GEN_INDEX",
		},
	},
	"debug": {
		Type:    boolType,
		Default: false,
		CLIFlag: cli.BoolFlag{
			Name:   "debug",
			Usage:  "show debug messages",
			EnvVar: "DEBUG",
		},
	},
	"logjson": {
		Type:    boolType,
		Default: false,
		CLIFlag: cli.BoolFlag{
			Name:   "log-json",
			Usage:  "output structured logs as json",
			EnvVar: "LOG_JSON",
		},
	},
	"loghealth": {
		Type:    boolType,
		Default: false,
		CLIFlag: cli.BoolFlag{
			Name:   "log-health",
			Usage:  "log inbound /health requests",
			EnvVar: "LOG_HEALTH",
		},
	},
	"disablemetrics": {
		Type:    boolType,
		Default: false,
		CLIFlag: cli.BoolFlag{
			Name:   "disable-metrics",
			Usage:  "disable Prometheus metrics",
			EnvVar: "DISABLE_METRICS",
		},
	},
	"disableapi": {
		Type:    boolType,
		Default: false,
		CLIFlag: cli.BoolFlag{
			Name:   "disable-api",
			Usage:  "disable all routes prefixed with /api",
			EnvVar: "DISABLE_API",
		},
	},
	"disabldelete": {
		Type:    boolType,
		Default: false,
		CLIFlag: cli.BoolFlag{
			Name:   "disable-delete",
			Usage:  "disable DELETE route",
			EnvVar: "DISABLE_DELETE",
		},
	},
	"disablestatefiles": {
		Type:    boolType,
		Default: false,
		CLIFlag: cli.BoolFlag{
			Name:   "disable-statefiles",
			Usage:  "disable use of index-cache.yaml",
			EnvVar: "DISABLE_STATEFILES",
		},
	},
	"allowoverwrite": {
		Type:    boolType,
		Default: false,
		CLIFlag: cli.BoolFlag{
			Name:   "allow-overwrite",
			Usage:  "allow chart versions to be re-uploaded without ?force querystring",
			EnvVar: "ALLOW_OVERWRITE",
		},
	},
	"disableforceoverwrite": {
		Type:    boolType,
		Default: false,
		CLIFlag: cli.BoolFlag{
			Name:   "disable-force-overwrite",
			Usage:  "do not allow chart versions to be re-uploaded, even with ?force querystring",
			EnvVar: "DISABLE_FORCE_OVERWRITE",
		},
	},
	"port": {
		Type:    intType,
		Default: 8080,
		CLIFlag: cli.IntFlag{
			Name:   "port",
			Usage:  "port to listen on",
			EnvVar: "PORT",
		},
	},
	"charturl": {
		Type:    stringType,
		Default: "",
		CLIFlag: cli.StringFlag{
			Name:   "chart-url",
			Usage:  "absolute url for .tgzs in index.yaml",
			EnvVar: "CHART_URL",
		},
	},
	"basicauth.user": {
		Type:    stringType,
		Default: "",
		CLIFlag: cli.StringFlag{
			Name:   "basic-auth-user",
			Usage:  "username for basic http authentication",
			EnvVar: "BASIC_AUTH_USER",
		},
	},
	"basicauth.pass": {
		Type:    stringType,
		Default: "",
		CLIFlag: cli.StringFlag{
			Name:   "basic-auth-pass",
			Usage:  "password for basic http authentication",
			EnvVar: "BASIC_AUTH_PASS",
		},
	},
	"authanonymousget": {
		Type:    boolType,
		Default: false,
		CLIFlag: cli.BoolFlag{
			Name:   "auth-anonymous-get",
			Usage:  "allow anonymous GET operations when auth is used",
			EnvVar: "AUTH_ANONYMOUS_GET",
		},
	},
	"tls.cert": {
		Type:    stringType,
		Default: "",
		CLIFlag: cli.StringFlag{
			Name:   "tls-cert",
			Usage:  "path to tls certificate chain file",
			EnvVar: "TLS_CERT",
		},
	},
	"tls.key": {
		Type:    stringType,
		Default: "",
		CLIFlag: cli.StringFlag{
			Name:   "tls-key",
			Usage:  "path to tls key file",
			EnvVar: "TLS_KEY",
		},
	},
	"tls.cacert": {
		Type:    stringType,
		Default: "",
		CLIFlag: cli.StringFlag{
			Name:   "tls-ca-cert",
			Usage:  "path to tls ca cert file",
			EnvVar: "TLS_CA_CERT",
		},
	},
	"cache.store": {
		Type:    stringType,
		Default: "",
		CLIFlag: cli.StringFlag{
			Name:   "cache",
			Usage:  "cache store, can be one of: redis",
			EnvVar: "CACHE",
		},
	},
	"cache.redis.addr": {
		Type:    stringType,
		Default: "",
		CLIFlag: cli.StringFlag{
			Name:   "cache-redis-addr",
			Usage:  "address of Redis service (host:port)",
			EnvVar: "CACHE_REDIS_ADDR",
		},
	},
	"cache.redis.password": {
		Type:    stringType,
		Default: "",
		CLIFlag: cli.StringFlag{
			Name:   "cache-redis-password",
			Usage:  "Redis requirepass server configuration",
			EnvVar: "CACHE_REDIS_PASSWORD",
		},
	},
	"cache.redis.db": {
		Type:    intType,
		Default: 0,
		CLIFlag: cli.IntFlag{
			Name:   "cache-redis-db",
			Usage:  "Redis database to be selected after connect",
			EnvVar: "CACHE_REDIS_DB",
			Value:  0,
		},
	},
	"storage.backend": {
		Type:    stringType,
		Default: "",
		CLIFlag: cli.StringFlag{
			Name:   "storage",
			Usage:  "storage backend, can be one of: local, amazon, google, oracle",
			EnvVar: "STORAGE",
		},
	},
	"storage.local.rootdir": {
		Type:    stringType,
		Default: "",
		CLIFlag: cli.StringFlag{
			Name:   "storage-local-rootdir",
			Usage:  "directory to store charts for local storage backend",
			EnvVar: "STORAGE_LOCAL_ROOTDIR",
		},
	},

	"storage.registry.repo": {
		Type:    stringType,
		Default: "",
		CLIFlag: cli.StringFlag{
			Name:   "storage-registry-repo",
			Usage:  "repo url for registry storage backend",
			EnvVar: "STORAGE_REGISTRY_REPO",
		},
	},


	"chartpostformfieldname": {
		Type:    stringType,
		Default: "chart",
		CLIFlag: cli.StringFlag{
			Name:   "chart-post-form-field-name",
			Usage:  "form field which will be queried for the chart file content",
			EnvVar: "CHART_POST_FORM_FIELD_NAME",
		},
	},
	"provpostformfieldname": {
		Type:    stringType,
		Default: "prov",
		CLIFlag: cli.StringFlag{
			Name:   "prov-post-form-field-name",
			Usage:  "form field which will be queried for the provenance file content",
			EnvVar: "PROV_POST_FORM_FIELD_NAME",
		},
	},
	"maxstorageobjects": {
		Type:    intType,
		Default: 0,
		CLIFlag: cli.IntFlag{
			Name:   "max-storage-objects",
			Usage:  "maximum number of objects allowed in storage (per tenant)",
			EnvVar: "MAX_STORAGE_OBJECTS",
		},
	},
	"maxuploadsize": {
		Type:    intType,
		Default: 1024 * 1024 * 20, // 20MB, per Helm's limit
		CLIFlag: cli.IntFlag{
			Name:   "max-upload-size",
			Usage:  "max size of post body (in bytes)",
			EnvVar: "MAX_UPLOAD_SIZE",
			Value:  1024 * 1024 * 20,
		},
	},
	"indexlimit": {
		Type:    intType,
		Default: 0,
		CLIFlag: cli.IntFlag{
			Name:   "index-limit",
			Usage:  "parallel scan limit for the repo indexer",
			EnvVar: "INDEX_LIMIT",
		},
	},
	"contextpath": {
		Type:    stringType,
		Default: "",
		CLIFlag: cli.StringFlag{
			Name:   "context-path",
			Usage:  "base context path",
			EnvVar: "CONTEXT_PATH",
		},
	},
	"depth": {
		Type:    intType,
		Default: 0,
		CLIFlag: cli.IntFlag{
			Name:   "depth",
			Usage:  "levels of nested repos for multitenancy",
			EnvVar: "DEPTH",
		},
	},
	"bearerauth": {
		Type:    boolType,
		Default: false,
		CLIFlag: cli.BoolFlag{
			Name:   "bearer-auth",
			Usage:  "enable bearer auth",
			EnvVar: "BEARER_AUTH",
		},
	},
	"authrealm": {
		Type:    stringType,
		Default: "",
		CLIFlag: cli.StringFlag{
			Name:   "auth-realm",
			Usage:  "authorization server url",
			EnvVar: "AUTH_REALM",
		},
	},
	"authservice": {
		Type:    stringType,
		Default: "",
		CLIFlag: cli.StringFlag{
			Name:   "auth-service",
			Usage:  "authorization server service name",
			EnvVar: "AUTH_SERVICE",
		},
	},
	"authcertpath": {
		Type:    stringType,
		Default: "",
		CLIFlag: cli.StringFlag{
			Name:   "auth-cert-path",
			Usage:  "path to authorization server public pem file",
			EnvVar: "AUTH_CERT_PATH",
		},
	},
	"depthdynamic": {
		Type:    boolType,
		Default: false,
		CLIFlag: cli.BoolFlag{
			Name:   "depth-dynamic",
			Usage:  "the length of repo variable",
			EnvVar: "DEPTH_DYNAMIC",
		},
	},
	"cors.alloworigin": {
		Type:    stringType,
		Default: "",
		CLIFlag: cli.StringFlag{
			Name:   "cors-alloworigin",
			Usage:  "value to set in the Access-Control-Allow-Origin HTTP header",
			EnvVar: "CORS_ALLOW_ORIGIN",
		},
	},
}

func populateCLIFlags() {
	CLIFlags = []cli.Flag{
		cli.StringFlag{
			Name:   "config, c",
			Usage:  "chartmuseum configuration file",
			EnvVar: "CONFIG",
		},
	}
	for _, configVar := range configVars {
		if flag := configVar.CLIFlag; flag != nil {
			CLIFlags = append(CLIFlags, flag)
		}
	}
}

func init() {
	populateCLIFlags()
}
