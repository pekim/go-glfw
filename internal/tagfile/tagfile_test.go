package tagfile

import (
	"encoding/xml"
	"testing"

	"github.com/stretchr/testify/assert"
)

func testUnmarshal(t *testing.T) Tagfile {
	t.Helper()

	var tagfile Tagfile
	err := xml.Unmarshal(tagfileContent, &tagfile)
	assert.NoError(t, err)
	return tagfile
}

func TestMapping(t *testing.T) {
	tagfile := testUnmarshal(t)

	assert.NotZero(t, len(tagfile.Compounds))

	compound, foundCompound := tagfile.findCompound("file", "glfw3.h")
	assert.True(t, foundCompound)
	assert.Equal(t, "file", compound.Kind)
	assert.NotEmpty(t, compound.Name)
	assert.NotEmpty(t, compound.Path)
	assert.NotEmpty(t, compound.Filename)

	member, foundMember := compound.findMember("glfwInit")
	assert.True(t, foundMember)
	assert.Equal(t, "function", member.Kind)
	assert.NotEmpty(t, member.Type)
	assert.NotEmpty(t, member.Name)
	assert.NotEmpty(t, member.Anchorfile)
	assert.NotEmpty(t, member.Anchor)
	assert.NotEmpty(t, member.Arglist)
}

func TestURL(t *testing.T) {
	tagfile := testUnmarshal(t)

	compound, foundCompound := tagfile.findCompound("file", "glfw3.h")
	assert.True(t, foundCompound)
	member, foundMember := compound.findMember("GLFW_FALSE")
	assert.True(t, foundMember)
	assert.Equal(t,
		"https://www.glfw.org/docs/latest/group__init.html#gac877fe3b627d21ef3a0a23e0a73ba8c5",
		member.URL(),
	)
}

func TestEntityURL(t *testing.T) {
	assert.Equal(t,
		"https://www.glfw.org/docs/latest/group__window.html#ga7cc0a09de172fa7250872046f8c4d2ca",
		EntityURLs()["GLFWimage"],
	)
}

func TestPageSectionURL(t *testing.T) {
	assert.Equal(t,
		URL{
			Title: "Support for versions of Windows older than XP",
			URL:   "https://www.glfw.org/docs/latest/moving_guide.html#moving_windows",
		},
		PageSectionURLs()["moving_windows"],
	)
}
