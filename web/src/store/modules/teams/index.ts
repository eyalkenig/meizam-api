import { mutations } from './mutations'
import { actions } from './actions'
import { TeamState } from './types'

export const initialStateData: TeamState = {
  teams: []
}

export default {
  state: { ...initialStateData },
  mutations,
  actions
}
