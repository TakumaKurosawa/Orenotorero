<template>
  <div>
    <h1>ユーザー認証ページ</h1>
    <SignIn></SignIn>
    <SignUp></SignUp>
  </div>
</template>

<script lang="ts">
import { Vue, Component } from 'vue-property-decorator'
import SignIn from '../../components/authentication/signIn.vue'
import SignUp from '../../components/authentication/signUp.vue'

@Component({
  components: {
    SignIn,
    SignUp
  }
})
export default class AuthenticationTop extends Vue {}
</script>
