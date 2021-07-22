package behavioral

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIterator(t *testing.T) {
	user1 := &User{
		Grade: 1,
	}
	user2 := &User{
		Grade: 2,
	}

	userCollection := &UserCollection{
		users: []*User{user1, user2},
	}

	iterator := userCollection.GetIterator()

	var max int
	for iterator.HasMore() {
		u := iterator.GetNext()
		if u.Grade > max {
			max += u.Grade
		}
	}
	assert.Equal(t, 3, max)
}
