/*
 * Copyright 2020 EPAM Systems
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package config

import "time"

const (
	NvidiaResourceName                      = "nvidia.com/gpu"
	RepositoryKubernetesType RepositoryType = "kubernetes"
	RepositoryVaultType      RepositoryType = "vault"
	RepositoryMemoryType     RepositoryType = "memory"
	RepositoryPostgresType   RepositoryType = "postgres"

	StoragePostgres StorageType = "postgres"
)

type RepositoryType string

type StorageType string

type ExternalUrl struct {
	// Human-readable name
	Name string `json:"name"`
	// Link to a resource
	URL string `json:"url"`
	// Optional link to an image which represents a type of the resource, for example the logo of Grafana
	ImageURL string `json:"imageUrl"`
}

type CommonConfig struct {
	// The collection of external urls, for example: metrics, edge, service catalog and so on
	ExternalURLs []ExternalUrl `json:"externalUrls"`
	// Kubernetes can consume the GPU resource in the <vendor>.com/gpu format.
	// For example, amd.com/gpu or nvidia.com/gpu.
	ResourceGPUName string `json:"resourceGpuName"`
	// Version of ODAHU platform
	Version string `json:"version"`
	// Database connection string
	DatabaseConnectionString string `json:"databaseConnectionString"`
	// OpenID token url
	OAuthOIDCTokenEndpoint string `json:"oauthOidcTokenEndpoint"`
	// How often launch new training
	LaunchPeriod time.Duration `json:"launchPeriod"`
	// Graceful shutdown timeout
	GracefulTimeout time.Duration `json:"gracefulTimeout"`
}

func NewDefaultCommonConfig() CommonConfig {
	return CommonConfig{
		ExternalURLs:    []ExternalUrl{},
		ResourceGPUName: NvidiaResourceName,
		Version:         "develop",
		LaunchPeriod:    time.Second * 3,
		GracefulTimeout: time.Second * 5,
	}
}
