// Licensed to Elasticsearch B.V under one or more agreements.
// Elasticsearch B.V. licenses this file to you under the Apache 2.0 License.
// See the LICENSE file in the project root for more information.
//
// Code generated, DO NOT EDIT

package elasticsearch_test

import (
	"fmt"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/elastic/go-elasticsearch/v8"
)

var (
	_ = fmt.Printf
	_ = os.Stdout
	_ = elasticsearch.NewDefaultClient
)

// <https://github.com/elastic/elasticsearch/blob/master/docs/reference/docs/index_.asciidoc#L393>
//
// --------------------------------------------------------------------------------
// PUT twitter/_doc/1?timeout=5m
// {
//     "user" : "kimchy",
//     "post_date" : "2009-11-15T14:12:12",
//     "message" : "trying out Elasticsearch"
// }
// --------------------------------------------------------------------------------

func Test_docs_index__b918d6b798da673a33e49b94f61dcdc0(t *testing.T) {
	es, _ := elasticsearch.NewDefaultClient()

	// tag:b918d6b798da673a33e49b94f61dcdc0[]
	res, err := es.Index(
		"twitter",
		strings.NewReader(`{
		  "user": "kimchy",
		  "post_date": "2009-11-15T14:12:12",
		  "message": "trying out Elasticsearch"
		}`),
		es.Index.WithDocumentID("1"),
		es.Index.WithTimeout(time.Duration(300000000000)),
		es.Index.WithPretty(),
	)
	fmt.Println(res, err)
	if err != nil { // SKIP
		t.Fatalf("Error getting the response: %s", err) // SKIP
	} // SKIP
	defer res.Body.Close() // SKIP
	// end:b918d6b798da673a33e49b94f61dcdc0[]
}
