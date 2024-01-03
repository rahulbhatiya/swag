package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/swaggo/swag/example/celler/model"
)

// Define the HTTPS_PROXY environment variable
var httpsProxy = ""

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

	// Set proxy environment variables programmatically
	os.Setenv("HTTPS_PROXY", httpsProxy)

	// make apiurl
	apiUrl := "https://artifactory-itg.test.devops.telekom.de/artifactory/api/release/bundles"

	// request, error := http.NewRequest("GET", apiUrl, nil)
	request, error := http.NewRequestWithContext(ctx, "GET", apiUrl, http.NoBody)

	if error != nil {
		fmt.Println(error)
	}

	request.Header.Set("Content-Type", "application/json; charset=utf-8")
	request.SetBasicAuth(username, pwd)

	// Create an HTTP client with proxy settings from the environment
	client := &http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyFromEnvironment,
		},
	}

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

// VerDeleteReleaseBundles godoc

// @Summary		    Delete ReleaseBundles
// @Description	    DELETE Releasebundles
// @Tags			DeleteReleaseBundles
// @Accept			json
// @Produce		    json
// @Param			username	query	string  false	"UserName"
// @Param			pwd	        query	string  false	"Password"
// @Param			BundleName	query	string  false	"Bundle Name"
// @Success		200	{object}	string
// @Failure		400	{object}	httputil.HTTPError
// @Failure		404	{object}	httputil.HTTPError
// @Failure		500	{object}	httputil.HTTPError
// @Router			/VerDeleteReleaseBundles [DELETE]
func (c *Controller) VerDeleteReleaseBundles(ctx *gin.Context) {

	username := ctx.Request.URL.Query().Get("username")
	pwd := ctx.Request.URL.Query().Get("pwd")
	bname := ctx.Request.URL.Query().Get("BundleName")

	fmt.Println("Deleting Release Bundles List...")

	// make GET request to API to get user by ID
	apiUrl := "https://artifactory-itg.test.devops.telekom.de/artifactory/api/release/bundles/" + bname
	// request, error := http.NewRequest("GET", apiUrl, nil)
	request, error := http.NewRequestWithContext(ctx, "DELETE", apiUrl, http.NoBody)

	if error != nil {
		fmt.Println(error)
	}

	// request.Header.Set("Content-Type", "application/json; charset=utf-8")
	request.SetBasicAuth(username, pwd)

	client := &http.Client{}
	response, error := client.Do(request)

	if error != nil {
		fmt.Println(error)
	}

	// responseBody, error := ioutil.ReadAll(response.Body)

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Errorf("cannot parse Artifactory response to list target bundles: %s", err)
	}

	if response.StatusCode != 200 {
		fmt.Errorf("cannot delete target release bundle %s: %s", response.Status, string(body))
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

	ctx.JSON(http.StatusOK, response.Status)
}
