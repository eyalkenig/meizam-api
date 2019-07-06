import Vue from 'vue'
import Vuex from 'vuex'
import teamsModule from './store/modules/teams/index'

Vue.use(Vuex)

export default new Vuex.Store({
  modules: {
    teams: teamsModule
  }
})
