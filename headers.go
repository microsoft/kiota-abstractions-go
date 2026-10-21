package abstractions

import "strings"

type header struct {
	headers map[string]map[string]struct{}
}

type void struct{}

var voidInstance void

func normalizeHeaderKey(key string) string {
	return strings.ToLower(strings.TrimSpace(key))
}

// Add adds a new header or appends a new value to an existing header
func (r *header) Add(key string, value string, additionalValues ...string) {
	if r == nil {
		return
	}
	normalizedKey := normalizeHeaderKey(key)
	if normalizedKey == "" {
		return
	}
	hasValue := value != ""
	hasAdditional := false
	for _, v := range additionalValues {
		if v != "" {
			hasAdditional = true
			break
		}
	}
	if !hasValue && !hasAdditional {
		return
	}
	if r.headers == nil {
		r.headers = make(map[string]map[string]struct{})
	}
	if r.headers[normalizedKey] == nil {
		r.headers[normalizedKey] = make(map[string]struct{})
	}
	if hasValue {
		r.headers[normalizedKey][value] = voidInstance
	}
	for _, v := range additionalValues {
		if v != "" {
			r.headers[normalizedKey][v] = voidInstance
		}
	}
}

// Get returns the values for the specific header
func (r *header) Get(key string) []string {
	if r == nil || r.headers == nil {
		return nil
	}
	normalizedKey := normalizeHeaderKey(key)
	if r.headers[normalizedKey] == nil {
		return make([]string, 0)
	}
	values := make([]string, 0, len(r.headers[normalizedKey]))
	for k := range r.headers[normalizedKey] {
		values = append(values, k)
	}
	return values
}

// Remove removes the specific header and all its values
func (r *header) Remove(key string) {
	if r == nil || r.headers == nil {
		return
	}
	normalizedKey := normalizeHeaderKey(key)
	delete(r.headers, normalizedKey)
}

// RemoveValue removes the value for the specific header
func (r *header) RemoveValue(key string, value string) {
	if r == nil || r.headers == nil {
		return
	}
	normalizedKey := normalizeHeaderKey(key)
	if r.headers[normalizedKey] == nil {
		return
	}
	delete(r.headers[normalizedKey], value)
	if len(r.headers[normalizedKey]) == 0 {
		delete(r.headers, normalizedKey)
	}
}

// ContainsKey checks if the key exists in the headers
func (r *header) ContainsKey(key string) bool {
	if r == nil || r.headers == nil {
		return false
	}
	normalizedKey := normalizeHeaderKey(key)
	return r.headers[normalizedKey] != nil
}

// Clear clears all headers
func (r *header) Clear() {
	if r == nil {
		return
	}
	r.headers = nil
}

// AddAll adds all headers from the other headers
func (r *header) AddAll(other *header) {
	if r == nil || other == nil || other.headers == nil {
		return
	}
	for k, v := range other.headers {
		for k2 := range v {
			r.Add(k, k2)
		}
	}
}

// ListKeys returns all the keys in the headers
func (r *header) ListKeys() []string {
	if r == nil || r.headers == nil {
		return make([]string, 0)
	}
	keys := make([]string, 0, len(r.headers))
	for k := range r.headers {
		keys = append(keys, k)
	}
	return keys
}

