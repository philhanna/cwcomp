package model

import "fmt"
import "encoding/json"

// ---------------------------------------------------------------------
// Type definitions
// ---------------------------------------------------------------------

// Point is a row and column pair
type Point struct {
	Row int `json:"r"`
	Col int `json:"c"`
}

// ---------------------------------------------------------------------
// Methods
// ---------------------------------------------------------------------

// Compare returns -1, 0, or 1 depending on whether this point is
// less than, equal to, or greater than another.
func (p *Point) Compare(other Point) int {
	switch {
	case p.Row < other.Row:
		return -1
	case p.Row > other.Row:
		return 1
	case p.Col < other.Col:
		return -1
	case p.Col > other.Col:
		return 1
	default:
		return 0
	}
}

// Equal is true if this point has the same row and column of another
// point.
func (p *Point) Equal(other Point) bool {
	return *p == other
}

// FromJSON creates a Point from its JSON representation. See ToJSON.
func (p *Point) FromJSON(jsonBlob []byte) error {
	err := json.Unmarshal(jsonBlob, p)
	return err
}

// String returns a string representation of this type
func (p *Point) String() string {
	return fmt.Sprintf("(%d,%d)", p.Row, p.Col)
}

// ToJSON creates a JSON representatoin from a Point. See FromJSON.
func (p *Point) ToJSON() ([]byte, error) {
	result, err := json.Marshal(p)
	return result, err
}
