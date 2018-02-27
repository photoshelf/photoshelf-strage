package protobuf

import (
	"net/http"

	"github.com/golang/protobuf/proto"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"golang.org/x/net/context"
)

func forwardReadResp(
	ctx context.Context,
	mux *runtime.ServeMux,
	marshaler runtime.Marshaler,
	w http.ResponseWriter,
	req *http.Request,
	resp proto.Message,
	opts ...func(context.Context, http.ResponseWriter, proto.Message) error,
) {
	photo, ok := resp.(*Photo)
	if !ok {
		runtime.ForwardResponseMessage(ctx, mux, marshaler, w, req, resp, opts...)
	}

	mineType := http.DetectContentType(photo.Image)

	w.Header().Set("Content-Type", mineType)
	w.WriteHeader(200)
	w.Write([]byte(photo.Image))
}

func init() {
	forward_PhotoService_Read_0 = forwardReadResp
}
