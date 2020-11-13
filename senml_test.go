package senml

import (
	"encoding/json"
	"reflect"
	"testing"
)

type result struct {
	BaseName string
	Name     string
	Unit     string
	Value    interface{}
}

//Taken from https://tools.ietf.org/html/rfc8428#section-5
const singleDataPoint = `
	[
		{"n":"urn:dev:ow:10e2073a01080063","u":"Cel","v":23.1}
	]
`

func TestSingleDataPoint(t *testing.T) {
	var pack Pack
	err := json.Unmarshal([]byte(singleDataPoint), &pack)
	if err != nil {
		t.Fatalf("error unmarshalling single data point: %s", err)
	}
	if len(pack.Entries) != 1 {
		t.Fatalf("Expected entries to be of length 1 but got %d", len(pack.Entries))
	}

	float64Entry, ok := pack.Entries[0].(Float64Entry)
	if !ok {
		t.Fatalf("Invalid entry type %s", reflect.ValueOf(&pack.Entries[0]).Elem().Type())
	}

	if float64Entry.Name != "urn:dev:ow:10e2073a01080063" {
		t.Fatalf("Invalid name expected %s got %s", "urn:dev:ow:10e2073a01080063", float64Entry.Name)
	}

	if float64Entry.Unit != "Cel" {
		t.Fatalf("Invalid unit expected %s got %s", "Cel", float64Entry.Unit)
	}

	if float64Entry.Value != 23.1 {
		t.Fatalf("Invalid value expected %f got %f", 23.1, float64Entry.Value)
	}
}

const multipleDataPoints = `
[
  {"bn":"urn:dev:ow:10e2073a01080063:","n":"voltage","u":"V","v":120.1},
  {"n":"current","u":"A","v":1.2}
]
`

func TestMultipleDataPoints(t *testing.T) {
	var pack Pack

	err := json.Unmarshal([]byte(multipleDataPoints), &pack)
	if err != nil {
		t.Error("error unmarshalling multiple data points: ", err)
	}
	if len(pack.Entries) != 2 {
		t.Errorf("Expected entries to be of length 2 but got %d", len(pack.Entries))
	}
	//TODO: implement testing for multiple
}

const multipleDataPointsRelativeTime = `
	[
	{"bn":"urn:dev:ow:10e2073a0108006:","bt":1.276020076001e+09,
	 "bu":"A","bver":5,
	 "n":"voltage","u":"V","v":120.1},
	{"n":"current","t":-5,"v":1.2},
	{"n":"current","t":-4,"v":1.3},
	{"n":"current","t":-3,"v":1.4},
	{"n":"current","t":-2,"v":1.5},
	{"n":"current","t":-1,"v":1.6},
	{"n":"current","v":1.7}
  ]
`

const multipleMeasurements = `
	[
	{"bn":"urn:dev:ow:10e2073a01080063","bt":1.320067464e+09,
	 "bu":"%RH","v":20},
	{"u":"lon","v":24.30621},
	{"u":"lat","v":60.07965},
	{"t":60,"v":20.3},
	{"u":"lon","t":60,"v":24.30622},
	{"u":"lat","t":60,"v":60.07965},
	{"t":120,"v":20.7},
	{"u":"lon","t":120,"v":24.30623},
	{"u":"lat","t":120,"v":60.07966},
	{"u":"%EL","t":150,"v":98},
	{"t":180,"v":21.2},
	{"u":"lon","t":180,"v":24.30628},
	{"u":"lat","t":180,"v":60.07967}
  ]
`

const resolved = `
[
	{"n":"urn:dev:ow:10e2073a01080063","u":"%RH","t":1.320067464e+09,
	 "v":20},
	{"n":"urn:dev:ow:10e2073a01080063","u":"lon","t":1.320067464e+09,
	 "v":24.30621},
	{"n":"urn:dev:ow:10e2073a01080063","u":"lat","t":1.320067464e+09,
	 "v":60.07965},
	{"n":"urn:dev:ow:10e2073a01080063","u":"%RH","t":1.320067524e+09,
	 "v":20.3},
	{"n":"urn:dev:ow:10e2073a01080063","u":"lon","t":1.320067524e+09,
	 "v":24.30622},
	{"n":"urn:dev:ow:10e2073a01080063","u":"lat","t":1.320067524e+09,
	 "v":60.07965},
	{"n":"urn:dev:ow:10e2073a01080063","u":"%RH","t":1.320067584e+09,
	 "v":20.7},
	{"n":"urn:dev:ow:10e2073a01080063","u":"lon","t":1.320067584e+09,
	 "v":24.30623},
	{"n":"urn:dev:ow:10e2073a01080063","u":"lat","t":1.320067584e+09,
	 "v":60.07966},
	{"n":"urn:dev:ow:10e2073a01080063","u":"%EL","t":1.320067614e+09,
	 "v":98},
	{"n":"urn:dev:ow:10e2073a01080063","u":"%RH","t":1.320067644e+09,
	 "v":21.2},
	{"n":"urn:dev:ow:10e2073a01080063","u":"lon","t":1.320067644e+09,
	 "v":24.30628},
	{"n":"urn:dev:ow:10e2073a01080063","u":"lat","t":1.320067644e+09,
	 "v":60.07967}
  ]
`

const multipleDataTypes = `
	[
    {"bn":"urn:dev:ow:10e2073a01080063:","n":"temp","u":"Cel","v":23.1},
    {"n":"label","vs":"Machine Room"},
    {"n":"open","vb":false},
    {"n":"nfv-reader","vd":"aGkgCg"}
  ]
`
