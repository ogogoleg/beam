// Licensed to the Apache Software Foundation (ASF) under one or more
// contributor license agreements.  See the NOTICE file distributed with
// this work for additional information regarding copyright ownership.
// The ASF licenses this file to You under the Apache License, Version 2.0
// (the "License"); you may not use this file except in compliance with
// the License.  You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package share

import "testing"

func TestSnippet_ID(t *testing.T) {
	tests := []struct {
		name    string
		snip    *Snippet
		want    string
		wantErr bool
	}{
		{
			name: "Snippet ID() in the usual case",
			snip: &Snippet{
				Codes: []*CodeDocument{{
					Name:   "MOCK_NAME",
					Code:   "MOCK_CODE",
					IsMain: false,
				}},
				Snippet: &SnippetDocument{
					Sdk:      1,
					PipeOpts: "MOCK_OPTIONS",
				},
				IdLength: 11,
				Salt:     "MOCK_SALT",
			},
			want:    "",
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			id, err := tt.snip.ID()
			if (err != nil) != tt.wantErr {
				t.Errorf("ID() error = %v, wantErr %v", err, tt.wantErr)
			}

			if err == nil {
				if len(id) != tt.snip.IdLength {
					t.Error("The ID length is not 11")
				}
				if "OnY1ocoN0EO" != id {
					t.Error("ID is wrong")
				}
			}
		})
	}
}

func TestCode_ID(t *testing.T) {
	code := &CodeDocument{
		Name:   "MOCK_NAME",
		Code:   "MOCK_CODE",
		IsMain: false,
	}

	tests := []struct {
		name    string
		snip    *Snippet
		code    *CodeDocument
		want    string
		wantErr bool
	}{
		{
			name: "CodeDocument ID() in the usual case",
			snip: &Snippet{
				Salt:  "MOCK_SALT",
				Codes: []*CodeDocument{code},
				Snippet: &SnippetDocument{
					Sdk:      1,
					PipeOpts: "MOCK_OPTIONS",
				},
				IdLength: 11,
			},
			code:    code,
			want:    "",
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			id, err := tt.code.ID(tt.snip)
			if (err != nil) != tt.wantErr {
				t.Errorf("ID() error = %v, wantErr %v", err, tt.wantErr)
			}

			if err == nil {
				if len(id) != tt.snip.IdLength {
					t.Error("The ID length is not 11")
				}
				if "XSnhl0HoUOd" != id {
					t.Error("ID is wrong")
				}
			}
		})
	}
}
