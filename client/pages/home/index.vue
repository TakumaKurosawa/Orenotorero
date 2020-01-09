<template>
  <div>
    <h1>ユーザーTOPページ</h1>
    <MenuList></MenuList>
    <RecentlyUsedBoardList :board-data="boardData"></RecentlyUsedBoardList>
    <PersonalBoardList :board-data="boardData"></PersonalBoardList>
  </div>
</template>

<script lang="ts">
import { Vue, Component } from 'nuxt-property-decorator'
import MenuList from '@/components/home/MenuList.vue'
import RecentlyUsedBoardList from '@/components/home/RecentlyUsedBoardList.vue'
import PersonalBoardList from '@/components/home/PersonalBoardList.vue'
import BoardData from '@/assets/mock/board.json'

@Component({
  components: {
    MenuList,
    RecentlyUsedBoardList,
    PersonalBoardList
  }
})
export default class Home extends Vue {
  boardData = BoardData
}
</script>
