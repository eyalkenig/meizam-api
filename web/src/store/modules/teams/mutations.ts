import { SET_TEAMS, TeamState, Team } from './types'

export const mutations = {
  [SET_TEAMS] (state: TeamState, teams: Team[]) {
    state.teams = teams
  }
}
