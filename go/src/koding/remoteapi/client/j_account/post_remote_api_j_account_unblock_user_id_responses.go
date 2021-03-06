package j_account

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"

	"koding/remoteapi/models"
)

// PostRemoteAPIJAccountUnblockUserIDReader is a Reader for the PostRemoteAPIJAccountUnblockUserID structure.
type PostRemoteAPIJAccountUnblockUserIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostRemoteAPIJAccountUnblockUserIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPostRemoteAPIJAccountUnblockUserIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostRemoteAPIJAccountUnblockUserIDOK creates a PostRemoteAPIJAccountUnblockUserIDOK with default headers values
func NewPostRemoteAPIJAccountUnblockUserIDOK() *PostRemoteAPIJAccountUnblockUserIDOK {
	return &PostRemoteAPIJAccountUnblockUserIDOK{}
}

/*PostRemoteAPIJAccountUnblockUserIDOK handles this case with default header values.

OK
*/
type PostRemoteAPIJAccountUnblockUserIDOK struct {
	Payload PostRemoteAPIJAccountUnblockUserIDOKBody
}

func (o *PostRemoteAPIJAccountUnblockUserIDOK) Error() string {
	return fmt.Sprintf("[POST /remote.api/JAccount.unblockUser/{id}][%d] postRemoteApiJAccountUnblockUserIdOK  %+v", 200, o.Payload)
}

func (o *PostRemoteAPIJAccountUnblockUserIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*PostRemoteAPIJAccountUnblockUserIDOKBody post remote API j account unblock user ID o k body
swagger:model PostRemoteAPIJAccountUnblockUserIDOKBody
*/
type PostRemoteAPIJAccountUnblockUserIDOKBody struct {
	models.JAccount

	models.DefaultResponse
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *PostRemoteAPIJAccountUnblockUserIDOKBody) UnmarshalJSON(raw []byte) error {

	var postRemoteAPIJAccountUnblockUserIDOKBodyAO0 models.JAccount
	if err := swag.ReadJSON(raw, &postRemoteAPIJAccountUnblockUserIDOKBodyAO0); err != nil {
		return err
	}
	o.JAccount = postRemoteAPIJAccountUnblockUserIDOKBodyAO0

	var postRemoteAPIJAccountUnblockUserIDOKBodyAO1 models.DefaultResponse
	if err := swag.ReadJSON(raw, &postRemoteAPIJAccountUnblockUserIDOKBodyAO1); err != nil {
		return err
	}
	o.DefaultResponse = postRemoteAPIJAccountUnblockUserIDOKBodyAO1

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o PostRemoteAPIJAccountUnblockUserIDOKBody) MarshalJSON() ([]byte, error) {
	var _parts [][]byte

	postRemoteAPIJAccountUnblockUserIDOKBodyAO0, err := swag.WriteJSON(o.JAccount)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, postRemoteAPIJAccountUnblockUserIDOKBodyAO0)

	postRemoteAPIJAccountUnblockUserIDOKBodyAO1, err := swag.WriteJSON(o.DefaultResponse)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, postRemoteAPIJAccountUnblockUserIDOKBodyAO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this post remote API j account unblock user ID o k body
func (o *PostRemoteAPIJAccountUnblockUserIDOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.JAccount.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := o.DefaultResponse.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
