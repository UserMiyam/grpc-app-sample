# grpc-app-sample
Udamy-gRPC-Go-Firebse-Docker

-環境構築
git clone[リポジトリ]<br>
Go1.25　install  
Docker　install  

REST APIの場合
データ形式: JSONやXMLなど、テキストベースのフォーマットが一般的
通信: HTTPのエンドポイント（GET/POST/PUT/DELETE など）を利用

gRPCの場合
データ形式: Protocol Buffers（通称 protobuf）が標準
.proto ファイルでデータ構造やサービス（RPCメソッド）を定義
その定義をもとに、サーバーとクライアントのコードを自動生成できる
通信: HTTP/2 をベースにしており、双方向ストリーミングなどもサポート

ポイント
gRPC自体はプロトコルバッファを「必須」としているわけではなく、
他のシリアライズ方式（例：JSON、FlatBuffers、Avroなど）も理論的には使える。
ただし、Googleが設計したgRPCの公式ツールチェーンでは protoが第一級サポート されており、
実務でもほぼ「gRPC = Protocol Buffers」で使われるのが一般的です。

