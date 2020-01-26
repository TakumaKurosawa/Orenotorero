<template>
  <v-container fluid>
    <v-row>
      <draggable v-model="boardData" class="list-group d-flex" group="kanban">
        <kanban
          v-for="(kanban, index) in boardData"
          :key="index"
          class="list-group-item mr-5"
          :kanban-index="index"
          :kanban="kanban"
        ></kanban>
      </draggable>
    </v-row>
  </v-container>
</template>

<script lang="ts">
import { Vue, Component } from 'nuxt-property-decorator'
import draggable from 'vuedraggable'
import kanban from './Kanban.vue'

@Component({
  components: {
    draggable,
    kanban
  }
})
export default class BoardCanvas extends Vue {
  get boardData() {
    return this.$store.getters['board/boardData']
  }

  set boardData(value: Array<object>) {
    this.$store.commit('board/updateBoardData', value)
  }
}
</script>
