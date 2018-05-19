package xunit

func Assert(res bool) {
	if !res {
		panic(res)
	}
}

func AssertNot(res bool) {
	if res {
		panic(res)
	}
}
