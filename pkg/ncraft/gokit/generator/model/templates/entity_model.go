package templates

const EntityModel = `
package model

import (
    "context"
	"sync"

	"github.com/chaos-io/chaos/logs"
	"github.com/chaos-io/chaos/db"

	"{{.Go.ImportPath}}"
)

var {{ToLowerCamel .Name}}Model *{{.Name}}Model
var {{ToLowerCamel .Name}}ModelOnce sync.Once

type {{.Name}}Model struct {
	DB *db.DB
}

func Get{{.Name}}Model() *{{.Name}}Model {
	{{ToLowerCamel .Name}}ModelOnce.Do(func() {
		{{ToLowerCamel .Name}}Model = New{{.Name}}Model()
	})

	return {{ToLowerCamel .Name}}Model
}

func New{{.Name}}Model() *{{.Name}}Model {
	m := &{{.Name}}Model{DB: initDB()}
	if !m.DB.Config.DisableAutoMigrate || !d.Migrator().HasTable(&{{.Go.PackageName}}.{{.Name}}{}) {
		if err := d.AutoMigrate(&{{.Go.PackageName}}.{{.Name}}{}); err != nil {
			logs.Error("Init {{.Name}}Model model err: ", err)
			panic(err)
		}
	}

	return m
}

func (m *{{.Name}}Model) Create(ctx context.Context, {{ToLowerCamel .Name}} *{{.Go.PackageName}}.{{.Name}}) (int64, error) {
	result := m.DB.WithContext(ctx).Create({{ToLowerCamel .Name}})
	return result.RowsAffected, result.Error
}

func (m *{{.Name}}Model) Get(ctx context.Context, id string) (*{{.Go.PackageName}}.{{.Name}}, error) {
	{{ToLowerCamel .Name}} := &{{.Go.PackageName}}.{{.Name}}{}
	return {{ToLowerCamel .Name}}, m.DB.WithContext(ctx).First({{ToLowerCamel .Name}}, "id = ?", id).Error
}

func (m *{{.Name}}Model) Delete(ctx context.Context, id string) (int64, error) {
	result := m.DB.WithContext(ctx).Where("id = ?", id).Delete(&{{.Go.PackageName}}.{{.Name}}{})
	return result.RowsAffected, result.Error
}

func (m *{{.Name}}Model) Update(ctx context.Context, {{ToLowerCamel .Name}} *{{.Go.PackageName}}.{{.Name}}) (int64, error) {
	result := m.DB.WithContext(ctx).Updates({{ToLowerCamel .Name}})
	return result.RowsAffected, result.Error
}

func (m *{{.Name}}Model) List(ctx context.Context, filter string, condition ...string) ([]*{{.Go.PackageName}}.{{.Name}}, error) {
	var {{ToLowerCamel .Name}} []*{{.Go.PackageName}}.{{.Name}}

	tx := m.DB.WithContext(ctx)
	// todo add condition

	return {{ToLowerCamel .Name}}, tx.Find(&{{ToLowerCamel .Name}}).Error
}

func (m *{{.Name}}Model) BatchCreate(ctx context.Context, {{ToLowerCamel .Name}} ...*{{.Go.PackageName}}.{{.Name}}) (int64, error) {
	result := m.DB.WithContext(ctx).CreateInBatches({{ToLowerCamel .Name}}, len({{ToLowerCamel .Name}}))
	return result.RowsAffected, result.Error
}

func (m *{{.Name}}Model) BatchGet(ctx context.Context, ids ...string) ([]*{{.Go.PackageName}}.{{.Name}}, error) {
	var {{ToLowerCamel .Name}} []*{{.Go.PackageName}}.{{.Name}}
	return {{ToLowerCamel .Name}}, m.DB.WithContext(ctx).Find(&{{ToLowerCamel .Name}}, ids).Error
}

func (m *{{.Name}}Model) BatchDelete(ctx context.Context, ids ...string) (int64, error) {
	result := m.DB.WithContext(ctx).Where("id = ?", ids).Delete(&{{.Go.PackageName}}.{{.Name}}{})
	return result.RowsAffected, result.Error
}

func (m *{{.Name}}Model) BatchUpdate(ctx context.Context, {{ToLowerCamel .Name}} []*{{.Go.PackageName}}.{{.Name}}) (int64, error) {
	result := m.DB.WithContext(ctx).Updates({{ToLowerCamel .Name}})
	return result.RowsAffected, result.Error
}
`
