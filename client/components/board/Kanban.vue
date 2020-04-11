<template>
  <v-col>
    <v-card light min-width="200px" max-width="200px">
      <v-card-title v-if="isEdit">
        {{ kanban.title }}
        <v-spacer></v-spacer>
        <v-icon @click="isEdit = false">mdi-pencil</v-icon>
        <v-icon @click="deleteKanban">mdi-close</v-icon>
      </v-card-title>
      <v-card-title v-if="!isEdit">
        <v-row>
          <v-col>
            <v-text-field
              v-model="kanbanTitle"
              append-icon="mdi-send"
              @click:append="changeKanbanTitle"
            ></v-text-field>
          </v-col>
        </v-row>
      </v-card-title>
      <draggable v-model="kanbanData" class="list-group" group="card">
        <card
          v-for="(card, index) in kanbanData"
          :key="index"
          class="list-group-item"
          :card-index="index"
          :card="card"
        ></card>
      </draggable>
      <v-card-actions>
        <v-btn v-if="!inputTitle" block text @click="inputTitle = true"
          >カード追加</v-btn
        >
        <div v-if="inputTitle">
          <v-col>
            <textField
              :text-rules="titleRules"
              :max-length="30"
              :text-label="'タイトルを入力'"
              :text-type="'text'"
              @submit="onReceiveCardTitle"
            ></textField>
          </v-col>
          <v-col>
            <Button
              :value="'カードを追加'"
              :is-valid="isValid"
              @action="createCard()"
            ></Button>
            <v-icon @click="inputTitle = false">mdi-close</v-icon>
          </v-col>
        </div>
      </v-card-actions>
    </v-card>
  </v-col>
</template>

<script lang="ts">
import { Vue, Component, Prop } from 'nuxt-property-decorator'
import draggable from 'vuedraggable'
import Card from './Card.vue'
import textField from '@/components/atom/TextField.vue'
import Button from '@/components/atom/Button.vue'

@Component({
  components: {
    draggable,
    Card,
    textField,
    Button
  }
})
export default class Kanban extends Vue {
  newCardTitle = ''
  inputTitle = false
  isValid = true
  kanbanTitle = ''
  isEdit = true

  created() {
    this.kanbanTitle = this.kanban.title
  }

  async createCard() {
    const payload = {
      title: this.newCardTitle,
      kanban_id: this.kanban.id,
      position: this.kanban.card.length + 1
    }
    await this.$axios
      .post('/card', payload, {
        headers: {
          Authorization: 'Bearer ' + this.$store.getters['auth/getAuthToken']
        }
      })
      .then((res: any) => {
        console.log(res.data)
        this.$emit('action')
        this.inputTitle = false
      })
      .catch((err: any) => {
        console.log(err)
      })
  }

  get kanbanData() {
    return this.$store.getters['board/boardData'][this.kanbanIndex].card
  }

  set kanbanData(value: Array<object>) {
    this.$store.commit('board/updateKanbanData', {
      index: this.kanbanIndex,
      value
    })
  }

  async deleteKanban() {
    await this.$axios
      .delete('/kanban', {
        headers: {
          Authorization: 'Bearer ' + this.$store.getters['auth/getAuthToken']
        },
        data: {
          id: this.kanban.id
        }
      })
      .then((res: any) => {
        console.log(res.data)
        this.$emit('action')
      })
      .catch((err: any) => {
        console.log(err)
      })
  }

  async changeKanbanTitle() {
    const payload = {
      id: this.kanban.id,
      title: this.kanbanTitle
    }
    await this.$axios
      .put('/kanban', payload, {
        headers: {
          Authorization: 'Bearer ' + this.$store.getters['auth/getAuthToken']
        }
      })
      .then((res: any) => {
        console.log(res.data)
        this.isEdit = true
        this.$emit('action')
      })
      .catch((err: any) => {
        console.log(err)
        this.isEdit = true
      })
  }
  async updatePosition() {
    interface position {
      kanbanId: Number
      cardArray: Array<Number>
    }
    const payload = this.$store.state.board.boardData.map(
      (kanban: { id: Number; card: [] }) => {
        const p: position = {
          kanbanId: kanban.id,
          cardArray: kanban.card.map((card: { id: Number }) => card.id)
        }
        return p
      }
    )
    await this.$axios
      .put('/position', payload, {
        headers: {
          Authorization: 'Bearer ' + this.$store.getters['auth/getAuthToken']
        }
      })
      .catch((err: any) => {
        console.log(err)
      })
  }

  onReceiveCardTitle(cardTitle: string) {
    this.newCardTitle = cardTitle
  }

  @Prop({ type: Array, required: true })
  titleRules!: Array<string>

  @Prop({ type: Object, required: true })
  kanban!: object

  @Prop({ type: Number, required: true })
  kanbanIndex!: number
}
</script>
