// Licensed to the Apache Software Foundation (ASF) under one or more
// contributor license agreements.  See the NOTICE file distributed with
// this work for additional information regarding copyright ownership.
// The ASF licenses this file to You under the Apache License, Version 2.0
// (the "License"); you may not use this file except in compliance with
// the License.  You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package apisix

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	v1 "github.com/api7/ingress-controller/pkg/types/apisix/v1"
)

// listRepsonse is the unified LIST response mapping of APISIX.
type listResponse struct {
	Count string `json:"count"`
	Node  node   `json:"node"`
}

type node struct {
	Key   string `json:"key"`
	Items items  `json:"nodes"`
}

type items []item

// items implements json.Unmarshaler interface.
// lua-cjson doesn't distinguish empty array and table,
// and by default empty array will be encoded as '{}'.
// We have to maintain the compatibility.
func (items *items) UnmarshalJSON(p []byte) error {
	if p[0] == '{' {
		if len(p) != 2 {
			return errors.New("unexpected non-empty object")
		}
		return nil
	}
	var data []item
	if err := json.Unmarshal(p, &data); err != nil {
		return err
	}
	*items = data
	return nil
}

type item struct {
	Key   string          `json:"key"`
	Value json.RawMessage `json:"value"`
}

type routeItem struct {
	UpstreamId *string                `json:"upstream_id"`
	ServiceId  *string                `json:"service_id"`
	Host       *string                `json:"host"`
	URI        *string                `json:"uri"`
	Desc       *string                `json:"desc"`
	Methods    []*string              `json:"methods"`
	Plugins    map[string]interface{} `json:"plugins"`
}

// route decodes item.value and converts it to v1.Route.
func (i *item) route(group string) (*v1.Route, error) {
	list := strings.Split(i.Key, "/")
	if len(list) < 1 {
		return nil, fmt.Errorf("bad route config key: %s", i.Key)
	}

	var route routeItem
	if err := json.Unmarshal(i.Value, &route); err != nil {
		return nil, err
	}

	name := route.Desc
	fullName := "unknown"
	if name != nil {
		fullName = *name
	}
	if group != "" {
		fullName = group + "_" + fullName
	}

	return &v1.Route{
		ID:         &list[len(list)-1],
		Group:      &group,
		FullName:   &fullName,
		Name:       route.Desc,
		Host:       route.Host,
		Path:       route.URI,
		Methods:    route.Methods,
		UpstreamId: route.UpstreamId,
		ServiceId:  route.ServiceId,
		Plugins:    (*v1.Plugins)(&route.Plugins),
	}, nil
}
