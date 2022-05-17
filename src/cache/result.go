package cache

type Result struct {
	Val interface{}
	Err error
}

func (r *Result) Value() interface{} {
	return r.Val
}

func (r *Result) Error() error {
	return r.Err
}

func (r *Result) IsError() bool {
	return r.Err != nil
}

func (r *Result) ValOrNil() interface{} {
	if r.Err != nil {
		return nil
	}

	return r.Val
}
