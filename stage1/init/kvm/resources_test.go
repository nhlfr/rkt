// Copyright 2015 The rkt Authors
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

package kvm

import (
	"testing"

	"github.com/appc/spec/schema/types"
)

func TestFindResources(t *testing.T) {
	tests := []struct {
		in types.Isolators

		wmem int64
		wcpu int64
	}{
		{
			types.Isolators{},

			defaultMem,
			0,
		},
		{
			types.Isolators([]types.Isolator{
				newIsolator(`
				{
					"name":     "resource/cpu",
					"value": {
						"limit": 100,
						"request": 100
						}
				}`),
			}),

			defaultMem,
			100,
		},
	}

	for i, tt := range tests {
		gmem, gcpu := findResources(tt.in)
		if gmem != tt.wmem {
			t.Errorf("#%d: got mem=%d, want %d", i, gmem, tt.wmem)
		}
		if gcpu != tt.wcpu {
			t.Errorf("#%d: got cpu=%d, want %d", i, gcpu, tt.wcpu)
		}
	}
}

func newIsolator(body string) (i types.Isolator) {
	err := i.UnmarshalJSON([]byte(body))
	if err != nil {
		panic(err)
	}
	return
}
