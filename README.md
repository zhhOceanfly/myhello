# myhello

- go test ./...
- git tag v0.1.0
- git push origin v0.1.0
- go list

- GOPROXY=proxy.golang.org go list -m github.com/zhhOceanfly/myhello@v0.1.0

- GOPROXY=http://gitea.oceanfly.com/api/packages/{owner}/go go install {package_name}
