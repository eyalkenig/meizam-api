<template>
  <div class="hello">
      <h1>Name: {{ this.name }}</h1>
      <h3>ID: {{ this.teamId }}</h3>
  </div>
</template>

<script lang="ts">
import { Component, Prop, Vue } from 'vue-property-decorator'
import { FETCH_TEAMS } from '../../store/modules/teams/types'

@Component
export default class Team extends Vue {
    @Prop() private teamId!: number;
    @Prop() private name!: string;
}
</script>

<style scoped>
</style>
