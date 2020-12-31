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
	"context"
	"errors"
	"net/http"
	"strings"
	"time"

	v1 "github.com/api7/ingress-controller/pkg/types/apisix/v1"
)

const (
	_defaultTimeout = 5 * time.Second
)

// Options contains parameters to customize APISIX client.
type Options struct {
	AdminKey string
	BaseURL  string
	Timeout  time.Duration
}

// Interface is the unified client tool to communicate with APISIX.
type Client interface {
	Route() Route
	Upstream() Upstream
}

// Route is the specific client interface to take over the create, update,
// listResource and delete for APISIX's Route resource.
type Route interface {
	List(context.Context, string) ([]*v1.Route, error)
	Create(context.Context, *v1.Route) (*v1.Route, error)
	Delete(context.Context, *v1.Route) error
	Update(context.Context, *v1.Route) error
}

// Upstream is the specific client interface to take over the create, update,
// listResource and delete for APISIX's Upstream resource.
type Upstream interface {
	List(context.Context, string) ([]*v1.Upstream, error)
	Create(context.Context, *v1.Upstream) (*v1.Upstream, error)
	Delete(context.Context, *v1.Upstream) error
	Update(context.Context, *v1.Upstream) error
}

type client struct {
	stub     *stub
	route    Route
	upstream Upstream
}

// NewClient creates an APISIX client to perform resources change pushing.
func NewClient(o *Options) (Client, error) {
	if o.BaseURL == "" {
		return nil, errors.New("empty base_url")
	}
	if o.Timeout == time.Duration(0) {
		o.Timeout = _defaultTimeout
	}
	o.BaseURL = strings.TrimSuffix(o.BaseURL, "/")

	stub := &stub{
		baseURL:  o.BaseURL,
		adminKey: o.AdminKey,
		cli: &http.Client{
			Timeout: o.Timeout,
			Transport: &http.Transport{
				ResponseHeaderTimeout: o.Timeout,
				ExpectContinueTimeout: o.Timeout,
			},
		},
	}
	cli := &client{
		stub:     stub,
		route:    newRouteClient(stub),
		upstream: newUpstreamClient(stub),
	}
	return cli, nil
}

// Route implements Client interface.
func (c *client) Route() Route {
	return c.route
}

// Upstream implements Client interface.
func (c *client) Upstream() Upstream {
	return c.upstream
}