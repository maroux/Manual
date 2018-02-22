// WARNING: Do not modify. This file is machine-generated by manual.

package models

// Base is machine-generated.
//
// The base interface.
type Base interface {
	// This is a foo property.
	GetFoo() (string, error)
}

// Baz is machine-generated.
type Baz struct {
	Qux bool `json:"qux"`

	// This is a foo property.
	Foo string `json:"foo"`
}

// Baz's conformance to Base is machine-generated.

// GetFoo is machine-generated.
//
// Gets Baz's Foo property.
func (o Baz) GetFoo() (string, error) {
	return o.Foo, nil
}

// Corge is machine-generated.
//
// This is the corge.
type Corge struct {
	Bar Baz `json:"bar"`
}

// List is machine-generated.
//
// A list.
type List interface {
	// A list of items.
	GetItems() ([]interface{}, error)
}

// ListBaz is machine-generated.
//
// A list of Baz.
type ListBaz struct {
	// Baz items.
	//
	// Items is nullable.
	Items *[]Baz `json:"items"`
}

// ListBaz's conformance to List is machine-generated.

// GetItems is machine-generated.
//
// Gets ListBaz's Items as List's property type.
func (o ListBaz) GetItems() ([]interface{}, error) {
	items := make([]interface{}, len(o.Items))
	for _, item := range o.Items {
		items = append(items, item)
	}
	return items, nil
}
