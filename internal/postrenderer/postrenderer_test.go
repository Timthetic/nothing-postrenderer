package postrenderer

import "testing"
import "github.com/stretchr/testify/require"

func TestRawManifestParses(t *testing.T) {
	manifest := []byte(`{"apiVersion": "v1", "kind": "ConfigMap", "metadata": {"name": "test"}, "data": {"key": "hello"}}`)
	err := parseManifest(manifest)
	require.NoError(t, err)
}
