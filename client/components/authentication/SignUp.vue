<template>
  <div>
    <h2>新規登録</h2>
    <v-form v-model="isValid">
      <TextField
        :text-label="'name'"
        :text-type="'text'"
        @submit="onReceiveName"
      ></TextField>
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
        :value="'新規登録'"
        :is-valid="isValid"
        @action="signUpAction()"
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
export default class SignUp extends Vue {
  name = ''
  email = ''
  password = ''
  isValid = false
  onReceiveName(nameData: string) {
    this.name = nameData
  }

  onReceiveEmail(emailData: string) {
    this.email = emailData
  }

  onReceivePassword(passwordData: string) {
    this.password = passwordData
  }

  signUpAction() {
    this.$store.dispatch('auth/signup', {
      name: this.name,
      email: this.email,
      password: this.password
    })
  }

  @Prop({ type: Array, required: true })
  emailRules!: Array<string>

  @Prop({ type: Array, required: true })
  passRules!: Array<string>
}
</script>
