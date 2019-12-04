<template>
  <div>
    <h1>プロフィールページ</h1>
    <Profile></Profile>
    <Config></Config>
  </div>
</template>

<script lang="ts">
import { Vue, Component } from 'vue-property-decorator'

@Component({
  components: {
    Profile: () => import('../../components/account/profile.vue'),
    Config: () => import('../../components/account/config.vue')
  }
})
export default class AccountTop extends Vue {}
</script>
