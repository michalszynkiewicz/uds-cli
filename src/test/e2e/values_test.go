// Copyright 2024 Defense Unicorns
// SPDX-License-Identifier: AGPL-3.0-or-later OR LicenseRef-Defense-Unicorns-Commercial

// Package test provides e2e tests for UDS.
package test

import (
	"fmt"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

// TestBundleWithZarfValuesCreate tests that bundles with values configuration can be created.
// NOTE: Full e2e testing of Zarf values requires Zarf to export the value.Values type publicly.
// Currently, UDS can parse and load values, but cannot pass them to Zarf's Deploy function
// because value.Values is in an internal package.
// See: https://github.com/zarf-dev/zarf/src/internal/value/value.go
func TestBundleWithZarfValuesCreate(t *testing.T) {
	bundleDir := "src/test/bundles/17-values"

	t.Run("create bundle with values configuration", func(t *testing.T) {
		// This test verifies that bundles with values configuration can be created
		// The values configuration is parsed and validated during creation
		_, stderr := runCmd(t, fmt.Sprintf("create %s --insecure --confirm -a %s", bundleDir, e2e.Arch))

		// Should not have errors about values configuration
		require.NotContains(t, stderr, "failed to parse")
		require.NotContains(t, stderr, "values file not found")
	})

	t.Run("inspect bundle shows values configuration", func(t *testing.T) {
		bundleTarballPath := filepath.Join(bundleDir, fmt.Sprintf("uds-bundle-values-test-%s-0.0.1.tar.zst", e2e.Arch))

		stdout, _ := runCmd(t, fmt.Sprintf("inspect %s", bundleTarballPath))

		// The bundle should contain the values-pkg package
		require.Contains(t, stdout, "values-pkg")
	})
}
