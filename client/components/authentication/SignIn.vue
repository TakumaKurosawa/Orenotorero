<template>
  <div>
    <h2>ログイン</h2>
    <BaseButton :is-valid="isValid" @submit="loginAction()"></BaseButton>
  </div>
</template>

<script lang="ts">
import { Vue, Component } from 'nuxt-property-decorator'
import BaseButton from '@/components/Button.vue'

@Component({
  components: {
    BaseButton
  }
})
export default class SignIn extends Vue {
  isValid = true
  loginAction() {
    console.log('ログインできますで')
  }
}
</script>
