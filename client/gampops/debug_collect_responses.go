package gampops

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/blablapolicja/go-gamp/models"
)

// DebugCollectReader is a Reader for the DebugCollect structure.
type DebugCollectReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DebugCollectReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDebugCollectOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDebugCollectOK creates a DebugCollectOK with default headers values
func NewDebugCollectOK() *DebugCollectOK {
	return &DebugCollectOK{}
}

/*DebugCollectOK handles this case with default header values.

Debug Report
*/
type DebugCollectOK struct {
	Payload Payload
}

//Payload is payload received in DebugCollectOK
type Payload struct {
	HitParsingResult []*models.HitParsingResult `json:"hitParsingResult"`
	ParserMsg        []*models.ParserMessage    `json:"parserMessage"`
}

func (o *DebugCollectOK) Error() string {
	return fmt.Sprintf("[POST /debug/collect][%d] debugCollectOK  %+v", 200, o.Payload)
}

func (o *DebugCollectOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
