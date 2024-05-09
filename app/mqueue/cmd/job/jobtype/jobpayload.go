package jobtype

import "looklook/app/order/model"

// DeferCloseHomestayOrderPayload defer close homestay order
type DeferCloseHomestayOrderPayload struct {
	Sn string
}

type DeferCloseTiktokRecord struct {
	ID int64
}
type DeferOpenTiktokRecord struct {
	ID int64
}

// PaySuccessNotifyUserPayload pay success notify user
type PaySuccessNotifyUserPayload struct {
	Order *model.HomestayOrder
}
