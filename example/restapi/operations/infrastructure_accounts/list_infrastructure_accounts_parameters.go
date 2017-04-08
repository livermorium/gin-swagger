package infrastructure_accounts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/gin-gonic/gin"
	"github.com/go-openapi/errors"
)

// ListInfrastructureAccountsParams contains all the bound params for the list infrastructure accounts operation
// typically these are obtained from a http.Request
//
// swagger:parameters listInfrastructureAccounts
type ListInfrastructureAccountsParams struct {
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls
func (o *ListInfrastructureAccountsParams) bindRequest(ctx *gin.Context) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
