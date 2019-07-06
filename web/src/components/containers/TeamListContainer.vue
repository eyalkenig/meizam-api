<template>
  <div>
      <div>Teams ({{ teams.length }}):</div>
      <Team v-for="team in teams"
            :key="team.id"
            v-bind:name="team.name"
            v-bind:team-id="team.id">
      </Team>
  </div>
</template>

<script lang="ts">
import { Component, Prop, Vue } from 'vue-property-decorator'
import { FETCH_TEAMS } from '../../store/modules/teams/types'
import Team from '../widgets/Team.vue'

@Component({
  components: {
    Team
  }
})
export default class TeamListContainer extends Vue {
  data () {
    return this.$store.state.teams
  }
  created () {
    this.$store.dispatch(FETCH_TEAMS)
  }
}
</script>

<style scoped>
</style>
