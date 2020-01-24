<template>
  <div>
    <h1>ユーザーTOPページ</h1>
    <MenuList></MenuList>
    <RecentlyUsedBoardList
      :board-data="recentAccessedBoard"
    ></RecentlyUsedBoardList>
    <PersonalBoardList :board-data="boardData"></PersonalBoardList>
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
  boardData = []

  async asyncData({ $axios, store }: any) {
    const boardData: Array<object> = await $axios
      .$get('/board', {
        headers: {
          Authorization: 'Bearer ' + store.getters['auth/getAuthToken']
        },
        data: {}
      })
      .catch((err: any) => {
        console.log(err)
      })
    return { boardData }
  }

  get recentAccessedBoard(): Array<object> {
    return this.boardData.slice(0, 4)
  }
}
</script>
