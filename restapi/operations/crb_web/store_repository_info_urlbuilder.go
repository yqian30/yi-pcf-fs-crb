package crb_web

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"errors"
	"net/url"
	golangswaggerpaths "path"
)

// StoreRepositoryInfoURL generates an URL for the store repository info operation
type StoreRepositoryInfoURL struct {
	_basePath string
}

// WithBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *StoreRepositoryInfoURL) WithBasePath(bp string) *StoreRepositoryInfoURL {
	o.SetBasePath(bp)
	return o
}

// SetBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *StoreRepositoryInfoURL) SetBasePath(bp string) {
	o._basePath = bp
}

// Build a url path and query string
func (o *StoreRepositoryInfoURL) Build() (*url.URL, error) {
	var result url.URL

	var _path = "/repositories"

	_basePath := o._basePath
	if _basePath == "" {
		_basePath = "/crb"
	}
	result.Path = golangswaggerpaths.Join(_basePath, _path)

	return &result, nil
}

// Must is a helper function to panic when the url builder returns an error
func (o *StoreRepositoryInfoURL) Must(u *url.URL, err error) *url.URL {
	if err != nil {
		panic(err)
	}
	if u == nil {
		panic("url can't be nil")
	}
	return u
}

// String returns the string representation of the path with query string
func (o *StoreRepositoryInfoURL) String() string {
	return o.Must(o.Build()).String()
}

// BuildFull builds a full url with scheme, host, path and query string
func (o *StoreRepositoryInfoURL) BuildFull(scheme, host string) (*url.URL, error) {
	if scheme == "" {
		return nil, errors.New("scheme is required for a full url on StoreRepositoryInfoURL")
	}
	if host == "" {
		return nil, errors.New("host is required for a full url on StoreRepositoryInfoURL")
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
func (o *StoreRepositoryInfoURL) StringFull(scheme, host string) string {
	return o.Must(o.BuildFull(scheme, host)).String()
}
