// WARNING: Do not modify. This file is machine-generated by swagger-gen.

package models

// Base is machine-generated.
//
// The base model object
type Base interface {
	// The date-time that this object was created at.
	GetCreatedAt() (*JSONDateTime, error)

	// The discriminator for the base model.
	GetType() (string, error)
}
