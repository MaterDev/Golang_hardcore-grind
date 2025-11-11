// Package restapijsonmarshalling provides the JSON marshalling challenge implementation point.
//
// Problem: Define Go structs that marshal to the JSON shape:
// {
//   "id": 123,
//   "user_name": "shipt_user",
//   "is_active": true,
//   "geo_location": {"lat": 33.52, "lon": -86.80}
// }
//
// Requirements
// - Implement the structs below so their marshalled JSON EXACTLY matches the shape above.
// - Use appropriate struct tags and types so keys and values align with the example.
// - The tests will marshal and validate the JSON shape.
//
// Examples
//   u := User{ID: 1, UserName: "alice", IsActive: true, GeoLocation: Geo{Lat: 1.2, Lon: 3.4}}
//   b, _ := json.Marshal(u)
//
// Running tests
// - From repo root you can run just this challenge via the harness:
//     ./scripts/run.sh -challenge rest-api-json-marshalling -v
package restapijsonmarshalling

// Geo represents the nested geo_location object with latitude and longitude.
type Geo struct {
}

// User represents the top-level object.
type User struct {
}
