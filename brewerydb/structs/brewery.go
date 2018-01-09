package structs

type Brewery struct {
	ID             string `json:"id,omitempty"`
	Name           string `json:"name,omitempty"`
	Description    string `json:"description,omitempty"`
	Website        string `json:"website,omitempty"`
	Established    string `json:"established,omitempty"`
	MailingListURL string `json:"mailingListUrl,omitempty"`
	IsOrganic      bool   `json:"isOrganic,omitempty"`
	Images         Images `json:"images,omitempty"`
}

func booleanVal(b []byte) (val bool, ok bool) {
	switch string(b) {
	case "true", "Y":
		return true, true
	case "false", "N", "":
		return false, true
	default:
		return false, false
	}
}
