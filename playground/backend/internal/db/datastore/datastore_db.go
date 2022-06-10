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

package datastore

import (
	"beam.apache.org/playground/backend/internal/logger"
	"beam.apache.org/playground/backend/internal/share"
	"cloud.google.com/go/datastore"
	"context"
	"time"
)

const (
	snippetCollection = "pg_snippets"
)

type Datastore struct {
	client *datastore.Client
}

func New(ctx context.Context, projectId string) (*Datastore, error) {
	client, err := datastore.NewClient(ctx, projectId)
	if err != nil {
		logger.Errorf("Datastore: connection to store: error during connection, err: %s\n", err.Error())
		return nil, err
	}
	return &Datastore{client: client}, nil
}

// PutSnippet puts the snippet to datastore database
func (f *Datastore) PutSnippet(ctx context.Context, id string, snip *share.SnippetDocument) error {
	key := datastore.NameKey(snippetCollection, id, nil)
	if _, err := f.client.Put(ctx, key, snip); err != nil {
		logger.Errorf("Datastore: PutSnippet(): error during snippet saving, err: %s\n", err.Error())
		return err
	}
	return nil
}

// GetSnippet returns the code snippet
func (f *Datastore) GetSnippet(ctx context.Context, id string) (*share.SnippetDocument, error) {
	key := datastore.NameKey(snippetCollection, id, nil)
	snip := new(share.SnippetDocument)
	if err := f.client.Get(ctx, key, snip); err != nil {
		logger.Errorf("Datastore: GetSnippet(): error during data getting, err: %s\n", err.Error())
		return nil, err
	}
	snip.LVisited = time.Now()
	snip.VisitCount += 1
	if err := f.PutSnippet(ctx, id, snip); err != nil {
		logger.Errorf("Datastore: GetSnippet(): error during data setting, err: %s\n", err.Error())
		return nil, err
	}
	return snip, nil
}
