package request


type SendMessageOneToOne struct {
	ReceiverMobileNumber string `json:"receiver_mobile_number" binding:"required"`
	Message              string `json:"message" binding:"required"`
}

type CreateChat struct {
	ReceiverMobileNumber string `json:"receiver_mobile_number" binding:"required"`
}
