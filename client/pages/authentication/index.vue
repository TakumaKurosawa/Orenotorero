<template>
  <div>
    <h1>{{ boo }}</h1>
    <v-tabs color="white" centered dark>
      <v-tabs-slider color="white"></v-tabs-slider>
      <v-tab v-on:click="indicateLogin = true">
        ログイン
      </v-tab>
      <v-tab v-on:click="indicateLogin = false">
        新規登録
      </v-tab>
    </v-tabs>
    <SignIn v-if="indicateLogin == true"></SignIn>
    <SignUp v-if="indicateLogin == false"></SignUp>
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
}
</script>
