<template>
  <v-text-field
    :rules="textRules"
    :counter="maxLen"
    :label="textLabel"
    :type="textType"
  ></v-text-field>
</template>

<script lang="ts">
import { Vue, Prop, Component } from 'vue-property-decorator'

@Component
export default class TextField extends Vue {
  @Prop({ type: Array, required: false })
  textRules: Array<string>

  @Prop({ type: Number, required: false })
  maxLen: number

  @Prop({ type: String, required: true })
  textLabel: string

  @Prop({ type: String, required: false })
  textType: string
}
</script>
