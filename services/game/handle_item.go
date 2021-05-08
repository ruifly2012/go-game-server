package game

import (
	"context"
	"errors"
	"fmt"

	"bitbucket.org/funplus/server/define"
	pbGlobal "bitbucket.org/funplus/server/proto/global"
	"bitbucket.org/funplus/server/services/game/item"
	"bitbucket.org/funplus/server/services/game/player"
)

func (m *MsgRegister) handleDelItem(ctx context.Context, p ...interface{}) error {
	acct := p[0].(*player.Account)
	msg, ok := p[1].(*pbGlobal.C2S_DelItem)
	if !ok {
		return errors.New("handleDelItem failed: recv message body error")
	}
	pl, err := m.am.GetPlayerByAccount(acct)
	if err != nil {
		return fmt.Errorf("handleDelItem.AccountExecute failed: %w", err)
	}

	it, err := pl.ItemManager().GetItem(msg.Id)
	if err != nil {
		return fmt.Errorf("handleDelItem.AccountExecute failed: %w", err)
	}

	// clear hero's equip id before delete item
	if it.GetType() == define.Item_TypeEquip {
		equip := it.(*item.Equip)
		equipObjID := equip.GetEquipObj()
		if equipObjID != -1 {
			if err := pl.HeroManager().TakeoffEquip(equipObjID, equip.EquipEnchantEntry.EquipPos); err != nil {
				return fmt.Errorf("TakeoffEquip failed: %w", err)
			}
		}
	}

	// delete item
	return pl.ItemManager().DeleteItem(msg.Id)
}

func (m *MsgRegister) handleUseItem(ctx context.Context, p ...interface{}) error {
	acct := p[0].(*player.Account)
	msg, ok := p[1].(*pbGlobal.C2S_UseItem)
	if !ok {
		return errors.New("handleUseItem failed: recv message body error")
	}
	pl, err := m.am.GetPlayerByAccount(acct)
	if err != nil {
		return fmt.Errorf("handleUseItem.AccountExecute failed: %w", err)
	}

	if err := pl.ItemManager().UseItem(msg.ItemId); err != nil {
		return fmt.Errorf("handleUseItem.AccountExecute failed: %w", err)
	}

	return nil
}

func (m *MsgRegister) handlePutonEquip(ctx context.Context, p ...interface{}) error {
	acct := p[0].(*player.Account)
	msg, ok := p[1].(*pbGlobal.C2S_PutonEquip)
	if !ok {
		return errors.New("handlePutonEquip failed: recv message body error")
	}
	pl, err := m.am.GetPlayerByAccount(acct)
	if err != nil {
		return fmt.Errorf("handlePutonEquip.AccountExecute failed: %w", err)
	}

	if err := pl.HeroManager().PutonEquip(msg.HeroId, msg.EquipId); err != nil {
		return fmt.Errorf("handlePutonEquip.AccountExecute failed: %w", err)
	}

	return nil
}

func (m *MsgRegister) handleTakeoffEquip(ctx context.Context, p ...interface{}) error {
	acct := p[0].(*player.Account)
	msg, ok := p[1].(*pbGlobal.C2S_TakeoffEquip)
	if !ok {
		return errors.New("handleTakeoffEquip failed: recv message body error")
	}
	pl, err := m.am.GetPlayerByAccount(acct)
	if err != nil {
		return fmt.Errorf("handleTakeoffEquip failed: %w", err)
	}

	return pl.HeroManager().TakeoffEquip(msg.HeroId, msg.Pos)
}

func (m *MsgRegister) handleEquipPromote(ctx context.Context, p ...interface{}) error {
	acct := p[0].(*player.Account)
	msg, ok := p[1].(*pbGlobal.C2S_EquipPromote)
	if !ok {
		return errors.New("handleEquipPromote failed: recv message body error")
	}

	pl, err := m.am.GetPlayerByAccount(acct)
	if err != nil {
		return fmt.Errorf("handleEquipPromote failed: %w", err)
	}

	return pl.ItemManager().EquipPromote(msg.ItemId)
}

