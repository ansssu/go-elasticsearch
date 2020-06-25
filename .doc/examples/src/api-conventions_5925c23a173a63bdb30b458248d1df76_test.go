// Licensed to Elasticsearch B.V under one or more agreements.
// Elasticsearch B.V. licenses this file to you under the Apache 2.0 License.
// See the LICENSE file in the project root for more information.
//
// Code generated, DO NOT EDIT

package elasticsearch_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/elastic/go-elasticsearch/v8"
)

var (
	_ = fmt.Printf
	_ = os.Stdout
	_ = elasticsearch.NewDefaultClient
)

// <https://github.com/elastic/elasticsearch/blob/master/docs/reference/api-conventions.asciidoc#L407>
//
// --------------------------------------------------------------------------------
// GET twitter/_settings?flat_settings=false
// --------------------------------------------------------------------------------

func Test_api_conventions_5925c23a173a63bdb30b458248d1df76(t *testing.T) {
	es, _ := elasticsearch.NewDefaultClient()

	// tag:5925c23a173a63bdb30b458248d1df76[]
	res, err := es.Indices.GetSettings(
		es.Indices.GetSettings.WithIndex([]string{"twitter"}...),
		es.Indices.GetSettings.WithFlatSettings(false),
	)
	fmt.Println(res, err)
	if err != nil { // SKIP
		t.Fatalf("Error getting the response: %s", err) // SKIP
	} // SKIP
	defer res.Body.Close() // SKIP
	// end:5925c23a173a63bdb30b458248d1df76[]
}
