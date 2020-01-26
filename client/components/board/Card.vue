<template>
  <v-card raised class="ma-3">
    <v-card-text class="text-truncate font-weight-bold subtitle-1">
      {{ card.title }}
    </v-card-text>
  </v-card>
</template>

<script lang="ts">
import { Vue, Component } from 'nuxt-property-decorator'
import { Prop } from '~/node_modules/nuxt-property-decorator'

@Component
export default class Card extends Vue {
  @Prop({ type: Object, required: true })
  card!: object

  @Prop({ type: Number, required: true })
  cardIndex!: number
}
</script>
