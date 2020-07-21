package repositories

import (
	"database/sql"
	"fmt"
	"loment/models"

	// nothing
	_ "github.com/go-sql-driver/mysql"
	"xorm.io/xorm"
	"xorm.io/xorm/names"
)

// CommentRepository repo
type CommentRepository struct {
	engine     *xorm.Engine
	dataSource string
	dbName     string
}

// Create repo
func Create(dataSource string, dbName string) *CommentRepository {
	repo := new(CommentRepository)
	repo.dbName = dbName
	repo.dataSource = dataSource
	repo.engine = nil
	return repo
}

// Start repo engine
func (repo *CommentRepository) Start(isDebug bool) error {
	engine, err := xorm.NewEngine("mysql", fmt.Sprintf("%s/%s", repo.dataSource, repo.dbName))
	if err != nil {
		return err
	}
	if isDebug {
		engine.ShowSQL(true)
	}
	engine.SetTableMapper(names.SameMapper{})
	engine.SetColumnMapper(names.SameMapper{})
	engine.Sync2(new(models.Comment))
	repo.engine = engine
	return err
}

// Stop repo engine
func (repo *CommentRepository) Stop() error {
	err := repo.engine.Close()
	return err
}

// EnsureExisits db exisis
func (repo *CommentRepository) EnsureExisits() error {
	db, err := sql.Open("mysql", fmt.Sprintf("%s/", repo.dataSource))
	if err != nil {
		return err
	}
	defer db.Close()
	_, err = db.Exec(fmt.Sprintf("create database if not exists %s", repo.dbName))
	return err
}

// Create new comment
func (repo *CommentRepository) Create(obj *models.Comment) error {
	_, err := repo.engine.Insert(obj)
	return err
}

// Get by id
func (repo *CommentRepository) Get(id string) (*models.Comment, error) {
	var result models.Comment
	has, err := repo.engine.Where("Id = ?", id).Get(&result)
	if has {
		return &result, err
	}
	return nil, err
}

// Update by id
func (repo *CommentRepository) Update(obj *models.Comment) error {
	_, err := repo.engine.Where("Id = ?", obj.Id).Update(obj)
	return err
}

// Delete by id
func (repo *CommentRepository) Delete(id string) (*models.Comment, error) {
	var obj models.Comment
	_, err := repo.engine.Where("Id = ?", id).Delete(&obj)
	return &obj, err
}

// Query comments
func (repo *CommentRepository) Query(query *models.CommentQuery) ([]models.Comment, error) {
	session := repo.engine.NewSession()
	if query.Id != "" {
		session = session.Where("Id = ?", query.Id)
	}
	if query.Uri != "" {
		session = session.Where("Uri = ?", query.Uri)
	}
	if query.Author != "" {
		session = session.Where("Author = ?", query.Author)
	}
	if query.Email != "" {
		session = session.Where("Email = ?", query.Email)
	}
	if query.Limit == 0 {
		query.Limit = 10
	}
	var result []models.Comment
	err := session.Limit(query.Limit, query.Offset).Find(&result)
	return result, err
}
