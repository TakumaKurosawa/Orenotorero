<template>
  <div>
    <h2>BoardCanvas</h2>
    <SignBoard></SignBoard>
  </div>
</template>

<script lang="ts">
import { Vue, Component } from 'vue-property-decorator'
import SignBoard from './signBoard.vue'

@Component({
  components: {
    SignBoard
  }
})
export default class BoardCanvas extends Vue {}
</script>
