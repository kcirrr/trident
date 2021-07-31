// Code generated by go-swagger; DO NOT EDIT.

package storage

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/netapp/trident/storage_drivers/ontap/api/rest/models"
)

// SnapshotPolicyScheduleCreateReader is a Reader for the SnapshotPolicyScheduleCreate structure.
type SnapshotPolicyScheduleCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SnapshotPolicyScheduleCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewSnapshotPolicyScheduleCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewSnapshotPolicyScheduleCreateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSnapshotPolicyScheduleCreateCreated creates a SnapshotPolicyScheduleCreateCreated with default headers values
func NewSnapshotPolicyScheduleCreateCreated() *SnapshotPolicyScheduleCreateCreated {
	return &SnapshotPolicyScheduleCreateCreated{}
}

/* SnapshotPolicyScheduleCreateCreated describes a response with status code 201, with default header values.

Created
*/
type SnapshotPolicyScheduleCreateCreated struct {
}

func (o *SnapshotPolicyScheduleCreateCreated) Error() string {
	return fmt.Sprintf("[POST /storage/snapshot-policies/{snapshot-policy.uuid}/schedules][%d] snapshotPolicyScheduleCreateCreated ", 201)
}

func (o *SnapshotPolicyScheduleCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewSnapshotPolicyScheduleCreateDefault creates a SnapshotPolicyScheduleCreateDefault with default headers values
func NewSnapshotPolicyScheduleCreateDefault(code int) *SnapshotPolicyScheduleCreateDefault {
	return &SnapshotPolicyScheduleCreateDefault{
		_statusCode: code,
	}
}

/* SnapshotPolicyScheduleCreateDefault describes a response with status code -1, with default header values.

 ONTAP Error Response Codes
| Error Code | Description |
| ---------- | ----------- |
| 1638407    | When adding schedule to a Snapshot copy policy, the count for that schedule must be specified. |
| 1638410    | Specified schedule already exists in snapshot policy. |
| 1638413    | Schedule not found. |
| 1638451    | This operation would result in total Snapshot copy count for the policy to exceed maximum supported count. |
| 1638508    | Another schedule has the same prefix within this policy. |
| 1638528    | This operation is not supported in a mixed-version cluster. |
| 1638531    | This operation is not supported because specified policy is owned by the cluster admin. |

*/
type SnapshotPolicyScheduleCreateDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the snapshot policy schedule create default response
func (o *SnapshotPolicyScheduleCreateDefault) Code() int {
	return o._statusCode
}

func (o *SnapshotPolicyScheduleCreateDefault) Error() string {
	return fmt.Sprintf("[POST /storage/snapshot-policies/{snapshot-policy.uuid}/schedules][%d] snapshot_policy_schedule_create default  %+v", o._statusCode, o.Payload)
}
func (o *SnapshotPolicyScheduleCreateDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *SnapshotPolicyScheduleCreateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}