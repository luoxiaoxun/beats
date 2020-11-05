// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

// Code generated by beats/dev-tools/cmd/asset/asset.go - DO NOT EDIT.

package elasticsearch

import (
	"github.com/elastic/beats/v7/libbeat/asset"
)

func init() {
	if err := asset.SetFields("metricbeat", "elasticsearch", asset.ModuleFieldsPri, AssetElasticsearch); err != nil {
		panic(err)
	}
}

// AssetElasticsearch returns asset data.
// This is the base64 encoded gzipped contents of module/elasticsearch.
func AssetElasticsearch() string {
	return "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"
}
