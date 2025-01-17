package main

// https://leetcode.com/problems/group-anagrams/description/

// Given an array of strings strs, group the anagrams together. You can return the answer in any order.
//
// strs[i] consists of lowercase English letters.
// 0 <= strs[i].length <= 100
// 1 <= strs.length <= 10^4

func strToKey(str string) [26]byte {
	var key [26]byte
	for _, r := range str {
		key[r-'a']++
	}
	return key
}

func groupAnagrams(strs []string) [][]string {
	groups := make(map[[26]byte][]string)

	for _, str := range strs {
		key := strToKey(str)
		groups[key] = append(groups[key], str)
	}

	result := make([][]string, 0, len(groups))
	for _, group := range groups {
		result = append(result, group)
	}
	return result
}
