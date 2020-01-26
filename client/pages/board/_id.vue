<template>
  <div>
    <h1>ボードページ</h1>
    <BoardHeader></BoardHeader>
    <Board></Board>
  </div>
</template>

<script lang="ts">
import { Vue, Component } from 'nuxt-property-decorator'
import Board from '~/components/board/Board.vue'
import BoardHeader from '@/components/board/BoardHeader.vue'

@Component({
  components: {
    Board,
    BoardHeader
  }
})
export default class BoardTop extends Vue {}
</script>
