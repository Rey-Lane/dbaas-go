package dbaas

import (
	"context"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

const prometheusMetricTokenID = "20d7bcf4-f8d6-4bf6-b8f6-46cb440a87f4"

const testPrometheusMetricTokenNotFoundResponse = `{
	"error": {
		"code": 404,
		"title": "Not Found",
		"message": "prometheusmetrictoken 123 not found."
	}
}`

const testPrometheusMetricTokenResponse = `{
	"id": "20d7bcf4-f8d6-4bf6-b8f6-46cb440a87f4",
	"created_at": "1970-01-01T00:00:00",
	"updated_at": "1970-01-01T00:00:00",
	"project_id": "123e4567e89b12d3a456426655440000",
	"name": "token",
	"value": "GlEDgjR4oWaOjxy4a4YMorlrj81Jb93cR5Zpww6lx9fJs50dv3NygIB2zs3not5I"
}`

const testPrometheusMetricTokensResponse = `{
	"prometheus-metrics-tokens": [
		{
			"id": "20d7bcf4-f8d6-4bf6-b8f6-46cb440a87f4",
			"created_at": "1970-01-01T00:00:00",
			"updated_at": "1970-01-01T00:00:00",
			"project_id": "123e4567e89b12d3a456426655440000",
			"name": "token",
			"value": "GlEDgjR4oWaOjxy4a4YMorlrj81Jb93cR5Zpww6lx9fJs50dv3NygIB2zs3not5I"
		},
		{
			"id": "20d7bcf4-f8d6-4bf6-b8f6-46cb440a87f5",
			"created_at": "1970-01-01T00:00:00",
			"updated_at": "1970-01-01T00:00:00",
			"project_id": "123e4567e89b12d3a456426655440000",
			"name": "token123",
			"value": "GlEDgjR4oWaOjxy4a4YMorlrj81Jb93cR5Zpww6lx9fJs50dv3NygIB2zs3not52"
		}
	]
}`

func TestPrometheusMetricTokens(t *testing.T) {
	httpmock.Activate()
	testClient := SetupTestClient()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", testClient.Endpoint+"/prometheus-metrics-tokens",
		httpmock.NewStringResponder(200, testPrometheusMetricTokensResponse))

	expected := []PrometheusMetricToken{
		{
			ID:        "20d7bcf4-f8d6-4bf6-b8f6-46cb440a87f4",
			CreatedAt: "1970-01-01T00:00:00",
			UpdatedAt: "1970-01-01T00:00:00",
			ProjectID: "123e4567e89b12d3a456426655440000",
			Name:      "token",
			Value:     "GlEDgjR4oWaOjxy4a4YMorlrj81Jb93cR5Zpww6lx9fJs50dv3NygIB2zs3not5I",
		},
		{
			ID:        "20d7bcf4-f8d6-4bf6-b8f6-46cb440a87f5",
			CreatedAt: "1970-01-01T00:00:00",
			UpdatedAt: "1970-01-01T00:00:00",
			ProjectID: "123e4567e89b12d3a456426655440000",
			Name:      "token123",
			Value:     "GlEDgjR4oWaOjxy4a4YMorlrj81Jb93cR5Zpww6lx9fJs50dv3NygIB2zs3not52",
		},
	}

	actual, err := testClient.PrometheusMetricTokens(context.Background())

	if assert.NoError(t, err) {
		assert.Equal(t, expected, actual)
	}
}

func TestPrometheusMetricToken(t *testing.T) {
	httpmock.Activate()
	testClient := SetupTestClient()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", testClient.Endpoint+"/prometheus-metrics-tokens/"+prometheusMetricTokenID,
		httpmock.NewStringResponder(200, testPrometheusMetricTokenResponse))

	expected := PrometheusMetricToken{
		ID:        "20d7bcf4-f8d6-4bf6-b8f6-46cb440a87f4",
		CreatedAt: "1970-01-01T00:00:00",
		UpdatedAt: "1970-01-01T00:00:00",
		ProjectID: "123e4567e89b12d3a456426655440000",
		Name:      "token",
		Value:     "GlEDgjR4oWaOjxy4a4YMorlrj81Jb93cR5Zpww6lx9fJs50dv3NygIB2zs3not5I",
	}

	actual, err := testClient.PrometheusMetricToken(context.Background(), prometheusMetricTokenID)

	if assert.NoError(t, err) {
		assert.Equal(t, expected, actual)
	}
}

