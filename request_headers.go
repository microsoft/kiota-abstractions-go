package abstractions

// RequestHeaders represents a collection of request headers
type RequestHeaders struct {
	header
}

// NewRequestHeaders creates a new RequestHeaders
func NewRequestHeaders() *RequestHeaders {
	return &RequestHeaders{
		header{make(map[string]map[string]struct{})},
	}
}

// Add adds a new header or appends a new value to an existing header
func (r *RequestHeaders) Add(key string, value string, additionalValues ...string) {
	if r == nil {
		return
	}
	r.header.Add(key, value, additionalValues...)
}

// Get returns the values for the specific header
func (r *RequestHeaders) Get(key string) []string {
	if r == nil {
		return nil
	}
	return r.header.Get(key)
}

// Remove removes the specific header and all its values
func (r *RequestHeaders) Remove(key string) {
	if r == nil {
		return
	}
	r.header.Remove(key)
}

// RemoveValue removes the value for the specific header
func (r *RequestHeaders) RemoveValue(key string, value string) {
	if r == nil {
		return
	}
	r.header.RemoveValue(key, value)
}

// ContainsKey checks if the key exists in the headers
func (r *RequestHeaders) ContainsKey(key string) bool {
	if r == nil {
		return false
	}
	return r.header.ContainsKey(key)
}

// Clear clears all headers
func (r *RequestHeaders) Clear() {
	if r == nil {
		return
	}
	r.header.Clear()
}

// ListKeys returns all the keys in the headers
func (r *RequestHeaders) ListKeys() []string {
	if r == nil {
		return make([]string, 0)
	}
	return r.header.ListKeys()
}

// AddAll adds all headers from the other headers
func (r *RequestHeaders) AddAll(other *RequestHeaders) {
	if r == nil || other == nil || other.headers == nil {
		return
	}
	for k, v := range other.headers {
		for k2 := range v {
			r.Add(k, k2)
		}
	}
}

// TryAdd adds the header if it's not already present
func (r *RequestHeaders) TryAdd(key string, value string) bool {
	if r == nil {
		return false
	}
	if r.ContainsKey(key) {
		return false
	}

	r.Add(key, value)
	return true
}