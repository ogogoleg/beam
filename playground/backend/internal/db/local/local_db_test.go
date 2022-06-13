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

package local_db

import (
	db "beam.apache.org/playground/backend/internal/db/entity"
	"context"
	"testing"
	"time"
)

var ctx context.Context
var localDb *LocalDB

func init() {
	ctx = context.Background()
	var err error
	localDb, err = New()
	if err != nil {
		panic(err)
	}
}

func TestLocalDB_PutSnippet(t *testing.T) {
	type args struct {
		ctx  context.Context
		id   string
		snip *db.SnippetEntity
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "PutSnippet() in the usual case",
			args: args{
				ctx: ctx,
				id:  "MOCK_ID",
				snip: &db.SnippetEntity{
					Sdk:      "SDK_GO",
					PipeOpts: "MOCK_OPTIONS",
					Origin:   db.PLAYGROUND,
					OwnerId:  "",
					Codes: []*db.CodeEntity{{
						Code:   "MOCK_CODE",
						IsMain: false,
					}},
				},
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := localDb.PutSnippet(tt.args.ctx, tt.args.id, tt.args.snip)
			if (err != nil) != tt.wantErr {
				t.Errorf("PutSnippet() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}

	delete(localDb.items, "MOCK_ID")
}

func TestLocalDB_GetSnippet(t *testing.T) {
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
			name:    "GetSnippet() when ID is not in the database",
			prepare: func() {},
			args: args{
				ctx: ctx,
				id:  "MOCK_ID",
			},
			wantErr: true,
		},
		{
			name: "GetSnippet() in the usual case",
			prepare: func() {
				_ = localDb.PutSnippet(ctx, "MOCK_ID", &db.SnippetEntity{
					Sdk:      "SDK_GO",
					PipeOpts: "MOCK_OPTIONS",
					Created:  nowDate,
					Origin:   db.PLAYGROUND,
					OwnerId:  "",
					Codes: []*db.CodeEntity{{
						Code:   "MOCK_CODE",
						IsMain: false,
					}},
				})
			},
			args: args{
				ctx: ctx,
				id:  "MOCK_ID",
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.prepare()
			snip, err := localDb.GetSnippet(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetSnippet() error = %v, wantErr %v", err, tt.wantErr)
			}

			if err == nil {
				if snip.Sdk != "SDK_GO" ||
					snip.Codes[0].Code != "MOCK_CODE" ||
					snip.PipeOpts != "MOCK_OPTIONS" ||
					snip.Origin != db.PLAYGROUND ||
					snip.OwnerId != "" {
					t.Error("GetSnippet() unexpected result")
				}
			}
		})
	}

	delete(localDb.items, "MOCK_ID")
}

func TestNew(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		{
			name:    "Initialize local database",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if _, err := New(); (err != nil) != tt.wantErr {
				t.Errorf("New() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
