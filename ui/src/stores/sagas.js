import {all, takeLatest, takeEvery, take} from 'redux-saga/effects'
import {DATA_LOAD, loadDataRequest} from "./reducers/basicData"

function * loadDataSaga() {
  yield takeEvery(DATA_LOAD, loadDataRequest)
}

export default function * root () {
  yield all([
    loadDataSaga()
  ])
}