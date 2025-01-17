//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

package azblob

import (
	"sort"

	"github.com/stretchr/testify/assert"
)

// tests general functionality
func (s *azblobTestSuite) TestBlobListWrapper() {
	_assert := assert.New(s.T())
	testName := s.T().Name()
	_context := getTestContext(testName)
	svcClient, err := getServiceClient(_context.recording, testAccountDefault, nil)
	if err != nil {
		s.Fail("Unable to fetch service client because " + err.Error())
	}

	containerName := generateContainerName(testName)
	containerClient, _ := getContainerClient(containerName, svcClient)

	_, err = containerClient.Create(ctx, nil)
	_assert.Nil(err)
	defer deleteContainer(_assert, containerClient)

	files := []string{"a123", "b234", "c345"}

	createNewBlobs(_assert, files, containerClient)

	pager := containerClient.ListBlobsFlat(nil)

	found := make([]string, 0)

	for pager.NextPage(ctx) {
		resp := pager.PageResponse()

		for _, blob := range resp.ContainerListBlobFlatSegmentResult.Segment.BlobItems {
			found = append(found, *blob.Name)
		}
	}
	_assert.Nil(pager.Err())

	sort.Strings(files)
	sort.Strings(found)

	_assert.EqualValues(found, files)
}

// tests that the buffer filling isn't a problem
func (s *azblobTestSuite) TestBlobListWrapperFullBuffer() {
	_assert := assert.New(s.T())
	testName := s.T().Name()
	_context := getTestContext(testName)
	svcClient, err := getServiceClient(_context.recording, testAccountDefault, nil)
	if err != nil {
		s.Fail("Unable to fetch service client because " + err.Error())
	}

	containerClient, _ := getContainerClient(generateContainerName(testName), svcClient)

	_, err = containerClient.Create(ctx, nil)
	_assert.Nil(err)
	defer deleteContainer(_assert, containerClient)

	files := []string{"a123", "b234", "c345"}

	createNewBlobs(_assert, files, containerClient)

	pager := containerClient.ListBlobsFlat(nil)

	found := make([]string, 0)

	for pager.NextPage(ctx) {
		resp := pager.PageResponse()

		for _, blob := range resp.ContainerListBlobFlatSegmentResult.Segment.BlobItems {
			found = append(found, *blob.Name)
		}
	}
	_assert.Nil(pager.Err())

	sort.Strings(files)
	sort.Strings(found)

	_assert.EqualValues(files, found)
}

func (s *azblobTestSuite) TestBlobListWrapperListingError() {
	_assert := assert.New(s.T())
	testName := s.T().Name()
	_context := getTestContext(testName)
	svcClient, err := getServiceClient(_context.recording, testAccountDefault, nil)
	if err != nil {
		s.Fail("Unable to fetch service client because " + err.Error())
	}

	containerClient, _ := getContainerClient(generateContainerName(testName), svcClient)

	pager := containerClient.ListBlobsFlat(nil)

	_assert.Equal(pager.NextPage(ctx), false)
	_assert.NotNil(pager.Err())
}
