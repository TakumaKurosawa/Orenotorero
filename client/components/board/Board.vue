<template>
  <div>
    <v-container fluid>
      <v-row>
        <draggable v-model="boardData" class="list-group d-flex" group="kanban">
          <kanban
            v-for="(kanban, index) in boardData"
            :key="index"
            class="list-group-item mr-5"
            :title-rules="titleRules"
            :kanban-index="index"
            :kanban="kanban"
            @action="getBoardData()"
          ></kanban>
        </draggable>
        <v-col>
          <v-card light min-width="200px" max-width="200px">
            <v-card-actions>
              <v-btn v-if="!inputTitle" block text @click="inputTitle = true"
                >カンバン追加</v-btn
              >
              <div v-if="inputTitle">
                <v-col>
                  <textField
                    :text-rules="titleRules"
                    :max-length="30"
                    :text-label="'タイトルを入力'"
                    :text-type="'text'"
                    @submit="onReceiveKanbanTitle"
                  ></textField>
                </v-col>
                <v-col>
                  <Button
                    :value="'カンバンを追加'"
                    :is-valid="isValid"
                    @action="createKanban()"
                  ></Button>
                  <v-icon @click="inputTitle = false">mdi-close</v-icon>
                </v-col>
              </div>
            </v-card-actions>
          </v-card>
        </v-col>
      </v-row>
    </v-container>
  </div>
</template>

<script lang="ts">
import { Vue, Component } from 'nuxt-property-decorator'
import draggable from 'vuedraggable'
import kanban from './Kanban.vue'
import textField from '@/components/atom/TextField.vue'
import Button from '@/components/atom/Button.vue'

@Component({
  components: {
    draggable,
    kanban,
    textField,
    Button
  }
})
export default class BoardCanvas extends Vue {
  newKanbanTitle = ''
  inputTitle = false
  isValid = true
  titleRules = [
    (v: string) => !!v || 'titleの入力は必須です',
    (v: string) => v.length <= 30 || '30文字以内で入力してください'
  ]

  created() {
    this.getBoardData()
  }

  async createKanban() {
    const payload = {
      title: this.newKanbanTitle,
      board_id: Number(this.$route.params.id),
      position: this.boardData.length + 1
    }
    await this.$axios
      .post('/kanban', payload, {
        headers: {
          Authorization: 'Bearer ' + this.$store.getters['auth/getAuthToken']
        }
      })
      .then((res: any) => {
        console.log(res.data)
        this.getBoardData()
        this.inputTitle = false
      })
      .catch((err: any) => {
        console.log(err)
      })
  }

  get boardData() {
    return this.$store.getters['board/boardData']
  }

  set boardData(value: Array<object>) {
    this.$store.commit('board/updateBoardData', value)
  }

  getBoardData() {
    this.$store.dispatch('board/getBoardData', {
      boardId: this.$route.params.id,
      token: this.$store.getters['auth/getAuthToken']
    })
  }

  onReceiveKanbanTitle(kanbanTitle: string) {
    this.newKanbanTitle = kanbanTitle
  }
}
</script>
