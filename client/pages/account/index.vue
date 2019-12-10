<template>
  <div>
    <h1>プロフィールページ</h1>
    <Profile></Profile>
    <Config></Config>
  </div>
</template>

<script lang="ts">
import { Vue, Component } from 'nuxt-property-decorator'
import Profile from '../../components/account/Profile.vue'
import Config from '../../components/account/Config.vue'

@Component({
  components: {
    Profile,
    Config
  }
})
export default class AccountTop extends Vue {}
</script>
