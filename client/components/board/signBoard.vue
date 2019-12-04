<template>
  <div>
    <h3>SignBoard</h3>
    <TaskCard></TaskCard>
  </div>
</template>

<script lang="ts">
import { Vue, Component } from 'vue-property-decorator'
import TaskCard from './taskCard.vue'

@Component({
  components: {
    TaskCard
  }
})
export default class SignBoard extends Vue {}
</script>
