package basic

import (
	"errors"
	"testing"

	"github.com/magiconair/properties/assert"
)

func TestErrorEqual(t *testing.T){
	err := errors.New("EOF")
	err2 := err
	assert.Equal(t, err, err2)


}