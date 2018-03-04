package controller

import (
	"errors"
	"github.com/photoshelf/photoshelf-storage/application/service"
	"github.com/photoshelf/photoshelf-storage/domain/model/photo"
	"github.com/photoshelf/photoshelf-storage/presentation/protobuf"
	"golang.org/x/net/context"
)

type photoControllerImpl struct {
	Service service.PhotoService `inject:""`
}

func NewPhotoController() protobuf.PhotoServiceServer {
	return &photoControllerImpl{}
}

func (ctrl *photoControllerImpl) Create(ctx context.Context, req *protobuf.Photo) (*protobuf.Id, error) {
	photograph := photo.New(req.Image)

	id, err := ctrl.Service.Save(*photograph)
	if err != nil {
		return nil, err
	}

	return &protobuf.Id{Value: id.Value()}, nil
}

func (ctrl *photoControllerImpl) Update(ctx context.Context, req *protobuf.Photo) (*protobuf.Empty, error) {
	id := req.Id
	if id == nil {
		return nil, errors.New("id not allows nil")
	}

	photograph := photo.Of(*photo.IdentifierOf(id.Value), req.Image)

	_, err := ctrl.Service.Save(*photograph)
	if err != nil {
		return nil, err
	}

	return &protobuf.Empty{}, nil
}

func (ctrl *photoControllerImpl) Read(ctx context.Context, req *protobuf.Id) (*protobuf.Photo, error) {
	id := photo.IdentifierOf(req.Value)
	photograph, err := ctrl.Service.Find(*id)
	if err != nil {
		return nil, err
	}
	return &protobuf.Photo{Id: &protobuf.Id{Value: photograph.Id().Value()}, Image: photograph.Image()}, nil
}

func (ctrl *photoControllerImpl) Delete(ctx context.Context, req *protobuf.Id) (*protobuf.Empty, error) {
	id := photo.IdentifierOf(req.Value)
	if err := ctrl.Service.Delete(*id); err != nil {
		return nil, err
	}
	return &protobuf.Empty{}, nil
}
