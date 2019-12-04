<template>
  <div>
    <h1>ボードページ</h1>
    <BoardHeader></BoardHeader>
    <h2>{{ boardData }}</h2>
    <BoardCanvas></BoardCanvas>
    <!--
          ここにボードコンポーネント配置
          -->
  </div>
</template>

<script lang="ts">
import { Vue, Component } from 'vue-property-decorator'

@Component({
  components: {
    BoardCanvas: () => import('../../components/board/boardcanvas.vue'),
    BoardHeader: () => import('../../components/board/boardheader.vue')
  }
})
export default class BoardTop extends Vue {
  boardData: object = this.$store.state.boardData
}
</script>
