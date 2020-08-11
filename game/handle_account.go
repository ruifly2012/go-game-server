package game

import (
	"context"
	"errors"
	"fmt"

	logger "github.com/sirupsen/logrus"
	"github.com/yokaiio/yokai_server/game/player"
	pbAccount "github.com/yokaiio/yokai_server/proto/account"
	"github.com/yokaiio/yokai_server/transport"
)

func (m *MsgHandler) handleAccountTest(ctx context.Context, sock transport.Socket, p *transport.Message) error {
	return nil
}

func (m *MsgHandler) handleAccountLogon(ctx context.Context, sock transport.Socket, p *transport.Message) error {
	msg, ok := p.Body.(*pbAccount.C2M_AccountLogon)
	if !ok {
		return errors.New("handleAccountLogon failed: cannot assert value to message")
	}

	err := m.g.am.AccountLogon(ctx, msg.UserId, msg.AccountId, msg.AccountName, sock)
	if err != nil {
		return fmt.Errorf("handleAccountLogon failed: %w", err)
	}

	return m.g.am.AccountExecute(sock, func(acct *player.Account) error {
		reply := &pbAccount.M2C_AccountLogon{
			RpcId:       msg.RpcId,
			UserId:      acct.UserId,
			AccountId:   acct.ID,
			PlayerId:    -1,
			PlayerName:  "",
			PlayerLevel: 0,
		}

		pl, err := m.g.am.GetPlayerByAccount(acct)
		if err != nil {
			return fmt.Errorf("handleAccountLogon.AccountExecute failed: %w", err)
		}

		reply.PlayerId = pl.GetID()
		reply.PlayerName = pl.GetName()
		reply.PlayerLevel = pl.GetLevel()
		acct.SendProtoMessage(reply)
		return nil
	})
}

func (m *MsgHandler) handleHeartBeat(ctx context.Context, sock transport.Socket, p *transport.Message) error {
	msg, ok := p.Body.(*pbAccount.C2M_HeartBeat)
	if !ok {
		return errors.New("handleHeartBeat failed: cannot assert value to message")
	}

	return m.g.am.AccountExecute(sock, func(acct *player.Account) error {
		acct.HeartBeat(msg.RpcId)
		return nil
	})
}

// todo after account logon
func (m *MsgHandler) handleAccountConnected(ctx context.Context, sock transport.Socket, p *transport.Message) error {
	if acct := m.g.am.GetAccountBySock(sock); acct != nil {
		accountID := p.Body.(*pbAccount.MC_AccountConnected).AccountId
		logger.WithFields(logger.Fields{
			"account_id": accountID,
		}).Info("account connected")

		// todo after connected
	}

	return nil
}

// client disconnect
func (m *MsgHandler) handleAccountDisconnect(ctx context.Context, sock transport.Socket, p *transport.Message) error {
	return player.ErrAccountDisconnect
}
