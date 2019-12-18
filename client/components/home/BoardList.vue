<template>
  <div>
    <h2>ボードリスト</h2>
    <LabelIcon :text-label="'最近の表示'" :icon-name="'mdi-clock'"></LabelIcon>
    <Card
      v-for="item in cardCount"
      :key="item.id"
      :card-title="item.title"
      :card-height="100"
      :card-width="180"
      :card-img="'https://cdn.vuetifyjs.com/images/cards/docks.jpg'"
    ></Card>
    <LabelIcon
      :text-label="'パーソナルボード'"
      :icon-name="'mdi-account-outline'"
    ></LabelIcon>
    <Card
      v-for="item in titles"
      :key="item.id"
      :card-title="item.title"
      :card-height="100"
      :card-width="180"
      :card-img="'https://cdn.vuetifyjs.com/images/cards/docks.jpg'"
    ></Card>
  </div>
</template>
<script lang="ts">
import { Vue, Component } from 'nuxt-property-decorator'
import LabelIcon from '@/components/atom/LabelIcon.vue'
import Card from '@/components/atom/Card.vue'
@Component({
  components: {
    LabelIcon,
    Card
  }
})
export default class BoardList extends Vue {
  titles = [
    { title: 'test1' },
    { title: 'test2' },
    { title: 'test3' },
    { title: 'test4' },
    { title: 'test5' },
    { title: 'test6' },
    { title: 'test7' },
    { title: 'test8' }
  ]

  get cardCount() {
    return this.titles.slice(0, 4)
  }
}
</script>
