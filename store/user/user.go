package user

import (
	"2SOMEone/core"
	"2SOMEone/util"
	"context"
	"errors"

	"gorm.io/gorm"
)

func NewUserStore(db *util.DB) core.UserStore {
	return &userStore{db: db}
}

type userStore struct {
	db *util.DB
}

func toUpdateParams(user *core.BasicUser) map[string]interface{} {
	return map[string]interface{}{
		"user_id":  user.UserID,
		"name":     user.Name,
		"password": user.Password,
		"avatar":   user.Avatar,
		"lang":     user.Lang,
		"email":    user.Email,
		"phone":    user.Phone,
	}
}

func update(db *util.DB, user *core.BasicUser) (int64, error) {
	updates := toUpdateParams(user)
	tx := db.Update().Model(user).Where("user_id = ?", user.UserID).Updates(updates)
	return tx.RowsAffected, tx.Error
}

func updateByPhone(db *util.DB, user *core.BasicUser) (int64, error) {
	updates := toUpdateParams(user)
	tx := db.Update().Model(user).Where("phone = ?", user.Phone).Updates(updates)
	return tx.RowsAffected, tx.Error
}

// func updateByEmail(db *util.DB, user *core.BasicUser) (int64, error) {
// 	updates := toUpdateParams(user)
// 	tx := db.Update().Model(user).Where("email = ?", user.Email).Updates(updates)
// 	return tx.RowsAffected, tx.Error
// }

func (s *userStore) Save(_ context.Context, user *core.BasicUser) error {
	return s.db.Tx(func(tx *util.DB) error {
		var rows int64
		var err error
		if user.UserID == "" {
			rows, err = updateByPhone(tx, user)
		} else {
			rows, err = update(tx, user)
		}
		if err != nil {
			return err
		}

		if rows == 0 {
			return tx.Update().Create(user).Error
		}

		return nil
	})
}

// func (s *userStore) SaveByEmail(_ context.Context, user *core.BasicUser) error {
// 	return s.db.Tx(func(tx *util.DB) error {
// 		// rows, err := update(tx, user)
// 		rows, err := updateByEmail(tx, user)
// 		if err != nil {
// 			return err
// 		}
// 		if rows == 0 {
// 			return tx.Update().Create(user).Error
// 		}
// 		return nil
// 	})
// }

func (s *userStore) FindByMixinID(_ context.Context, mixinID string) (*core.BasicUser, error) {
	user := core.BasicUser{}
	if err := s.db.View().Where("mixin_id = ?", mixinID).Take(&user).Error; err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	} else if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	return &user, nil
}

func (s *userStore) FindByPhone(_ context.Context, phone string) (*core.BasicUser, error) {
	user := core.BasicUser{}
	if err := s.db.View().Where("phone = ?", phone).Take(&user).Error; err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	} else if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &user, nil
}

// func (s *userStore) FindByEmail(_ context.Context, email string) (*core.BasicUser, error) {
// 	user := core.BasicUser{}
// 	if err := s.db.View().Where("email = ?", email).Take(&user).Error; err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
// 		return nil, err
// 	}
// 	return &user, nil
// }

func (s *userStore) FindByName(_ context.Context, name string) (*core.BasicUser, error) {
	user := core.BasicUser{}
	if err := s.db.View().Where("name = ?", name).Take(&user).Error; err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	} else if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	return &user, nil
}
func (s *userStore) FindByUserID(_ context.Context, user_id string) (*core.BasicUser, error) {
	user := core.BasicUser{}
	if err := s.db.View().Where("user_id = ?", user_id).Take(&user).Error; err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	} else if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	return &user, nil
}

func (s *userStore) DeleteByUserID(_ context.Context, user_id string) error {
	if err := s.db.Update().Where("user_id = ?", user_id).Delete(&core.BasicUser{}).Error; err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}
	return nil
}
