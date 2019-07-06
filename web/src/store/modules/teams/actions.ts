import { FETCH_TEAMS, SET_TEAMS } from './types'
import MeizamApi from '../../../services/meizam-api-service'

interface ContextCommit {
  commit: any;
}
export const actions = {
  [FETCH_TEAMS] ({ commit }: ContextCommit) {
    return MeizamApi.fetchTeams()
      .then((response) =>
        commit(SET_TEAMS, response.data))
  }
}