package service

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"mcx002/filestoreService/src"
	"net/url"
	"testing"

	"github.com/stretchr/testify/mock"
)

type RepositoryMock struct {
	mock.Mock
}

func (r *RepositoryMock) GetMinioPresignedPutUrl(ctx context.Context, objName string) (*url.URL, error) {
	args := r.Called(ctx)
	return args.Get(0).(*url.URL), args.Error(1)
}

func (r *RepositoryMock) GetMinioPresignedGetUrl(ctx context.Context, objName string) (*url.URL, error) {
	args := r.Called(ctx, objName)
	return args.Get(0).(*url.URL), args.Error(1)
}

func (r *RepositoryMock) GetCacheImageUrl(ctx context.Context, filename string) (string, error) {
	args := r.Called(ctx, filename)
	return args.String(0), args.Error(1)
}

func (r *RepositoryMock) SetCacheImageUrl(ctx context.Context, filename string, url string) error {
	args := r.Called(ctx, filename)
	return args.Error(0)
}

func TestGetUploadUrl(t *testing.T) {
	repoMock := &RepositoryMock{}
	svc := NewService(&src.AppConfig{}, repoMock)
	ctx := context.Background()
	expectedUrl := &url.URL{Host: "0.0.0.0"}

	repoMock.On("GetMinioPresignedPutUrl", ctx).Return(&url.URL{}, fmt.Errorf("error"))
	_, _, err := svc.GetUploadUrl(ctx)
	if err == nil {
		t.Errorf("got err nil; want err not nil")
	}

	ctx = context.WithValue(ctx, "secondTest", "1")
	repoMock.On("GetMinioPresignedPutUrl", ctx).Return(expectedUrl, nil)
	_, requestedUrl, _ := svc.GetUploadUrl(ctx)
	if expectedUrl.String() != requestedUrl {
		t.Errorf("got url %s; want url %s", requestedUrl, expectedUrl.String())
	}
}

func TestGetImageUrl(t *testing.T) {
	log.SetOutput(ioutil.Discard)

	repoMock := &RepositoryMock{}
	svc := NewService(&src.AppConfig{}, repoMock)
	ctx := context.Background()
	expectedUrl := &url.URL{Host: "0.0.0.0"}

	expectedUrlString := "http://success.com"
	repoMock.On("GetCacheImageUrl", ctx, "file1").Return(expectedUrlString, nil)
	val, _ := svc.GetImageUrl(ctx, "file1")
	if val != expectedUrlString {
		t.Errorf("got url %s; want url %s", val, expectedUrlString)
	}

	repoMock.On("GetCacheImageUrl", ctx, "file2").Return("", fmt.Errorf("no cache"))
	repoMock.On("GetMinioPresignedGetUrl", ctx, "file2").Return(&url.URL{}, fmt.Errorf("no image found"))
	_, err := svc.GetImageUrl(ctx, "file2")
	if err == nil {
		t.Errorf("got err nil; want err not nil")
	}

	repoMock.On("GetCacheImageUrl", ctx, "file3").Return("", fmt.Errorf("no cache"))
	repoMock.On("GetMinioPresignedGetUrl", ctx, "file3").Return(expectedUrl, nil)
	repoMock.On("SetCacheImageUrl", ctx, "file3").Return(fmt.Errorf("failed to set the cache"))
	val, _ = svc.GetImageUrl(ctx, "file3")
	if val != expectedUrl.String() {
		t.Errorf("got val = %s; want val = %s", val, expectedUrl.String())
	}

	repoMock.On("GetCacheImageUrl", ctx, "file4").Return("", fmt.Errorf("no cache"))
	repoMock.On("GetMinioPresignedGetUrl", ctx, "file4").Return(expectedUrl, nil)
	repoMock.On("SetCacheImageUrl", ctx, "file4").Return(nil)
	val, _ = svc.GetImageUrl(ctx, "file4")
	if val != expectedUrl.String() {
		t.Errorf("got val = %s; want val = %s", val, expectedUrl.String())
	}
}
