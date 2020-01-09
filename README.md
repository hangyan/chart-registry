# ChartRegistry

[![Go Report Card](https://goreportcard.com/badge/github.com/hangyan/chart-registry)](https://goreportcard.com/report/github.com/hangyan/chart-registry)


*ChartRegistry* is an open-source **[Helm Chart Repository](https://github.com/helm/helm-www/blob/master/content/docs/topics/chart_repository.md)** server written in Go (Golang), mainly use OCI registry as storage backend. 

Powered by some great Go technology:
- [helm/helm](https://github.com/helm/helm) - for working with charts and oci storage
- [chartmuseum/chartmuseum](https://github.com/chartmuseum/chartmuseum) - for api 




### Using Docker

We will need the following things to start:

* A docker network 
* An OCI registry, which will act as the storage backend for chart-registry
* ChartRegistry Service


```bash
docker network create registry

docker run -d --restart=always --name=registry --network=registry \
	-e REGISTRY_HTTP_ADDR=0.0.0.0:443 \
	-e REGISTRY_HTTP_TLS_CERTIFICATE=/certs/domain.crt \
	-e REGISTRY_HTTP_TLS_KEY=/certs/domain.key \
	-p 443:443 \
	hangyan/https-registry:2

docker run -d -p 8080:8080 --restart=always --name=chart-registry  \
	-e DEBUG=1 -e STORAGE=registry -e STORAGE_REGISTRY_REPO=registry \
	--network=registry hangyan/chart-registry:latest
```

Then, we can use a Helm(2 or 3) client to fetch/upload charts in this repo


```bash
helm repo add oci http://127.0.0.1:8080
helm repo update

# create a simple chart to test
helm create simple-pod
helm package simple-pod
curl -v --fail -F chart=@simple-0.1.0.tgz http://127.0.0.1:8080/api/charts

# update and fetch
helm repo update
helm fetch oci/simple-pod

```


For other functionality, please check the [chartmuseum/chartmuseum](https://github.com/chartmuseum/chartmuseum) project.
