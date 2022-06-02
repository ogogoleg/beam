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
	"beam.apache.org/playground/backend/internal/logger"
	"beam.apache.org/playground/backend/internal/share"
	"cloud.google.com/go/firestore"
	"context"
	"encoding/json"
	"google.golang.org/api/iterator"
	"time"
)

const (
	snippetCollection = "pg_snippets"
	codeCollection    = "pg_codes"
)

type Firestore struct {
	client *firestore.Client
}

func New(ctx context.Context, projectId string) (*Firestore, error) {
	client, err := firestore.NewClient(ctx, projectId)
	if err != nil {
		logger.Errorf("Firestore: connect to store: error during connection, err: %s\n", err.Error())
		return nil, err
	}
	return &Firestore{client: client}, nil
}

// PutSnippet puts the snippet to firestore database
func (f *Firestore) PutSnippet(ctx context.Context, id string, snip *share.Snippet) error {
	batch := f.client.Batch()
	snipDoc := f.client.Collection(snippetCollection).Doc(id)
	batch.Set(snipDoc, snip.Snippet)

	for _, code := range snip.Codes {
		codeId, err := code.ID(snip)
		if err != nil {
			logger.Errorf("Firestore: PutSnippet(): error during code id generation, err: %s\n", err.Error())
			return err
		}
		codeDoc := snipDoc.Collection(codeCollection).Doc(codeId)
		batch.Set(codeDoc, code)
	}

	_, err := batch.Commit(ctx)
	if err != nil {
		logger.Errorf("Firestore: PutSnippet(): error during the snippet saving, err: %s\n", err.Error())
		return err
	}

	return nil
}

// GetSnippet returns the code snippet
func (f *Firestore) GetSnippet(ctx context.Context, id string) (*share.Snippet, error) {
	snipFSDoc := f.client.Collection(snippetCollection).Doc(id)
	dsnap, err := snipFSDoc.Get(ctx)
	if err != nil {
		logger.Errorf("Firestore: GetSnippet(): error during snippets getting, err: %s\n", err.Error())
		return nil, err
	}
	var snipDoc share.SnippetDocument
	var snip share.Snippet
	jsonStr, err := json.Marshal(dsnap.Data())
	if err != nil {
		logger.Errorf("Firestore: GetSnippet(): error during data transformation, err: %s\n", err.Error())
		return nil, err
	}
	if err := json.Unmarshal(jsonStr, &snipDoc); err != nil {
		logger.Errorf("Firestore: GetSnippet(): error during data transformation, err: %s\n", err.Error())
		return nil, err
	}
	snip.Snippet = &snipDoc

	codeIter := snipFSDoc.Collection(codeCollection).Documents(ctx)
	for {
		doc, err := codeIter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			logger.Errorf("Firestore: GetSnippet(): error during codes getting, err: %s\n", err.Error())
			return nil, err
		}
		jsonStr, err := json.Marshal(doc.Data())
		if err != nil {
			logger.Errorf("Firestore: GetSnippet(): error during data transformation, err: %s\n", err.Error())
			return nil, err
		}
		var code share.CodeDocument
		if err := json.Unmarshal(jsonStr, &code); err != nil {
			logger.Errorf("Firestore: GetSnippet(): error during data transformation, err: %s\n", err.Error())
			return nil, err
		}
		snip.Codes = append(snip.Codes, &code)
	}

	snip.Snippet.LVisited = time.Now()
	snip.Snippet.VisitCount += 1
	_, err = f.client.Collection(snippetCollection).Doc(id).Set(ctx, snip.Snippet)
	if err != nil {
		logger.Errorf("Firestore: GetSnippet(): error during data setting, err: %s\n", err.Error())
		return nil, err
	}

	return &snip, nil
}
