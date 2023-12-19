package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/swaggo/swag/example/celler/model"
)

// ListReleaseBundles godoc
//
//	@Summary		List ReleaseBundles
//	@Description	get Releasebundles
//	@Tags			ListReleaseBundles
//	@Accept			json
//	@Produce		json
//	@Param			username	query	string  false	"UserName"
//	@Param			pwd	        query	string  false	"Password"
//	@Success		200	{object}	model.ArtifactoryReleaseBundleSummary
//	@Failure		400	{object}	httputil.HTTPError
//	@Failure		404	{object}	httputil.HTTPError
//	@Failure		500	{object}	httputil.HTTPError
//	@Router			/ListReleaseBundles [get]
func (c *Controller) ListReleaseBundles(ctx *gin.Context) {

	username := ctx.Request.URL.Query().Get("username")
	pwd := ctx.Request.URL.Query().Get("pwd")

	fmt.Println("Getting Release Bundles List...")

	// make GET request to API to get user by ID
	apiUrl := "https://artifactory.devops.telekom.de/artifactory/api/release/bundles"
	// request, error := http.NewRequest("GET", apiUrl, nil)
	request, error := http.NewRequestWithContext(ctx, "GET", apiUrl, http.NoBody)

	if error != nil {
		fmt.Println(error)
	}

	request.Header.Set("Content-Type", "application/json; charset=utf-8")
	request.SetBasicAuth(username, pwd)

	client := &http.Client{}
	response, error := client.Do(request)

	if error != nil {
		fmt.Println(error)
	}

	// responseBody, error := ioutil.ReadAll(response.Body)

	body, err := ioutil.ReadAll(response.Body)
	// if err != nil {
	// 	return nil, err
	// }

	var r model.ArtifactoryReleaseBundles
	err = json.Unmarshal(body, &r)
	if err != nil {
		fmt.Errorf("cannot parse Artifactory response to list target bundles: %s", err)
	}

	var foundBundles []model.ArtifactoryReleaseBundleSummary

	for name, versions := range r.Bundles {
		for _, v := range versions {
			foundBundles = append(foundBundles,
				model.ArtifactoryReleaseBundleSummary{
					Name:    name,
					Version: v.Version,
					Status:  v.Status.String(),
					Type:    "TARGET",
				},
			)
		}
	}

	if error != nil {
		fmt.Println(error)
	}

	// formattedData := formatJSON(responseBody)
	// fmt.Println("Status: ", response.Status)
	// fmt.Println("Response body: ", formattedData)

	// if error != nil {
	// 	httputil.NewError(ctx, http.StatusNotFound, error)
	// 	return
	// }

	ctx.JSON(http.StatusOK, foundBundles)
}

// DeleteReleaseBundles godoc
//
//	@Summary		Delete ReleaseBundles
//	@Description	get Releasebundles
//	@Tags			DeleteReleaseBundles
//	@Accept			json
//	@Produce		json
//	@Param			username	query	string  false	"UserName"
//	@Param			pwd	        query	string  false	"Password"
//	@Param			bundlever	query	string  false	"Bundle Version"
//	@Success		200	{object}	controller.ArtifactoryReleaseBundleSummary
//	@Failure		400	{object}	httputil.HTTPError
//	@Failure		404	{object}	httputil.HTTPError
//	@Failure		500	{object}	httputil.HTTPError
//	@Router			/ListReleaseBundles [get]
// func (c *Controller) DeleteReleaseBundles(ctx *gin.Context) {
// 	// q := ctx.Request.URL.Query().Get("q")
// 	// accounts, err := model.AccountsAll(q)

// 	username := ctx.Request.URL.Query().Get("username")
// 	pwd := ctx.Request.URL.Query().Get("pwd")

// 	fmt.Println("Deleting Release Bundles List...")

// 	// make GET request to API to get user by ID
// 	apiUrl := "https://artifactory.devops.telekom.de/artifactory/api/release/bundles/%s"
// 	// request, error := http.NewRequest("GET", apiUrl, nil)
// 	request, error := http.NewRequestWithContext(ctx, "GET", apiUrl, http.NoBody)

// 	// if error != nil {
// 	// 	fmt.Println(error)
// 	// }

// 	request.Header.Set("Content-Type", "application/json; charset=utf-8")
// 	request.SetBasicAuth("rahul.bhatiya@t-systems.com", "cmVmdGtuOjAxOjE3MzMzMDE2MTc6dWNCOGkxNlNrcDA0ZngweUIwWmh6cXNQeUN4")

// 	client := &http.Client{}
// 	response, error := client.Do(request)

// 	// if error != nil {
// 	// 	fmt.Println(error)
// 	// }

// 	// responseBody, error := ioutil.ReadAll(response.Body)

// 	body, err := ioutil.ReadAll(response.Body)
// 	// if err != nil {
// 	// 	return nil, err
// 	// }

// 	var r ArtifactoryReleaseBundles
// 	err = json.Unmarshal(body, &r)
// 	if err != nil {
// 		fmt.Errorf("cannot parse Artifactory response to list target bundles: %s", err)
// 	}

// 	var foundBundles []ArtifactoryReleaseBundleSummary

// 	for name, versions := range r.Bundles {
// 		for _, v := range versions {
// 			foundBundles = append(foundBundles,
// 				ArtifactoryReleaseBundleSummary{
// 					Name:    name,
// 					Version: v.Version,
// 					Status:  v.Status.String(),
// 					Type:    "TARGET",
// 				},
// 			)
// 		}
// 	}

// 	if error != nil {
// 		fmt.Println(error)
// 	}

// 	// formattedData := formatJSON(responseBody)
// 	// fmt.Println("Status: ", response.Status)
// 	// fmt.Println("Response body: ", formattedData)

// 	// if error != nil {
// 	// 	httputil.NewError(ctx, http.StatusNotFound, error)
// 	// 	return
// 	// }

// 	ctx.JSON(http.StatusOK, foundBundles)
// }
