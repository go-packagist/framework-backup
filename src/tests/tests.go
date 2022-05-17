package tests

type Tests struct {
	test string
}

func New(test string) *Tests {
	return &Tests{test}
}