func (m *MsgRegister) handleEquipStarup(ctx context.Context, p ...interface{}) error {
	acct := p[0].(*player.Account)
	msg, ok := p[1].(*pbGlobal.C2S_EquipStarup)
	if !ok {
		return errors.New("handleEquipStarup failed: recv message body error")
	}

	pl, err := m.am.GetPlayerByAccount(acct)
	if err != nil {
		return fmt.Errorf("handleEquipStarup failed: %w", err)
	}

	return pl.ItemManager().EquipStarup(msg.GetItemId(), msg.GetStuffItems())
}

func (m *MsgRegister) handleEquipLevelup(ctx context.Context, p ...interface{}) error {
	acct := p[0].(*player.Account)
	msg, ok := p[1].(*pbGlobal.C2S_EquipLevelup)
	if !ok {
		return errors.New("handleEquipLevelup failed: recv message body error")
	}
	pl, err := m.am.GetPlayerByAccount(acct)
	if err != nil {
		return fmt.Errorf("handleEquipLevelup.AccountExecute failed: %w", err)
	}

	return pl.ItemManager().EquipLevelup(msg.GetItemId(), msg.GetStuffItems(), msg.GetExpItems())
}

func (m *MsgRegister) handlePutonCrystal(ctx context.Context, p ...interface{}) error {
	acct := p[0].(*player.Account)
	msg, ok := p[1].(*pbGlobal.C2S_PutonCrystal)
	if !ok {
		return errors.New("handlePutonCrystal failed: recv message body error")
	}

	pl, err := m.am.GetPlayerByAccount(acct)
	if err != nil {
		return fmt.Errorf("handlePutonCrystal failed: %w", err)
	}

	if err := pl.HeroManager().PutonCrystal(msg.HeroId, msg.CrystalId); err != nil {
		return fmt.Errorf("handlePutonCrystal.AccountExecute failed: %w", err)
	}

	return nil
}

func (m *MsgRegister) handleTakeoffCrystal(ctx context.Context, p ...interface{}) error {
	acct := p[0].(*player.Account)
	msg, ok := p[1].(*pbGlobal.C2S_TakeoffCrystal)
	if !ok {
		return errors.New("handleTakeoffCrystal failed: recv message body error")
	}

	pl, err := m.am.GetPlayerByAccount(acct)
	if err != nil {
		return fmt.Errorf("handleTakeoffCrystal failed: %w", err)
	}

	if err := pl.HeroManager().TakeoffCrystal(msg.HeroId, msg.Pos); err != nil {
		return fmt.Errorf("handleTakeoffCrystal failed: %w", err)
	}

	return nil
}

func (m *MsgRegister) handleCrystalLevelup(ctx context.Context, p ...interface{}) error {
	acct := p[0].(*player.Account)
	msg, ok := p[1].(*pbGlobal.C2S_CrystalLevelup)
	if !ok {
		return errors.New("handleCrystalLevelup failed: recv message body error")
	}

	pl, err := m.am.GetPlayerByAccount(acct)
	if err != nil {
		return fmt.Errorf("handleCrystalLevelup failed: %w", err)
	}

	return pl.ItemManager().CrystalLevelup(msg.GetItemId(), msg.GetStuffItems(), msg.GetExpItems())
}

func (m *MsgRegister) handleTestCrystalRandom(ctx context.Context, p ...interface{}) error {
	acct := p[0].(*player.Account)
	msg, ok := p[1].(*pbGlobal.C2S_TestCrystalRandom)
	if !ok {
		return errors.New("handleTestCrystalRandom failed: recv message body error")
	}

	pl, err := m.am.GetPlayerByAccount(acct)
	if err != nil {
		return fmt.Errorf("handleTestCrystalRandom failed: %w", err)
	}

	return pl.ItemManager().CrystalBulkRandom(msg.CrystalRandomNum)
}
