package senml

//Implementation is based on: https://tools.ietf.org/html/draft-ietf-core-senml-13

import (
	"encoding/json"
	"errors"
	"fmt"
)

type Pack struct {
	Resolve bool //Not used yet. This will indicate if we want to resolve the entries when marshalling/unmarshalling
	Entries []Entry
}

func (s *Pack) UnmarshalJSON(data []byte) error {
	entries := make([]json.RawMessage, 0)
	if err := json.Unmarshal(data, &entries); err != nil {
		return err
	}
	for _, v := range entries {
		fields := make(map[string]interface{})
		if err := json.Unmarshal(v, &fields); err != nil {
			return err
		}
		var en Entry
		if _, ok := fields["v"]; ok {
			enn := Float64Entry{}
			if err := json.Unmarshal(v, &enn); err != nil {
				return err
			}
			en = enn
		} else if _, ok := fields["vs"]; ok {
			enn := StringEntry{}
			if err := json.Unmarshal(v, &enn); err != nil {
				return err
			}
			en = enn
		} else if _, ok := fields["vb"]; ok {
			enn := BooleanEntry{}
			if err := json.Unmarshal(v, &enn); err != nil {
				return err
			}
			en = enn
		} else if _, ok := fields["vd"]; ok {
			enn := DataEntry{}
			if err := json.Unmarshal(v, &enn); err != nil {
				return err
			}
			en = enn
		}
		s.Entries = append(s.Entries, en)
	}
	return nil
}

func (s *Pack) MarshalJSON() ([]byte, error) {
	return nil, errors.New("JSON marshalling is not implemented yet")
}

type Entry interface {
	ToString() string
}

//Can we somehow extract the fields and share it across the different types of entries?

type Float64Entry struct {
	BaseName    string  `json:"bn,omitempty"`
	BaseTime    int64   `json:"bt,omitempty"`
	BaseUnit    string  `json:"bu,omitempty"`
	BaseValue   string  `json:"bv,omitempty"`
	BaseSum     string  `json:"bs,omitempty"`
	BaseVersion string  `json:"bver,omitempty"`
	Name        string  `json:"n"`
	Unit        string  `json:"u"`
	Sum         string  `json:"s"`
	Time        int64   `json:"t"`
	UpdateTime  int64   `json:"ut"`
	Value       float64 `json:"v"`
}

func (en Float64Entry) ToString() string {
	//Return comma seperated struct
	return fmt.Sprintf("%#v", en)
}

type StringEntry struct {
	BaseName    string `json:"bn,omitempty"`
	BaseTime    int64  `json:"bt,omitempty"`
	BaseUnit    string `json:"bu,omitempty"`
	BaseValue   string `json:"bv,omitempty"`
	BaseSum     string `json:"bs,omitempty"`
	BaseVersion string `json:"bver,omitempty"`
	Name        string `json:"n"`
	Unit        string `json:"u"`
	Sum         string `json:"s"`
	Time        int64  `json:"t"`
	UpdateTime  int64  `json:"ut"`
	StringValue string `json:"vs"`
}

func (en StringEntry) ToString() string {
	//Return comma seperated struct
	return fmt.Sprintf("%#v", en)
}

type BooleanEntry struct {
	BaseName     string `json:"bn,omitempty"`
	BaseTime     int64  `json:"bt,omitempty"`
	BaseUnit     string `json:"bu,omitempty"`
	BaseValue    string `json:"bv,omitempty"`
	BaseSum      string `json:"bs,omitempty"`
	BaseVersion  string `json:"bver,omitempty"`
	Name         string `json:"n"`
	Unit         string `json:"u"`
	Sum          string `json:"s"`
	Time         int64  `json:"t"`
	UpdateTime   int64  `json:"ut"`
	BooleanValue bool   `json:"vb"`
}

func (en BooleanEntry) ToString() string {
	//Return comma seperated struct
	return fmt.Sprintf("%#v", en)
}

type DataEntry struct {
	BaseName    string `json:"bn,omitempty"`
	BaseTime    int64  `json:"bt,omitempty"`
	BaseUnit    string `json:"bu,omitempty"`
	BaseValue   string `json:"bv,omitempty"`
	BaseSum     string `json:"bs,omitempty"`
	BaseVersion string `json:"bver,omitempty"`
	Name        string `json:"n"`
	Unit        string `json:"u"`
	Sum         string `json:"s"`
	Time        int64  `json:"t"`
	UpdateTime  int64  `json:"ut"`
	DataValue   string `json:"vd"`
}

func (en DataEntry) ToString() string {
	//Return comma seperated struct
	return fmt.Sprintf("%#v", en)
}
