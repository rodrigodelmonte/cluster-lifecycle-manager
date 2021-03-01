// Code generated by go-swagger; DO NOT EDIT.

package infrastructure_accounts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"
	"io"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/zalando-incubator/cluster-lifecycle-manager/pkg/cluster-registry/models"
)

// ListInfrastructureAccountsReader is a Reader for the ListInfrastructureAccounts structure.
type ListInfrastructureAccountsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListInfrastructureAccountsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListInfrastructureAccountsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewListInfrastructureAccountsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewListInfrastructureAccountsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewListInfrastructureAccountsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewListInfrastructureAccountsOK creates a ListInfrastructureAccountsOK with default headers values
func NewListInfrastructureAccountsOK() *ListInfrastructureAccountsOK {
	return &ListInfrastructureAccountsOK{}
}

/* ListInfrastructureAccountsOK describes a response with status code 200, with default header values.

List of all infrastructure accounts.
*/
type ListInfrastructureAccountsOK struct {
	Payload *ListInfrastructureAccountsOKBody
}

func (o *ListInfrastructureAccountsOK) Error() string {
	return fmt.Sprintf("[GET /infrastructure-accounts][%d] listInfrastructureAccountsOK  %+v", 200, o.Payload)
}
func (o *ListInfrastructureAccountsOK) GetPayload() *ListInfrastructureAccountsOKBody {
	return o.Payload
}

func (o *ListInfrastructureAccountsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ListInfrastructureAccountsOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListInfrastructureAccountsUnauthorized creates a ListInfrastructureAccountsUnauthorized with default headers values
func NewListInfrastructureAccountsUnauthorized() *ListInfrastructureAccountsUnauthorized {
	return &ListInfrastructureAccountsUnauthorized{}
}

/* ListInfrastructureAccountsUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type ListInfrastructureAccountsUnauthorized struct {
}

func (o *ListInfrastructureAccountsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /infrastructure-accounts][%d] listInfrastructureAccountsUnauthorized ", 401)
}

func (o *ListInfrastructureAccountsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewListInfrastructureAccountsForbidden creates a ListInfrastructureAccountsForbidden with default headers values
func NewListInfrastructureAccountsForbidden() *ListInfrastructureAccountsForbidden {
	return &ListInfrastructureAccountsForbidden{}
}

/* ListInfrastructureAccountsForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type ListInfrastructureAccountsForbidden struct {
}

func (o *ListInfrastructureAccountsForbidden) Error() string {
	return fmt.Sprintf("[GET /infrastructure-accounts][%d] listInfrastructureAccountsForbidden ", 403)
}

func (o *ListInfrastructureAccountsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewListInfrastructureAccountsInternalServerError creates a ListInfrastructureAccountsInternalServerError with default headers values
func NewListInfrastructureAccountsInternalServerError() *ListInfrastructureAccountsInternalServerError {
	return &ListInfrastructureAccountsInternalServerError{}
}

/* ListInfrastructureAccountsInternalServerError describes a response with status code 500, with default header values.

Unexpected error
*/
type ListInfrastructureAccountsInternalServerError struct {
	Payload *models.Error
}

func (o *ListInfrastructureAccountsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /infrastructure-accounts][%d] listInfrastructureAccountsInternalServerError  %+v", 500, o.Payload)
}
func (o *ListInfrastructureAccountsInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *ListInfrastructureAccountsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*ListInfrastructureAccountsOKBody list infrastructure accounts o k body
swagger:model ListInfrastructureAccountsOKBody
*/
type ListInfrastructureAccountsOKBody struct {

	// items
	Items []*models.InfrastructureAccount `json:"items"`
}

// Validate validates this list infrastructure accounts o k body
func (o *ListInfrastructureAccountsOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateItems(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ListInfrastructureAccountsOKBody) validateItems(formats strfmt.Registry) error {
	if swag.IsZero(o.Items) { // not required
		return nil
	}

	for i := 0; i < len(o.Items); i++ {
		if swag.IsZero(o.Items[i]) { // not required
			continue
		}

		if o.Items[i] != nil {
			if err := o.Items[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("listInfrastructureAccountsOK" + "." + "items" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this list infrastructure accounts o k body based on the context it is used
func (o *ListInfrastructureAccountsOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateItems(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ListInfrastructureAccountsOKBody) contextValidateItems(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Items); i++ {

		if o.Items[i] != nil {
			if err := o.Items[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("listInfrastructureAccountsOK" + "." + "items" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *ListInfrastructureAccountsOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ListInfrastructureAccountsOKBody) UnmarshalBinary(b []byte) error {
	var res ListInfrastructureAccountsOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
