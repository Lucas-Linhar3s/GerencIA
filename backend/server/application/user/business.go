package user

import (
	"github.com/gin-gonic/gin"

	"github.com/Lucas-Linhar3s/GerencIA/backend/server/di"
	"github.com/Lucas-Linhar3s/GerencIA/backend/server/domain/user"
	"github.com/Lucas-Linhar3s/GerencIA/backend/server/responses"
)

func Find(ctx *gin.Context, params map[string]interface{}) (*[]UserRes, error) {
	const msg = "Error finding users"
	var (
		service = user.GetService(user.GetRepository(di.GetDatabase()))
	)

	res, err := service.Find(params)
	responses.CheckError(msg, err)

	if len(res) == 0 {
		return nil, responses.ErrNotFound
	}

	var resp = make([]UserRes, len(res))
	for r := range res {
		resp[r] = UserRes{
			ID:        res[r].ID,
			Nome:      res[r].Name,
			Whatsapp:  res[r].Whatsapp,
			CreatedAt: res[r].CreatedAt,
			UpdatedAt: res[r].UpdatedAt,
		}
	}

	return &resp, nil
}

func Create(ctx *gin.Context, req *UserReq) error {
	const msg = "Error creating user"
	var (
		service = user.GetService(user.GetRepository(di.GetDatabase()))
	)

	user := user.UserModel{
		Name:     &req.Nome,
		Whatsapp: &req.Whatsapp,
	}

	err := service.Create(&user)
	responses.CheckError(msg, err)

	return nil
}

func Update(ctx *gin.Context, req *UserReq) error {
	const msg = "Error updating user"
	var (
		service = user.GetService(user.GetRepository(di.GetDatabase()))
	)

	user := user.UserModel{
		ID:       &req.ID,
		Name:     &req.Nome,
		Whatsapp: &req.Whatsapp,
	}

	err := service.Update(&user)
	responses.CheckError(msg, err)

	return nil
}

func Delete(ctx *gin.Context, id *string) error {
	const msg = "Error deleting user"
	var (
		service = user.GetService(user.GetRepository(di.GetDatabase()))
	)

	err := service.Delete(id)
	responses.CheckError(msg, err)

	return nil
}
