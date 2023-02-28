package IDUtil

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func ObjectID() string {
	h := 16

	// Get ms in ms
	ms := time.Now().UTC().UnixMilli()

	// Convert int64 to string
	strTime := strconv.FormatInt(ms, 10)

	// Get the substring from position 0 to 8
	strTime = strTime[:8]

	// Converts `n` into hexadecimal value
	x := func() string {
		return fmt.Sprintf("%x", rand.Intn(h))
	}

	// Add a random hexadecimal value 16x
	for i := 0; i < h; i++ {
		strTime += x()
	}

	return strTime
}
