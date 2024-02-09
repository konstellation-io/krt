module github.com/konstellation-io/krt

go 1.21

toolchain go1.21.4

require (
	github.com/stretchr/testify v1.8.4
	gopkg.in/yaml.v3 v3.0.1
	k8s.io/apimachinery v0.29.1
)

require k8s.io/utils v0.0.0-20230726121419-3b25d923346b // indirect

require (
	github.com/alecthomas/units v0.0.0-20211218093645-b94a6e3cc137
	github.com/creasty/defaults v1.7.0
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
)
