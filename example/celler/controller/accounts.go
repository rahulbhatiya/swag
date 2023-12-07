package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

// ShowAccount godoc
//
//	@Summary		Show an account
//	@Description	get string by ID
//	@Tags			accounts
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"Account ID"
//	@Success		200	{object}	model.Account
//	@Failure		400	{object}	httputil.HTTPError
//	@Failure		404	{object}	httputil.HTTPError
//	@Failure		500	{object}	httputil.HTTPError
//	@Router			/accounts/{id} [get]
// func (c *Controller) ShowAccount(ctx *gin.Context) {
// 	id := ctx.Param("id")
// 	aid, err := strconv.Atoi(id)
// 	if err != nil {
// 		httputil.NewError(ctx, http.StatusBadRequest, err)
// 		return
// 	}
// 	account, err := model.AccountOne(aid)
// 	if err != nil {
// 		httputil.NewError(ctx, http.StatusNotFound, err)
// 		return
// 	}
// 	ctx.JSON(http.StatusOK, account)
// }

// ListReleaseBundles godoc
//
//	@Summary		List ReleaseBundles
//	@Description	get Releasebundles
//	@Tags			accounts
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	controller.ArtifactoryReleaseBundleSummary
//	@Failure		400	{object}	httputil.HTTPError
//	@Failure		404	{object}	httputil.HTTPError
//	@Failure		500	{object}	httputil.HTTPError
//	@Router			/accounts [get]
func (c *Controller) ListReleaseBundles(ctx *gin.Context) {
	// q := ctx.Request.URL.Query().Get("q")
	// accounts, err := model.AccountsAll(q)

	fmt.Println("Getting Release Bundles List...")

	// make GET request to API to get user by ID
	apiUrl := "https://artifactory.devops.telekom.de/artifactory/api/release/bundles"
	// request, error := http.NewRequest("GET", apiUrl, nil)
	request, error := http.NewRequestWithContext(ctx, "GET", apiUrl, http.NoBody)

	// if error != nil {
	// 	fmt.Println(error)
	// }

	request.Header.Set("Content-Type", "application/json; charset=utf-8")
	request.SetBasicAuth("rahul.bhatiya@t-systems.com", "cmVmdGtuOjAxOjE3MzMzMDE2MTc6dWNCOGkxNlNrcDA0ZngweUIwWmh6cXNQeUN4")

	client := &http.Client{}
	response, error := client.Do(request)

	// if error != nil {
	// 	fmt.Println(error)
	// }

	// responseBody, error := ioutil.ReadAll(response.Body)

	body, err := ioutil.ReadAll(response.Body)
	// if err != nil {
	// 	return nil, err
	// }

	var r ArtifactoryReleaseBundles
	err = json.Unmarshal(body, &r)
	if err != nil {
		fmt.Errorf("cannot parse Artifactory response to list target bundles: %s", err)
	}

	var foundBundles []ArtifactoryReleaseBundleSummary

	for name, versions := range r.Bundles {
		for _, v := range versions {
			foundBundles = append(foundBundles,
				ArtifactoryReleaseBundleSummary{
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

func formatJSON(custInfo interface{}) string {
	val, err := json.MarshalIndent(custInfo, "", "    ")
	if err != nil {
		return ""
	}
	return string(val)
}

type ArtifactoryReleaseBundleSummary struct {
	Name    string
	Version string
	Created string
	Status  string
	Type    string
}

// BundleVersionStatus is an alias type for bundles statuses defined below
type BundleVersionStatus string

func (s *BundleVersionStatus) String() string { return string(*s) }

// ArtifactoryReleaseBundles is a set of bundles in an Artifactory response
type ArtifactoryReleaseBundles struct {
	Bundles map[string][]ArtifactoryReleaseBundleVersionStatus
}

// ArtifactoryReleaseBundleVersionStatus descripes a release bundle version in Artifactory response
type ArtifactoryReleaseBundleVersionStatus struct {
	Version string              `json:"version"`
	Created string              `json:"created"`
	Status  BundleVersionStatus `json:"status"`
}

// AccountsAll example
// func AccountsAll(q string) ([]Account, error) {
// 	if q == "" {
// 		return accounts, nil
// 	}
// 	as := []Account{}
// 	for k, v := range accounts {
// 		if q == v.Name {
// 			as = append(as, accounts[k])
// 		}
// 	}
// 	return as, nil
// }

// GetTargetBundles retrieve full list of target bundles
// func (cln *Client) GetTargetBundles(ctx context.Context) ([]ArtifactoryReleaseBundleSummary, error) {
// func (c *Controller) GetTargetBundles(ctx *gin.Context) {
// 	url := "https://artifactory.devops.telekom.de/artifactory/api/release/bundles"
// 	req, err := c.mkReq(ctx, http.MethodGet, url, http.NoBody)

// 	resp, err := c.httpClient.Do(req)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer resp.Body.Close()

// 	if resp.StatusCode != 200 {
// 		return nil, fmt.Errorf("cannot list target bundles: %s", resp.Status)
// 	}

// 	body, err := ioutil.ReadAll(resp.Body)
// 	if err != nil {
// 		return nil, err
// 	}

// 	var r ArtifactoryReleaseBundles
// 	err = json.Unmarshal(body, &r)
// 	if err != nil {
// 		return nil, fmt.Errorf("cannot parse Artifactory response to list target bundles: %s", err)
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

// 	return foundBundles, nil
// }

// func RB() {
//     fmt.Println("Getting Release Bundles List...")

//     // make GET request to API to get user by ID
//     apiUrl := "https://artifactory.devops.telekom.de/artifactory/api/release/bundles"
//     request, error := http.NewRequest("GET", apiUrl, nil)

//     if error != nil {
//         fmt.Println(error)
//     }

//     request.Header.Set("Content-Type", "application/json; charset=utf-8")
// 	request.SetBasicAuth("rahul.bhatiya@t-systems.com", "cmVmdGtuOjAxOjE3MzMzMDE2MTc6dWNCOGkxNlNrcDA0ZngweUIwWmh6cXNQeUN4")

//     client := &http.Client{}
//     response, error := client.Do(request)

//     if error != nil {
//         fmt.Println(error)
//     }

//     responseBody, error := io.ReadAll(response.Body)

//     if error != nil {
//         fmt.Println(error)
//     }

//     formattedData := formatJSON(responseBody)
//     fmt.Println("Status: ", response.Status)
//     fmt.Println("Response body: ", formattedData)

//     // clean up memory after execution
//    defer response.Body.Close()
// }

// func (c *Controller) mkReq(ctx *gin.Context, method, url string, body io.Reader) (*http.Request, error) {
// 	r, err := http.NewRequestWithContext(ctx, method, url, body)
// 	if err != nil {
// 		return nil, err
// 	}

// 	r.SetBasicAuth("rahul.bhatiya@t-systems.com", "cmVmdGtuOjAxOjE3MzMzMDE2MTc6dWNCOGkxNlNrcDA0ZngweUIwWmh6cXNQeUN4")
// 	return r, nil
// }

// AddAccount godoc
//
//	@Summary		Add an account
//	@Description	add by json account
//	@Tags			accounts
//	@Accept			json
//	@Produce		json
//	@Param			account	body		model.AddAccount	true	"Add account"
//	@Success		200		{object}	model.Account
//	@Failure		400		{object}	httputil.HTTPError
//	@Failure		404		{object}	httputil.HTTPError
//	@Failure		500		{object}	httputil.HTTPError
//	@Router			/accounts [post]
// func (c *Controller) AddAccount(ctx *gin.Context) {
// 	var addAccount model.AddAccount
// 	if err := ctx.ShouldBindJSON(&addAccount); err != nil {
// 		httputil.NewError(ctx, http.StatusBadRequest, err)
// 		return
// 	}
// 	if err := addAccount.Validation(); err != nil {
// 		httputil.NewError(ctx, http.StatusBadRequest, err)
// 		return
// 	}
// 	account := model.Account{
// 		Name: addAccount.Name,
// 	}
// 	lastID, err := account.Insert()
// 	if err != nil {
// 		httputil.NewError(ctx, http.StatusBadRequest, err)
// 		return
// 	}
// 	account.ID = lastID
// 	ctx.JSON(http.StatusOK, account)
// }

// UpdateAccount godoc
//
//	@Summary		Update an account
//	@Description	Update by json account
//	@Tags			accounts
//	@Accept			json
//	@Produce		json
//	@Param			id		path		int					true	"Account ID"
//	@Param			account	body		model.UpdateAccount	true	"Update account"
//	@Success		200		{object}	model.Account
//	@Failure		400		{object}	httputil.HTTPError
//	@Failure		404		{object}	httputil.HTTPError
//	@Failure		500		{object}	httputil.HTTPError
//	@Router			/accounts/{id} [patch]
// func (c *Controller) UpdateAccount(ctx *gin.Context) {
// 	id := ctx.Param("id")
// 	aid, err := strconv.Atoi(id)
// 	if err != nil {
// 		httputil.NewError(ctx, http.StatusBadRequest, err)
// 		return
// 	}
// 	var updateAccount model.UpdateAccount
// 	if err := ctx.ShouldBindJSON(&updateAccount); err != nil {
// 		httputil.NewError(ctx, http.StatusBadRequest, err)
// 		return
// 	}
// 	account := model.Account{
// 		ID:   aid,
// 		Name: updateAccount.Name,
// 	}
// 	err = account.Update()
// 	if err != nil {
// 		httputil.NewError(ctx, http.StatusNotFound, err)
// 		return
// 	}
// 	ctx.JSON(http.StatusOK, account)
// }

// DeleteAccount godoc
//
//	@Summary		Delete an account
//	@Description	Delete by account ID
//	@Tags			accounts
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"Account ID"	Format(int64)
//	@Success		204	{object}	model.Account
//	@Failure		400	{object}	httputil.HTTPError
//	@Failure		404	{object}	httputil.HTTPError
//	@Failure		500	{object}	httputil.HTTPError
//	@Router			/accounts/{id} [delete]
// func (c *Controller) DeleteAccount(ctx *gin.Context) {
// 	id := ctx.Param("id")
// 	aid, err := strconv.Atoi(id)
// 	if err != nil {
// 		httputil.NewError(ctx, http.StatusBadRequest, err)
// 		return
// 	}
// 	err = model.Delete(aid)
// 	if err != nil {
// 		httputil.NewError(ctx, http.StatusNotFound, err)
// 		return
// 	}
// 	ctx.JSON(http.StatusNoContent, gin.H{})
// }

// UploadAccountImage godoc
//
//	@Summary		Upload account image
//	@Description	Upload file
//	@Tags			accounts
//	@Accept			multipart/form-data
//	@Produce		json
//	@Param			id		path		int		true	"Account ID"
//	@Param			file	formData	file	true	"account image"
//	@Success		200		{object}	controller.Message
//	@Failure		400		{object}	httputil.HTTPError
//	@Failure		404		{object}	httputil.HTTPError
//	@Failure		500		{object}	httputil.HTTPError
//	@Router			/accounts/{id}/images [post]
// func (c *Controller) UploadAccountImage(ctx *gin.Context) {
// 	id, err := strconv.Atoi(ctx.Param("id"))
// 	if err != nil {
// 		httputil.NewError(ctx, http.StatusBadRequest, err)
// 		return
// 	}
// 	file, err := ctx.FormFile("file")
// 	if err != nil {
// 		httputil.NewError(ctx, http.StatusBadRequest, err)
// 		return
// 	}
// 	ctx.JSON(http.StatusOK, Message{Message: fmt.Sprintf("upload complete userID=%d filename=%s", id, file.Filename)})
// }
