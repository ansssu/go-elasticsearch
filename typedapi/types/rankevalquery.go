// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

// Code generated from the elasticsearch-specification DO NOT EDIT.
// https://github.com/elastic/elasticsearch-specification/tree/6e0fb6b929f337b62bf0676bdf503e061121fad2

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"strconv"
)

// RankEvalQuery type.
//
// https://github.com/elastic/elasticsearch-specification/blob/6e0fb6b929f337b62bf0676bdf503e061121fad2/specification/_global/rank_eval/types.ts#L111-L114
type RankEvalQuery struct {
	Query Query `json:"query"`
	Size  *int  `json:"size,omitempty"`
}

func (s *RankEvalQuery) UnmarshalJSON(data []byte) error {

	dec := json.NewDecoder(bytes.NewReader(data))

	for {
		t, err := dec.Token()
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			return err
		}

		switch t {

		case "query":
			if err := dec.Decode(&s.Query); err != nil {
				return err
			}

		case "size":

			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return err
				}
				s.Size = &value
			case float64:
				f := int(v)
				s.Size = &f
			}

		}
	}
	return nil
}

// NewRankEvalQuery returns a RankEvalQuery.
func NewRankEvalQuery() *RankEvalQuery {
	r := &RankEvalQuery{}

	return r
}
