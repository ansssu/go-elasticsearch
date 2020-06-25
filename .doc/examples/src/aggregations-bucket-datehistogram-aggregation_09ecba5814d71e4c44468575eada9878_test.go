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

// <https://github.com/elastic/elasticsearch/blob/master/docs/reference/aggregations/bucket/datehistogram-aggregation.asciidoc#L214>
//
// --------------------------------------------------------------------------------
// POST /sales/_search?size=0
// {
//     "aggs" : {
//         "sales_over_time" : {
//             "date_histogram" : {
//                 "field" : "date",
//                 "fixed_interval" : "30d"
//             }
//         }
//     }
// }
// --------------------------------------------------------------------------------

func Test_aggregations_bucket_datehistogram_aggregation_09ecba5814d71e4c44468575eada9878(t *testing.T) {
	es, _ := elasticsearch.NewDefaultClient()

	// tag:09ecba5814d71e4c44468575eada9878[]
	res, err := es.Search(
		es.Search.WithIndex("sales"),
		es.Search.WithBody(strings.NewReader(`{
		  "aggs": {
		    "sales_over_time": {
		      "date_histogram": {
		        "field": "date",
		        "fixed_interval": "30d"
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
	// end:09ecba5814d71e4c44468575eada9878[]
}