func TestPrometheusMetricTokenNotFound(t *testing.T) {
	httpmock.Activate()
	testClient := SetupTestClient()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", testClient.Endpoint+"/prometheus-metrics-tokens/123",
		httpmock.NewStringResponder(404, testPrometheusMetricTokenNotFoundResponse))

	expected := &DBaaSAPIError{}
	expected.APIError.Code = 404
	expected.APIError.Title = ErrorNotFoundTitle
	expected.APIError.Message = "prometheusmetrictoken 123 not found."

	_, err := testClient.PrometheusMetricToken(context.Background(), "123")

	assert.ErrorAs(t, err, &expected)
}

func TestCreatePrometheusMetricToken(t *testing.T) {
	httpmock.Activate()
	testClient := SetupTestClient()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("POST", testClient.Endpoint+"/prometheus-metrics-tokens",
		func(req *http.Request) (*http.Response, error) {
			if err := json.NewDecoder(req.Body).Decode(&PrometheusMetricTokenCreateOpts{}); err != nil {
				return httpmock.NewStringResponse(400, ""), err
			}

			tokens := make(map[string]PrometheusMetricToken)
			tokens["prometheus-metrics-token"] = PrometheusMetricToken{
				ID:        "20d7bcf4-f8d6-4bf6-b8f6-46cb440a87f4",
				CreatedAt: "1970-01-01T00:00:00",
				UpdatedAt: "1970-01-01T00:00:00",
				ProjectID: "123e4567e89b12d3a456426655440000",
				Name:      "token",
				Value:     "GlEDgjR4oWaOjxy4a4YMorlrj81Jb93cR5Zpww6lx9fJs50dv3NygIB2zs3not5I",
			}

			resp, err := httpmock.NewJsonResponse(200, tokens)
			if err != nil {
				return httpmock.NewStringResponse(500, ""), err
			}

			return resp, nil
		})

	createPrometheusMetricTokenOpts := PrometheusMetricTokenCreateOpts{
		Name: "token",
	}

	expected := PrometheusMetricToken{
		ID:        "20d7bcf4-f8d6-4bf6-b8f6-46cb440a87f4",
		CreatedAt: "1970-01-01T00:00:00",
		UpdatedAt: "1970-01-01T00:00:00",
		ProjectID: "123e4567e89b12d3a456426655440000",
		Name:      "token",
		Value:     "GlEDgjR4oWaOjxy4a4YMorlrj81Jb93cR5Zpww6lx9fJs50dv3NygIB2zs3not5I",
	}

	actual, err := testClient.CreatePrometheusMetricToken(context.Background(), createPrometheusMetricTokenOpts)

	if assert.NoError(t, err) {
		assert.Equal(t, expected, actual)
	}
}

func TestUpdatePrometheusMetricToken(t *testing.T) {
	httpmock.Activate()
	testClient := SetupTestClient()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("PUT", testClient.Endpoint+"/prometheus-metrics-tokens/"+prometheusMetricTokenID,
		func(req *http.Request) (*http.Response, error) {
			if err := json.NewDecoder(req.Body).Decode(&PrometheusMetricTokenUpdateOpts{}); err != nil {
				return httpmock.NewStringResponse(400, ""), err
			}

			token := PrometheusMetricToken{
				ID:        "20d7bcf4-f8d6-4bf6-b8f6-46cb440a87f4",
				CreatedAt: "1970-01-01T00:00:00",
				UpdatedAt: "1970-01-01T00:00:00",
				ProjectID: "123e4567e89b12d3a456426655440000",
				Name:      "token123",
				Value:     "GlEDgjR4oWaOjxy4a4YMorlrj81Jb93cR5Zpww6lx9fJs50dv3NygIB2zs3not5I",
			}

			resp, err := httpmock.NewJsonResponse(200, token)
			if err != nil {
				return httpmock.NewStringResponse(500, ""), err
			}

			return resp, nil
		})

	updatePrometheusMetricTokenOpts := PrometheusMetricTokenUpdateOpts{
		Name: "token123",
	}

	expected := PrometheusMetricToken{
		ID:        "20d7bcf4-f8d6-4bf6-b8f6-46cb440a87f4",
		CreatedAt: "1970-01-01T00:00:00",
		UpdatedAt: "1970-01-01T00:00:00",
		ProjectID: "123e4567e89b12d3a456426655440000",
		Name:      "token123",
		Value:     "GlEDgjR4oWaOjxy4a4YMorlrj81Jb93cR5Zpww6lx9fJs50dv3NygIB2zs3not5I",
	}

	actual, err := testClient.UpdatePrometheusMetricToken(context.Background(), prometheusMetricTokenID, updatePrometheusMetricTokenOpts)

	if assert.NoError(t, err) {
		assert.Equal(t, expected, actual)
	}
}
