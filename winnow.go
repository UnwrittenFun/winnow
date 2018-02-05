package winnow

import (
	"encoding/json"
)

// A Grant declares the operations
// a user can do on a document in the collection
// that satisfies the given match
type Grant struct {
	Collection string                 `json:"collection"`
	Match      map[string]interface{} `json:"match"`
	Operations []string               `json:"operations"`
}

// A Winnow instance contains a list of permission grants
// which can then be queried using the Can method
type Winnow struct {
	Grants []Grant `json:"grants"`
}

// Can will return true if the grants in the winnow instance
// permit the requested operation on the collection, matching
// properties on the supplied document
func (w Winnow) Can(op, col string, doc interface{}) bool {
	for _, grant := range w.Grants {
		if grant.Collection != col || !contains(grant.Operations, op) {
			continue
		}

		if compare(remarshalToMap(doc), grant.Match) {
			return true
		}
	}

	return false
}

func contains(slice []string, item string) bool {
	set := make(map[string]struct{}, len(slice))
	for _, s := range slice {
		set[s] = struct{}{}
	}

	_, ok := set[item]
	return ok
}

func remarshalToMap(doc interface{}) (res map[string]interface{}) {
	data, err := json.Marshal(doc)
	if err != nil {
		// Ignore error, returning an empty map.
		return
	}

	if err := json.Unmarshal(data, &res); err != nil {
		// Ignore error, returning an empty map.
		return
	}

	return
}

func compare(doc, test map[string]interface{}) bool {
	for key, value := range test {
		if doc[key] != value {
			return false
		}
	}

	return true
}
