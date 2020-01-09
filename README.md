# ChartRegistry

[![Go Report Card](https://goreportcard.com/badge/github.com/hangyan/chart-registry)](https://goreportcard.com/report/github.com/hangyan/chart-registry)



*ChartRegistry* is an open-source **[Helm Chart Repository](https://github.com/helm/helm-www/blob/master/content/docs/topics/chart_repository.md)** server written in Go (Golang),
mainly use OCI registry as storage backend. 

Powered by some great Go technology:
- [helm/helm](https://github.com/helm/helm) - for working with charts
- [chartmuseum/auth](https://github.com/chartmuseum/chartmuseum) - for auth




### Docker Image
Available via [Docker Hub](https://hub.docker.com/r/chartmuseum/chartmuseum/).

Example usage (local storage):
```bash
docker run --rm -it \
  -p 8080:8080 \
  -e DEBUG=1 \
  -e STORAGE=registry \
  -e STORAGE_REGISTRY_REPO=localhost:5000 \
  -v $(pwd)/charts:/charts \
  hangyan/chart-registry:latest
```
