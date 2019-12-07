<template>
  <v-btn :disabled="!isValid" @click="submit">完了</v-btn>
</template>

<script lang="ts">
import { Vue, Prop, Component } from 'vue-property-decorator'

@Component
export default class MyButton extends Vue {
  submit() {
    this.$emit('submit')
  }

  @Prop({ type: Boolean, required: true })
  isValid: boolean
}
</script>
