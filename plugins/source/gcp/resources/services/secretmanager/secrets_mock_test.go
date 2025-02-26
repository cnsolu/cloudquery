// Code generated by codegen; DO NOT EDIT.

package secretmanager

import (
	"context"
	"fmt"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/cloudquery/plugins/source/gcp/client"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"net"
	"testing"

	"cloud.google.com/go/secretmanager/apiv1"

	pb "google.golang.org/genproto/googleapis/cloud/secretmanager/v1"

	"google.golang.org/api/option"
)

func createSecrets() (*client.Services, error) {
	fakeServer := &fakeSecretsServer{}
	l, err := net.Listen("tcp", "localhost:0")
	if err != nil {
		return nil, fmt.Errorf("failed to listen: %w", err)
	}
	gsrv := grpc.NewServer()
	pb.RegisterSecretManagerServiceServer(gsrv, fakeServer)
	fakeServerAddr := l.Addr().String()
	go func() {
		if err := gsrv.Serve(l); err != nil {
			panic(err)
		}
	}()

	// Create a client.
	svc, err := secretmanager.NewClient(context.Background(),
		option.WithEndpoint(fakeServerAddr),
		option.WithoutAuthentication(),
		option.WithGRPCDialOption(grpc.WithTransportCredentials(insecure.NewCredentials())),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create grpc client: %w", err)
	}

	return &client.Services{
		SecretmanagerClient: svc,
	}, nil
}

type fakeSecretsServer struct {
	pb.UnimplementedSecretManagerServiceServer
}

func (f *fakeSecretsServer) ListSecrets(context.Context, *pb.ListSecretsRequest) (*pb.ListSecretsResponse, error) {
	resp := pb.ListSecretsResponse{}
	if err := faker.FakeObject(&resp); err != nil {
		return nil, fmt.Errorf("failed to fake data: %w", err)
	}
	resp.NextPageToken = ""
	return &resp, nil
}

func TestSecrets(t *testing.T) {
	client.MockTestHelper(t, Secrets(), createSecrets, client.TestOptions{})
}
