package debt

type IDebt interface {
	Find(params map[string]interface{}) ([]DebtModel, error)
	Create(req *DebtModel) error
	Update(req *DebtModel) error
	Delete(id *string) error
}
