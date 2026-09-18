package postrenderer

import (
	"io"
	"log/slog"
	"os"
	"strings"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/serializer/yaml"
)

func RunPostrenderer() error {
	manifest, err := io.ReadAll(os.Stdin)
	if err != nil {
		return err
	}
	err = parseManifest(manifest)
	if err != nil {
		return err
	}
	_, err = io.WriteString(os.Stdout, string(manifest))
	if err != nil {
		return err
	}
	slog.Info("Run successfully")
	return nil
}

func parseManifest(manifest []byte) error {
	documents := strings.Split(string(manifest), "---")
	deserializer := yaml.NewDecodingSerializer(unstructured.UnstructuredJSONScheme)
	for _, doc := range documents {
		if _, _, err := deserializer.Decode([]byte(doc), nil, new(unstructured.Unstructured)); err != nil {
			return err
		}
	}
	return nil
}
