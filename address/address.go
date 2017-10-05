package address

import (
	"strings"

	"github.com/eyecuelab/kit/stringslice"
)

//Address represents a physical location
type Address struct {
	Street, Extension, POBox string
	Locality                 string //city
	Region                   string //state
	PostalCode               string //zip
	Country                  string
}

//filterOutComponentsMissingFromReciever returns a copy of b, except that where a.component == "", b.component == ""
func (a *Address) filterOutComponentsMissingFromReciever(b Address) Address {
	if a.Street == "" {
		b.Street = ""
	}
	if a.Extension == "" {
		b.Extension = ""
	}
	if a.POBox == "" {
		b.POBox = ""
	}
	if a.Locality == "" {
		b.Locality = ""
	}
	if a.Region == "" {
		b.Region = ""
	}
	if a.PostalCode == "" {
		b.PostalCode = ""
	}
	if a.Country == "" {
		b.Country = ""
	}
	return b
}

//SharedComponents returns copies of a and b, except if a.component == "" || b.component == "",
//c.component == "" && d.component == ""
func SharedComponents(a, b Address) (Address, Address) {
	b = a.filterOutComponentsMissingFromReciever(b)
	a = b.filterOutComponentsMissingFromReciever(a)
	return a, b
}

//StringSliceFromNonempty returns a stringslice of the non-empty components of a
func (a *Address) StringSliceFromNonempty() []string {
	return stringslice.AppendIfNonEmpty(make([]string, 0, 7),
		a.Street, a.Extension, a.POBox, a.Locality, a.Region, a.PostalCode, a.Country)
}

func (a *Address) StringSlice() []string {
	return []string{}
}

//Returns true if all components of the address are ""
func (a *Address) Empty() bool {
	isEmpty := func(s string) bool {
		return s == ""
	}
	return stringslice.All(a.StringSlice(), isEmpty)

}
func (a *Address) String() string {
	//PO BOX deliberately ignored for now; 9/7/2017
	components := []string{
		a.Street, a.Extension, a.POBox,
		a.Locality, a.Region, a.PostalCode, a.Country,
	}
	return strings.Join(stringslice.NonEmpty(components), ", ")
}

//NonEmptyComponents is the number of nonempty components of a; that is, an address with only a steet has 1,
//street and locality is 2, ...
func (a *Address) NonEmptyComponents() int {
	return sumBools(
		a.Street != "",
		a.Extension != "",
		a.POBox != "",
		a.Locality != "",
		a.Region != "",
		a.PostalCode != "",
		a.Country != "")
}

func sumBools(bools ...bool) int {
	sum := 0
	for _, b := range bools {
		if b {
			sum++
		}
	}
	return sum
}
