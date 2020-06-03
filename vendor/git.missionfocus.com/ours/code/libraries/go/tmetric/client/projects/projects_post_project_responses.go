// Code generated by go-swagger; DO NOT EDIT.

package projects

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"git.missionfocus.com/ours/code/libraries/go/tmetric/models"
)

// ProjectsPostProjectReader is a Reader for the ProjectsPostProject structure.
type ProjectsPostProjectReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ProjectsPostProjectReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewProjectsPostProjectOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewProjectsPostProjectOK creates a ProjectsPostProjectOK with default headers values
func NewProjectsPostProjectOK() *ProjectsPostProjectOK {
	return &ProjectsPostProjectOK{}
}

/*ProjectsPostProjectOK handles this case with default header values.

OK
*/
type ProjectsPostProjectOK struct {
	Payload *models.Project
}

func (o *ProjectsPostProjectOK) Error() string {
	return fmt.Sprintf("[POST /api/accounts/{accountId}/projects][%d] projectsPostProjectOK  %+v", 200, o.Payload)
}

func (o *ProjectsPostProjectOK) GetPayload() *models.Project {
	return o.Payload
}

func (o *ProjectsPostProjectOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Project)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}