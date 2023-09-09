package requests

type Header struct {
	PayerPaymentRequisitionID       int       `json:"PayerPaymentRequisitionID"`
	IsCancelled		                *bool     `json:"IsCancelled"`
}
