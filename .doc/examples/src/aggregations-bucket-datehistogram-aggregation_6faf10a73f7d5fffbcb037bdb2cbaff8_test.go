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

	"github.com/elastic/go-elasticsearch/v8"
)

var (
	_ = fmt.Printf
	_ = os.Stdout
	_ = elasticsearch.NewDefaultClient
)

// <https://github.com/elastic/elasticsearch/blob/master/docs/reference/aggregations/bucket/datehistogram-aggregation.asciidoc#L669>
//
// --------------------------------------------------------------------------------
// POST /sales/_search?size=0
// {
//     "aggs": {
//         "dayOfWeek": {
//             "terms": {
//                 "script": {
//                     "lang": "painless",
//                     "source": "doc['date'].value.dayOfWeekEnum.value"
//                 }
//             }
//         }
//     }
// }
// --------------------------------------------------------------------------------

func Test_aggregations_bucket_datehistogram_aggregation_6faf10a73f7d5fffbcb037bdb2cbaff8(t *testing.T) {
	es, _ := elasticsearch.NewDefaultClient()

	// tag:6faf10a73f7d5fffbcb037bdb2cbaff8[]
	res, err := es.Search(
		es.Search.WithIndex("sales"),
		es.Search.WithBody(strings.NewReader(`{
		  "aggs": {
		    "dayOfWeek": {
		      "terms": {
		        "script": {
		          "lang": "painless",
		          "source": "doc['date'].value.dayOfWeekEnum.value"
		        }
		      }
		    }
		  }
		}`)),
		es.Search.WithSize(0),
		es.Search.WithPretty(),
	)
	fmt.Println(res, err)
	if err != nil { // SKIP
		t.Fatalf("Error getting the response: %s", err) // SKIP
	} // SKIP
	defer res.Body.Close() // SKIP
	// end:6faf10a73f7d5fffbcb037bdb2cbaff8[]
}
