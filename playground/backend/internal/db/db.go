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

package db

import (
	"beam.apache.org/playground/backend/internal/share"
	"context"
)

// Database represents data type that needed to use specific database
type Database string

// DATASTORE represents value indicates database as datastore
// LOCAL represents value indicates database for local usage or testing
const (
	DATASTORE Database = "datastore"
	LOCAL     Database = "local"
)

func (db Database) String() string {
	return string(db)
}

type SnippetDB interface {
	PutSnippet(ctx context.Context, id string, snip *share.SnippetDocument) error

	GetSnippet(ctx context.Context, id string) (*share.SnippetDocument, error)
}
