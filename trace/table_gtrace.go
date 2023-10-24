// Code generated by gtrace. DO NOT EDIT.

package trace

import (
	"context"
)

// tableComposeOptions is a holder of options
type tableComposeOptions struct {
	panicCallback func(e interface{})
}

// TableOption specified Table compose option
type TableComposeOption func(o *tableComposeOptions)

// WithTablePanicCallback specified behavior on panic
func WithTablePanicCallback(cb func(e interface{})) TableComposeOption {
	return func(o *tableComposeOptions) {
		o.panicCallback = cb
	}
}

// Compose returns a new Table which has functional fields composed both from t and x.
func (t *Table) Compose(x *Table, opts ...TableComposeOption) *Table {
	var ret Table
	options := tableComposeOptions{}
	for _, opt := range opts {
		if opt != nil {
			opt(&options)
		}
	}
	{
		h1 := t.OnInit
		h2 := x.OnInit
		ret.OnInit = func(t TableInitStartInfo) func(TableInitDoneInfo) {
			if options.panicCallback != nil {
				defer func() {
					if e := recover(); e != nil {
						options.panicCallback(e)
					}
				}()
			}
			var r, r1 func(TableInitDoneInfo)
			if h1 != nil {
				r = h1(t)
			}
			if h2 != nil {
				r1 = h2(t)
			}
			return func(t TableInitDoneInfo) {
				if options.panicCallback != nil {
					defer func() {
						if e := recover(); e != nil {
							options.panicCallback(e)
						}
					}()
				}
				if r != nil {
					r(t)
				}
				if r1 != nil {
					r1(t)
				}
			}
		}
	}
	{
		h1 := t.OnClose
		h2 := x.OnClose
		ret.OnClose = func(t TableCloseStartInfo) func(TableCloseDoneInfo) {
			if options.panicCallback != nil {
				defer func() {
					if e := recover(); e != nil {
						options.panicCallback(e)
					}
				}()
			}
			var r, r1 func(TableCloseDoneInfo)
			if h1 != nil {
				r = h1(t)
			}
			if h2 != nil {
				r1 = h2(t)
			}
			return func(t TableCloseDoneInfo) {
				if options.panicCallback != nil {
					defer func() {
						if e := recover(); e != nil {
							options.panicCallback(e)
						}
					}()
				}
				if r != nil {
					r(t)
				}
				if r1 != nil {
					r1(t)
				}
			}
		}
	}
	{
		h1 := t.OnDo
		h2 := x.OnDo
		ret.OnDo = func(t TableDoStartInfo) func(TableDoIntermediateInfo) func(TableDoDoneInfo) {
			if options.panicCallback != nil {
				defer func() {
					if e := recover(); e != nil {
						options.panicCallback(e)
					}
				}()
			}
			var r, r1 func(TableDoIntermediateInfo) func(TableDoDoneInfo)
			if h1 != nil {
				r = h1(t)
			}
			if h2 != nil {
				r1 = h2(t)
			}
			return func(info TableDoIntermediateInfo) func(TableDoDoneInfo) {
				if options.panicCallback != nil {
					defer func() {
						if e := recover(); e != nil {
							options.panicCallback(e)
						}
					}()
				}
				var r2, r3 func(TableDoDoneInfo)
				if r != nil {
					r2 = r(info)
				}
				if r1 != nil {
					r3 = r1(info)
				}
				return func(t TableDoDoneInfo) {
					if options.panicCallback != nil {
						defer func() {
							if e := recover(); e != nil {
								options.panicCallback(e)
							}
						}()
					}
					if r2 != nil {
						r2(t)
					}
					if r3 != nil {
						r3(t)
					}
				}
			}
		}
	}
	{
		h1 := t.OnDoTx
		h2 := x.OnDoTx
		ret.OnDoTx = func(t TableDoTxStartInfo) func(TableDoTxIntermediateInfo) func(TableDoTxDoneInfo) {
			if options.panicCallback != nil {
				defer func() {
					if e := recover(); e != nil {
						options.panicCallback(e)
					}
				}()
			}
			var r, r1 func(TableDoTxIntermediateInfo) func(TableDoTxDoneInfo)
			if h1 != nil {
				r = h1(t)
			}
			if h2 != nil {
				r1 = h2(t)
			}
			return func(info TableDoTxIntermediateInfo) func(TableDoTxDoneInfo) {
				if options.panicCallback != nil {
					defer func() {
						if e := recover(); e != nil {
							options.panicCallback(e)
						}
					}()
				}
				var r2, r3 func(TableDoTxDoneInfo)
				if r != nil {
					r2 = r(info)
				}
				if r1 != nil {
					r3 = r1(info)
				}
				return func(t TableDoTxDoneInfo) {
					if options.panicCallback != nil {
						defer func() {
							if e := recover(); e != nil {
								options.panicCallback(e)
							}
						}()
					}
					if r2 != nil {
						r2(t)
					}
					if r3 != nil {
						r3(t)
					}
				}
			}
		}
	}
	{
		h1 := t.OnCreateSession
		h2 := x.OnCreateSession
		ret.OnCreateSession = func(t TableCreateSessionStartInfo) func(TableCreateSessionIntermediateInfo) func(TableCreateSessionDoneInfo) {
			if options.panicCallback != nil {
				defer func() {
					if e := recover(); e != nil {
						options.panicCallback(e)
					}
				}()
			}
			var r, r1 func(TableCreateSessionIntermediateInfo) func(TableCreateSessionDoneInfo)
			if h1 != nil {
				r = h1(t)
			}
			if h2 != nil {
				r1 = h2(t)
			}
			return func(info TableCreateSessionIntermediateInfo) func(TableCreateSessionDoneInfo) {
				if options.panicCallback != nil {
					defer func() {
						if e := recover(); e != nil {
							options.panicCallback(e)
						}
					}()
				}
				var r2, r3 func(TableCreateSessionDoneInfo)
				if r != nil {
					r2 = r(info)
				}
				if r1 != nil {
					r3 = r1(info)
				}
				return func(t TableCreateSessionDoneInfo) {
					if options.panicCallback != nil {
						defer func() {
							if e := recover(); e != nil {
								options.panicCallback(e)
							}
						}()
					}
					if r2 != nil {
						r2(t)
					}
					if r3 != nil {
						r3(t)
					}
				}
			}
		}
	}
	{
		h1 := t.OnSessionNew
		h2 := x.OnSessionNew
		ret.OnSessionNew = func(t TableSessionNewStartInfo) func(TableSessionNewDoneInfo) {
			if options.panicCallback != nil {
				defer func() {
					if e := recover(); e != nil {
						options.panicCallback(e)
					}
				}()
			}
			var r, r1 func(TableSessionNewDoneInfo)
			if h1 != nil {
				r = h1(t)
			}
			if h2 != nil {
				r1 = h2(t)
			}
			return func(t TableSessionNewDoneInfo) {
				if options.panicCallback != nil {
					defer func() {
						if e := recover(); e != nil {
							options.panicCallback(e)
						}
					}()
				}
				if r != nil {
					r(t)
				}
				if r1 != nil {
					r1(t)
				}
			}
		}
	}
	{
		h1 := t.OnSessionDelete
		h2 := x.OnSessionDelete
		ret.OnSessionDelete = func(t TableSessionDeleteStartInfo) func(TableSessionDeleteDoneInfo) {
			if options.panicCallback != nil {
				defer func() {
					if e := recover(); e != nil {
						options.panicCallback(e)
					}
				}()
			}
			var r, r1 func(TableSessionDeleteDoneInfo)
			if h1 != nil {
				r = h1(t)
			}
			if h2 != nil {
				r1 = h2(t)
			}
			return func(t TableSessionDeleteDoneInfo) {
				if options.panicCallback != nil {
					defer func() {
						if e := recover(); e != nil {
							options.panicCallback(e)
						}
					}()
				}
				if r != nil {
					r(t)
				}
				if r1 != nil {
					r1(t)
				}
			}
		}
	}
	{
		h1 := t.OnSessionKeepAlive
		h2 := x.OnSessionKeepAlive
		ret.OnSessionKeepAlive = func(t TableKeepAliveStartInfo) func(TableKeepAliveDoneInfo) {
			if options.panicCallback != nil {
				defer func() {
					if e := recover(); e != nil {
						options.panicCallback(e)
					}
				}()
			}
			var r, r1 func(TableKeepAliveDoneInfo)
			if h1 != nil {
				r = h1(t)
			}
			if h2 != nil {
				r1 = h2(t)
			}
			return func(t TableKeepAliveDoneInfo) {
				if options.panicCallback != nil {
					defer func() {
						if e := recover(); e != nil {
							options.panicCallback(e)
						}
					}()
				}
				if r != nil {
					r(t)
				}
				if r1 != nil {
					r1(t)
				}
			}
		}
	}
	{
		h1 := t.OnSessionQueryPrepare
		h2 := x.OnSessionQueryPrepare
		ret.OnSessionQueryPrepare = func(t TablePrepareDataQueryStartInfo) func(TablePrepareDataQueryDoneInfo) {
			if options.panicCallback != nil {
				defer func() {
					if e := recover(); e != nil {
						options.panicCallback(e)
					}
				}()
			}
			var r, r1 func(TablePrepareDataQueryDoneInfo)
			if h1 != nil {
				r = h1(t)
			}
			if h2 != nil {
				r1 = h2(t)
			}
			return func(t TablePrepareDataQueryDoneInfo) {
				if options.panicCallback != nil {
					defer func() {
						if e := recover(); e != nil {
							options.panicCallback(e)
						}
					}()
				}
				if r != nil {
					r(t)
				}
				if r1 != nil {
					r1(t)
				}
			}
		}
	}
	{
		h1 := t.OnSessionQueryExecute
		h2 := x.OnSessionQueryExecute
		ret.OnSessionQueryExecute = func(t TableExecuteDataQueryStartInfo) func(TableExecuteDataQueryDoneInfo) {
			if options.panicCallback != nil {
				defer func() {
					if e := recover(); e != nil {
						options.panicCallback(e)
					}
				}()
			}
			var r, r1 func(TableExecuteDataQueryDoneInfo)
			if h1 != nil {
				r = h1(t)
			}
			if h2 != nil {
				r1 = h2(t)
			}
			return func(t TableExecuteDataQueryDoneInfo) {
				if options.panicCallback != nil {
					defer func() {
						if e := recover(); e != nil {
							options.panicCallback(e)
						}
					}()
				}
				if r != nil {
					r(t)
				}
				if r1 != nil {
					r1(t)
				}
			}
		}
	}
	{
		h1 := t.OnSessionQueryExplain
		h2 := x.OnSessionQueryExplain
		ret.OnSessionQueryExplain = func(t TableExplainQueryStartInfo) func(TableExplainQueryDoneInfo) {
			if options.panicCallback != nil {
				defer func() {
					if e := recover(); e != nil {
						options.panicCallback(e)
					}
				}()
			}
			var r, r1 func(TableExplainQueryDoneInfo)
			if h1 != nil {
				r = h1(t)
			}
			if h2 != nil {
				r1 = h2(t)
			}
			return func(t TableExplainQueryDoneInfo) {
				if options.panicCallback != nil {
					defer func() {
						if e := recover(); e != nil {
							options.panicCallback(e)
						}
					}()
				}
				if r != nil {
					r(t)
				}
				if r1 != nil {
					r1(t)
				}
			}
		}
	}
	{
		h1 := t.OnSessionQueryStreamExecute
		h2 := x.OnSessionQueryStreamExecute
		ret.OnSessionQueryStreamExecute = func(t TableSessionQueryStreamExecuteStartInfo) func(TableSessionQueryStreamExecuteIntermediateInfo) func(TableSessionQueryStreamExecuteDoneInfo) {
			if options.panicCallback != nil {
				defer func() {
					if e := recover(); e != nil {
						options.panicCallback(e)
					}
				}()
			}
			var r, r1 func(TableSessionQueryStreamExecuteIntermediateInfo) func(TableSessionQueryStreamExecuteDoneInfo)
			if h1 != nil {
				r = h1(t)
			}
			if h2 != nil {
				r1 = h2(t)
			}
			return func(t TableSessionQueryStreamExecuteIntermediateInfo) func(TableSessionQueryStreamExecuteDoneInfo) {
				if options.panicCallback != nil {
					defer func() {
						if e := recover(); e != nil {
							options.panicCallback(e)
						}
					}()
				}
				var r2, r3 func(TableSessionQueryStreamExecuteDoneInfo)
				if r != nil {
					r2 = r(t)
				}
				if r1 != nil {
					r3 = r1(t)
				}
				return func(t TableSessionQueryStreamExecuteDoneInfo) {
					if options.panicCallback != nil {
						defer func() {
							if e := recover(); e != nil {
								options.panicCallback(e)
							}
						}()
					}
					if r2 != nil {
						r2(t)
					}
					if r3 != nil {
						r3(t)
					}
				}
			}
		}
	}
	{
		h1 := t.OnSessionQueryStreamRead
		h2 := x.OnSessionQueryStreamRead
		ret.OnSessionQueryStreamRead = func(t TableSessionQueryStreamReadStartInfo) func(TableSessionQueryStreamReadIntermediateInfo) func(TableSessionQueryStreamReadDoneInfo) {
			if options.panicCallback != nil {
				defer func() {
					if e := recover(); e != nil {
						options.panicCallback(e)
					}
				}()
			}
			var r, r1 func(TableSessionQueryStreamReadIntermediateInfo) func(TableSessionQueryStreamReadDoneInfo)
			if h1 != nil {
				r = h1(t)
			}
			if h2 != nil {
				r1 = h2(t)
			}
			return func(t TableSessionQueryStreamReadIntermediateInfo) func(TableSessionQueryStreamReadDoneInfo) {
				if options.panicCallback != nil {
					defer func() {
						if e := recover(); e != nil {
							options.panicCallback(e)
						}
					}()
				}
				var r2, r3 func(TableSessionQueryStreamReadDoneInfo)
				if r != nil {
					r2 = r(t)
				}
				if r1 != nil {
					r3 = r1(t)
				}
				return func(t TableSessionQueryStreamReadDoneInfo) {
					if options.panicCallback != nil {
						defer func() {
							if e := recover(); e != nil {
								options.panicCallback(e)
							}
						}()
					}
					if r2 != nil {
						r2(t)
					}
					if r3 != nil {
						r3(t)
					}
				}
			}
		}
	}
	{
		h1 := t.OnSessionTransactionBegin
		h2 := x.OnSessionTransactionBegin
		ret.OnSessionTransactionBegin = func(t TableSessionTransactionBeginStartInfo) func(TableSessionTransactionBeginDoneInfo) {
			if options.panicCallback != nil {
				defer func() {
					if e := recover(); e != nil {
						options.panicCallback(e)
					}
				}()
			}
			var r, r1 func(TableSessionTransactionBeginDoneInfo)
			if h1 != nil {
				r = h1(t)
			}
			if h2 != nil {
				r1 = h2(t)
			}
			return func(t TableSessionTransactionBeginDoneInfo) {
				if options.panicCallback != nil {
					defer func() {
						if e := recover(); e != nil {
							options.panicCallback(e)
						}
					}()
				}
				if r != nil {
					r(t)
				}
				if r1 != nil {
					r1(t)
				}
			}
		}
	}
	{
		h1 := t.OnSessionTransactionExecute
		h2 := x.OnSessionTransactionExecute
		ret.OnSessionTransactionExecute = func(t TableTransactionExecuteStartInfo) func(TableTransactionExecuteDoneInfo) {
			if options.panicCallback != nil {
				defer func() {
					if e := recover(); e != nil {
						options.panicCallback(e)
					}
				}()
			}
			var r, r1 func(TableTransactionExecuteDoneInfo)
			if h1 != nil {
				r = h1(t)
			}
			if h2 != nil {
				r1 = h2(t)
			}
			return func(t TableTransactionExecuteDoneInfo) {
				if options.panicCallback != nil {
					defer func() {
						if e := recover(); e != nil {
							options.panicCallback(e)
						}
					}()
				}
				if r != nil {
					r(t)
				}
				if r1 != nil {
					r1(t)
				}
			}
		}
	}
	{
		h1 := t.OnSessionTransactionExecuteStatement
		h2 := x.OnSessionTransactionExecuteStatement
		ret.OnSessionTransactionExecuteStatement = func(t TableTransactionExecuteStatementStartInfo) func(TableTransactionExecuteStatementDoneInfo) {
			if options.panicCallback != nil {
				defer func() {
					if e := recover(); e != nil {
						options.panicCallback(e)
					}
				}()
			}
			var r, r1 func(TableTransactionExecuteStatementDoneInfo)
			if h1 != nil {
				r = h1(t)
			}
			if h2 != nil {
				r1 = h2(t)
			}
			return func(t TableTransactionExecuteStatementDoneInfo) {
				if options.panicCallback != nil {
					defer func() {
						if e := recover(); e != nil {
							options.panicCallback(e)
						}
					}()
				}
				if r != nil {
					r(t)
				}
				if r1 != nil {
					r1(t)
				}
			}
		}
	}
	{
		h1 := t.OnSessionTransactionCommit
		h2 := x.OnSessionTransactionCommit
		ret.OnSessionTransactionCommit = func(t TableSessionTransactionCommitStartInfo) func(TableSessionTransactionCommitDoneInfo) {
			if options.panicCallback != nil {
				defer func() {
					if e := recover(); e != nil {
						options.panicCallback(e)
					}
				}()
			}
			var r, r1 func(TableSessionTransactionCommitDoneInfo)
			if h1 != nil {
				r = h1(t)
			}
			if h2 != nil {
				r1 = h2(t)
			}
			return func(t TableSessionTransactionCommitDoneInfo) {
				if options.panicCallback != nil {
					defer func() {
						if e := recover(); e != nil {
							options.panicCallback(e)
						}
					}()
				}
				if r != nil {
					r(t)
				}
				if r1 != nil {
					r1(t)
				}
			}
		}
	}
	{
		h1 := t.OnSessionTransactionRollback
		h2 := x.OnSessionTransactionRollback
		ret.OnSessionTransactionRollback = func(t TableSessionTransactionRollbackStartInfo) func(TableSessionTransactionRollbackDoneInfo) {
			if options.panicCallback != nil {
				defer func() {
					if e := recover(); e != nil {
						options.panicCallback(e)
					}
				}()
			}
			var r, r1 func(TableSessionTransactionRollbackDoneInfo)
			if h1 != nil {
				r = h1(t)
			}
			if h2 != nil {
				r1 = h2(t)
			}
			return func(t TableSessionTransactionRollbackDoneInfo) {
				if options.panicCallback != nil {
					defer func() {
						if e := recover(); e != nil {
							options.panicCallback(e)
						}
					}()
				}
				if r != nil {
					r(t)
				}
				if r1 != nil {
					r1(t)
				}
			}
		}
	}
	{
		h1 := t.OnPoolStateChange
		h2 := x.OnPoolStateChange
		ret.OnPoolStateChange = func(t TablePoolStateChangeInfo) {
			if options.panicCallback != nil {
				defer func() {
					if e := recover(); e != nil {
						options.panicCallback(e)
					}
				}()
			}
			if h1 != nil {
				h1(t)
			}
			if h2 != nil {
				h2(t)
			}
		}
	}
	{
		h1 := t.OnPoolSessionAdd
		h2 := x.OnPoolSessionAdd
		ret.OnPoolSessionAdd = func(info TablePoolSessionAddInfo) {
			if options.panicCallback != nil {
				defer func() {
					if e := recover(); e != nil {
						options.panicCallback(e)
					}
				}()
			}
			if h1 != nil {
				h1(info)
			}
			if h2 != nil {
				h2(info)
			}
		}
	}
	{
		h1 := t.OnPoolSessionRemove
		h2 := x.OnPoolSessionRemove
		ret.OnPoolSessionRemove = func(info TablePoolSessionRemoveInfo) {
			if options.panicCallback != nil {
				defer func() {
					if e := recover(); e != nil {
						options.panicCallback(e)
					}
				}()
			}
			if h1 != nil {
				h1(info)
			}
			if h2 != nil {
				h2(info)
			}
		}
	}
	{
		h1 := t.OnPoolSessionNew
		h2 := x.OnPoolSessionNew
		ret.OnPoolSessionNew = func(t TablePoolSessionNewStartInfo) func(TablePoolSessionNewDoneInfo) {
			if options.panicCallback != nil {
				defer func() {
					if e := recover(); e != nil {
						options.panicCallback(e)
					}
				}()
			}
			var r, r1 func(TablePoolSessionNewDoneInfo)
			if h1 != nil {
				r = h1(t)
			}
			if h2 != nil {
				r1 = h2(t)
			}
			return func(t TablePoolSessionNewDoneInfo) {
				if options.panicCallback != nil {
					defer func() {
						if e := recover(); e != nil {
							options.panicCallback(e)
						}
					}()
				}
				if r != nil {
					r(t)
				}
				if r1 != nil {
					r1(t)
				}
			}
		}
	}
	{
		h1 := t.OnPoolSessionClose
		h2 := x.OnPoolSessionClose
		ret.OnPoolSessionClose = func(t TablePoolSessionCloseStartInfo) func(TablePoolSessionCloseDoneInfo) {
			if options.panicCallback != nil {
				defer func() {
					if e := recover(); e != nil {
						options.panicCallback(e)
					}
				}()
			}
			var r, r1 func(TablePoolSessionCloseDoneInfo)
			if h1 != nil {
				r = h1(t)
			}
			if h2 != nil {
				r1 = h2(t)
			}
			return func(t TablePoolSessionCloseDoneInfo) {
				if options.panicCallback != nil {
					defer func() {
						if e := recover(); e != nil {
							options.panicCallback(e)
						}
					}()
				}
				if r != nil {
					r(t)
				}
				if r1 != nil {
					r1(t)
				}
			}
		}
	}
	{
		h1 := t.OnPoolPut
		h2 := x.OnPoolPut
		ret.OnPoolPut = func(t TablePoolPutStartInfo) func(TablePoolPutDoneInfo) {
			if options.panicCallback != nil {
				defer func() {
					if e := recover(); e != nil {
						options.panicCallback(e)
					}
				}()
			}
			var r, r1 func(TablePoolPutDoneInfo)
			if h1 != nil {
				r = h1(t)
			}
			if h2 != nil {
				r1 = h2(t)
			}
			return func(t TablePoolPutDoneInfo) {
				if options.panicCallback != nil {
					defer func() {
						if e := recover(); e != nil {
							options.panicCallback(e)
						}
					}()
				}
				if r != nil {
					r(t)
				}
				if r1 != nil {
					r1(t)
				}
			}
		}
	}
	{
		h1 := t.OnPoolGet
		h2 := x.OnPoolGet
		ret.OnPoolGet = func(t TablePoolGetStartInfo) func(TablePoolGetDoneInfo) {
			if options.panicCallback != nil {
				defer func() {
					if e := recover(); e != nil {
						options.panicCallback(e)
					}
				}()
			}
			var r, r1 func(TablePoolGetDoneInfo)
			if h1 != nil {
				r = h1(t)
			}
			if h2 != nil {
				r1 = h2(t)
			}
			return func(t TablePoolGetDoneInfo) {
				if options.panicCallback != nil {
					defer func() {
						if e := recover(); e != nil {
							options.panicCallback(e)
						}
					}()
				}
				if r != nil {
					r(t)
				}
				if r1 != nil {
					r1(t)
				}
			}
		}
	}
	{
		h1 := t.OnPoolWait
		h2 := x.OnPoolWait
		ret.OnPoolWait = func(t TablePoolWaitStartInfo) func(TablePoolWaitDoneInfo) {
			if options.panicCallback != nil {
				defer func() {
					if e := recover(); e != nil {
						options.panicCallback(e)
					}
				}()
			}
			var r, r1 func(TablePoolWaitDoneInfo)
			if h1 != nil {
				r = h1(t)
			}
			if h2 != nil {
				r1 = h2(t)
			}
			return func(t TablePoolWaitDoneInfo) {
				if options.panicCallback != nil {
					defer func() {
						if e := recover(); e != nil {
							options.panicCallback(e)
						}
					}()
				}
				if r != nil {
					r(t)
				}
				if r1 != nil {
					r1(t)
				}
			}
		}
	}
	return &ret
}
func (t *Table) onInit(t1 TableInitStartInfo) func(TableInitDoneInfo) {
	fn := t.OnInit
	if fn == nil {
		return func(TableInitDoneInfo) {
			return
		}
	}
	res := fn(t1)
	if res == nil {
		return func(TableInitDoneInfo) {
			return
		}
	}
	return res
}
func (t *Table) onClose(t1 TableCloseStartInfo) func(TableCloseDoneInfo) {
	fn := t.OnClose
	if fn == nil {
		return func(TableCloseDoneInfo) {
			return
		}
	}
	res := fn(t1)
	if res == nil {
		return func(TableCloseDoneInfo) {
			return
		}
	}
	return res
}
func (t *Table) onDo(t1 TableDoStartInfo) func(info TableDoIntermediateInfo) func(TableDoDoneInfo) {
	fn := t.OnDo
	if fn == nil {
		return func(TableDoIntermediateInfo) func(TableDoDoneInfo) {
			return func(TableDoDoneInfo) {
				return
			}
		}
	}
	res := fn(t1)
	if res == nil {
		return func(TableDoIntermediateInfo) func(TableDoDoneInfo) {
			return func(TableDoDoneInfo) {
				return
			}
		}
	}
	return func(info TableDoIntermediateInfo) func(TableDoDoneInfo) {
		res := res(info)
		if res == nil {
			return func(TableDoDoneInfo) {
				return
			}
		}
		return res
	}
}
func (t *Table) onDoTx(t1 TableDoTxStartInfo) func(info TableDoTxIntermediateInfo) func(TableDoTxDoneInfo) {
	fn := t.OnDoTx
	if fn == nil {
		return func(TableDoTxIntermediateInfo) func(TableDoTxDoneInfo) {
			return func(TableDoTxDoneInfo) {
				return
			}
		}
	}
	res := fn(t1)
	if res == nil {
		return func(TableDoTxIntermediateInfo) func(TableDoTxDoneInfo) {
			return func(TableDoTxDoneInfo) {
				return
			}
		}
	}
	return func(info TableDoTxIntermediateInfo) func(TableDoTxDoneInfo) {
		res := res(info)
		if res == nil {
			return func(TableDoTxDoneInfo) {
				return
			}
		}
		return res
	}
}
func (t *Table) onCreateSession(t1 TableCreateSessionStartInfo) func(info TableCreateSessionIntermediateInfo) func(TableCreateSessionDoneInfo) {
	fn := t.OnCreateSession
	if fn == nil {
		return func(TableCreateSessionIntermediateInfo) func(TableCreateSessionDoneInfo) {
			return func(TableCreateSessionDoneInfo) {
				return
			}
		}
	}
	res := fn(t1)
	if res == nil {
		return func(TableCreateSessionIntermediateInfo) func(TableCreateSessionDoneInfo) {
			return func(TableCreateSessionDoneInfo) {
				return
			}
		}
	}
	return func(info TableCreateSessionIntermediateInfo) func(TableCreateSessionDoneInfo) {
		res := res(info)
		if res == nil {
			return func(TableCreateSessionDoneInfo) {
				return
			}
		}
		return res
	}
}
func (t *Table) onSessionNew(t1 TableSessionNewStartInfo) func(TableSessionNewDoneInfo) {
	fn := t.OnSessionNew
	if fn == nil {
		return func(TableSessionNewDoneInfo) {
			return
		}
	}
	res := fn(t1)
	if res == nil {
		return func(TableSessionNewDoneInfo) {
			return
		}
	}
	return res
}
func (t *Table) onSessionDelete(t1 TableSessionDeleteStartInfo) func(TableSessionDeleteDoneInfo) {
	fn := t.OnSessionDelete
	if fn == nil {
		return func(TableSessionDeleteDoneInfo) {
			return
		}
	}
	res := fn(t1)
	if res == nil {
		return func(TableSessionDeleteDoneInfo) {
			return
		}
	}
	return res
}
func (t *Table) onSessionKeepAlive(t1 TableKeepAliveStartInfo) func(TableKeepAliveDoneInfo) {
	fn := t.OnSessionKeepAlive
	if fn == nil {
		return func(TableKeepAliveDoneInfo) {
			return
		}
	}
	res := fn(t1)
	if res == nil {
		return func(TableKeepAliveDoneInfo) {
			return
		}
	}
	return res
}
func (t *Table) onSessionQueryPrepare(t1 TablePrepareDataQueryStartInfo) func(TablePrepareDataQueryDoneInfo) {
	fn := t.OnSessionQueryPrepare
	if fn == nil {
		return func(TablePrepareDataQueryDoneInfo) {
			return
		}
	}
	res := fn(t1)
	if res == nil {
		return func(TablePrepareDataQueryDoneInfo) {
			return
		}
	}
	return res
}
func (t *Table) onSessionQueryExecute(t1 TableExecuteDataQueryStartInfo) func(TableExecuteDataQueryDoneInfo) {
	fn := t.OnSessionQueryExecute
	if fn == nil {
		return func(TableExecuteDataQueryDoneInfo) {
			return
		}
	}
	res := fn(t1)
	if res == nil {
		return func(TableExecuteDataQueryDoneInfo) {
			return
		}
	}
	return res
}
func (t *Table) onSessionQueryExplain(t1 TableExplainQueryStartInfo) func(TableExplainQueryDoneInfo) {
	fn := t.OnSessionQueryExplain
	if fn == nil {
		return func(TableExplainQueryDoneInfo) {
			return
		}
	}
	res := fn(t1)
	if res == nil {
		return func(TableExplainQueryDoneInfo) {
			return
		}
	}
	return res
}
func (t *Table) onSessionQueryStreamExecute(t1 TableSessionQueryStreamExecuteStartInfo) func(TableSessionQueryStreamExecuteIntermediateInfo) func(TableSessionQueryStreamExecuteDoneInfo) {
	fn := t.OnSessionQueryStreamExecute
	if fn == nil {
		return func(TableSessionQueryStreamExecuteIntermediateInfo) func(TableSessionQueryStreamExecuteDoneInfo) {
			return func(TableSessionQueryStreamExecuteDoneInfo) {
				return
			}
		}
	}
	res := fn(t1)
	if res == nil {
		return func(TableSessionQueryStreamExecuteIntermediateInfo) func(TableSessionQueryStreamExecuteDoneInfo) {
			return func(TableSessionQueryStreamExecuteDoneInfo) {
				return
			}
		}
	}
	return func(t TableSessionQueryStreamExecuteIntermediateInfo) func(TableSessionQueryStreamExecuteDoneInfo) {
		res := res(t)
		if res == nil {
			return func(TableSessionQueryStreamExecuteDoneInfo) {
				return
			}
		}
		return res
	}
}
func (t *Table) onSessionQueryStreamRead(t1 TableSessionQueryStreamReadStartInfo) func(TableSessionQueryStreamReadIntermediateInfo) func(TableSessionQueryStreamReadDoneInfo) {
	fn := t.OnSessionQueryStreamRead
	if fn == nil {
		return func(TableSessionQueryStreamReadIntermediateInfo) func(TableSessionQueryStreamReadDoneInfo) {
			return func(TableSessionQueryStreamReadDoneInfo) {
				return
			}
		}
	}
	res := fn(t1)
	if res == nil {
		return func(TableSessionQueryStreamReadIntermediateInfo) func(TableSessionQueryStreamReadDoneInfo) {
			return func(TableSessionQueryStreamReadDoneInfo) {
				return
			}
		}
	}
	return func(t TableSessionQueryStreamReadIntermediateInfo) func(TableSessionQueryStreamReadDoneInfo) {
		res := res(t)
		if res == nil {
			return func(TableSessionQueryStreamReadDoneInfo) {
				return
			}
		}
		return res
	}
}
func (t *Table) onSessionTransactionBegin(t1 TableSessionTransactionBeginStartInfo) func(TableSessionTransactionBeginDoneInfo) {
	fn := t.OnSessionTransactionBegin
	if fn == nil {
		return func(TableSessionTransactionBeginDoneInfo) {
			return
		}
	}
	res := fn(t1)
	if res == nil {
		return func(TableSessionTransactionBeginDoneInfo) {
			return
		}
	}
	return res
}
func (t *Table) onSessionTransactionExecute(t1 TableTransactionExecuteStartInfo) func(TableTransactionExecuteDoneInfo) {
	fn := t.OnSessionTransactionExecute
	if fn == nil {
		return func(TableTransactionExecuteDoneInfo) {
			return
		}
	}
	res := fn(t1)
	if res == nil {
		return func(TableTransactionExecuteDoneInfo) {
			return
		}
	}
	return res
}
func (t *Table) onSessionTransactionExecuteStatement(t1 TableTransactionExecuteStatementStartInfo) func(TableTransactionExecuteStatementDoneInfo) {
	fn := t.OnSessionTransactionExecuteStatement
	if fn == nil {
		return func(TableTransactionExecuteStatementDoneInfo) {
			return
		}
	}
	res := fn(t1)
	if res == nil {
		return func(TableTransactionExecuteStatementDoneInfo) {
			return
		}
	}
	return res
}
func (t *Table) onSessionTransactionCommit(t1 TableSessionTransactionCommitStartInfo) func(TableSessionTransactionCommitDoneInfo) {
	fn := t.OnSessionTransactionCommit
	if fn == nil {
		return func(TableSessionTransactionCommitDoneInfo) {
			return
		}
	}
	res := fn(t1)
	if res == nil {
		return func(TableSessionTransactionCommitDoneInfo) {
			return
		}
	}
	return res
}
func (t *Table) onSessionTransactionRollback(t1 TableSessionTransactionRollbackStartInfo) func(TableSessionTransactionRollbackDoneInfo) {
	fn := t.OnSessionTransactionRollback
	if fn == nil {
		return func(TableSessionTransactionRollbackDoneInfo) {
			return
		}
	}
	res := fn(t1)
	if res == nil {
		return func(TableSessionTransactionRollbackDoneInfo) {
			return
		}
	}
	return res
}
func (t *Table) onPoolStateChange(t1 TablePoolStateChangeInfo) {
	fn := t.OnPoolStateChange
	if fn == nil {
		return
	}
	fn(t1)
}
func (t *Table) onPoolSessionAdd(info TablePoolSessionAddInfo) {
	fn := t.OnPoolSessionAdd
	if fn == nil {
		return
	}
	fn(info)
}
func (t *Table) onPoolSessionRemove(info TablePoolSessionRemoveInfo) {
	fn := t.OnPoolSessionRemove
	if fn == nil {
		return
	}
	fn(info)
}
func (t *Table) onPoolSessionNew(t1 TablePoolSessionNewStartInfo) func(TablePoolSessionNewDoneInfo) {
	fn := t.OnPoolSessionNew
	if fn == nil {
		return func(TablePoolSessionNewDoneInfo) {
			return
		}
	}
	res := fn(t1)
	if res == nil {
		return func(TablePoolSessionNewDoneInfo) {
			return
		}
	}
	return res
}
func (t *Table) onPoolSessionClose(t1 TablePoolSessionCloseStartInfo) func(TablePoolSessionCloseDoneInfo) {
	fn := t.OnPoolSessionClose
	if fn == nil {
		return func(TablePoolSessionCloseDoneInfo) {
			return
		}
	}
	res := fn(t1)
	if res == nil {
		return func(TablePoolSessionCloseDoneInfo) {
			return
		}
	}
	return res
}
func (t *Table) onPoolPut(t1 TablePoolPutStartInfo) func(TablePoolPutDoneInfo) {
	fn := t.OnPoolPut
	if fn == nil {
		return func(TablePoolPutDoneInfo) {
			return
		}
	}
	res := fn(t1)
	if res == nil {
		return func(TablePoolPutDoneInfo) {
			return
		}
	}
	return res
}
func (t *Table) onPoolGet(t1 TablePoolGetStartInfo) func(TablePoolGetDoneInfo) {
	fn := t.OnPoolGet
	if fn == nil {
		return func(TablePoolGetDoneInfo) {
			return
		}
	}
	res := fn(t1)
	if res == nil {
		return func(TablePoolGetDoneInfo) {
			return
		}
	}
	return res
}
func (t *Table) onPoolWait(t1 TablePoolWaitStartInfo) func(TablePoolWaitDoneInfo) {
	fn := t.OnPoolWait
	if fn == nil {
		return func(TablePoolWaitDoneInfo) {
			return
		}
	}
	res := fn(t1)
	if res == nil {
		return func(TablePoolWaitDoneInfo) {
			return
		}
	}
	return res
}
func TableOnInit(t *Table, c *context.Context) func(limit int) {
	var p TableInitStartInfo
	p.Context = c
	res := t.onInit(p)
	return func(limit int) {
		var p TableInitDoneInfo
		p.Limit = limit
		res(p)
	}
}
func TableOnClose(t *Table, c *context.Context) func(error) {
	var p TableCloseStartInfo
	p.Context = c
	res := t.onClose(p)
	return func(e error) {
		var p TableCloseDoneInfo
		p.Error = e
		res(p)
	}
}
func TableOnDo(t *Table, c *context.Context, iD string, label string, idempotent bool, nestedCall bool) func(error) func(attempts int, _ error) {
	var p TableDoStartInfo
	p.Context = c
	p.ID = iD
	p.Label = label
	p.Idempotent = idempotent
	p.NestedCall = nestedCall
	res := t.onDo(p)
	return func(e error) func(int, error) {
		var p TableDoIntermediateInfo
		p.Error = e
		res := res(p)
		return func(attempts int, e error) {
			var p TableDoDoneInfo
			p.Attempts = attempts
			p.Error = e
			res(p)
		}
	}
}
func TableOnDoTx(t *Table, c *context.Context, iD string, label string, idempotent bool, nestedCall bool) func(error) func(attempts int, _ error) {
	var p TableDoTxStartInfo
	p.Context = c
	p.ID = iD
	p.Label = label
	p.Idempotent = idempotent
	p.NestedCall = nestedCall
	res := t.onDoTx(p)
	return func(e error) func(int, error) {
		var p TableDoTxIntermediateInfo
		p.Error = e
		res := res(p)
		return func(attempts int, e error) {
			var p TableDoTxDoneInfo
			p.Attempts = attempts
			p.Error = e
			res(p)
		}
	}
}
func TableOnCreateSession(t *Table, c *context.Context) func(error) func(session tableSessionInfo, attempts int, _ error) {
	var p TableCreateSessionStartInfo
	p.Context = c
	res := t.onCreateSession(p)
	return func(e error) func(tableSessionInfo, int, error) {
		var p TableCreateSessionIntermediateInfo
		p.Error = e
		res := res(p)
		return func(session tableSessionInfo, attempts int, e error) {
			var p TableCreateSessionDoneInfo
			p.Session = session
			p.Attempts = attempts
			p.Error = e
			res(p)
		}
	}
}
func TableOnSessionNew(t *Table, c *context.Context) func(session tableSessionInfo, _ error) {
	var p TableSessionNewStartInfo
	p.Context = c
	res := t.onSessionNew(p)
	return func(session tableSessionInfo, e error) {
		var p TableSessionNewDoneInfo
		p.Session = session
		p.Error = e
		res(p)
	}
}
func TableOnSessionDelete(t *Table, c *context.Context, session tableSessionInfo) func(error) {
	var p TableSessionDeleteStartInfo
	p.Context = c
	p.Session = session
	res := t.onSessionDelete(p)
	return func(e error) {
		var p TableSessionDeleteDoneInfo
		p.Error = e
		res(p)
	}
}
func TableOnSessionKeepAlive(t *Table, c *context.Context, session tableSessionInfo) func(error) {
	var p TableKeepAliveStartInfo
	p.Context = c
	p.Session = session
	res := t.onSessionKeepAlive(p)
	return func(e error) {
		var p TableKeepAliveDoneInfo
		p.Error = e
		res(p)
	}
}
func TableOnSessionQueryPrepare(t *Table, c *context.Context, session tableSessionInfo, query string) func(result tableDataQuery, _ error) {
	var p TablePrepareDataQueryStartInfo
	p.Context = c
	p.Session = session
	p.Query = query
	res := t.onSessionQueryPrepare(p)
	return func(result tableDataQuery, e error) {
		var p TablePrepareDataQueryDoneInfo
		p.Result = result
		p.Error = e
		res(p)
	}
}
func TableOnSessionQueryExecute(t *Table, c *context.Context, session tableSessionInfo, query tableDataQuery, parameters tableQueryParameters, keepInCache bool) func(tx tableTransactionInfo, prepared bool, result tableResult, _ error) {
	var p TableExecuteDataQueryStartInfo
	p.Context = c
	p.Session = session
	p.Query = query
	p.Parameters = parameters
	p.KeepInCache = keepInCache
	res := t.onSessionQueryExecute(p)
	return func(tx tableTransactionInfo, prepared bool, result tableResult, e error) {
		var p TableExecuteDataQueryDoneInfo
		p.Tx = tx
		p.Prepared = prepared
		p.Result = result
		p.Error = e
		res(p)
	}
}
func TableOnSessionQueryExplain(t *Table, c *context.Context, session tableSessionInfo, query string) func(aST string, plan string, _ error) {
	var p TableExplainQueryStartInfo
	p.Context = c
	p.Session = session
	p.Query = query
	res := t.onSessionQueryExplain(p)
	return func(aST string, plan string, e error) {
		var p TableExplainQueryDoneInfo
		p.AST = aST
		p.Plan = plan
		p.Error = e
		res(p)
	}
}
func TableOnSessionQueryStreamExecute(t *Table, c *context.Context, session tableSessionInfo, query tableDataQuery, parameters tableQueryParameters) func(error) func(error) {
	var p TableSessionQueryStreamExecuteStartInfo
	p.Context = c
	p.Session = session
	p.Query = query
	p.Parameters = parameters
	res := t.onSessionQueryStreamExecute(p)
	return func(e error) func(error) {
		var p TableSessionQueryStreamExecuteIntermediateInfo
		p.Error = e
		res := res(p)
		return func(e error) {
			var p TableSessionQueryStreamExecuteDoneInfo
			p.Error = e
			res(p)
		}
	}
}
func TableOnSessionQueryStreamRead(t *Table, c *context.Context, session tableSessionInfo) func(error) func(error) {
	var p TableSessionQueryStreamReadStartInfo
	p.Context = c
	p.Session = session
	res := t.onSessionQueryStreamRead(p)
	return func(e error) func(error) {
		var p TableSessionQueryStreamReadIntermediateInfo
		p.Error = e
		res := res(p)
		return func(e error) {
			var p TableSessionQueryStreamReadDoneInfo
			p.Error = e
			res(p)
		}
	}
}
func TableOnSessionTransactionBegin(t *Table, c *context.Context, session tableSessionInfo) func(tx tableTransactionInfo, _ error) {
	var p TableSessionTransactionBeginStartInfo
	p.Context = c
	p.Session = session
	res := t.onSessionTransactionBegin(p)
	return func(tx tableTransactionInfo, e error) {
		var p TableSessionTransactionBeginDoneInfo
		p.Tx = tx
		p.Error = e
		res(p)
	}
}
func TableOnSessionTransactionExecute(t *Table, c *context.Context, session tableSessionInfo, tx tableTransactionInfo, query tableDataQuery, parameters tableQueryParameters) func(result tableResult, _ error) {
	var p TableTransactionExecuteStartInfo
	p.Context = c
	p.Session = session
	p.Tx = tx
	p.Query = query
	p.Parameters = parameters
	res := t.onSessionTransactionExecute(p)
	return func(result tableResult, e error) {
		var p TableTransactionExecuteDoneInfo
		p.Result = result
		p.Error = e
		res(p)
	}
}
func TableOnSessionTransactionExecuteStatement(t *Table, c *context.Context, session tableSessionInfo, tx tableTransactionInfo, statementQuery tableDataQuery, parameters tableQueryParameters) func(result tableResult, _ error) {
	var p TableTransactionExecuteStatementStartInfo
	p.Context = c
	p.Session = session
	p.Tx = tx
	p.StatementQuery = statementQuery
	p.Parameters = parameters
	res := t.onSessionTransactionExecuteStatement(p)
	return func(result tableResult, e error) {
		var p TableTransactionExecuteStatementDoneInfo
		p.Result = result
		p.Error = e
		res(p)
	}
}
func TableOnSessionTransactionCommit(t *Table, c *context.Context, session tableSessionInfo, tx tableTransactionInfo) func(error) {
	var p TableSessionTransactionCommitStartInfo
	p.Context = c
	p.Session = session
	p.Tx = tx
	res := t.onSessionTransactionCommit(p)
	return func(e error) {
		var p TableSessionTransactionCommitDoneInfo
		p.Error = e
		res(p)
	}
}
func TableOnSessionTransactionRollback(t *Table, c *context.Context, session tableSessionInfo, tx tableTransactionInfo) func(error) {
	var p TableSessionTransactionRollbackStartInfo
	p.Context = c
	p.Session = session
	p.Tx = tx
	res := t.onSessionTransactionRollback(p)
	return func(e error) {
		var p TableSessionTransactionRollbackDoneInfo
		p.Error = e
		res(p)
	}
}
func TableOnPoolStateChange(t *Table, size int, event string) {
	var p TablePoolStateChangeInfo
	p.Size = size
	p.Event = event
	t.onPoolStateChange(p)
}
func TableOnPoolSessionAdd(t *Table, session tableSessionInfo) {
	var p TablePoolSessionAddInfo
	p.Session = session
	t.onPoolSessionAdd(p)
}
func TableOnPoolSessionRemove(t *Table, session tableSessionInfo) {
	var p TablePoolSessionRemoveInfo
	p.Session = session
	t.onPoolSessionRemove(p)
}
func TableOnPoolSessionNew(t *Table, c *context.Context) func(session tableSessionInfo, _ error) {
	var p TablePoolSessionNewStartInfo
	p.Context = c
	res := t.onPoolSessionNew(p)
	return func(session tableSessionInfo, e error) {
		var p TablePoolSessionNewDoneInfo
		p.Session = session
		p.Error = e
		res(p)
	}
}
func TableOnPoolSessionClose(t *Table, c *context.Context, session tableSessionInfo) func() {
	var p TablePoolSessionCloseStartInfo
	p.Context = c
	p.Session = session
	res := t.onPoolSessionClose(p)
	return func() {
		var p TablePoolSessionCloseDoneInfo
		res(p)
	}
}
func TableOnPoolPut(t *Table, c *context.Context, session tableSessionInfo) func(error) {
	var p TablePoolPutStartInfo
	p.Context = c
	p.Session = session
	res := t.onPoolPut(p)
	return func(e error) {
		var p TablePoolPutDoneInfo
		p.Error = e
		res(p)
	}
}
func TableOnPoolGet(t *Table, c *context.Context) func(session tableSessionInfo, attempts int, _ error) {
	var p TablePoolGetStartInfo
	p.Context = c
	res := t.onPoolGet(p)
	return func(session tableSessionInfo, attempts int, e error) {
		var p TablePoolGetDoneInfo
		p.Session = session
		p.Attempts = attempts
		p.Error = e
		res(p)
	}
}
func TableOnPoolWait(t *Table, c *context.Context) func(session tableSessionInfo, _ error) {
	var p TablePoolWaitStartInfo
	p.Context = c
	res := t.onPoolWait(p)
	return func(session tableSessionInfo, e error) {
		var p TablePoolWaitDoneInfo
		p.Session = session
		p.Error = e
		res(p)
	}
}
