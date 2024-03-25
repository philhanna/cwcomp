package acrosslite

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

/*
	type fields struct {
		Size         int
		Name         string
		Title        string
		Author       string
		Copyright    string
		Grid         []string
		AcrossClues  map[int]string
		DownClues    map[int]string
		CreatedDate  time.Time
		ModifiedDate time.Time
	}
*/

func CreateTestObject() *AcrossLite {
	var pal = NewAcrossLite()
	pal.Size = 15
	pal.Name = "disney"
	pal.Title = "Failed the Audition"
	pal.Author = "Jack London"
	pal.Copyright = "Not copyrighted"
	pal.Grid = []string{
		"FATE.AWASH.AWOL",
		"LIES.CURIO.SHOE",
		"ELECTORATE.SIZE",
		"ASS.ERST.DIETED",
		"...CENT.HOSTESS",
		"REFITS.JEWISH..",
		"ARITH.KERNS.OAF",
		"NILE.ANNES.DUPE",
		"DEI.OVENS.LOSER",
		"..BODILY.RACERS",
		"GLUTEAL.PEPS...",
		"RESIST.SLUE.SKI",
		"OTTO.REPUBLICAN",
		"OMES.IRATE.RAMS",
		"MERE.XENON.ABET",
	}
	pal.AcrossClues = map[int]string{
		1: "Clue for Fate",
		5: "Clue for Awash",
	}
	pal.DownClues = map[int]string{
		1: "Clue for Flea",
		2: "Clue for Ails",
	}
	t := time.Now()
	pal.CreatedDate = t.Add(-(time.Hour * 24))
	pal.ModifiedDate = t
	return pal
}

func TestNewAcrossLite(t *testing.T) {
	var pal = NewAcrossLite()
	assert.Equal(t, 0, len(pal.Grid))
	assert.Equal(t, 0, len(pal.AcrossClues))
	assert.Equal(t, 0, len(pal.DownClues))
}

func TestAcrossLite_GetAcrossClues(t *testing.T) {
	tests := []struct {
		name string
		pal  *AcrossLite
		want map[int]string
	}{
		{pal: NewAcrossLite(), want: make(map[int]string)},
		{
			pal: func() *AcrossLite {
				pal := NewAcrossLite()
				pal.AcrossClues[3] = "Hello"
				return pal
			}(),
			want: map[int]string{3: "Hello"},
		},
		{
			pal: func() *AcrossLite {
				pal := NewAcrossLite()
				pal.SetAcrossClues(map[int]string{1: "New Clue"})
				return pal
			}(),
			want: map[int]string{1: "New Clue"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.pal.GetAcrossClues())
		})
	}
}

func TestAcrossLite_GetAuthor(t *testing.T) {
	tests := []struct {
		name string
		pal  *AcrossLite
		want string
	}{
		{pal: NewAcrossLite(), want: ""},
		{
			pal: func() *AcrossLite {
				pal := NewAcrossLite()
				pal.SetAuthor("Jack London")
				return pal
			}(),
			want: "Jack London",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.pal.GetAuthor())
		})
	}
}

func TestAcrossLite_GetCell(t *testing.T) {
	tests := []struct {
		name    string
		pal     *AcrossLite
		r       int
		c       int
		letter  byte
		want    byte
		wantErr bool
	}{
		{
			pal:     NewAcrossLite(),
			wantErr: true,
		},
		{r: -3, c: 3, wantErr: true},
		{r: 1, c: 1, letter: 'X', want: 'X'},
		{r: 1, c: 1, want: 'F'},
		{r: 1, c: 5, want: 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var pal *AcrossLite
			if tt.pal == nil {
				pal = CreateTestObject()
			} else {
				pal = tt.pal
			}
			if tt.letter != 0 {
				pal.SetCell(tt.r, tt.c, tt.letter)
			}
			switch tt.wantErr {
			case false:
				have, err := pal.GetCell(tt.r, tt.c)
				assert.Nil(t, err)
				want := tt.want
				assert.Equal(t, want, have)
			case true:
				_, err := pal.GetCell(tt.r, tt.c)
				assert.NotNil(t, err)
			}
		})
	}
}

func TestAcrossLite_GetCopyright(t *testing.T) {
	tests := []struct {
		name string
		pal  *AcrossLite
		want string
	}{
		{name: "Good", pal: CreateTestObject(), want: "Not copyrighted"},
		{name: "Bad", pal: NewAcrossLite()},
		{name: "TestSetter", pal: func() *AcrossLite {
			pal := CreateTestObject()
			pal.SetCopyright("1992")
			return pal
		}(), want: "1992"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.pal.GetCopyright())
		})
	}
}

func TestAcrossLite_GetCreatedDate(t *testing.T) {
	now := time.Now()

	tests := []struct {
		name string
		pal  *AcrossLite
		want time.Time
	}{
		{
			pal: func() *AcrossLite {
				pal := CreateTestObject()
				pal.SetCreatedDate(now)
				return pal
			}(),
			want: now,
		},
		{
			pal:  NewAcrossLite(),
			want: now,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			const DATE_FORMAT = "2006-01-02"
			want := tt.want.Format(DATE_FORMAT)
			have := tt.pal.GetCreatedDate().Format(DATE_FORMAT)
			assert.Equal(t, want, have)
		})
	}
}

