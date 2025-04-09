package latlgn

import (
	"context"
	"errors"
	"log"
	"platform/viacep"
	"strconv"

	"github.com/alexliesenfeld/opencage"
)

const key string = "YOUR_KEY"

type ParamsFindCep struct {
	From  int    `json:"from"`
	Size  int    `json:"size"`
	Query string `json:"query_string"`
}

type FindCep struct {
	ID     string        `json:"id"`
	Params ParamsFindCep `json:"params"`
}

func LatitudeLongitudeByCEP(cep viacep.CEP) (lat, lgn string, err error) {

	query := cep.Logradouro + " " + cep.Localidade + " " + cep.UF + " " + "Brasil"
	lat, lgn, err = callToOpenCageData(query)
	if err != nil {
		return "", "", err
	}

	return lat, lgn, nil
}

func callToOpenCageData(query string) (lat, lgn string, err error) {
	client := opencage.New(key)

	// Can be used to control timeouts (e.g., using context.WithDeadline), cancellation, etc.
	ctx := context.Background()

	// Call the API with default parameters.
	// Also shows how to use forward geocoding API using a location.
	response, err := client.Geocode(ctx, query, nil)
	if err != nil {
		log.Panicln("quota exceeded")
	}

	// Or set your own API parameters.
	// Also shows how to use reverse geocoding API using latitude and longitude.
	//response, err = client.Geocode(ctx, "52.3877830 9.7334394", &opencage.GeocodingParams{
	//	RoadInfo:  true,
	//	Proximity: []float32{1.0, -1.0},
	//	Language:  "de",
	//	})

	if response.Status.Code != 200 {
		return "", "", errors.New("deu ruim no sistema do cep")
	}

	if len(response.Results) == 0 {
		return "", "", errors.New("não venho nenhum resultado")
	}

	lat = strconv.FormatFloat(response.Results[0].Geometry.Lat, 'f', -1, 64)
	lgn = strconv.FormatFloat(response.Results[0].Geometry.Lng, 'f', -1, 64)

	return lat, lgn, nil
}

// func callToFindCep(query string) (lat, lgn string, err error) {
// 	findCep := FindCep{
// 		ID: "pesquisa_endereco",
// 		Params: ParamsFindCep{
// 			From:  0,
// 			Size:  2,
// 			Query: query,
// 		},
// 	}

// 	payload, _ := json.Marshal(&findCep)

// 	// Can be used to control timeouts (e.g., using context.WithDeadline), cancellation, etc.
// 	ctx := context.Background()

// 	// Call the API with default parameters.
// 	// Also shows how to use forward geocoding API using a location.
// 	response, err := client.Geocode(ctx, query, nil)
// 	if err != nil {
// 		log.Panicln("quota exceeded")
// 	}

// 	// Or set your own API parameters.
// 	// Also shows how to use reverse geocoding API using latitude and longitude.
// 	//response, err = client.Geocode(ctx, "52.3877830 9.7334394", &opencage.GeocodingParams{
// 	//	RoadInfo:  true,
// 	//	Proximity: []float32{1.0, -1.0},
// 	//	Language:  "de",
// 	//	})

// 	if response.Status.Code != 200 {
// 		return "", "", errors.New("deu ruim no sistema do cep")
// 	}

// 	if len(response.Results) == 0 {
// 		return "", "", errors.New("não venho nenhum resultado")
// 	}

// 	lat = strconv.FormatFloat(response.Results[0].Geometry.Lat, 'f', -1, 64)
// 	lgn = strconv.FormatFloat(response.Results[0].Geometry.Lng, 'f', -1, 64)

// 	return lat, lgn, nil
// }
