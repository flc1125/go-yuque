package yuque

import "fmt"

// Ptr is a helper function to get the pointer of a value.
func Ptr[T any](v T) *T {
	return &v
}

// parseID is a helper function to parse the id.
// It supports string and int type.
func parseID(id any) (string, error) {
	switch v := id.(type) {
	case string:
		return v, nil
	case int:
		return fmt.Sprintf("%d", v), nil
	}
	return "", fmt.Errorf("invalid id type: %T", id)
}
