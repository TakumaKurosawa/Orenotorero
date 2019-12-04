<template>
  <div>
    <h1>ユーザー認証ページ</h1>
    <SignIn></SignIn>
    <SignUp></SignUp>
  </div>
</template>

<script lang="ts">
import { Vue, Component } from 'vue-property-decorator'

@Component({
  components: {
    SignIn: () => import('../../components/authentication/signin.vue'),
    SignUp: () => import('../../components/authentication/signup.vue')
  }
})
export default class AuthenticationTop extends Vue {}
</script>
