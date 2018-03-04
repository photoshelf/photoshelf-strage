package protobuf

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"io"
	"io/ioutil"
	"mime/multipart"
	"reflect"
	"regexp"
)

type MultipartPhotoMarshaler struct {
}

func (*MultipartPhotoMarshaler) Marshal(v interface{}) ([]byte, error) {
	switch v.(type) {
	case *Id:
		id, _ := v.(*Id)
		return json.Marshal(map[string]string{"id": id.Value})
	case *Empty:
		return nil, nil
	default:
		return nil, fmt.Errorf("photoshelf: wrong type %s", reflect.TypeOf(v))
	}
}

func (*MultipartPhotoMarshaler) Unmarshal(data []byte, v interface{}) error {
	photo, err := toPhoto(v)
	if err != nil {
		return err
	}
	photo.Image = data

	return nil
}

func (*MultipartPhotoMarshaler) NewDecoder(r io.Reader) runtime.Decoder {
	return &PhotoDecoder{r}
}

func (*MultipartPhotoMarshaler) NewEncoder(w io.Writer) runtime.Encoder {
	return &PhotoEncoder{w}
}

func (*MultipartPhotoMarshaler) ContentType() string {
	return "application/json"
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
	return errors.New("photoshelf: not implement yet")
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
