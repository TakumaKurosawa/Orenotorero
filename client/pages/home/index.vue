<template>
  <div>
    <h1>ユーザーTOPページ</h1>
    <MenuList></MenuList>
    <BoardList></BoardList>
  </div>
</template>

<script lang="ts">
import { Vue, Component } from 'vue-property-decorator'

@Component({
  components: {
    MenuList: () => import('../../components/home/menulist.vue'),
    BoardList: () => import('../../components/home/boardlist.vue')
  }
})
export default class ome extends Vue {}
</script>
