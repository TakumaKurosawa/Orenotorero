<template>
  <div>
    <h2>ボード一覧</h2>
    <!-- <LabelIcon :text-label="'test'" :icon-name="'mdi-delete'"></LabelIcon> -->
  </div>
</template>

<script lang="ts">
import { Vue, Component } from 'nuxt-property-decorator'
import LabelIcon from '@/components/atom/LabelIcon.vue'

@Component({
  components: {
    LabelIcon
  }
})
export default class BoardList extends Vue {}
</script>
