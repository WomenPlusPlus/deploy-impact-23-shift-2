package utils

func IIF[T any](condition bool, yes, no T) T {
	if condition {
		return yes
	}
	return no
}

func PTR[T any](v T) *T {
	return &v
}
