package query

type GetAccountQuery struct {
	Id string
}

func NewGetAccountQuery(id string) *GetAccountQuery {
	return &GetAccountQuery{
		Id: id,
	}
}
