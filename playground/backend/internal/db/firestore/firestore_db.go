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
	snippetCollection = "snippets"
	codeCollection    = "codes"
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
	snipDoc := f.client.Collection(snippetCollection).Doc(id)
	_, err := snipDoc.Set(ctx, snip)
	if err != nil {
		logger.Errorf("Firestore: PutSnippet(): error during data setting, err: %s\n", err.Error())
		return err
	}

	for _, code := range snip.Codes {
		code.SnipId = id
		codeId, err := code.ID(snip.Salt)
		if err != nil {
			logger.Errorf("Firestore: PutSnippet(): error during code id generation, err: %s\n", err.Error())
			return err
		}
		_, err = f.client.Collection(codeCollection).Doc(codeId).Set(ctx, code)
		if err != nil {
			logger.Errorf("Firestore: PutSnippet(): error during data setting, err: %s\n", err.Error())
			return err
		}
	}
	return nil
}

// GetSnippet returns the code snippet
func (f *Firestore) GetSnippet(ctx context.Context, id string) (*share.Snippet, error) {
	dsnap, err := f.client.Collection(snippetCollection).Doc(id).Get(ctx)
	if err != nil {
		logger.Errorf("Firestore: GetSnippet(): error during snippets getting, err: %s\n", err.Error())
		return nil, err
	}
	var snip share.Snippet
	jsonStr, err := json.Marshal(dsnap.Data())
	if err != nil {
		logger.Errorf("Firestore: GetSnippet(): error during data transformation, err: %s\n", err.Error())
		return nil, err
	}
	if err := json.Unmarshal(jsonStr, &snip); err != nil {
		logger.Errorf("Firestore: GetSnippet(): error during data transformation, err: %s\n", err.Error())
		return nil, err
	}

	codeIter := f.client.Collection(codeCollection).Where("snipId", "==", id).Documents(ctx)
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
		var code share.Code
		if err := json.Unmarshal(jsonStr, &code); err != nil {
			logger.Errorf("Firestore: GetSnippet(): error during data transformation, err: %s\n", err.Error())
			return nil, err
		}
		snip.Codes = append(snip.Codes, code)
	}

	snip.LastVisited = time.Now()
	snip.VisitCount += 1
	_, err = f.client.Collection(snippetCollection).Doc(id).Set(ctx, snip)
	if err != nil {
		logger.Errorf("Firestore: GetSnippet(): error during data setting, err: %s\n", err.Error())
		return nil, err
	}

	return &snip, nil
}
