import {getRequest} from "../../utils/requests"
import {call, put} from "redux-saga/effects"

const initialState = {
  error: null
}

export const BLOCKCHAIN_LOAD = 'BLOCKCHAIN_LOAD'
export const BLOCKCHAIN_LOAD_SUCCESS = 'BLOCKCHAIN_LOAD_SUCCESS'
export const BLOCKCHAIN_LOAD_ERROR = 'BLOCKCHAIN_LOAD_ERROR'


export const loadBlockchain = () => ({type: BLOCKCHAIN_LOAD})

export default (state = initialState, action) => {
  switch (action.type) {

    case BLOCKCHAIN_LOAD_SUCCESS:
      return {
        ...state,
        blockchain: action.blockchain
      }

    case BLOCKCHAIN_LOAD_ERROR:
      return {
        ...state,
        error: "Error while loading blockchain"
      }

    default:
      return state
  }
}


export function * getBlockchainRequest() {
  try {
    const resp = yield call(getRequest, "/blockchain")

    yield put({type: BLOCKCHAIN_LOAD_SUCCESS, blockchain: resp})
    
  } catch {
    yield put({type: BLOCKCHAIN_LOAD_ERROR})
  }
}
