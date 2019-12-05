<template>
  <div>
    <h1>ユーザーTOPページ</h1>
    <MenuList></MenuList>
    <BoardList></BoardList>
  </div>
</template>

<script lang="ts">
import { Vue, Component } from 'vue-property-decorator'
import MenuList from '../../components/home/MenuList.vue'
import BoardList from '../../components/home/BoardList.vue'

@Component({
  components: {
    MenuList,
    BoardList
  }
})
export default class ome extends Vue {}
</script>
