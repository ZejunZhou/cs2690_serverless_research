package user

import (
	"fmt"
	"net"
	"time"

	"github.com/delimitrou/DeathStarBench/tree/master/hotelReservation/registry"
	pb "github.com/delimitrou/DeathStarBench/tree/master/hotelReservation/services/user/proto"
	"github.com/delimitrou/DeathStarBench/tree/master/hotelReservation/tls"
	"github.com/google/uuid"
	"github.com/grpc-ecosystem/grpc-opentracing/go/otgrpc"
	"github.com/opentracing/opentracing-go"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
)

const name = "srv-user"

// Server implements the user service
type Server struct {
	pb.UnimplementedUserServer

	users map[string]string
	uuid  string

	Tracer      opentracing.Tracer
	Registry    *registry.Client
	Port        int
	IpAddr      string
	MongoClient *mongo.Client
}

// Run starts the server
func (s *Server) Run() error {
	if s.Port == 0 {
		return fmt.Errorf("server port must be set")
	}

	if s.users == nil {
		s.users = loadUsers(s.MongoClient)
	}

	s.uuid = uuid.New().String()

	opts := []grpc.ServerOption{
		grpc.KeepaliveParams(keepalive.ServerParameters{
			Timeout: 120 * time.Second,
		}),
		grpc.KeepaliveEnforcementPolicy(keepalive.EnforcementPolicy{
			PermitWithoutStream: true,
		}),
		grpc.UnaryInterceptor(
			otgrpc.OpenTracingServerInterceptor(s.Tracer),
		),
	}

	if tlsopt := tls.GetServerOpt(); tlsopt != nil {
		opts = append(opts, tlsopt)
	}

	srv := grpc.NewServer(opts...)

	pb.RegisterUserServer(srv, s)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", s.Port))
	if err != nil {
		log.Fatal().Msgf("failed to listen: %v", err)
	}

	err = s.Registry.Register(name, s.uuid, s.IpAddr, s.Port)
	if err != nil {
		return fmt.Errorf("failed register: %v", err)
	}
	log.Info().Msg("Successfully registered in consul")

	return srv.Serve(lis)
}

// Shutdown cleans up any processes
func (s *Server) Shutdown() {
	s.Registry.Deregister(s.uuid)
}

func (s *Server) CheckUser(ctx context.Context, req *pb.Request) (*pb.Result, error) {
	res := new(pb.Result)

	// log.Info().Msg("CheckUser")
	// log.Info().Msgf("Request user name %s", req.Username)
	res.Correct = false
	if truePass, found := s.users[req.Username]; found {
		// log.Info().Msgf("Database password: %s, Input password: %s", truePass, req.Password)
		res.Correct = req.Password == truePass
	}

	// log.Info().Msgf("CheckUser result: %v", res.Correct)

	return res, nil
}

// loadUsers loads hotel users from mongodb.
func loadUsers(client *mongo.Client) map[string]string {
    collection := client.Database("user-db").Collection("user")
    curr, err := collection.Find(context.TODO(), bson.D{})
    if err != nil {
        log.Error().Msgf("Failed to get users data: %v", err)
        return nil
    }

    var users []User
    if err := curr.All(context.TODO(), &users); err != nil {
        log.Error().Msgf("Failed to decode users data: %v", err)
        return nil
    }

    res := make(map[string]string)
    for _, user := range users {
        res[user.Username] = user.Password
        log.Info().Msgf("Loaded user - Username: %s, Password: %s", user.Username, user.Password)
    }

    log.Info().Msgf("All loaded users: %v", res) // 输出所有加载的用户数据
    log.Info().Msg("Completed loading users")

    return res
}

type User struct {
	Username string `bson:"username"`
	Password string `bson:"password"`
}
