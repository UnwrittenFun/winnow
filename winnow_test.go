package winnow

import "testing"

func TestCan(t *testing.T) {
	w := &Winnow{
		Grants: []Grant{
			{
				Collection: "posts",
				Match: map[string]interface{}{
					"authorId": "1",
				},
				Operations: []string{"read"},
			},
			{
				Collection: "pages",
				Operations: []string{"read"},
			},
		},
	}

	// Test operation and document matching
	if !w.Can("read", "posts", map[string]interface{}{"authorId": "1"}) {
		t.Errorf("Can read posts(authorId=1) was false, expected true")
	}

	// Test operation matching
	if w.Can("update", "posts", map[string]interface{}{"authorId": "1"}) {
		t.Errorf("Can update posts(authorId=1) was true, expected false")
	}

	// Test document matching
	if w.Can("read", "posts", map[string]interface{}{"authorId": "2"}) {
		t.Errorf("Can read posts(authorId=2) was true, expected false")
	}

	// Test grant with no match
	if !w.Can("read", "pages", map[string]interface{}{"authorId": "1"}) {
		t.Errorf("Can read pages(authorId=1) was false, expected true")
	}

	// Test Can with no document
	if w.Can("read", "posts", nil) {
		t.Errorf("Can read posts(authorId=1) was true, expected false")
	}
}
