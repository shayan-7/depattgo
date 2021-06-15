package creational

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPrototype(t *testing.T) {

	// Testing the File concrete type
	originFile := NewFile("original.go", true, Root)
	clonedFile := originFile.Clone()

	assert.Equal(t, originFile.owner, clonedFile.GetOwner())
	assert.Equal(t, originFile.readonly, clonedFile.GetReadOnly())

	clonedFile.SetReadOnly(false)
	assert.NotEqual(t, originFile.readonly, clonedFile.GetReadOnly())

	clonedFile.SetOwner(Guest)
	assert.NotEqual(t, originFile.owner, clonedFile.GetOwner())

	// Testing the Directory concrete type
	originDirectory := NewFile("/tmp/original/", true, Root)
	clonedDirectory := originFile.Clone()

	assert.Equal(t, originDirectory.owner, clonedDirectory.GetOwner())
	assert.Equal(t, originDirectory.readonly, clonedDirectory.GetReadOnly())

	clonedDirectory.SetReadOnly(false)
	assert.NotEqual(t, originDirectory.readonly, clonedDirectory.GetReadOnly())

	clonedDirectory.SetOwner(Guest)
	assert.NotEqual(t, originDirectory.owner, clonedDirectory.GetOwner())
}
