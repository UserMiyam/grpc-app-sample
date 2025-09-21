package main

//パッケージ宣言とインポート文法
import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"

	hellopb "grpc-app-sample/gen/api" //生成されたコードのインポート

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// 型定義
// hellopb.UnimplementedGreetingServiceServer の全メソッドが myServer で使用可能
// 継承ではなく「委譲」の仕組み
type myServer struct {
	hellopb.UnimplementedGreetingServiceServer
}

// 関数定義とポインタレシー
func NewMyServer() *myServer {
	return &myServer{}
}

// メソッド定義とレシーバー
func (s *myServer) Hello(ctx context.Context, req *hellopb.HelloRequest) (*hellopb.HelloResponse, error) {
	// return
	return &hellopb.HelloResponse{
		Message: fmt.Sprintf("Hello, %s!", req.GetName()),
	}, nil
}

// main関数と変数宣言
func main() {
	// 8080番ポートのListenerを作成
	port := 8080
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	//エラーハンドリングパターン
	if err != nil {
		panic(err)
	}
	//変数宣言と代入
	s := grpc.NewServer()
	//関数呼び出し
	hellopb.RegisterGreetingServiceServer(s, NewMyServer())

	//この行がリフレクション機能を有効にしている。おかげでgrpcurlが使える！
	reflection.Register(s)

	go func() {
		log.Printf("start gRPC server port: %v", port)
		s.Serve(listener)
	}()
	//チャンネル（Channel）とシグナル（Signal）
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	//ログ出力とメソッド呼び出し
	log.Println("stopping gRPC server...")
	s.GracefulStop()
}
