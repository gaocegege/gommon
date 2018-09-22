package testutil

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"testing"

	"gopkg.in/yaml.v2"
)

func ReadFixture(t *testing.T, path string) []byte {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		t.Fatalf("can't read fixture %s: %v", path, err)
	}
	return b
}

func ReadJsonTo(t *testing.T, path string, v interface{}) {
	b := ReadFixture(t, path)
	if err := json.Unmarshal(b, v); err != nil {
		t.Fatalf("can't unmarshal fixture %s %v", path, err)
	}
}

func ReadYAMLTo(t *testing.T, path string, v interface{}) {
	b := ReadFixture(t, path)
	if err := yaml.Unmarshal(b, v); err != nil {
		t.Fatalf("can't unmarhsal YAML fixture %s %v", path, err)
	}
}

// ReadYAMLToStrict uses strict mode when decoding, if unknown fields shows up in YAML but not in struct it will error
func ReadYAMLToStrict(t *testing.T, path string, v interface{}) {
	b := ReadFixture(t, path)
	if err := yaml.UnmarshalStrict(b, v); err != nil {
		t.Fatalf("can't unmarhsal YAML fixture %s in strict mode %v", path, err)
	}
}

func WriteFixture(t *testing.T, path string, data []byte) {
	err := ioutil.WriteFile(path, data, 0664)
	if err != nil {
		t.Fatalf("can't write fixture %s: %v", path, err)
	}
}

func SaveAsJson(t *testing.T, v interface{}, file string) {
	b, err := json.Marshal(v)
	if err != nil {
		t.Fatalf("failed to encode as json: %s %v", err, v)
		return
	}
	if err := ioutil.WriteFile(file, b, 0664); err != nil {
		t.Fatalf("failed to save file %s: %v", file, err)
		return
	}
}

func SaveAsJsonf(t *testing.T, v interface{}, format string, args ...interface{}) {
	file := fmt.Sprintf(format, args...)
	b, err := json.Marshal(v)
	if err != nil {
		t.Fatalf("failed to encode as json: %s %v", err, v)
		return
	}
	if err := ioutil.WriteFile(file, b, 0664); err != nil {
		t.Fatalf("failed to save file %s: %v", file, err)
		return
	}
}

func SaveAsPrettyJson(t *testing.T, v interface{}, file string) {
	b, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		t.Fatalf("failed to encode as json: %s %v", err, v)
		return
	}
	if err := ioutil.WriteFile(file, b, 0664); err != nil {
		t.Fatalf("failed to save file %s: %v", file, err)
		return
	}
	t.Logf("saved json to %s", file)
}

func SaveAsPrettyJsonf(t *testing.T, v interface{}, format string, args ...interface{}) {
	file := fmt.Sprintf(format, args...)
	b, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		t.Fatalf("failed to encode as json %v", v)
		return
	}
	if err := ioutil.WriteFile(file, b, 0664); err != nil {
		t.Fatalf("failed to save file %s: %v", file, err)
		return
	}
	t.Logf("saved json to %s", file)
}

func SaveAsYAML(t *testing.T, v interface{}, file string) {
	b, err := yaml.Marshal(v)
	if err != nil {
		t.Fatalf("failed to encode as YAML: %s %v", err, v)
		return
	}
	if err := ioutil.WriteFile(file, b, 0664); err != nil {
		t.Fatalf("failed to save file %s: %v", file, err)
		return
	}
	t.Logf("saved YAML to %s", file)
}

func SaveAsYAMLf(t *testing.T, v interface{}, format string, args ...interface{}) {
	file := fmt.Sprintf(format, args...)
	b, err := yaml.Marshal(v)
	if err != nil {
		t.Fatalf("failed to encode as YAML: %s %v", err, v)
		return
	}
	if err := ioutil.WriteFile(file, b, 0664); err != nil {
		t.Fatalf("failed to save file %s: %v", file, err)
		return
	}
	t.Logf("saved YAML to %s", file)
}
