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

package entity

import "testing"

func TestSchema_ID(t *testing.T) {
	tests := []struct {
		name    string
		schema  *Schema
		want    string
		wantErr bool
	}{
		{
			name: "Schema ID() in the usual case",
			schema: &Schema{
				IDInfo: IDInfo{
					Salt:     "MOCK_SALT",
					IdLength: 11,
				},
				Schema: &SchemaEntity{
					Version: "MOCK_VERSION",
					Descr:   "MOCK_DESCRIPTION",
				},
			},
			want:    "C5At43gZ2Hy",
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			id, err := tt.schema.ID()
			if (err != nil) != tt.wantErr {
				t.Errorf("ID() error = %v, wantErr %v", err, tt.wantErr)
			}

			if err == nil {
				if len(id) != tt.schema.IdLength {
					t.Error("The ID length is not 11")
				}
				if tt.want != id {
					t.Error("ID is wrong")
				}
			}
		})
	}
}
