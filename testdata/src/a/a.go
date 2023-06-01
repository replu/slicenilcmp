package a

func f() {
	var basicSlice []int
	if basicSlice == nil { // want "suggestion: use len func for empty check"
		print(basicSlice)
	}

	if len(basicSlice) == 0 { // ok
		print(basicSlice)
	}

	a := basicSlice == nil // want "suggestion: use len func for empty check"
	print(a)

	if err := (basicSlice == nil); err { // want "suggestion: use len func for empty check"
		print(basicSlice)
	}

	var structSlice []struct{}
	if structSlice == nil { // want "suggestion: use len func for empty check"
		print(structSlice)
	}

	if len(structSlice) == 0 { // ok
		print(structSlice)
	}
}
