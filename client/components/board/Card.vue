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
            ><v-text-field v-model="cardTitle" type="text"
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
        <v-card-text>
          期限:
          <v-menu v-model="deadlinePicker">
            <template v-slot:activator="{ on }">
              <v-btn outlined v-on="on">
                <template v-if="deadline">{{
                  new Date(deadline).toLocaleDateString()
                }}</template>
                <template v-else-if="!card.dead_line">なし</template>
                <template v-else>
                  {{ new Date(card.dead_line).toLocaleDateString() }}
                </template>
              </v-btn>
            </template>
            <v-date-picker v-model="deadline"></v-date-picker>
          </v-menu>
        </v-card-text>
        <v-card-title>
          説明:
        </v-card-title>
        <v-card-text>
          {{ card.describe }}
        </v-card-text>
        <v-card-actions>
          <v-spacer></v-spacer>
          <v-btn color="primary" @click="updateDeadline">保存</v-btn>
        </v-card-actions>
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
  deadline = ''
  deadlinePicker = false
  isEdit = false
  cardTitle = ''

  created() {
    this.cardTitle = this.card.title
  }

  @Prop({ type: Object, required: true })
  card!: object

  @Prop({ type: Number, required: true })
  cardIndex!: number

  updateDeadline() {
    const payload = {
      deadline: this.deadline + ' 00:00:00',
      id: this.card.id
    }
    console.log(payload)
    this.$axios
      .put('/card/deadline', payload, {
        headers: {
          Authorization: 'Bearer ' + this.$store.getters['auth/getAuthToken']
        }
      })
      .then((res: any) => {
        console.log(res.data)
      })
      .catch((err: any) => {
        console.log(err)
      })
    this.dialog = false
  }

  toggleIsEdit(): void {
    this.isEdit = !this.isEdit
  }

  updateCardTitle(): void {
    const payload = {
      id: this.card.id,
      title: this.cardTitle,
      token: this.$store.getters['auth/getAuthToken']
    }

    console.log(payload)
    this.$store.dispatch('board/updateCardTitle', payload)
  }
}
</script>
