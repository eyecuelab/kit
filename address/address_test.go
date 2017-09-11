package address

import (
	"reflect"
	"testing"

	"googlemaps.github.io/maps"
	"gopkg.in/mgo.v2/bson"
)

func TestFromFactualRecord(t *testing.T) {
	type args struct {
		factualRecord bson.M
	}
	tests := []struct {
		name string
		args args
		want Address
	}{
		{"hell", args{hellFactualRecord}, hellAddress},
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

func TestAddress_SharedComponentsOf(t *testing.T) {
	type args struct {
		b Address
	}
	tests := []struct {
		name     string
		reciever Address
		arg      Address
		want     Address
	}{
		{"copy", heavenAddress, hellAddress, hellWithHeavenFieldsOnly},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			argCopy := tt.arg
			a := tt.reciever
			if got := a.SharedComponentsOf(tt.arg); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Address.SharedComponentsOf() = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(tt.arg, argCopy) {
				t.Errorf("should not modify %v", tt.arg)
			}
		})
	}
}

var hellFactualRecord = bson.M{
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

var heavenAddress = Address{
	Street:    "1 N Heaven St",
	Extension: "#18",
	Locality:  "San Diego",
	Region:    "CA",
	Country:   "us",
}

var hellWithHeavenFieldsOnly = Address{
	Street:    "1 S Hell St",
	Extension: "#666",
	Locality:  "Newark",
	Region:    "NJ",
	Country:   "us",
}

func TestFromGoogleAddressComponents(t *testing.T) {
	type args struct {
		components []maps.AddressComponent
	}
	tests := []struct {
		name        string
		args        args
		wantAddress Address
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAddress := FromGoogleAddressComponents(tt.args.components); !reflect.DeepEqual(gotAddress, tt.wantAddress) {
				t.Errorf("FromGoogleAddressComponents() = %v, want %v", gotAddress, tt.wantAddress)
			}
		})
	}
}
