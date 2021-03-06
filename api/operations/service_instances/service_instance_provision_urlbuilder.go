// Code generated by go-swagger; DO NOT EDIT.

package service_instances

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"errors"
	"net/url"
	golangswaggerpaths "path"
	"strings"

	"github.com/go-openapi/swag"
)

// ServiceInstanceProvisionURL generates an URL for the service instance provision operation
type ServiceInstanceProvisionURL struct {
	InstanceID string

	AcceptsIncomplete *bool

	_basePath string
	// avoid unkeyed usage
	_ struct{}
}

// WithBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *ServiceInstanceProvisionURL) WithBasePath(bp string) *ServiceInstanceProvisionURL {
	o.SetBasePath(bp)
	return o
}

// SetBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *ServiceInstanceProvisionURL) SetBasePath(bp string) {
	o._basePath = bp
}

// Build a url path and query string
func (o *ServiceInstanceProvisionURL) Build() (*url.URL, error) {
	var _result url.URL

	var _path = "/v2/service_instances/{instance_id}"

	instanceID := o.InstanceID
	if instanceID != "" {
		_path = strings.Replace(_path, "{instance_id}", instanceID, -1)
	} else {
		return nil, errors.New("instanceId is required on ServiceInstanceProvisionURL")
	}

	_basePath := o._basePath
	_result.Path = golangswaggerpaths.Join(_basePath, _path)

	qs := make(url.Values)

	var acceptsIncompleteQ string
	if o.AcceptsIncomplete != nil {
		acceptsIncompleteQ = swag.FormatBool(*o.AcceptsIncomplete)
	}
	if acceptsIncompleteQ != "" {
		qs.Set("accepts_incomplete", acceptsIncompleteQ)
	}

	_result.RawQuery = qs.Encode()

	return &_result, nil
}

// Must is a helper function to panic when the url builder returns an error
func (o *ServiceInstanceProvisionURL) Must(u *url.URL, err error) *url.URL {
	if err != nil {
		panic(err)
	}
	if u == nil {
		panic("url can't be nil")
	}
	return u
}

// String returns the string representation of the path with query string
func (o *ServiceInstanceProvisionURL) String() string {
	return o.Must(o.Build()).String()
}

// BuildFull builds a full url with scheme, host, path and query string
func (o *ServiceInstanceProvisionURL) BuildFull(scheme, host string) (*url.URL, error) {
	if scheme == "" {
		return nil, errors.New("scheme is required for a full url on ServiceInstanceProvisionURL")
	}
	if host == "" {
		return nil, errors.New("host is required for a full url on ServiceInstanceProvisionURL")
	}

	base, err := o.Build()
	if err != nil {
		return nil, err
	}

	base.Scheme = scheme
	base.Host = host
	return base, nil
}

// StringFull returns the string representation of a complete url
func (o *ServiceInstanceProvisionURL) StringFull(scheme, host string) string {
	return o.Must(o.BuildFull(scheme, host)).String()
}
