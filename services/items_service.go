package users_services

var (
	ItemService itemServicesInterface = &ItemServices{}
)

type ItemServices struct {
	itemServicesInterface
}

type itemServicesInterface interface {
	GetItem()
	ChangeItem()
}

func (item *ItemServices) GetItem() {

}

func (item *ItemServices) ChangeItem() {

}