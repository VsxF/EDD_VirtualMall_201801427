package orders

import "testing"

func TestOrders_Insertar(t *testing.T) {
	type fields struct {
		Layer   int
		HeaderX *HeaderList
		HeaderY *HeaderList
	}
	type args struct {
		date          string
		store         string
		department    string
		qualification int
		products      []ProductsCodes
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Orders{
				Layer:   tt.fields.Layer,
				HeaderX: tt.fields.HeaderX,
				HeaderY: tt.fields.HeaderY,
			}
			m.Insertar(tt.args.date, tt.args.store, tt.args.department, tt.args.qualification, tt.args.products)
		})
	}
}
