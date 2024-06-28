package debt

import (
	sq "github.com/Masterminds/squirrel"

	"github.com/Lucas-Linhar3s/GerencIA/backend/server/database"
)

type SQDebt struct {
	Db *database.Database
}

func (r *SQDebt) Create(req *DebtModel) error {
	if err := r.Db.Builder.
		Insert("dividas").
		Columns(
			"id",
			"usuario_id",
			"cartao_id",
			"valor_compra",
			"parcelas",
			"descricao",
			"pago",
			"created_at",
		).
		Values(req.ID, req.UserID, req.CardID, req.ValueDebt, req.Parcels, req.Description, req.IsPaid, req.CreatedAt).
		Suffix("RETURNING id").
		Scan(new(string)); err != nil {
		return err
	}
	return nil
}

func (r *SQDebt) Find(params map[string]interface{}) ([]DebtModel, error) {
	consulta := r.Db.Builder.
		Select(
			"DIV.id",
			"DIV.descricao",
			"DIV.usuario_id",
			"DIV.cartao_id",
			"DIV.valor_compra",
			"DIV.created_at",
			"DIV.updated_at",
			"DIV.pago",
			"DIV.parcelas",
			"CAR.nome",
			"USU.nome",
			"COUNT(DIV.id) OVER() AS total",
		).
		From("dividas DIV").
		Join("cartoes CAR ON DIV.cartao_id = CAR.id").
		Join("usuarios USU ON DIV.usuario_id = USU.id")

	if params != nil {
		consulta = consulta.Where(
			sq.Or{
				sq.Eq{"DIV.usuario_id": params["UserId"]},
				sq.Eq{"DIV.cartao_id": params["CardId"]},
				sq.Expr("DIV.descricao LIKE ? COLLATE NOCASE", addConcat(params, "Description")),
				sq.Expr("CAR.nome LIKE ? COLLATE NOCASE", addConcat(params, "CardName")),
				sq.Expr("USU.nome LIKE ? COLLATE NOCASE", addConcat(params, "UserName")),
			},
		)

	}

	rows, err := consulta.Query()
	if err != nil {
		return nil, err
	}

	var debts []DebtModel
	for rows.Next() {
		var debt DebtModel
		if err := rows.Scan(
			&debt.ID,
			&debt.Description,
			&debt.UserID,
			&debt.CardID,
			&debt.ValueDebt,
			&debt.CreatedAt,
			&debt.UpdatedAt,
			&debt.IsPaid,
			&debt.Parcels,
			&debt.CardName,
			&debt.UserName,
			&debt.Total,
		); err != nil {
			return nil, err
		}
		debts = append(debts, debt)
	}

	return debts, nil
}

func (r *SQDebt) Update(req *DebtModel) error {
	_, err := r.Db.Builder.
		Update("dividas").
		SetMap(sq.Eq{
			"descricao":    req.Description,
			"parcelas":     req.Parcels,
			"valor_compra": req.ValueDebt,
			"updated_at":   req.UpdatedAt,
			"pago":         req.IsPaid,
		}).
		Where(sq.Eq{"id": req.ID}).
		Exec()
	if err != nil {
		return err
	}

	return nil
}

func (r *SQDebt) Delete(id *string) error {
	if _, err := r.Db.Builder.
		Delete("dividas").
		Where(sq.Eq{"id": id}).
		Exec(); err != nil {
		return err
	}
	return nil
}

func addConcat(params map[string]interface{}, variable string) string {
	value := params[variable].(*string)

	if value == nil {
		return ""
	}
	return "%" + *value + "%"
}
