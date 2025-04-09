package latlgn_test

import (
	"platform/latlgn"
	"platform/viacep"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLatitudeLongitudeByCEP(t *testing.T) {
	cep, err := viacep.SearchCEP("65900520")
	assert.NoError(t, err)
	lat, lgn, err := latlgn.LatitudeLongitudeByCEP(cep)
	assert.NoError(t, err)
	assert.Equal(t, "-5.5203511", lat)
	assert.Equal(t, "-47.4824274", lgn)
}
