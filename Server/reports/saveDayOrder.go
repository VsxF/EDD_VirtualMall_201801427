package reports

import (
	"fmt"
	"strconv"
	ord "../data/orders"
)
// "19-05-2019"
//30-11-2018

func GetDayOrders(orders *ord.Calendar, date string) ord.AuxOrders {
	aux := []byte(date)
	day2search, _ := strconv.Atoi(string(aux[:2]))
	month2search, _ := strconv.Atoi(string(aux[3]) + string(aux[4]))
	year2search, _ := strconv.Atoi(string(aux[6:]))

	fmt.Println("-")
	var response ord.AuxOrders
	searchAVL(orders.Root, year2search, month2search, day2search, &response)
	return response
}

func searchAVL(year *ord.Year, y2s, m2s, d2s int, orders *ord.AuxOrders) {
	if y2s < year.Year {
		searchAVL(year.Left, y2s, m2s, d2s, orders)
	} else if y2s > year.Year {
		searchAVL(year.Right, y2s, m2s, d2s, orders)
	} else {
		*orders = searchMonth(year.Months, m2s, d2s)
	}
}

func searchMonth(months *ord.Months, m2s, d2s int) ord.AuxOrders {
	next := months.Start
	prev := months.Lastest
	monthSearched := next

	for i := 0; i < months.Size; i++ {
		if next.Month == m2s {
			monthSearched = next
			break
		} else if prev.Month == m2s {
			monthSearched = prev
			break
		}
		next = next.Next
		prev = prev.Previous
	}
	return searchDayOrders(monthSearched.Orders, d2s)
}

//Recorrer Matrix
func searchDayOrders(orders *ord.Orders, d2s int) ord.AuxOrders {
	x0 := orders.HeaderX.First //Primer dia eje X
	xf := orders.HeaderX.Last //Ultimo dia eje X

	delta0 := d2s-x0.Order
	deltaf := xf.Order-d2s

	if delta0 <= deltaf {
		return searchHeaderX(x0, true, d2s)
	} else {
		return searchHeaderX(xf, false, d2s)
	}	
}

func searchHeaderX(header *ord.Header, next bool, d2s int) ord.AuxOrders {
	var orders ord.AuxOrders
	if next {
		for header != nil {
			if header.Order == d2s {
				orders.Orders = append(orders.Orders, *header.InnerList.First)
			} else if header.Order < d2s {
				break
			}
			header = header.Next
		}
	} else {
		for header != nil {
			if header.Order == d2s {
				orders.Orders = append(orders.Orders, *header.InnerList.First)
			} else if header.Order > d2s {
				break
			}
			header = header.Previous
		}
	}
	return orders
}


