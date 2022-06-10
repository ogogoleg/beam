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

package share

import (
	pb "beam.apache.org/playground/backend/internal/api/v1"
	"beam.apache.org/playground/backend/internal/errors"
	"beam.apache.org/playground/backend/internal/logger"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"io"
	"sort"
	"strings"
	"time"
	"unicode"
)

type Origin int32

const (
	TOUR_OF_BEAM Origin = 0 //will be used in Tour of Beam project
	PLAYGROUND   Origin = 1
)

func (s Origin) Value() int32 {
	return int32(s)
}

type CodeDocument struct {
	Name     string `datastore:"name"`
	Code     string `datastore:"code"`
	CntxLine int32  `datastore:"cntxLine"`
	IsMain   bool   `datastore:"isMain"`
}

type SnippetDocument struct {
	OwnerId    string          `datastore:"ownerId"`
	Sdk        pb.Sdk          `datastore:"sdk"`
	PipeOpts   string          `datastore:"pipeOpts"`
	Created    time.Time       `datastore:"created"`
	LVisited   time.Time       `datastore:"lVisited"`
	Origin     Origin          `datastore:"origin"`
	VisitCount int             `datastore:"visitCount"`
	Codes      []*CodeDocument `datastore:"codes"`
}

type Snippet struct {
	Salt     string
	IdLength int
	Snippet  *SnippetDocument
}

// ID generates id according to content of a snippet
func (s *Snippet) ID() (string, error) {
	hash := sha256.New()
	if _, err := io.WriteString(hash, s.Salt); err != nil {
		logger.Errorf("ID(): error while writing ID and salt: %s", err.Error())
		return "", errors.InternalError("Error during ID generation", "Error with writing ID and salt")
	}
	var codes []string
	for _, v := range s.Snippet.Codes {
		codes = append(codes, spaceStringsBuilder(v.Code)+spaceStringsBuilder(v.Name))
	}
	sort.Strings(codes)
	var content string
	for i, v := range codes {
		content += v
		if i == len(codes)-1 {
			content += fmt.Sprintf("%v%s", s.Snippet.Sdk, spaceStringsBuilder(s.Snippet.PipeOpts))
		}
	}
	hash.Write([]byte(content))
	sum := hash.Sum(nil)
	b := make([]byte, base64.URLEncoding.EncodedLen(len(sum)))
	base64.URLEncoding.Encode(b, sum)
	hashLen := s.IdLength
	for hashLen <= len(b) && b[hashLen-1] == '_' {
		hashLen++
	}
	return string(b)[:hashLen], nil
}

// ID generates id according to content of a code and its name
func (c *CodeDocument) ID(snip *Snippet) (string, error) {
	hash := sha256.New()
	if _, err := io.WriteString(hash, snip.Salt); err != nil {
		logger.Errorf("ID(): error while writing ID and salt: %s", err.Error())
		return "", errors.InternalError("Error during ID generation", "Error with writing ID and salt")
	}
	content := fmt.Sprintf("%s%s", spaceStringsBuilder(c.Code), spaceStringsBuilder(c.Name))
	hash.Write([]byte(content))
	sum := hash.Sum(nil)
	b := make([]byte, base64.URLEncoding.EncodedLen(len(sum)))
	base64.URLEncoding.Encode(b, sum)
	hashLen := snip.IdLength
	for hashLen <= len(b) && b[hashLen-1] == '_' {
		hashLen++
	}
	return string(b)[:hashLen], nil
}

func spaceStringsBuilder(str string) string {
	var b strings.Builder
	b.Grow(len(str))
	for _, ch := range str {
		if !unicode.IsSpace(ch) {
			b.WriteRune(ch)
		}
	}
	return b.String()
}
