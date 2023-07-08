package game

// func (g *gameHandler) get(ctx context.Context, id int) (*model.Game, error) {
// 	tx := g.db.Begin()
// 	isCommitted := false
// 	defer func() {
// 		if !isCommitted {
// 			tx.Rollback()
// 		}
// 	}()

// 	var targetGame model.Game

// 	tx.Model(&model.Game{}).Where("id = ?", id).Scan(&targetGame)

// 	tx.Commit()

// 	if tx.Error != nil {
// 		return nil, tx.Error
// 	}
// 	isCommitted = true

// 	return &targetGame, nil
// }

// func (g *GameHandler) GetByUserID(ctx context.Context, userID int) (*model.Game, error) {
// 	tx := g.db.Begin()
// 	isCommitted := false
// 	defer func() {
// 		if !isCommitted {
// 			tx.Rollback()
// 		}
// 	}()

// 	var targetGame model.Game

// 	tx.Model(&model.Game{}).Where("host_id = ?", userID).
// 		Or("guest_id1 = ?", userID).
// 		Or("guest_id2 = ?", userID).
// 		Or("guest_id3 = ?", userID).
// 		Scan(targetGame)

// 	tx.Commit()

// 	if tx.Error != nil {
// 		return nil, tx.Error
// 	}
// 	isCommitted = true

// 	return &targetGame, nil
// }
