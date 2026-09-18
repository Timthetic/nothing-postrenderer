# Nothing Postrenderer

A postrenderer that just tries to parse all the objects in the manifest along with some example helm charts. This reproduces
an issue where a manifest that `yaml.NewDecodingSerializer(unstructured.UnstructuredJSONScheme)` would have been able
to parse does not parse when passed through a Helm postrenderer that does not mutate the input.

This behavior is a little surprising to me for a few reasons:
1. This won't cause an issue with Helm 3, it seems to have to do with Helm 4's use of `kyaml` to inject annotations.
2. `yaml.NewDecodingSerializer(unstructured.UnstructuredJSONScheme)` appears to assume that anything with a leading `{` is pure JSON (see [here](https://github.com/kubernetes/kubernetes/blob/672e05c830ff5d4d7dfccbc87e1fde5b4be47d10/staging/src/k8s.io/apimachinery/pkg/util/yaml/decoder.go#L475))

## Usage

* `make` to generate the postrenderer binary in the build folder
* `make test` shows the original manifest parsing without passing through a postrenderer
* `HELM_PLUGINS="./example-plugins" helm template successful-chart --post-renderer postrenderer` shows the postrenderer running for a helm chart that is formated as YAML that is not JSON
* `HELM_PLUGINS="./example-plugins" helm template erroring-chart --post-renderer postrenderer` shows the postrenderer failing on a helm chart that is formated as JSON (all JSON is valid YAML)


