package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/techschool/simplebank/db/util"
)

func createRandomAccount(t *testing.T) Account {
	arg := CreateAccountParams{
		Owner:    util.RandomOwner(),
		Balance:  util.RandomMoney(),
		Currency: util.RandomCurrency(),
	}
	account, err := testQueries.CreateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, account)
	require.Equal(t, arg.Owner, account.Owner)
	require.Equal(t, arg.Balance, account.Balance)
	require.Equal(t, arg.Currency, account.Currency)
	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)

	return account
}

func TestCreateAccount(t *testing.T) {
	createRandomAccount(t)
}

func TestGetAccount(t *testing.T) {
	newAcc := createRandomAccount(t)
	foundAcc, err := testQueries.GetAccount(context.Background(), newAcc.ID)
	require.NoError(t, err)
	require.NotEmpty(t, foundAcc)

	require.Equal(t, newAcc.ID, foundAcc.ID)
	require.Equal(t, newAcc.Owner, foundAcc.Owner)
	require.Equal(t, newAcc.Balance, foundAcc.Balance)
	require.Equal(t, newAcc.Currency, foundAcc.Currency)
	require.WithinDuration(t, newAcc.CreatedAt, foundAcc.CreatedAt, time.Second)
}

func TestUpdateAccount(t *testing.T) {
	newAcc := createRandomAccount(t)
	arg := UpdateAccountParams{
		ID:      newAcc.ID,
		Balance: util.RandomMoney(),
	}
	updatedAcc, err := testQueries.UpdateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, updatedAcc)

	require.Equal(t, newAcc.ID, updatedAcc.ID)
	require.Equal(t, newAcc.Owner, updatedAcc.Owner)
	require.Equal(t, arg.Balance, updatedAcc.Balance)
	require.Equal(t, newAcc.Currency, updatedAcc.Currency)
	require.WithinDuration(t, newAcc.CreatedAt, updatedAcc.CreatedAt, time.Second)
}

func TestDelete(t *testing.T) {
	newAcc := createRandomAccount(t)

	err := testQueries.DeleteAccount(context.Background(), newAcc.ID)
	require.NoError(t, err)

	foundAcc, err := testQueries.GetAccount(context.Background(), newAcc.ID)
	require.Error(t, err) // supposed to be err
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, foundAcc)
}

func TestListAccounts(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomAccount(t)
	}

	arg := ListAccountsParams{
		Limit:  5,
		Offset: 5,
	}

	accounts, err := testQueries.ListAccounts(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, accounts, 5)

	for _, account := range accounts {
		require.NotEmpty(t, account)
	}
}
