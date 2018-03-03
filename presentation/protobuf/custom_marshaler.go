package protobuf

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"io"
	"io/ioutil"
	"mime/multipart"
	"reflect"
	"regexp"
)

type CustomMarshaler struct{}

func (c *CustomMarshaler) Marshal(v interface{}) ([]byte, error) {
	photo, err := toPhoto(v)
	if err != nil {
		return nil, err
	}

	return photo.Image, nil
}

func (c *CustomMarshaler) Unmarshal(data []byte, v interface{}) error {
	photo, err := toPhoto(v)
	if err != nil {
		return err
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
	photo, err := toPhoto(v)
	if err != nil {
		return err
	}

	data, err := ioutil.ReadAll(decoder.r)
	if err != nil {
		return err
	}

	boundary, err := messageBoundary(data)
	if err != nil {
		return err
	}

	reader := multipart.NewReader(bytes.NewReader(data), boundary)
	for {
		part, err := reader.NextPart()
		if err == io.EOF {
			return errors.New("photoshelf: EOF")
		}
		if err != nil {
			return err
		}
		if part.FormName() == "photo" {
			photoData, err := ioutil.ReadAll(part)
			if err != nil {
				return err
			}

			photo.Image = photoData

			return nil
		}
	}

	return errors.New("photoshelf: form photo not found")
}

type PhotoEncoder struct {
	w io.Writer
}

func (encoder *PhotoEncoder) Encode(v interface{}) error {
	photo, ok := v.(*Photo)
	if !ok {
		return errors.New("(´;ω;｀)")
	}
	if _, err := encoder.w.Write(photo.Image); err != nil {
		return err
	}

	return nil
}

func messageBoundary(data []byte) (string, error) {
	assigned := regexp.MustCompile("--(.+)\r\n")
	boundary := assigned.FindSubmatch(data)[1]

	return string(boundary), nil
}

func toPhoto(v interface{}) (*Photo, error) {
	photo, ok := v.(*Photo)
	if !ok {
		return nil, fmt.Errorf("photoshelf: second argument is required Photo type, but %s", reflect.TypeOf(v))
	}
	return photo, nil
}
