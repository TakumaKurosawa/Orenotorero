<template>
  <div>
    <h3>SignBoard</h3>
    <TaskCard></TaskCard>
  </div>
</template>

<script lang="ts">
import { Vue, Component } from 'nuxt-property-decorator'
import TaskCard from './TaskCard.vue'

@Component({
  components: {
    TaskCard
  }
})
export default class SignBoard extends Vue {}
</script>
