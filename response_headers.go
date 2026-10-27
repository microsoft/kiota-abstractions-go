package abstractions

// ResponseHeaders represents a collection of response headers
type ResponseHeaders struct {
	header
}

// NewResponseHeaders creates a new ResponseHeaders
func NewResponseHeaders() *ResponseHeaders {
	return &ResponseHeaders{
		header{make(map[string]map[string]struct{})},
	}
}

// Add adds a new header or appends a new value to an existing header
func (r *ResponseHeaders) Add(key string, value string, additionalValues ...string) {
	if r == nil {
		return
	}
	r.header.Add(key, value, additionalValues...)
}

// Get returns the values for the specific header
func (r *ResponseHeaders) Get(key string) []string {
	if r == nil {
		return nil
	}
	return r.header.Get(key)
}

// Remove removes the specific header and all its values
func (r *ResponseHeaders) Remove(key string) {
	if r == nil {
		return
	}
	r.header.Remove(key)
}

// RemoveValue removes the value for the specific header
func (r *ResponseHeaders) RemoveValue(key string, value string) {
	if r == nil {
		return
	}
	r.header.RemoveValue(key, value)
}

// ContainsKey checks if the key exists in the headers
func (r *ResponseHeaders) ContainsKey(key string) bool {
	if r == nil {
		return false
	}
	return r.header.ContainsKey(key)
}

// Clear clears all headers
func (r *ResponseHeaders) Clear() {
	if r == nil {
		return
	}
	r.header.Clear()
}

// ListKeys returns all the keys in the headers
func (r *ResponseHeaders) ListKeys() []string {
	if r == nil {
		return make([]string, 0)
	}
	return r.header.ListKeys()
}

// AddAll adds all headers from the other headers
func (r *ResponseHeaders) AddAll(other *ResponseHeaders) {
	if r == nil || other == nil || other.headers == nil {
		return
	}
	r.header.AddAll(&other.header)
}

// TryAdd adds the header if it's not already present. It returns true if the header was added,
// or false if it already existed, the receiver is nil, or the key/value was rejected.
func (r *ResponseHeaders) TryAdd(key string, value string) bool {
	if r == nil {
		return false
	}
	if r.ContainsKey(key) {
		return false
	}

	r.Add(key, value)
	return r.ContainsKey(key)
}
