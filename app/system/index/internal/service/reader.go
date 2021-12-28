package service

import "beegoframe/app/dao"

var Reader = new(readerService)

type readerService struct{}

func (s *readerService) Create() error {
	return nil
}

func (s *readerService) Delete() error {
	return nil
}

func (s *readerService) Update() error {
	return nil
}

func (s *readerService) Get() (interface{}, error) {
	var cont []struct {
		Name            string `json:"name"`
		Gender          string `json:"gender"`
		Email           string `json:"email" alias:"t1"`
		AccountPassword string `json:"account_password" alias:"t1" field:"password"`
	}
	readerIds := []int64{1, 2, 3}
	dao.Reader.
		LeftJoin(dao.Account, dao.AliasT1).On(dao.Reader.Fields.AccountId).
		Where(dao.Reader.Fields.Id.Alias().In(len(readerIds)), readerIds).
		And(dao.Account.Fields.Email.AliasT2().Greater(), "").
		Limit(10).
		Rows(&cont)
	return cont, nil
}
