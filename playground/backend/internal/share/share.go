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
	"time"
)

const (
	idLength = 11
)

type Origin int32

const (
	TOUR_OF_BEAM Origin = 0 //will be used in Tour of Beam project
	PLAYGROUND   Origin = 1
)

func (s Origin) Value() int32 {
	return int32(s)
}

type Code struct {
	Name     string `firestore:"name"`
	Code     string `firestore:"code"`
	CntxLine int32  `firestore:"cntxLine"`
	IsMain   bool   `firestore:"isMain"`
	SnpId    string `firestore:"snpId"`
}

type Snippet struct {
	Salt       string    `firestore:"-"`
	OwnerId    string    `firestore:"ownerId"`
	Sdk        pb.Sdk    `firestore:"sdk"`
	PipeOpts   string    `firestore:"pipeOpts"`
	Created    time.Time `firestore:"created"`
	LVisited   time.Time `firestore:"lVisited"`
	Origin     Origin    `firestore:"origin"`
	VisitCount int       `firestore:"visitCount"`
	Codes      []Code    `firestore:"-"`
}

// ID generates id according to content of a snippet
func (s *Snippet) ID() (string, error) {
	hash := sha256.New()
	if _, err := io.WriteString(hash, s.Salt); err != nil {
		logger.Errorf("ID(): error while writing ID and salt: %s", err.Error())
		return "", errors.InternalError("Error during ID generation", "Error with writing ID and salt")
	}
	var codes []string
	for _, v := range s.Codes {
		codes = append(codes, v.Code)
	}
	sort.Strings(codes)
	var content string
	for i, v := range codes {
		content += v
		if i == len(codes)-1 {
			content += fmt.Sprintf("%v%s", s.Sdk, s.PipeOpts)
		}
	}
	hash.Write([]byte(content))
	sum := hash.Sum(nil)
	b := make([]byte, base64.URLEncoding.EncodedLen(len(sum)))
	base64.URLEncoding.Encode(b, sum)
	hashLen := idLength
	for hashLen <= len(b) && b[hashLen-1] == '_' {
		hashLen++
	}
	return string(b)[:hashLen], nil
}

// ID generates id according to content of a code and its name
func (c *Code) ID(salt string) (string, error) {
	hash := sha256.New()
	if _, err := io.WriteString(hash, salt); err != nil {
		logger.Errorf("ID(): error while writing ID and salt: %s", err.Error())
		return "", errors.InternalError("Error during ID generation", "Error with writing ID and salt")
	}
	content := fmt.Sprintf("%s%s%s", c.Code, c.Name, c.SnpId)
	hash.Write([]byte(content))
	sum := hash.Sum(nil)
	b := make([]byte, base64.URLEncoding.EncodedLen(len(sum)))
	base64.URLEncoding.Encode(b, sum)
	hashLen := idLength
	for hashLen <= len(b) && b[hashLen-1] == '_' {
		hashLen++
	}
	return string(b)[:hashLen], nil
}
