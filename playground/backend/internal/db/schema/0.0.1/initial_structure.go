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

package migration

import (
	pb "beam.apache.org/playground/backend/internal/api/v1"
	"beam.apache.org/playground/backend/internal/db/entity"
	"beam.apache.org/playground/backend/internal/db/schema"
)

const (
	version     = "0.0.1"
	description = "Data initialization: code snippet, schema versions, SDK"
)

type InitialStructure struct {
}

func (is *InitialStructure) InitiateData(args *schema.DBArgs) error {
	//init snippets
	dummyStr := "dummy"
	idInfo := entity.IDInfo{
		IdLength: args.AppEnv.IdLength(),
		Salt:     args.AppEnv.PlaygroundSalt(),
	}
	snip := &entity.Snippet{
		IDInfo: idInfo,
		Snippet: &entity.SnippetEntity{
			OwnerId:  dummyStr,
			PipeOpts: dummyStr,
		},
		Codes: []*entity.CodeEntity{
			{
				Name: dummyStr,
				Code: dummyStr,
			},
		},
	}
	snipId, err := snip.ID()
	if err != nil {
		return err
	}
	if err = args.Db.PutSnippet(args.Ctx, snipId, snip); err != nil {
		return err
	}

	//init schema versions
	schemaEntity := &entity.SchemaEntity{Descr: description}
	if err = args.Db.PutSchemaVersion(args.Ctx, version, schemaEntity); err != nil {
		return err
	}

	//init sdks
	var sdkEntities []*entity.SDKEntity
	for _, sdk := range pb.Sdk_name {
		sdkEntities = append(sdkEntities, &entity.SDKEntity{
			Name:           sdk,
			DefaultExample: "",
		})
	}
	if err = args.Db.PutSDKs(args.Ctx, sdkEntities); err != nil {
		return err
	}

	return nil
}

func (is *InitialStructure) GetVersion() string {
	return version
}

func (is *InitialStructure) GetDescription() string {
	return description
}
