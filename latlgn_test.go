package latlgn_test

import (
	"platform/latlgn"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLatitudeLongitudeByCEP(t *testing.T) {
	cep := latlgn.CEP{
		// @todo
	}
	lat, lgn, err := latlgn.LatitudeLongitudeByCEP(cep)
	assert.NoError(t, err)
	assert.Equal(t, "-5.5203511", lat)
	assert.Equal(t, "-47.4824274", lgn)
}
