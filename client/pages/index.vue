<template>
  <div>
    <nuxt-link to="/authentication">ログインページに移動<br /></nuxt-link>
    <nuxt-link to="/home">ユーザートップページに移動<br /></nuxt-link>
    <nuxt-link to="/board">ボードページに移動<br /></nuxt-link>
    <nuxt-link to="/account">プロフィールページに移動</nuxt-link>
  </div>
</template>

<script lang="ts">
import { Vue, Component } from 'nuxt-property-decorator'

@Component
export default class Top extends Vue {
  created() {
    if (this.$store.state.auth.isAuth) {
      this.$router.push('/home')
    } else {
      this.$router.push('/authentication')
    }
  }
}
</script>
