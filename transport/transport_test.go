package transport

import (
	"context"
	"log"
	"sync"
	"testing"

	pbAccount "github.com/yokaiio/yokai_server/proto/account"
	"github.com/yokaiio/yokai_server/transport/codec"
)

var (
	trSrv  = NewTransport("tcp")
	regSrv = NewTransportRegister()
	trCli  = NewTransport("tcp")
	regCli = NewTransportRegister()
	wg     sync.WaitGroup
)

func handleServerSocket(ctx context.Context, sock Socket) {
	defer func() {
		sock.Close()
	}()

	for {
		select {
		case <-ctx.Done():
			break
		default:
		}

		msg, h, err := sock.Recv(regSrv)
		if err != nil {
			log.Printf("tcp server handle socket error: %w", err)
			return
		}

		h.Fn(ctx, sock, msg)
	}
}

func handleClient(ctx context.Context, sock Socket, p *Message) {
	msg, ok := p.Body.(*pbAccount.C2M_AccountLogon)
	if !ok {
		log.Fatalf("handleClient failed")
	}

	var sendMsg Message
	sendMsg.Type = BodyProtobuf
	sendMsg.Name = "M2C_AccountLogon"
	sendMsg.Body = &pbAccount.M2C_AccountLogon{
		RpcId:      2,
		Error:      0,
		Message:    "OK",
		PlayerName: msg.AccountName,
	}

	sock.Send(&sendMsg)
	wg.Done()
}

func handleServer(ctx context.Context, sock Socket, p *Message) {
	msg, ok := p.Body.(*pbAccount.M2C_AccountLogon)
	if !ok {
		log.Fatalf("handleServer failed")
	}

	if msg.PlayerName != "test_name" {
		log.Fatalf("handleServer failed")
	}

	wg.Done()
}

func TestTransportTcp(t *testing.T) {

	// tcp server
	trSrv.Init(
		Timeout(DefaultDialTimeout),
		Codec(&codec.ProtoBufMarshaler{}),
	)

	regSrv.RegisterProtobufMessage(&pbAccount.C2M_AccountLogon{}, handleClient)

	wg.Add(1)
	go func() {
		err := trSrv.ListenAndServe(context.TODO(), ":7030", handleServerSocket)
		if err != nil {
			log.Fatal("TcpServer ListenAndServe failed:%w", err)
		}
	}()

	// tcp client
	trCli.Init(
		Timeout(DefaultDialTimeout),
	)

	regCli.RegisterProtobufMessage(&pbAccount.M2C_AccountLogon{}, handleServer)

	sockClient, err := trCli.Dial("127.0.0.1:7030")
	if err != nil {
		log.Fatalf("unexpected dial err:%w", err)
	}

	wg.Add(1)
	msg := &Message{
		Type: BodyProtobuf,
		Name: "yokai_account.C2M_AccountLogon",
		Body: &pbAccount.C2M_AccountLogon{
			RpcId:       1,
			UserId:      1,
			AccountId:   1,
			AccountName: "test_name",
		},
	}

	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		defer func() {
			sockClient.Close()
		}()

		for {
			select {
			case <-ctx.Done():
				return
			default:
			}

			if msg, h, err := sockClient.Recv(regCli); err != nil {
				log.Fatalf("Unexpected recv err:%w", err)
			} else {
				h.Fn(ctx, sockClient, msg)
			}
		}
	}()

	if err := sockClient.Send(msg); err != nil {
		log.Fatalf("client send socket failed:%w", err)
	}

	wg.Wait()
	cancel()
}
