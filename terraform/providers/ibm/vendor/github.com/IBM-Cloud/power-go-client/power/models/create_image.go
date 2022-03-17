// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// CreateImage create image
// swagger:model CreateImage
type CreateImage struct {

	// Cloud Storage access key; required for import image
	AccessKey string `json:"accessKey,omitempty"`

	// Cloud Storage bucket name; bucket-name[/optional/folder]; required for import image
	BucketName string `json:"bucketName,omitempty"`

	// Type of Disk; will be ignored if storagePool or affinityPolicy is provided; Used only when importing an image from cloud storage.
	DiskType string `json:"diskType,omitempty"`

	// Cloud Storage image filename; required for import image
	ImageFilename string `json:"imageFilename,omitempty"`

	// Image ID of existing source image; required for copy image
	ImageID string `json:"imageID,omitempty"`

	// Name to give created image; required for import image
	ImageName string `json:"imageName,omitempty"`

	// (deprecated - replaced by region, imageFilename and bucketName) Path to image starting with service endpoint and ending with image filename
	ImagePath string `json:"imagePath,omitempty"`

	// Image OS Type, required if importing a raw image; raw images can only be imported using the command line interface
	// Enum: [aix ibmi rhel sles]
	OsType string `json:"osType,omitempty"`

	// Cloud Storage Region; only required to access IBM Cloud Storage
	Region string `json:"region,omitempty"`

	// Cloud Storage secret key; required for import image
	SecretKey string `json:"secretKey,omitempty"`

	// Source of the image
	// Required: true
	// Enum: [root-project url]
	Source *string `json:"source"`

	// The storage affinity data; ignored if storagePool is provided; Used only when importing an image from cloud storage.
	StorageAffinity *StorageAffinity `json:"storageAffinity,omitempty"`

	// Storage pool where the image will be loaded; if provided then storageAffinity and diskType will be ignored; Used only when importing an image from cloud storage.
	StoragePool string `json:"storagePool,omitempty"`
}

// Validate validates this create image
func (m *CreateImage) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateOsType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSource(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStorageAffinity(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var createImageTypeOsTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["aix","ibmi","rhel","sles"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		createImageTypeOsTypePropEnum = append(createImageTypeOsTypePropEnum, v)
	}
}

const (

	// CreateImageOsTypeAix captures enum value "aix"
	CreateImageOsTypeAix string = "aix"

	// CreateImageOsTypeIbmi captures enum value "ibmi"
	CreateImageOsTypeIbmi string = "ibmi"

	// CreateImageOsTypeRhel captures enum value "rhel"
	CreateImageOsTypeRhel string = "rhel"

	// CreateImageOsTypeSles captures enum value "sles"
	CreateImageOsTypeSles string = "sles"
)

// prop value enum
func (m *CreateImage) validateOsTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, createImageTypeOsTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *CreateImage) validateOsType(formats strfmt.Registry) error {

	if swag.IsZero(m.OsType) { // not required
		return nil
	}

	// value enum
	if err := m.validateOsTypeEnum("osType", "body", m.OsType); err != nil {
		return err
	}

	return nil
}

var createImageTypeSourcePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["root-project","url"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		createImageTypeSourcePropEnum = append(createImageTypeSourcePropEnum, v)
	}
}

const (

	// CreateImageSourceRootProject captures enum value "root-project"
	CreateImageSourceRootProject string = "root-project"

	// CreateImageSourceURL captures enum value "url"
	CreateImageSourceURL string = "url"
)

// prop value enum
func (m *CreateImage) validateSourceEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, createImageTypeSourcePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *CreateImage) validateSource(formats strfmt.Registry) error {

	if err := validate.Required("source", "body", m.Source); err != nil {
		return err
	}

	// value enum
	if err := m.validateSourceEnum("source", "body", *m.Source); err != nil {
		return err
	}

	return nil
}

func (m *CreateImage) validateStorageAffinity(formats strfmt.Registry) error {

	if swag.IsZero(m.StorageAffinity) { // not required
		return nil
	}

	if m.StorageAffinity != nil {
		if err := m.StorageAffinity.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("storageAffinity")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CreateImage) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreateImage) UnmarshalBinary(b []byte) error {
	var res CreateImage
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}