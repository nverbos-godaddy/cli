package storetest

import (
	"fmt"
	"os"
	"time"

	"gopkg.in/yaml.v3"

	"github.com/openfga/openfga/pkg/server"
	"github.com/openfga/openfga/pkg/storage/memory"
)

// ServerConfig mirrors a subset of the OpenFGA server configuration
// (see https://github.com/openfga/openfga/blob/main/.config-schema.json)
// that is relevant for the embedded test server used by `fga model test`.
type ServerConfig struct {
	MaxTypesPerAuthorizationModel    *int           `yaml:"maxTypesPerAuthorizationModel"`
	MaxAuthorizationModelSizeInBytes *int           `yaml:"maxAuthorizationModelSizeInBytes"`
	MaxTuplesPerWrite                *int           `yaml:"maxTuplesPerWrite"`
	ResolveNodeLimit                 *uint32        `yaml:"resolveNodeLimit"`
	ResolveNodeBreadthLimit          *uint32        `yaml:"resolveNodeBreadthLimit"`
	ListObjectsDeadline              *time.Duration `yaml:"listObjectsDeadline"`
	ListObjectsMaxResults            *uint32        `yaml:"listObjectsMaxResults"`
	ListUsersDeadline                *time.Duration `yaml:"listUsersDeadline"`
	ListUsersMaxResults              *uint32        `yaml:"listUsersMaxResults"`
	RequestTimeout                   *time.Duration `yaml:"requestTimeout"`
}

// ReadServerConfigFromFile reads a server configuration YAML file and returns
// a ServerConfig. If the path is empty, a zero-value config is returned.
func ReadServerConfigFromFile(path string) (ServerConfig, error) {
	var config ServerConfig

	if path == "" {
		return config, nil
	}

	data, err := os.ReadFile(path)
	if err != nil {
		return config, fmt.Errorf("failed to read server config file %s: %w", path, err)
	}

	if err := yaml.Unmarshal(data, &config); err != nil {
		return config, fmt.Errorf("failed to parse server config file %s: %w", path, err)
	}

	return config, nil
}

// MemoryOptions returns storage options for the in-memory datastore.
func (c *ServerConfig) MemoryOptions() []memory.StorageOption {
	var opts []memory.StorageOption

	if c.MaxTypesPerAuthorizationModel != nil {
		opts = append(opts, memory.WithMaxTypesPerAuthorizationModel(*c.MaxTypesPerAuthorizationModel))
	}

	if c.MaxTuplesPerWrite != nil {
		opts = append(opts, memory.WithMaxTuplesPerWrite(*c.MaxTuplesPerWrite))
	}

	return opts
}

// ServerOptions returns server options for the embedded OpenFGA server.
func (c *ServerConfig) ServerOptions() []server.OpenFGAServiceV1Option {
	var opts []server.OpenFGAServiceV1Option

	if c.MaxAuthorizationModelSizeInBytes != nil {
		opts = append(opts, server.WithMaxAuthorizationModelSizeInBytes(*c.MaxAuthorizationModelSizeInBytes))
	}

	if c.ResolveNodeLimit != nil {
		opts = append(opts, server.WithResolveNodeLimit(*c.ResolveNodeLimit))
	}

	if c.ResolveNodeBreadthLimit != nil {
		opts = append(opts, server.WithResolveNodeBreadthLimit(*c.ResolveNodeBreadthLimit))
	}

	if c.ListObjectsDeadline != nil {
		opts = append(opts, server.WithListObjectsDeadline(*c.ListObjectsDeadline))
	}

	if c.ListObjectsMaxResults != nil {
		opts = append(opts, server.WithListObjectsMaxResults(*c.ListObjectsMaxResults))
	}

	if c.ListUsersDeadline != nil {
		opts = append(opts, server.WithListUsersDeadline(*c.ListUsersDeadline))
	}

	if c.ListUsersMaxResults != nil {
		opts = append(opts, server.WithListUsersMaxResults(*c.ListUsersMaxResults))
	}

	if c.RequestTimeout != nil {
		opts = append(opts, server.WithRequestTimeout(*c.RequestTimeout))
	}

	return opts
}
