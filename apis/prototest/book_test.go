package prototest

import (
	"errors"
	"fmt"
	pb "grpc-boot-starter/apis/protov1"
	"testing"

	"github.com/bufbuild/protovalidate-go"
	"github.com/stretchr/testify/assert"
)

// TestBook on book creation and validation
func TestBook(t *testing.T) {
	book := &pb.BookCreateInput{
		Title:  "111",
		Price:  10,
		Amount: 0,
	}
	//
	err := protovalidate.Validate(book)
	// fmt.Printf("err: %v", err)
	assert.NotNil(t, err)
	//
	var valErr *protovalidate.ValidationError
	if ok := errors.As(err, &valErr); ok {
		//pb := valErr.ToProto()
		// fmt.Printf("%T", pb)
		for _, violation := range valErr.Violations {
			fmt.Println(*violation.FieldPath)
		}
	}
}

func TestBookMaxLength(t *testing.T) {
	book := &pb.BookCreateInput{
		Title:  "1234567890123",
		Price:  10,
		Amount: 10,
	}
	//
	err := protovalidate.Validate(book)
	fmt.Printf("err: %v", err)
	assert.NotNil(t, err)
	//
	var valErr *protovalidate.ValidationError
	if ok := errors.As(err, &valErr); ok {
		pb := valErr.ToProto()
		fmt.Println(pb)
	}
}

func TestBookEmail(t *testing.T) {
	email := "abc"
	book := &pb.BookCreateInput{
		Title:    "123456789",
		Price:    10,
		Amount:   10,
		Email:    email,
		CoverUrl: "http://example.com/img/a/b/c.png",
	}
	//
	err := protovalidate.Validate(book)
	fmt.Printf("err: %v", err)
	assert.NotNil(t, err)
	//
	var valErr *protovalidate.ValidationError
	if ok := errors.As(err, &valErr); ok {
		pb := valErr.ToProto()
		fmt.Println(pb)
	}
}
