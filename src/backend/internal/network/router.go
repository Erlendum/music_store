package network

import (
	"backend/internal/network/handlers"
	"net/http"
)

func NewRouter(h *handlers.Handlers) *http.ServeMux {
	mux := &http.ServeMux{}
	mux.HandleFunc("/create_instrument", http.HandlerFunc(h.InstrumentHandler.Create))
	mux.HandleFunc("/instruments", http.HandlerFunc(h.InstrumentHandler.GetList))
	mux.HandleFunc("/delete_instrument", http.HandlerFunc(h.InstrumentHandler.Delete))
	mux.HandleFunc("/update_instrument", http.HandlerFunc(h.InstrumentHandler.Update))
	mux.HandleFunc("/create_user", http.HandlerFunc(h.UserHandler.Create))
	mux.HandleFunc("/get_user", http.HandlerFunc(h.UserHandler.Get))
	mux.HandleFunc("/comparison_list", http.HandlerFunc(h.UserHandler.GetComparisonList))
	mux.HandleFunc("/add_instrument_to_comparison_list", http.HandlerFunc(h.ComparisonListHandler.AddInstrument))
	mux.HandleFunc("/delete_instrument_from_comparison_list", http.HandlerFunc(h.ComparisonListHandler.DeleteInstrument))
	mux.HandleFunc("/create_discount", http.HandlerFunc(h.DiscountHandler.Create))
	mux.HandleFunc("/create_for_all_discount", http.HandlerFunc(h.DiscountHandler.CreateForAll))
	mux.HandleFunc("/discounts", http.HandlerFunc(h.DiscountHandler.GetList))
	mux.HandleFunc("/delete_discount", http.HandlerFunc(h.DiscountHandler.Delete))
	mux.HandleFunc("/update_discount", http.HandlerFunc(h.DiscountHandler.Update))
	mux.HandleFunc("/create_order", http.HandlerFunc(h.OrderHandler.Create))
	mux.HandleFunc("/orders", http.HandlerFunc(h.OrderHandler.GetList))
	mux.HandleFunc("/ordersForAll", http.HandlerFunc(h.OrderHandler.GetListForAll))
	mux.HandleFunc("/update_order", http.HandlerFunc(h.OrderHandler.Update))
	mux.HandleFunc("/order_elements", http.HandlerFunc(h.OrderHandler.GetOrderElements))
	return mux
}
