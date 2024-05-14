package couponService

import (
	"coupon_service/internal/repository/memdb"
	"coupon_service/internal/service/entity"
	"fmt"
	"reflect"
	"testing"
)

type mockRepository struct {
	findByCodeFunc func(string) (*entity.Coupon, error)
}

func (m *mockRepository) FindByCode(code string) (*entity.Coupon, error) {
	return m.findByCodeFunc(code)
}

func (m *mockRepository) Save(coupon entity.Coupon) error {
	return nil
}

func TestCouponNew(t *testing.T) {
	type args struct {
		repo Repository
	}
	tests := []struct {
		name string
		args args
		want CouponService
	}{
		{"initialize service", args{repo: nil}, CouponService{repo: nil}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.repo); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCouponService_ApplyDiscount(t *testing.T) {
	type fields struct {
		repo Repository
	}
	type args struct {
		basket entity.Basket
		code   string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantB   *entity.Basket
		wantErr bool
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := CouponService{
				repo: tt.fields.repo,
			}
			gotB, err := s.UpdateBasket(tt.args.basket, tt.args.code)
			if (err != nil) != tt.wantErr {
				t.Errorf("ApplyDiscount() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotB, tt.wantB) {
				t.Errorf("ApplyDiscount() gotB = %v, want %v", gotB, tt.wantB)
			}
		})
	}
}

func TestCouponService_CreateCoupon(t *testing.T) {
	type fields struct {
		repo Repository
	}
	type args struct {
		discount       int
		code           string
		minBasketValue int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   any
	}{
		{"Apply 10%", fields{memdb.New()}, args{10, "SuperDiscount", 55}, nil},
		{"Apply 50%", fields{memdb.New()}, args{50, "MegaDiscount", 100}, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := CouponService{
				repo: tt.fields.repo,
			}

			_, err := s.Insert(tt.args.discount, tt.args.code, tt.args.minBasketValue)
			if err != nil {
				t.Errorf("Unexpected error: %v", err)
			}
		})
	}
}

func TestCouponService_UpdateBasket(t *testing.T) {
	type fields struct {
		repo Repository
	}
	type args struct {
		basket entity.Basket
		code   string
	}
	tests := []struct {
		name          string
		fields        fields
		args          args
		wantB         *entity.Basket
		wantErr       bool
		expectedValue int
	}{
		{
			name: "Valid coupon code",
			fields: fields{
				repo: &mockRepository{
					findByCodeFunc: func(code string) (*entity.Coupon, error) {
						return &entity.Coupon{
							Discount: 10,
						}, nil
					},
				},
			},
			args: args{
				basket: entity.Basket{
					Value: 100,
				},
				code: "VALIDTESTCODE",
			},
			wantB: &entity.Basket{
				Value:                 100,
				AppliedDiscount:       10,
				ApplicationSuccessful: true,
			},
			wantErr:       false,
			expectedValue: 100,
		},
		{
			name: "Invalid coupon code",
			fields: fields{
				repo: &mockRepository{
					findByCodeFunc: func(code string) (*entity.Coupon, error) {
						return nil, fmt.Errorf("coupon not found")
					},
				},
			},
			args: args{
				basket: entity.Basket{
					Value: 100,
				},
				code: "INVALIDCODE",
			},
			wantB:         nil,
			wantErr:       true,
			expectedValue: 100,
		},
		{
			name: "Negative value basket",
			fields: fields{
				repo: &mockRepository{
					findByCodeFunc: func(code string) (*entity.Coupon, error) {
						return &entity.Coupon{
							Discount: 10,
						}, nil
					},
				},
			},
			args: args{
				basket: entity.Basket{
					Value: -100,
				},
				code: "TESTCODE",
			},
			wantB:         nil,
			wantErr:       true,
			expectedValue: -100,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := CouponService{
				repo: tt.fields.repo,
			}

			gotB, err := s.UpdateBasket(tt.args.basket, tt.args.code)
			if (err != nil) != tt.wantErr {
				t.Errorf("UpdateBasket() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotB, tt.wantB) {
				t.Errorf("UpdateBasket() gotB = %v, want %v", gotB, tt.wantB)
			}
			if gotB != nil && gotB.Value != tt.expectedValue {
				t.Errorf("UpdateBasket() gotB.Value = %v, expectedValue %v", gotB.Value, tt.expectedValue)
			}
		})
	}
}

func TestCouponService_GetByCodes(t *testing.T) {
	type fields struct {
		repo Repository
	}
	type args struct {
		codes []string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []entity.Coupon
		wantErr bool
	}{
		{
			name: "Valid codes",
			fields: fields{
				repo: &mockRepository{
					findByCodeFunc: func(code string) (*entity.Coupon, error) {
						return &entity.Coupon{
							Discount: 10,
						}, nil
					},
				},
			},
			args: args{
				codes: []string{"CODE1", "CODE2", "CODE3"},
			},
			want: []entity.Coupon{
				{
					Discount: 10,
				},
				{
					Discount: 10,
				},
				{
					Discount: 10,
				},
			},
			wantErr: false,
		},
		{
			name: "Invalid codes",
			fields: fields{
				repo: &mockRepository{
					findByCodeFunc: func(code string) (*entity.Coupon, error) {
						return nil, fmt.Errorf("coupon not found")
					},
				},
			},
			args: args{
				codes: []string{"CODE1", "CODE2", "CODE3"},
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := CouponService{
				repo: tt.fields.repo,
			}
			got, err := s.GetByCodes(tt.args.codes)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetByCodes() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetByCodes() got = %v, want %v", got, tt.want)
			}
		})
	}
}
