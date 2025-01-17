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

package types

import "fmt"

// Array represents BSON array.
type Array []any

// Get returns a value at the given index.
func (a Array) Get(index int) (any, error) {
	if l := len(a); index < 0 || index >= l {
		return nil, fmt.Errorf("types.Array.Get: index %d is out of bounds [0-%d)", index, l)
	}

	return a[index], nil
}

// GetByPath returns a value by path - a sequence of indexes and keys.
func (a Array) GetByPath(path ...string) (any, error) {
	return getByPath(a, path...)
}

// Set sets the value at the given index.
func (a Array) Set(index int, value any) error {
	if l := len(a); index < 0 || index >= l {
		return fmt.Errorf("types.Array.Set: index %d is out of bounds [0-%d)", index, l)
	}

	if err := validateValue(value); err != nil {
		return fmt.Errorf("types.Array.Set: %w", err)
	}

	a[index] = value
	return nil
}
