package main

import "testing"

func TestShouldGenerateValidMessagesSimple(t *testing.T) {
	rules := map[int]Rule{
		0: {
			M: [][]int{
				{1, 2},
			},
		},
		1: {
			R: 'a',
		},
		2: {
			M: [][]int{
				{1, 3},
				{3, 1},
			},
		},
		3: {
			R: 'b',
		},
	}
	expected := []string{
		"aab",
		"aba",
	}

	resetCache()
	result := generateValidMessages(rules, 0, true)
	if len(result) != len(expected) {
		t.Errorf("got %+v, expected %+v", result, expected)
		return
	}
	for i := range result {
		if result[i] != expected[i] {
			t.Errorf("got %+v, expected %+v", result, expected)
		}
	}
}

func TestShouldGenerateValidMessagesComplex(t *testing.T) {
	rules := map[int]Rule{
		0: {
			M: [][]int{
				{4, 1, 5},
			},
		},
		1: {
			M: [][]int{
				{2, 3},
				{3, 2},
			},
		},
		2: {
			M: [][]int{
				{4, 4},
				{5, 5},
			},
		},
		3: {
			M: [][]int{
				{4, 5},
				{5, 4},
			},
		},
		4: {
			R: 'a',
		},
		5: {
			R: 'b',
		},
	}
	expected := []string{
		"aaaabb",
		"aaabab",
		"abbabb",
		"abbbab",
		"aabaab",
		"aabbbb",
		"abaaab",
		"ababbb",
	}

	resetCache()
	result := generateValidMessages(rules, 0, true)
	if len(result) != len(expected) {
		t.Errorf("got %+v, expected %+v", result, expected)
		return
	}
	for _, i := range expected {
		var found bool
		for _, j := range result {
			if i == j {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("got %+v, expected %+v", result, expected)
		}
	}
}

func TestShouldGenerateValidMessagesLoop(t *testing.T) {
	rules := map[int]Rule{
		0: {
			M: [][]int{
				{1, 2},
			},
		},
		1: {
			M: [][]int{
				{3},
				{3, 1},
			},
		},
		2: {
			M: [][]int{
				{3, 4},
				{3, 1},
			},
		},
		3: {
			R: 'a',
		},
		4: {
			R: 'b',
		},
	}
	expected := []string{
		"aaab",
	}

	resetCache()
	result := generateValidMessages(rules, 0, true)
	if len(result) != len(expected) {
		t.Errorf("got %+v, expected %+v", result, expected)
		return
	}
	for i := range result {
		if result[i] != expected[i] {
			t.Errorf("got %+v, expected %+v", result, expected)
		}
	}
}

func TestTask1ShouldFindResult(t *testing.T) {
	rules := map[int]Rule{
		0: {
			M: [][]int{
				{4, 1, 5},
			},
		},
		1: {
			M: [][]int{
				{2, 3},
				{3, 2},
			},
		},
		2: {
			M: [][]int{
				{4, 4},
				{5, 5},
			},
		},
		3: {
			M: [][]int{
				{4, 5},
				{5, 4},
			},
		},
		4: {
			R: 'a',
		},
		5: {
			R: 'b',
		},
	}
	messages := []string{
		"ababbb",
		"bababa",
		"abbbab",
		"aaabbb",
		"aaaabbb",
	}
	expected := 2

	resetCache()
	result := Task1(rules, messages, true)
	if result != expected {
		t.Errorf("got %d, expected %d", result, expected)
	}
}

func TestTask2ShouldFindResult(t *testing.T) {
	rules := map[int]Rule{
		42: {
			M: [][]int{{
				9, 14},
				{10, 1},
			},
		},
		9: {
			M: [][]int{
				{14, 27},
				{1, 26},
			},
		},
		10: {
			M: [][]int{
				{23, 14},
				{28, 1},
			},
		},
		1: {
			R: 'a',
		},
		11: {
			M: [][]int{
				{42, 31},
			},
		},
		5: {
			M: [][]int{
				{1, 14},
				{15, 1},
			},
		},
		19: {
			M: [][]int{
				{14, 1},
				{14, 14},
			},
		},
		12: {
			M: [][]int{
				{24, 14},
				{19, 1},
			},
		},
		16: {
			M: [][]int{
				{15, 1},
				{14, 14},
			},
		},
		31: {
			M: [][]int{
				{14, 17},
				{1, 13},
			},
		},
		6: {
			M: [][]int{
				{14, 14},
				{1, 14},
			},
		},
		2: {
			M: [][]int{
				{1, 24},
				{14, 4},
			},
		},
		0: {
			M: [][]int{
				{8, 11},
			},
		},
		13: {
			M: [][]int{
				{14, 3},
				{1, 12},
			},
		},
		15: {
			M: [][]int{
				{1},
				{14},
			},
		},
		17: {
			M: [][]int{
				{14, 2},
				{1, 7},
			},
		},
		23: {
			M: [][]int{
				{25, 1},
				{22, 14},
			},
		},
		28: {
			M: [][]int{
				{16, 1},
			},
		},
		4: {
			M: [][]int{
				{1, 1},
			},
		},
		20: {
			M: [][]int{
				{14, 14},
				{1, 15},
			},
		},
		3: {
			M: [][]int{
				{5, 14},
				{16, 1},
			},
		},
		27: {
			M: [][]int{
				{1, 6},
				{14, 18},
			},
		},
		14: {
			R: 'b',
		},
		21: {
			M: [][]int{
				{14, 1},
				{1, 14},
			},
		},
		25: {
			M: [][]int{
				{1, 1},
				{1, 14},
			},
		},
		22: {
			M: [][]int{
				{14, 14},
			},
		},
		8: {
			M: [][]int{
				{42},
			},
		},
		26: {
			M: [][]int{
				{14, 22},
				{1, 20},
			},
		},
		18: {
			M: [][]int{
				{15, 15},
			},
		},
		7: {
			M: [][]int{
				{14, 5},
				{1, 21},
			},
		},
		24: {
			M: [][]int{
				{14, 1},
			},
		},
	}
	messages := []string{
		"abbbbbabbbaaaababbaabbbbabababbbabbbbbbabaaaa",
		"bbabbbbaabaabba",
		"babbbbaabbbbbabbbbbbaabaaabaaa",
		"aaabbbbbbaaaabaababaabababbabaaabbababababaaa",
		"bbbbbbbaaaabbbbaaabbabaaa",
		"bbbababbbbaaaaaaaabbababaaababaabab",
		"ababaaaaaabaaab",
		"ababaaaaabbbaba",
		"baabbaaaabbaaaababbaababb",
		"abbbbabbbbaaaababbbbbbaaaababb",
		"aaaaabbaabaaaaababaa",
		"aaaabbaaaabbaaa",
		"aaaabbaabbaaaaaaabbbabbbaaabbaabaaa",
		"babaaabbbaaabaababbaabababaaab",
		"aabbbbbaabbbaaaaaabbbbbababaaaaabbaaabba",
	}
	expected := 12

	resetCache()
	result := Task2(rules, messages, false)
	if result != expected {
		t.Errorf("got %d, expected %d", result, expected)
	}
}
