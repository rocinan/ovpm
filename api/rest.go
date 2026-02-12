package api

import (
	"fmt"
	"google.golang.org/protobuf/encoding/protojson"
	"net/http"
	"strings"

	"github.com/asaskevich/govalidator"
	"github.com/rocinan/ovpm/api/pb"
	assetfs "github.com/elazarl/go-bindata-assetfs"
	"github.com/go-openapi/runtime/middleware"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/sirupsen/logrus"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

// NewRESTServer returns a new REST server.
func NewRESTServer(grpcPort string) (http.Handler, context.CancelFunc, error) {
	mux := http.NewServeMux()
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	if !govalidator.IsNumeric(grpcPort) {
		return nil, cancel, fmt.Errorf("grpcPort should be numeric")
	}
	endPoint := fmt.Sprintf("localhost:%s", grpcPort)
	ctx = NewOriginTypeContext(ctx, OriginTypeREST)
	gmux := runtime.NewServeMux(runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.JSONPb{
		MarshalOptions: protojson.MarshalOptions{
			UseProtoNames:   true,
			EmitUnpopulated: true,
		},
		UnmarshalOptions: protojson.UnmarshalOptions{
			DiscardUnknown: true,
		},
	}))
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err := pb.RegisterVPNServiceHandlerFromEndpoint(ctx, gmux, endPoint, opts)
	if err != nil {
		return nil, cancel, err
	}

	err = pb.RegisterUserServiceHandlerFromEndpoint(ctx, gmux, endPoint, opts)
	if err != nil {
		return nil, cancel, err
	}

	err = pb.RegisterNetworkServiceHandlerFromEndpoint(ctx, gmux, endPoint, opts)
	if err != nil {
		return nil, cancel, err
	}

	err = pb.RegisterAuthServiceHandlerFromEndpoint(ctx, gmux, endPoint, opts)
	if err != nil {
		return nil, cancel, err
	}

	mux.HandleFunc("/api/specs/", specsHandler)
	mware := middleware.Redoc(middleware.RedocOpts{
		BasePath: "/api/docs/",
		SpecURL:  "/api/specs/user.swagger.json",
		Path:     "user",
	}, gmux)
	mware = middleware.Redoc(middleware.RedocOpts{
		BasePath: "/api/docs/",
		SpecURL:  "/api/specs/vpn.swagger.json",
		Path:     "vpn",
	}, mware)
	mware = middleware.Redoc(middleware.RedocOpts{
		BasePath: "/api/docs/",
		SpecURL:  "/api/specs/network.swagger.json",
		Path:     "network",
	}, mware)
	mware = middleware.Redoc(middleware.RedocOpts{
		BasePath: "/api/docs/",
		SpecURL:  "/api/specs/auth.swagger.json",
		Path:     "auth",
	}, mware)
	mux.Handle("/api/", mware)

	return allowCORS(mux), cancel, nil
}
func preflightHandler(w http.ResponseWriter, r *http.Request) {
	headers := []string{"Content-Type", "Accept", "Authorization"}
	w.Header().Set("Access-Control-Allow-Headers", strings.Join(headers, ","))
	methods := []string{"GET", "HEAD", "POST", "PUT", "DELETE"}
	w.Header().Set("Access-Control-Allow-Methods", strings.Join(methods, ","))
	w.Header().Set("Access-Control-Expose-Headers", "Access-Control-Allow-Origin")
	logrus.Debugf("rest: preflight request for %s", r.URL.Path)
	return
}

func allowCORS(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if origin := r.Header.Get("Origin"); origin != "" {
			w.Header().Set("Access-Control-Allow-Origin", origin)
			if r.Method == "OPTIONS" && r.Header.Get("Access-Control-Request-Method") != "" {
				preflightHandler(w, r)
				return
			}
		}
		h.ServeHTTP(w, r)
	})
}
