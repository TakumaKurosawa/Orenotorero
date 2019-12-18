<template>
  <div>
    <h2>ボードリスト</h2>
    <LabelIcon
      :text-label="'パーソナルボード'"
      :icon-name="'mdi-account-outline'"
    ></LabelIcon>
    <Card
      v-for="item in cardTitles"
      :key="item.id"
      :card-title="item.title"
      :card-height="100"
      :card-width="180"
      :card-img="'https://cdn.vuetifyjs.com/images/cards/docks.jpg'"
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
  cardTitles!: Array<string>
}
</script>
