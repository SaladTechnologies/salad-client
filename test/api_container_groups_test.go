/*
SaladCloud API

Testing ContainerGroupsAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package saladclient

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/SaladTechnologies/salad-client"
)

func Test_saladclient_ContainerGroupsAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ContainerGroupsAPIService CreateContainerGroup", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var organizationName string
		var projectName string

		resp, httpRes, err := apiClient.ContainerGroupsAPI.CreateContainerGroup(context.Background(), organizationName, projectName).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ContainerGroupsAPIService DeleteContainerGroup", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var organizationName string
		var projectName string
		var containerGroupName string

		httpRes, err := apiClient.ContainerGroupsAPI.DeleteContainerGroup(context.Background(), organizationName, projectName, containerGroupName).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ContainerGroupsAPIService GetContainerGroup", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var organizationName string
		var projectName string
		var containerGroupName string

		resp, httpRes, err := apiClient.ContainerGroupsAPI.GetContainerGroup(context.Background(), organizationName, projectName, containerGroupName).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ContainerGroupsAPIService GetContainerGroupInstance", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var organizationName string
		var projectName string
		var containerGroupName string
		var containerGroupInstanceId string

		resp, httpRes, err := apiClient.ContainerGroupsAPI.GetContainerGroupInstance(context.Background(), organizationName, projectName, containerGroupName, containerGroupInstanceId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ContainerGroupsAPIService ListContainerGroupInstances", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var organizationName string
		var projectName string
		var containerGroupName string

		resp, httpRes, err := apiClient.ContainerGroupsAPI.ListContainerGroupInstances(context.Background(), organizationName, projectName, containerGroupName).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ContainerGroupsAPIService ListContainerGroups", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var organizationName string
		var projectName string

		resp, httpRes, err := apiClient.ContainerGroupsAPI.ListContainerGroups(context.Background(), organizationName, projectName).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ContainerGroupsAPIService ReallocateContainerGroupInstance", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var organizationName string
		var projectName string
		var containerGroupName string
		var containerGroupInstanceId string

		httpRes, err := apiClient.ContainerGroupsAPI.ReallocateContainerGroupInstance(context.Background(), organizationName, projectName, containerGroupName, containerGroupInstanceId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ContainerGroupsAPIService RecreateContainerGroupInstance", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var organizationName string
		var projectName string
		var containerGroupName string
		var containerGroupInstanceId string

		httpRes, err := apiClient.ContainerGroupsAPI.RecreateContainerGroupInstance(context.Background(), organizationName, projectName, containerGroupName, containerGroupInstanceId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ContainerGroupsAPIService RestartContainerGroupInstance", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var organizationName string
		var projectName string
		var containerGroupName string
		var containerGroupInstanceId string

		httpRes, err := apiClient.ContainerGroupsAPI.RestartContainerGroupInstance(context.Background(), organizationName, projectName, containerGroupName, containerGroupInstanceId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ContainerGroupsAPIService StartContainerGroup", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var organizationName string
		var projectName string
		var containerGroupName string

		httpRes, err := apiClient.ContainerGroupsAPI.StartContainerGroup(context.Background(), organizationName, projectName, containerGroupName).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ContainerGroupsAPIService StopContainerGroup", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var organizationName string
		var projectName string
		var containerGroupName string

		httpRes, err := apiClient.ContainerGroupsAPI.StopContainerGroup(context.Background(), organizationName, projectName, containerGroupName).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ContainerGroupsAPIService UpdateContainerGroup", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var organizationName string
		var projectName string
		var containerGroupName string

		resp, httpRes, err := apiClient.ContainerGroupsAPI.UpdateContainerGroup(context.Background(), organizationName, projectName, containerGroupName).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
