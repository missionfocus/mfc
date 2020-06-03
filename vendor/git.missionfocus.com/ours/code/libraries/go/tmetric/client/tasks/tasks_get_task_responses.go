// Code generated by go-swagger; DO NOT EDIT.

package tasks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"git.missionfocus.com/ours/code/libraries/go/tmetric/models"
)

// TasksGetTaskReader is a Reader for the TasksGetTask structure.
type TasksGetTaskReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TasksGetTaskReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewTasksGetTaskOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewTasksGetTaskOK creates a TasksGetTaskOK with default headers values
func NewTasksGetTaskOK() *TasksGetTaskOK {
	return &TasksGetTaskOK{}
}

/*TasksGetTaskOK handles this case with default header values.

OK
*/
type TasksGetTaskOK struct {
	Payload *models.ProjectTask
}

func (o *TasksGetTaskOK) Error() string {
	return fmt.Sprintf("[GET /api/accounts/{accountId}/tasks/{taskId}][%d] tasksGetTaskOK  %+v", 200, o.Payload)
}

func (o *TasksGetTaskOK) GetPayload() *models.ProjectTask {
	return o.Payload
}

func (o *TasksGetTaskOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProjectTask)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}