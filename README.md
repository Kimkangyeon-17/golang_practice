# Go 필수 명령어 가이드

## 1. 프로젝트 초기화

### go mod init
```bash
go mod init [모듈명]
```
- 새 Go 모듈을 생성하고 go.mod 파일을 만듭니다
- 예시: `go mod init github.com/username/project`
- 프로젝트 시작 시 가장 먼저 실행하는 명령어

## 2. 코드 실행

### go run
```bash
go run [파일명.go]
```
- Go 파일을 컴파일하고 즉시 실행합니다
- 개발 중 빠른 테스트에 유용
- 예시: `go run main.go`

### go build
```bash
go build
go build [파일명.go]
go build -o [실행파일명]
```
- 실행 가능한 바이너리 파일을 생성합니다
- `-o` 옵션으로 출력 파일명 지정 가능
- 예시: `go build -o myapp main.go`

## 3. 의존성 관리

### go get
```bash
go get [패키지 경로]
go get -u [패키지 경로]
```
- 외부 패키지를 다운로드하고 설치합니다
- `-u` 옵션: 패키지를 최신 버전으로 업데이트
- 예시: `go get github.com/gin-gonic/gin`

### go mod tidy
```bash
go mod tidy
```
- 사용하지 않는 의존성을 제거하고 필요한 의존성을 추가합니다
- go.mod와 go.sum 파일을 정리합니다
- 의존성 변경 후 항상 실행 권장

### go mod download
```bash
go mod download
```
- go.mod에 명시된 모든 의존성을 다운로드합니다

## 4. 테스트

### go test
```bash
go test
go test ./...
go test -v
go test -cover
```
- `_test.go` 파일의 테스트를 실행합니다
- `./...`: 현재 디렉토리와 하위 디렉토리의 모든 테스트 실행
- `-v`: 자세한 출력
- `-cover`: 코드 커버리지 표시

## 5. 코드 포맷팅 및 검사

### go fmt
```bash
go fmt
go fmt ./...
```
- Go 코드를 표준 스타일로 자동 포맷팅합니다
- 코드 커밋 전에 실행 권장

### go vet
```bash
go vet
go vet ./...
```
- 코드의 잠재적인 오류를 검사합니다
- 컴파일은 되지만 문제가 될 수 있는 코드를 찾아냅니다

## 6. 기타 유용한 명령어

### go install
```bash
go install [패키지 경로]
```
- 패키지를 컴파일하고 $GOPATH/bin에 설치합니다
- CLI 도구 설치에 주로 사용

### go clean
```bash
go clean
```
- 빌드로 생성된 파일들을 삭제합니다

### go version
```bash
go version
```
- 설치된 Go 버전을 확인합니다

### go env
```bash
go env
go env GOPATH
```
- Go 환경 변수를 확인합니다

## 일반적인 개발 워크플로우

```bash
# 1. 프로젝트 초기화
go mod init myproject

# 2. 코드 작성 후 실행
go run main.go

# 3. 외부 패키지 추가
go get github.com/some/package

# 4. 의존성 정리
go mod tidy

# 5. 코드 포맷팅
go fmt ./...

# 6. 테스트 실행
go test ./...

# 7. 빌드
go build -o myapp

# 8. 실행 파일 실행
./myapp
```

## 실전 예시

### 예시 1: 간단한 프로젝트 만들고 실행하기
```bash
# 프로젝트 초기화
go mod init goproject/ex1.1

# main.go 파일 작성 후 빌드
go build

# Windows에서 실행
.\ex1.1.exe

# Linux/Mac에서 실행
./ex1.1
```

### 예시 2: 빠른 개발 및 테스트
```bash
# 초기화
go mod init goproject/hello

# 바로 실행 (빌드 없이)
go run main.go

# 수정 후 다시 실행
go run main.go
```

### 예시 3: 외부 패키지 사용
```bash
# 프로젝트 초기화
go mod init goproject/webserver

# 웹 프레임워크 설치
go get github.com/gin-gonic/gin

# 의존성 정리
go mod tidy

# 실행
go run main.go
```

## 추가 팁

- **개발 시**: `go run`으로 빠르게 테스트
- **배포 전**: `go build`로 실행 파일 생성
- **의존성 변경 후**: 반드시 `go mod tidy` 실행
- **커밋 전**: `go fmt`, `go vet`, `go test` 실행 습관화