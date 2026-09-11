package glfw

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func testInit(t *testing.T) {
	t.Helper()
	assert.NoError(t, Initialise())
	if Init() != TRUE {
		errMessage, err := GetError()
		assert.Fail(t, "Init failed", "%d : %q", err, errMessage)
	}
}

func TestVersion(t *testing.T) {
	testInit(t)

	major, minor, rev := GetVersion()
	assert.Equal(t, Int(3), major)
	assert.GreaterOrEqual(t, minor, Int(3))

	version := GetVersionString()
	expectedPrefix := fmt.Sprintf("%d.%d.%d", major, minor, rev)
	assert.True(t, strings.HasPrefix(version, expectedPrefix))
}

func TestStructPointerResult(t *testing.T) {
	testInit(t)

	assert.NotNil(t, GetPrimaryMonitor())
}

func TestCreateWindow(t *testing.T) {
	testInit(t)

	assert.NotNil(t, CreateWindow(800, 600, "title", nil, nil))
}

func TestNewCallback(t *testing.T) {
	testInit(t)

	assert.NotNil(t, ErrorCallbackNew(func(_error_code Int, _description string) {}))
	assert.NotNil(t, CursorposCallbackNew(func(_window *Window, _xpos, _ypos Double) {}))
}
