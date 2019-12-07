<template>
  <div>
    <h2>ログイン</h2>
    <myButton :is-valid="isValid" @submit="loginAction()"></myButton>
  </div>
</template>

<script lang="ts">
import { Vue, Component } from 'vue-property-decorator'
import myButton from '../../components/button.vue'

@Component({
  components: {
    myButton
  }
})
export default class SignIn extends Vue {
  isValid = true
  loginAction() {
    console.log('ログインできますで')
  }
}
</script>
