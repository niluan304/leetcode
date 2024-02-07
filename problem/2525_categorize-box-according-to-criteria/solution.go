package main

// Deprecated: Although it can run, its maintainability is poor.
func categorizeBox(length int, width int, height int, mass int) string {
	volume := length * width * height
	isBulky := volume >= 1e9 || length >= 1e4 || width >= 1e4 || height >= 1e4
	isHeavy := mass >= 100

	// 虽然可以运行，但是可维护性差
	switch {
	case isBulky && isHeavy:
		return "Both"
	case isBulky:
		return "Bulky"
	case isHeavy:
		return "Heavy"
	default:
		return "Neither"
	}
}

func categorizeBox2(length int, width int, height int, mass int) string {
	type Box struct{ isBulky, isHeavy bool }

	volume := length * width * height
	var box = Box{
		isBulky: volume >= 1e9 || length >= 1e4 || width >= 1e4 || height >= 1e4,
		isHeavy: mass >= 100,
	}
	switch box {
	case Box{isBulky: true, isHeavy: true}:
		return "Both"
	case Box{isBulky: true, isHeavy: false}:
		return "Bulky"
	case Box{isBulky: false, isHeavy: true}:
		return "Heavy"
	default:
		return "Neither"
	}
}

func categorizeBox3(length int, width int, height int, mass int) string {
	volume := length * width * height
	isBulky := volume >= 1e9 || length >= 1e4 || width >= 1e4 || height >= 1e4
	isHeavy := mass >= 100

	if isBulky {
		if isHeavy {
			return "Both"
		} else {
			return "Bulky"
		}
	}

	if isHeavy {
		return "Heavy"
	} else {
		return "Neither"
	}
}
