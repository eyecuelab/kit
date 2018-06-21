package nullconv

import (
	"gopkg.in/volatiletech/null.v6"
	"time"
)

func NewTime(t time.Time) null.Time {
	return null.Time{
		Time: t,
		Valid: !t.IsZero(),
	}
}

func NewString(s string) null.String {
	return null.String{
		String: s,
		Valid: s != "",
	}
}

func NewInt(i int) null.Int {
	return null.Int{
		Int: i,
		Valid: i != 0,
	}
}
