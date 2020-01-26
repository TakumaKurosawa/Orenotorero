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
      <v-card-actions>
        <v-btn block text>カード追加</v-btn>
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
