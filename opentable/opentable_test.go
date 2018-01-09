package opentable

import (
	"testing"
)

func TestBadRequestError_Error(t *testing.T) {
	tests := []struct {
		name string
		err  BadRequestError
		want string
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.err.Error(); got != tt.want {
				t.Errorf("BadRequestError.Error() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_accessToken(t *testing.T) {
	tests := []struct {
		name  string
		want  string
		want1 bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := accessToken()
			if got != tt.want {
				t.Errorf("accessToken() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("accessToken() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_priceOK(t *testing.T) {
	type args struct {
		price int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := priceOK(tt.args.price); got != tt.want {
				t.Errorf("priceOK() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_perPageOK(t *testing.T) {
	type args struct {
		perPage int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := perPageOK(tt.args.perPage); got != tt.want {
				t.Errorf("perPageOK() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_pageOK(t *testing.T) {
	type args struct {
		page int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pageOK(tt.args.page); got != tt.want {
				t.Errorf("pageOK() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_request_validate(t *testing.T) {
	type fields struct {
		Price   int
		Name    string
		Street  string
		State   string
		City    string
		Zip     string
		Country string
		Page    int
		PerPage int
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &request{
				Price:   tt.fields.Price,
				Name:    tt.fields.Name,
				Street:  tt.fields.Street,
				State:   tt.fields.State,
				City:    tt.fields.City,
				Zip:     tt.fields.Zip,
				Country: tt.fields.Country,
				Page:    tt.fields.Page,
				PerPage: tt.fields.PerPage,
			}
			if err := r.validate(); (err != nil) != tt.wantErr {
				t.Errorf("request.validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

type byKey []param

func (bk byKey) Len() int           { return len(bk) }
func (bk byKey) Less(i, j int) bool { return bk[i].key < bk[j].key }
func (bk byKey) Swap(i, j int)      { bk[i], bk[j] = bk[j], bk[i] }

func Test_request_Get(t *testing.T) {

}
