<template>
  <div>
    <h1>ユーザー認証ページ</h1>
    <v-tabs color="white" centered dark>
      <v-tabs-slider color="white"></v-tabs-slider>
      <v-tab @click="indicateLogin = true">
        ログイン
      </v-tab>
      <v-tab @click="indicateLogin = false">
        新規登録
      </v-tab>
    </v-tabs>
    <SignIn v-if="indicateLogin"></SignIn>
    <SignUp v-if="!indicateLogin"></SignUp>
  </div>
</template>

<script lang="ts">
import { Vue, Component } from 'vue-property-decorator'
import SignIn from '../../components/authentication/SignIn.vue'
import SignUp from '../../components/authentication/SignUp.vue'

@Component({
  components: {
    SignIn,
    SignUp
  }
})
export default class AuthenticationTop extends Vue {
  indicateLogin: boolean = true
  emailRules = [
    (v: string) => !!v || 'emailの入力は必須です',
    (v: string) => /.+@.+/.test(v) || 'emailが正しくありません'
  ]

  passRules = [
    (v: string) => !!v || 'passの入力は必須です',
    (v: string) => v.length <= 10 || 'passは10文字以内で入力してください'
  ]
}
</script>
