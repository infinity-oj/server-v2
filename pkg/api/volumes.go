package api

import (
	"bytes"
	"fmt"

	"github.com/go-resty/resty/v2"

	"github.com/infinity-oj/server-v2/pkg/models"
)

type VolumeAPI interface {
	CreateVolume() (*models.Volume, error)
	CreateDirectory(volumeName, directory string) (*models.Volume, error)
	CreateFile(volumeName, filename string, file []byte) (*models.Volume, error)
	DownloadVolume(volumeName, directory string) ([]byte, error)
}

type volumeAPI struct {
	client *resty.Client
}

func (a *volumeAPI) DownloadVolume(volumeName, directory string) ([]byte, error) {
	resp, err := a.client.R().
		SetPathParams(map[string]string{
			"volumeName": volumeName,
		}).
		SetQueryParam("dirname", directory).
		Get("/volume/{volumeName}/download")
	if err != nil {
		return nil, err
	}
	return resp.Body(), nil
}

func (a *volumeAPI) CreateDirectory(volumeName, dirname string) (*models.Volume, error) {
	volume := &models.Volume{}

	_, err := a.client.R().
		SetBody(map[string]string{
			"dirname": dirname,
		}).
		SetResult(volume).
		Post(fmt.Sprintf("/volume/%s/directory", volumeName))

	if err != nil {
		return nil, err
	}

	return volume, nil
}

func (a *volumeAPI) CreateFile(volumeName, filename string, file []byte) (*models.Volume, error) {
	volume := &models.Volume{}

	_, err := a.client.R().
		SetFileReader(
			"file", filename, bytes.NewReader(file)).
		SetResult(volume).
		Post(fmt.Sprintf("/volume/%s/file", volumeName))

	if err != nil {
		return nil, err
	}

	return volume, nil
}

func (a *volumeAPI) CreateVolume() (*models.Volume, error) {
	volume := &models.Volume{}

	_, err := a.client.R().
		SetResult(volume).
		Post("/volume")
	if err != nil {
		return nil, err
	}

	return volume, nil
}

func NewVolumeAPI(client *resty.Client) VolumeAPI {
	return &volumeAPI{
		client: client,
	}
}
