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
          <span v-if="isEdit"
            ><v-text-field :value="card.title" type="text"
          /></span>
          <span v-else>{{ card.title }}</span>
          <a v-if="isEdit" @click="updateCardTitle"
            ><v-icon>mdi-send</v-icon></a
          >
          <a v-if="isEdit" @click="toggleIsEdit"
            ><v-icon>mdi-pencil-remove</v-icon></a
          >
          <a v-else @click="toggleIsEdit"
            ><v-icon>mdi-square-edit-outline</v-icon></a
          >
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
  isEdit = false

  @Prop({ type: Object, required: true })
  card!: object

  @Prop({ type: Number, required: true })
  cardIndex!: number

  toggleIsEdit(): void {
    this.isEdit = !this.isEdit
  }

  updateCardTitle(): void {
    const payload = {
      id: this.cardIndex,
      title: this.card.title,
      token: this.$store.getters['auth/getAuthToken']
    }

    console.log(payload)
    this.$store.dispatch('board/updateCardTitle', payload)
  }
}
</script>
