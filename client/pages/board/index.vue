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
import BoardCanvas from '../../components/board/BoardCanvas.vue'
import BoardHeader from '../../components/board/BoardHeader.vue'

@Component({
  components: {
    BoardCanvas,
    BoardHeader
  }
})
export default class BoardTop extends Vue {
  boardData: object = this.$store.state.boardData
}
</script>
