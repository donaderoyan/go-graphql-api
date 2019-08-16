package schema

import (
	gql "github.com/mattdamon108/gqlmerge/lib"
)

func NewSchema() string {
	schema := gql.Merge("  ", "./src/schema")
	return *schema
}
