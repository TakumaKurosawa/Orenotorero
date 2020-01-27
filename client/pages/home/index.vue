<template>
  <div>
    <h1>ユーザーTOPページ</h1>
    <MenuList></MenuList>
    <RecentlyUsedBoardList
      :board-data="recentAccessedBoard"
    ></RecentlyUsedBoardList>
    <PersonalBoardList :board-data="myBoardsData"></PersonalBoardList>
    <v-dialog v-model="isOpenCreateBoard" width="500">
      <template v-slot:activator="{ on }">
        <v-card
          class="d-flex align-center justify-center"
          :width="180"
          :height="100"
          color="gray"
          v-on="on"
        >
          <v-card-title class="subtitle-1">新しいボードの作成</v-card-title>
        </v-card>
      </template>
      <v-card>
        <v-card-title class="headline grey" primary-title>
          ボード作成
        </v-card-title>

        <v-divider></v-divider>

        <v-form v-model="isValid">
          <v-text-field
            v-model="title"
            :rules="rules"
            label="ボードタイトルを入力"
            required
          ></v-text-field>
          <v-text-field
            v-model="img"
            :rules="rules"
            label="url"
            required
          ></v-text-field>
          <v-card-actions>
            <v-btn :disabled="!isValid" class="mr-4" @click="createBoard()">
              ボードを作成
            </v-btn>
          </v-card-actions>
        </v-form>
      </v-card>
    </v-dialog>
  </div>
</template>

<script lang="ts">
import { Vue, Component } from 'nuxt-property-decorator'
import MenuList from '@/components/home/MenuList.vue'
import RecentlyUsedBoardList from '@/components/home/RecentlyUsedBoardList.vue'
import PersonalBoardList from '@/components/home/PersonalBoardList.vue'

@Component({
  components: {
    MenuList,
    RecentlyUsedBoardList,
    PersonalBoardList
  }
})
export default class Home extends Vue {
  isValid = false
  isOpenCreateBoard = false
  title = ''
  img =
    'data:image/jpeg;base64,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'

  rules = [(v: string) => !!v || '入力は必須です']

  m() {
    this.$store.dispatch('myBoards/fetchMyBoardsData')
  }

  createBoard() {
    this.isOpenCreateBoard = false
    this.$store.dispatch('myBoards/createBoard', {
      title: this.title,
      img: this.img
    })
  }

  get myBoardsData(): Array<object> {
    return this.$store.getters['myBoards/myBoardsData']
  }

  get recentAccessedBoard(): Array<object> {
    return this.myBoardsData.slice(0, 4)
  }
}
</script>
