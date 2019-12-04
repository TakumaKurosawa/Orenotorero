<template>
  <div>
    <h2>BoardCanvas</h2>
    <SignBoard></SignBoard>
  </div>
</template>

<script lang="ts">
import { Vue, Component } from 'vue-property-decorator'

@Component({
  components: {
    SignBoard: () => import('./signboard.vue')
  }
})
export default class BoardCanvas extends Vue {}
</script>
