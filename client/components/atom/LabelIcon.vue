<template>
  <p>
    <v-icon>{{ iconName }}</v-icon
    >{{ textLabel }}
  </p>
</template>

<script lang="ts">
import { Vue, Prop, Component } from 'nuxt-property-decorator'

@Component
export default class BaseButton extends Vue {
  @Prop({ type: String, required: true })
  textLabel!: string

  @Prop({ type: String, required: false, default: '' })
  iconName!: string
}
</script>
