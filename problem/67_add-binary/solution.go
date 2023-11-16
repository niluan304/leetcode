package main

func addBinary(a string, b string) string {
	var sum = []byte{'0'}

	n1, n2 := len(a), len(b)
	for i := 0; i < n1 || i < n2; i++ {
		num := sum[len(sum)-1]
		if i < n1 {
			num += (a[n1-i-1]) - '0'
		}
		if i < n2 {
			num += (b[n2-i-1]) - '0'
		}

		var flag byte = '0'
		if num > 1+'0' {
			num -= 2
			flag += 1
		}
		sum[len(sum)-1] = num
		sum = append(sum, flag)
	}

	if sum[len(sum)-1] == '0' {
		sum = sum[:len(sum)-1]
	}
	n := len(sum)
	for i := 0; i < n/2; i++ {
		sum[i], sum[n-i-1] = sum[n-i-1], sum[i]
	}
	return string(sum)
}

func addBinary2(a string, b string) string {
	if len(a) < len(b) {
		a, b = b, a
	}

	var sum = []byte("0" + a)

	for i, j := len(sum)-1, len(b)-1; i > 0; i, j = i-1, j-1 {
		if j >= 0 {
			sum[i] += b[j] - '0'
		}

		if sum[i] >= '2' {
			sum[i] -= 2
			sum[i-1] += 1
		}
	}

	if sum[0] == '0' {
		return string(sum[1:])
	}
	return string(sum)
}
