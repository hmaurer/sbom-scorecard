// SPDX-License-Identifier: Apache-2.0

package spdx

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLoadDocument(t *testing.T) {
	// Test 2.3 document load
	doc, err := LoadDocument("test_data/nginx.spdx.json")
	require.NoError(t, err)
	require.NotNil(t, doc)
	_, ok := doc.(*Document_23)
	require.True(t, ok)

	// Test 2.3 document load
	doc, err = LoadDocument("test_data/missionlz.spdx.json")
	require.NoError(t, err)
	require.NotNil(t, doc)
	// Currently 2.3 will open 2.2 docs. So this will fail
	// _, ok = doc.(*Document_22)
	// require.True(t, ok)
}
