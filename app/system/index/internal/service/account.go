package service

import "beegoframe/app/dao"

var Account = new(accountService)

type accountService struct{}

func (s *accountService) Create() error {
	data := &struct {
		Email    string `json:"email"`
		Phone    string `json:"phone"`
		Password string `json:"password"`
	}{
		Email:    "example@qq.com",
		Phone:    "13282113456",
		Password: "123456",
	}
	if _, err := dao.Account.SetData(data).Insert(); err != nil {
		return err
	}
	return nil
}

func (s *accountService) Delete() error {
	if _, err := dao.Account.Filter(dao.Account.Fields.Id.Exact(), 1).Delete(); err != nil {
		return err
	}
	return nil
}

func (s *accountService) Update() error {
	data := &struct {
		Password string `json:"password"`
	}{
		Password: "654321",
	}

	// email = 'example@qq.com' OR phone = '13282113456'
	cond := dao.NewCond().
		Or(dao.Account.Fields.Email.Exact(), "example@qq.com").
		Or(dao.Account.Fields.Phone.Exact(), "13282113456")

	// deleted = true
	cond2 := dao.NewCond().
		And(dao.Account.Fields.Deleted.Exact(), false)

	// deleted = true AND ( email = 'example@qq.com' OR phone = '13282113456' )
	cond3 := cond2.AndCond(cond)

	// UPDATE FROM account SET password = '654321' WHERE ...^
	if _, err := dao.Account.SetCond(cond3).SetData(data).Update(); err != nil {
		return err
	}
	return nil
}

func (s *accountService) Get() (interface{}, error) {
	var cont []struct {
		Phone      string `json:"phone"`
		Password   string `json:"password"`
		CreateTime string `json:"create_time"`
	}
	// SELECT phone, password, create_time FROM account WHERE phone LIKE '%132%';
	if _, err := dao.Account.Filter(dao.Account.Fields.Phone.Contains(), "132").Rows(&cont); err != nil {
		return nil, err
	}

	return nil, nil
}
