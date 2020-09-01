module github.com/cnabio/cnab-go

go 1.13

require (
	get.porter.sh/porter v0.28.1
	github.com/Masterminds/semver v1.5.0
	github.com/bmizerany/assert v0.0.0-20160611221934-b7ed37b82869
	github.com/docker/cli v0.0.0-20191017083524-a8ff7f821017
	github.com/docker/distribution v2.7.1+incompatible
	github.com/docker/docker v1.4.2-0.20181229214054-f76d6a078d88
	github.com/docker/go v1.5.1-1
	github.com/globalsign/mgo v0.0.0-20181015135952-eeefdecb41b8
	github.com/gobuffalo/packr/v2 v2.8.0
	github.com/hashicorp/go-hclog v0.14.1
	github.com/hashicorp/go-multierror v1.1.0
	github.com/mitchellh/copystructure v1.0.0
	github.com/oklog/ulid v1.3.1
	github.com/pivotal/image-relocation v0.0.0-20191111101224-e94aff6df06c
	github.com/pkg/errors v0.9.1
	github.com/qri-io/jsonschema v0.1.1
	github.com/stretchr/testify v1.5.1
	github.com/xeipuuv/gojsonschema v1.2.0
	gopkg.in/yaml.v2 v2.2.4
	k8s.io/api v0.0.0-20191016110408-35e52d86657a
	k8s.io/apimachinery v0.0.0-20191004115801-a2eda9f80ab8
	k8s.io/client-go v0.0.0-20191016111102-bec269661e48
)

replace (
	github.com/docker/docker => github.com/moby/moby v0.7.3-0.20190826074503-38ab9da00309
	github.com/hashicorp/go-plugin => github.com/carolynvs/go-plugin v1.0.1-acceptstdin
)
