package models

type WorkOrder struct {
	Id        int64  `json:"id" form:"id"`
	Partner   string `json:"partner"`
	Merchants string `json:"merchants"`
	Contacts  string `json:"contacts"`
	Phone     string `json:"phone"`
}

func OrderList() []map[string]string {
	orders, _ := ORM.QueryString("select * from work_order")

	return orders
}
