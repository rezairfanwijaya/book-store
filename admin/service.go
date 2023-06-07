package admin

type IService interface {
	Login(input InputAdminLogin) (Admin, int, error)
}

type service struct {
	repoAdmin IRepository
}

func NewService(repoAdmin IRepository) *service {
	return &service{repoAdmin}
}
