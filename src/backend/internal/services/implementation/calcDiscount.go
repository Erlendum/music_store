package servicesImplementation

import (
	"backend/internal/models"
	"backend/internal/pkg/errors/repositoryErrors"
	"backend/internal/repository"
	"backend/internal/services"
	"strconv"
	"strings"
	"time"
)

const (
	PERCENT     = "Процентная"
	BIRTH       = "Именинная"
	MALE        = "Мужской"
	FEMALE      = "Женский"
	GIFT        = "Подарочная"
	GiftNumElem = 3
)

type calcDiscountServiceImplementation struct {
	discountRepository repository.DiscountRepository
}

func NewCalcDiscountServiceImplementation(discountRepository repository.DiscountRepository) services.CalcDiscountService {
	return &calcDiscountServiceImplementation{
		discountRepository: discountRepository,
	}
}

func (c *calcDiscountServiceImplementation) CalcDiscount(user *models.User, instruments []models.Instrument) ([]models.Instrument, error) {
	for i, instrument := range instruments {
		discounts, err := c.discountRepository.GetSpecificList(instrument.InstrumentId, user.UserId)
		if err != nil && err == repositoryErrors.ObjectDoesNotExists {
			return instruments, err
		} else if err != nil {
			return nil, err
		}
		var maxPercent uint64 = 0
		count := c.countOfInstruments(instruments, instrument.InstrumentId)
		maxN := -1
		maxM := -1
		dateNow := time.Now()
		for _, discount := range discounts {
			elems := strings.Split(discount.Type, " ")
			switch elems[0] {
			case PERCENT, BIRTH:
				if dateNow.Before(discount.DateEnd) && dateNow.After(discount.DateBegin) && discount.Amount > maxPercent {
					_, m1, d1 := dateNow.Date()
					_, m2, d2 := discount.DateBegin.Date()
					if elems[0] == BIRTH && (m1 != m2 || d1 != d2) {
						continue
					}
					maxPercent = discount.Amount
				}
			case MALE, FEMALE:
				if dateNow.Before(discount.DateEnd) && dateNow.After(discount.DateBegin) && discount.Amount > maxPercent && user.Gender == models.UserGender(discount.Type) {
					maxPercent = discount.Amount
				}
			case GIFT:
				c.calcGift(count, elems, &maxN, &maxM, dateNow, discount)
			}
		}
		if maxPercent > 0 {
			instruments[i].Price -= instruments[i].Price * maxPercent / 100.0
		}
		for k := 0; k < maxM; k++ {
			newInstrument := instrument
			newInstrument.Price = 0
			instruments = append(instruments, newInstrument)
		}
	}
	return instruments, nil
}

func (c *calcDiscountServiceImplementation) calcGift(count int, elems []string, maxN, maxM *int, dateNow time.Time, discount models.Discount) {
	if len(elems) != GiftNumElem {
		return
	}
	n, err := strconv.Atoi(elems[1])
	if err != nil {
		return
	}
	m, err := strconv.Atoi(elems[2])
	if err != nil {
		return
	}

	if dateNow.Before(discount.DateEnd) && dateNow.After(discount.DateBegin) && count >= n && n > *maxN {
		*maxN = n
		*maxM = m
	}

}
func (c *calcDiscountServiceImplementation) countOfInstruments(instruments []models.Instrument, instrumentId uint64) int {
	k := 0
	for _, instrument := range instruments {
		if instrument.InstrumentId == instrumentId {
			k++
		}
	}
	return k
}
