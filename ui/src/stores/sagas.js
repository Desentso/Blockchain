import {all, takeLatest, takeEvery, take} from 'redux-saga/effects'
import {
  DATA_LOAD, loadDataRequest, 
  BALANCE_GET, getBalanceRequest,
  TRANSACTIONS_GET, getTransactionsRequest
} from "./reducers/basicData"

import {BLOCKCHAIN_LOAD, getBlockchainRequest} from "./reducers/blockchain"

function * loadDataSaga() {
  yield takeEvery(DATA_LOAD, loadDataRequest)
}

function * getBalanceSaga() {
  yield takeEvery(BALANCE_GET, getBalanceRequest)
}

function * getTransactionsSaga() {
  yield takeEvery(TRANSACTIONS_GET, getTransactionsRequest)
}

function * getBlockchainSaga() {
  yield takeEvery(BLOCKCHAIN_LOAD, getBlockchainRequest)
}

export default function * root () {
  yield all([
    loadDataSaga(),
    getBalanceSaga(),
    getTransactionsSaga(),
    getBlockchainSaga()
  ])
}