package user

import (
	"2SOMEone/core"
	"2SOMEone/util"
	"context"
	"errors"

	"gorm.io/gorm"
)

func New(db *util.DB) core.UserStore {
	return &userStore{db: db}
}

type userStore struct {
	db *util.DB
}

func toUpdateParams(user *core.User) map[string]interface{} {
	return map[string]interface{}{
		"name":         user.Name,
		"password":     user.Password,
		"avatar":       user.Avatar,
		"access_token": user.AccessToken,
		"lang":         user.Lang,
		"role":         user.Role,
		"mixin_id":     user.MixinID,
		"user_id":      user.UserID,
		"buid":         user.Buid,
		// "email":        user.Email,
		// "phone":        user.Phone,
		// "code":         user.Code,
		// "balence":      user.Balence,
	}
}

func update(db *util.DB, user *core.User) (int64, error) {
	updates := toUpdateParams(user)
	tx := db.Update().Model(user).Where("user_id = ?", user.UserID).Updates(updates)
	return tx.RowsAffected, tx.Error
}

// func updateByEmail(db *util.DB, user *core.User) (int64, error) {
// 	updates := toUpdateParams(user)
// 	tx := db.Update().Model(user).Where("email = ?", user.Email).Updates(updates)
// 	return tx.RowsAffected, tx.Error
// }

func (s *userStore) Save(_ context.Context, user *core.User) error {
	return s.db.Tx(func(tx *util.DB) error {
		rows, err := update(tx, user)
		if err != nil {
			return err
		}

		if rows == 0 {
			return tx.Update().Create(user).Error
		}

		return nil
	})
}

// func (s *userStore) SaveByEmail(_ context.Context, user *core.User) error {
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

func (s *userStore) FindByMixinID(_ context.Context, mixinID string) (*core.User, error) {
	user := core.User{}
	if err := s.db.View().Where("mixin_id = ?", mixinID).Take(&user).Error; err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	return &user, nil
}

func (s *userStore) FindByPhone(_ context.Context, phone string) (*core.User, error) {
	user := core.User{}
	if err := s.db.View().Where("phone = ?", phone).Take(&user).Error; err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	return &user, nil
}

// func (s *userStore) FindByEmail(_ context.Context, email string) (*core.User, error) {
// 	user := core.User{}
// 	if err := s.db.View().Where("email = ?", email).Take(&user).Error; err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
// 		return nil, err
// 	}
// 	return &user, nil
// }

func (s *userStore) FindByName(_ context.Context, name string) (*core.User, error) {
	user := core.User{}
	if err := s.db.View().Where("name = ?", name).Take(&user).Error; err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	return &user, nil
}
func (s *userStore) FindByUserID(_ context.Context, user_id string) (*core.User, error) {
	user := core.User{}
	if err := s.db.View().Where("user_id = ?", user_id).Take(&user).Error; err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	return &user, nil
}

func (s *userStore) DeleteByUserID(_ context.Context, user_id string) error {
	if err := s.db.Update().Where("user_id = ?", user_id).Delete(&core.User{}).Error; err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}
	return nil
}
