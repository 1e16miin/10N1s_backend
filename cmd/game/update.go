package game

import (
	"context"

	"github.com/10n1s-backend/cmd/model"
)

func (g *GameHandler) Update(ctx context.Context, game *model.Game) error {
	// tx
	if err := validateUpdate(ctx, game); err != nil {
		return err
	}

	// data 가공
	if err := g.updateGameToDB(ctx, game); err != nil {
		return err
	}

	// commit
	return nil
}

func validateUpdate(ctx context.Context, game *model.Game) error {
	return nil
}

func (g *GameHandler) updateGameToDB(ctx context.Context, game *model.Game) error {
	tx := g.db.Begin()
	isCommitted := false
	defer func() {
		if !isCommitted {
			tx.Rollback()
		}
	}()

	tx.Model(&model.Game{}).Updates(game)

	tx.Commit()
	if tx.Error != nil {
		return tx.Error
	}
	isCommitted = true
	return tx.Error
}