func TestAcrossLite_GetDownClues(t *testing.T) {
	tests := []struct {
		name string
		pal  *AcrossLite
		want map[int]string
	}{
		{
			pal:  CreateTestObject(),
			want: map[int]string{1: "Clue for Flea"},
		},
		{
			pal: func() *AcrossLite {
				pal := CreateTestObject()
				pal.SetDownClues(map[int]string{1: "New clue"})
				return pal
			}(),
			want: map[int]string{1: "New clue"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			wantMap := tt.want
			want := wantMap[1]
			haveMap := tt.pal.GetDownClues()
			have := haveMap[1]
			assert.Equal(t, want, have)
		})
	}
}

func TestAcrossLite_GetGrid(t *testing.T) {
	tests := []struct {
		name string
		pal  *AcrossLite
		want []string
	}{
		{pal: CreateTestObject(), want: []string{
			"FATE.AWASH.AWOL",
			"LIES.CURIO.SHOE",
			"ELECTORATE.SIZE",
			"ASS.ERST.DIETED",
			"...CENT.HOSTESS",
			"REFITS.JEWISH..",
			"ARITH.KERNS.OAF",
			"NILE.ANNES.DUPE",
			"DEI.OVENS.LOSER",
			"..BODILY.RACERS",
			"GLUTEAL.PEPS...",
			"RESIST.SLUE.SKI",
			"OTTO.REPUBLICAN",
			"OMES.IRATE.RAMS",
			"MERE.XENON.ABET",
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			want := tt.want
			have := tt.pal.GetGrid()
			assert.Equal(t, want, have)
		})
	}
}

func TestAcrossLite_GetModifiedDate(t *testing.T) {
	const DATE_FORMAT = "2006-01-02"

	now := time.Now()
	isLeap := func(year int) bool {
		if year%400 == 0 {
			return true
		}
		if year%100 == 0 {
			return false
		}
		return year%4 == 0
	}

	maxDays := func(year, month int) int {
		switch month {
		case 2:
			if isLeap(year) {
				return 29
			} else {
				return 28
			}
		case 1, 3, 5, 7, 8, 10, 12:
			return 31
		default:
			return 30
		}
	}

	tests := []struct {
		name string
		pal  *AcrossLite
		want time.Time
	}{
		{
			pal: NewAcrossLite(),
			want: func() time.Time {
				year, month, day := now.Date()
				day += 3
				if day > maxDays(year, int(month)) {
					month++
					if month > 12 {
						year++
						month = 1
					}
				}
				s := fmt.Sprintf("%04d-%02d-%02d", year, month, day)
				newTime, _ := time.Parse(DATE_FORMAT, s)
				return newTime
			}(),
		},
		{
			pal: func() *AcrossLite {
				pal := CreateTestObject()
				pal.SetModifiedDate(now)
				return pal
			}(),
			want: now,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			want := tt.want.Format(DATE_FORMAT)
			have := tt.pal.GetModifiedDate().Format(DATE_FORMAT)
			assert.Equal(t, want, have)
		})
	}

}

func TestAcrossLite_GetName(t *testing.T) {
	tests := []struct {
		name string
		pal  *AcrossLite
		want string
	}{
		{
			pal:  CreateTestObject(),
			want: "disney",
		},
		{
			pal: func() *AcrossLite {
				pal := CreateTestObject()
				pal.SetName("waffle")
				return pal
			}(),
			want: "waffle",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			want := tt.want
			have := tt.pal.GetName()
			assert.Equal(t, want, have)
		})
	}
}

func TestAcrossLite_GetNotepad(t *testing.T) {
	tests := []struct {
		name string
		pal  *AcrossLite
	}{
		{
			name: "Good",
			pal:  CreateTestObject(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			have := tt.pal.GetNotepad()
			assert.Contains(t, have, "created")
			assert.Contains(t, have, "modified")
		})
	}
}

func TestAcrossLite_GetTitle(t *testing.T) {
	tests := []struct {
		name string
		pal  *AcrossLite
		want string
	}{
		{pal: NewAcrossLite(), want: ""},
		{
			pal: func() *AcrossLite {
				pal := NewAcrossLite()
				pal.SetTitle("Whoops!")
				return pal
			}(),
			want: "Whoops!",
		},
		{
			pal:  CreateTestObject(),
			want: "Failed the Audition",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.pal.GetTitle())
		})
	}
}

func TestAcrossLite_SetCell(t *testing.T) {
	tests := []struct {
		name    string
		pal     *AcrossLite
		r       int
		c       int
		letter  byte
		wantErr bool
	}{
		{
			pal: func() *AcrossLite {
				pal := CreateTestObject()
				pal.SetCell(1, 1, '\x00')
				return pal
			}(),
			r:      1,
			c:      1,
			letter: '\x00',
		},
		{
			pal: func() *AcrossLite {
				pal := CreateTestObject()
				pal.SetSize(0)
				return pal
			}(),
			wantErr: true,
		},
		{
			pal:     CreateTestObject(),
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pal := tt.pal

			switch tt.wantErr {
			case false:
				want := tt.letter
				have, err := pal.GetCell(tt.r, tt.c)
				assert.Nil(t, err)
				assert.Equal(t, want, have)
			case true:
				if pal.GetSize() == 0 {
					err := pal.SetCell(1, 1, 'a')
					assert.NotNil(t, err)
				} else {
					err := pal.SetCell(-3, -6, 'a')
					assert.NotNil(t, err)
				}
			}
		})
	}
}

func TestAcrossLite_SetSize(t *testing.T) {
	tests := []struct {
		name    string
		pal     *AcrossLite
		newSize int
	}{
		{
			pal:     CreateTestObject(),
			newSize: 21,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pal := tt.pal
			want := tt.newSize
			pal.SetSize(tt.newSize)
			have := pal.GetSize()
			assert.Equal(t, want, have)
		})
	}
}
