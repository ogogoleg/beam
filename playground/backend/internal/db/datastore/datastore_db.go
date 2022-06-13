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
	"time"
)

const (
	snippetKind = "pg_snippets"
	schemaKind  = "pg_schema_versions"
	sdkKind     = "pg_sdks"
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
func (d *Datastore) PutSnippet(ctx context.Context, id string, snip *entity.SnippetEntity) error {
	key := getNameKey(snippetKind, id)
	if _, err := d.client.Put(ctx, key, snip); err != nil {
		logger.Errorf("Datastore: PutSnippet(): error during entity saving, err: %s\n", err.Error())
		return err
	}
	return nil
}

// GetSnippet returns the snippet entity by identifier
func (d *Datastore) GetSnippet(ctx context.Context, id string) (*entity.SnippetEntity, error) {
	key := getNameKey(snippetKind, id)
	snip := new(entity.SnippetEntity)
	if err := d.client.Get(ctx, key, snip); err != nil {
		logger.Errorf("Datastore: GetSnippet(): error during data getting, err: %s\n", err.Error())
		return nil, err
	}
	snip.LVisited = time.Now()
	snip.VisitCount += 1
	if err := d.PutSnippet(ctx, id, snip); err != nil {
		logger.Errorf("Datastore: GetSnippet(): error during data setting, err: %s\n", err.Error())
		return nil, err
	}
	return snip, nil
}

// PutSchemaVersion puts the schema entity to datastore
func (d *Datastore) PutSchemaVersion(ctx context.Context, id string, schema *entity.SchemaEntity) error {
	key := getNameKey(schemaKind, id)
	if _, err := d.client.Put(ctx, key, schema); err != nil {
		logger.Errorf("Datastore: PutSchemaVersion(): error during entity saving, err: %s\n", err.Error())
		return err
	}
	return nil
}

// PutSDKs puts the SDK entity to datastore
func (d *Datastore) PutSDKs(ctx context.Context, sdks []*entity.SDKEntity) error {
	var keys []*datastore.Key
	for _, sdk := range sdks {
		keys = append(keys, getNameKey(sdkKind, sdk.Name))
	}
	if _, err := d.client.PutMulti(ctx, keys, sdks); err != nil {
		logger.Errorf("Datastore: PutSDK(): error during entity saving, err: %s\n", err.Error())
		return err
	}
	return nil
}

// getNameKey returns the datastore key
func getNameKey(kind, id string) *datastore.Key {
	key := datastore.NameKey(kind, id, nil)
	key.Namespace = "Playground" //TODO should it get from env?
	return key
}
