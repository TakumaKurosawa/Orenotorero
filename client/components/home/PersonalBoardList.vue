<template>
  <div>
    <h2>ボードリスト</h2>
    <LabelIcon
      :text-label="'パーソナルボード'"
      :icon-name="'mdi-account-outline'"
    ></LabelIcon>
    <Card
      v-for="item in boardData"
      :key="item.id"
      :card-title="item.name"
      :card-height="100"
      :card-width="180"
      :card-img="item.img"
    ></Card>
  </div>
</template>
<script lang="ts">
import { Vue, Prop, Component } from 'nuxt-property-decorator'
import LabelIcon from '@/components/atom/LabelIcon.vue'
import Card from '@/components/atom/Card.vue'

@Component({
  components: {
    LabelIcon,
    Card
  }
})
export default class BoardList extends Vue {
  @Prop({ type: Array, required: true })
  boardData!: Array<string>
}
</script>
