package game

import (
	"context"
	"errors"
	"fmt"

	pbGlobal "bitbucket.org/funplus/server/proto/global"
	"bitbucket.org/funplus/server/services/game/player"
)

func (m *MsgRegister) handleCollectionActive(ctx context.Context, p ...interface{}) error {
	acct := p[0].(*player.Account)
	msg, ok := p[1].(*pbGlobal.C2S_CollectionActive)
	if !ok {
		return errors.New("handleCollectionActive failed: recv message body error")
	}
	pl, err := m.am.GetPlayerByAccount(acct)
	if err != nil {
		return fmt.Errorf("GetPlayerByAccount failed: %w", err)
	}

	return pl.CollectionManager().CollectionActive(msg.TypeId)
}

func (m *MsgRegister) handleCollectionStarup(ctx context.Context, p ...interface{}) error {
	acct := p[0].(*player.Account)
	msg, ok := p[1].(*pbGlobal.C2S_CollectionStarup)
	if !ok {
		return errors.New("handleCollectionStarup failed: recv message body error")
	}
	pl, err := m.am.GetPlayerByAccount(acct)
	if err != nil {
		return fmt.Errorf("GetPlayerByAccount failed: %w", err)
	}

	return pl.CollectionManager().CollectionStarup(msg.TypeId)
}
