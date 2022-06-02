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

package firestore

import (
	"beam.apache.org/playground/backend/internal/share"
	"context"
	"os"
	"testing"
	"time"
)

const (
	firestoreEmulatorHostKey   = "FIRESTORE_EMULATOR_HOST"
	firestoreEmulatorHostValue = "localhost:8082"
	firestoreEmulatorProjectId = "dummy-emulator-firestore-project"
)

var firestoreDb *Firestore
var ctx context.Context

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	teardown()
	os.Exit(code)
}

func setup() {
	firestoreEmulatorHost := os.Getenv(firestoreEmulatorHostKey)
	if firestoreEmulatorHost == "" {
		if err := os.Setenv(firestoreEmulatorHostKey, firestoreEmulatorHostValue); err != nil {
			panic(err)
		}
	}
	ctx = context.Background()
	var err error
	firestoreDb, err = New(ctx, firestoreEmulatorProjectId)
	if err != nil {
		panic(err)
	}
}

func teardown() {
	if err := firestoreDb.client.Close(); err != nil {
		panic(err)
	}
}

func TestFirestore_PutSnippet(t *testing.T) {
	type args struct {
		ctx  context.Context
		id   string
		snip *share.Snippet
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "PutSnippet() in the usual case",
			args: args{ctx: ctx, id: "MOCK_ID", snip: &share.Snippet{
				IdLength: 11,
				Codes: []*share.CodeDocument{{
					Code:   "MOCK_CODE",
					IsMain: false,
				}},
				Snippet: &share.SnippetDocument{
					Sdk:      1,
					PipeOpts: "MOCK_OPTIONS",
				},
			}},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := firestoreDb.PutSnippet(tt.args.ctx, tt.args.id, tt.args.snip)
			if err != nil {
				t.Error("PutSnippet() method failed")
			}
		})
	}

	cleanData(t)
}

func TestFirestore_GetSnippet(t *testing.T) {
	nowDate := time.Now()
	type args struct {
		ctx context.Context
		id  string
	}
	tests := []struct {
		name    string
		prepare func()
		args    args
		wantErr bool
	}{
		{
			name:    "GetSnippet() with id that is no in the database",
			prepare: func() {},
			args:    args{ctx: ctx, id: "MOCK_ID"},
			wantErr: true,
		},
		{
			name: "GetSnippet() in the usual case",
			prepare: func() {
				_ = firestoreDb.PutSnippet(ctx, "MOCK_ID", &share.Snippet{
					IdLength: 11,
					Codes: []*share.CodeDocument{{
						Code:   "MOCK_CODE",
						IsMain: false,
					}},
					Snippet: &share.SnippetDocument{
						Sdk:      1,
						PipeOpts: "MOCK_OPTIONS",
						Created:  nowDate,
						Origin:   share.PLAYGROUND,
						OwnerId:  "",
					},
				})
			},
			args:    args{ctx: ctx, id: "MOCK_ID"},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.prepare()
			snip, err := firestoreDb.GetSnippet(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetSnippet() error = %v, wantErr %v", err, tt.wantErr)
			}

			if err == nil {
				if snip.Snippet.Sdk != 1 ||
					snip.Codes[0].Code != "MOCK_CODE" ||
					snip.Snippet.PipeOpts != "MOCK_OPTIONS" ||
					snip.Snippet.Created.Local() != nowDate.Local() ||
					snip.Snippet.Origin != share.PLAYGROUND ||
					snip.Snippet.OwnerId != "" {
					t.Error("GetSnippet() unexpected result")
				}
			}
		})
	}

	cleanData(t)
}

func TestNew(t *testing.T) {
	type args struct {
		ctx       context.Context
		projectId string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "Initialize firestore database",
			args:    args{ctx: ctx, projectId: firestoreEmulatorProjectId},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := New(ctx, firestoreEmulatorProjectId)
			if (err != nil) != tt.wantErr {
				t.Errorf("New() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func cleanData(t *testing.T) {
	_, err := firestoreDb.client.Collection(snippetCollection).Doc("MOCK_ID").Delete(ctx)
	if err != nil {
		t.Error("Error during data cleaning after the test")
	}
}
