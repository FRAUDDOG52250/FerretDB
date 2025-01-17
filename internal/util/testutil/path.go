// Copyright 2021 FerretDB Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package testutil

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/FerretDB/FerretDB/internal/types"
)

// GetByPath returns a value by path - a sequence of indexes and keys.
func GetByPath(tb testing.TB, str any, path ...string) any {
	tb.Helper()

	var res any
	var err error
	switch str := str.(type) {
	case types.Array:
		res, err = str.GetByPath(path...)
	case types.Document:
		res, err = str.GetByPath(path...)
	default:
		err = fmt.Errorf("can't access %T by path", str)
	}

	require.NoError(tb, err)
	return res
}

// SetByPath sets the value by path - a sequence of indexes and keys.
//
// The path must exist.
func SetByPath(tb testing.TB, str any, value any, path ...string) {
	tb.Helper()

	l := len(path)
	require.NotZero(tb, l, "path is empty")

	for i, p := range path {
		last := i == l-1
		switch s := str.(type) {
		case types.Array:
			index, err := strconv.Atoi(p)
			require.NoError(tb, err)

			str, err = s.Get(index)
			require.NoError(tb, err)

			if last {
				err = s.Set(index, value)
				require.NoError(tb, err)
			}

		case types.Document:
			var err error
			str, err = s.Get(p)
			require.NoError(tb, err)

			if last {
				err = s.Set(p, value)
				require.NoError(tb, err)
			}

		default:
			tb.Fatalf("can't access %T by path %q", str, p)
		}
	}
}

// CompareAndSetByPath asserts that two values with the same path in two objects (documents or arrays)
// are within a given delta, then updates the expected object with the actual value.
func CompareAndSetByPath(tb testing.TB, expected, actual any, delta float64, path ...string) {
	tb.Helper()

	expectedV := GetByPath(tb, expected, path...)
	actualV := GetByPath(tb, actual, path...)
	assert.IsType(tb, expectedV, actualV)
	assert.InDelta(tb, expectedV, actualV, delta)

	SetByPath(tb, expected, actualV, path...)
}
