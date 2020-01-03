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

package main

import (
	"fmt"
	"github.com/hangyan/chart-registry/pkg/storage"
	"log"
	"os"
	"strings"

	"github.com/hangyan/chart-registry/pkg/cache"
	"github.com/hangyan/chart-registry/pkg/chartmuseum"
	"github.com/hangyan/chart-registry/pkg/config"

	"github.com/urfave/cli"
)

var (
	crash = log.Fatal

	newServer = chartmuseum.NewServer

	// Version is the semantic version (added at compile time)
	Version string

	// Revision is the git commit id (added at compile time)
	Revision string
)

func main() {
	app := cli.NewApp()
	app.Name = "ChartRegistry"
	app.Version = fmt.Sprintf("%s (build %s)", Version, Revision)
	app.Usage = "Helm Chart Repository Hosted on Docker Registry"
	app.Action = cliHandler
	app.Flags = config.CLIFlags
	app.Run(os.Args)
}

func cliHandler(c *cli.Context) {
	conf := config.NewConfig()
	err := conf.UpdateFromCLIContext(c)
	if err != nil {
		crash(err)
	}

	backend := backendFromConfig(conf)
	store := storeFromConfig(conf)

	options := chartmuseum.ServerOptions{
		StorageBackend:         backend,
		ExternalCacheStore:     store,
		ChartURL:               conf.GetString("charturl"),
		TlsCert:                conf.GetString("tls.cert"),
		TlsKey:                 conf.GetString("tls.key"),
		TlsCACert:              conf.GetString("tls.cacert"),
		Username:               conf.GetString("basicauth.user"),
		Password:               conf.GetString("basicauth.pass"),
		ChartPostFormFieldName: conf.GetString("chartpostformfieldname"),
		ProvPostFormFieldName:  conf.GetString("provpostformfieldname"),
		ContextPath:            conf.GetString("contextpath"),
		LogJSON:                conf.GetBool("logjson"),
		LogHealth:              conf.GetBool("loghealth"),
		Debug:                  conf.GetBool("debug"),
		EnableAPI:              !conf.GetBool("disableapi"),
		DisableDelete:          conf.GetBool("disabledelete"),
		UseStatefiles:          !conf.GetBool("disablestatefiles"),
		AllowOverwrite:         conf.GetBool("allowoverwrite"),
		AllowForceOverwrite:    !conf.GetBool("disableforceoverwrite"),
		EnableMetrics:          !conf.GetBool("disablemetrics"),
		AnonymousGet:           conf.GetBool("authanonymousget"),
		GenIndex:               conf.GetBool("genindex"),
		MaxStorageObjects:      conf.GetInt("maxstorageobjects"),
		IndexLimit:             conf.GetInt("indexlimit"),
		Depth:                  conf.GetInt("depth"),
		MaxUploadSize:          conf.GetInt("maxuploadsize"),
		BearerAuth:             conf.GetBool("bearerauth"),
		AuthRealm:              conf.GetString("authrealm"),
		AuthService:            conf.GetString("authservice"),
		AuthCertPath:           conf.GetString("authcertpath"),
		DepthDynamic:           conf.GetBool("depthdynamic"),
		CORSAllowOrigin:        conf.GetString("cors.alloworigin"),
	}

	server, err := newServer(options)
	if err != nil {
		crash(err)
	}

	server.Listen(conf.GetInt("port"))
}

func backendFromConfig(conf *config.Config) storage.Backend {
	crashIfConfigMissingVars(conf, []string{"storage.backend"})

	var backend storage.Backend

	storageFlag := strings.ToLower(conf.GetString("storage.backend"))
	switch storageFlag {
	case "local":
		backend = localBackendFromConfig(conf)

	default:
		crash("Unsupported storage backend: ", storageFlag)
	}

	return backend
}

func localBackendFromConfig(conf *config.Config) storage.Backend {
	crashIfConfigMissingVars(conf, []string{"storage.local.rootdir"})
	return storage.Backend(storage.NewLocalFilesystemBackend(
		conf.GetString("storage.local.rootdir"),
	))
}

func storeFromConfig(conf *config.Config) cache.Store {
	if conf.GetString("cache.store") == "" {
		return nil
	}

	var store cache.Store

	cacheFlag := strings.ToLower(conf.GetString("cache.store"))
	switch cacheFlag {
	case "redis":
		store = redisCacheFromConfig(conf)
	default:
		crash("Unsupported cache store: ", cacheFlag)
	}

	return store
}

func redisCacheFromConfig(conf *config.Config) cache.Store {
	crashIfConfigMissingVars(conf, []string{"cache.redis.addr"})
	return cache.Store(cache.NewRedisStore(
		conf.GetString("cache.redis.addr"),
		conf.GetString("cache.redis.password"),
		conf.GetInt("cache.redis.db"),
	))
}

func crashIfConfigMissingVars(conf *config.Config, vars []string) {
	missing := []string{}
	for _, v := range vars {
		if conf.GetString(v) == "" {
			flag := config.GetCLIFlagFromVarName(v)
			missing = append(missing, fmt.Sprintf("--%s", flag))
		}
	}
	if len(missing) > 0 {
		crash("Missing required flags(s): ", strings.Join(missing, ", "))
	}
}
