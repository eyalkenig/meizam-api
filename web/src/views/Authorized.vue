<template>
  <div class="hello">
    <h1>Hey</h1>
    <h2>You Are Authorized!</h2>
  </div>
</template>

<script lang="ts">
import { Component, Prop, Vue } from 'vue-property-decorator'
import auth from '../services/auth-service'

@Component
export default class Authorized extends Vue {
  async created () {
    try {
      await auth.handleAuthentication()
    } catch (e) {
      this.$router.push('/')
    }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
h3 {
  margin: 40px 0 0;
}
ul {
  list-style-type: none;
  padding: 0;
}
li {
  display: inline-block;
  margin: 0 10px;
}
a {
  color: #42b983;
}
</style>
