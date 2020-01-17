<template>
  <div>
    <h2>新規登録</h2>
    <v-form v-model="isValid">
      <TextField :text-label="'name'" :text-type="'text'"></TextField>
      <TextField
        :text-rules="emailRules"
        :text-label="'email'"
        :text-type="'email'"
      ></TextField>
      <TextField
        :text-rules="passRules"
        :max-length="10"
        :text-label="'password'"
        :text-type="'password'"
      ></TextField>
      <BaseButton
        :value="'新規登録'"
        :is-valid="isValid"
        @action="signUpAction()"
      ></BaseButton>
    </v-form>
    <v-btn @click="signUpAction()"> あくしよんテスト</v-btn>
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
  isValid = false
  signUpAction() {
    this.$store.dispatch('auth/signup', {
      name: 'kenshin',
      email: 'hogehoge@gmail.com',
      password: 'passpass'
    })
  }

  @Prop({ type: Array, required: true })
  emailRules!: Array<string>

  @Prop({ type: Array, required: true })
  passRules!: Array<string>
}
</script>
