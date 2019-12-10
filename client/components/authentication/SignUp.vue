<template>
  <div>
    <h2>新規登録</h2>
    <TextField
      :text-rules="emailRules"
      :text-label="'email'"
      :text-type="'email'"
    ></TextField>
    <TextField
      :text-rules="passRules"
      :max-len="32"
      :text-label="'password'"
      :text-type="'password'"
    ></TextField>
  </div>
</template>

<script lang="ts">
import { Vue, Prop, Component } from 'nuxt-property-decorator'
import TextField from '../../components/TextField.vue'

@Component({
  components: {
    TextField
  }
})
export default class SignUp extends Vue {
  @Prop({ type: Array, required: true })
  emailRules: Array<string>

  @Prop({ type: Array, required: true })
  passRules: Array<string>
}
</script>
