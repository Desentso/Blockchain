import {combineReducers} from 'redux'
import data from "./basicData"
import blockchain from "./blockchain"
import mining from "./mining"

export default combineReducers({
	data,
	blockchain,
	mining
})