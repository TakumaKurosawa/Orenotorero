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
  @Prop({ type: Array, required: true })
  textLabel!: Array<string>

  @Prop({ type: Array, required: false, default: '' })
  iconName!: Array<string>
}
</script>
