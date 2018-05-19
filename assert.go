package xunit

func Assert(res bool) {
	if !res {
		panic(res)
	}
}
