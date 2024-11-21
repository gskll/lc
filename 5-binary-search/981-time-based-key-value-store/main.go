package main

import "fmt"

// https://leetcode.com/problems/time-based-key-value-store/description/

// Design a time-based key-value data structure that can store multiple values for the same key at different time stamps and retrieve the key's value at a certain timestamp.
//
// Implement the TimeMap class:
//
// TimeMap() Initializes the object of the data structure.
// void set(String key, String value, int timestamp) Stores the key key with the value value at the given time timestamp.
// String get(String key, int timestamp) Returns a value such that set was called previously, with timestamp_prev <= timestamp. If there are multiple such values, it returns the value associated with the largest timestamp_prev. If there are no values, it returns "".
//
// 1 <= key.length, value.length <= 100
// key and value consist of lowercase English letters and digits.
// 1 <= timestamp <= 107
// All the timestamps timestamp of set are strictly increasing.
// At most 2 * 105 calls will be made to set and get.

type entry struct {
	value string
	ts    int
}

type TimeMap struct {
	items map[string][]entry
}

func Constructor() TimeMap {
	return TimeMap{items: make(map[string][]entry)}
}

func (this *TimeMap) Set(key string, value string, timestamp int) {
	this.items[key] = append(this.items[key], entry{value: value, ts: timestamp})
}

func (this *TimeMap) Get(key string, timestamp int) string {
	values := this.items[key]
	if len(values) == 0 || timestamp < values[0].ts {
		return ""
	}
	if timestamp >= values[len(values)-1].ts {
		return values[len(values)-1].value
	}

	l, r := 0, len(values)-1
	for l < r {
		i := l + (r-l+1)/2

		if values[i].ts <= timestamp {
			l = i
		} else {
			r = i - 1
		}
	}
	return values[l].value
}

/**
 * Your TimeMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Set(key,value,timestamp);
 * param_2 := obj.Get(key,timestamp);
 */

func testTimeMap() {
	// Create TimeMap instance
	timeMap := Constructor()

	// Test set/get operations
	tests := []struct {
		operation string
		key       string
		value     string
		timestamp int
		expected  string
	}{
		// First test case
		{"set", "foo", "bar", 1, ""},
		{"get", "foo", "", 1, "bar"},
		{"get", "foo", "", 3, "bar"},
		{"set", "foo", "bar2", 4, ""},
		{"get", "foo", "", 4, "bar2"},
		{"get", "foo", "", 5, "bar2"},

		// Second test case
		{"set", "love", "high", 10, ""},
		{"set", "love", "low", 20, ""},
		{"get", "love", "", 5, ""},
		{"get", "love", "", 10, "high"},
		{"get", "love", "", 15, "high"},
		{"get", "love", "", 20, "low"},
		{"get", "love", "", 25, "low"},
	}

	fmt.Println("Testing TimeMap operations:")
	for i, test := range tests {
		if test.operation == "set" {
			timeMap.Set(test.key, test.value, test.timestamp)
			fmt.Printf("Operation %d: Set(%s, %s, %d)\n",
				i, test.key, test.value, test.timestamp)
		} else {
			result := timeMap.Get(test.key, test.timestamp)
			fmt.Printf("Operation %d: Get(%s, %d) = %s",
				i, test.key, test.timestamp, result)
			if result == test.expected {
				fmt.Println(" ✓")
			} else {
				fmt.Printf(" ✗ (expected %s)\n", test.expected)
			}
		}
	}
}

func main() {
	testTimeMap()
}
