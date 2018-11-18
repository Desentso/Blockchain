const initialState = {mining: false, intervalID: null}

export const TOGGLE_MINING = 'TOGGLE_MINING'

export const toggleMining = (intervalID) => ({type: TOGGLE_MINING, intervalID})

export default (state = initialState, action) => {
  switch (action.type) {

    case TOGGLE_MINING:
      return {
        ...state,
        mining: !state.mining,
        intervalID: action.intervalID
      }

    default:
      return state
  }
}
