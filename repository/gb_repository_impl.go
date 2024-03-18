package repository

import (
	"context"
	"database/sql"
	"errors"
	"goland-crud/helper"
	"goland-crud/model"
)

type GrowBoxRepositoryImpl struct {
	Db *sql.DB
}

func NewGrowBoxRepository(Db *sql.DB) GrowBoxRepository {
	return &GrowBoxRepositoryImpl{Db: Db}
}

// Save implements GrowBoxRepository.
func (b *GrowBoxRepositoryImpl) Save(ctx context.Context, GrowBox model.GrowBox) {
	tx, err := b.Db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	sql_exec := "insert into grow_box(namebox, description, count_fans, filtration, dimensions, \"Automation\") values ($1,$2,$3,$4,$5,$6)"

	_, err = tx.ExecContext(ctx, sql_exec, GrowBox.Name, GrowBox.Description, GrowBox.Count_fans, GrowBox.Filtration, GrowBox.Dimensions, GrowBox.Automation)
	helper.PanicIfError(err)
}

// Delete implements GrowBoxRepository.
func (b *GrowBoxRepositoryImpl) Delete(ctx context.Context, GrowBoxId int) {
	tx, err := b.Db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	sql_exec := "delete from grow_box where id=$1"
	_, errExec := tx.ExecContext(ctx, sql_exec, GrowBoxId)
	helper.PanicIfError(errExec)

}

// FindAll implements GrowBoxRepository.
func (b *GrowBoxRepositoryImpl) FindAll(ctx context.Context) []model.GrowBox {
	tx, err := b.Db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	sql_exec := "select * from grow_box"
	result, errQuery := tx.QueryContext(ctx, sql_exec)
	helper.PanicIfError(errQuery)
	defer result.Close()

	var GrowBoxs []model.GrowBox

	for result.Next() {
		GrowBox := model.GrowBox{}
		err := result.Scan(&GrowBox.Id, &GrowBox.Name,
			&GrowBox.Description, &GrowBox.Count_fans,
			&GrowBox.Filtration, &GrowBox.Dimensions,
			&GrowBox.Automation)
		helper.PanicIfError(err)

		GrowBoxs = append(GrowBoxs, GrowBox)
	}

	return GrowBoxs
}

// FindById implements GrowBoxRepository.
func (b *GrowBoxRepositoryImpl) FindById(ctx context.Context, GrowBoxId int) (model.GrowBox, error) {
	tx, err := b.Db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	sql_exec := "select * from grow_box where id=$1"
	result, errQuery := tx.QueryContext(ctx, sql_exec, GrowBoxId)
	helper.PanicIfError(errQuery)
	defer result.Close()

	GrowBox := model.GrowBox{}

	if result.Next() {
		err := result.Scan(&GrowBox.Id, &GrowBox.Name,
			&GrowBox.Description, &GrowBox.Count_fans,
			&GrowBox.Filtration, &GrowBox.Dimensions,
			&GrowBox.Automation)
		helper.PanicIfError(err)
		return GrowBox, nil
	} else {
		return GrowBox, errors.New("growbox id not found")
	}
}

// Update implements GrowBoxRepository.
func (b *GrowBoxRepositoryImpl) Update(ctx context.Context, GrowBox model.GrowBox) {
	tx, err := b.Db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	sql_exec := "update public.grow_box set namebox=$1, description=$2, count_fans=$3, filtration=$4, dimensions=$5, \"Automation\"=$6 where id=$7"
	_, err = tx.ExecContext(ctx, sql_exec, GrowBox.Name, GrowBox.Description,
		GrowBox.Count_fans, GrowBox.Filtration,
		GrowBox.Dimensions, GrowBox.Automation, GrowBox.Id)
	helper.PanicIfError(err)
}
