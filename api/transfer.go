package api

import (
	"database/sql"
	"fmt"
	"net/http"

	db "github.com/ArsalanKm/simple_bank/db/sqlc"
	"github.com/gin-gonic/gin"
)

func (server *Server) createTransfer(ctx *gin.Context) {

	var req transferRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	if !server.validAccount(ctx, req.FromAccountId, req.Currency) {

		return
	}

	if !server.validAccount(ctx, req.ToAccountId, req.Currency) {

		return
	}

	arg := db.TransferTxParams{
		FromAccountID: sql.NullInt64{Valid: true, Int64: req.FromAccountId},
		ToAccountID:   sql.NullInt64{Valid: true, Int64: req.ToAccountId},
		Amount:        req.Amount,
	}

	result, err := server.store.TransferTx(ctx, arg)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, result)
}

func (server *Server) validAccount(ctx *gin.Context, accountId int64, currency string) bool {

	account, err := server.store.GetAccount(ctx, accountId)

	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return false
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return false
	}

	if account.Currency != currency {
		err := fmt.Errorf("account [%d] currency mismatch : %s vs %s", accountId, account.Currency, currency)

		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return false
	}

	return true
}
