package namespace

import (
	pb "github.com/authzed/spicedb/pkg/REDACTEDapi/api"
)

// Manager is a subset of the datastore interface that can read (and possibly cache) namespaces.
type Manager interface {
	// ReadNamespace reads a namespace definition and version and returns it if found.
	ReadNamespace(nsName string) (*pb.NamespaceDefinition, uint64, error)

	// CheckNamespaceAndRelation checks that the specified namespace and relation exist in the
	// datastore.
	CheckNamespaceAndRelation(namespace, relation string, allowEllipsis bool) error
}
