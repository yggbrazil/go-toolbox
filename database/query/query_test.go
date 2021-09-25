package query

import "testing"

func TestNew(t *testing.T) {
	var filter = struct {
		Name  string
		Phone string
	}{
		Name:  "Betiolo",
		Phone: "12345678",
	}

	q := New("SELECT * FROM Person WHERE ?", true)

	q.AddIf(filter.Name == "Betiolo", " AND Person.Name = ? ", filter.Name)
	q.AddIf(filter.Phone != "", " AND Person.Phone = ? ", filter.Phone)
	q.Add(" ORDER BY Person.Id")

	fullQuery := `SELECT * FROM Person WHERE ? AND Person.Name = ?  AND Person.Phone = ?  ORDER BY Person.Id`

	if q.String() != fullQuery {
		t.Error("query err")
	}

	if len(q.Args()) != 3 {
		t.Error("args err")
	}
}
