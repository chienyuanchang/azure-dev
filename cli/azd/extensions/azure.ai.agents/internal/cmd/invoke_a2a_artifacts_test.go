// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

package cmd

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestExtractA2ATextPrefersTaskArtifacts(t *testing.T) {
	t.Parallel()
	for _, version := range []string{"0.3", "1.0"} {
		t.Run(version, func(t *testing.T) {
			t.Parallel()
			// A native SDK task can finish with a generic lifecycle message while
			// its actual answer is carried in artifacts. Both dialects support this.
			task := map[string]any{
				"status": map[string]any{
					"message": map[string]any{"parts": []any{map[string]any{"text": "Done"}}},
				},
				"artifacts": []any{
					map[string]any{"parts": []any{map[string]any{"text": "native answer"}}},
				},
			}
			result := task
			if version == "1.0" {
				result = map[string]any{"task": task}
			}
			body, err := json.Marshal(map[string]any{"jsonrpc": "2.0", "result": result})
			require.NoError(t, err)
			require.Equal(t, "native answer", extractA2AText(body, version))
		})
	}
}
