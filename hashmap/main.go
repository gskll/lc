package hashmap

import (
	"fmt"
	"math"
)

/*
Tombstone Mechanism in Open Addressing Hash Tables

In a hash table using open addressing (like this one), we need tombstones (marked as 'Deleted' entries)
rather than just clearing deleted entries. Here's why:

Example scenario:
Let's say items A, B, C, D all hash to index 0, and due to probing end up like this:
Index: 0    1    2    3
Items: A -> B -> C -> D

If we delete B and just clear its slot (set to Empty):
Index: 0    1    2    3
Items: A -> _ -> C -> D

Now if we try to find D:
1. Hash D -> get index 0
2. Check index 0 -> find A
3. Check index 1 -> find Empty slot
4. Stop searching because Empty slot means end of probe sequence
5. PROBLEM: We never reach D even though it's in the table!

With tombstones, when we delete B:
Index: 0    1    2    3
Items: A -> † -> C -> D  (where † is a tombstone)

Now when finding D:
1. Hash D -> get index 0
2. Check index 0 -> find A
3. Check index 1 -> find tombstone, keep probing
4. Check index 2 -> find C
5. Check index 3 -> find D (success!)

The tombstone tells us "an item was here, keep probing" while Empty means
"nothing was ever here, stop probing." This maintains probe sequence integrity.

Tombstones can be fully removed during resize operations since we rebuild all probe
sequences from scratch at that point.
*/

// EntryState represents the state of a hashmap entry
type EntryState int

const (
	Empty EntryState = iota
	Occupied
	Deleted
)

// Entry represents a key-value pair in the hashmap
type Entry struct {
	Key   string
	Val   any
	State EntryState
}

// HashMap implements a hash table with linear probing
type HashMap struct {
	items         []Entry
	size          int     // Number of occupied entries
	deleted       int     // Number of deleted entries
	loadThreshold float64 // Threshold for resizing
}

// Config holds hashmap configuration options
type Config struct {
	InitialCapacity int
	LoadThreshold   float64
}

// DefaultConfig returns default configuration values
func DefaultConfig() Config {
	return Config{
		InitialCapacity: 16,
		LoadThreshold:   0.75,
	}
}

// NewHashMap creates a new HashMap with default configuration
func NewHashMap() *HashMap {
	return NewHashMapWithConfig(DefaultConfig())
}

// NewHashMapWithConfig creates a new HashMap with custom configuration
func NewHashMapWithConfig(config Config) *HashMap {
	if config.InitialCapacity < 1 {
		config.InitialCapacity = DefaultConfig().InitialCapacity
	}
	if config.LoadThreshold <= 0 || config.LoadThreshold >= 1 {
		config.LoadThreshold = DefaultConfig().LoadThreshold
	}

	return &HashMap{
		items:         make([]Entry, config.InitialCapacity),
		loadThreshold: config.LoadThreshold,
	}
}

// Size returns the number of entries in the map
func (h *HashMap) Size() int {
	return h.size
}

// Capacity returns the current capacity of the map
func (h *HashMap) Capacity() int {
	return len(h.items)
}

// LoadFactor returns the current load factor of the map
func (h *HashMap) LoadFactor() float64 {
	return float64(h.size+h.deleted) / float64(len(h.items))
}

// Insert adds or updates a key-value pair in the map
func (h *HashMap) Insert(key string, value any) error {
	if key == "" {
		return fmt.Errorf("empty key not allowed")
	}

	if h.LoadFactor() >= h.loadThreshold {
		if err := h.resize(len(h.items) * 2); err != nil {
			return fmt.Errorf("resize failed: %v", err)
		}
	}

	hash := h.hash(key)
	idx := hash
	firstDeletedIdx := -1

	// Linear probing with deleted entry tracking
	for i := 0; i < len(h.items); i++ {
		entry := &h.items[idx]

		switch entry.State {
		case Empty:
			// Insert at first deleted slot if found, otherwise at empty slot
			insertIdx := idx
			if firstDeletedIdx != -1 {
				insertIdx = firstDeletedIdx
				h.deleted--
			}

			h.items[insertIdx] = Entry{
				Key:   key,
				Val:   value,
				State: Occupied,
			}
			h.size++
			return nil

		case Deleted:
			// Remember first deleted slot
			if firstDeletedIdx == -1 {
				firstDeletedIdx = idx
			}

		case Occupied:
			if entry.Key == key {
				// Update existing entry
				entry.Val = value
				return nil
			}
		}

		idx = (idx + 1) % len(h.items)
	}

	return fmt.Errorf("no available slots")
}

// Get retrieves a value by key
func (h *HashMap) Get(key string) (any, bool) {
	idx := h.probe(key)
	if idx == -1 {
		return nil, false
	}
	return h.items[idx].Val, true
}

// Delete removes an entry by key
func (h *HashMap) Delete(key string) bool {
	idx := h.probe(key)
	if idx == -1 {
		return false
	}

	h.items[idx].State = Deleted
	h.size--
	h.deleted++

	// Consider shrinking if load factor is very low
	if h.LoadFactor() < h.loadThreshold/4 && len(h.items) > DefaultConfig().InitialCapacity {
		h.resize(len(h.items) / 2)
	}

	return true
}

// Contains checks if a key exists in the map
func (h *HashMap) Contains(key string) bool {
	return h.probe(key) != -1
}

// Clear removes all entries from the map
func (h *HashMap) Clear() {
	h.items = make([]Entry, DefaultConfig().InitialCapacity)
	h.size = 0
	h.deleted = 0
}

// Range iterates over all entries in the map
func (h *HashMap) Range(fn func(key string, value any) bool) {
	for i := range h.items {
		if h.items[i].State == Occupied {
			if !fn(h.items[i].Key, h.items[i].Val) {
				break
			}
		}
	}
}

// probe finds the index of a key using linear probing
func (h *HashMap) probe(key string) int {
	if key == "" || len(h.items) == 0 {
		return -1
	}

	idx := h.hash(key)
	for i := 0; i < len(h.items); i++ {
		entry := &h.items[idx]

		switch entry.State {
		case Empty:
			return -1
		case Occupied:
			if entry.Key == key {
				return idx
			}
		}

		idx = (idx + 1) % len(h.items)
	}

	return -1
}

// resize changes the capacity of the hashmap
func (h *HashMap) resize(newCapacity int) error {
	if newCapacity < h.size {
		return fmt.Errorf("new capacity %d is too small for current size %d", newCapacity, h.size)
	}

	oldItems := h.items
	h.items = make([]Entry, newCapacity)
	h.size = 0
	h.deleted = 0

	// Reinsert all occupied entries
	for _, entry := range oldItems {
		if entry.State == Occupied {
			if err := h.Insert(entry.Key, entry.Val); err != nil {
				return fmt.Errorf("failed to reinsert entry during resize: %v", err)
			}
		}
	}

	return nil
}

// hash generates an index for a key
func (h *HashMap) hash(key string) int {
	// FNV-1a hash
	hash := uint64(14695981039346656037)
	for i := 0; i < len(key); i++ {
		hash ^= uint64(key[i])
		hash *= 1099511628211
	}
	return int(hash % uint64(len(h.items)))
}

// String returns a string representation of the hashmap
func (h *HashMap) String() string {
	var s string
	s += fmt.Sprintf("Size: %d, Capacity: %d, Load Factor: %.2f\n", h.size, len(h.items), h.LoadFactor())
	h.Range(func(key string, value any) bool {
		s += fmt.Sprintf("%s: %v\n", key, value)
		return true
	})
	return s
}
