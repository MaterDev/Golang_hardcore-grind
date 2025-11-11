// Package restapijsonmarshalling contains tests for the Basic REST API JSON Marshalling challenge.
// Implement the structs in solution.go so these tests pass.
package restapijsonmarshalling

import (
	"encoding/json"
	"math"
	"testing"
)

func TestMarshalUser_JSONShape(t *testing.T) {
	u := User{
		ID:         123,
		UserName:   "shipt_user",
		IsActive:   true,
		GeoLocation: Geo{Lat: 33.52, Lon: -86.80},
	}

	b, err := json.Marshal(u)
	if err != nil {
		t.Fatalf("marshal error: %v", err)
	}

	// Unmarshal to a generic map so we can validate key names and values
	var got map[string]any
	if err := json.Unmarshal(b, &got); err != nil {
		t.Fatalf("unmarshal roundtrip: %v", err)
	}

	// Top-level keys must exist with exact snake_case names
	id, ok := got["id"].(float64)
	if !ok || int(id) != 123 {
		t.Fatalf("expected id=123, got %v (ok=%v)", got["id"], ok)
	}
	uname, ok := got["user_name"].(string)
	if !ok || uname != "shipt_user" {
		t.Fatalf("expected user_name=shipt_user, got %v (ok=%v)", got["user_name"], ok)
	}
	active, ok := got["is_active"].(bool)
	if !ok || active != true {
		t.Fatalf("expected is_active=true, got %v (ok=%v)", got["is_active"], ok)
	}

	geo, ok := got["geo_location"].(map[string]any)
	if !ok {
		t.Fatalf("expected geo_location object, got %T", got["geo_location"])
	}
	lat, ok := geo["lat"].(float64)
	if !ok || !floatEq(lat, 33.52) {
		t.Fatalf("expected lat≈33.52, got %v (ok=%v)", geo["lat"], ok)
	}
	lon, ok := geo["lon"].(float64)
	if !ok || !floatEq(lon, -86.8) {
		t.Fatalf("expected lon≈-86.8, got %v (ok=%v)", geo["lon"], ok)
	}
}

func floatEq(a, b float64) bool { return math.Abs(a-b) < 1e-9 }

func BenchmarkMarshalUser(b *testing.B) {
	u := User{ID: 123, UserName: "shipt_user", IsActive: true, GeoLocation: Geo{Lat: 33.52, Lon: -86.80}}
	for i := 0; i < b.N; i++ {
		_, _ = json.Marshal(u)
	}
}
