package orders

//import "fmt"

type Month struct {
	Next *Month
	Previous *Month
	Name string
	Month int
	Orders *Orders
	HasOrders bool
}

type Months struct {
	Start *Month
	Lastest *Month
	Size int
}

func NewMonth() *Month {
	return &Month{}
}

func NewMonthsList() *Months {
	start, last := createCalendar()
	return &Months{start, last, 12}
}

func createCalendar() (*Month, *Month) {
	dic := &Month{nil, nil, "Diciembre", 12, NewOrders(), false}
	nov := &Month{dic, nil, "Noviembre", 11, NewOrders(), false}
	oct := &Month{nov, nil, "Octubre", 10, NewOrders(), false}
	sep := &Month{oct, nil, "Septiembre", 9, NewOrders(), false}
	ago := &Month{sep, nil, "Agosto", 8, NewOrders(), false}
	jul := &Month{ago, nil, "Julio", 7, NewOrders(), false}
	jun := &Month{jul, nil, "Junio", 6, NewOrders(), false}
	may := &Month{jun, nil, "Mayo", 5, NewOrders(), false}
	abr := &Month{may, nil, "Abril", 4, NewOrders(), false}
	mar := &Month{abr, nil, "Marzo", 3, NewOrders(), false}
	feb := &Month{mar, nil, "Febrero", 2, NewOrders(), false}
	ene := &Month{feb, nil, "Enero", 1, NewOrders(), false}

	feb.Previous = ene
	mar.Previous = feb
	abr.Previous = mar
	may.Previous = abr
	jun.Previous = may
	jul.Previous = jun
	ago.Previous = jul
	sep.Previous = ago
	oct.Previous = sep
	nov.Previous = oct
	dic.Previous = nov

	return ene, dic
}

func (months *Months) setMonth(month int, orders *Orders) {
	if orders != nil {
		aux := orders.HeaderX.First.InnerList.First	
		next := months.Start
		
		for next.Month != month {
			next = next.Next
		}
		
		InsertarOrden(next.Orders, aux.Date, aux.Store, aux.Department, aux.Qualification, aux.Products)
		next.HasOrders = true
	}		
}

	

func (Months *Months) AddMonths(st *Months) {
	if Months.Size > 0 {
		if st.Size > 0 {
			st.Start.Previous = Months.Lastest
			*Months.Lastest.Next = *st.Start
			
			*Months.Lastest = *st.Lastest
		}
	} else {
		*Months = *st
	}
}

// func (vt *Vector) DeleteMonth(delete Month, index int) bool {
// 	if vt.Vector[index].Months.Size > 0 {
// 		Month := vt.Vector[index].Months.Start
		
// 		for Month != nil {
// 			name := strings.ToLower(Month.Name)
// 			delname := strings.ToLower(delete.Name)
			
// 			if name == delname && delete.Qualification == Month.Qualification {
				
// 				if Month.Next == nil && Month.Previous == nil {
// 					vt.Vector[index].Months = &Months{nil, nil, 0}

// 				} else if Month.Next != nil && Month.Previous != nil {
// 					Month.Previous.Next = Month.Next
// 					Month.Next.Previous = Month.Previous
// 					vt.Vector[index].Months.Size--

// 				} else  if Month.Next != nil {
// 					Month.Next.Previous = nil
// 					vt.Vector[index].Months.Start = Month.Next
// 					vt.Vector[index].Months.Size--

// 				} else if Month.Previous != nil {
// 					Month.Previous.Next = nil
// 					vt.Vector[index].Months.Lastest = Month.Previous
// 					vt.Vector[index].Months.Size--
// 				} 
// 				return true
// 			} 
// 			Month = Month.Next
// 		}
// 	}
// 	return false
// }
