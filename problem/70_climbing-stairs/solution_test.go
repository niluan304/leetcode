package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_climbStairs(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []func(n int) int{
		//bruteForce,
		climbStairs,
		climbStairs2,
		//leetcode,
		//endlessCheng,
	}

	for _, f := range fs {
		err := tests.RunFunc(t, f, samples, targetCaseNum)
		if err != nil {
			t.Error(err)
		}
	}
}

var samples = `
2
2

3
3


` +
	`
4
5

5
8

6
13

7
21

8
34

9
55

10
89

11
144

12
233

13
377

14
610

15
987

16
1597

17
2584

18
4181

19
6765

20
10946

21
17711

22
28657

23
46368

24
75025

25
121393

26
196418

27
317811

28
514229

29
832040

30
1346269

31
2178309

32
3524578

33
5702887

34
9227465

35
14930352

36
24157817

37
39088169

38
63245986

39
102334155

40
165580141

41
267914296

42
433494437

43
701408733

44
1134903170

45
1836311903

`
