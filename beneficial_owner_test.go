package dwolla

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBeneficialOwnerServiceRetrieve(t *testing.T) {
	c := newMockClient(200, filepath.Join("testdata", "beneficial-owner.json"))
	res, err := c.BeneficialOwner.Retrieve("00cb67f2-768c-4ee3-ac81-73bc4faf9c2b")

	assert.Nil(t, err)
	assert.NotNil(t, res)
	assert.Equal(t, res.ID, "00cb67f2-768c-4ee3-ac81-73bc4faf9c2b")
}

func TestBeneficialOwnerServiceRetrieveError(t *testing.T) {
	c := newMockClient(404, filepath.Join("testdata", "resource-not-found.json"))
	res, err := c.BeneficialOwner.Retrieve("00cb67f2-768c-4ee3-ac81-73bc4faf9c2b")

	assert.Error(t, err)
	assert.Nil(t, res)
}

func TestBeneficialOwnerServiceRemove(t *testing.T) {
	c := newMockClient(200, filepath.Join("testdata", "beneficial-owner.json"))
	err := c.BeneficialOwner.Remove("00cb67f2-768c-4ee3-ac81-73bc4faf9c2b")

	assert.Nil(t, err)
}

func TestBeneficialOwnerServiceUpdate(t *testing.T) {
	c := newMockClient(200, filepath.Join("testdata", "beneficial-owner.json"))
	res, err := c.BeneficialOwner.Update("00cb67f2-768c-4ee3-ac81-73bc4faf9c2b", &BeneficialOwnerRequest{
		FirstName: "John",
		LastName:  "Doe",
	})

	assert.Nil(t, err)
	assert.NotNil(t, res)
	assert.Equal(t, res.ID, "00cb67f2-768c-4ee3-ac81-73bc4faf9c2b")
}

func TestBeneficialOwnerServiceUpdateError(t *testing.T) {
	c := newMockClient(404, filepath.Join("testdata", "resource-not-found.json"))
	res, err := c.BeneficialOwner.Update("00cb67f2-768c-4ee3-ac81-73bc4faf9c2b", &BeneficialOwnerRequest{
		FirstName: "John",
		LastName:  "Doe",
	})

	assert.Error(t, err)
	assert.Nil(t, res)
}

func TestBeneficialOwnerCreateDocument(t *testing.T) {
	c := newMockClient(200, filepath.Join("testdata", "document.json"))

	owner := &BeneficialOwner{Resource: Resource{client: c, Links: Links{"self": Link{Href: "https://api-sandbox.dwolla.com/beneficial-owners/07d59716-ef22-4fe6-98e8-f3190233dfb8"}}}}

	f, _ := os.Open(filepath.Join("testdata", "document-upload-success.png"))
	res, err := owner.CreateDocument(&DocumentRequest{
		Type:     DocumentTypePassport,
		FileName: f.Name(),
		File:     f,
	})

	assert.Nil(t, err)
	assert.NotNil(t, res)
}

func TestBeneficialOwnerCreateDocumentError(t *testing.T) {
	c := newMockClient(404, filepath.Join("testdata", "resource-not-found.json"))

	owner := &BeneficialOwner{Resource: Resource{client: c, Links: Links{"self": Link{Href: "https://api-sandbox.dwolla.com/beneficial-owners/07d59716-ef22-4fe6-98e8-f3190233dfb8"}}}}
	f, _ := os.Open(filepath.Join("testdata", "document-upload-success.png"))
	res, err := owner.CreateDocument(&DocumentRequest{
		Type:     DocumentTypePassport,
		FileName: f.Name(),
		File:     f,
	})

	assert.Error(t, err)
	assert.Nil(t, res)
}