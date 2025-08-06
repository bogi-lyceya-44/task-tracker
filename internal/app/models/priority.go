//go:generate go-enum --output-suffix=.generated --marshal --ptr --names --values

package models

// ENUM(low, medium, high, critical)
type Priority int32
