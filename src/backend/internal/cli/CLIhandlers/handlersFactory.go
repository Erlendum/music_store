package CLIhandlers

import "backend/internal/services"

type HandlersServicesFields struct {
	CalcDiscountService   services.CalcDiscountService
	ComparisonListService services.ComparisonListService
	DiscountService       services.DiscountService
	InstrumentService     services.InstrumentService
	UserService           services.UserService
}

type Handlers struct {
	InstrumentHandler     *InstrumentHandler
	UserHandler           *UserHandler
	DiscountHandler       *DiscountHandler
	ComparisonListHandler *ComparisonListHandler
}

func NewHandlers(services HandlersServicesFields) *Handlers {
	return &Handlers{InstrumentHandler: NewInstrumentHandler(services.InstrumentService),
		UserHandler:           NewUserHandler(services.UserService),
		DiscountHandler:       NewDiscountHandler(services.DiscountService),
		ComparisonListHandler: NewComparisonListHandler(services.ComparisonListService)}
}
