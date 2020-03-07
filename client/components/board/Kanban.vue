<template>
  <v-col>
    <v-card light min-width="200px" max-width="200px">
      <v-card-title>{{ kanban.title }}</v-card-title>
      <draggable v-model="kanbanData" class="list-group" group="card">
        <card
          v-for="(card, index) in kanbanData"
          :key="index"
          class="list-group-item"
          :card-index="index"
          :card="card"
        ></card>
      </draggable>
      <v-card v-if="isOpenForm" raised class="ma-3">
        <v-form>
          <v-text-field label="カード名を入力"></v-text-field>
        </v-form>
      </v-card>
      <v-card-actions v-if="isOpenForm">
        <v-btn color="primary">カード追加</v-btn>
        <v-btn icon @click="isOpenForm = !isOpenForm">
          <v-icon>
            mdi-close
          </v-icon>
        </v-btn>
      </v-card-actions>
      <v-card-actions v-else>
        <v-btn block text @click="isOpenForm = !isOpenForm"
          >さらにカード追加</v-btn
        >
      </v-card-actions>
    </v-card>
  </v-col>
</template>

<script lang="ts">
import { Vue, Component, Prop } from 'nuxt-property-decorator'
import draggable from 'vuedraggable'
import Card from './Card.vue'

@Component({
  components: {
    draggable,
    Card
  }
})
export default class Kanban extends Vue {
  isOpenForm: boolean = false
  inputCardName: string = ''
  async createCard(title: string, kanbanIndex: number, position: number) {
    const payload = {
      title,
      kanban_id: kanbanIndex,
      position
    }
    await this.$axios
      .post('/card', payload, {
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

  @Prop({ type: Object, required: true })
  kanban!: object

  @Prop({ type: Number, required: true })
  kanbanIndex!: number
}
</script>
