// Code generated by goagen v1.3.0, DO NOT EDIT.
//
// API "Slash API": Application Media Types
//
// Command:
// $ goagen
// --design=github.com/maurodelazeri/slash-auth/design
// --out=$(GOPATH)/src/github.com/maurodelazeri/slash-auth
// --version=v1.3.0

package client

import (
	"github.com/goadesign/goa"
	"net/http"
)

// A user account (default view)
//
// Identifier: application/ft.account+json; view=default
type FtAccount struct {
	// Email of the user
	Email string `form:"email" json:"email" xml:"email"`
	// First name of the user
	FirstName string `form:"first_name" json:"first_name" xml:"first_name"`
	// ID of the user
	ID int `form:"id" json:"id" xml:"id"`
	// Last name of the user
	LastName string `form:"last_name" json:"last_name" xml:"last_name"`
}

// Validate validates the FtAccount media type instance.
func (mt *FtAccount) Validate() (err error) {

	if mt.FirstName == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "first_name"))
	}
	if mt.LastName == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "last_name"))
	}
	if mt.Email == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "email"))
	}
	return
}

// DecodeFtAccount decodes the FtAccount instance encoded in resp body.
func (c *Client) DecodeFtAccount(resp *http.Response) (*FtAccount, error) {
	var decoded FtAccount
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// FtAccountCollection is the media type for an array of FtAccount (default view)
//
// Identifier: application/ft.account+json; type=collection; view=default
type FtAccountCollection []*FtAccount

// Validate validates the FtAccountCollection media type instance.
func (mt FtAccountCollection) Validate() (err error) {
	for _, e := range mt {
		if e != nil {
			if err2 := e.Validate(); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// DecodeFtAccountCollection decodes the FtAccountCollection instance encoded in resp body.
func (c *Client) DecodeFtAccountCollection(resp *http.Response) (FtAccountCollection, error) {
	var decoded FtAccountCollection
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return decoded, err
}

// A JWT for a user (default view)
//
// Identifier: application/ft.token+json; view=default
type FtToken struct {
	// A JWT token
	Token string `form:"token" json:"token" xml:"token"`
}

// Validate validates the FtToken media type instance.
func (mt *FtToken) Validate() (err error) {
	if mt.Token == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "token"))
	}
	return
}

// DecodeFtToken decodes the FtToken instance encoded in resp body.
func (c *Client) DecodeFtToken(resp *http.Response) (*FtToken, error) {
	var decoded FtToken
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// DecodeErrorResponse decodes the ErrorResponse instance encoded in resp body.
func (c *Client) DecodeErrorResponse(resp *http.Response) (*goa.ErrorResponse, error) {
	var decoded goa.ErrorResponse
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}
