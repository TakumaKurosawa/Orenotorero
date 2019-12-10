<template>
  <v-btn :disabled="!isValid" @click="submit">{{ text }}</v-btn>
</template>

<script lang="ts">
import { Vue, Prop, Component } from 'nuxt-property-decorator'

@Component
export default class BaseButton extends Vue {
  submit() {
    this.$emit('submit')
  }

  @Prop({ type: Boolean, required: true })
  isValid!: boolean

  @Prop({ type: String, default: '完了' })
  text!: string
}
</script>
