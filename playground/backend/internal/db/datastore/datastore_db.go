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
	"beam.apache.org/playground/backend/internal/db/entity"
	"beam.apache.org/playground/backend/internal/logger"
	"cloud.google.com/go/datastore"
	"context"
	"google.golang.org/api/iterator"
	"time"
)

const (
	snippetKind = "pg_snippets"
	schemaKind  = "pg_schema_versions"
	sdkKind     = "pg_sdks"
	codeKind    = "pg_codes"

	namespace = "Playground" //TODO should it get from env?
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

// PutSnippet puts the snippet entity to datastore
func (d *Datastore) PutSnippet(ctx context.Context, id string, snip *entity.Snippet) error {
	if snip == nil {
		logger.Errorf("Datastore: PutSnippet(): snippet is nil")
		return nil
	}
	key := getNameKey(snippetKind, id, nil)
	if _, err := d.client.Put(ctx, key, snip.Snippet); err != nil {
		logger.Errorf("Datastore: PutSnippet(): error during the snippet entity saving, err: %s\n", err.Error())
		return err
	}

	var keys []*datastore.Key
	for _, code := range snip.Codes {
		codeId, err := code.ID(snip)
		if err != nil {
			logger.Errorf("Datastore: PutSnippet(): error during the code ID generation, err: %s\n", err.Error())
			return err
		}
		keys = append(keys, getNameKey(codeKind, codeId, key))
	}

	if _, err := d.client.PutMulti(ctx, keys, snip.Codes); err != nil {
		logger.Errorf("Datastore: PutSnippet(): error during the code entity saving, err: %s\n", err.Error())
		return err
	}

	return nil
}

// GetSnippet returns the snippet entity by identifier
func (d *Datastore) GetSnippet(ctx context.Context, id string) (*entity.SnippetEntity, error) {
	key := getNameKey(snippetKind, id, nil)
	snip := new(entity.SnippetEntity)
	if err := d.client.Get(ctx, key, snip); err != nil {
		logger.Errorf("Datastore: GetSnippet(): error during snippet getting, err: %s\n", err.Error())
		return nil, err
	}
	snip.LVisited = time.Now()
	snip.VisitCount += 1
	if _, err := d.client.Put(ctx, key, snip); err != nil {
		logger.Errorf("Datastore: GetSnippet(): error during snippet setting, err: %s\n", err.Error())
		return nil, err
	}
	return snip, nil
}

// PutSchemaVersion puts the schema entity to datastore
func (d *Datastore) PutSchemaVersion(ctx context.Context, id string, schema *entity.SchemaEntity) error {
	if schema == nil {
		logger.Errorf("Datastore: PutSchemaVersion(): schema version is nil")
		return nil
	}
	key := getNameKey(schemaKind, id, nil)
	if _, err := d.client.Put(ctx, key, schema); err != nil {
		logger.Errorf("Datastore: PutSchemaVersion(): error during entity saving, err: %s\n", err.Error())
		return err
	}
	return nil
}

// PutSDKs puts the SDK entity to datastore
func (d *Datastore) PutSDKs(ctx context.Context, sdks []*entity.SDKEntity) error {
	if sdks == nil || len(sdks) == 0 {
		logger.Errorf("Datastore: PutSDKs(): sdks are empty")
		return nil
	}
	var keys []*datastore.Key
	for _, sdk := range sdks {
		keys = append(keys, getNameKey(sdkKind, sdk.Name, nil))
	}
	if _, err := d.client.PutMulti(ctx, keys, sdks); err != nil {
		logger.Errorf("Datastore: PutSDK(): error during entity saving, err: %s\n", err.Error())
		return err
	}
	return nil
}

//GetCodes returns the code entities by parent identifier
func (d *Datastore) GetCodes(ctx context.Context, parentId string) ([]*entity.CodeEntity, error) {
	snipId := getNameKey(snippetKind, parentId, nil)
	query := datastore.NewQuery(codeKind).Ancestor(snipId).Namespace(namespace)
	it := d.client.Run(ctx, query)
	var codes []*entity.CodeEntity
	for {
		var code entity.CodeEntity
		_, err := it.Next(&code)
		if err == iterator.Done {
			break
		}
		if err != nil {
			logger.Errorf("Datastore: GetCodes(): error during code getting, err: %s\n", err.Error())
		}
		codes = append(codes, &code)
	}
	return codes, nil
}

// getNameKey returns the datastore key
func getNameKey(kind, id string, parentId *datastore.Key) *datastore.Key {
	key := datastore.NameKey(kind, id, nil)
	if parentId != nil {
		key.Parent = parentId
	}
	key.Namespace = namespace
	return key
}
