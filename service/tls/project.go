package tls

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"
)

func (c *LsClient) CreateProject(request *CreateProjectRequest) (r *CreateProjectResponse, e error) {
	reqHeaders := map[string]string{
		"Content-Type": "application/json",
	}

	reqBody := map[string]string{
		"ProjectName": request.ProjectName,
		"Description": request.Description,
		"Region":      request.Region,
	}

	bytesBody, err := json.Marshal(reqBody)
	if err != nil {
		return nil, err
	}

	rawResponse, err := c.Request(http.MethodPost, PathCreateProject, nil, c.assembleHeader(request.CommonRequest, reqHeaders), bytesBody)
	if err != nil {
		return nil, err
	}
	defer rawResponse.Body.Close()

	responseBody, err := ioutil.ReadAll(rawResponse.Body)
	if err != nil {
		return nil, err
	}

	var response = &CreateProjectResponse{}
	response.FillRequestId(rawResponse)
	if err := json.Unmarshal(responseBody, response); err != nil {
		return nil, err
	}

	return response, nil
}

func (c *LsClient) DeleteProject(request *DeleteProjectRequest) (r *CommonResponse, e error) {
	reqHeaders := map[string]string{
		"Content-Type": "application/json",
	}

	reqBody := map[string]string{
		"ProjectId": request.ProjectID,
	}

	jsonBody, err := json.Marshal(reqBody)
	if err != nil {
		return nil, err
	}

	rawResponse, err := c.Request(http.MethodDelete, PathDeleteProject, nil, c.assembleHeader(request.CommonRequest, reqHeaders), jsonBody)
	if err != nil {
		return nil, err
	}
	defer rawResponse.Body.Close()

	if rawResponse.StatusCode != http.StatusOK {
		errorMessage, err := ioutil.ReadAll(rawResponse.Body)
		if err != nil {
			return nil, err
		}
		return nil, errors.New(string(errorMessage))
	}

	response := &CommonResponse{}
	response.FillRequestId(rawResponse)

	return response, nil
}

func (c *LsClient) DescribeProject(request *DescribeProjectRequest) (r *DescribeProjectResponse, e error) {
	reqHeaders := map[string]string{
		"Content-Type": "application/json",
	}

	params := map[string]string{
		"ProjectId": request.ProjectID,
	}

	body := map[string]string{}
	bytesBody, err := json.Marshal(body)

	rawResponse, err := c.Request(http.MethodGet, PathDescribeProject, params, c.assembleHeader(request.CommonRequest, reqHeaders), bytesBody)
	if err != nil {
		return nil, err
	}
	defer rawResponse.Body.Close()

	responseBody, err := ioutil.ReadAll(rawResponse.Body)
	if err != nil {
		return nil, err
	}

	var response = &DescribeProjectResponse{}
	response.FillRequestId(rawResponse)

	if err := json.Unmarshal(responseBody, response); err != nil {
		return nil, err
	}

	return response, nil
}

func (c *LsClient) DescribeProjects(request *DescribeProjectsRequest) (r *DescribeProjectsResponse, e error) {
	reqHeaders := map[string]string{
		"Content-Type": "application/json",
	}

	params := make(map[string]string)
	if len(request.ProjectName) != 0 {
		params["ProjectName"] = request.ProjectName
	}

	if len(request.ProjectID) != 0 {
		params["ProjectId"] = request.ProjectID
	}

	if request.PageNumber != 0 {
		params["PageNumber"] = strconv.Itoa(request.PageNumber)
	}

	if request.PageSize != 0 {
		params["PageSize"] = strconv.Itoa(request.PageSize)
	}

	params["IsFullName"] = strconv.FormatBool(request.IsFullName)

	body := map[string]string{}
	bytesBody, err := json.Marshal(body)

	rawResponse, err := c.Request(http.MethodGet, PathDescribeProjects, params, c.assembleHeader(request.CommonRequest, reqHeaders), bytesBody)
	if err != nil {
		return nil, err
	}
	defer rawResponse.Body.Close()

	responseBody, err := ioutil.ReadAll(rawResponse.Body)
	if err != nil {
		return nil, err
	}

	var response = &DescribeProjectsResponse{}
	response.FillRequestId(rawResponse)

	if err := json.Unmarshal(responseBody, response); err != nil {
		return nil, err
	}

	return response, nil
}

func (c *LsClient) ModifyProject(request *ModifyProjectRequest) (r *CommonResponse, e error) {
	reqHeaders := map[string]string{
		"Content-Type": "application/json",
	}

	reqBody := map[string]string{
		"ProjectId": request.ProjectID,
	}

	if request.ProjectName != nil {
		reqBody["ProjectName"] = *request.ProjectName
	}

	if request.Description != nil {
		reqBody["Description"] = *request.Description
	}

	jsonBody, err := json.Marshal(reqBody)
	if err != nil {
		return nil, err
	}

	rawResponse, err := c.Request(http.MethodPut, PathModifyProject, nil, c.assembleHeader(request.CommonRequest, reqHeaders), jsonBody)
	if err != nil {
		return nil, err
	}
	defer rawResponse.Body.Close()

	if rawResponse.StatusCode != http.StatusOK {
		errorMessage, err := ioutil.ReadAll(rawResponse.Body)
		if err != nil {
			return nil, err
		}
		return nil, errors.New(string(errorMessage))
	}
	response := &CommonResponse{}
	response.FillRequestId(rawResponse)

	return response, nil
}
