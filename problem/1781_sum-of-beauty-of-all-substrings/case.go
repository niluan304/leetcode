package sum_of_beauty_of_all_substrings

import (
	"testing"

	"leetcode/tests"
)

var cases = func() []tests.Case[string, int] {
	return []tests.Case[string, int]{
		{Input: "aab", Except: 1},
		{Input: "abaacc", Except: 12},
		{Input: "aabcb", Except: 5},
		{Input: "kcgdnprqxcmpcavjzjgvgekzsvoejxwrdsidzitpzcegxrrrnayndadtqwqswuinzyhdewzzvukqbzobylcporryqpurrzzmidrjcgtfoeyycr" +
			"sbpbinbzweirmlamaajudtaermybbopxknkwalbnowfsevnodehzdalgailizfvnenmfuatxieorjaybxilmjsslalgeecmsbqwdjntfoaizbpm" +
			"xekrtesrguepsevaymfwetnddblkbrirhrxrxvrpnqtazyurmwmlxtxczsypiycedwdgyzelbyapgfmedpzbfjfmbtydaqnshncgciqhatuzzpj" +
			"klomxxqkdvrcwpotadandkwkfnrgiugpxyfjzzwkfdlvytfufxpsdwgmrqzufghuyq", Except: 670620},
	}
}

var funcs = tests.NewFunc[string, int]()

func AddCases(c func() []tests.Case[string, int]) {
	_cases := cases()
	cases = func() []tests.Case[string, int] {
		return append(_cases, c()...)
	}
}

func AddFunc(f ...func(string) int) {
	funcs = append(funcs, tests.NewFunc(f...)...)
}

func Unit(t *testing.T) {
	tests.Unit(t, tests.Test[string, int]{
		Solution: funcs,
		Cases:    cases,
		IsEqual:  nil,
	})
}

func Bench(b *testing.B) {
	tests.Bench(b, tests.Test[string, int]{
		Solution: funcs,
		Cases:    cases,
		IsEqual:  nil,
	})
}
