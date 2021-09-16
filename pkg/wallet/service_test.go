package wallet

import (
	"github.com/kobilov-web-dev/wallet/pkg/types"
	"reflect"
	"testing"
)

func TestService_FindAccountByID_success(t *testing.T) { 
	svc := &Service{}

	account := &types.Account{
		ID:      1,
	}
	svc.accounts = append(svc.accounts, account)

	expected := &types.Account{
		ID:      1,
	}

	result, _ := svc.FindAccountByID(1)

	if !reflect.DeepEqual(expected, result) {
		t.Errorf("invalid result, expected: %v, actual: %v", expected, result)
	}
}

func TestService_FindAccountByID_NotFound(t *testing.T) { 
	svc := &Service{}

	expected := ErrAccountNotFound

	_, result := svc.FindAccountByID(1)

	if !reflect.DeepEqual(expected, result) {
		t.Errorf("invalid result, expected: %v, actual: %v", expected, result)
	}
}
