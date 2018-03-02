package protobuf

import (
	"io"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"errors"
	"reflect"
	"fmt"
)

type CustomMarshaler struct {}

func (c *CustomMarshaler) Marshal(v interface{}) ([]byte, error) {
	fmt.Println(reflect.TypeOf(v))
	photo, ok := v.(*Photo)
	if !ok {
		return nil, errors.New("(´;ω;｀)")
	}

	return photo.Image, nil
}

func (c *CustomMarshaler) Unmarshal(data []byte, v interface{}) error {
	photo, ok := v.(*Photo)
	if !ok {
		return errors.New("(´;ω;｀)")
	}
	photo.Image = data

	return nil
}

func (c *CustomMarshaler) NewDecoder(r io.Reader) runtime.Decoder {
	return &PhotoDecoder{r}
}

func (c *CustomMarshaler) NewEncoder(w io.Writer) runtime.Encoder {
	return &PhotoEncoder{w}
}

func (c *CustomMarshaler) ContentType() string {
	return "multipart/form-data"
}

type PhotoDecoder struct {
	r io.Reader
}

func (decoder *PhotoDecoder) Decode(v interface{}) error {
	fmt.Println(reflect.TypeOf(v))
	photo, ok := v.(*Photo)
	if !ok {
		return errors.New("(´;ω;｀)")
	}
	if _, err := decoder.r.Read(photo.Image); err != nil {
		return err
	}

	return nil
}

type PhotoEncoder struct {
	w io.Writer
}

func (encoder *PhotoEncoder) Encode(v interface{}) error {
	fmt.Println(reflect.TypeOf(v))
	photo, ok := v.(*Photo)
	if !ok {
		return errors.New("(´;ω;｀)")
	}
	if _, err := encoder.w.Write(photo.Image); err != nil {
		return err
	}

	return nil
}
