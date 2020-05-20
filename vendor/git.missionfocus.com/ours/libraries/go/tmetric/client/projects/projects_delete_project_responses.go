// Code generated by go-swagger; DO NOT EDIT.

package projects

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// ProjectsDeleteProjectReader is a Reader for the ProjectsDeleteProject structure.
type ProjectsDeleteProjectReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ProjectsDeleteProjectReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewProjectsDeleteProjectNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewProjectsDeleteProjectNoContent creates a ProjectsDeleteProjectNoContent with default headers values
func NewProjectsDeleteProjectNoContent() *ProjectsDeleteProjectNoContent {
	return &ProjectsDeleteProjectNoContent{}
}

/*ProjectsDeleteProjectNoContent handles this case with default header values.

No Content
*/
type ProjectsDeleteProjectNoContent struct {
}

func (o *ProjectsDeleteProjectNoContent) Error() string {
	return fmt.Sprintf("[DELETE /api/accounts/{accountId}/projects/{projectid}][%d] projectsDeleteProjectNoContent ", 204)
}

func (o *ProjectsDeleteProjectNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
