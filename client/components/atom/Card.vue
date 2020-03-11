<template>
  <v-card
    :width="cardWidth"
    :height="cardHeight"
    :img="cardImg"
    :color="cardColor"
    @click="submit()"
  >
    <v-card-title>{{ cardTitle }}</v-card-title>
  </v-card>
</template>

<script lang="ts">
import { Vue, Prop, Component } from 'nuxt-property-decorator'

@Component
export default class TextField extends Vue {
  submit() {
    this.$emit('action')
  }

  @Prop({ type: String, required: true })
  cardTitle!: string

  @Prop({ type: String, required: false })
  cardImg!: string

  @Prop({ type: String, required: false })
  cardColor!: string

  @Prop({ type: Number, required: true })
  cardWidth!: number

  @Prop({ type: Number, required: true })
  cardHeight!: number
}
</script>
