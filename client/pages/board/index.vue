<template>
  <div>
    <h1>ボードページ</h1>
    <h2>{{ boardData }}</h2>
    <!--
          ここにボードコンポーネント配置
          -->
  </div>
</template>

<script lang="ts">
import { Vue, Component } from 'vue-property-decorator'

@Component
export default class BoardTop extends Vue {
  boardData: object = this.$store.state.boardData
}
</script>
