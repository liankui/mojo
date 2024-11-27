package templates

const ExampleEntityModel = `
package model

import (
	"context"
	"sync"

	"github.com/chaos-io/chaos/db"
	"github.com/chaos-io/chaos/logs"

	"github.com/liankui/mojo/example/auth/go/pkg/auth"
)

var (
	user *User
	userOnce sync.Once	
)

type User struct {
	DB *db.DB
}

func GetUserModel() *User {
	userOnce.Do(func() {
		user = NewUser()
	})

	return user
}

func NewUser() *User {
	m := &User{DB: initDB()}
	if !m.DB.Config.DisableAutoMigrate || !d.Migrator().HasTable(&auth.User{}) {
		if err := d.AutoMigrate(&auth.User{}); err != nil {
			logs.Error("Init User model error: ", err)
			panic(err)
		}
	}

	return m
}

func (m *User) Create(ctx context.Context, user *auth.User) (int64, error) {
	result := m.DB.WithContext(ctx).Create(user)
	return result.RowsAffected, result.Error
}

func (m *User) Get(ctx context.Context, id string) (*auth.User, error) {
	user := &auth.User{}
	return user, m.DB.WithContext(ctx).First(user, "id = ?", id).Error
}

func (m *User) Delete(ctx context.Context, id string) (int64, error) {
	result := m.DB.WithContext(ctx).Where("id = ?", id).Delete(&auth.User{})
	return result.RowsAffected, result.Error
}

func (m *User) Update(ctx context.Context, user *auth.User) (int64, error) {
	result := m.DB.WithContext(ctx).Updates(user)
	return result.RowsAffected, result.Error
}

func (m *User) List(ctx context.Context, filter string, condition ...string) ([]*auth.User, error) {
	var user []*auth.User

	tx := m.DB.WithContext(ctx)
	// todo add condition

	return user, tx.Find(&user).Error
}

func (m *User) BatchCreate(ctx context.Context, user ...*auth.User) (int64, error) {
	result := m.DB.WithContext(ctx).CreateInBatches(user, len(user))
	return result.RowsAffected, result.Error
}

func (m *User) BatchGet(ctx context.Context, ids ...string) ([]*auth.User, error) {
	var user []*auth.User
	return user, m.DB.WithContext(ctx).Find(&user, ids).Error
}

func (m *User) BatchDelete(ctx context.Context, ids ...string) (int64, error) {
	result := m.DB.WithContext(ctx).Where("id = ?", ids).Delete(&auth.User{})
	return result.RowsAffected, result.Error
}

func (m *User) BatchUpdate(ctx context.Context, user []*auth.User) (int64, error) {
	result := m.DB.WithContext(ctx).Updates(user)
	return result.RowsAffected, result.Error
}
`
