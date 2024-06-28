package utils

import "time"

// ToStringPointer transforma uma string em um ponteiro para string
func ToStringPointer(s string) *string {
	return &s
}

// FromIntPointer transforma um ponteiro para int em um int
func FromStringPointer(ptr *string) string {
	// Se o ponteiro for nulo, retornamos zero (ou qualquer outro valor padrão que você preferir)
	if ptr == nil {
		return ""
	}
	return *ptr
}

// ToIntPointer transforma um int em um ponteiro para int
func ToIntPointer(i int) *int {
	return &i
}

// FromIntPointer transforma um ponteiro para int em um int
func FromIntPointer(ptr *int) int {
	// Se o ponteiro for nulo, retornamos zero (ou qualquer outro valor padrão que você preferir)
	if ptr == nil {
		return 0
	}
	return *ptr
}

// FromTimePointer transforma um ponteiro para time.Time em um time.Time
func FromTimePointer(ptr *time.Time) time.Time {
	// Se o ponteiro for nulo, retornamos o valor zero de time.Time
	if ptr == nil {
		return time.Time{}
	}
	return *ptr
}

// ToTimePointer transforma um time.Time em um ponteiro para time.Time
func ToTimePointer(t time.Time) *time.Time {
	return &t
}

func ToBoolPointer(b bool) *bool {
	return &b
}
