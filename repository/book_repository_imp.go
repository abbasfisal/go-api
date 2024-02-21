package repository

import (
	"context"
	"database/sql"
	"github.com/abbasfisal/go-api/helper"
	"github.com/abbasfisal/go-api/model"
)

type BookRepositoryImpl struct {
	Db *sql.DB
}

func NewBookRepository(Db *sql.DB) BookRepository {
	return &BookRepositoryImpl{Db: Db}
}

func (b *BookRepositoryImpl) Save(ctx context.Context, book model.Book) {
	//TODO implement me
	panic("implement me")
}

func (b *BookRepositoryImpl) Update(ctx context.Context, book model.Book) {
	//TODO implement me
	panic("implement me")
}

func (b *BookRepositoryImpl) Delete(ctx context.Context, bookId int) {
	tx, err := b.Db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	SQL := `delete from books where id=$1`
	_, errExec := tx.ExecContext(ctx, SQL, bookId)
	helper.PanicIfError(errExec)

}

func (b *BookRepositoryImpl) FindById(ctx context.Context, bookId int) (model.Book, error) {
	//TODO implement me
	panic("implement me")
}

func (b *BookRepositoryImpl) FindAll(ctx context.Context) []model.Book {
	//TODO implement me
	panic("implement me")
}
