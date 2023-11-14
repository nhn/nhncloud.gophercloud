# Extensions for Gophercloud: an NHN Cloud SDK for Go

NHN Cloud 서비스 사용을 위한 [Gophercloud](https://github.com/gophercloud/gophercloud) 의 Go SDK 확장 프로그램 입니다.

다음과 같은 NHN Cloud 서비스의 SDK 를 포함하고 있습니다.

* VPC
* VPC subnet

## How to install

코드 내에 다음과 같이 `nhncloud.gophercloud` 패키지를 참조합니다.

```go
import "github.com/nhn/nhncloud.gophercloud"
```

그런 다음 `go.mod`를 업데이트 합니다.

```shell
go mod tidy
```

---

혹은 `go get` 명령어를 사용하여 패키지를 가져올 수 있습니다.

```shell
go get github.com/nhn/nhncloud.gophercloud
```