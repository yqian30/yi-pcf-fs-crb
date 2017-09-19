package crb_web

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"

	strfmt "github.com/go-openapi/strfmt"
)

// NewDeleteCopyParams creates a new DeleteCopyParams object
// with the default values initialized.
func NewDeleteCopyParams() DeleteCopyParams {
	var ()
	return DeleteCopyParams{}
}

// DeleteCopyParams contains all the bound params for the delete copy operation
// typically these are obtained from a http.Request
//
// swagger:parameters deleteCopy
type DeleteCopyParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request

	/*The id of the copy instance.
	  Required: true
	  In: path
	*/
	CopyID string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls
func (o *DeleteCopyParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error
	o.HTTPRequest = r

	rCopyID, rhkCopyID, _ := route.Params.GetOK("copyId")
	if err := o.bindCopyID(rCopyID, rhkCopyID, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *DeleteCopyParams) bindCopyID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	o.CopyID = raw

	return nil
}