package server

import (
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/golang/protobuf/jsonpb"
	"github.com/photoshelf/photoshelf-storage/presentation/protobuf"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"net"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestGateway(t *testing.T) {
	t.Run("GET /api/v1/photos/identifier", func(t *testing.T) {
		listener, err := net.Listen("tcp", ":0")

		gw, err := NewGateway(listener.Addr().String())
		if err != nil {
			t.Fatal(err)
		}

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockGrpcServer := protobuf.NewMockPhotoServiceServer(ctrl)
		mockGrpcServer.EXPECT().
			Read(gomock.Any(), &protobuf.Id{"identifier"}).
			Times(1).
			Return(&protobuf.Photo{
				Id:    &protobuf.Id{"id_test"},
				Image: []byte("Hello World."),
			}, nil)
		mockGrpcServer.EXPECT().
			Read(gomock.Any(), gomock.Any()).
			AnyTimes().
			Return(nil, errors.New("shouldn't call this"))

		s := grpc.NewServer()
		protobuf.RegisterPhotoServiceServer(s, mockGrpcServer)

		go func() {
			if err := s.Serve(listener); err != nil {
				t.Error("Fail to start grpc mock server")
			}
			defer s.Stop()
		}()

		req := httptest.NewRequest("GET", "/api/v1/photos/identifier", nil)
		rec := httptest.NewRecorder()

		gw.ServeHTTP(rec, req)

		actual := rec.Body.Bytes()
		expected := []byte("Hello World.")

		assert.Equal(t, 200, rec.Result().StatusCode)
		assert.Equal(t, expected, actual)
	})

	t.Run("POST /api/v1/photos", func(t *testing.T) {
		listener, err := net.Listen("tcp", ":0")
		if err != nil {
			t.Fatal(err)
		}

		gw, err := NewGateway(listener.Addr().String())
		if err != nil {
			t.Fatal(err)
		}

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockGrpcServer := protobuf.NewMockPhotoServiceServer(ctrl)
		mockGrpcServer.EXPECT().
			Create(gomock.Any(), &protobuf.Photo{Image: []byte("Hello World.")}).
			Times(1).
			Return(&protobuf.Id{"test_id"}, nil)
		mockGrpcServer.EXPECT().
			Create(gomock.Any(), gomock.Any()).
			AnyTimes().
			Return(nil, errors.New("shouldn't call this"))

		s := grpc.NewServer()
		protobuf.RegisterPhotoServiceServer(s, mockGrpcServer)

		go func() {
			if err := s.Serve(listener); err != nil {
				t.Error("Fail to start grpc mock server")
			}
			defer s.Stop()
		}()

		reqBody := "{\"image\":\"SGVsbG8gV29ybGQu\"}" // encoded base64 "Hello World."
		req := httptest.NewRequest("POST", "/api/v1/photos", strings.NewReader(reqBody))
		rec := httptest.NewRecorder()

		gw.ServeHTTP(rec, req)

		res := &protobuf.Id{}
		jsonpb.Unmarshal(rec.Body, res)

		actual := res.Value
		expected := "test_id"

		assert.Equal(t, 201, rec.Result().StatusCode)
		assert.Equal(t, expected, actual)
	})

	t.Run("PUT /api/v1/photos/identifier", func(t *testing.T) {
		listener, err := net.Listen("tcp", ":0")

		gw, err := NewGateway(listener.Addr().String())
		if err != nil {
			t.Fatal(err)
		}

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockGrpcServer := protobuf.NewMockPhotoServiceServer(ctrl)
		mockGrpcServer.EXPECT().
			Update(gomock.Any(), &protobuf.Photo{Id: &protobuf.Id{"identifier"}, Image: []byte("Hello World.")}).
			Times(1).
			Return(&protobuf.Empty{}, nil)
		mockGrpcServer.EXPECT().
			Update(gomock.Any(), gomock.Any()).
			AnyTimes().
			Return(nil, errors.New("shouldn't call this"))

		s := grpc.NewServer()
		protobuf.RegisterPhotoServiceServer(s, mockGrpcServer)

		go func() {
			if err := s.Serve(listener); err != nil {
				t.Fatal(err)
			}
			defer s.Stop()
		}()

		reqBody := "{\"id\":{\"value\":\"identifier\"},\"image\":\"SGVsbG8gV29ybGQu\"}"
		req := httptest.NewRequest("PUT", "/api/v1/photos/identifier", strings.NewReader(reqBody))
		rec := httptest.NewRecorder()

		gw.ServeHTTP(rec, req)

		res := &protobuf.Empty{}
		jsonpb.Unmarshal(rec.Body, res)

		actual := res
		expected := &protobuf.Empty{}

		assert.Equal(t, 200, rec.Result().StatusCode)
		assert.Equal(t, expected, actual)
	})

	t.Run("DELETE /api/v1/photos/identifier", func(t *testing.T) {
		listener, err := net.Listen("tcp", ":0")

		gw, err := NewGateway(listener.Addr().String())
		if err != nil {
			t.Fatal(err)
		}

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockGrpcServer := protobuf.NewMockPhotoServiceServer(ctrl)
		mockGrpcServer.EXPECT().
			Delete(gomock.Any(), &protobuf.Id{"identifier"}).
			Times(1).
			Return(&protobuf.Empty{}, nil)
		mockGrpcServer.EXPECT().
			Delete(gomock.Any(), gomock.Any()).
			AnyTimes().
			Return(nil, errors.New("shouldn't call this"))

		s := grpc.NewServer()
		protobuf.RegisterPhotoServiceServer(s, mockGrpcServer)

		go func() {
			if err := s.Serve(listener); err != nil {
				t.Error("Fail to start grpc mock server")
			}
			defer s.Stop()
		}()

		req := httptest.NewRequest("DELETE", "/api/v1/photos/identifier", nil)
		rec := httptest.NewRecorder()

		gw.ServeHTTP(rec, req)

		res := &protobuf.Empty{}
		jsonpb.Unmarshal(rec.Body, res)

		actual := res
		expected := &protobuf.Empty{}

		assert.Equal(t, 200, rec.Result().StatusCode)
		assert.Equal(t, expected, actual)
	})
}
