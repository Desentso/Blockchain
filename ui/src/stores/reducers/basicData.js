import {getRequest, postRequest} from "../../utils/requests"
import {call, put, select} from "redux-saga/effects"

const initialState = {abc: 123}

export const DATA_LOAD = 'DATA_LOAD'
export const DATA_LOAD_SUCCESS = 'DATA_LOAD_SUCCESS'
export const DATA_LOAD_ERROR = 'DATA_LOAD_ERROR'

export const BALANCE_GET = "BALANCE_GET"
export const BALANCE_GET_SUCCESS = "BALANCE_GET_SUCCESS"
export const BALANCE_GET_ERROR = "BALANCE_GET_ERROR"

export const TRANSACTIONS_GET = "TRANSACTIONS_GET"
export const TRANSACTIONS_GET_SUCCESS = "TRANSACTIONS_GET_SUCCESS"
export const TRANSACTIONS_GET_ERROR = "BALANCE_GET_ERROR"


export const loadData = () => ({type: DATA_LOAD})
export const getBalance = () => ({type: BALANCE_GET})
export const getTransactions = () => ({type: TRANSACTIONS_GET})

export default (state = initialState, action) => {
  switch (action.type) {

    case DATA_LOAD_SUCCESS:
      return {
        ...state,
        address: action.address
      }

    case DATA_LOAD_ERROR:
      return {
        ...state,
        error: "Error while loading address"
      }

    case BALANCE_GET_SUCCESS:
      return {
        ...state,
        balance: action.balance
      }

    case BALANCE_GET_ERROR:
      return {
        ...state,
        error: "Error while loading balance"
      }

    case TRANSACTIONS_GET_SUCCESS:
      return {
        ...state,
        finishedTransactions: action.transactions.finished,
        pendingTransactions: action.transactions.pending
      }

    case TRANSACTIONS_GET_ERROR:
      return {
        ...state,
        error: "Error while loading transactions"
      }

    default:
      return state
  }
}

export function * loadDataRequest() {
  try {
    const resp = yield call(getRequest, "/utils/getOwnAddress")

    yield put({type: DATA_LOAD_SUCCESS, address: resp.address})
  } catch {
    yield put({type: DATA_LOAD_ERROR})
  }
}

export function * getBalanceRequest() {
  try {
    const resp = yield call(getRequest, "/utils/getBalance")

    yield put({type: BALANCE_GET_SUCCESS, balance: resp.balance})
  } catch {
    yield put({type: BALANCE_GET_ERROR})
  }
}

export function * getTransactionsRequest() {
  try {
    const state = yield select()
    const resp = yield call(postRequest, "/utils/transactions", {address: state.data.address})

    yield put({type: TRANSACTIONS_GET_SUCCESS, transactions: resp})
    
  } catch {
    yield put({type: TRANSACTIONS_GET_ERROR})
  }
}
