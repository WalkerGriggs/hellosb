package state

import (
	memdb "github.com/hashicorp/go-memdb"
)

type schemaFactory func() *memdb.TableSchema

func (s *StateStore) setupDB() (*memdb.MemDB, error) {
	return memdb.NewMemDB(stateStoreSchema())
}

func stateStoreSchema() *memdb.DBSchema {
	db := &memdb.DBSchema{
		Tables: make(map[string]*memdb.TableSchema),
	}

	factories := []schemaFactory{
		serviceInstanceTableSchema,
		serviceBindingTableSchema,
	}

	for _, fn := range factories {
		schema := fn()
		db.Tables[schema.Name] = schema
	}

	return db
}

func serviceInstanceTableSchema() *memdb.TableSchema {
	return &memdb.TableSchema{
		Name: "service_instance",
		Indexes: map[string]*memdb.IndexSchema{
			"id": &memdb.IndexSchema{
				Name:    "id",
				Unique:  true,
				Indexer: &memdb.StringFieldIndex{Field: "id"},
			},
		},
	}
}

func serviceBindingTableSchema() *memdb.TableSchema {
	return &memdb.TableSchema{
		Name: "service_binding",
		Indexes: map[string]*memdb.IndexSchema{
			"id": &memdb.IndexSchema{
				Name:    "id",
				Unique:  true,
				Indexer: &memdb.StringFieldIndex{Field: "id"},
			},
		},
	}
}
