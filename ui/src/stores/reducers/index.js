import {combineReducers} from 'redux'
import data from "./basicData"
import blockchain from "./blockchain"

export default combineReducers({
	data,
	blockchain
})