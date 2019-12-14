<template>
  <v-text-field
    :rules="textRules"
    :counter="maxLength"
    :label="textLabel"
    :type="textType"
    :value="textValue"
  ></v-text-field>
</template>

<script lang="ts">
import { Vue, Prop, Component } from 'vue-property-decorator'

@Component
export default class TextField extends Vue {
  @Prop({ type: Array, required: false })
  textRules!: Array<string>

  @Prop({ type: Number, required: false })
  maxLength!: number

  @Prop({ type: String, required: true })
  textLabel!: string

  @Prop({ type: String, required: false })
  textType!: string

  @Prop({ type: String, required: false, default: '' })
  textValue!: string
}
</script>
