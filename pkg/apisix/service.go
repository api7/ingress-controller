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
	"bytes"
	"context"
	"encoding/json"

	"github.com/api7/ingress-controller/pkg/log"
	v1 "github.com/api7/ingress-controller/pkg/types/apisix/v1"
	"go.uber.org/zap"
)

type serviceClient struct {
	url         string
	clusterName string
	cluster     *cluster
}

type serviceItem struct {
	UpstreamId *string                 `json:"upstream_id,omitempty"`
	Plugins    *map[string]interface{} `json:"plugins,omitempty"`
	Desc       *string                 `json:"desc,omitempty"`
}

func newServiceClient(c *cluster) Service {
	return &serviceClient{
		url:         c.baseURL + "/services",
		clusterName: c.name,
		cluster:     c,
	}
}

// Get only looks up the cache, it's not necessary to access APISIX, since all resources
// are created by Create, which reflects the change to cache in turn, so if resource
// is not in cache, it's not in APISIX either.
func (s *serviceClient) Get(_ context.Context, fullname string) (*v1.Service, error) {
	return s.cluster.cache.GetService(fullname)
}

// List is only used in cache warming up. So here just pass through
// to APISIX.
func (s *serviceClient) List(ctx context.Context) ([]*v1.Service, error) {
	log.Infow("try to list services in APISIX", zap.String("url", s.url))

	upsItems, err := s.cluster.listResource(ctx, s.url)
	if err != nil {
		log.Errorf("failed to list upstreams: %s", err)
		return nil, err
	}

	var items []*v1.Service
	for i, item := range upsItems.Node.Items {
		svc, err := item.service(s.clusterName)
		if err != nil {
			log.Errorw("failed to convert service item",
				zap.String("url", s.url),
				zap.String("service_key", item.Key),
				zap.Error(err),
			)
			return nil, err
		}
		items = append(items, svc)
		log.Infof("list service #%d, body: %s", i, string(item.Value))
	}
	return items, nil
}

func (s *serviceClient) Create(ctx context.Context, obj *v1.Service) (*v1.Service, error) {
	if err := s.cluster.Ready(ctx); err != nil {
		return nil, err
	}
	log.Infow("try to create service", zap.String("full_name", *obj.FullName))

	body, err := json.Marshal(serviceItem{
		UpstreamId: obj.UpstreamId,
		Plugins:    (*map[string]interface{})(obj.Plugins),
		Desc:       obj.Name,
	})
	if err != nil {
		return nil, err
	}

	log.Infow("creating service", zap.ByteString("body", body), zap.String("url", s.url))
	resp, err := s.cluster.createResource(ctx, s.url, bytes.NewReader(body))
	if err != nil {
		log.Errorf("failed to create service: %s", err)
		return nil, err
	}
	var clusterName string
	if obj.Group != nil {
		clusterName = *obj.Group
	}
	svc, err := resp.Item.service(clusterName)
	if err != nil {
		return nil, err
	}
	if err := s.cluster.cache.InsertService(svc); err != nil {
		log.Errorf("failed to reflect service create to cache: %s", err)
		return nil, err
	}
	return svc, nil
}

func (s *serviceClient) Delete(ctx context.Context, obj *v1.Service) error {
	if err := s.cluster.Ready(ctx); err != nil {
		return err
	}
	log.Infof("delete service, id:%s", *obj.ID)
	url := s.url + "/" + *obj.ID
	if err := s.cluster.deleteResource(ctx, url); err != nil {
		return err
	}
	if err := s.cluster.cache.DeleteService(obj); err != nil {
		log.Errorf("failed to reflect service delete to cache: %s", err)
		return err
	}
	return nil
}

func (s *serviceClient) Update(ctx context.Context, obj *v1.Service) (*v1.Service, error) {
	if err := s.cluster.Ready(ctx); err != nil {
		return nil, err
	}
	log.Infof("update service, id:%s", *obj.ID)

	body, err := json.Marshal(serviceItem{
		UpstreamId: obj.UpstreamId,
		Plugins:    (*map[string]interface{})(obj.Plugins),
		Desc:       obj.Name,
	})
	if err != nil {
		return nil, err
	}

	url := s.url + "/" + *obj.ID
	log.Infow("creating service", zap.ByteString("body", body), zap.String("url", url))
	resp, err := s.cluster.updateResource(ctx, url, bytes.NewReader(body))
	if err != nil {
		return nil, err
	}
	var clusterName string
	if obj.Group != nil {
		clusterName = *obj.Group
	}
	svc, err := resp.Item.service(clusterName)
	if err != nil {
		return nil, err
	}
	if err := s.cluster.cache.InsertService(obj); err != nil {
		log.Errorf("failed to reflect service update to cache: %s", err)
		return nil, err
	}
	return svc, nil
}
