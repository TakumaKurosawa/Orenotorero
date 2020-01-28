<template>
  <div>
    <v-card raised class="ma-3">
      <v-card-text
        class="text-truncate font-weight-bold subtitle-1"
        @click.stop="dialog = true"
      >
        {{ card.title }}
      </v-card-text>
    </v-card>
    <v-dialog v-model="dialog" max-width="500">
      <v-card>
        <v-card-actions>
          <v-spacer></v-spacer>
          <v-icon @click="dialog = false">mdi-close</v-icon>
        </v-card-actions>
        <v-card-title class="headline">
          {{ card.title }}
        </v-card-title>
        <v-card-title>
          説明:
        </v-card-title>
        <v-card-text>
          {{ card.describe }}
        </v-card-text>
      </v-card>
    </v-dialog>
  </div>
</template>

<script lang="ts">
import { Vue, Component } from 'nuxt-property-decorator'
import { Prop } from '~/node_modules/nuxt-property-decorator'

@Component
export default class Card extends Vue {
  dialog = false
  @Prop({ type: Object, required: true })
  card!: object

  @Prop({ type: Number, required: true })
  cardIndex!: number
}
</script>
