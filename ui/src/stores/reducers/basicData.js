import {getRequest} from "../../utils/requests"
import {call, put} from "redux-saga/effects"

const initialState = {abc: 123}

export const DATA_LOAD = 'DATA_LOAD'
export const DATA_LOAD_SUCCESS = 'DATA_LOAD_SUCCESS'
export const DATA_LOAD_ERROR = 'DATA_LOAD_ERROR'

export const loadData = () => ({type: DATA_LOAD})

export default (state = initialState, action) => {
  switch (action.type) {
    /*case DATA_LOAD:
      console.log("load data")
      return state*/

    case DATA_LOAD_SUCCESS:
      state = action.data
      return state

    case DATA_LOAD_ERROR:
      state.error = "Error while loading data"
      return state

    default:
      return state
  }
}

export function * loadDataRequest() {
  //yield
  try {
    const resp = yield call(getRequest, "/utils/getOwnAddress")

    yield put({type: DATA_LOAD_SUCCESS, data: resp})
  } catch {
    yield put({type: DATA_LOAD_ERROR})
  }

}