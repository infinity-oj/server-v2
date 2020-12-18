package api

import (
	"fmt"

	"github.com/go-resty/resty/v2"

	"github.com/infinity-oj/server-v2/pkg/models"
)

type AccountService interface {
	Create(username, password, email string) (*models.Account, error)
	Login(username, password string) error
	Test() (*models.Account, error)
}

type accountService struct {
	client *resty.Client
}

func (s *accountService) Create(username, password, email string) (*models.Account, error) {

	account := &models.Account{}

	request := map[string]interface{}{
		"username": username,
		"password": password,
		"email":    email,
	}

	_, err := s.client.R().
		SetBody(request).
		SetResult(account).
		Post("/account/application")
	if err != nil {
		return nil, err
	}

	// Explore response object
	//fmt.Println("Response Info:")
	//fmt.Println("  ", resp.Request.URL)
	//fmt.Println("  Error      :", err)
	//fmt.Println("  Status Code:", resp.StatusCode())
	//fmt.Println("  Status     :", resp.Status())
	//fmt.Println("  Proto      :", resp.Proto())
	//fmt.Println("  Time       :", resp.Time())
	//fmt.Println("  Received At:", resp.ReceivedAt())
	//fmt.Println("  Body       :\n", resp)
	//fmt.Println()

	return account, nil
}

func (s *accountService) Login(username, password string) error {

	request := map[string]interface{}{
		"username": username,
		"password": password,
	}

	_, err := s.client.R().
		SetBody(request).
		Post("/session/principal")
	if err != nil {
		return err
	}

	// Explore response object
	//fmt.Println("Response Info:")
	//fmt.Println("  ", resp.Request.URL)
	//fmt.Println("  Error      :", err)
	//fmt.Println("  Status Code:", resp.StatusCode())
	//fmt.Println("  Status     :", resp.Status())
	//fmt.Println("  Proto      :", resp.Proto())
	//fmt.Println("  Time       :", resp.Time())
	//fmt.Println("  Received At:", resp.ReceivedAt())
	//fmt.Println("  Body       :\n", resp)
	//fmt.Println()
	err = Jar.Save()
	if err != nil {
		fmt.Println(err)
	}
	return err
}

func (s *accountService) Test() (*models.Account, error) {

	account := &models.Account{}

	_, err := s.client.R().
		SetResult(account).
		Get("/session/principal")
	if err != nil {
		return nil, err
	}

	// Explore response object
	//fmt.Println("Response Info:")
	//fmt.Println("  ", resp.Request.URL)
	//fmt.Println("  Error      :", err)
	//fmt.Println("  Status Code:", resp.StatusCode())
	//fmt.Println("  Status     :", resp.Status())
	//fmt.Println("  Proto      :", resp.Proto())
	//fmt.Println("  Time       :", resp.Time())
	//fmt.Println("  Received At:", resp.ReceivedAt())
	//fmt.Println("  Body       :\n", resp)
	//fmt.Println()
	return account, err
}

func NewAccountService(client *resty.Client) AccountService {
	return &accountService{
		client: client,
	}
}