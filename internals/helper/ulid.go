package helper

import (
	"math/rand"
	"time"

	"github.com/Songmu/flextime"
	"github.com/oklog/ulid"
)

// NewULID @see https://qiita.com/kai_kou/items/b4ac2d316920e08ac75a
func NewULID() ulid.ULID {
	t := flextime.Now()
	entropy := ulid.Monotonic(rand.New(rand.NewSource(t.UnixNano())), 0)
	return ulid.MustNew(ulid.Timestamp(t), entropy)
}

func IsULID(str string) bool {
	id, err := ulid.Parse(str)
	if err != nil {
		return false
	}
	return str == id.String()
}

func TimeToULID(t time.Time) (ulid.ULID, error) {
	entropy := ulid.Monotonic(rand.New(rand.NewSource(t.UnixNano())), 0)
	return ulid.New(ulid.Timestamp(t), entropy)
}
