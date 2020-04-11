package main

import (
	"fmt"
	"sort"
)

type Entry struct {
	value     string
	timestamp int
}

type TimeMap struct {
	key2entries map[string][]Entry
}

/** Initialize your data structure here. */
func Constructor() TimeMap {
	return TimeMap{
		key2entries: make(map[string][]Entry),
	}
}

func (this *TimeMap) Set(key string, value string, timestamp int) {
	entries, ok := this.key2entries[key]

	entry := Entry{
		value:     value,
		timestamp: timestamp,
	}
	entry.value = value
	entry.timestamp = timestamp

	if ok {
		entries = append(entries, entry)
		this.key2entries[key] = entries
		return
	}

	this.key2entries[key] = []Entry{entry}
}

func (this *TimeMap) Get(key string, timestamp int) string {
	if entries, ok := this.key2entries[key]; ok {
		idx := sort.Search(len(entries), func(i int) bool {
			return timestamp < entries[i].timestamp
		})
		if len(entries) == 0 || idx == 0 {
			return ""
		}
		if idx == -1 {
			return entries[len(entries)-1].value
		}
		return entries[idx-1].value
	}

	return ""
}

func main() {
	kv := Constructor()
	kv.Set("foo", "bar", 1)       // store the key "foo" and value "bar" along with timestamp = 1
	fmt.Println(kv.Get("foo", 1)) // output "bar"
	fmt.Println(kv.Get("foo", 3)) // output "bar" since there is no value corresponding to foo at timestamp 3 and timestamp 2, then the only value is at timestamp 1 ie "bar"
	kv.Set("foo", "bar2", 4)
	fmt.Println(kv.Get("foo", 4)) // output "bar2"
	fmt.Println(kv.Get("foo", 5)) //output "bar2"
}
