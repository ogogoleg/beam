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

package local_db

import (
	"beam.apache.org/playground/backend/internal/share"
	"context"
	"fmt"
	"sync"
)

type LocalDB struct {
	sync.RWMutex
	items map[string]interface{}
}

func New() (*LocalDB, error) {
	items := make(map[string]interface{})
	ls := &LocalDB{items: items}
	return ls, nil
}

// PutSnippet puts the snippet to local database
func (f *LocalDB) PutSnippet(_ context.Context, id string, snip *share.Snippet) error {
	f.Lock()
	defer f.Unlock()
	f.items[id] = snip
	return nil
}

// GetSnippet returns the code snippet
func (f *LocalDB) GetSnippet(_ context.Context, id string) (*share.Snippet, error) {
	f.RLock()
	value, found := f.items[id]
	if !found {
		f.RUnlock()
		return nil, fmt.Errorf("value with id: %s not found", id)
	}
	f.RUnlock()
	snippet, _ := value.(*share.Snippet)
	return snippet, nil
}
