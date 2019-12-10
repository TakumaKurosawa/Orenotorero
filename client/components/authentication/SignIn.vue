<template>
  <div>
    <h2>ログイン</h2>
    <MyButton :is-valid="isValid" @submit="loginAction()"></MyButton>
  </div>
</template>

<script lang="ts">
import { Vue, Component } from 'nuxt-property-decorator'
import MyButton from '../../components/button.vue'

@Component({
  components: {
    MyButton
  }
})
export default class SignIn extends Vue {
  isValid = true
  loginAction() {
    console.log('ログインできますで')
  }
}
</script>
