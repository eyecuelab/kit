package address

import (
	"reflect"
	"testing"

	"gopkg.in/mgo.v2/bson"
)

var hellRecord = bson.M{
	"address":          "1 S Hell St",
	"existence":        "0.7",
	"admin_region":     nil,
	"postcode":         "66666",
	"email":            "fakeplace@fake.com",
	"hours_display":    nil,
	"post_town":        nil,
	"name":             "TEST DATABSE",
	"address_extended": "#666",
	"created_at":       `ISODate("2017-09-06T19:18:19.185Z")`,
	"chain_id":         nil,
	"po_box":           nil,
	"fax":              nil,
	"source":           "factual",
	"matched":          false,
	"hours":            nil,
	"locality":         "Newark",
	"country":          "us",
	"region":           "NJ",
	"chain_name":       nil,
	"tel":              "(312) 427-4410",
	"latitude":         666,
	"longitude":        666,
}

var hellAddress = Address{
	Street:     "1 S Hell St",
	Extension:  "#666",
	POBox:      "",
	Locality:   "Newark",
	Region:     "NJ",
	PostalCode: "66666",
	Country:    "us",
}

func TestFromFactualRecord(t *testing.T) {
	type args struct {
		factualRecord bson.M
	}
	tests := []struct {
		name string
		args args
		want Address
	}{
		{"hell", args{hellRecord}, hellAddress},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FromFactualRecord(tt.args.factualRecord); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FromFactualRecord() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAddress_String(t *testing.T) {
	tests := []struct {
		name string
		arg  Address
		want string
	}{
		{"hell", hellAddress, "1 S Hell St, #666, Newark, NJ, 66666, us"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.arg.String(); got != tt.want {
				t.Errorf("Address.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
