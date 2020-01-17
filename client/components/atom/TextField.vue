<template>
  <v-text-field
    v-model="textData"
    :rules="textRules"
    :counter="maxLength"
    :label="textLabel"
    :type="textType"
    :value="textValue"
    @input="emitValue"
  ></v-text-field>
</template>

<script lang="ts">
import { Vue, Prop, Component } from 'nuxt-property-decorator'

@Component
export default class TextField extends Vue {
  textData = ''
  emitValue() {
    this.$emit('submit', this.textData)
  }

  @Prop({ type: Array, required: false, default: [] })
  textRules!: Array<string>

  @Prop({ type: Number, required: false, default: 10 })
  maxLength!: number

  @Prop({ type: String, required: true })
  textLabel!: string

  @Prop({ type: String, required: false, default: 'text' })
  textType!: string

  @Prop({ type: String, required: false, default: '' })
  textValue!: string
}
</script>
