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

import (
	"beam.apache.org/playground/backend/internal/utils"
	"fmt"
	"sort"
	"time"
)

type Origin int32

const (
	TOUR_OF_BEAM Origin = 0 //will be used in Tour of Beam project
	PLAYGROUND   Origin = 1
)

func (s Origin) Value() int32 {
	return int32(s)
}

type CodeEntity struct {
	Name     string `datastore:"name"`
	Code     string `datastore:"code"`
	CntxLine int32  `datastore:"cntxLine"`
	IsMain   bool   `datastore:"isMain"`
}

type SnippetEntity struct {
	OwnerId    string        `datastore:"ownerId"`
	Sdk        string        `datastore:"sdk"`
	PipeOpts   string        `datastore:"pipeOpts"`
	Created    time.Time     `datastore:"created"`
	LVisited   time.Time     `datastore:"lVisited"`
	Origin     Origin        `datastore:"origin"`
	VisitCount int           `datastore:"visitCount"`
	Codes      []*CodeEntity `datastore:"codes"`
	SchVer     string        `datastore:"schVer"`
}

type Snippet struct {
	IDInfo
	Snippet *SnippetEntity
}

// ID generates id according to content of the entity
func (s *Snippet) ID() (string, error) {
	var codes []string
	for _, v := range s.Snippet.Codes {
		codes = append(codes, utils.RemoveSpaces(v.Code)+utils.RemoveSpaces(v.Name))
	}
	sort.Strings(codes)
	var content string
	for i, v := range codes {
		content += v
		if i == len(codes)-1 {
			content += fmt.Sprintf("%v%s", s.Snippet.Sdk, utils.RemoveSpaces(s.Snippet.PipeOpts))
		}
	}
	id, err := utils.ID(s.Salt, content, s.IdLength)
	if err != nil {
		return "", err
	}
	return id, nil
}

// ID generates id according to content of a code and its name
func (c *CodeEntity) ID(snip *Snippet) (string, error) {
	content := fmt.Sprintf("%s%s", utils.RemoveSpaces(c.Code), utils.RemoveSpaces(c.Name))
	id, err := utils.ID(snip.Salt, content, snip.IdLength)
	if err != nil {
		return "", err
	}
	return id, nil
}
