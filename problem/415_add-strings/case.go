package add_strings

import (
	"math"
	"strconv"
	"testing"

	"leetcode/tests"
)

type Input struct{ Num1, Num2 string }

var cases = func() []tests.Case[Input, string] {
	return []tests.Case[Input, string]{
		{Input: Input{Num1: "11", Num2: "123"}, Except: "134"},
		{Input: Input{Num1: "456", Num2: "77"}, Except: "533"},
		{Input: Input{Num1: "0", Num2: "0"}, Except: "0"},
		{Input: Input{
			Num1: strconv.FormatUint(math.MaxUint64, 10),
			Num2: strconv.FormatUint(math.MaxUint64, 10)},
			Except: "36893488147419103230"},
		{Input: Input{
			Num1: "987987978978242342342328345645645645745745790483290480423479879",
			Num2: "000000000879789789789789745645645127884732749209487129074590812"},
			Except: "987987979858032132132118091291290773630478539692777609498070691"},
		{Input: Input{
			Num1: "9999999999999999549291066784979473595300225087383" +
				"524118479625982517885450291174622154390152298057300" +
				"868772377386949310916067328",
			Num2: "19999999999999999098582133569958947190600450174767" +
				"0482369592519650357709005823492443087803045961146017" +
				"375447547738986218321347068000000"},
			Except: "2000000099999999909858208849906562568854780970478" +
				"95569753116638129983691523708942734262425200351298315" +
				"432748416511363605270657984067328"},
	}
}

type Func func(string, string) string

func adaptor(f Func) func(in Input) (out string) {
	return func(in Input) (out string) {
		return f(in.Num1, in.Num2)
	}
}

var funcs = tests.NewFuncWithAdaptor(adaptor)

func AddCases(c func() []tests.Case[Input, string]) {
	_cases := cases()
	cases = func() []tests.Case[Input, string] {
		return append(_cases, c()...)
	}
}

func AddFunc(f ...Func) {
	funcs = append(funcs, tests.NewFuncWithAdaptor(adaptor, f...)...)
}

func Unit(t *testing.T) {
	tests.Unit(t, cases, funcs...)
}

var checkResult = true

func Bench(b *testing.B) {
	checkResult = false
	tests.Bench(b, cases, funcs...)
}
