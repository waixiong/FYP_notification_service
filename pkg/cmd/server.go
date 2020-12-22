package cmd

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"path/filepath"

	firebase "firebase.google.com/go"
	"getitqec.com/server/mailnotification/pkg/commons"
	"getitqec.com/server/mailnotification/pkg/logger"
	"google.golang.org/api/option"

	// "getitqec.com/server/mailnotification/pkg/handlers"
	"getitqec.com/server/mailnotification/pkg/model"
	"getitqec.com/server/mailnotification/pkg/protocol/grpc"
	"getitqec.com/server/mailnotification/pkg/protocol/grpcClient"
	"getitqec.com/server/mailnotification/pkg/protocol/rest"
	service "getitqec.com/server/mailnotification/pkg/service/v1"
	// "getitqec.com/server/mailnotification/pkg/protocol/rest"
)

var (
	tls      = flag.Bool("tls", true, "Connection uses TLS if true, else plain TCP")
	certFile = flag.String("cert", "configs/key/certs/mycert.pem", "The TLS cert file")
	keyFile  = flag.String("key", "configs/key/private/mykey.pem", "The TLS key file")
	//jsonDBFile = flag.String("json_db_file", "", "A json file containing a list of features") 1250 1500 || 1500 2000
	port = flag.Int("port", 8080, "The server port")
	//svc  = &dynamodb.DynamoDB{}
)

// Config is configuration for Server
type Config struct {
	GRPCPort string
	HTTPPort string
	// LogLevel is global log level: Debug(-1), Info(0), Warn(1), Error(2), DPanic(3), Panic(4), Fatal(5)
	LogLevel int
	// LogTimeFormat is print time format for logger e.g. 2006-01-02T15:04:05Z07:00
	LogTimeFormat string
}

// RunServer runs gRPC server and HTTP gateway
func RunServer() error {
	ctx := context.Background()

	// get configuration
	cfg := &Config{GRPCPort: commons.ENVVariable("GRPC_SERVICE_PORT"), HTTPPort: commons.ENVVariable("REST_SERVICE_PORT"), LogLevel: -1}

	// initialize logger
	if err := logger.Init(cfg.LogLevel, cfg.LogTimeFormat); err != nil {
		return fmt.Errorf("failed to initialize logger: %v", err)
	}

	mongoDB, err := commons.InitMongoDB(ctx)
	if err != nil {
		return fmt.Errorf("error getting connect mongo client: %v", err)
	}
	defer mongoDB.Disconnect(ctx)
	mmodel := model.InitMailModel(mongoDB, &http.Client{})

	// load PATH_TO_SERVICE_FILE.json base on domain
	opt := option.WithCredentialsFile("./configs/key/fcm.json")
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}
	nmodel := model.InitNotificationModel(mongoDB, app)

	// // initialize handlers
	// handler := handlers.NewHandler(model, mapClient)
	mnService := service.NewServer(mmodel, nmodel)

	// SSL Key
	certFilePath, _ := filepath.Abs(commons.ENVVariable("CRT_PATH"))
	keyFilePath, _ := filepath.Abs(commons.ENVVariable("KEY_PATH"))
	grpcClient.CertFilePath = certFilePath
	grpcClient.ServerAddr = commons.ENVVariable("USER_SERVER_ADDR")

	// run HTTP gateway
	go func() {
		fmt.Println("Run REST")
		_ = rest.RunServer(ctx, cfg.GRPCPort, cfg.HTTPPort, certFilePath, keyFilePath)
	}()
	// fmt.Println("Run REST")
	// _ = rest.RunServer(ctx, cfg.GRPCPort, cfg.HTTPPort, certFilePath, keyFilePath)
	fmt.Println("Run gRPC")
	return grpc.RunServer(ctx, mnService, cfg.GRPCPort, certFilePath, keyFilePath)
}
