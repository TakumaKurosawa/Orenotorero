<template>
  <div>
    <h3>SignBoard</h3>
    <TaskCard></TaskCard>
  </div>
</template>

<script lang="ts">
import { Vue, Component } from 'vue-property-decorator'

@Component({
  components: {
    TaskCard: () => import('./taskcard.vue')
  }
})
export default class SignBoard extends Vue {}
</script>
