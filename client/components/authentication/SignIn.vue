<template>
  <div>
    <h2>ログイン</h2>
    <v-form v-model="isValid">
      <TextField
        :text-rules="emailRules"
        :text-label="'email'"
        :text-type="'email'"
        @submit="onReceiveEmail"
      ></TextField>
      <TextField
        :text-rules="passRules"
        :max-length="10"
        :text-label="'password'"
        :text-type="'password'"
        @submit="onReceivePassword"
      ></TextField>
      <BaseButton
        :value="'ログイン'"
        :is-valid="isValid"
        @action="loginAction()"
      ></BaseButton>
    </v-form>
  </div>
</template>

<script lang="ts">
import { Vue, Prop, Component } from 'nuxt-property-decorator'
import TextField from '@/components/atom/TextField.vue'
import BaseButton from '@/components/atom/Button.vue'

@Component({
  components: {
    TextField,
    BaseButton
  }
})
export default class SignIn extends Vue {
  email = ''
  password = ''
  isValid = true
  onReceiveEmail(emailData: string) {
    this.email = emailData
  }

  onReceivePassword(passwordData: string) {
    this.password = passwordData
  }

  async loginAction() {
    await this.$store.dispatch('auth/login', {
      email: this.email,
      password: this.password
    })
    if (this.$store.state.auth.isAuth) {
      this.$router.push('/home')
    }
  }

  @Prop({ type: Array, required: true })
  emailRules!: Array<string>

  @Prop({ type: Array, required: true })
  passRules!: Array<string>
}
</script>
