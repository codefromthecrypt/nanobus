package apex_test

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/nanobus/nanobus/spec/apex"
)

func TestParse(t *testing.T) {
	expectedBytes, err := os.ReadFile("testdata/expected.json")
	if err != nil {
		t.FailNow()
	}

	name, loader := apex.Apex()
	assert.Equal(t, "apex", name)
	namespaces, err := loader(context.Background(), map[string]interface{}{
		"filename": "testdata/spec.apexlang",
	}, nil)
	require.NoError(t, err)
	require.Len(t, namespaces, 1)

	actualBytes, err := json.MarshalIndent(namespaces[0], "", "  ")
	require.NoError(t, err)
	fmt.Println(string(actualBytes))

	var expected, actual interface{}
	require.NoError(t, json.Unmarshal(expectedBytes, &expected))
	require.NoError(t, json.Unmarshal(actualBytes, &actual))

	assert.Equal(t, expected, actual)
}
